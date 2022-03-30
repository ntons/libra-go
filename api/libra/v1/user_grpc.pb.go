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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// 查询详细信息
	Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserGetResponse, error)
	// 扩展查询
	Query(ctx context.Context, in *UserQueryRequest, opts ...grpc.CallOption) (*UserQueryResponse, error)
	// 封禁
	Ban(ctx context.Context, in *UserBanRequest, opts ...grpc.CallOption) (*UserBanResponse, error)
	// 扩展封禁
	Block(ctx context.Context, in *UserBlockRequest, opts ...grpc.CallOption) (*UserBlockResponse, error)
	// 绑定
	BindAcctId(ctx context.Context, in *UserBindAcctIdRequest, opts ...grpc.CallOption) (*UserBindAcctIdResponse, error)
	Login(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	Logout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error)
	Bind(ctx context.Context, in *UserBindRequest, opts ...grpc.CallOption) (*UserBindResponse, error)
	Unbind(ctx context.Context, in *UserUnbindRequest, opts ...grpc.CallOption) (*UserUnbindResponse, error)
	SetMetadata(ctx context.Context, in *UserSetMetadataRequest, opts ...grpc.CallOption) (*UserSetMetadataResponse, error)
	GetMetadata(ctx context.Context, in *UserGetMetadataRequest, opts ...grpc.CallOption) (*UserGetMetadataResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserGetResponse, error) {
	out := new(UserGetResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Query(ctx context.Context, in *UserQueryRequest, opts ...grpc.CallOption) (*UserQueryResponse, error) {
	out := new(UserQueryResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Ban(ctx context.Context, in *UserBanRequest, opts ...grpc.CallOption) (*UserBanResponse, error) {
	out := new(UserBanResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/Ban", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Block(ctx context.Context, in *UserBlockRequest, opts ...grpc.CallOption) (*UserBlockResponse, error) {
	out := new(UserBlockResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/Block", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) BindAcctId(ctx context.Context, in *UserBindAcctIdRequest, opts ...grpc.CallOption) (*UserBindAcctIdResponse, error) {
	out := new(UserBindAcctIdResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/BindAcctId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Logout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error) {
	out := new(UserLogoutResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Bind(ctx context.Context, in *UserBindRequest, opts ...grpc.CallOption) (*UserBindResponse, error) {
	out := new(UserBindResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/Bind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Unbind(ctx context.Context, in *UserUnbindRequest, opts ...grpc.CallOption) (*UserUnbindResponse, error) {
	out := new(UserUnbindResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/Unbind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetMetadata(ctx context.Context, in *UserSetMetadataRequest, opts ...grpc.CallOption) (*UserSetMetadataResponse, error) {
	out := new(UserSetMetadataResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/SetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetMetadata(ctx context.Context, in *UserGetMetadataRequest, opts ...grpc.CallOption) (*UserGetMetadataResponse, error) {
	out := new(UserGetMetadataResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.User/GetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// 查询详细信息
	Get(context.Context, *UserGetRequest) (*UserGetResponse, error)
	// 扩展查询
	Query(context.Context, *UserQueryRequest) (*UserQueryResponse, error)
	// 封禁
	Ban(context.Context, *UserBanRequest) (*UserBanResponse, error)
	// 扩展封禁
	Block(context.Context, *UserBlockRequest) (*UserBlockResponse, error)
	// 绑定
	BindAcctId(context.Context, *UserBindAcctIdRequest) (*UserBindAcctIdResponse, error)
	Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	Logout(context.Context, *UserLogoutRequest) (*UserLogoutResponse, error)
	Bind(context.Context, *UserBindRequest) (*UserBindResponse, error)
	Unbind(context.Context, *UserUnbindRequest) (*UserUnbindResponse, error)
	SetMetadata(context.Context, *UserSetMetadataRequest) (*UserSetMetadataResponse, error)
	GetMetadata(context.Context, *UserGetMetadataRequest) (*UserGetMetadataResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Get(context.Context, *UserGetRequest) (*UserGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServer) Query(context.Context, *UserQueryRequest) (*UserQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedUserServer) Ban(context.Context, *UserBanRequest) (*UserBanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ban not implemented")
}
func (UnimplementedUserServer) Block(context.Context, *UserBlockRequest) (*UserBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Block not implemented")
}
func (UnimplementedUserServer) BindAcctId(context.Context, *UserBindAcctIdRequest) (*UserBindAcctIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindAcctId not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Logout(context.Context, *UserLogoutRequest) (*UserLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServer) Bind(context.Context, *UserBindRequest) (*UserBindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bind not implemented")
}
func (UnimplementedUserServer) Unbind(context.Context, *UserUnbindRequest) (*UserUnbindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unbind not implemented")
}
func (UnimplementedUserServer) SetMetadata(context.Context, *UserSetMetadataRequest) (*UserSetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMetadata not implemented")
}
func (UnimplementedUserServer) GetMetadata(context.Context, *UserGetMetadataRequest) (*UserGetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Get(ctx, req.(*UserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Query(ctx, req.(*UserQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Ban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Ban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/Ban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Ban(ctx, req.(*UserBanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Block_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Block(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/Block",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Block(ctx, req.(*UserBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_BindAcctId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBindAcctIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).BindAcctId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/BindAcctId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).BindAcctId(ctx, req.(*UserBindAcctIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Logout(ctx, req.(*UserLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Bind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Bind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/Bind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Bind(ctx, req.(*UserBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Unbind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUnbindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Unbind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/Unbind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Unbind(ctx, req.(*UserUnbindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/SetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetMetadata(ctx, req.(*UserSetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.User/GetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetMetadata(ctx, req.(*UserGetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _User_Get_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _User_Query_Handler,
		},
		{
			MethodName: "Ban",
			Handler:    _User_Ban_Handler,
		},
		{
			MethodName: "Block",
			Handler:    _User_Block_Handler,
		},
		{
			MethodName: "BindAcctId",
			Handler:    _User_BindAcctId_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _User_Logout_Handler,
		},
		{
			MethodName: "Bind",
			Handler:    _User_Bind_Handler,
		},
		{
			MethodName: "Unbind",
			Handler:    _User_Unbind_Handler,
		},
		{
			MethodName: "SetMetadata",
			Handler:    _User_SetMetadata_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _User_GetMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libra/v1/user.proto",
}
