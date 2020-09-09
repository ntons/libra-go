package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"

	"google.golang.org/grpc"
	hw "google.golang.org/grpc/examples/helloworld/helloworld"

	acct_v1 "github.com/ntons/libra-go/api/acct/v1"
	"github.com/ntons/libra-go/client/sdk"
)

type AppClient struct {
	*sdk.Client
	hw.GreeterClient
}

func Dial(
	addr, appId string, opts ...grpc.DialOption) (_ *AppClient, err error) {
	cli, err := sdk.Dial(addr, appId, opts...)
	if err != nil {
		return
	}
	return &AppClient{
		Client:        cli,
		GreeterClient: hw.NewGreeterClient(cli),
	}, nil
}

func run(addr, appId string) (err error) {
	cli, err := Dial(addr, appId, grpc.WithInsecure())
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
	var addr string
	flag.StringVar(&addr, "addr", "", "address to server")
	flag.Parse()
	if err := run(addr, "example"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
