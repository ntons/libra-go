package main

import (
	"fmt"
	"os"

	"github.com/ntons/libra-go/server/sdk"
	hw "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
	s := sdk.NewServer()
	if err := s.Init(); err != nil {
		fmt.Println("failed to init server: ", err)
		os.Exit(1)
	}
	x := newGreeterServer()
	hw.RegisterGreeterServer(s.GRPCServer(), x)
	s.Serve()
}
