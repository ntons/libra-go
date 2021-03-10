package main

import (
	"context"
	"flag"
	"os"
	"sync"

	v1pb "github.com/ntons/libra-go/api/v1"
	"github.com/ntons/log-go"
	"google.golang.org/grpc"
	hellopb "google.golang.org/grpc/examples/helloworld/helloworld"
	anypb "google.golang.org/protobuf/types/known/anypb"

	sdk "github.com/ntons/libra-go/sdk/server"
)

type GreeterServer struct {
	hellopb.UnimplementedGreeterServer
	api *sdk.Client
}

func (x GreeterServer) SayHello(
	ctx context.Context, req *hellopb.HelloRequest) (
	resp *hellopb.HelloReply, err error) {
	if call, ok := sdk.FromIncomingContext(ctx); ok {
		key := &v1pb.EntryKey{Kind: "role", Id: call.RoleId}
		getReq := &v1pb.DatabaseGetRequest{
			Key:         key,
			LockOptions: &v1pb.DistlockLockOptions{},
		}
		getReq.AddIfNotFound, _ = anypb.New(&Archive{RoleId: call.RoleId})
		getResp, err := x.api.Database.Get(ctx, getReq)
		if err != nil {
			return nil, err
		}
		defer func() {
			if err != nil {
				req := &v1pb.DistlockUnlockRequest{
					LockToken: getResp.LockToken,
				}
				x.api.Distlock.Unlock(ctx, req)
			}
		}()

		var archive Archive
		if err := getResp.Data.UnmarshalTo(&archive); err != nil {
			return nil, err
		}

		archive.History = append(archive.History, req.GetName())

		setReq := &v1pb.DatabaseSetRequest{
			Key:       key,
			LockToken: getResp.LockToken,
			UnlockOptions: &v1pb.DistlockUnlockOptions{
				EvenOnFailure: true,
			},
		}
		setReq.Data, _ = anypb.New(&archive)
		setResp, err := x.api.Database.Set(ctx, setReq)
		if err != nil {
			return nil, err
		}
		log.Debugf(
			"archive revision for %s: %d",
			call.RoleId, setResp.GetRevision())
	}
	return &hellopb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	var (
		listen string
		api    string
	)
	flag.StringVar(&listen, "listen", ":80", "Server listen address")
	flag.StringVar(&api, "api", "", "Libra api access point")
	flag.Parse()

	if api == "" {
		log.Errorf("requre api url")
		os.Exit(1)
	}

	log.Debugf("serve on: \"%s\"", listen)
	log.Debugf("libra api: \"%s\"", api)

	cli, err := sdk.Dial(api, grpc.WithInsecure())
	if err != nil {
		log.Debugf("failed to dail to libra: %v", err)
		os.Exit(1)
	}
	cli.AddAuth("greeter", "3f67ae95ed060e33d5ac351db031f1c6")

	srv := sdk.NewServer()
	hellopb.RegisterGreeterServer(srv, &GreeterServer{api: cli})

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)

	quit := make(chan struct{}, 1)
	go func() {
		defer func() { quit <- struct{}{}; wg.Done() }()
		if err := srv.ListenAndServe(listen); err != nil {
			log.Errorf("exit with error: %v", err)
		}
	}()
	defer srv.Stop()

	select {
	case <-quit:
	case <-sdk.WaitForSignals(sdk.SIGINT, sdk.SIGTERM):
	}
	log.Debugf("terminated")
}
