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

// LeaderboardClient is the client API for Leaderboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeaderboardClient interface {
	// set entry score by id
	// set entry info if info field not empty
	SetScore(ctx context.Context, in *LeaderboardSetScoreRequest, opts ...grpc.CallOption) (*LeaderboardSetScoreResponse, error)
	// incr entry score by id, if id not exists, add it with score
	// set entry info if info filed not empty
	IncrScore(ctx context.Context, in *LeaderboardIncrScoreRequest, opts ...grpc.CallOption) (*LeaderboardIncrScoreResponse, error)
	// get entries by range
	GetRange(ctx context.Context, in *LeaderboardGetRangeRequest, opts ...grpc.CallOption) (*LeaderboardGetRangeResponse, error)
	// get entries by id
	GetById(ctx context.Context, in *LeaderboardGetByIdRequest, opts ...grpc.CallOption) (*LeaderboardGetByIdResponse, error)
	// remove entries by id
	RemoveById(ctx context.Context, in *LeaderboardRemoveByIdRequest, opts ...grpc.CallOption) (*LeaderboardRemoveByIdResponse, error)
	// set entry info by id
	SetInfo(ctx context.Context, in *LeaderboardSetInfoRequest, opts ...grpc.CallOption) (*LeaderboardSetInfoResponse, error)
}

type leaderboardClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaderboardClient(cc grpc.ClientConnInterface) LeaderboardClient {
	return &leaderboardClient{cc}
}

func (c *leaderboardClient) SetScore(ctx context.Context, in *LeaderboardSetScoreRequest, opts ...grpc.CallOption) (*LeaderboardSetScoreResponse, error) {
	out := new(LeaderboardSetScoreResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/SetScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderboardClient) IncrScore(ctx context.Context, in *LeaderboardIncrScoreRequest, opts ...grpc.CallOption) (*LeaderboardIncrScoreResponse, error) {
	out := new(LeaderboardIncrScoreResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/IncrScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderboardClient) GetRange(ctx context.Context, in *LeaderboardGetRangeRequest, opts ...grpc.CallOption) (*LeaderboardGetRangeResponse, error) {
	out := new(LeaderboardGetRangeResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/GetRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderboardClient) GetById(ctx context.Context, in *LeaderboardGetByIdRequest, opts ...grpc.CallOption) (*LeaderboardGetByIdResponse, error) {
	out := new(LeaderboardGetByIdResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderboardClient) RemoveById(ctx context.Context, in *LeaderboardRemoveByIdRequest, opts ...grpc.CallOption) (*LeaderboardRemoveByIdResponse, error) {
	out := new(LeaderboardRemoveByIdResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/RemoveById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderboardClient) SetInfo(ctx context.Context, in *LeaderboardSetInfoRequest, opts ...grpc.CallOption) (*LeaderboardSetInfoResponse, error) {
	out := new(LeaderboardSetInfoResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/SetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaderboardServer is the server API for Leaderboard service.
// All implementations must embed UnimplementedLeaderboardServer
// for forward compatibility
type LeaderboardServer interface {
	// set entry score by id
	// set entry info if info field not empty
	SetScore(context.Context, *LeaderboardSetScoreRequest) (*LeaderboardSetScoreResponse, error)
	// incr entry score by id, if id not exists, add it with score
	// set entry info if info filed not empty
	IncrScore(context.Context, *LeaderboardIncrScoreRequest) (*LeaderboardIncrScoreResponse, error)
	// get entries by range
	GetRange(context.Context, *LeaderboardGetRangeRequest) (*LeaderboardGetRangeResponse, error)
	// get entries by id
	GetById(context.Context, *LeaderboardGetByIdRequest) (*LeaderboardGetByIdResponse, error)
	// remove entries by id
	RemoveById(context.Context, *LeaderboardRemoveByIdRequest) (*LeaderboardRemoveByIdResponse, error)
	// set entry info by id
	SetInfo(context.Context, *LeaderboardSetInfoRequest) (*LeaderboardSetInfoResponse, error)
	mustEmbedUnimplementedLeaderboardServer()
}

// UnimplementedLeaderboardServer must be embedded to have forward compatible implementations.
type UnimplementedLeaderboardServer struct {
}

func (UnimplementedLeaderboardServer) SetScore(context.Context, *LeaderboardSetScoreRequest) (*LeaderboardSetScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetScore not implemented")
}
func (UnimplementedLeaderboardServer) IncrScore(context.Context, *LeaderboardIncrScoreRequest) (*LeaderboardIncrScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrScore not implemented")
}
func (UnimplementedLeaderboardServer) GetRange(context.Context, *LeaderboardGetRangeRequest) (*LeaderboardGetRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRange not implemented")
}
func (UnimplementedLeaderboardServer) GetById(context.Context, *LeaderboardGetByIdRequest) (*LeaderboardGetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedLeaderboardServer) RemoveById(context.Context, *LeaderboardRemoveByIdRequest) (*LeaderboardRemoveByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveById not implemented")
}
func (UnimplementedLeaderboardServer) SetInfo(context.Context, *LeaderboardSetInfoRequest) (*LeaderboardSetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInfo not implemented")
}
func (UnimplementedLeaderboardServer) mustEmbedUnimplementedLeaderboardServer() {}

// UnsafeLeaderboardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeaderboardServer will
// result in compilation errors.
type UnsafeLeaderboardServer interface {
	mustEmbedUnimplementedLeaderboardServer()
}

func RegisterLeaderboardServer(s grpc.ServiceRegistrar, srv LeaderboardServer) {
	s.RegisterService(&Leaderboard_ServiceDesc, srv)
}

func _Leaderboard_SetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardSetScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).SetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/SetScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).SetScore(ctx, req.(*LeaderboardSetScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leaderboard_IncrScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardIncrScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).IncrScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/IncrScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).IncrScore(ctx, req.(*LeaderboardIncrScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leaderboard_GetRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardGetRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).GetRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/GetRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).GetRange(ctx, req.(*LeaderboardGetRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leaderboard_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardGetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).GetById(ctx, req.(*LeaderboardGetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leaderboard_RemoveById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardRemoveByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).RemoveById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/RemoveById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).RemoveById(ctx, req.(*LeaderboardRemoveByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leaderboard_SetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardSetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).SetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/SetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).SetInfo(ctx, req.(*LeaderboardSetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Leaderboard_ServiceDesc is the grpc.ServiceDesc for Leaderboard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Leaderboard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.Leaderboard",
	HandlerType: (*LeaderboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetScore",
			Handler:    _Leaderboard_SetScore_Handler,
		},
		{
			MethodName: "IncrScore",
			Handler:    _Leaderboard_IncrScore_Handler,
		},
		{
			MethodName: "GetRange",
			Handler:    _Leaderboard_GetRange_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _Leaderboard_GetById_Handler,
		},
		{
			MethodName: "RemoveById",
			Handler:    _Leaderboard_RemoveById_Handler,
		},
		{
			MethodName: "SetInfo",
			Handler:    _Leaderboard_SetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/leaderboard.proto",
}
