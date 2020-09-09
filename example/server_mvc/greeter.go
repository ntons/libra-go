package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	hw "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/status"

	"github.com/ntons/libra-go/example/protos"
	"github.com/ntons/libra-go/server/mvc"
)

type Role struct {
	protos.Model
}

func (m *Role) Foo() {
	// TODO custom methods
}

type greeterServer struct {
	hw.UnimplementedGreeterServer
}

func newGreeterServer() *greeterServer {
	return &greeterServer{}
}

func (*greeterServer) SayHello(ctx context.Context, req *hw.HelloRequest) (
	reply *hw.HelloReply, err error) {
	call, ok := mvc.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "bad context")
	}
	var role Role
	if err = call.LoadModel(
		ctx, call.GetRoleId(), &role,
		mvc.WithAddIfNotFound(&Role{}),
	); err != nil {
		call.Warnf("failed to load model: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to load model")
	}
	role.History = append(role.History, fmt.Sprintf(
		"%s receive hello from %s", time.Now(), req.Name))
	call.Infof("say hello from %v", req.Name)
	call.Infof("history: %v", role.History)
	reply = &hw.HelloReply{Message: fmt.Sprintf("Hello %s!", req.Name)}
	return
}
