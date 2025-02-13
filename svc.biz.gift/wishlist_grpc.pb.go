// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: svc.biz.gift/wishlist.proto

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
	LiveWishlist_SetWishlist_FullMethodName       = "/svc.biz.gift.LiveWishlist/SetWishlist"
	LiveWishlist_GetByRoomId_FullMethodName       = "/svc.biz.gift.LiveWishlist/GetByRoomId"
	LiveWishlist_UpdateWishGifts_FullMethodName   = "/svc.biz.gift.LiveWishlist/UpdateWishGifts"
	LiveWishlist_SetActiveStatus_FullMethodName   = "/svc.biz.gift.LiveWishlist/SetActiveStatus"
	LiveWishlist_SetAutomodeStatus_FullMethodName = "/svc.biz.gift.LiveWishlist/SetAutomodeStatus"
	LiveWishlist_ExecAutoModeTask_FullMethodName  = "/svc.biz.gift.LiveWishlist/ExecAutoModeTask"
)

// LiveWishlistClient is the client API for LiveWishlist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 心愿单
type LiveWishlistClient interface {
	// 设置心愿单
	SetWishlist(ctx context.Context, in *SetWishlistReq, opts ...grpc.CallOption) (*SetWishlistResp, error)
	// 查询指定房间的心愿单信息
	GetByRoomId(ctx context.Context, in *GetWishlistByRoomIdReq, opts ...grpc.CallOption) (*WishlistInfoResp, error)
	// 修改心愿单礼物
	UpdateWishGifts(ctx context.Context, in *UpdateGiftsReq, opts ...grpc.CallOption) (*WishlistInfoResp, error)
	// 设置心愿单开启状态
	SetActiveStatus(ctx context.Context, in *EnabledStatusInfo, opts ...grpc.CallOption) (*EnabledStatusInfo, error)
	// 设置自动模式(每天自动刷新心愿单)开启状态
	SetAutomodeStatus(ctx context.Context, in *EnabledStatusInfo, opts ...grpc.CallOption) (*EnabledStatusInfo, error)
	// 执行自动模式逻辑(定时任务调用)
	ExecAutoModeTask(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type liveWishlistClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveWishlistClient(cc grpc.ClientConnInterface) LiveWishlistClient {
	return &liveWishlistClient{cc}
}

func (c *liveWishlistClient) SetWishlist(ctx context.Context, in *SetWishlistReq, opts ...grpc.CallOption) (*SetWishlistResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetWishlistResp)
	err := c.cc.Invoke(ctx, LiveWishlist_SetWishlist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveWishlistClient) GetByRoomId(ctx context.Context, in *GetWishlistByRoomIdReq, opts ...grpc.CallOption) (*WishlistInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WishlistInfoResp)
	err := c.cc.Invoke(ctx, LiveWishlist_GetByRoomId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveWishlistClient) UpdateWishGifts(ctx context.Context, in *UpdateGiftsReq, opts ...grpc.CallOption) (*WishlistInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WishlistInfoResp)
	err := c.cc.Invoke(ctx, LiveWishlist_UpdateWishGifts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveWishlistClient) SetActiveStatus(ctx context.Context, in *EnabledStatusInfo, opts ...grpc.CallOption) (*EnabledStatusInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnabledStatusInfo)
	err := c.cc.Invoke(ctx, LiveWishlist_SetActiveStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveWishlistClient) SetAutomodeStatus(ctx context.Context, in *EnabledStatusInfo, opts ...grpc.CallOption) (*EnabledStatusInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnabledStatusInfo)
	err := c.cc.Invoke(ctx, LiveWishlist_SetAutomodeStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveWishlistClient) ExecAutoModeTask(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, LiveWishlist_ExecAutoModeTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveWishlistServer is the server API for LiveWishlist service.
// All implementations must embed UnimplementedLiveWishlistServer
// for forward compatibility.
//
// 心愿单
type LiveWishlistServer interface {
	// 设置心愿单
	SetWishlist(context.Context, *SetWishlistReq) (*SetWishlistResp, error)
	// 查询指定房间的心愿单信息
	GetByRoomId(context.Context, *GetWishlistByRoomIdReq) (*WishlistInfoResp, error)
	// 修改心愿单礼物
	UpdateWishGifts(context.Context, *UpdateGiftsReq) (*WishlistInfoResp, error)
	// 设置心愿单开启状态
	SetActiveStatus(context.Context, *EnabledStatusInfo) (*EnabledStatusInfo, error)
	// 设置自动模式(每天自动刷新心愿单)开启状态
	SetAutomodeStatus(context.Context, *EnabledStatusInfo) (*EnabledStatusInfo, error)
	// 执行自动模式逻辑(定时任务调用)
	ExecAutoModeTask(context.Context, *EmptyRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedLiveWishlistServer()
}

// UnimplementedLiveWishlistServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLiveWishlistServer struct{}

func (UnimplementedLiveWishlistServer) SetWishlist(context.Context, *SetWishlistReq) (*SetWishlistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWishlist not implemented")
}
func (UnimplementedLiveWishlistServer) GetByRoomId(context.Context, *GetWishlistByRoomIdReq) (*WishlistInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByRoomId not implemented")
}
func (UnimplementedLiveWishlistServer) UpdateWishGifts(context.Context, *UpdateGiftsReq) (*WishlistInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWishGifts not implemented")
}
func (UnimplementedLiveWishlistServer) SetActiveStatus(context.Context, *EnabledStatusInfo) (*EnabledStatusInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetActiveStatus not implemented")
}
func (UnimplementedLiveWishlistServer) SetAutomodeStatus(context.Context, *EnabledStatusInfo) (*EnabledStatusInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAutomodeStatus not implemented")
}
func (UnimplementedLiveWishlistServer) ExecAutoModeTask(context.Context, *EmptyRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecAutoModeTask not implemented")
}
func (UnimplementedLiveWishlistServer) mustEmbedUnimplementedLiveWishlistServer() {}
func (UnimplementedLiveWishlistServer) testEmbeddedByValue()                      {}

// UnsafeLiveWishlistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveWishlistServer will
// result in compilation errors.
type UnsafeLiveWishlistServer interface {
	mustEmbedUnimplementedLiveWishlistServer()
}

func RegisterLiveWishlistServer(s grpc.ServiceRegistrar, srv LiveWishlistServer) {
	// If the following call pancis, it indicates UnimplementedLiveWishlistServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LiveWishlist_ServiceDesc, srv)
}

func _LiveWishlist_SetWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetWishlistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveWishlistServer).SetWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveWishlist_SetWishlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveWishlistServer).SetWishlist(ctx, req.(*SetWishlistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveWishlist_GetByRoomId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWishlistByRoomIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveWishlistServer).GetByRoomId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveWishlist_GetByRoomId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveWishlistServer).GetByRoomId(ctx, req.(*GetWishlistByRoomIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveWishlist_UpdateWishGifts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGiftsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveWishlistServer).UpdateWishGifts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveWishlist_UpdateWishGifts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveWishlistServer).UpdateWishGifts(ctx, req.(*UpdateGiftsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveWishlist_SetActiveStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnabledStatusInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveWishlistServer).SetActiveStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveWishlist_SetActiveStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveWishlistServer).SetActiveStatus(ctx, req.(*EnabledStatusInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveWishlist_SetAutomodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnabledStatusInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveWishlistServer).SetAutomodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveWishlist_SetAutomodeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveWishlistServer).SetAutomodeStatus(ctx, req.(*EnabledStatusInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveWishlist_ExecAutoModeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveWishlistServer).ExecAutoModeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiveWishlist_ExecAutoModeTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveWishlistServer).ExecAutoModeTask(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LiveWishlist_ServiceDesc is the grpc.ServiceDesc for LiveWishlist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveWishlist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.gift.LiveWishlist",
	HandlerType: (*LiveWishlistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetWishlist",
			Handler:    _LiveWishlist_SetWishlist_Handler,
		},
		{
			MethodName: "GetByRoomId",
			Handler:    _LiveWishlist_GetByRoomId_Handler,
		},
		{
			MethodName: "UpdateWishGifts",
			Handler:    _LiveWishlist_UpdateWishGifts_Handler,
		},
		{
			MethodName: "SetActiveStatus",
			Handler:    _LiveWishlist_SetActiveStatus_Handler,
		},
		{
			MethodName: "SetAutomodeStatus",
			Handler:    _LiveWishlist_SetAutomodeStatus_Handler,
		},
		{
			MethodName: "ExecAutoModeTask",
			Handler:    _LiveWishlist_ExecAutoModeTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.gift/wishlist.proto",
}
