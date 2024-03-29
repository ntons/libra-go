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
	// set entries
	Set(ctx context.Context, in *LeaderboardSetRequest, opts ...grpc.CallOption) (*LeaderboardSetResponse, error)
	// set entry info by id
	SetInfo(ctx context.Context, in *LeaderboardSetInfoRequest, opts ...grpc.CallOption) (*LeaderboardSetInfoResponse, error)
	// get entries by rank
	GetByRank(ctx context.Context, in *LeaderboardGetByRankRequest, opts ...grpc.CallOption) (*LeaderboardGetByRankResponse, error)
	// get entries by id
	GetById(ctx context.Context, in *LeaderboardGetByIdRequest, opts ...grpc.CallOption) (*LeaderboardGetByIdResponse, error)
	// random entries by score
	GetByScore(ctx context.Context, in *LeaderboardGetByScoreRequest, opts ...grpc.CallOption) (*LeaderboardGetByScoreResponse, error)
	// remove entries by id
	RemoveById(ctx context.Context, in *LeaderboardRemoveByIdRequest, opts ...grpc.CallOption) (*LeaderboardRemoveByIdResponse, error)
	// only update metadata of chart
	Touch(ctx context.Context, in *LeaderboardTouchRequest, opts ...grpc.CallOption) (*LeaderboardTouchResponse, error)
	// deprecated, will be removed later
	SetScore(ctx context.Context, in *LeaderboardSetScoreRequest, opts ...grpc.CallOption) (*LeaderboardSetScoreResponse, error)
	// deprecated, will be removed later
	GetRange(ctx context.Context, in *LeaderboardGetRangeRequest, opts ...grpc.CallOption) (*LeaderboardGetRangeResponse, error)
}

type leaderboardClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaderboardClient(cc grpc.ClientConnInterface) LeaderboardClient {
	return &leaderboardClient{cc}
}

