// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: svc.infra.link/instruction.proto

package link

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LinkInstruction_RemoveSession_FullMethodName = "/svc.infra.link.LinkInstruction/RemoveSession"
	LinkInstruction_RemoveAccount_FullMethodName = "/svc.infra.link.LinkInstruction/RemoveAccount"
	LinkInstruction_RemoveDevice_FullMethodName  = "/svc.infra.link.LinkInstruction/RemoveDevice"
)

// LinkInstructionClient is the client API for LinkInstruction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service of instruction
type LinkInstructionClient interface {
	// 移除（踢）连接
	RemoveSession(ctx context.Context, in *RemoveSessionRequest, opts ...grpc.CallOption) (*RemoveSessionResponse, error)
	// 移除（踢）账号
	RemoveAccount(ctx context.Context, in *RemoveAccountRequest, opts ...grpc.CallOption) (*RemoveAccountResponse, error)
	// 移除（踢）设备
	RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error)
}

type linkInstructionClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkInstructionClient(cc grpc.ClientConnInterface) LinkInstructionClient {
	return &linkInstructionClient{cc}
}

func (c *linkInstructionClient) RemoveSession(ctx context.Context, in *RemoveSessionRequest, opts ...grpc.CallOption) (*RemoveSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveSessionResponse)
	err := c.cc.Invoke(ctx, LinkInstruction_RemoveSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkInstructionClient) RemoveAccount(ctx context.Context, in *RemoveAccountRequest, opts ...grpc.CallOption) (*RemoveAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveAccountResponse)
	err := c.cc.Invoke(ctx, LinkInstruction_RemoveAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkInstructionClient) RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveDeviceResponse)
	err := c.cc.Invoke(ctx, LinkInstruction_RemoveDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkInstructionServer is the server API for LinkInstruction service.
// All implementations must embed UnimplementedLinkInstructionServer
// for forward compatibility.
//
// Service of instruction
type LinkInstructionServer interface {
	// 移除（踢）连接
	RemoveSession(context.Context, *RemoveSessionRequest) (*RemoveSessionResponse, error)
	// 移除（踢）账号
	RemoveAccount(context.Context, *RemoveAccountRequest) (*RemoveAccountResponse, error)
	// 移除（踢）设备
	RemoveDevice(context.Context, *RemoveDeviceRequest) (*RemoveDeviceResponse, error)
	mustEmbedUnimplementedLinkInstructionServer()
}

// UnimplementedLinkInstructionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLinkInstructionServer struct{}

func (UnimplementedLinkInstructionServer) RemoveSession(context.Context, *RemoveSessionRequest) (*RemoveSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSession not implemented")
}
func (UnimplementedLinkInstructionServer) RemoveAccount(context.Context, *RemoveAccountRequest) (*RemoveAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAccount not implemented")
}
func (UnimplementedLinkInstructionServer) RemoveDevice(context.Context, *RemoveDeviceRequest) (*RemoveDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDevice not implemented")
}
func (UnimplementedLinkInstructionServer) mustEmbedUnimplementedLinkInstructionServer() {}
func (UnimplementedLinkInstructionServer) testEmbeddedByValue()                         {}

// UnsafeLinkInstructionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkInstructionServer will
// result in compilation errors.
type UnsafeLinkInstructionServer interface {
	mustEmbedUnimplementedLinkInstructionServer()
}

func RegisterLinkInstructionServer(s grpc.ServiceRegistrar, srv LinkInstructionServer) {
	// If the following call pancis, it indicates UnimplementedLinkInstructionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LinkInstruction_ServiceDesc, srv)
}

func _LinkInstruction_RemoveSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkInstructionServer).RemoveSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkInstruction_RemoveSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkInstructionServer).RemoveSession(ctx, req.(*RemoveSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkInstruction_RemoveAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkInstructionServer).RemoveAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkInstruction_RemoveAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkInstructionServer).RemoveAccount(ctx, req.(*RemoveAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkInstruction_RemoveDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkInstructionServer).RemoveDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkInstruction_RemoveDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkInstructionServer).RemoveDevice(ctx, req.(*RemoveDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkInstruction_ServiceDesc is the grpc.ServiceDesc for LinkInstruction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkInstruction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.infra.link.LinkInstruction",
	HandlerType: (*LinkInstructionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoveSession",
			Handler:    _LinkInstruction_RemoveSession_Handler,
		},
		{
			MethodName: "RemoveAccount",
			Handler:    _LinkInstruction_RemoveAccount_Handler,
		},
		{
			MethodName: "RemoveDevice",
			Handler:    _LinkInstruction_RemoveDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.infra.link/instruction.proto",
}
