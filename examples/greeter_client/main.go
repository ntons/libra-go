package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	v1pb "github.com/ntons/libra-go/api/v1"
	sdk "github.com/ntons/libra-go/client"
)

var (
	appId = "example"
)

type GreeterClient struct {
	*sdk.Client
	pb.GreeterClient
}

func dial(addr string) (_ *GreeterClient, err error) {
	cli, err := sdk.Dial(appId, addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	return &GreeterClient{
		Client:        cli,
		GreeterClient: pb.NewGreeterClient(cli),
	}, nil
}

func login(ctx context.Context, cli *GreeterClient) (err error) {
	user, err := cli.Login(ctx, &v1pb.DevLoginState{Username: "userxxxx"})
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
	return
}

func sayHello(ctx context.Context, cli *GreeterClient) (err error) {
	resp, err := cli.SayHello(ctx, &pb.HelloRequest{Name: "foo"})
	if err != nil {
		return fmt.Errorf("failed to say hello: %v", err)
	}
	fmt.Println("echo reply: ", resp.Message)
	return
}

func main() {
	var address string
	flag.StringVar(&address, "a", "127.0.0.1:10000", "server address")
	flag.Parse()

	cli, err := dial(address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err = login(ctx, cli); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := sayHello(ctx, cli); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
