package main

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"

	chatpb "github.com/ntons/libra-go/api/sdk/example/chat"
	"github.com/ntons/libra-go/server/sdk"
	"github.com/ntons/log-go"

	dbpb "github.com/ntons/libra-go/api/db/v1"
	gwpb "github.com/ntons/libra-go/api/gw/v1"
)

type chatServer struct {
	chatpb.UnimplementedChatServer
	dbapi dbpb.DBClient
	gwapi gwpb.GatewayClient
}

func create(b json.RawMessage) (svc sdk.Service, err error) {
	log.Infof("create chat service: %s", b)
	return &chatServer{}, nil
}

func (srv *chatServer) Register(
	gs *grpc.Server, conn grpc.ClientConnInterface) error {
	chatpb.RegisterChatServer(gs, srv)
	srv.dbapi = dbpb.NewDBClient(conn)
	srv.gwapi = gwpb.NewGatewayClient(conn)
	return nil
}
func (srv *chatServer) Stop() {
}

func (srv *chatServer) Send(
	ctx context.Context, msg *chatpb.SendRequest) (
	resp *chatpb.SendResponse, err error) {
	return
}
