package server

import (
	"context"

	v1pb "github.com/ntons/libra-go/api/v1"
	log "github.com/ntons/log-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	pb "google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

type callKey struct{}

func withCall(ctx context.Context, call *Call) context.Context {
	return context.WithValue(ctx, callKey{}, call)
}
func CallFromContext(ctx context.Context) (call *Call, ok bool) {
	call, ok = ctx.Value(callKey{}).(*Call)
	return
}

type Call struct {
	*libraClients
	log.Recorder

	appId  string
	userId string
	roleId string

	cc grpc.ClientConnInterface
}

func (c *Call) AppId() string  { return c.appId }
func (c *Call) UserId() string { return c.userId }
func (c *Call) RoleId() string { return c.roleId }

func (c *Call) SendTo(
	ctx context.Context, roleId string, msg pb.Message) (err error) {
	if c.Gateway == nil {
		return status.Errorf(codes.Internal, "not connect to libra")
	}
	req := &v1pb.GatewaySendRequest{}
	if req.Data, err = anypb.New(msg); err != nil {
		return
	}
	if err = grpc.SetHeader(
		ctx, metadata.Pairs(xLibraRoleId, roleId)); err != nil {
		return
	}
	if _, err = c.Gateway.Send(ctx, req); err != nil {
		return
	}
	return
}
