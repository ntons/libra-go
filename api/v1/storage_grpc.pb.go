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

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseClient interface {
	RegisterSchema(ctx context.Context, in *DatabaseRegisterSchemaRequest, opts ...grpc.CallOption) (*DatabaseRegisterSchemaResponse, error)
	Get(ctx context.Context, in *DatabaseGetRequest, opts ...grpc.CallOption) (*DatabaseGetResponse, error)
	Set(ctx context.Context, in *DatabaseSetRequest, opts ...grpc.CallOption) (*DatabaseSetResponse, error)
}

type databaseClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseClient(cc grpc.ClientConnInterface) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) RegisterSchema(ctx context.Context, in *DatabaseRegisterSchemaRequest, opts ...grpc.CallOption) (*DatabaseRegisterSchemaResponse, error) {
	out := new(DatabaseRegisterSchemaResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Database/RegisterSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Get(ctx context.Context, in *DatabaseGetRequest, opts ...grpc.CallOption) (*DatabaseGetResponse, error) {
	out := new(DatabaseGetResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Database/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Set(ctx context.Context, in *DatabaseSetRequest, opts ...grpc.CallOption) (*DatabaseSetResponse, error) {
	out := new(DatabaseSetResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Database/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
// All implementations must embed UnimplementedDatabaseServer
// for forward compatibility
type DatabaseServer interface {
	RegisterSchema(context.Context, *DatabaseRegisterSchemaRequest) (*DatabaseRegisterSchemaResponse, error)
	Get(context.Context, *DatabaseGetRequest) (*DatabaseGetResponse, error)
	Set(context.Context, *DatabaseSetRequest) (*DatabaseSetResponse, error)
	mustEmbedUnimplementedDatabaseServer()
}

// UnimplementedDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServer struct {
}

func (UnimplementedDatabaseServer) RegisterSchema(context.Context, *DatabaseRegisterSchemaRequest) (*DatabaseRegisterSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSchema not implemented")
}
func (UnimplementedDatabaseServer) Get(context.Context, *DatabaseGetRequest) (*DatabaseGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDatabaseServer) Set(context.Context, *DatabaseSetRequest) (*DatabaseSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedDatabaseServer) mustEmbedUnimplementedDatabaseServer() {}

// UnsafeDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServer will
// result in compilation errors.
type UnsafeDatabaseServer interface {
	mustEmbedUnimplementedDatabaseServer()
}

func RegisterDatabaseServer(s grpc.ServiceRegistrar, srv DatabaseServer) {
	s.RegisterService(&Database_ServiceDesc, srv)
}

func _Database_RegisterSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseRegisterSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).RegisterSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Database/RegisterSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).RegisterSchema(ctx, req.(*DatabaseRegisterSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Database/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Get(ctx, req.(*DatabaseGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Database/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Set(ctx, req.(*DatabaseSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Database_ServiceDesc is the grpc.ServiceDesc for Database service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Database_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterSchema",
			Handler:    _Database_RegisterSchema_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Database_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Database_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/storage.proto",
}

// MailboxClient is the client API for Mailbox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailboxClient interface {
	List(ctx context.Context, in *MailboxListRequest, opts ...grpc.CallOption) (*MailboxListResponse, error)
	Push(ctx context.Context, in *MailboxPushRequest, opts ...grpc.CallOption) (*MailboxPushResponse, error)
	Pull(ctx context.Context, in *MailboxPullRequest, opts ...grpc.CallOption) (*MailboxPullResponse, error)
}

type mailboxClient struct {
	cc grpc.ClientConnInterface
}

func NewMailboxClient(cc grpc.ClientConnInterface) MailboxClient {
	return &mailboxClient{cc}
}

func (c *mailboxClient) List(ctx context.Context, in *MailboxListRequest, opts ...grpc.CallOption) (*MailboxListResponse, error) {
	out := new(MailboxListResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Mailbox/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailboxClient) Push(ctx context.Context, in *MailboxPushRequest, opts ...grpc.CallOption) (*MailboxPushResponse, error) {
	out := new(MailboxPushResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Mailbox/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailboxClient) Pull(ctx context.Context, in *MailboxPullRequest, opts ...grpc.CallOption) (*MailboxPullResponse, error) {
	out := new(MailboxPullResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Mailbox/Pull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailboxServer is the server API for Mailbox service.
// All implementations must embed UnimplementedMailboxServer
// for forward compatibility
type MailboxServer interface {
	List(context.Context, *MailboxListRequest) (*MailboxListResponse, error)
	Push(context.Context, *MailboxPushRequest) (*MailboxPushResponse, error)
	Pull(context.Context, *MailboxPullRequest) (*MailboxPullResponse, error)
	mustEmbedUnimplementedMailboxServer()
}

// UnimplementedMailboxServer must be embedded to have forward compatible implementations.
type UnimplementedMailboxServer struct {
}

func (UnimplementedMailboxServer) List(context.Context, *MailboxListRequest) (*MailboxListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMailboxServer) Push(context.Context, *MailboxPushRequest) (*MailboxPushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedMailboxServer) Pull(context.Context, *MailboxPullRequest) (*MailboxPullResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedMailboxServer) mustEmbedUnimplementedMailboxServer() {}

// UnsafeMailboxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailboxServer will
// result in compilation errors.
type UnsafeMailboxServer interface {
	mustEmbedUnimplementedMailboxServer()
}

func RegisterMailboxServer(s grpc.ServiceRegistrar, srv MailboxServer) {
	s.RegisterService(&Mailbox_ServiceDesc, srv)
}

func _Mailbox_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailboxListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailboxServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Mailbox/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailboxServer).List(ctx, req.(*MailboxListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mailbox_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailboxPushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailboxServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Mailbox/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailboxServer).Push(ctx, req.(*MailboxPushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mailbox_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailboxPullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailboxServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Mailbox/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailboxServer).Pull(ctx, req.(*MailboxPullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mailbox_ServiceDesc is the grpc.ServiceDesc for Mailbox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mailbox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.Mailbox",
	HandlerType: (*MailboxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Mailbox_List_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Mailbox_Push_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _Mailbox_Pull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/storage.proto",
}
