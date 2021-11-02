package client

import (
	"context"
	"fmt"
	"io"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	v1pb "github.com/ntons/libra-go/api/libra/v1"
)

const (
	xLibraToken  = "x-libra-token"
	xLibraTicket = "x-libra-ticket"
)

type Client struct {
	grpc.ClientConnInterface

	// libra clients
	user    v1pb.UserClient
	role    v1pb.RoleClient
	gateway v1pb.GatewayClient

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
}

// dail to libra edge proxy
func Dial(
	appId string, addr string, opts ...grpc.DialOption) (_ *Client, err error) {
	cli := &Client{
		appId:   appId,
		msgChan: make(chan proto.Message, 64),
	}
	if cli.ClientConnInterface, err = grpc.Dial(
		addr,
		append(
			opts,
			grpc.WithUnaryInterceptor(cli.interceptUnary),
			grpc.WithStreamInterceptor(cli.interceptStream),
		)...,
	); err != nil {
		return
	}
	cli.user = v1pb.NewUserClient(cli)
	cli.role = v1pb.NewRoleClient(cli)
	cli.gateway = v1pb.NewGatewayClient(cli)
	cli.ctx, cli.cancel = context.WithCancel(context.Background())
	return cli, nil
}

func (cli *Client) Close() {
	cli.cancel()
	cli.wg.Wait()
}

func (cli *Client) Login(
	ctx context.Context, state proto.Message) (user *v1pb.UserData, err error) {
	req := &v1pb.UserLoginRequest{AppId: cli.appId}
	if req.State, err = anypb.New(state); err != nil {
		return
	}
	var header metadata.MD
	resp, err := cli.user.Login(ctx, req, grpc.Header(&header))
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

func (cli *Client) ListRoles(
	ctx context.Context) (roles []*v1pb.RoleData, err error) {
	req := &v1pb.RoleListRequest{}
	resp, err := cli.role.List(ctx, req)
	if err != nil {
		return
	}
	return resp.Roles, nil
}

func (cli *Client) CreateRole(
	ctx context.Context, index uint32) (role *v1pb.RoleData, err error) {
	req := &v1pb.RoleCreateRequest{Index: index}
	resp, err := cli.role.Create(ctx, req)
	if err != nil {
		return
	}
	role = resp.Role
	return
}

func (cli *Client) SignIn(ctx context.Context, roleId string) (err error) {
	req := &v1pb.RoleSignInRequest{RoleId: roleId}
	var header metadata.MD
	if _, err = cli.role.SignIn(ctx, req, grpc.Header(&header)); err != nil {
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

func (cli *Client) Connect(ctx context.Context) (err error) {
	req := &v1pb.GatewayConnectRequest{}
	stream, err := cli.gateway.Connect(ctx, req)
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

func (cli *Client) interceptUnary(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	return invoker(cli.withState(ctx), method, req, reply, cc, opts...)
}
func (cli *Client) interceptStream(
	ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
	method string, streamer grpc.Streamer, opts ...grpc.CallOption) (
	grpc.ClientStream, error) {
	return streamer(cli.withState(ctx), desc, cc, method, opts...)
}
func (cli *Client) withState(ctx context.Context) context.Context {
	md := metadata.MD{}
	if cli.token != "" {
		md.Set(xLibraToken, cli.token)
	}
	if cli.ticket != "" {
		md.Set(xLibraTicket, cli.ticket)
	}
	return metadata.NewOutgoingContext(ctx, md)
}
