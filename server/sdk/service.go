package sdk

import (
	"encoding/json"

	"google.golang.org/grpc"
)

type Service interface {
	// register grpc service
	// srv:  grpc server which provide services from
	// conn: grpc connection to libra access point
	Register(srv *grpc.Server, conn grpc.ClientConnInterface) error

	// stop serving, clean resources
	Stop()
}

// ServiceFactory create service
// cfg: raw json from config file
type ServiceFactory func(cfg json.RawMessage) (svc Service, err error)
