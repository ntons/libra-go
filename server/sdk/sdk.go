package sdk

import (
	"context"

	gwpb "github.com/ntons/libra-go/api/gw/v1"
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
	log.Recorder
	// session values
	GetAppId() string
	GetUserId() string
	GetRoleId() string
	// push data to client
	OnReply(OnReplyFunc)
	PushTo(ctx context.Context, roleId string, msg pb.Message) error
}

type sdk struct {
	log.Recorder
	conn    grpc.ClientConnInterface
	appId   string
	userId  string
	roleId  string
	onReply OnReplyFunc
}

func (x *sdk) GetAppId() string       { return x.appId }
func (x *sdk) GetUserId() string      { return x.userId }
func (x *sdk) GetRoleId() string      { return x.roleId }
func (x *sdk) OnReply(fn OnReplyFunc) { x.onReply = fn }

func (x *sdk) PushTo(
	ctx context.Context, roleId string, msg pb.Message) (err error) {
	req := &gwpb.PushRequest{RoleId: roleId}
	if req.Data, err = anypb.New(msg); err != nil {
		return
	}
	if _, err = gwpb.NewGatewayClient(x.conn).Push(ctx, req); err != nil {
		return
	}
	return
}
