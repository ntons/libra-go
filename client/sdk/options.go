package sdk

import (
	"google.golang.org/grpc"
)

type dialOptions struct {
	// grpc dial opts
	dialOptions []grpc.DialOption
}

type DialOption interface {
	apply(o *dialOptions)
}

type fnDialOption struct {
	fn func(*dialOptions)
}

func (fo fnDialOption) apply(o *dialOptions) { fo.fn(o) }

func WithGrpcDialOption(opts ...grpc.DialOption) DialOption {
	return fnDialOption{func(o *dialOptions) { o.dialOptions = opts }}
}
