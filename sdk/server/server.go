package server

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/ntons/log-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	xRequestId             = "x-request-id"
	xLibraAppId            = "x-libra-app-id"
	xLibraAppSecret        = "x-libra-app-secret"
	xLibraRoleId           = "x-libra-role-id"
	xLibraTrustedAppId     = "x-libra-trusted-app-id"
	xLibraTrustedUserId    = "x-libra-trusted-user-id"
	xLibraTrustedRoleId    = "x-libra-trusted-role-id"
	xLibraTrustedRoleIndex = "x-libra-trusted-role-index"
)

type callKey struct{}

func FromIncomingContext(ctx context.Context) (call *Call, ok bool) {
	call, ok = ctx.Value(callKey{}).(*Call)
	return
}

type Call struct {
	log.Recorder
	// these fields can be truested
	AppId     string
	UserId    string
	RoleId    string
	RoleIndex int32
}

type Server struct {
	// grpc server, implements ServiceRegistrar
	s *grpc.Server
	grpc.ServiceRegistrar
	// life-time control
	ctx  context.Context
	stop context.CancelFunc
	// health status
	status *health.Server
}

func NewServer(opts ...grpc.ServerOption) *Server {
	srv := &Server{status: health.NewServer()}
	srv.ctx, srv.stop = context.WithCancel(context.Background())
	grpcsrv := grpc.NewServer(append(
		opts,
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime: 30 * time.Second,
			},
		),
		grpc.ChainUnaryInterceptor(srv.intercept),
	)...)
	srv.s, srv.ServiceRegistrar = grpcsrv, grpcsrv
	return srv
}

func (srv *Server) ListenAndServe(addr string) (err error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	defer lis.Close()
	return srv.Serve(lis)
}

func (srv *Server) Serve(lis net.Listener) (err error) {
	grpc_health_v1.RegisterHealthServer(srv, srv.status)
	defer srv.status.Shutdown()
	go srv.s.Serve(lis)
	defer srv.s.GracefulStop()
	<-srv.Done()
	return
}

func (srv *Server) Stop() { srv.stop() }

func (srv *Server) Done() <-chan struct{} { return srv.ctx.Done() }

func (srv *Server) intercept(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	if info.Server == srv.status {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "metadata required")
	}
	log.Debugw("call-in", "method", info.FullMethod, "metadata", md)
	fromMD := func(key string) string {
		if values := md.Get(key); len(values) > 0 {
			return values[0]
		}
		return ""
	}
	appId := fromMD(xLibraTrustedAppId)
	userId := fromMD(xLibraTrustedUserId)
	roleId := fromMD(xLibraTrustedRoleId)
	if appId == "" || userId == "" || roleId == "" {
		return nil, status.Errorf(
			codes.PermissionDenied, "require trusted session")
	}
	roleIndex, err := strconv.ParseInt(fromMD(xLibraTrustedRoleIndex), 10, 32)
	if err != nil {
		return nil, status.Errorf(
			codes.PermissionDenied, "require trusted session")
	}
	call := &Call{
		Recorder: log.With(log.M{
			"app":     appId,
			"user":    userId,
			"role":    roleId,
			"request": fromMD(xRequestId),
		}),
		AppId:     appId,
		UserId:    userId,
		RoleId:    roleId,
		RoleIndex: int32(roleIndex),
	}
	defer func() {
		if r := recover(); r != nil {
			call.Errorf("recover from %v", r)
			var ok bool
			if err, ok = r.(error); !ok {
				err = status.Errorf(codes.Aborted, "%s", r)
			}
		}
	}()
	resp, err = handler(context.WithValue(ctx, callKey{}, call), req)
	return
}
