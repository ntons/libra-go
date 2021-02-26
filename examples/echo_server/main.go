package main

import (
	"context"
	"flag"
	"sync"
	"time"

	"github.com/ntons/libra-go/server"
	"github.com/ntons/log-go"

	examplespb "github.com/ntons/libra-go/api/examples"
)

type EchoServer struct {
	examplespb.UnimplementedEchoServer
}

func (EchoServer) Echo(
	ctx context.Context, req *examplespb.EchoRequest) (
	resp *examplespb.EchoResponse, err error) {
	resp = &examplespb.EchoResponse{Content: req.Content}
	return
}
func (EchoServer) Repeat(
	req *examplespb.EchoRepeatRequest, stream examplespb.Echo_RepeatServer) (err error) {
	for i := int32(0); i < req.Count; i++ {
		if err = stream.Send(&examplespb.EchoRepeatResponse{
			Content: req.Content,
			Seq:     i,
		}); err != nil {
			return
		}
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case <-time.After(time.Duration(req.Interval) * time.Millisecond):
		}
	}
	return
}

func main() {
	var (
		listen string
		api    string
	)
	flag.StringVar(&listen, "listen", ":80", "Server listen address")
	flag.StringVar(&api, "api", "", "Libra api access point")
	flag.Parse()

	log.Debugf("echo server serve on: \"%s\"", listen)
	log.Debugf("libra api: \"%s\"", api)

	srv := server.NewServer()
	echo := &EchoServer{}
	examplespb.RegisterEchoServer(srv, echo)

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		srv.ListenAndServe(listen)
	}()
	defer srv.Shutdown()

	server.WaitForSignals(server.SIGINT, server.SIGTERM)
}
