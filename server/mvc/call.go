package mvc

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	dbapi_v1 "github.com/ntons/libra-go/api/dbapi/v1"
	sdk_v1 "github.com/ntons/libra-go/api/sdk/v1"
	"github.com/ntons/libra-go/server/sdk"
)

// anypb.New with panic
func newAny(msg proto.Message) *anypb.Any {
	any, err := anypb.New(msg)
	if err != nil {
		panic(err)
	}
	return any
}

type Call interface {
	sdk.Call
	GetModel(
		ctx context.Context,
		id string, opts ...GetOption) (proto.Message, error)
	LoadModel(
		ctx context.Context,
		id string, model proto.Message, opts ...GetOption) error
}

type modelContext struct {
	model  proto.Message // model message
	token  string        // db token
	origin *anypb.Any    // original model snapshot
}

type call struct {
	sdk.Call
	loaded map[string]modelContext
}

func FromContext(ctx context.Context) (Call, bool) {
	if sdkCall, ok := sdk.FromContext(ctx); ok {
		x := &call{
			Call:   sdkCall,
			loaded: make(map[string]modelContext),
		}
		sdkCall.OnReply(x.commit)
		return x, true
	}
	return nil, false
}

type getOptions struct {
	addIfNotFound proto.Message
}

type GetOption interface {
	apply(o *getOptions)
}

type funcGetOption struct {
	fn func(o *getOptions)
}

func (x funcGetOption) apply(o *getOptions) {
	x.fn(o)
}

func WithAddIfNotFound(model proto.Message) GetOption {
	if model == nil {
		panic(fmt.Errorf("model must not be nil"))
	}
	return funcGetOption{func(o *getOptions) { o.addIfNotFound = model }}
}

func (call *call) GetModel(
	ctx context.Context,
	id string, opts ...GetOption) (model proto.Message, err error) {
	if e, ok := call.loaded[id]; ok {
		return e.model, nil
	}
	origin, token, err := call.getModel(ctx, id, opts...)
	if err != nil {
		return
	}
	if model, err = origin.UnmarshalNew(); err != nil {
		return
	}
	call.loaded[id] = modelContext{
		model:  model,
		token:  token,
		origin: origin,
	}
	return
}

func (call *call) LoadModel(
	ctx context.Context, id string, model proto.Message, opts ...GetOption) (
	err error) {
	if _, ok := call.loaded[id]; ok {
		return fmt.Errorf("model %s already loaded", id)
	}
	origin, token, err := call.getModel(ctx, id, opts...)
	if err != nil {
		return
	}
	if err = origin.UnmarshalTo(model); err != nil {
		return
	}
	call.loaded[id] = modelContext{
		model:  model,
		token:  token,
		origin: origin,
	}
	return
}

func (call *call) getModel(
	ctx context.Context, id string, opts ...GetOption) (
	_ *anypb.Any, token string, err error) {
	var o getOptions
	for _, opt := range opts {
		opt.apply(&o)
	}
	req := &dbapi_v1.ArchiveGetRequest{Id: id, WithLock: true}
	if o.addIfNotFound != nil {
		req.AddIfNotFound = &dbapi_v1.Archive{
			Id:    id,
			Model: newAny(o.addIfNotFound),
		}
	}
	resp, err := call.GetDBV1().GetModel(ctx, req)
	if err != nil {
		return
	}
	return resp.Archive.Model, resp.Token, nil
}

func (call *call) commit(ctx context.Context, handleErr error) (firstErr error) {
	firstErr = handleErr
	for id, e := range call.loaded {
		if err := func() (err error) {
			req := &dbapi_v1.ArchiveSetRequest{
				Archive:    &dbapi_v1.Archive{Id: id},
				Token:      e.token,
				WithUnlock: true,
			}
			if model := newAny(e.model); proto.Equal(model, e.origin) {
				if _, err = call.GetDBV1().Unlock(ctx, req); err != nil {
					return
				}
			} else {
				req.Archive.Model = model
				if _, err = call.GetDBV1().SetModel(ctx, req); err != nil {
					return
				}
				notice := &sdk_v1.UpdateModelNotice{Model: model}
				if err = call.PushTo(ctx, id, notice); err != nil {
					return
				}
			}
			return
		}(); err != nil {
			if firstErr == nil {
				firstErr = err
			}
		}
	}
	return
}
