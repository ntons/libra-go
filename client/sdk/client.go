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

const (
	xLibraToken  = "x-libra-token"
	xLibraTicket = "x-libra-ticket"
)

type Client interface {
	grpc.ClientConnInterface
	Close()
	Recv(ctx context.Context) (proto.Message, error)
	Login(ctx context.Context, state proto.Message) (*v1pb.UserData, error)
	ListRoles(ctx context.Context) ([]*v1pb.RoleData, error)
	CreateRole(ctx context.Context, index uint32) (*v1pb.RoleData, error)
	SignIn(ctx context.Context, roleId string) error
	Connect(ctx context.Context) error
}

type client struct {
	grpc.ClientConnInterface

	appId  string
	userId string // assigned by login
	roleId string // assigned by sign in

	// session
	token  string
	ticket string

	msgChan chan proto.Message

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	gwapi   v1pb.GatewayClient
	userapi v1pb.UserClient
	roleapi v1pb.RoleClient
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
	cli.gwapi = v1pb.NewGatewayClient(cli)
	cli.userapi = v1pb.NewUserClient(cli)
	cli.roleapi = v1pb.NewRoleClient(cli)
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
	ctx context.Context, state proto.Message) (user *v1pb.UserData, err error) {
	req := &v1pb.UserLoginRequest{AppId: cli.appId}
	if req.State, err = anypb.New(state); err != nil {
		return
	}
	var header metadata.MD
	resp, err := cli.userapi.Login(ctx, req, grpc.Header(&header))
	if err != nil {
		return
	}
	if v := header.Get(xLibraToken); len(v) != 1 {
		return nil, fmt.Errorf("miss token in response")
	} else {
		cli.token = v[0]
	}
	cli.userId = resp.User.Id
	return resp.User, nil
}

func (cli *client) ListRoles(
	ctx context.Context) (roles []*v1pb.RoleData, err error) {
	req := &v1pb.RoleListRequest{}
	resp, err := cli.roleapi.List(ctx, req)
	if err != nil {
		return
	}
	return resp.Roles, nil
}

func (cli *client) CreateRole(
	ctx context.Context, index uint32) (role *v1pb.RoleData, err error) {
	req := &v1pb.RoleCreateRequest{Index: index}
	resp, err := cli.roleapi.Create(ctx, req)
	if err != nil {
		return
	}
	role = resp.Role
	return
}

func (cli *client) SignIn(ctx context.Context, roleId string) (err error) {
	req := &v1pb.RoleSignInRequest{RoleId: roleId}
	var header metadata.MD
	if _, err = cli.roleapi.SignIn(ctx, req, grpc.Header(&header)); err != nil {
		return
	}
	if v := header.Get(xLibraTicket); len(v) != 1 {
		return fmt.Errorf("miss ticket in response")
	} else {
		cli.ticket = v[0]
	}
	cli.roleId = roleId
	return
}
func (cli *client) Connect(ctx context.Context) (err error) {
	req := &v1pb.GatewayConnectRequest{}
	stream, err := cli.gwapi.Connect(ctx, req)
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
	if cli.token != "" {
		md.Set(xLibraToken, cli.token)
	}
	if cli.ticket != "" {
		md.Set(xLibraTicket, cli.ticket)
	}
	return metadata.NewOutgoingContext(ctx, md)
}
