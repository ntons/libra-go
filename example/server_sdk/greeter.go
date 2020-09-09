package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	hw "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"

	dbapi_v1 "github.com/ntons/libra-go/api/dbapi/v1"
	"github.com/ntons/libra-go/example/protos"
	"github.com/ntons/libra-go/server/sdk"
)

type greeterServer struct {
	hw.UnimplementedGreeterServer
}

func newGreeterServer() *greeterServer {
	return &greeterServer{}
}

func (*greeterServer) SayHello(
	ctx context.Context, req *hw.HelloRequest) (
	reply *hw.HelloReply, err error) {
	call, ok := sdk.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	req1 := &dbapi_v1.ArchiveGetRequest{
		Id:            call.GetRoleId(),
		AddIfNotFound: &dbapi_v1.Archive{},
		WithLock:      true,
	}
	req1.AddIfNotFound.Model, _ = anypb.New(&protos.Model{})
	resp1, err := call.GetDBV1().GetArchive(ctx, req1)
	if err != nil {
		call.Warnf("failed to get archive: %v", err)
		return
	}

	model := &protos.Model{}
	if err = resp1.Archive.Model.UnmarshalTo(model); err != nil {
		call.Warnf("failed to unmarshal model: %v", err)
		return
	}
	model.History = append(model.History,
		fmt.Sprintf("%s receive hello from %s", time.Now(), req.Name))

	req2 := &dbapi_v1.ArchiveSetRequest{
		Archive:    resp1.Archive,
		Token:      resp1.Token,
		WithUnlock: true,
	}
	req2.Archive.Model, _ = anypb.New(model)
	if _, err = call.GetDBV1().SetArchive(ctx, req2); err != nil {
		call.Warnf("failed to set archive: %v", err)
		return
	}

	call.Infof("say hello from %v", req.Name)
	call.Infof("history: %v", model.History)
	reply = &hw.HelloReply{
		Message: fmt.Sprintf("Hello %s!", req.Name),
	}

	return
}
