package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"

	acctpb "github.com/ntons/libra-go/api/acct/v1"
	echopb "github.com/ntons/libra-go/api/sdk/example/echo"
	sdkpb "github.com/ntons/libra-go/api/sdk/v1"
	"github.com/ntons/libra-go/client/sdk"
)

var (
	appId = "example"
)

type EchoClient struct {
	sdk.Client
	echopb.EchoClient
}

func Dial() (_ *EchoClient, err error) {
	cli, err := sdk.Dial(
		appId,
		"127.0.0.1:10000",
		sdk.WithGrpcDialOption(grpc.WithInsecure()),
	)
	if err != nil {
		return
	}
	return &EchoClient{
		Client:     cli,
		EchoClient: echopb.NewEchoClient(cli),
	}, nil
}

func run() (err error) {
	cli, err := Dial()
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}
	defer cli.Close()

	var wg sync.WaitGroup
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	wg.Add(1)
	go func() {
		defer wg.Done()
		if msg, err := cli.Recv(ctx); err != nil {
			fmt.Println("failed to recv: ", err)
		} else {
			update := msg.(*sdkpb.UpdateModelNotice)
			fmt.Println(update.Model.UnmarshalNew())
		}
	}()

	resp, err := cli.Send(ctx, &echopb.SendRequest{Content: "foo"})
	if err != nil {
		return fmt.Errorf("failed to say hello: %v", err)
	}
	fmt.Println("reply: ", resp)

	<-time.After(time.Second)
	return
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
