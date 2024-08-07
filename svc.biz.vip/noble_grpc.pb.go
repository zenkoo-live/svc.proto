// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.3
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
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Noble_CreateNoble_FullMethodName        = "/svc.biz.vip.Noble/CreateNoble"
	Noble_GetNobleByLevel_FullMethodName    = "/svc.biz.vip.Noble/GetNobleByLevel"
	Noble_GetNobleList_FullMethodName       = "/svc.biz.vip.Noble/GetNobleList"
	Noble_UpdateNobleByLevel_FullMethodName = "/svc.biz.vip.Noble/UpdateNobleByLevel"
)

// NobleClient is the client API for Noble service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 贵族
type NobleClient interface {
	// CreateNoble 创建
	CreateNoble(ctx context.Context, in *CreateNobleReq, opts ...grpc.CallOption) (*CreateNobleResp, error)
	// GetNobleByLevel 查询
	GetNobleByLevel(ctx context.Context, in *GetNobleByLevelReq, opts ...grpc.CallOption) (*GetNobleByLevelResp, error)
	// CreateNoble 查询列表
	GetNobleList(ctx context.Context, in *GetNobleListReq, opts ...grpc.CallOption) (*GetNobleListResp, error)
	// UpdateNobleByLevel 更新
	UpdateNobleByLevel(ctx context.Context, in *UpdateNobleByLevelReq, opts ...grpc.CallOption) (*UpdateNobleByLevelResp, error)
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

func (c *nobleClient) GetNobleByLevel(ctx context.Context, in *GetNobleByLevelReq, opts ...grpc.CallOption) (*GetNobleByLevelResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNobleByLevelResp)
	err := c.cc.Invoke(ctx, Noble_GetNobleByLevel_FullMethodName, in, out, cOpts...)
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

func (c *nobleClient) UpdateNobleByLevel(ctx context.Context, in *UpdateNobleByLevelReq, opts ...grpc.CallOption) (*UpdateNobleByLevelResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNobleByLevelResp)
	err := c.cc.Invoke(ctx, Noble_UpdateNobleByLevel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NobleServer is the server API for Noble service.
// All implementations must embed UnimplementedNobleServer
// for forward compatibility
//
// 贵族
type NobleServer interface {
	// CreateNoble 创建
	CreateNoble(context.Context, *CreateNobleReq) (*CreateNobleResp, error)
	// GetNobleByLevel 查询
	GetNobleByLevel(context.Context, *GetNobleByLevelReq) (*GetNobleByLevelResp, error)
	// CreateNoble 查询列表
	GetNobleList(context.Context, *GetNobleListReq) (*GetNobleListResp, error)
	// UpdateNobleByLevel 更新
	UpdateNobleByLevel(context.Context, *UpdateNobleByLevelReq) (*UpdateNobleByLevelResp, error)
	mustEmbedUnimplementedNobleServer()
}

// UnimplementedNobleServer must be embedded to have forward compatible implementations.
type UnimplementedNobleServer struct {
}

func (UnimplementedNobleServer) CreateNoble(context.Context, *CreateNobleReq) (*CreateNobleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNoble not implemented")
}
func (UnimplementedNobleServer) GetNobleByLevel(context.Context, *GetNobleByLevelReq) (*GetNobleByLevelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNobleByLevel not implemented")
}
func (UnimplementedNobleServer) GetNobleList(context.Context, *GetNobleListReq) (*GetNobleListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNobleList not implemented")
}
func (UnimplementedNobleServer) UpdateNobleByLevel(context.Context, *UpdateNobleByLevelReq) (*UpdateNobleByLevelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNobleByLevel not implemented")
}
func (UnimplementedNobleServer) mustEmbedUnimplementedNobleServer() {}

// UnsafeNobleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NobleServer will
// result in compilation errors.
type UnsafeNobleServer interface {
	mustEmbedUnimplementedNobleServer()
}

func RegisterNobleServer(s grpc.ServiceRegistrar, srv NobleServer) {
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

func _Noble_GetNobleByLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNobleByLevelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleServer).GetNobleByLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Noble_GetNobleByLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleServer).GetNobleByLevel(ctx, req.(*GetNobleByLevelReq))
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

func _Noble_UpdateNobleByLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNobleByLevelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleServer).UpdateNobleByLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Noble_UpdateNobleByLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleServer).UpdateNobleByLevel(ctx, req.(*UpdateNobleByLevelReq))
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
			MethodName: "GetNobleByLevel",
			Handler:    _Noble_GetNobleByLevel_Handler,
		},
		{
			MethodName: "GetNobleList",
			Handler:    _Noble_GetNobleList_Handler,
		},
		{
			MethodName: "UpdateNobleByLevel",
			Handler:    _Noble_UpdateNobleByLevel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.vip/noble.proto",
}
