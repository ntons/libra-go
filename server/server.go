package server

import (
	"context"
	"net"
	"time"

	"github.com/ntons/log-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	v1pb "github.com/ntons/libra-go/api/v1"
)

const (
	xLibraAppId         = "x-libra-app-id"
	xLibraRoleId        = "x-libra-role-id"
	xLibraTrustedAppId  = "x-libra-trusted-app-id"
	xLibraTrustedUserId = "x-libra-trusted-user-id"
	xLibraTrustedRoleId = "x-libra-trusted-role-id"
)

type libraClients struct {
	Gateway     v1pb.GatewayClient
	Database    v1pb.DatabaseClient
	Mailbox     v1pb.MailboxClient
	Leaderboard v1pb.LeaderboardClient
	BubbleChart v1pb.BubbleChartClient
}

type Service interface {
	// register grpc service
	Register(s *grpc.Server) error
	// stop serving, clean resources
	Stop()
}

type Server struct {
	// grpc server, implements ServiceRegistrar
	s *grpc.Server
	grpc.ServiceRegistrar

	// libra api clients, only work after DialToLibra
	libraClients

	// life-time control
	ctx    context.Context
	cancel context.CancelFunc

	// health status
	status *health.Server
}

func NewServer(opts ...grpc.ServerOption) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	srv := &Server{ctx: ctx, cancel: cancel, status: health.NewServer()}

	grpcsrv := grpc.NewServer(append(
		opts,
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime: 30 * time.Second,
			},
		),
		grpc.ChainUnaryInterceptor(srv.interceptCallIn),
	)...)
	srv.s, srv.ServiceRegistrar = grpcsrv, grpcsrv
	return srv
}

func (srv *Server) DialToLibra(
	addr string, opts ...grpc.DialOption) (err error) {
	cc, err := grpc.Dial(addr, append(
		opts,
		grpc.WithChainUnaryInterceptor(srv.interceptCallOut),
	)...)
	if err != nil {
		return
	}
	srv.libraClients = libraClients{
		Gateway:     v1pb.NewGatewayClient(cc),
		Database:    v1pb.NewDatabaseClient(cc),
		Mailbox:     v1pb.NewMailboxClient(cc),
		Leaderboard: v1pb.NewLeaderboardClient(cc),
		BubbleChart: v1pb.NewBubbleChartClient(cc),
	}
	return
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

func (srv *Server) Shutdown() { srv.cancel() }

func (srv *Server) Done() <-chan struct{} { return srv.ctx.Done() }

func fetchMetadata(md metadata.MD, key string) string {
	if values := md.Get(key); len(values) > 0 {
		return values[0]
	}
	return ""
}

func (srv *Server) interceptCallIn(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	if info.Server == srv.status {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "metadata required")
	}
	log.Debugf("call-in metadata: %v", md)
	appId := fetchMetadata(md, xLibraTrustedAppId)
	userId := fetchMetadata(md, xLibraTrustedUserId)
	roleId := fetchMetadata(md, xLibraTrustedRoleId)
	if appId == "" || userId == "" || roleId == "" {
		return nil, status.Errorf(codes.PermissionDenied, "session required")
	}
	call := &Call{
		libraClients: &srv.libraClients,
		Recorder: log.With(log.M{
			"app":     appId,
			"user":    userId,
			"role":    roleId,
			"request": fetchMetadata(md, "x-request-id"),
		}),
		appId:  appId,
		userId: userId,
		roleId: roleId,
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
	resp, err = handler(withCall(ctx, call), req)
	return
}

func (srv *Server) interceptCallOut(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) (err error) {
	log.Debugf("call-out %s", method)
	call, ok := CallFromContext(ctx)
	if !ok { // call must exist
		return status.Errorf(codes.Internal, "")
	}
	// set outgoing metadata
	md := metadata.Pairs(xLibraAppId, call.AppId())
	return invoker(
		metadata.NewOutgoingContext(ctx, md),
		method, req, reply, cc, opts...)
}
