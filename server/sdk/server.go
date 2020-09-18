package sdk

import (
	"context"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ntons/log-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	xUserId = "x-libra-user-id"
	xRoleId = "x-libra-role-id"
)

type Server struct {
	// implement health service
	health *health.Server
	// connection to libra access point
	conn grpc.ClientConnInterface
	// quit routine
	ctx    context.Context
	cancel context.CancelFunc
	// registered svcs
	svcs []Service
}

func NewServer() (srv *Server) {
	srv = &Server{
		health: health.NewServer(),
	}
	srv.ctx, srv.cancel = context.WithCancel(context.Background())
	return
}

func (srv *Server) RegisterService(svc Service) {
	srv.svcs = append(srv.svcs, svc)
}

// dial to libra access point
func (srv *Server) Dial(addr string, opts ...DialOption) (err error) {
	var o = &dialOptions{}
	for _, opt := range opts {
		opt.apply(o)
	}
	if srv.conn, err = grpc.Dial(
		addr,
		append([]grpc.DialOption{
			grpc.WithChainUnaryInterceptor(srv.interceptClient),
		}, o.dialOptions...)...,
	); err != nil {
		return
	}
	return
}

func (srv *Server) ListenAndServe(addr string, opts ...ServeOption) (err error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	defer lis.Close()
	return srv.Serve(lis)
}

func (srv *Server) Serve(lis net.Listener, opts ...ServeOption) (err error) {
	var o = &serveOptions{}
	for _, opt := range opts {
		opt.apply(o)
	}

	gs := grpc.NewServer(append([]grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime: 30 * time.Second,
			},
		),
		grpc.ChainUnaryInterceptor(srv.interceptServer),
	}, o.serverOptions...)...)
	defer gs.GracefulStop()

	for _, svc := range srv.svcs {
		svc.Register(gs, srv.conn)
		defer svc.Stop()
	}

	healthpb.RegisterHealthServer(gs, srv.health)
	defer srv.health.Shutdown()

	go gs.Serve(lis)

	<-srv.ctx.Done()
	return
}

// shutdown server
func (srv *Server) Shutdown() {
	srv.cancel()
}

// block util term signal or shutdown
func (srv *Server) WaitForTerm() {
	sig := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGPIPE)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sig)
	select {
	case <-sig:
	case <-srv.ctx.Done():
	}
	return
}

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

func (srv *Server) interceptServer(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	if info.Server == srv.health {
		return handler(ctx, req)
	}
	log.Infof("intecept: %v", info)
	/*
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
	*/
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
	sdk := &sdk{
		ClientConnInterface: srv.conn,
		Recorder: log.With(log.M{
			"userId":    userId,
			"roleId":    roleId,
			"requestId": getRequestIdFromMetadata(md),
		}),
		appId:  appId,
		userId: userId,
		roleId: roleId,
	}
	log.Infof("sdk: %#v", sdk)
	out := metadata.MD{}
	for key, vals := range md {
		if strings.HasPrefix(key, "x-libra-") {
			out[key] = vals
		}
	}
	ctx = metadata.NewOutgoingContext(ctx, out)
	resp, err = handler(context.WithValue(ctx, sdkKey{}, sdk), req)
	if sdk.onReply != nil {
		err = sdk.onReply(ctx, err)
	}
	return
}

func (srv *Server) interceptClient(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (
	err error) {
	log.Infof("rpc call: %s", method)
	return invoker(ctx, method, req, reply, cc, opts...)
}
