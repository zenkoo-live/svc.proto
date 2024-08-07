// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.3
// source: svc.biz.vip/fanbase.proto

package vip

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Fanbase_CreateFanbase_FullMethodName             = "/svc.biz.vip.Fanbase/CreateFanbase"
	Fanbase_GetFanbaseByStreamerID_FullMethodName    = "/svc.biz.vip.Fanbase/GetFanbaseByStreamerID"
	Fanbase_GetFanbaseByName_FullMethodName          = "/svc.biz.vip.Fanbase/GetFanbaseByName"
	Fanbase_UpdateFanbaseByStreamerID_FullMethodName = "/svc.biz.vip.Fanbase/UpdateFanbaseByStreamerID"
)

// FanbaseClient is the client API for Fanbase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 粉丝团
type FanbaseClient interface {
	// CreateFanbase 创建粉丝团
	CreateFanbase(ctx context.Context, in *CreateFanbaseReq, opts ...grpc.CallOption) (*CreateFanbaseResp, error)
	// GetFanbaseByStreamerID 获取粉丝团
	GetFanbaseByStreamerID(ctx context.Context, in *GetFanbaseByStreamerIDResp, opts ...grpc.CallOption) (*GetFanbaseResp, error)
	// GetFanbaseByName 通过名称获取粉丝团
	GetFanbaseByName(ctx context.Context, in *GetFanbaseByNameReq, opts ...grpc.CallOption) (*GetFanbaseResp, error)
	// UpdateFanbaseByStreamerID 更新粉丝团
	UpdateFanbaseByStreamerID(ctx context.Context, in *UpdateFanbaseByStreamerIDReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type fanbaseClient struct {
	cc grpc.ClientConnInterface
}

func NewFanbaseClient(cc grpc.ClientConnInterface) FanbaseClient {
	return &fanbaseClient{cc}
}

func (c *fanbaseClient) CreateFanbase(ctx context.Context, in *CreateFanbaseReq, opts ...grpc.CallOption) (*CreateFanbaseResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFanbaseResp)
	err := c.cc.Invoke(ctx, Fanbase_CreateFanbase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseClient) GetFanbaseByStreamerID(ctx context.Context, in *GetFanbaseByStreamerIDResp, opts ...grpc.CallOption) (*GetFanbaseResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFanbaseResp)
	err := c.cc.Invoke(ctx, Fanbase_GetFanbaseByStreamerID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseClient) GetFanbaseByName(ctx context.Context, in *GetFanbaseByNameReq, opts ...grpc.CallOption) (*GetFanbaseResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFanbaseResp)
	err := c.cc.Invoke(ctx, Fanbase_GetFanbaseByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseClient) UpdateFanbaseByStreamerID(ctx context.Context, in *UpdateFanbaseByStreamerIDReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Fanbase_UpdateFanbaseByStreamerID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FanbaseServer is the server API for Fanbase service.
// All implementations must embed UnimplementedFanbaseServer
// for forward compatibility
//
// 粉丝团
type FanbaseServer interface {
	// CreateFanbase 创建粉丝团
	CreateFanbase(context.Context, *CreateFanbaseReq) (*CreateFanbaseResp, error)
	// GetFanbaseByStreamerID 获取粉丝团
	GetFanbaseByStreamerID(context.Context, *GetFanbaseByStreamerIDResp) (*GetFanbaseResp, error)
	// GetFanbaseByName 通过名称获取粉丝团
	GetFanbaseByName(context.Context, *GetFanbaseByNameReq) (*GetFanbaseResp, error)
	// UpdateFanbaseByStreamerID 更新粉丝团
	UpdateFanbaseByStreamerID(context.Context, *UpdateFanbaseByStreamerIDReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedFanbaseServer()
}

// UnimplementedFanbaseServer must be embedded to have forward compatible implementations.
type UnimplementedFanbaseServer struct {
}

func (UnimplementedFanbaseServer) CreateFanbase(context.Context, *CreateFanbaseReq) (*CreateFanbaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFanbase not implemented")
}
func (UnimplementedFanbaseServer) GetFanbaseByStreamerID(context.Context, *GetFanbaseByStreamerIDResp) (*GetFanbaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFanbaseByStreamerID not implemented")
}
func (UnimplementedFanbaseServer) GetFanbaseByName(context.Context, *GetFanbaseByNameReq) (*GetFanbaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFanbaseByName not implemented")
}
func (UnimplementedFanbaseServer) UpdateFanbaseByStreamerID(context.Context, *UpdateFanbaseByStreamerIDReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFanbaseByStreamerID not implemented")
}
func (UnimplementedFanbaseServer) mustEmbedUnimplementedFanbaseServer() {}

// UnsafeFanbaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FanbaseServer will
// result in compilation errors.
type UnsafeFanbaseServer interface {
	mustEmbedUnimplementedFanbaseServer()
}

func RegisterFanbaseServer(s grpc.ServiceRegistrar, srv FanbaseServer) {
	s.RegisterService(&Fanbase_ServiceDesc, srv)
}

func _Fanbase_CreateFanbase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFanbaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseServer).CreateFanbase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fanbase_CreateFanbase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseServer).CreateFanbase(ctx, req.(*CreateFanbaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fanbase_GetFanbaseByStreamerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFanbaseByStreamerIDResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseServer).GetFanbaseByStreamerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fanbase_GetFanbaseByStreamerID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseServer).GetFanbaseByStreamerID(ctx, req.(*GetFanbaseByStreamerIDResp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fanbase_GetFanbaseByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFanbaseByNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseServer).GetFanbaseByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fanbase_GetFanbaseByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseServer).GetFanbaseByName(ctx, req.(*GetFanbaseByNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fanbase_UpdateFanbaseByStreamerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFanbaseByStreamerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseServer).UpdateFanbaseByStreamerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fanbase_UpdateFanbaseByStreamerID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseServer).UpdateFanbaseByStreamerID(ctx, req.(*UpdateFanbaseByStreamerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Fanbase_ServiceDesc is the grpc.ServiceDesc for Fanbase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fanbase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.vip.Fanbase",
	HandlerType: (*FanbaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFanbase",
			Handler:    _Fanbase_CreateFanbase_Handler,
		},
		{
			MethodName: "GetFanbaseByStreamerID",
			Handler:    _Fanbase_GetFanbaseByStreamerID_Handler,
		},
		{
			MethodName: "GetFanbaseByName",
			Handler:    _Fanbase_GetFanbaseByName_Handler,
		},
		{
			MethodName: "UpdateFanbaseByStreamerID",
			Handler:    _Fanbase_UpdateFanbaseByStreamerID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.vip/fanbase.proto",
}
