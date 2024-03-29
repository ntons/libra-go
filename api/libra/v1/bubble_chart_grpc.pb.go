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

// BubbleChartClient is the client API for BubbleChart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BubbleChartClient interface {
	// append entries to the end of chart
	Append(ctx context.Context, in *BubbleChartAppendRequest, opts ...grpc.CallOption) (*BubbleChartAppendResponse, error)
	// swap 2 entries by id
	SwapById(ctx context.Context, in *BubbleChartSwapByIdRequest, opts ...grpc.CallOption) (*BubbleChartSwapByIdResponse, error)
	// swap 2 entries by rank
	SwapByRank(ctx context.Context, in *BubbleChartSwapByRankRequest, opts ...grpc.CallOption) (*BubbleChartSwapByRankResponse, error)
	// get entries by range
	GetRange(ctx context.Context, in *BubbleChartGetRangeRequest, opts ...grpc.CallOption) (*BubbleChartGetRangeResponse, error)
	// get entries by id
	GetById(ctx context.Context, in *BubbleChartGetByIdRequest, opts ...grpc.CallOption) (*BubbleChartGetByIdResponse, error)
	// remove entries by id
	RemoveById(ctx context.Context, in *BubbleChartRemoveByIdRequest, opts ...grpc.CallOption) (*BubbleChartRemoveByIdResponse, error)
	// set entry info by id
	SetInfo(ctx context.Context, in *BubbleChartSetInfoRequest, opts ...grpc.CallOption) (*BubbleChartSetInfoResponse, error)
}

type bubbleChartClient struct {
	cc grpc.ClientConnInterface
}

func NewBubbleChartClient(cc grpc.ClientConnInterface) BubbleChartClient {
	return &bubbleChartClient{cc}
}

