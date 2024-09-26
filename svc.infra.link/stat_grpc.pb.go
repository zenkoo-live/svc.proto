// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: svc.infra.link/stat.proto

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
	LinkStat_OnlineSessionCount_FullMethodName = "/svc.infra.link.LinkStat/OnlineSessionCount"
	LinkStat_OnlineAccountCount_FullMethodName = "/svc.infra.link.LinkStat/OnlineAccountCount"
	LinkStat_OnlineDeviceCount_FullMethodName  = "/svc.infra.link.LinkStat/OnlineDeviceCount"
	LinkStat_OnlineSessionList_FullMethodName  = "/svc.infra.link.LinkStat/OnlineSessionList"
	LinkStat_OnlineAccountList_FullMethodName  = "/svc.infra.link.LinkStat/OnlineAccountList"
	LinkStat_OnlineDeviceList_FullMethodName   = "/svc.infra.link.LinkStat/OnlineDeviceList"
	LinkStat_CheckSession_FullMethodName       = "/svc.infra.link.LinkStat/CheckSession"
	LinkStat_CheckAccount_FullMethodName       = "/svc.infra.link.LinkStat/CheckAccount"
	LinkStat_CheckDevice_FullMethodName        = "/svc.infra.link.LinkStat/CheckDevice"
	LinkStat_RoomLiveStat_FullMethodName       = "/svc.infra.link.LinkStat/RoomLiveStat"
	LinkStat_CheckAccountRoom_FullMethodName   = "/svc.infra.link.LinkStat/CheckAccountRoom"
	LinkStat_CheckSessionRoom_FullMethodName   = "/svc.infra.link.LinkStat/CheckSessionRoom"
	LinkStat_CheckDeviceRoom_FullMethodName    = "/svc.infra.link.LinkStat/CheckDeviceRoom"
	LinkStat_Refresh_FullMethodName            = "/svc.infra.link.LinkStat/Refresh"
)

// LinkStatClient is the client API for LinkStat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service of stat
type LinkStatClient interface {
	// 获取在线连接数
	OnlineSessionCount(ctx context.Context, in *OnlineCountRequest, opts ...grpc.CallOption) (*OnlineCountResponse, error)
	// 获取在线账号数
	OnlineAccountCount(ctx context.Context, in *OnlineCountRequest, opts ...grpc.CallOption) (*OnlineCountResponse, error)
	// 获取在线设备数
	OnlineDeviceCount(ctx context.Context, in *OnlineCountRequest, opts ...grpc.CallOption) (*OnlineCountResponse, error)
	// 获取在线连接列表
	OnlineSessionList(ctx context.Context, in *OnlineListRequest, opts ...grpc.CallOption) (*OnlineSessionListResponse, error)
	// 获取在线账号列表
	OnlineAccountList(ctx context.Context, in *OnlineListRequest, opts ...grpc.CallOption) (*OnlineAccountListResponse, error)
	// 获取在线设备列表
	OnlineDeviceList(ctx context.Context, in *OnlineListRequest, opts ...grpc.CallOption) (*OnlineDeviceListResponse, error)
	// 检查连接是否在线
	CheckSession(ctx context.Context, in *CheckSessionRequest, opts ...grpc.CallOption) (*CheckSessionResponse, error)
	// 检查账号是否在线
	CheckAccount(ctx context.Context, in *CheckAccountRequest, opts ...grpc.CallOption) (*CheckAccountResponse, error)
	// 检查设备是否在线
	CheckDevice(ctx context.Context, in *CheckDeviceRequest, opts ...grpc.CallOption) (*CheckDeviceResponse, error)
	// 直播间统计数据
	RoomLiveStat(ctx context.Context, in *RoomLiveStatRequest, opts ...grpc.CallOption) (*RoomLiveStatResponse, error)
	// 检查账号是否在房间内
	CheckAccountRoom(ctx context.Context, in *CheckAccountRoomRequest, opts ...grpc.CallOption) (*CheckAccountResponse, error)
	// 检查session是否在房间内
	CheckSessionRoom(ctx context.Context, in *CheckSessionRoomRequest, opts ...grpc.CallOption) (*CheckSessionResponse, error)
	// 检查设备是否在房间内
	CheckDeviceRoom(ctx context.Context, in *CheckDeviceRoomRequest, opts ...grpc.CallOption) (*CheckDeviceResponse, error)
	// 刷新统计
	Refresh(ctx context.Context, in *RefreshStatRequest, opts ...grpc.CallOption) (*RefreshStatResponse, error)
}

type linkStatClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkStatClient(cc grpc.ClientConnInterface) LinkStatClient {
	return &linkStatClient{cc}
}

func (c *linkStatClient) OnlineSessionCount(ctx context.Context, in *OnlineCountRequest, opts ...grpc.CallOption) (*OnlineCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnlineCountResponse)
	err := c.cc.Invoke(ctx, LinkStat_OnlineSessionCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) OnlineAccountCount(ctx context.Context, in *OnlineCountRequest, opts ...grpc.CallOption) (*OnlineCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnlineCountResponse)
	err := c.cc.Invoke(ctx, LinkStat_OnlineAccountCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) OnlineDeviceCount(ctx context.Context, in *OnlineCountRequest, opts ...grpc.CallOption) (*OnlineCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnlineCountResponse)
	err := c.cc.Invoke(ctx, LinkStat_OnlineDeviceCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) OnlineSessionList(ctx context.Context, in *OnlineListRequest, opts ...grpc.CallOption) (*OnlineSessionListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnlineSessionListResponse)
	err := c.cc.Invoke(ctx, LinkStat_OnlineSessionList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) OnlineAccountList(ctx context.Context, in *OnlineListRequest, opts ...grpc.CallOption) (*OnlineAccountListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnlineAccountListResponse)
	err := c.cc.Invoke(ctx, LinkStat_OnlineAccountList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) OnlineDeviceList(ctx context.Context, in *OnlineListRequest, opts ...grpc.CallOption) (*OnlineDeviceListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnlineDeviceListResponse)
	err := c.cc.Invoke(ctx, LinkStat_OnlineDeviceList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) CheckSession(ctx context.Context, in *CheckSessionRequest, opts ...grpc.CallOption) (*CheckSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckSessionResponse)
	err := c.cc.Invoke(ctx, LinkStat_CheckSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) CheckAccount(ctx context.Context, in *CheckAccountRequest, opts ...grpc.CallOption) (*CheckAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAccountResponse)
	err := c.cc.Invoke(ctx, LinkStat_CheckAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) CheckDevice(ctx context.Context, in *CheckDeviceRequest, opts ...grpc.CallOption) (*CheckDeviceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckDeviceResponse)
	err := c.cc.Invoke(ctx, LinkStat_CheckDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) RoomLiveStat(ctx context.Context, in *RoomLiveStatRequest, opts ...grpc.CallOption) (*RoomLiveStatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoomLiveStatResponse)
	err := c.cc.Invoke(ctx, LinkStat_RoomLiveStat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) CheckAccountRoom(ctx context.Context, in *CheckAccountRoomRequest, opts ...grpc.CallOption) (*CheckAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAccountResponse)
	err := c.cc.Invoke(ctx, LinkStat_CheckAccountRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) CheckSessionRoom(ctx context.Context, in *CheckSessionRoomRequest, opts ...grpc.CallOption) (*CheckSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckSessionResponse)
	err := c.cc.Invoke(ctx, LinkStat_CheckSessionRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) CheckDeviceRoom(ctx context.Context, in *CheckDeviceRoomRequest, opts ...grpc.CallOption) (*CheckDeviceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckDeviceResponse)
	err := c.cc.Invoke(ctx, LinkStat_CheckDeviceRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkStatClient) Refresh(ctx context.Context, in *RefreshStatRequest, opts ...grpc.CallOption) (*RefreshStatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshStatResponse)
	err := c.cc.Invoke(ctx, LinkStat_Refresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkStatServer is the server API for LinkStat service.
// All implementations must embed UnimplementedLinkStatServer
// for forward compatibility.
//
// Service of stat
type LinkStatServer interface {
	// 获取在线连接数
	OnlineSessionCount(context.Context, *OnlineCountRequest) (*OnlineCountResponse, error)
	// 获取在线账号数
	OnlineAccountCount(context.Context, *OnlineCountRequest) (*OnlineCountResponse, error)
	// 获取在线设备数
	OnlineDeviceCount(context.Context, *OnlineCountRequest) (*OnlineCountResponse, error)
	// 获取在线连接列表
	OnlineSessionList(context.Context, *OnlineListRequest) (*OnlineSessionListResponse, error)
	// 获取在线账号列表
	OnlineAccountList(context.Context, *OnlineListRequest) (*OnlineAccountListResponse, error)
	// 获取在线设备列表
	OnlineDeviceList(context.Context, *OnlineListRequest) (*OnlineDeviceListResponse, error)
	// 检查连接是否在线
	CheckSession(context.Context, *CheckSessionRequest) (*CheckSessionResponse, error)
	// 检查账号是否在线
	CheckAccount(context.Context, *CheckAccountRequest) (*CheckAccountResponse, error)
	// 检查设备是否在线
	CheckDevice(context.Context, *CheckDeviceRequest) (*CheckDeviceResponse, error)
	// 直播间统计数据
	RoomLiveStat(context.Context, *RoomLiveStatRequest) (*RoomLiveStatResponse, error)
	// 检查账号是否在房间内
	CheckAccountRoom(context.Context, *CheckAccountRoomRequest) (*CheckAccountResponse, error)
	// 检查session是否在房间内
	CheckSessionRoom(context.Context, *CheckSessionRoomRequest) (*CheckSessionResponse, error)
	// 检查设备是否在房间内
	CheckDeviceRoom(context.Context, *CheckDeviceRoomRequest) (*CheckDeviceResponse, error)
	// 刷新统计
	Refresh(context.Context, *RefreshStatRequest) (*RefreshStatResponse, error)
	mustEmbedUnimplementedLinkStatServer()
}

// UnimplementedLinkStatServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLinkStatServer struct{}

func (UnimplementedLinkStatServer) OnlineSessionCount(context.Context, *OnlineCountRequest) (*OnlineCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineSessionCount not implemented")
}
func (UnimplementedLinkStatServer) OnlineAccountCount(context.Context, *OnlineCountRequest) (*OnlineCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineAccountCount not implemented")
}
func (UnimplementedLinkStatServer) OnlineDeviceCount(context.Context, *OnlineCountRequest) (*OnlineCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineDeviceCount not implemented")
}
func (UnimplementedLinkStatServer) OnlineSessionList(context.Context, *OnlineListRequest) (*OnlineSessionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineSessionList not implemented")
}
func (UnimplementedLinkStatServer) OnlineAccountList(context.Context, *OnlineListRequest) (*OnlineAccountListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineAccountList not implemented")
}
func (UnimplementedLinkStatServer) OnlineDeviceList(context.Context, *OnlineListRequest) (*OnlineDeviceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineDeviceList not implemented")
}
func (UnimplementedLinkStatServer) CheckSession(context.Context, *CheckSessionRequest) (*CheckSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSession not implemented")
}
func (UnimplementedLinkStatServer) CheckAccount(context.Context, *CheckAccountRequest) (*CheckAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccount not implemented")
}
func (UnimplementedLinkStatServer) CheckDevice(context.Context, *CheckDeviceRequest) (*CheckDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDevice not implemented")
}
func (UnimplementedLinkStatServer) RoomLiveStat(context.Context, *RoomLiveStatRequest) (*RoomLiveStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomLiveStat not implemented")
}
func (UnimplementedLinkStatServer) CheckAccountRoom(context.Context, *CheckAccountRoomRequest) (*CheckAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccountRoom not implemented")
}
func (UnimplementedLinkStatServer) CheckSessionRoom(context.Context, *CheckSessionRoomRequest) (*CheckSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSessionRoom not implemented")
}
func (UnimplementedLinkStatServer) CheckDeviceRoom(context.Context, *CheckDeviceRoomRequest) (*CheckDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDeviceRoom not implemented")
}
func (UnimplementedLinkStatServer) Refresh(context.Context, *RefreshStatRequest) (*RefreshStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedLinkStatServer) mustEmbedUnimplementedLinkStatServer() {}
func (UnimplementedLinkStatServer) testEmbeddedByValue()                  {}

// UnsafeLinkStatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkStatServer will
// result in compilation errors.
type UnsafeLinkStatServer interface {
	mustEmbedUnimplementedLinkStatServer()
}

func RegisterLinkStatServer(s grpc.ServiceRegistrar, srv LinkStatServer) {
	// If the following call pancis, it indicates UnimplementedLinkStatServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LinkStat_ServiceDesc, srv)
}

func _LinkStat_OnlineSessionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).OnlineSessionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_OnlineSessionCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).OnlineSessionCount(ctx, req.(*OnlineCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_OnlineAccountCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).OnlineAccountCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_OnlineAccountCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).OnlineAccountCount(ctx, req.(*OnlineCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_OnlineDeviceCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).OnlineDeviceCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_OnlineDeviceCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).OnlineDeviceCount(ctx, req.(*OnlineCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_OnlineSessionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).OnlineSessionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_OnlineSessionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).OnlineSessionList(ctx, req.(*OnlineListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_OnlineAccountList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).OnlineAccountList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_OnlineAccountList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).OnlineAccountList(ctx, req.(*OnlineListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_OnlineDeviceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).OnlineDeviceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_OnlineDeviceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).OnlineDeviceList(ctx, req.(*OnlineListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_CheckSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).CheckSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_CheckSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).CheckSession(ctx, req.(*CheckSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_CheckAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).CheckAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_CheckAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).CheckAccount(ctx, req.(*CheckAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_CheckDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).CheckDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_CheckDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).CheckDevice(ctx, req.(*CheckDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_RoomLiveStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomLiveStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).RoomLiveStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_RoomLiveStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).RoomLiveStat(ctx, req.(*RoomLiveStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_CheckAccountRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccountRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).CheckAccountRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_CheckAccountRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).CheckAccountRoom(ctx, req.(*CheckAccountRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_CheckSessionRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSessionRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).CheckSessionRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_CheckSessionRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).CheckSessionRoom(ctx, req.(*CheckSessionRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_CheckDeviceRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDeviceRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).CheckDeviceRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_CheckDeviceRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).CheckDeviceRoom(ctx, req.(*CheckDeviceRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkStat_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStatServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkStat_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStatServer).Refresh(ctx, req.(*RefreshStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkStat_ServiceDesc is the grpc.ServiceDesc for LinkStat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkStat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.infra.link.LinkStat",
	HandlerType: (*LinkStatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnlineSessionCount",
			Handler:    _LinkStat_OnlineSessionCount_Handler,
		},
		{
			MethodName: "OnlineAccountCount",
			Handler:    _LinkStat_OnlineAccountCount_Handler,
		},
		{
			MethodName: "OnlineDeviceCount",
			Handler:    _LinkStat_OnlineDeviceCount_Handler,
		},
		{
			MethodName: "OnlineSessionList",
			Handler:    _LinkStat_OnlineSessionList_Handler,
		},
		{
			MethodName: "OnlineAccountList",
			Handler:    _LinkStat_OnlineAccountList_Handler,
		},
		{
			MethodName: "OnlineDeviceList",
			Handler:    _LinkStat_OnlineDeviceList_Handler,
		},
		{
			MethodName: "CheckSession",
			Handler:    _LinkStat_CheckSession_Handler,
		},
		{
			MethodName: "CheckAccount",
			Handler:    _LinkStat_CheckAccount_Handler,
		},
		{
			MethodName: "CheckDevice",
			Handler:    _LinkStat_CheckDevice_Handler,
		},
		{
			MethodName: "RoomLiveStat",
			Handler:    _LinkStat_RoomLiveStat_Handler,
		},
		{
			MethodName: "CheckAccountRoom",
			Handler:    _LinkStat_CheckAccountRoom_Handler,
		},
		{
			MethodName: "CheckSessionRoom",
			Handler:    _LinkStat_CheckSessionRoom_Handler,
		},
		{
			MethodName: "CheckDeviceRoom",
			Handler:    _LinkStat_CheckDeviceRoom_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _LinkStat_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.infra.link/stat.proto",
}
