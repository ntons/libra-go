package main

import (
	"context"
	"flag"
	"sync"

	"github.com/ntons/log-go"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	sdk "github.com/ntons/libra-go/server"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (GreeterServer) SayHello(
	ctx context.Context, req *pb.HelloRequest) (
	resp *pb.HelloReply, err error) {
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	var (
		listen string
		api    string
	)
	flag.StringVar(&listen, "listen", ":80", "Server listen address")
	flag.StringVar(&api, "api", "", "Libra api access point")
	flag.Parse()

	log.Debugf("serve on: \"%s\"", listen)
	log.Debugf("libra api: \"%s\"", api)

	srv := sdk.NewServer()
	echo := &GreeterServer{}
	pb.RegisterGreeterServer(srv, echo)

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		srv.ListenAndServe(listen)
	}()
	defer srv.Shutdown()

	sdk.WaitForSignals(sdk.SIGINT, sdk.SIGTERM)
}
