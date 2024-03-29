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

// SubstrIndex1ServiceClient is the client API for SubstrIndex1Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubstrIndex1ServiceClient interface {
	// 更新索引
	Update(ctx context.Context, in *SubstrIndex1_UpdateRequest, opts ...grpc.CallOption) (*SubstrIndex1_UpdateResponse, error)
	// 删除索引
	Remove(ctx context.Context, in *SubstrIndex1_RemoveRequest, opts ...grpc.CallOption) (*SubstrIndex1_RemoveResponse, error)
	// 搜索对象
	Search(ctx context.Context, in *SubstrIndex1_SearchRequest, opts ...grpc.CallOption) (*SubstrIndex1_SearchResponse, error)
}

type substrIndex1ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubstrIndex1ServiceClient(cc grpc.ClientConnInterface) SubstrIndex1ServiceClient {
	return &substrIndex1ServiceClient{cc}
}

func (c *substrIndex1ServiceClient) Update(ctx context.Context, in *SubstrIndex1_UpdateRequest, opts ...grpc.CallOption) (*SubstrIndex1_UpdateResponse, error) {
	out := new(SubstrIndex1_UpdateResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.SubstrIndex1Service/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *substrIndex1ServiceClient) Remove(ctx context.Context, in *SubstrIndex1_RemoveRequest, opts ...grpc.CallOption) (*SubstrIndex1_RemoveResponse, error) {
	out := new(SubstrIndex1_RemoveResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.SubstrIndex1Service/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *substrIndex1ServiceClient) Search(ctx context.Context, in *SubstrIndex1_SearchRequest, opts ...grpc.CallOption) (*SubstrIndex1_SearchResponse, error) {
	out := new(SubstrIndex1_SearchResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.SubstrIndex1Service/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubstrIndex1ServiceServer is the server API for SubstrIndex1Service service.
// All implementations must embed UnimplementedSubstrIndex1ServiceServer
// for forward compatibility
type SubstrIndex1ServiceServer interface {
	// 更新索引
	Update(context.Context, *SubstrIndex1_UpdateRequest) (*SubstrIndex1_UpdateResponse, error)
	// 删除索引
	Remove(context.Context, *SubstrIndex1_RemoveRequest) (*SubstrIndex1_RemoveResponse, error)
	// 搜索对象
	Search(context.Context, *SubstrIndex1_SearchRequest) (*SubstrIndex1_SearchResponse, error)
	mustEmbedUnimplementedSubstrIndex1ServiceServer()
}

// UnimplementedSubstrIndex1ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubstrIndex1ServiceServer struct {
}

func (UnimplementedSubstrIndex1ServiceServer) Update(context.Context, *SubstrIndex1_UpdateRequest) (*SubstrIndex1_UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSubstrIndex1ServiceServer) Remove(context.Context, *SubstrIndex1_RemoveRequest) (*SubstrIndex1_RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedSubstrIndex1ServiceServer) Search(context.Context, *SubstrIndex1_SearchRequest) (*SubstrIndex1_SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSubstrIndex1ServiceServer) mustEmbedUnimplementedSubstrIndex1ServiceServer() {}

// UnsafeSubstrIndex1ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubstrIndex1ServiceServer will
// result in compilation errors.
type UnsafeSubstrIndex1ServiceServer interface {
	mustEmbedUnimplementedSubstrIndex1ServiceServer()
}

func RegisterSubstrIndex1ServiceServer(s grpc.ServiceRegistrar, srv SubstrIndex1ServiceServer) {
	s.RegisterService(&SubstrIndex1Service_ServiceDesc, srv)
}

func _SubstrIndex1Service_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubstrIndex1_UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubstrIndex1ServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.SubstrIndex1Service/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubstrIndex1ServiceServer).Update(ctx, req.(*SubstrIndex1_UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubstrIndex1Service_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubstrIndex1_RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubstrIndex1ServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.SubstrIndex1Service/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubstrIndex1ServiceServer).Remove(ctx, req.(*SubstrIndex1_RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubstrIndex1Service_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubstrIndex1_SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubstrIndex1ServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.SubstrIndex1Service/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubstrIndex1ServiceServer).Search(ctx, req.(*SubstrIndex1_SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubstrIndex1Service_ServiceDesc is the grpc.ServiceDesc for SubstrIndex1Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubstrIndex1Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.SubstrIndex1Service",
	HandlerType: (*SubstrIndex1ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _SubstrIndex1Service_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _SubstrIndex1Service_Remove_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _SubstrIndex1Service_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libra/v1/indexing.proto",
}
