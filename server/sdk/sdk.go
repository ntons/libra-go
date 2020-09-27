package sdk

import (
	"context"

	v1pb "github.com/ntons/libra-go/api/v1"
	log "github.com/ntons/log-go"
	"google.golang.org/grpc"
	pb "google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

type sdkKey struct{}

func FromContext(ctx context.Context) (sdk SDK, ok bool) {
	sdk, ok = ctx.Value(sdkKey{}).(SDK)
	return
}

type OnReplyFunc func(context.Context, error) error

type SDK interface {
	// connection to libra services
	grpc.ClientConnInterface
	// logger
	log.Recorder
	// session values
	GetAppId() string
	GetUserId() string
	GetRoleId() string
	// hook before reply
	OnReply(OnReplyFunc)
	// push data to client
	PushTo(ctx context.Context, roleId string, msg pb.Message) error
}

type sdk struct {
	grpc.ClientConnInterface
	log.Recorder
	appId   string
	userId  string
	roleId  string
	onReply OnReplyFunc
}

func (x *sdk) GetAppId() string  { return x.appId }
func (x *sdk) GetUserId() string { return x.userId }
func (x *sdk) GetRoleId() string { return x.roleId }

func (x *sdk) OnReply(callback OnReplyFunc) { x.onReply = callback }

func (x *sdk) PushTo(
	ctx context.Context, roleId string, msg pb.Message) (err error) {
	req := &v1pb.GatewayPushRequest{RoleId: roleId}
	if req.Data, err = anypb.New(msg); err != nil {
		return
	}
	if _, err = v1pb.NewGatewayClient(x).Push(ctx, req); err != nil {
		return
	}
	return
}
