package sdk

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/ghodss/yaml"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	dbapi_v1 "github.com/ntons/libra-go/api/dbapi/v1"
	gwapi_v1 "github.com/ntons/libra-go/api/gwapi/v1"
	sdk_v1 "github.com/ntons/libra-go/api/sdk/v1"
	"github.com/ntons/log-go"
	"github.com/ntons/tongo/template"
)

const (
	xUserId = "x-libra-user-id"
	xRoleId = "x-libra-role-id"
)

type Server struct {
	// grpc server
	gRPCServer *grpc.Server
	// grpc health service
	health *healthServer
	// quit routine
	ctx    context.Context
	cancel context.CancelFunc
	// service listener
	listener net.Listener
	// api
	dbv1 dbapi_v1.DBClient
	gwv1 gwapi_v1.GatewayClient
}

func NewServer() (s *Server) {
	s = &Server{}
	s.gRPCServer = grpc.NewServer(grpc.UnaryInterceptor(s.intercept))
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.health = newHealth(s.ctx)
	grpc_health_v1.RegisterHealthServer(s.gRPCServer, s.health)
	return
}

func (s *Server) initLog(cfg *sdk_v1.ServerConfig) (err error) {
	zcfg := zap.NewProductionConfig()
	zcfg.Sampling = nil
	logger, err := zcfg.Build(zap.AddCaller())
	if err != nil {
		return
	}
	log.SetZapLogger(logger)
	return
}

func (s *Server) initApi(cfg *sdk_v1.ServerConfig) (err error) {
	var conn *grpc.ClientConn
	if conn, err = grpc.Dial(
		cfg.AccessPoint.Address,
		grpc.WithAuthority(cfg.AccessPoint.Authority),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(s.dbIntercept),
	); err != nil {
		log.Errorf("failed to dail to db: %v")
		return
	}
	s.dbv1 = dbapi_v1.NewDBClient(conn)

	if conn, err = grpc.Dial(
		cfg.AccessPoint.Address,
		grpc.WithAuthority(cfg.AccessPoint.Authority),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(s.gwIntercept),
	); err != nil {
		log.Errorf("failed to dail to gw: %v")
		return
	}
	s.gwv1 = gwapi_v1.NewGatewayClient(conn)

	return
}

func readConfig(fp string) (b []byte, err error) {
	if b, err = template.RenderFile(fp, nil); err != nil {
		return
	}
	switch ext := filepath.Ext(fp); ext {
	case ".json":
		return
	case ".yml", ".yaml":
		return yaml.YAMLToJSON(b)
	default:
		return nil, fmt.Errorf("unknown file extension: %v", ext)
	}
}

func (s *Server) Init() (err error) {
	var fp string
	flag.StringVar(&fp, "c", "", "config file")
	flag.Parse()
	if fp == "" {
		return fmt.Errorf("failed to parse options")
	}
	cfg := &sdk_v1.ServerConfig{}
	if b, err := readConfig(fp); err != nil {
		return err
	} else if err = json.Unmarshal(b, cfg); err != nil {
		return err
	}
	if err = s.initLog(cfg); err != nil {
		return
	}
	if err = s.initApi(cfg); err != nil {
		return
	}
	if s.listener, err = net.Listen("tcp", cfg.Bind); err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	log.Infof("listen on %s", cfg.Bind)
	return
}

func (s *Server) Serve() (err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = s.gRPCServer.Serve(s.listener); err != nil {
			return
		}
	}()
	defer s.gRPCServer.GracefulStop()

	sig := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGPIPE)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sig)
	select {
	case <-sig:
		s.Stop()
	case <-s.ctx.Done():
	}
	return
}

func (s *Server) Stop() { s.cancel() }

// SomeService.RegisterSomeServer(s.GRPC(), SomeImplement)
func (s *Server) GRPCServer() *grpc.Server { return s.gRPCServer }

func getValueFromMetadata(md metadata.MD, key string) string {
	if values := md.Get(key); len(values) > 0 {
		return values[0]
	}
	return ""
}

// remove hyphen from uuid
func getRequestIdFromMetadata(md metadata.MD) string {
	return strings.ReplaceAll(
		getValueFromMetadata(md, "x-request-id"), "-", "")
}

func (s *Server) intercept(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		// business layer panic check
		if r := recover(); r != nil {
			log.Errorf("recover: %v", r)
			var ok bool
			if err, ok = r.(error); !ok {
				err = status.Errorf(codes.Aborted, "%s", r)
			}
		}
	}()
	if info.Server == s.health {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "no metadata")
	}
	// retrieve session values from md
	appId := getValueFromMetadata(md, "x-libra-app-id")
	userId := getValueFromMetadata(md, "x-libra-user-id")
	roleId := getValueFromMetadata(md, "x-libra-role-id")
	if appId == "" || userId == "" || roleId == "" {
		return nil, status.Errorf(codes.PermissionDenied, "no session")
	}
	rec := log.With(log.M{
		"userId":    userId,
		"roleId":    roleId,
		"requestId": getRequestIdFromMetadata(md),
	})
	call := &call{
		Recorder: rec,
		appId:    appId,
		userId:   userId,
		roleId:   roleId,
		dbv1:     s.dbv1,
		gwv1:     s.gwv1,
	}
	out := metadata.MD{}
	for key, vals := range md {
		if strings.HasPrefix(key, "x-libra-") {
			out[key] = vals
		}
	}
	ctx = metadata.NewOutgoingContext(ctx, out)
	resp, err = handler(withCall(ctx, call), req)
	if call.onReply != nil {
		err = call.onReply(ctx, err)
	}
	return
}

func (s *Server) dbIntercept(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (
	err error) {
	log.Infof("invoke db: %s", method)
	return invoker(ctx, method, req, reply, cc, opts...)
}
func (s *Server) gwIntercept(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (
	err error) {
	log.Infof("invoke gw: %s", method)
	return invoker(ctx, method, req, reply, cc, opts...)
}
