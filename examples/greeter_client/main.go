package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"

	hellopb "google.golang.org/grpc/examples/helloworld/helloworld"

	v1pb "github.com/ntons/libra-go/api/libra/v1"
	sdk "github.com/ntons/libra-go/sdk/client"
)

type GreeterClient struct {
	*sdk.Client
	hellopb.GreeterClient
}

func dial(addr string) (_ *GreeterClient, err error) {
	cli, err := sdk.Dial("greeter", addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	return &GreeterClient{
		Client:        cli,
		GreeterClient: hellopb.NewGreeterClient(cli),
	}, nil
}

func login(
	ctx context.Context, cli *GreeterClient, username string) (err error) {
	user, err := cli.Login(ctx, &v1pb.DevLoginState{Username: username})
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
	resp, err := cli.SayHello(ctx, &hellopb.HelloRequest{Name: "World"})
	if err != nil {
		return fmt.Errorf("failed to say hello: %v", err)
	}
	fmt.Println("echo reply: ", resp.Message)
	return
}

func main() {
	var (
		addr string
		user string
	)
	flag.StringVar(&addr, "a", "127.0.0.1:10000", "server address")
	flag.StringVar(&user, "u", "", "username")
	flag.Parse()
	if user == "" {
		fmt.Println("require username")
		os.Exit(1)
	}

	cli, err := dial(addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err = login(ctx, cli, user); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := sayHello(ctx, cli); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