func (c *leaderboardClient) Set(ctx context.Context, in *LeaderboardSetRequest, opts ...grpc.CallOption) (*LeaderboardSetResponse, error) {
	out := new(LeaderboardSetResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/Set", in, out, opts...)
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

func (c *leaderboardClient) GetByRank(ctx context.Context, in *LeaderboardGetByRankRequest, opts ...grpc.CallOption) (*LeaderboardGetByRankResponse, error) {
	out := new(LeaderboardGetByRankResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/GetByRank", in, out, opts...)
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

func (c *leaderboardClient) GetByScore(ctx context.Context, in *LeaderboardGetByScoreRequest, opts ...grpc.CallOption) (*LeaderboardGetByScoreResponse, error) {
	out := new(LeaderboardGetByScoreResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/GetByScore", in, out, opts...)
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

func (c *leaderboardClient) Touch(ctx context.Context, in *LeaderboardTouchRequest, opts ...grpc.CallOption) (*LeaderboardTouchResponse, error) {
	out := new(LeaderboardTouchResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/Touch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderboardClient) SetScore(ctx context.Context, in *LeaderboardSetScoreRequest, opts ...grpc.CallOption) (*LeaderboardSetScoreResponse, error) {
	out := new(LeaderboardSetScoreResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Leaderboard/SetScore", in, out, opts...)
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

// LeaderboardServer is the server API for Leaderboard service.
// All implementations must embed UnimplementedLeaderboardServer
// for forward compatibility
type LeaderboardServer interface {
	// set entries
	Set(context.Context, *LeaderboardSetRequest) (*LeaderboardSetResponse, error)
	// set entry info by id
	SetInfo(context.Context, *LeaderboardSetInfoRequest) (*LeaderboardSetInfoResponse, error)
	// get entries by rank
	GetByRank(context.Context, *LeaderboardGetByRankRequest) (*LeaderboardGetByRankResponse, error)
	// get entries by id
	GetById(context.Context, *LeaderboardGetByIdRequest) (*LeaderboardGetByIdResponse, error)
	// random entries by score
	GetByScore(context.Context, *LeaderboardGetByScoreRequest) (*LeaderboardGetByScoreResponse, error)
	// remove entries by id
	RemoveById(context.Context, *LeaderboardRemoveByIdRequest) (*LeaderboardRemoveByIdResponse, error)
	// only update metadata of chart
	Touch(context.Context, *LeaderboardTouchRequest) (*LeaderboardTouchResponse, error)
	// deprecated, will be removed later
	SetScore(context.Context, *LeaderboardSetScoreRequest) (*LeaderboardSetScoreResponse, error)
	// deprecated, will be removed later
	GetRange(context.Context, *LeaderboardGetRangeRequest) (*LeaderboardGetRangeResponse, error)
	mustEmbedUnimplementedLeaderboardServer()
}

// UnimplementedLeaderboardServer must be embedded to have forward compatible implementations.
type UnimplementedLeaderboardServer struct {
}

func (UnimplementedLeaderboardServer) Set(context.Context, *LeaderboardSetRequest) (*LeaderboardSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedLeaderboardServer) SetInfo(context.Context, *LeaderboardSetInfoRequest) (*LeaderboardSetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInfo not implemented")
}
func (UnimplementedLeaderboardServer) GetByRank(context.Context, *LeaderboardGetByRankRequest) (*LeaderboardGetByRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByRank not implemented")
}
func (UnimplementedLeaderboardServer) GetById(context.Context, *LeaderboardGetByIdRequest) (*LeaderboardGetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedLeaderboardServer) GetByScore(context.Context, *LeaderboardGetByScoreRequest) (*LeaderboardGetByScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByScore not implemented")
}
func (UnimplementedLeaderboardServer) RemoveById(context.Context, *LeaderboardRemoveByIdRequest) (*LeaderboardRemoveByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveById not implemented")
}
func (UnimplementedLeaderboardServer) Touch(context.Context, *LeaderboardTouchRequest) (*LeaderboardTouchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Touch not implemented")
}
func (UnimplementedLeaderboardServer) SetScore(context.Context, *LeaderboardSetScoreRequest) (*LeaderboardSetScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetScore not implemented")
}
func (UnimplementedLeaderboardServer) GetRange(context.Context, *LeaderboardGetRangeRequest) (*LeaderboardGetRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRange not implemented")
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

func _Leaderboard_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).Set(ctx, req.(*LeaderboardSetRequest))
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

func _Leaderboard_GetByRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardGetByRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).GetByRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/GetByRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).GetByRank(ctx, req.(*LeaderboardGetByRankRequest))
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

func _Leaderboard_GetByScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardGetByScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).GetByScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/GetByScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).GetByScore(ctx, req.(*LeaderboardGetByScoreRequest))
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

func _Leaderboard_Touch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardTouchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServer).Touch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Leaderboard/Touch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServer).Touch(ctx, req.(*LeaderboardTouchRequest))
	}
	return interceptor(ctx, in, info, handler)
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

// Leaderboard_ServiceDesc is the grpc.ServiceDesc for Leaderboard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Leaderboard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.Leaderboard",
	HandlerType: (*LeaderboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Leaderboard_Set_Handler,
		},
		{
			MethodName: "SetInfo",
			Handler:    _Leaderboard_SetInfo_Handler,
		},
		{
			MethodName: "GetByRank",
			Handler:    _Leaderboard_GetByRank_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _Leaderboard_GetById_Handler,
		},
		{
			MethodName: "GetByScore",
			Handler:    _Leaderboard_GetByScore_Handler,
		},
		{
			MethodName: "RemoveById",
			Handler:    _Leaderboard_RemoveById_Handler,
		},
		{
			MethodName: "Touch",
			Handler:    _Leaderboard_Touch_Handler,
		},
		{
			MethodName: "SetScore",
			Handler:    _Leaderboard_SetScore_Handler,
		},
		{
			MethodName: "GetRange",
			Handler:    _Leaderboard_GetRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libra/v1/leaderboard.proto",
}
