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
const _ = grpc.SupportPackageIsVersion6

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	// Login verify the login state, return user and token as passport
	// a user will be created automatically if not exists
	Login(ctx context.Context, in *AccountLoginRequest, opts ...grpc.CallOption) (*AccountLoginResponse, error)
	// Bind another acct_id to user
	Bind(ctx context.Context, in *AccountBindRequest, opts ...grpc.CallOption) (*AccountBindResponse, error)
	// List roles which belonging to the user
	ListRoles(ctx context.Context, in *AccountListRolesRequest, opts ...grpc.CallOption) (*AccountListRolesResponse, error)
	// Create role to the user
	CreateRole(ctx context.Context, in *AccountCreateRoleRequest, opts ...grpc.CallOption) (*AccountCreateRoleResponse, error)
	// Sign in the role, get a ticket as passport
	SignIn(ctx context.Context, in *AccountSignInRequest, opts ...grpc.CallOption) (*AccountSignInResponse, error)
	// Set metadata of the user
	SetUserMetadata(ctx context.Context, in *AccountSetUserMetadataRequest, opts ...grpc.CallOption) (*AccountSetUserMetadataResponse, error)
	// Set metadata of a role
	SetRoleMetadata(ctx context.Context, in *AccountSetRoleMetadataRequest, opts ...grpc.CallOption) (*AccountSetRoleMetadataResponse, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Login(ctx context.Context, in *AccountLoginRequest, opts ...grpc.CallOption) (*AccountLoginResponse, error) {
	out := new(AccountLoginResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Account/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Bind(ctx context.Context, in *AccountBindRequest, opts ...grpc.CallOption) (*AccountBindResponse, error) {
	out := new(AccountBindResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Account/Bind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) ListRoles(ctx context.Context, in *AccountListRolesRequest, opts ...grpc.CallOption) (*AccountListRolesResponse, error) {
	out := new(AccountListRolesResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Account/ListRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) CreateRole(ctx context.Context, in *AccountCreateRoleRequest, opts ...grpc.CallOption) (*AccountCreateRoleResponse, error) {
	out := new(AccountCreateRoleResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Account/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) SignIn(ctx context.Context, in *AccountSignInRequest, opts ...grpc.CallOption) (*AccountSignInResponse, error) {
	out := new(AccountSignInResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Account/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) SetUserMetadata(ctx context.Context, in *AccountSetUserMetadataRequest, opts ...grpc.CallOption) (*AccountSetUserMetadataResponse, error) {
	out := new(AccountSetUserMetadataResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Account/SetUserMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) SetRoleMetadata(ctx context.Context, in *AccountSetRoleMetadataRequest, opts ...grpc.CallOption) (*AccountSetRoleMetadataResponse, error) {
	out := new(AccountSetRoleMetadataResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Account/SetRoleMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	// Login verify the login state, return user and token as passport
	// a user will be created automatically if not exists
	Login(context.Context, *AccountLoginRequest) (*AccountLoginResponse, error)
	// Bind another acct_id to user
	Bind(context.Context, *AccountBindRequest) (*AccountBindResponse, error)
	// List roles which belonging to the user
	ListRoles(context.Context, *AccountListRolesRequest) (*AccountListRolesResponse, error)
	// Create role to the user
	CreateRole(context.Context, *AccountCreateRoleRequest) (*AccountCreateRoleResponse, error)
	// Sign in the role, get a ticket as passport
	SignIn(context.Context, *AccountSignInRequest) (*AccountSignInResponse, error)
	// Set metadata of the user
	SetUserMetadata(context.Context, *AccountSetUserMetadataRequest) (*AccountSetUserMetadataResponse, error)
	// Set metadata of a role
	SetRoleMetadata(context.Context, *AccountSetRoleMetadataRequest) (*AccountSetRoleMetadataResponse, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) Login(context.Context, *AccountLoginRequest) (*AccountLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAccountServer) Bind(context.Context, *AccountBindRequest) (*AccountBindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bind not implemented")
}
func (*UnimplementedAccountServer) ListRoles(context.Context, *AccountListRolesRequest) (*AccountListRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (*UnimplementedAccountServer) CreateRole(context.Context, *AccountCreateRoleRequest) (*AccountCreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (*UnimplementedAccountServer) SignIn(context.Context, *AccountSignInRequest) (*AccountSignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedAccountServer) SetUserMetadata(context.Context, *AccountSetUserMetadataRequest) (*AccountSetUserMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserMetadata not implemented")
}
func (*UnimplementedAccountServer) SetRoleMetadata(context.Context, *AccountSetRoleMetadataRequest) (*AccountSetRoleMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoleMetadata not implemented")
}
func (*UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*AccountLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Bind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Bind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Account/Bind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Bind(ctx, req.(*AccountBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Account/ListRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).ListRoles(ctx, req.(*AccountListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Account/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).CreateRole(ctx, req.(*AccountCreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountSignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Account/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SignIn(ctx, req.(*AccountSignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_SetUserMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountSetUserMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SetUserMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Account/SetUserMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SetUserMetadata(ctx, req.(*AccountSetUserMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_SetRoleMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountSetRoleMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SetRoleMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Account/SetRoleMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SetRoleMetadata(ctx, req.(*AccountSetRoleMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
		{
			MethodName: "Bind",
			Handler:    _Account_Bind_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _Account_ListRoles_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _Account_CreateRole_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Account_SignIn_Handler,
		},
		{
			MethodName: "SetUserMetadata",
			Handler:    _Account_SetUserMetadata_Handler,
		},
		{
			MethodName: "SetRoleMetadata",
			Handler:    _Account_SetRoleMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/account.proto",
}