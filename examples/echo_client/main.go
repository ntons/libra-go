package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"

	acctpb "github.com/ntons/libra-go/api/acct/v1"
	echopb "github.com/ntons/libra-go/api/sdk/example/echo"
	"github.com/ntons/libra-go/client/mvc"
	"github.com/ntons/libra-go/client/sdk"
)

var (
	appId = "example"
)

type EchoClient struct {
	mvc.Client
	echopb.EchoClient
}

func dial(addr string) (_ *EchoClient, err error) {
	cli, err := sdk.Dial(
		appId, addr,
		sdk.WithGrpcDialOption(grpc.WithInsecure()),
	)
	if err != nil {
		return
	}
	return &EchoClient{
		Client:     mvc.New(cli),
		EchoClient: echopb.NewEchoClient(cli),
	}, nil
}

func login(ctx context.Context, cli *EchoClient) (err error) {
	user, err := cli.Login(ctx, &acctpb.Dev{Username: "userxxxx"})
	if err != nil {
		return fmt.Errorf("failed to login: %v", err)
	}
	fmt.Println("login: ", user)

	roles, err := cli.ListRoles(ctx)
	if err != nil {
		return fmt.Errorf("failed to list roles: %v", err)
	}
	fmt.Println("list roles:", roles)

	if len(roles) == 0 {
		role, err := cli.CreateRole(ctx, 1)
		if err != nil {
			return fmt.Errorf("failed to create role: %v", err)
		}
		fmt.Println("create role: ", role)
		roles = append(roles, role)
	}
	if err := cli.SignIn(ctx, roles[0].Id); err != nil {
		return fmt.Errorf("failed to sign in: %v", err)
	}
	fmt.Println("sign in: ok")

	go func() {
		fmt.Println(cli.Recv(ctx))
		//if msg, err := cli.Recv(ctx); err != nil {
		//	fmt.Println("failed to recv: ", err)
		//} else {
		//	update := msg.(*sdkpb.UpdateModelNotice)
		//	fmt.Println(update.Model.UnmarshalNew())
		//}
	}()
	return
}

func echo(ctx context.Context, cli *EchoClient, content string) (err error) {
	resp, err := cli.Send(ctx, &echopb.SendRequest{Content: content})
	if err != nil {
		return fmt.Errorf("failed to say hello: %v", err)
	}
	fmt.Println("reply: ", resp)

	<-time.After(time.Second)
	return
}

func main() {
	var (
		address string
		content string
	)
	flag.StringVar(&address, "a", "127.0.0.1:10000", "server address")
	flag.StringVar(&content, "m", "hello", "message content")
	flag.Parse()

	cli, err := dial(address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err = login(ctx, cli); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := echo(ctx, cli, content); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	<-ctx.Done()
}