func (c *bubbleChartClient) Append(ctx context.Context, in *BubbleChartAppendRequest, opts ...grpc.CallOption) (*BubbleChartAppendResponse, error) {
	out := new(BubbleChartAppendResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.BubbleChart/Append", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bubbleChartClient) SwapById(ctx context.Context, in *BubbleChartSwapByIdRequest, opts ...grpc.CallOption) (*BubbleChartSwapByIdResponse, error) {
	out := new(BubbleChartSwapByIdResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.BubbleChart/SwapById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bubbleChartClient) SwapByRank(ctx context.Context, in *BubbleChartSwapByRankRequest, opts ...grpc.CallOption) (*BubbleChartSwapByRankResponse, error) {
	out := new(BubbleChartSwapByRankResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.BubbleChart/SwapByRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bubbleChartClient) GetRange(ctx context.Context, in *BubbleChartGetRangeRequest, opts ...grpc.CallOption) (*BubbleChartGetRangeResponse, error) {
	out := new(BubbleChartGetRangeResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.BubbleChart/GetRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bubbleChartClient) GetById(ctx context.Context, in *BubbleChartGetByIdRequest, opts ...grpc.CallOption) (*BubbleChartGetByIdResponse, error) {
	out := new(BubbleChartGetByIdResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.BubbleChart/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bubbleChartClient) RemoveById(ctx context.Context, in *BubbleChartRemoveByIdRequest, opts ...grpc.CallOption) (*BubbleChartRemoveByIdResponse, error) {
	out := new(BubbleChartRemoveByIdResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.BubbleChart/RemoveById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bubbleChartClient) SetInfo(ctx context.Context, in *BubbleChartSetInfoRequest, opts ...grpc.CallOption) (*BubbleChartSetInfoResponse, error) {
	out := new(BubbleChartSetInfoResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.BubbleChart/SetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BubbleChartServer is the server API for BubbleChart service.
// All implementations must embed UnimplementedBubbleChartServer
// for forward compatibility
type BubbleChartServer interface {
	// append entries to the end of chart
	Append(context.Context, *BubbleChartAppendRequest) (*BubbleChartAppendResponse, error)
	// swap 2 entries by id
	SwapById(context.Context, *BubbleChartSwapByIdRequest) (*BubbleChartSwapByIdResponse, error)
	// swap 2 entries by rank
	SwapByRank(context.Context, *BubbleChartSwapByRankRequest) (*BubbleChartSwapByRankResponse, error)
	// get entries by range
	GetRange(context.Context, *BubbleChartGetRangeRequest) (*BubbleChartGetRangeResponse, error)
	// get entries by id
	GetById(context.Context, *BubbleChartGetByIdRequest) (*BubbleChartGetByIdResponse, error)
	// remove entries by id
	RemoveById(context.Context, *BubbleChartRemoveByIdRequest) (*BubbleChartRemoveByIdResponse, error)
	// set entry info by id
	SetInfo(context.Context, *BubbleChartSetInfoRequest) (*BubbleChartSetInfoResponse, error)
	mustEmbedUnimplementedBubbleChartServer()
}

// UnimplementedBubbleChartServer must be embedded to have forward compatible implementations.
type UnimplementedBubbleChartServer struct {
}

func (UnimplementedBubbleChartServer) Append(context.Context, *BubbleChartAppendRequest) (*BubbleChartAppendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Append not implemented")
}
func (UnimplementedBubbleChartServer) SwapById(context.Context, *BubbleChartSwapByIdRequest) (*BubbleChartSwapByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapById not implemented")
}
func (UnimplementedBubbleChartServer) SwapByRank(context.Context, *BubbleChartSwapByRankRequest) (*BubbleChartSwapByRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapByRank not implemented")
}
func (UnimplementedBubbleChartServer) GetRange(context.Context, *BubbleChartGetRangeRequest) (*BubbleChartGetRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRange not implemented")
}
func (UnimplementedBubbleChartServer) GetById(context.Context, *BubbleChartGetByIdRequest) (*BubbleChartGetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBubbleChartServer) RemoveById(context.Context, *BubbleChartRemoveByIdRequest) (*BubbleChartRemoveByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveById not implemented")
}
func (UnimplementedBubbleChartServer) SetInfo(context.Context, *BubbleChartSetInfoRequest) (*BubbleChartSetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInfo not implemented")
}
func (UnimplementedBubbleChartServer) mustEmbedUnimplementedBubbleChartServer() {}

// UnsafeBubbleChartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BubbleChartServer will
// result in compilation errors.
type UnsafeBubbleChartServer interface {
	mustEmbedUnimplementedBubbleChartServer()
}

func RegisterBubbleChartServer(s grpc.ServiceRegistrar, srv BubbleChartServer) {
	s.RegisterService(&BubbleChart_ServiceDesc, srv)
}

func _BubbleChart_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BubbleChartAppendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BubbleChartServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.BubbleChart/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BubbleChartServer).Append(ctx, req.(*BubbleChartAppendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BubbleChart_SwapById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BubbleChartSwapByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BubbleChartServer).SwapById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.BubbleChart/SwapById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BubbleChartServer).SwapById(ctx, req.(*BubbleChartSwapByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BubbleChart_SwapByRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BubbleChartSwapByRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BubbleChartServer).SwapByRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.BubbleChart/SwapByRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BubbleChartServer).SwapByRank(ctx, req.(*BubbleChartSwapByRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BubbleChart_GetRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BubbleChartGetRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BubbleChartServer).GetRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.BubbleChart/GetRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BubbleChartServer).GetRange(ctx, req.(*BubbleChartGetRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BubbleChart_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BubbleChartGetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BubbleChartServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.BubbleChart/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BubbleChartServer).GetById(ctx, req.(*BubbleChartGetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BubbleChart_RemoveById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BubbleChartRemoveByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BubbleChartServer).RemoveById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.BubbleChart/RemoveById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BubbleChartServer).RemoveById(ctx, req.(*BubbleChartRemoveByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BubbleChart_SetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BubbleChartSetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BubbleChartServer).SetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.BubbleChart/SetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BubbleChartServer).SetInfo(ctx, req.(*BubbleChartSetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BubbleChart_ServiceDesc is the grpc.ServiceDesc for BubbleChart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BubbleChart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.BubbleChart",
	HandlerType: (*BubbleChartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Append",
			Handler:    _BubbleChart_Append_Handler,
		},
		{
			MethodName: "SwapById",
			Handler:    _BubbleChart_SwapById_Handler,
		},
		{
			MethodName: "SwapByRank",
			Handler:    _BubbleChart_SwapByRank_Handler,
		},
		{
			MethodName: "GetRange",
			Handler:    _BubbleChart_GetRange_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _BubbleChart_GetById_Handler,
		},
		{
			MethodName: "RemoveById",
			Handler:    _BubbleChart_RemoveById_Handler,
		},
		{
			MethodName: "SetInfo",
			Handler:    _BubbleChart_SetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libra/v1/bubble_chart.proto",
}
