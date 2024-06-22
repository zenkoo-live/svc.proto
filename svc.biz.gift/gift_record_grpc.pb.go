// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: svc.biz.gift/gift_record.proto

package gift

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
	GiftRecord_GetSendRecordList_FullMethodName = "/svc.biz.gift.GiftRecord/GetSendRecordList"
	GiftRecord_GetGetRecordList_FullMethodName  = "/svc.biz.gift.GiftRecord/GetGetRecordList"
	GiftRecord_GetLiveStat_FullMethodName       = "/svc.biz.gift.GiftRecord/GetLiveStat"
)

// GiftRecordClient is the client API for GiftRecord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GiftRecordClient interface {
	// GetSendRecordList 送礼记录
	GetSendRecordList(ctx context.Context, in *GetSendRecordListReq, opts ...grpc.CallOption) (*GetSendRecordListResp, error)
	// GetGetRecordList 收礼记录
	GetGetRecordList(ctx context.Context, in *GetGetRecordListReq, opts ...grpc.CallOption) (*GetGetRecordListResp, error)
	// GetLiveStat 直播统计
	GetLiveStat(ctx context.Context, in *GetLiveStatReq, opts ...grpc.CallOption) (*GetLiveStatResp, error)
}

type giftRecordClient struct {
	cc grpc.ClientConnInterface
}

func NewGiftRecordClient(cc grpc.ClientConnInterface) GiftRecordClient {
	return &giftRecordClient{cc}
}

func (c *giftRecordClient) GetSendRecordList(ctx context.Context, in *GetSendRecordListReq, opts ...grpc.CallOption) (*GetSendRecordListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSendRecordListResp)
	err := c.cc.Invoke(ctx, GiftRecord_GetSendRecordList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftRecordClient) GetGetRecordList(ctx context.Context, in *GetGetRecordListReq, opts ...grpc.CallOption) (*GetGetRecordListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGetRecordListResp)
	err := c.cc.Invoke(ctx, GiftRecord_GetGetRecordList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftRecordClient) GetLiveStat(ctx context.Context, in *GetLiveStatReq, opts ...grpc.CallOption) (*GetLiveStatResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLiveStatResp)
	err := c.cc.Invoke(ctx, GiftRecord_GetLiveStat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GiftRecordServer is the server API for GiftRecord service.
// All implementations must embed UnimplementedGiftRecordServer
// for forward compatibility
type GiftRecordServer interface {
	// GetSendRecordList 送礼记录
	GetSendRecordList(context.Context, *GetSendRecordListReq) (*GetSendRecordListResp, error)
	// GetGetRecordList 收礼记录
	GetGetRecordList(context.Context, *GetGetRecordListReq) (*GetGetRecordListResp, error)
	// GetLiveStat 直播统计
	GetLiveStat(context.Context, *GetLiveStatReq) (*GetLiveStatResp, error)
	mustEmbedUnimplementedGiftRecordServer()
}

// UnimplementedGiftRecordServer must be embedded to have forward compatible implementations.
type UnimplementedGiftRecordServer struct {
}

func (UnimplementedGiftRecordServer) GetSendRecordList(context.Context, *GetSendRecordListReq) (*GetSendRecordListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendRecordList not implemented")
}
func (UnimplementedGiftRecordServer) GetGetRecordList(context.Context, *GetGetRecordListReq) (*GetGetRecordListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGetRecordList not implemented")
}
func (UnimplementedGiftRecordServer) GetLiveStat(context.Context, *GetLiveStatReq) (*GetLiveStatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiveStat not implemented")
}
func (UnimplementedGiftRecordServer) mustEmbedUnimplementedGiftRecordServer() {}

// UnsafeGiftRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GiftRecordServer will
// result in compilation errors.
type UnsafeGiftRecordServer interface {
	mustEmbedUnimplementedGiftRecordServer()
}

func RegisterGiftRecordServer(s grpc.ServiceRegistrar, srv GiftRecordServer) {
	s.RegisterService(&GiftRecord_ServiceDesc, srv)
}

func _GiftRecord_GetSendRecordList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendRecordListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftRecordServer).GetSendRecordList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GiftRecord_GetSendRecordList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftRecordServer).GetSendRecordList(ctx, req.(*GetSendRecordListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiftRecord_GetGetRecordList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGetRecordListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftRecordServer).GetGetRecordList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GiftRecord_GetGetRecordList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftRecordServer).GetGetRecordList(ctx, req.(*GetGetRecordListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiftRecord_GetLiveStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLiveStatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftRecordServer).GetLiveStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GiftRecord_GetLiveStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftRecordServer).GetLiveStat(ctx, req.(*GetLiveStatReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GiftRecord_ServiceDesc is the grpc.ServiceDesc for GiftRecord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GiftRecord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.gift.GiftRecord",
	HandlerType: (*GiftRecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSendRecordList",
			Handler:    _GiftRecord_GetSendRecordList_Handler,
		},
		{
			MethodName: "GetGetRecordList",
			Handler:    _GiftRecord_GetGetRecordList_Handler,
		},
		{
			MethodName: "GetLiveStat",
			Handler:    _GiftRecord_GetLiveStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.gift/gift_record.proto",
}
