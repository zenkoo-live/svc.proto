// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: svc.infra.link/message.proto

package link

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	LinkMessage_SendGlobal_FullMethodName  = "/svc.infra.link.LinkMessage/SendGlobal"
	LinkMessage_SendGroup_FullMethodName   = "/svc.infra.link.LinkMessage/SendGroup"
	LinkMessage_SendAccount_FullMethodName = "/svc.infra.link.LinkMessage/SendAccount"
	LinkMessage_SendDevice_FullMethodName  = "/svc.infra.link.LinkMessage/SendDevice"
	LinkMessage_SendSession_FullMethodName = "/svc.infra.link.LinkMessage/SendSession"
)

// LinkMessageClient is the client API for LinkMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service of message
type LinkMessageClient interface {
	// 发送全局消息
	SendGlobal(ctx context.Context, in *SendGlobalRequest, opts ...grpc.CallOption) (*SendGlobalResponse, error)
	// 发送群组消息
	SendGroup(ctx context.Context, in *SendGroupRequest, opts ...grpc.CallOption) (*SendGroupResponse, error)
	// 发送账号消息（对单人)
	SendAccount(ctx context.Context, in *SendAccountRequest, opts ...grpc.CallOption) (*SendAccountResponse, error)
	// 发送设备消息
	SendDevice(ctx context.Context, in *SendDeviceRequest, opts ...grpc.CallOption) (*SendDeviceResponse, error)
	// 发送连接消息
	SendSession(ctx context.Context, in *SendSessionRequest, opts ...grpc.CallOption) (*SendSessionResponse, error)
}

type linkMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkMessageClient(cc grpc.ClientConnInterface) LinkMessageClient {
	return &linkMessageClient{cc}
}

func (c *linkMessageClient) SendGlobal(ctx context.Context, in *SendGlobalRequest, opts ...grpc.CallOption) (*SendGlobalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendGlobalResponse)
	err := c.cc.Invoke(ctx, LinkMessage_SendGlobal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkMessageClient) SendGroup(ctx context.Context, in *SendGroupRequest, opts ...grpc.CallOption) (*SendGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendGroupResponse)
	err := c.cc.Invoke(ctx, LinkMessage_SendGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkMessageClient) SendAccount(ctx context.Context, in *SendAccountRequest, opts ...grpc.CallOption) (*SendAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendAccountResponse)
	err := c.cc.Invoke(ctx, LinkMessage_SendAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkMessageClient) SendDevice(ctx context.Context, in *SendDeviceRequest, opts ...grpc.CallOption) (*SendDeviceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendDeviceResponse)
	err := c.cc.Invoke(ctx, LinkMessage_SendDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkMessageClient) SendSession(ctx context.Context, in *SendSessionRequest, opts ...grpc.CallOption) (*SendSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendSessionResponse)
	err := c.cc.Invoke(ctx, LinkMessage_SendSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkMessageServer is the server API for LinkMessage service.
// All implementations must embed UnimplementedLinkMessageServer
// for forward compatibility
//
// Service of message
type LinkMessageServer interface {
	// 发送全局消息
	SendGlobal(context.Context, *SendGlobalRequest) (*SendGlobalResponse, error)
	// 发送群组消息
	SendGroup(context.Context, *SendGroupRequest) (*SendGroupResponse, error)
	// 发送账号消息（对单人)
	SendAccount(context.Context, *SendAccountRequest) (*SendAccountResponse, error)
	// 发送设备消息
	SendDevice(context.Context, *SendDeviceRequest) (*SendDeviceResponse, error)
	// 发送连接消息
	SendSession(context.Context, *SendSessionRequest) (*SendSessionResponse, error)
	mustEmbedUnimplementedLinkMessageServer()
}

// UnimplementedLinkMessageServer must be embedded to have forward compatible implementations.
type UnimplementedLinkMessageServer struct {
}

func (UnimplementedLinkMessageServer) SendGlobal(context.Context, *SendGlobalRequest) (*SendGlobalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGlobal not implemented")
}
func (UnimplementedLinkMessageServer) SendGroup(context.Context, *SendGroupRequest) (*SendGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroup not implemented")
}
func (UnimplementedLinkMessageServer) SendAccount(context.Context, *SendAccountRequest) (*SendAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAccount not implemented")
}
func (UnimplementedLinkMessageServer) SendDevice(context.Context, *SendDeviceRequest) (*SendDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDevice not implemented")
}
func (UnimplementedLinkMessageServer) SendSession(context.Context, *SendSessionRequest) (*SendSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSession not implemented")
}
func (UnimplementedLinkMessageServer) mustEmbedUnimplementedLinkMessageServer() {}

// UnsafeLinkMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkMessageServer will
// result in compilation errors.
type UnsafeLinkMessageServer interface {
	mustEmbedUnimplementedLinkMessageServer()
}

func RegisterLinkMessageServer(s grpc.ServiceRegistrar, srv LinkMessageServer) {
	s.RegisterService(&LinkMessage_ServiceDesc, srv)
}

func _LinkMessage_SendGlobal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGlobalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkMessageServer).SendGlobal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkMessage_SendGlobal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkMessageServer).SendGlobal(ctx, req.(*SendGlobalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkMessage_SendGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkMessageServer).SendGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkMessage_SendGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkMessageServer).SendGroup(ctx, req.(*SendGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkMessage_SendAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkMessageServer).SendAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkMessage_SendAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkMessageServer).SendAccount(ctx, req.(*SendAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkMessage_SendDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkMessageServer).SendDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkMessage_SendDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkMessageServer).SendDevice(ctx, req.(*SendDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkMessage_SendSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkMessageServer).SendSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkMessage_SendSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkMessageServer).SendSession(ctx, req.(*SendSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkMessage_ServiceDesc is the grpc.ServiceDesc for LinkMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.infra.link.LinkMessage",
	HandlerType: (*LinkMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendGlobal",
			Handler:    _LinkMessage_SendGlobal_Handler,
		},
		{
			MethodName: "SendGroup",
			Handler:    _LinkMessage_SendGroup_Handler,
		},
		{
			MethodName: "SendAccount",
			Handler:    _LinkMessage_SendAccount_Handler,
		},
		{
			MethodName: "SendDevice",
			Handler:    _LinkMessage_SendDevice_Handler,
		},
		{
			MethodName: "SendSession",
			Handler:    _LinkMessage_SendSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.infra.link/message.proto",
}
