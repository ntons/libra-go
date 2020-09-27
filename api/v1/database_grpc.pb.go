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

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseClient interface {
	RegisterSchema(ctx context.Context, in *DatabaseRegisterSchemaRequest, opts ...grpc.CallOption) (*DatabaseRegisterSchemaResponse, error)
	// lock/unlock
	Lock(ctx context.Context, in *DatabaseLockRequest, opts ...grpc.CallOption) (*DatabaseLockResponse, error)
	Unlock(ctx context.Context, in *DatabaseUnlockRequest, opts ...grpc.CallOption) (*DatabaseUnlockResponse, error)
	// get/set data
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

func (c *databaseClient) Lock(ctx context.Context, in *DatabaseLockRequest, opts ...grpc.CallOption) (*DatabaseLockResponse, error) {
	out := new(DatabaseLockResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Database/Lock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Unlock(ctx context.Context, in *DatabaseUnlockRequest, opts ...grpc.CallOption) (*DatabaseUnlockResponse, error) {
	out := new(DatabaseUnlockResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Database/Unlock", in, out, opts...)
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
	// lock/unlock
	Lock(context.Context, *DatabaseLockRequest) (*DatabaseLockResponse, error)
	Unlock(context.Context, *DatabaseUnlockRequest) (*DatabaseUnlockResponse, error)
	// get/set data
	Get(context.Context, *DatabaseGetRequest) (*DatabaseGetResponse, error)
	Set(context.Context, *DatabaseSetRequest) (*DatabaseSetResponse, error)
	mustEmbedUnimplementedDatabaseServer()
}

// UnimplementedDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServer struct {
}

func (*UnimplementedDatabaseServer) RegisterSchema(context.Context, *DatabaseRegisterSchemaRequest) (*DatabaseRegisterSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSchema not implemented")
}
func (*UnimplementedDatabaseServer) Lock(context.Context, *DatabaseLockRequest) (*DatabaseLockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (*UnimplementedDatabaseServer) Unlock(context.Context, *DatabaseUnlockRequest) (*DatabaseUnlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlock not implemented")
}
func (*UnimplementedDatabaseServer) Get(context.Context, *DatabaseGetRequest) (*DatabaseGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDatabaseServer) Set(context.Context, *DatabaseSetRequest) (*DatabaseSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedDatabaseServer) mustEmbedUnimplementedDatabaseServer() {}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
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

func _Database_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Database/Lock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Lock(ctx, req.(*DatabaseLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Unlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseUnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Unlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Database/Unlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Unlock(ctx, req.(*DatabaseUnlockRequest))
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

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "libra.v1.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterSchema",
			Handler:    _Database_RegisterSchema_Handler,
		},
		{
			MethodName: "Lock",
			Handler:    _Database_Lock_Handler,
		},
		{
			MethodName: "Unlock",
			Handler:    _Database_Unlock_Handler,
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
	Metadata: "v1/database.proto",
}
