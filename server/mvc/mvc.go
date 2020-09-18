package mvc

import (
	"context"
	"fmt"

	log "github.com/ntons/log-go"
	pb "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	dbpb "github.com/ntons/libra-go/api/db/v1"
	sdkpb "github.com/ntons/libra-go/api/sdk/v1"
	"github.com/ntons/libra-go/server/sdk"
)

// anypb.New with panic
func newAny(msg pb.Message) *anypb.Any {
	any, err := anypb.New(msg)
	if err != nil {
		panic(err)
	}
	return any
}

type MVC interface {
	sdk.SDK
	// GetModel get model from db or cache
	GetModel(ctx context.Context, id string, opts ...GetOption) (pb.Message, error)
	// LoadModel is similar as GetModel, but unmarshal to a user defined message
	LoadModel(ctx context.Context, id string, model pb.Message, opts ...GetOption) error
}

// cache item
type item struct {
	model pb.Message // model message
	token string     // db token
	data  *anypb.Any // original model data
}

type mvc struct {
	sdk.SDK
	cache map[string]*item
	dbapi dbpb.DBClient
}

func FromContext(ctx context.Context) (MVC, bool) {
	if sdk, ok := sdk.FromContext(ctx); ok {
		x := &mvc{
			SDK:   sdk,
			cache: make(map[string]*item),
			dbapi: dbpb.NewDBClient(sdk),
		}
		sdk.OnReply(x.submit)
		return x, true
	}
	return nil, false
}

type getOptions struct {
	addIfNotFound pb.Message
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

func WithAddIfNotFound(model pb.Message) GetOption {
	if model == nil {
		panic(fmt.Errorf("model must not be nil"))
	}
	return funcGetOption{func(o *getOptions) { o.addIfNotFound = model }}
}

func (x *mvc) GetModel(ctx context.Context, id string, opts ...GetOption) (model pb.Message, err error) {
	if it, ok := x.cache[id]; ok {
		return it.model, nil
	}
	data, token, err := x.getModel(ctx, id, opts...)
	if err != nil {
		return
	}
	if model, err = data.UnmarshalNew(); err != nil {
		return
	}
	x.cache[id] = &item{model, token, data}
	return
}

func (x *mvc) LoadModel(ctx context.Context, id string, model pb.Message, opts ...GetOption) (err error) {
	if it, ok := x.cache[id]; ok {
		pb.Reset(model)
		pb.Merge(model, it.model)
		it.model = model
		return
	}
	data, token, err := x.getModel(ctx, id, opts...)
	if err != nil {
		return
	}
	if err = data.UnmarshalTo(model); err != nil {
		return
	}
	x.cache[id] = &item{model, token, data}
	return
}

func (x *mvc) getModel(
	ctx context.Context, id string, opts ...GetOption) (
	_ *anypb.Any, token string, err error) {
	var o getOptions
	for _, opt := range opts {
		opt.apply(&o)
	}
	req := &dbpb.ArchiveGetRequest{Id: id, WithLock: true}
	if o.addIfNotFound != nil {
		req.AddIfNotFound = &dbpb.Archive{
			Id:    id,
			Model: newAny(o.addIfNotFound),
		}
	}
	log.Infof("getModel req: %#v", req)
	resp, err := x.dbapi.GetModel(ctx, req)
	if err != nil {
		return
	}
	log.Infof("getModel resp: %#v", resp)
	return resp.Archive.Model, resp.Token, nil
}

func (x *mvc) submit(ctx context.Context, handleErr error) (firstErr error) {
	firstErr = handleErr
	for id, it := range x.cache {
		if err := func() (err error) {
			req := &dbpb.ArchiveSetRequest{
				Archive:    &dbpb.Archive{Id: id},
				Token:      it.token,
				WithUnlock: true,
			}
			if data := newAny(it.model); pb.Equal(data, it.data) {
				if _, err = x.dbapi.Unlock(ctx, req); err != nil {
					return
				}
			} else {
				req.Archive.Model = data
				if _, err = x.dbapi.SetModel(ctx, req); err != nil {
					return
				}
				msg := &sdkpb.UpdateModelNotice{Model: data}
				if err = x.PushTo(ctx, id, msg); err != nil {
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
