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

	v1pb "github.com/ntons/libra-go/api/v1"
)

type Client interface {
	grpc.ClientConnInterface
	Close()
	Recv(ctx context.Context) (proto.Message, error)
	Login(ctx context.Context, state proto.Message) (*v1pb.User, error)
	ListRoles(ctx context.Context) ([]*v1pb.Role, error)
	CreateRole(ctx context.Context, index int32) (*v1pb.Role, error)
	SignIn(ctx context.Context, roleId string) error
}

type client struct {
	grpc.ClientConnInterface

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

	gwapi v1pb.AccessClient
	ptapi v1pb.AccountClient
}

func Dial(appId string, addr string, opts ...DialOption) (_ Client, err error) {
	var o dialOptions
	for _, opt := range opts {
		opt.apply(&o)
	}
	cli := &client{
		appId:   appId,
		msgChan: make(chan proto.Message, 64),
	}
	if cli.ClientConnInterface, err = grpc.Dial(
		addr, append([]grpc.DialOption{
			grpc.WithUnaryInterceptor(cli.unaryInterceptor),
			grpc.WithStreamInterceptor(cli.streamInterceptor),
		}, o.dialOptions...)...,
	); err != nil {
		return
	}
	cli.gwapi = v1pb.NewAccessClient(cli)
	cli.ptapi = v1pb.NewAccountClient(cli)
	cli.ctx, cli.cancel = context.WithCancel(context.Background())
	return cli, nil
}

func (cli *client) Close() {
	cli.cancel()
	cli.wg.Wait()
}

func (cli *client) Recv(ctx context.Context) (msg proto.Message, err error) {
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case <-cli.ctx.Done():
		err = cli.ctx.Err()
	case msg = <-cli.msgChan:
	}
	return
}

func (cli *client) Login(
	ctx context.Context, state proto.Message) (user *v1pb.User, err error) {
	req := &v1pb.LoginRequest{
		AppId: cli.appId,
	}
	if req.State, err = anypb.New(state); err != nil {
		return
	}
	resp, err := cli.ptapi.Login(ctx, req)
	if err != nil {
		return
	}
	cli.token = resp.Token
	cli.userId = resp.User.Id
	return resp.User, nil
}

func (cli *client) ListRoles(
	ctx context.Context) (roles []*v1pb.Role, err error) {
	req := &v1pb.ListRolesRequest{
		AppId: cli.appId,
		Token: cli.token,
	}
	resp, err := cli.ptapi.ListRoles(ctx, req)
	if err != nil {
		return
	}
	roles = resp.Roles
	return
}

func (cli *client) CreateRole(
	ctx context.Context, index int32) (role *v1pb.Role, err error) {
	req := &v1pb.CreateRoleRequest{
		AppId: cli.appId,
		Token: cli.token,
		Index: index,
	}
	resp, err := cli.ptapi.CreateRole(ctx, req)
	if err != nil {
		return
	}
	role = resp.Role
	return
}

func (cli *client) SignIn(ctx context.Context, roleId string) (err error) {
	if err = cli.signIn(ctx, roleId); err != nil {
		return
	}
	if err = cli.access(ctx); err != nil {
		return
	}
	return
}
func (cli *client) signIn(ctx context.Context, roleId string) (err error) {
	req := &v1pb.SignInRequest{
		AppId:  cli.appId,
		Token:  cli.token,
		RoleId: roleId,
	}
	var resp *v1pb.SignInResponse
	if resp, err = cli.ptapi.SignIn(ctx, req); err != nil {
		return
	}
	cli.ticket = resp.Ticket
	cli.roleId = roleId
	return
}
func (cli *client) access(ctx context.Context) (err error) {
	req := &v1pb.AccessRequest{}
	stream, err := cli.gwapi.Access(ctx, req)
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

func (cli *client) unaryInterceptor(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	return invoker(cli.addMd(ctx, method), method, req, reply, cc, opts...)
}
func (cli *client) streamInterceptor(
	ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
	method string, streamer grpc.Streamer, opts ...grpc.CallOption) (
	grpc.ClientStream, error) {
	return streamer(cli.addMd(ctx, method), desc, cc, method, opts...)
}
func (cli *client) addMd(
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
