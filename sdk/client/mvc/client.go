package mvc

import (
	"context"

	"github.com/ntons/log-go"
	"google.golang.org/protobuf/proto"

	sdkpb "github.com/ntons/libra-go/api/sdk/v1"
	sdk "github.com/ntons/libra-go/client"
)

type client struct {
	sdk.Client
	model proto.Message
}

func New(cli sdk.Client) Client {
	return &client{Client: cli}
}

func (cli *client) GetModel() proto.Message { return cli.model }

func (cli *client) Recv(ctx context.Context) (msg proto.Message, err error) {
	for {
		if msg, err = cli.Client.Recv(ctx); err != nil {
			return
		}
		switch msg.(type) {
		case *sdkpb.UpdateModelNotice:
			cli.updateModel(msg.(*sdkpb.UpdateModelNotice))
		default:
			return
		}
	}
}

func (cli *client) updateModel(msg *sdkpb.UpdateModelNotice) {
	log.Infof("update model: %v", msg)
}
