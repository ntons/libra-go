package main

import (
	"context"
	"encoding/json"

	"github.com/ntons/log-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dbpb "github.com/ntons/libra-go/api/db/v1"
	gwpb "github.com/ntons/libra-go/api/gw/v1"
	echopb "github.com/ntons/libra-go/api/sdk/example/echo"
	"github.com/ntons/libra-go/server/mvc"
	"github.com/ntons/libra-go/server/sdk"
)

type echoServer struct {
	echopb.UnimplementedEchoServer
	dbapi dbpb.DBClient
	gwapi gwpb.GatewayClient
}

func create(b json.RawMessage) (svc sdk.Service, err error) {
	log.Infof("create echo service: %s", b)
	return &echoServer{}, nil
}

func (srv *echoServer) Register(
	gs *grpc.Server, conn grpc.ClientConnInterface) error {
	echopb.RegisterEchoServer(gs, srv)
	srv.dbapi = dbpb.NewDBClient(conn)
	srv.gwapi = gwpb.NewGatewayClient(conn)
	return nil
}
func (srv *echoServer) Stop() {
}

func (srv *echoServer) Send(
	ctx context.Context, req *echopb.SendRequest) (
	resp *echopb.SendResponse, err error) {
	x, ok := mvc.FromContext(ctx)
	if !ok {
		log.Warnf("failed to get mvc from context")
		err = status.Errorf(codes.Internal, "failed to get mvc from context")
		return
	}
	var model echopb.Model
	if err = x.LoadModel(
		ctx, x.GetRoleId(), &model,
		mvc.WithAddIfNotFound(&model)); err != nil {
		log.Warnf("failed to load model: %v", err)
		err = status.Errorf(codes.Internal, "failed to load model")
		return
	}
	model.History = append(model.History, req.Content)
	if len(model.History) > 10 {
		model.History = model.History[:10]
	}
	resp = &echopb.SendResponse{Content: req.Content}
	return
}
