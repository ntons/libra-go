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

// MailboxClient is the client API for Mailbox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailboxClient interface {
	List(ctx context.Context, in *MailboxListRequest, opts ...grpc.CallOption) (*MailboxListResponse, error)
	Push(ctx context.Context, in *MailboxPushRequest, opts ...grpc.CallOption) (*MailboxPushResponse, error)
	Pull(ctx context.Context, in *MailboxPullRequest, opts ...grpc.CallOption) (*MailboxPullResponse, error)
	// 支持批量收件人
	Send(ctx context.Context, in *MailboxSendRequest, opts ...grpc.CallOption) (*MailboxSendResponse, error)
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

func (c *mailboxClient) Send(ctx context.Context, in *MailboxSendRequest, opts ...grpc.CallOption) (*MailboxSendResponse, error) {
	out := new(MailboxSendResponse)
	err := c.cc.Invoke(ctx, "/libra.v1.Mailbox/Send", in, out, opts...)
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
	// 支持批量收件人
	Send(context.Context, *MailboxSendRequest) (*MailboxSendResponse, error)
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
func (UnimplementedMailboxServer) Send(context.Context, *MailboxSendRequest) (*MailboxSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
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

func _Mailbox_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailboxSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailboxServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libra.v1.Mailbox/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailboxServer).Send(ctx, req.(*MailboxSendRequest))
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
		{
			MethodName: "Send",
			Handler:    _Mailbox_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libra/v1/mailbox.proto",
}
