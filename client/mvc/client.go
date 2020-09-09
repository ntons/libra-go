package mvc

import (
	"context"

	"google.golang.org/protobuf/proto"

	sdk_v1 "github.com/ntons/libra-go/api/sdk/v1"
	"github.com/ntons/libra-go/client/sdk"
)

type Client struct {
	*sdk.Client
	model proto.Message
}

func New(cli *sdk.Client, model proto.Message) *Client {
	return &Client{Client: cli}
}

func (cli *Client) GetModel() proto.Message { return cli.model }

func (cli *Client) Recv(ctx context.Context) (msg proto.Message, err error) {
	for {
		if msg, err = cli.Client.Recv(ctx); err != nil {
			return
		}
		switch msg.(type) {
		case *sdk_v1.UpdateModelNotice:
			cli.updateModel(msg.(*sdk_v1.UpdateModelNotice))
		default:
			return
		}
	}
}

func (cli *Client) updateModel(msg *sdk_v1.UpdateModelNotice) {
	// TODO How model works?
}
