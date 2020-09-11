package sdk

import (
	"google.golang.org/grpc"

	sdk_v1 "github.com/ntons/libra-go/api/sdk/v1"
)

type dialOptions struct {
	// grpc dial opts
	opts []grpc.DialOption
	// particular endpoint for libra
	libraAccessPoint *sdk_v1.Endpoint
}

type DialOption interface {
	apply(o *dialOptions)
}

type fnDialOption struct {
	fn func(*dialOptions)
}

func (fo fnDialOption) apply(o *dialOptions) { fo.fn(o) }

func WithLibraAccessPoint(ap *sdk_v1.Endpoint) DialOption {
	return fnDialOption{func(o *dialOptions) {
		o.libraAccessPoint = ap
	}}
}
func WithInsecure() DialOption {
	return fnDialOption{func(o *dialOptions) {
		o.opts = append(o.opts, grpc.WithInsecure())
	}}
}
