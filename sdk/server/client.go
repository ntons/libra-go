package server

import (
	"context"
	"sync"

	"github.com/ntons/grpc-compressor/lz4"
	v1pb "github.com/ntons/libra-go/api/libra/v1"
	"github.com/ntons/log-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func WithMD(ctx context.Context, kv ...string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(kv...))
}
func WithAuth(ctx context.Context, appId, appSecret string) context.Context {
	return WithMD(ctx, xLibraAppId, appId, xLibraAppSecret, appSecret)
}

type Client struct {
	Gateway     v1pb.GatewayClient
	Database    v1pb.DatabaseClient
	Mailbox     v1pb.MailboxClient
	Distlock    v1pb.DistlockClient
	Leaderboard v1pb.LeaderboardClient
	BubbleChart v1pb.BubbleChartClient
	// appId -> appSecret
	authMu  sync.Mutex
	authMap map[string]string
}

func Dial(addr string, opts ...grpc.DialOption) (_ *Client, err error) {
	cli := &Client{
		authMap: make(map[string]string),
	}
	conn, err := grpc.Dial(addr, append(
		opts,
		grpc.WithChainUnaryInterceptor(cli.intercept),
		// enable lz4 compression by default
		grpc.WithDefaultCallOptions(grpc.UseCompressor(lz4.Name)),
	)...)
	if err != nil {
		return
	}
	cli.Gateway = v1pb.NewGatewayClient(conn)
	cli.Database = v1pb.NewDatabaseClient(conn)
	cli.Mailbox = v1pb.NewMailboxClient(conn)
	cli.Leaderboard = v1pb.NewLeaderboardClient(conn)
	cli.BubbleChart = v1pb.NewBubbleChartClient(conn)
	return cli, nil
}

func (cli *Client) AddAuth(appId, appSecret string) {
	cli.authMu.Lock()
	defer cli.authMu.Unlock()
	cli.authMap[appId] = appSecret
}

func (cli *Client) findAppSecret(appId string) (appSecret string, ok bool) {
	cli.authMu.Lock()
	defer cli.authMu.Unlock()
	appSecret, ok = cli.authMap[appId]
	return
}

func (cli *Client) intercept(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) (err error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok || len(md[xLibraAppId]) == 0 || len(md[xLibraAppSecret]) == 0 {
		// no app auth metadata assigned
		// try to match app id and secret by incoming metadata
		md, ok = metadata.FromIncomingContext(ctx)
		if !ok || len(md[xLibraTrustedAppId]) == 0 {
			return status.Errorf(codes.Unauthenticated, "")
		}
		appId := md[xLibraTrustedAppId][0]
		appSecret, ok := cli.findAppSecret(appId)
		if !ok {
			return status.Errorf(codes.Unauthenticated, "")
		}
		ctx = WithAuth(ctx, appId, appSecret)
		md, _ = metadata.FromOutgoingContext(ctx)
	}
	log.Debugw("call-out", "method", method, "metadata", md)
	return invoker(ctx, method, req, reply, cc, opts...)
}
