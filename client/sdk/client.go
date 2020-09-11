package sdk

import (
	"context"
	"fmt"
	"io"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	gwapi_v1 "github.com/ntons/libra-go/api/gwapi/v1"
	ptapi_v1 "github.com/ntons/libra-go/api/ptapi/v1"
	sdk_v1 "github.com/ntons/libra-go/api/sdk/v1"
)

type Client struct {
	*grpc.ClientConn

	appId  string
	userId string // assigned after login
	roleId string // assigned after sigi in

	// session
	token  string
	ticket string

	msgChan chan proto.Message

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	gwv1 gwapi_v1.AccessClient
	ptv1 ptapi_v1.AccountClient
}

func Dial(appId string, ap *sdk_v1.Endpoint, opts ...DialOption) (_ *Client, err error) {
	var o dialOptions
	for _, opt := range opts {
		opt.apply(&o)
	}
	cli := &Client{
		appId:   appId,
		msgChan: make(chan proto.Message, 64),
	}
	conn, err := cli.dial(ap, o.opts)
	if err != nil {
		return
	}
	cli.ClientConn = conn
	if o.libraAccessPoint != nil {
		if conn, err = cli.dial(o.libraAccessPoint, o.opts); err != nil {
			return
		}
	}
	cli.gwv1 = gwapi_v1.NewAccessClient(conn)
	cli.ptv1 = ptapi_v1.NewAccountClient(conn)
	cli.ctx, cli.cancel = context.WithCancel(context.Background())
	return cli, nil
}

// dial to endpoint if endpoint is not nil, otherwise conn will be returned
func (cli *Client) dial(
	ap *sdk_v1.Endpoint, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(ap.Address, append(
		opts,
		grpc.WithAuthority(ap.Authority),
		grpc.WithUnaryInterceptor(cli.unaryInterceptor),
		grpc.WithStreamInterceptor(cli.streamInterceptor),
	)...)
}

func (cli *Client) Close() {
	cli.cancel()
	cli.wg.Wait()
}

func (cli *Client) Recv(ctx context.Context) (msg proto.Message, err error) {
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case <-cli.ctx.Done():
		err = cli.ctx.Err()
	case msg = <-cli.msgChan:
	}
	return
}

func (cli *Client) Login(
	ctx context.Context, state proto.Message) (user *ptapi_v1.User, err error) {
	req := &ptapi_v1.LoginRequest{
		AppId: cli.appId,
	}
	if req.State, err = anypb.New(state); err != nil {
		return
	}
	resp, err := cli.ptv1.Login(ctx, req)
	if err != nil {
		return
	}
	cli.token = resp.Token
	cli.userId = resp.User.Id
	return resp.User, nil
}

func (cli *Client) ListRoles(
	ctx context.Context) (roles []*ptapi_v1.Role, err error) {
	req := &ptapi_v1.ListRolesRequest{
		AppId: cli.appId,
		Token: cli.token,
	}
	resp, err := cli.ptv1.ListRoles(ctx, req)
	if err != nil {
		return
	}
	roles = resp.Roles
	return
}

func (cli *Client) CreateRole(
	ctx context.Context, index int32) (role *ptapi_v1.Role, err error) {
	req := &ptapi_v1.CreateRoleRequest{
		AppId: cli.appId,
		Token: cli.token,
		Index: index,
	}
	resp, err := cli.ptv1.CreateRole(ctx, req)
	if err != nil {
		return
	}
	role = resp.Role
	return
}

func (cli *Client) SignIn(ctx context.Context, roleId string) (err error) {
	if err = cli.ptSignIn(ctx, roleId); err != nil {
		return
	}
	if err = cli.gwSignIn(ctx); err != nil {
		return
	}
	return
}
func (cli *Client) ptSignIn(ctx context.Context, roleId string) (err error) {
	req := &ptapi_v1.SignInRequest{
		AppId:  cli.appId,
		Token:  cli.token,
		RoleId: roleId,
	}
	var resp *ptapi_v1.SignInResponse
	if resp, err = cli.ptv1.SignIn(ctx, req); err != nil {
		return
	}
	cli.ticket = resp.Ticket
	cli.roleId = roleId
	return
}
func (cli *Client) gwSignIn(ctx context.Context) (err error) {
	req := &gwapi_v1.SignInRequest{}
	stream, err := cli.gwv1.SignIn(ctx, req)
	if err != nil {
		return
	}
	if _, err = stream.Recv(); err != nil {
		return
	}
	cli.wg.Add(1)
	go func() {
		defer cli.wg.Done()
		for {
			// RecvMsg blocks until it receives a message into m or the
			// stream is done. It returns io.EOF when the stream completes
			// successfully. On any other error, the stream is aborted
			// and the error contains the RPC status.
			any, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					fmt.Println("failed to receive: ", err)
				}
				return
			}
			msg, err := any.UnmarshalNew()
			if err != nil {
				fmt.Println("failed to unmarshal: ", err)
				continue
			}
			select {
			case <-cli.ctx.Done():
				return
			case cli.msgChan <- msg:
			}
		}
	}()
	return
}

func (cli *Client) unaryInterceptor(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	return invoker(cli.addMd(ctx, method), method, req, reply, cc, opts...)
}
func (cli *Client) streamInterceptor(
	ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
	method string, streamer grpc.Streamer, opts ...grpc.CallOption) (
	grpc.ClientStream, error) {
	return streamer(cli.addMd(ctx, method), desc, cc, method, opts...)
}
func (cli *Client) addMd(
	ctx context.Context, method string) context.Context {
	md := metadata.MD{}
	md.Set("x-libra-app-id", cli.appId)
	if cli.userId != "" {
		md.Set("x-libra-user-id", cli.userId)
	}
	if cli.roleId != "" {
		md.Set("x-libra-role-id", cli.roleId)
	}
	if cli.token != "" {
		md.Set("x-libra-token", cli.token)
	}
	if cli.ticket != "" {
		md.Set("x-libra-ticket", cli.ticket)
	}
	return metadata.NewOutgoingContext(ctx, md)
}
