// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package libra_gwapi_v1

import (
	context "context"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccessClient is the client API for Access service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessClient interface {
	// Establish the push stream
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (Access_SignInClient, error)
}

type accessClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessClient(cc grpc.ClientConnInterface) AccessClient {
	return &accessClient{cc}
}

func (c *accessClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (Access_SignInClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Access_serviceDesc.Streams[0], "/libra.gwapi.v1.Access/SignIn", opts...)
	if err != nil {
		return nil, err
	}
	x := &accessSignInClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Access_SignInClient interface {
	Recv() (*any.Any, error)
	grpc.ClientStream
}

type accessSignInClient struct {
	grpc.ClientStream
}

func (x *accessSignInClient) Recv() (*any.Any, error) {
	m := new(any.Any)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccessServer is the server API for Access service.
// All implementations must embed UnimplementedAccessServer
// for forward compatibility
type AccessServer interface {
	// Establish the push stream
	SignIn(*SignInRequest, Access_SignInServer) error
	mustEmbedUnimplementedAccessServer()
}

// UnimplementedAccessServer must be embedded to have forward compatible implementations.
type UnimplementedAccessServer struct {
}

func (*UnimplementedAccessServer) SignIn(*SignInRequest, Access_SignInServer) error {
	return status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedAccessServer) mustEmbedUnimplementedAccessServer() {}

func RegisterAccessServer(s *grpc.Server, srv AccessServer) {
	s.RegisterService(&_Access_serviceDesc, srv)
}

func _Access_SignIn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SignInRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccessServer).SignIn(m, &accessSignInServer{stream})
}

type Access_SignInServer interface {
	Send(*any.Any) error
	grpc.ServerStream
}

type accessSignInServer struct {
	grpc.ServerStream
}

func (x *accessSignInServer) Send(m *any.Any) error {
	return x.ServerStream.SendMsg(m)
}

var _Access_serviceDesc = grpc.ServiceDesc{
	ServiceName: "libra.gwapi.v1.Access",
	HandlerType: (*AccessServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SignIn",
			Handler:       _Access_SignIn_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gwapi/v1/access.proto",
}
