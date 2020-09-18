package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	hw "google.golang.org/grpc/examples/helloworld/helloworld"

	acct_v1 "github.com/ntons/libra-go/api/acct/v1"
	_ "github.com/ntons/libra-go/api/sdk/v1"
	"github.com/ntons/libra-go/client/sdk"
)

var (
	appId = "example"
)

type AppClient struct {
	*sdk.Client
	hw.GreeterClient
}

func Dial() (_ *AppClient, err error) {
	cli, err := sdk.Dial(
		appId,
		&sdk.Endpoint{Target: "127.0.0.1:80"},
		sdk.WithInsecure(),
		sdk.WithGwEndpoint(&sdk.Endpoint{Target: "127.0.0.1:8080"}),
		sdk.WithPtEndpoint(&sdk.Endpoint{Target: "127.0.0.1:8080"}),
	)
	if err != nil {
		return
	}
	return &AppClient{
		Client:        cli,
		GreeterClient: hw.NewGreeterClient(cli),
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

	user, err := cli.Login(ctx, &acct_v1.Dev{Username: "userxxxx"})
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
			fmt.Println("recv: ", msg)
		}
	}()

	x := hw.NewGreeterClient(cli)
	resp, err := x.SayHello(ctx, &hw.HelloRequest{Name: "foo"})
	if err != nil {
		return fmt.Errorf("failed to say hello: %v", err)
	}
	fmt.Println("reply: ", resp)
	return
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
