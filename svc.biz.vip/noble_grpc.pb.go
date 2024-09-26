// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: svc.biz.vip/noble.proto

package vip

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
	Noble_CreateNoble_FullMethodName  = "/svc.biz.vip.Noble/CreateNoble"
	Noble_GetNoble_FullMethodName     = "/svc.biz.vip.Noble/GetNoble"
	Noble_UpdateNoble_FullMethodName  = "/svc.biz.vip.Noble/UpdateNoble"
	Noble_GetNobleList_FullMethodName = "/svc.biz.vip.Noble/GetNobleList"
)

// NobleClient is the client API for Noble service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 贵族
type NobleClient interface {
	// CreateNoble 创建
	CreateNoble(ctx context.Context, in *CreateNobleReq, opts ...grpc.CallOption) (*CreateNobleResp, error)
	// GetNoble 查询
	GetNoble(ctx context.Context, in *GetNobleReq, opts ...grpc.CallOption) (*GetNobleResp, error)
	// UpdateNoble 查询
	UpdateNoble(ctx context.Context, in *UpdateNobleReq, opts ...grpc.CallOption) (*UpdateNobleResp, error)
	// GetNobleList 查询列表
	GetNobleList(ctx context.Context, in *GetNobleListReq, opts ...grpc.CallOption) (*GetNobleListResp, error)
}

type nobleClient struct {
	cc grpc.ClientConnInterface
}

func NewNobleClient(cc grpc.ClientConnInterface) NobleClient {
	return &nobleClient{cc}
}

func (c *nobleClient) CreateNoble(ctx context.Context, in *CreateNobleReq, opts ...grpc.CallOption) (*CreateNobleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNobleResp)
	err := c.cc.Invoke(ctx, Noble_CreateNoble_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleClient) GetNoble(ctx context.Context, in *GetNobleReq, opts ...grpc.CallOption) (*GetNobleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNobleResp)
	err := c.cc.Invoke(ctx, Noble_GetNoble_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleClient) UpdateNoble(ctx context.Context, in *UpdateNobleReq, opts ...grpc.CallOption) (*UpdateNobleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNobleResp)
	err := c.cc.Invoke(ctx, Noble_UpdateNoble_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleClient) GetNobleList(ctx context.Context, in *GetNobleListReq, opts ...grpc.CallOption) (*GetNobleListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNobleListResp)
	err := c.cc.Invoke(ctx, Noble_GetNobleList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NobleServer is the server API for Noble service.
// All implementations must embed UnimplementedNobleServer
// for forward compatibility.
//
// 贵族
type NobleServer interface {
	// CreateNoble 创建
	CreateNoble(context.Context, *CreateNobleReq) (*CreateNobleResp, error)
	// GetNoble 查询
	GetNoble(context.Context, *GetNobleReq) (*GetNobleResp, error)
	// UpdateNoble 查询
	UpdateNoble(context.Context, *UpdateNobleReq) (*UpdateNobleResp, error)
	// GetNobleList 查询列表
	GetNobleList(context.Context, *GetNobleListReq) (*GetNobleListResp, error)
	mustEmbedUnimplementedNobleServer()
}

// UnimplementedNobleServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNobleServer struct{}

func (UnimplementedNobleServer) CreateNoble(context.Context, *CreateNobleReq) (*CreateNobleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNoble not implemented")
}
func (UnimplementedNobleServer) GetNoble(context.Context, *GetNobleReq) (*GetNobleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoble not implemented")
}
func (UnimplementedNobleServer) UpdateNoble(context.Context, *UpdateNobleReq) (*UpdateNobleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNoble not implemented")
}
func (UnimplementedNobleServer) GetNobleList(context.Context, *GetNobleListReq) (*GetNobleListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNobleList not implemented")
}
func (UnimplementedNobleServer) mustEmbedUnimplementedNobleServer() {}
func (UnimplementedNobleServer) testEmbeddedByValue()               {}

// UnsafeNobleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NobleServer will
// result in compilation errors.
type UnsafeNobleServer interface {
	mustEmbedUnimplementedNobleServer()
}

func RegisterNobleServer(s grpc.ServiceRegistrar, srv NobleServer) {
	// If the following call pancis, it indicates UnimplementedNobleServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Noble_ServiceDesc, srv)
}

func _Noble_CreateNoble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNobleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleServer).CreateNoble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Noble_CreateNoble_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleServer).CreateNoble(ctx, req.(*CreateNobleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noble_GetNoble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNobleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleServer).GetNoble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Noble_GetNoble_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleServer).GetNoble(ctx, req.(*GetNobleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noble_UpdateNoble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNobleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleServer).UpdateNoble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Noble_UpdateNoble_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleServer).UpdateNoble(ctx, req.(*UpdateNobleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noble_GetNobleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNobleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleServer).GetNobleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Noble_GetNobleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleServer).GetNobleList(ctx, req.(*GetNobleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Noble_ServiceDesc is the grpc.ServiceDesc for Noble service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Noble_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.vip.Noble",
	HandlerType: (*NobleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNoble",
			Handler:    _Noble_CreateNoble_Handler,
		},
		{
			MethodName: "GetNoble",
			Handler:    _Noble_GetNoble_Handler,
		},
		{
			MethodName: "UpdateNoble",
			Handler:    _Noble_UpdateNoble_Handler,
		},
		{
			MethodName: "GetNobleList",
			Handler:    _Noble_GetNobleList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.vip/noble.proto",
}
