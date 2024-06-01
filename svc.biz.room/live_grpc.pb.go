// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: svc.biz.room/live.proto

package room

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

const (
	Live_GetLive_FullMethodName  = "/svc.biz.room.Live/GetLive"
	Live_MGetLive_FullMethodName = "/svc.biz.room.Live/MGetLive"
	Live_ListLive_FullMethodName = "/svc.biz.room.Live/ListLive"
)

// LiveClient is the client API for Live service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveClient interface {
	// 查询直播间信息
	GetLive(ctx context.Context, in *GetLiveReq, opts ...grpc.CallOption) (*GetLiveResp, error)
	// 批量获取直播间信息
	MGetLive(ctx context.Context, in *MGetLiveReq, opts ...grpc.CallOption) (*MGetLiveResp, error)
	// 获取在播直播间列表
	ListLive(ctx context.Context, in *ListLiveReq, opts ...grpc.CallOption) (*ListLiveResp, error)
}

type liveClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveClient(cc grpc.ClientConnInterface) LiveClient {
	return &liveClient{cc}
}

func (c *liveClient) GetLive(ctx context.Context, in *GetLiveReq, opts ...grpc.CallOption) (*GetLiveResp, error) {
	out := new(GetLiveResp)
	err := c.cc.Invoke(ctx, Live_GetLive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveClient) MGetLive(ctx context.Context, in *MGetLiveReq, opts ...grpc.CallOption) (*MGetLiveResp, error) {
	out := new(MGetLiveResp)
	err := c.cc.Invoke(ctx, Live_MGetLive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveClient) ListLive(ctx context.Context, in *ListLiveReq, opts ...grpc.CallOption) (*ListLiveResp, error) {
	out := new(ListLiveResp)
	err := c.cc.Invoke(ctx, Live_ListLive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveServer is the server API for Live service.
// All implementations must embed UnimplementedLiveServer
// for forward compatibility
type LiveServer interface {
	// 查询直播间信息
	GetLive(context.Context, *GetLiveReq) (*GetLiveResp, error)
	// 批量获取直播间信息
	MGetLive(context.Context, *MGetLiveReq) (*MGetLiveResp, error)
	// 获取在播直播间列表
	ListLive(context.Context, *ListLiveReq) (*ListLiveResp, error)
	mustEmbedUnimplementedLiveServer()
}

// UnimplementedLiveServer must be embedded to have forward compatible implementations.
type UnimplementedLiveServer struct {
}

func (UnimplementedLiveServer) GetLive(context.Context, *GetLiveReq) (*GetLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLive not implemented")
}
func (UnimplementedLiveServer) MGetLive(context.Context, *MGetLiveReq) (*MGetLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetLive not implemented")
}
func (UnimplementedLiveServer) ListLive(context.Context, *ListLiveReq) (*ListLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLive not implemented")
}
func (UnimplementedLiveServer) mustEmbedUnimplementedLiveServer() {}

// UnsafeLiveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveServer will
// result in compilation errors.
type UnsafeLiveServer interface {
	mustEmbedUnimplementedLiveServer()
}

func RegisterLiveServer(s grpc.ServiceRegistrar, srv LiveServer) {
	s.RegisterService(&Live_ServiceDesc, srv)
}

func _Live_GetLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServer).GetLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Live_GetLive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServer).GetLive(ctx, req.(*GetLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Live_MGetLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServer).MGetLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Live_MGetLive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServer).MGetLive(ctx, req.(*MGetLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Live_ListLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServer).ListLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Live_ListLive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServer).ListLive(ctx, req.(*ListLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Live_ServiceDesc is the grpc.ServiceDesc for Live service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Live_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.room.Live",
	HandlerType: (*LiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLive",
			Handler:    _Live_GetLive_Handler,
		},
		{
			MethodName: "MGetLive",
			Handler:    _Live_MGetLive_Handler,
		},
		{
			MethodName: "ListLive",
			Handler:    _Live_ListLive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.room/live.proto",
}
