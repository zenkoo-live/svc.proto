// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: svc.biz.gift/gift.proto

package gift

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
	Gift_Add_FullMethodName              = "/svc.biz.gift.Gift/Add"
	Gift_Get_FullMethodName              = "/svc.biz.gift.Gift/Get"
	Gift_Update_FullMethodName           = "/svc.biz.gift.Gift/Update"
	Gift_ListAdmin_FullMethodName        = "/svc.biz.gift.Gift/ListAdmin"
	Gift_ListOnlineByType_FullMethodName = "/svc.biz.gift.Gift/ListOnlineByType"
	Gift_ListOnlineAll_FullMethodName    = "/svc.biz.gift.Gift/ListOnlineAll"
	Gift_Send_FullMethodName             = "/svc.biz.gift.Gift/Send"
)

// GiftClient is the client API for Gift service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GiftClient interface {
	// Add 添加礼物
	Add(ctx context.Context, in *GiftAddReq, opts ...grpc.CallOption) (*GiftAddResp, error)
	// Get 查询礼物
	Get(ctx context.Context, in *GiftGetReq, opts ...grpc.CallOption) (*GiftGetResp, error)
	// Update 更新礼物
	Update(ctx context.Context, in *GiftUpdateReq, opts ...grpc.CallOption) (*GiftUpdateResp, error)
	// ListAdmin 后台查询礼物列表接口
	ListAdmin(ctx context.Context, in *ListAdminReq, opts ...grpc.CallOption) (*ListAdminResp, error)
	// ListOnlineByType 前台房间礼物查询接口
	ListOnlineByType(ctx context.Context, in *ListOnlineByTypeReq, opts ...grpc.CallOption) (*ListOnlineResp, error)
	// ListOnlineAll 所有礼物的缓存接口
	ListOnlineAll(ctx context.Context, in *ListOnlineAllReq, opts ...grpc.CallOption) (*ListOnlineResp, error)
	// Send 送礼物接口
	Send(ctx context.Context, in *GiftSendReq, opts ...grpc.CallOption) (*GiftSendResp, error)
}

type giftClient struct {
	cc grpc.ClientConnInterface
}

func NewGiftClient(cc grpc.ClientConnInterface) GiftClient {
	return &giftClient{cc}
}

func (c *giftClient) Add(ctx context.Context, in *GiftAddReq, opts ...grpc.CallOption) (*GiftAddResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GiftAddResp)
	err := c.cc.Invoke(ctx, Gift_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftClient) Get(ctx context.Context, in *GiftGetReq, opts ...grpc.CallOption) (*GiftGetResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GiftGetResp)
	err := c.cc.Invoke(ctx, Gift_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftClient) Update(ctx context.Context, in *GiftUpdateReq, opts ...grpc.CallOption) (*GiftUpdateResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GiftUpdateResp)
	err := c.cc.Invoke(ctx, Gift_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftClient) ListAdmin(ctx context.Context, in *ListAdminReq, opts ...grpc.CallOption) (*ListAdminResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAdminResp)
	err := c.cc.Invoke(ctx, Gift_ListAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftClient) ListOnlineByType(ctx context.Context, in *ListOnlineByTypeReq, opts ...grpc.CallOption) (*ListOnlineResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOnlineResp)
	err := c.cc.Invoke(ctx, Gift_ListOnlineByType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftClient) ListOnlineAll(ctx context.Context, in *ListOnlineAllReq, opts ...grpc.CallOption) (*ListOnlineResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOnlineResp)
	err := c.cc.Invoke(ctx, Gift_ListOnlineAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftClient) Send(ctx context.Context, in *GiftSendReq, opts ...grpc.CallOption) (*GiftSendResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GiftSendResp)
	err := c.cc.Invoke(ctx, Gift_Send_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GiftServer is the server API for Gift service.
// All implementations must embed UnimplementedGiftServer
// for forward compatibility.
type GiftServer interface {
	// Add 添加礼物
	Add(context.Context, *GiftAddReq) (*GiftAddResp, error)
	// Get 查询礼物
	Get(context.Context, *GiftGetReq) (*GiftGetResp, error)
	// Update 更新礼物
	Update(context.Context, *GiftUpdateReq) (*GiftUpdateResp, error)
	// ListAdmin 后台查询礼物列表接口
	ListAdmin(context.Context, *ListAdminReq) (*ListAdminResp, error)
	// ListOnlineByType 前台房间礼物查询接口
	ListOnlineByType(context.Context, *ListOnlineByTypeReq) (*ListOnlineResp, error)
	// ListOnlineAll 所有礼物的缓存接口
	ListOnlineAll(context.Context, *ListOnlineAllReq) (*ListOnlineResp, error)
	// Send 送礼物接口
	Send(context.Context, *GiftSendReq) (*GiftSendResp, error)
	mustEmbedUnimplementedGiftServer()
}

// UnimplementedGiftServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGiftServer struct{}

func (UnimplementedGiftServer) Add(context.Context, *GiftAddReq) (*GiftAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedGiftServer) Get(context.Context, *GiftGetReq) (*GiftGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGiftServer) Update(context.Context, *GiftUpdateReq) (*GiftUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGiftServer) ListAdmin(context.Context, *ListAdminReq) (*ListAdminResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAdmin not implemented")
}
func (UnimplementedGiftServer) ListOnlineByType(context.Context, *ListOnlineByTypeReq) (*ListOnlineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOnlineByType not implemented")
}
func (UnimplementedGiftServer) ListOnlineAll(context.Context, *ListOnlineAllReq) (*ListOnlineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOnlineAll not implemented")
}
func (UnimplementedGiftServer) Send(context.Context, *GiftSendReq) (*GiftSendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedGiftServer) mustEmbedUnimplementedGiftServer() {}
func (UnimplementedGiftServer) testEmbeddedByValue()              {}

// UnsafeGiftServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GiftServer will
// result in compilation errors.
type UnsafeGiftServer interface {
	mustEmbedUnimplementedGiftServer()
}

func RegisterGiftServer(s grpc.ServiceRegistrar, srv GiftServer) {
	// If the following call pancis, it indicates UnimplementedGiftServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Gift_ServiceDesc, srv)
}

func _Gift_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiftAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gift_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServer).Add(ctx, req.(*GiftAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gift_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiftGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gift_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServer).Get(ctx, req.(*GiftGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gift_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiftUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gift_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServer).Update(ctx, req.(*GiftUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gift_ListAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServer).ListAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gift_ListAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServer).ListAdmin(ctx, req.(*ListAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gift_ListOnlineByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOnlineByTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServer).ListOnlineByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gift_ListOnlineByType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServer).ListOnlineByType(ctx, req.(*ListOnlineByTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gift_ListOnlineAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOnlineAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServer).ListOnlineAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gift_ListOnlineAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServer).ListOnlineAll(ctx, req.(*ListOnlineAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gift_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiftSendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gift_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServer).Send(ctx, req.(*GiftSendReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Gift_ServiceDesc is the grpc.ServiceDesc for Gift service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gift_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.gift.Gift",
	HandlerType: (*GiftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Gift_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Gift_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Gift_Update_Handler,
		},
		{
			MethodName: "ListAdmin",
			Handler:    _Gift_ListAdmin_Handler,
		},
		{
			MethodName: "ListOnlineByType",
			Handler:    _Gift_ListOnlineByType_Handler,
		},
		{
			MethodName: "ListOnlineAll",
			Handler:    _Gift_ListOnlineAll_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Gift_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.gift/gift.proto",
}
