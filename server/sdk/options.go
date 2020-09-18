package sdk

import (
	"google.golang.org/grpc"
)

type serveOptions struct {
	serverOptions []grpc.ServerOption
}

type ServeOption interface {
	apply(*serveOptions)
}

type fnServeOption struct {
	fn func(*serveOptions)
}

func (fo fnServeOption) apply(o *serveOptions) { fo.fn(o) }

func WithGrpcServerOption(opts ...grpc.ServerOption) ServeOption {
	return fnServeOption{func(o *serveOptions) { o.serverOptions = opts }}
}

type dialOptions struct {
	dialOptions []grpc.DialOption
}
type DialOption interface {
	apply(*dialOptions)
}

type fnDialOption struct {
	fn func(*dialOptions)
}

func (fo fnDialOption) apply(o *dialOptions) { fo.fn(o) }

func WithGrpcDialOption(opts ...grpc.DialOption) DialOption {
	return fnDialOption{func(o *dialOptions) { o.dialOptions = opts }}
}
