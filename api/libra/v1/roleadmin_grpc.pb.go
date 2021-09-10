// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RoleAdminClient is the client API for RoleAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleAdminClient interface {
	Get(ctx context.Context, in *RoleAdminGetRequest, opts ...grpc.CallOption) (*RoleAdminGetResponse, error)
}

type roleAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleAdminClient(cc grpc.ClientConnInterface) RoleAdminClient {
	return &roleAdminClient{cc}
}

func (c *roleAdminClient) Get(ctx context.Context, in *RoleAdminGetRequest, opts ...grpc.CallOption) (*RoleAdminGetResponse, error) {
	out := new(RoleAdminGetResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.RoleAdmin/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleAdminServer is the server API for RoleAdmin service.
// All implementations must embed UnimplementedRoleAdminServer
// for forward compatibility
type RoleAdminServer interface {
	Get(context.Context, *RoleAdminGetRequest) (*RoleAdminGetResponse, error)
	mustEmbedUnimplementedRoleAdminServer()
}

// UnimplementedRoleAdminServer must be embedded to have forward compatible implementations.
type UnimplementedRoleAdminServer struct {
}

func (UnimplementedRoleAdminServer) Get(context.Context, *RoleAdminGetRequest) (*RoleAdminGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRoleAdminServer) mustEmbedUnimplementedRoleAdminServer() {}

// UnsafeRoleAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleAdminServer will
// result in compilation errors.
type UnsafeRoleAdminServer interface {
	mustEmbedUnimplementedRoleAdminServer()
}

func RegisterRoleAdminServer(s grpc.ServiceRegistrar, srv RoleAdminServer) {
	s.RegisterService(&RoleAdmin_ServiceDesc, srv)
}

func _RoleAdmin_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAdminGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleAdminServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.RoleAdmin/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleAdminServer).Get(ctx, req.(*RoleAdminGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleAdmin_ServiceDesc is the grpc.ServiceDesc for RoleAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.RoleAdmin",
	HandlerType: (*RoleAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RoleAdmin_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libra/v1/roleadmin.proto",
}
