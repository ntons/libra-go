package main

import (
	"github.com/ntons/libra-go/server/sdk"
	"github.com/ntons/log-go"
)

func main() {
	if err := sdk.Serve(create); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
