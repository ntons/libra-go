package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"google.golang.org/grpc"

	echopb "github.com/ntons/libra-go/api/examples/echo"
	v1pb "github.com/ntons/libra-go/api/v1"
	"github.com/ntons/libra-go/client"
)

var (
	appId = "example"
)

type EchoClient struct {
	*client.Client
	echopb.EchoClient
}

func dial(addr string) (_ *EchoClient, err error) {
	cli, err := client.Dial(appId, addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	return &EchoClient{
		Client:     cli,
		EchoClient: echopb.NewEchoClient(cli),
	}, nil
}

func login(ctx context.Context, cli *EchoClient) (err error) {
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
	/*
		if err := cli.Connect(ctx); err != nil {
			return fmt.Errorf("failed to access: %v", err)
		}
	*/
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

func echo(ctx context.Context, cli *EchoClient) (err error) {
	resp, err := cli.Echo(ctx, &echopb.EchoRequest{Content: "hello"})
	if err != nil {
		return fmt.Errorf("failed to say hello: %v", err)
	}
	fmt.Println("echo reply: ", resp.Content)
	return
}

func repeat(ctx context.Context, cli *EchoClient) (err error) {
	stream, err := cli.Repeat(ctx, &echopb.EchoRepeatRequest{
		Content:  "hello",
		Interval: 100,
		Count:    10,
	})
	if err != nil {
		return fmt.Errorf("failed to repeat: %v", err)
	}
	for i := 1; ; i++ {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to recv: %v", err)
		}
		fmt.Printf("repeat reply %02d: %v\n", i, resp.Content)
	}
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
	if err := echo(ctx, cli); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := repeat(ctx, cli); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
