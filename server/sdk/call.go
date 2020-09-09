package sdk

import (
	"context"

	dbapi_v1 "github.com/ntons/libra-go/api/dbapi/v1"
	gwapi_v1 "github.com/ntons/libra-go/api/gwapi/v1"
	log "github.com/ntons/log-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type callKey struct{}

func withCall(ctx context.Context, call Call) context.Context {
	return context.WithValue(ctx, callKey{}, call)
}
func FromContext(ctx context.Context) (call Call, ok bool) {
	call, ok = ctx.Value(callKey{}).(Call)
	return
}

type OnReplyFunc func(context.Context, error) error

// Call track every rpc call (in/out)
// provide session informations(userId, roleId)
// provide access api to libra services
type Call interface {
	log.Recorder
	// session values
	GetAppId() string
	GetUserId() string
	GetRoleId() string
	// sdk clients
	GetDBV1() dbapi_v1.DBClient
	GetGWV1() gwapi_v1.GatewayClient
	// push data to client
	OnReply(OnReplyFunc)
	PushTo(ctx context.Context, roleId string, msg proto.Message) error
}

type call struct {
	log.Recorder
	appId   string
	userId  string
	roleId  string
	dbv1    dbapi_v1.DBClient
	gwv1    gwapi_v1.GatewayClient
	onReply OnReplyFunc
}

func (call *call) GetAppId() string                { return call.appId }
func (call *call) GetUserId() string               { return call.userId }
func (call *call) GetRoleId() string               { return call.roleId }
func (call *call) GetDBV1() dbapi_v1.DBClient      { return call.dbv1 }
func (call *call) GetGWV1() gwapi_v1.GatewayClient { return call.gwv1 }

func (call *call) OnReply(fn OnReplyFunc) {
	call.onReply = fn
}
func (call *call) PushTo(
	ctx context.Context, roleId string, msg proto.Message) (err error) {
	req := &gwapi_v1.PushRequest{RoleId: roleId}
	if req.Data, err = anypb.New(msg); err != nil {
		return
	}
	if _, err = call.gwv1.Push(ctx, req); err != nil {
		return
	}
	return
}
