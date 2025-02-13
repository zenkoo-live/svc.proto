// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: svc.biz.vip/noble_member.proto

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
	NobleMember_JoinNoble_FullMethodName                            = "/svc.biz.vip.NobleMember/JoinNoble"
	NobleMember_RenewNoble_FullMethodName                           = "/svc.biz.vip.NobleMember/RenewNoble"
	NobleMember_UpgradeNoble_FullMethodName                         = "/svc.biz.vip.NobleMember/UpgradeNoble"
	NobleMember_GetNobleMember_FullMethodName                       = "/svc.biz.vip.NobleMember/GetNobleMember"
	NobleMember_MGetNobleMember_FullMethodName                      = "/svc.biz.vip.NobleMember/MGetNobleMember"
	NobleMember_GetNobleMemberList_FullMethodName                   = "/svc.biz.vip.NobleMember/GetNobleMemberList"
	NobleMember_CountNobleMember_FullMethodName                     = "/svc.biz.vip.NobleMember/CountNobleMember"
	NobleMember_GetOnlineNobleMemberListByStreamerID_FullMethodName = "/svc.biz.vip.NobleMember/GetOnlineNobleMemberListByStreamerID"
	NobleMember_GetNobleOrders_FullMethodName                       = "/svc.biz.vip.NobleMember/GetNobleOrders"
	NobleMember_GetNobleOrderStat_FullMethodName                    = "/svc.biz.vip.NobleMember/GetNobleOrderStat"
)

// NobleMemberClient is the client API for NobleMember service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 贵族
type NobleMemberClient interface {
	// JoinNoble 加入贵族
	JoinNoble(ctx context.Context, in *JoinNobleReq, opts ...grpc.CallOption) (*JoinNobleResp, error)
	// RenewNoble 续费贵族
	RenewNoble(ctx context.Context, in *RenewNobleReq, opts ...grpc.CallOption) (*RenewNobleResp, error)
	// UpgradeNoble 升级贵族
	UpgradeNoble(ctx context.Context, in *UpgradeNobleReq, opts ...grpc.CallOption) (*UpgradeNobleResp, error)
	// GetNobleMember 获取成员贵族信息
	GetNobleMember(ctx context.Context, in *GetNobleMemberReq, opts ...grpc.CallOption) (*GetNobleMemberResp, error)
	// MGetNobleMember 批量获取成员贵族信息
	MGetNobleMember(ctx context.Context, in *MGetNobleMemberReq, opts ...grpc.CallOption) (*MGetNobleMemberResp, error)
	// GetNobleMemberList 获取贵族成员列表（streamer_id传空字符串取所有）
	GetNobleMemberList(ctx context.Context, in *GetNobleMemberListReq, opts ...grpc.CallOption) (*GetNobleMemberListResp, error)
	// CountNobleMember 获取成员总数
	CountNobleMember(ctx context.Context, in *CountNobleMemberReq, opts ...grpc.CallOption) (*CountNobleMemberResp, error)
	// GetOnlineNobleMemberListByStreamerID 获取主播贵族在线成员列表
	GetOnlineNobleMemberListByStreamerID(ctx context.Context, in *GetOnlineNobleMemberListByStreamerIDReq, opts ...grpc.CallOption) (*GetNobleMemberListResp, error)
	// GetNobleOrders 获取贵族订单列表
	GetNobleOrders(ctx context.Context, in *GetNobleOrdersReq, opts ...grpc.CallOption) (*GetNobleOrdersResp, error)
	// GetNobleOrderStat 获取贵族订单统计
	GetNobleOrderStat(ctx context.Context, in *GetNobleOrderStatReq, opts ...grpc.CallOption) (*GetNobleOrderStatResp, error)
}

type nobleMemberClient struct {
	cc grpc.ClientConnInterface
}

func NewNobleMemberClient(cc grpc.ClientConnInterface) NobleMemberClient {
	return &nobleMemberClient{cc}
}

func (c *nobleMemberClient) JoinNoble(ctx context.Context, in *JoinNobleReq, opts ...grpc.CallOption) (*JoinNobleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinNobleResp)
	err := c.cc.Invoke(ctx, NobleMember_JoinNoble_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberClient) RenewNoble(ctx context.Context, in *RenewNobleReq, opts ...grpc.CallOption) (*RenewNobleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RenewNobleResp)
	err := c.cc.Invoke(ctx, NobleMember_RenewNoble_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberClient) UpgradeNoble(ctx context.Context, in *UpgradeNobleReq, opts ...grpc.CallOption) (*UpgradeNobleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpgradeNobleResp)
	err := c.cc.Invoke(ctx, NobleMember_UpgradeNoble_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberClient) GetNobleMember(ctx context.Context, in *GetNobleMemberReq, opts ...grpc.CallOption) (*GetNobleMemberResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNobleMemberResp)
	err := c.cc.Invoke(ctx, NobleMember_GetNobleMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberClient) MGetNobleMember(ctx context.Context, in *MGetNobleMemberReq, opts ...grpc.CallOption) (*MGetNobleMemberResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MGetNobleMemberResp)
	err := c.cc.Invoke(ctx, NobleMember_MGetNobleMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberClient) GetNobleMemberList(ctx context.Context, in *GetNobleMemberListReq, opts ...grpc.CallOption) (*GetNobleMemberListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNobleMemberListResp)
	err := c.cc.Invoke(ctx, NobleMember_GetNobleMemberList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberClient) CountNobleMember(ctx context.Context, in *CountNobleMemberReq, opts ...grpc.CallOption) (*CountNobleMemberResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountNobleMemberResp)
	err := c.cc.Invoke(ctx, NobleMember_CountNobleMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberClient) GetOnlineNobleMemberListByStreamerID(ctx context.Context, in *GetOnlineNobleMemberListByStreamerIDReq, opts ...grpc.CallOption) (*GetNobleMemberListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNobleMemberListResp)
	err := c.cc.Invoke(ctx, NobleMember_GetOnlineNobleMemberListByStreamerID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberClient) GetNobleOrders(ctx context.Context, in *GetNobleOrdersReq, opts ...grpc.CallOption) (*GetNobleOrdersResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNobleOrdersResp)
	err := c.cc.Invoke(ctx, NobleMember_GetNobleOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberClient) GetNobleOrderStat(ctx context.Context, in *GetNobleOrderStatReq, opts ...grpc.CallOption) (*GetNobleOrderStatResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNobleOrderStatResp)
	err := c.cc.Invoke(ctx, NobleMember_GetNobleOrderStat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NobleMemberServer is the server API for NobleMember service.
// All implementations must embed UnimplementedNobleMemberServer
// for forward compatibility.
//
// 贵族
type NobleMemberServer interface {
	// JoinNoble 加入贵族
	JoinNoble(context.Context, *JoinNobleReq) (*JoinNobleResp, error)
	// RenewNoble 续费贵族
	RenewNoble(context.Context, *RenewNobleReq) (*RenewNobleResp, error)
	// UpgradeNoble 升级贵族
	UpgradeNoble(context.Context, *UpgradeNobleReq) (*UpgradeNobleResp, error)
	// GetNobleMember 获取成员贵族信息
	GetNobleMember(context.Context, *GetNobleMemberReq) (*GetNobleMemberResp, error)
	// MGetNobleMember 批量获取成员贵族信息
	MGetNobleMember(context.Context, *MGetNobleMemberReq) (*MGetNobleMemberResp, error)
	// GetNobleMemberList 获取贵族成员列表（streamer_id传空字符串取所有）
	GetNobleMemberList(context.Context, *GetNobleMemberListReq) (*GetNobleMemberListResp, error)
	// CountNobleMember 获取成员总数
	CountNobleMember(context.Context, *CountNobleMemberReq) (*CountNobleMemberResp, error)
	// GetOnlineNobleMemberListByStreamerID 获取主播贵族在线成员列表
	GetOnlineNobleMemberListByStreamerID(context.Context, *GetOnlineNobleMemberListByStreamerIDReq) (*GetNobleMemberListResp, error)
	// GetNobleOrders 获取贵族订单列表
	GetNobleOrders(context.Context, *GetNobleOrdersReq) (*GetNobleOrdersResp, error)
	// GetNobleOrderStat 获取贵族订单统计
	GetNobleOrderStat(context.Context, *GetNobleOrderStatReq) (*GetNobleOrderStatResp, error)
	mustEmbedUnimplementedNobleMemberServer()
}

// UnimplementedNobleMemberServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNobleMemberServer struct{}

func (UnimplementedNobleMemberServer) JoinNoble(context.Context, *JoinNobleReq) (*JoinNobleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinNoble not implemented")
}
func (UnimplementedNobleMemberServer) RenewNoble(context.Context, *RenewNobleReq) (*RenewNobleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewNoble not implemented")
}
func (UnimplementedNobleMemberServer) UpgradeNoble(context.Context, *UpgradeNobleReq) (*UpgradeNobleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeNoble not implemented")
}
func (UnimplementedNobleMemberServer) GetNobleMember(context.Context, *GetNobleMemberReq) (*GetNobleMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNobleMember not implemented")
}
func (UnimplementedNobleMemberServer) MGetNobleMember(context.Context, *MGetNobleMemberReq) (*MGetNobleMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetNobleMember not implemented")
}
func (UnimplementedNobleMemberServer) GetNobleMemberList(context.Context, *GetNobleMemberListReq) (*GetNobleMemberListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNobleMemberList not implemented")
}
func (UnimplementedNobleMemberServer) CountNobleMember(context.Context, *CountNobleMemberReq) (*CountNobleMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountNobleMember not implemented")
}
func (UnimplementedNobleMemberServer) GetOnlineNobleMemberListByStreamerID(context.Context, *GetOnlineNobleMemberListByStreamerIDReq) (*GetNobleMemberListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineNobleMemberListByStreamerID not implemented")
}
func (UnimplementedNobleMemberServer) GetNobleOrders(context.Context, *GetNobleOrdersReq) (*GetNobleOrdersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNobleOrders not implemented")
}
func (UnimplementedNobleMemberServer) GetNobleOrderStat(context.Context, *GetNobleOrderStatReq) (*GetNobleOrderStatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNobleOrderStat not implemented")
}
func (UnimplementedNobleMemberServer) mustEmbedUnimplementedNobleMemberServer() {}
func (UnimplementedNobleMemberServer) testEmbeddedByValue()                     {}

// UnsafeNobleMemberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NobleMemberServer will
// result in compilation errors.
type UnsafeNobleMemberServer interface {
	mustEmbedUnimplementedNobleMemberServer()
}

func RegisterNobleMemberServer(s grpc.ServiceRegistrar, srv NobleMemberServer) {
	// If the following call pancis, it indicates UnimplementedNobleMemberServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NobleMember_ServiceDesc, srv)
}

func _NobleMember_JoinNoble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinNobleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).JoinNoble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_JoinNoble_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).JoinNoble(ctx, req.(*JoinNobleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleMember_RenewNoble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewNobleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).RenewNoble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_RenewNoble_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).RenewNoble(ctx, req.(*RenewNobleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleMember_UpgradeNoble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeNobleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).UpgradeNoble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_UpgradeNoble_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).UpgradeNoble(ctx, req.(*UpgradeNobleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleMember_GetNobleMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNobleMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).GetNobleMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_GetNobleMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).GetNobleMember(ctx, req.(*GetNobleMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleMember_MGetNobleMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetNobleMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).MGetNobleMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_MGetNobleMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).MGetNobleMember(ctx, req.(*MGetNobleMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleMember_GetNobleMemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNobleMemberListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).GetNobleMemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_GetNobleMemberList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).GetNobleMemberList(ctx, req.(*GetNobleMemberListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleMember_CountNobleMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountNobleMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).CountNobleMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_CountNobleMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).CountNobleMember(ctx, req.(*CountNobleMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleMember_GetOnlineNobleMemberListByStreamerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOnlineNobleMemberListByStreamerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).GetOnlineNobleMemberListByStreamerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_GetOnlineNobleMemberListByStreamerID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).GetOnlineNobleMemberListByStreamerID(ctx, req.(*GetOnlineNobleMemberListByStreamerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleMember_GetNobleOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNobleOrdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).GetNobleOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_GetNobleOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).GetNobleOrders(ctx, req.(*GetNobleOrdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NobleMember_GetNobleOrderStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNobleOrderStatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NobleMemberServer).GetNobleOrderStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NobleMember_GetNobleOrderStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NobleMemberServer).GetNobleOrderStat(ctx, req.(*GetNobleOrderStatReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NobleMember_ServiceDesc is the grpc.ServiceDesc for NobleMember service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NobleMember_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.vip.NobleMember",
	HandlerType: (*NobleMemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinNoble",
			Handler:    _NobleMember_JoinNoble_Handler,
		},
		{
			MethodName: "RenewNoble",
			Handler:    _NobleMember_RenewNoble_Handler,
		},
		{
			MethodName: "UpgradeNoble",
			Handler:    _NobleMember_UpgradeNoble_Handler,
		},
		{
			MethodName: "GetNobleMember",
			Handler:    _NobleMember_GetNobleMember_Handler,
		},
		{
			MethodName: "MGetNobleMember",
			Handler:    _NobleMember_MGetNobleMember_Handler,
		},
		{
			MethodName: "GetNobleMemberList",
			Handler:    _NobleMember_GetNobleMemberList_Handler,
		},
		{
			MethodName: "CountNobleMember",
			Handler:    _NobleMember_CountNobleMember_Handler,
		},
		{
			MethodName: "GetOnlineNobleMemberListByStreamerID",
			Handler:    _NobleMember_GetOnlineNobleMemberListByStreamerID_Handler,
		},
		{
			MethodName: "GetNobleOrders",
			Handler:    _NobleMember_GetNobleOrders_Handler,
		},
		{
			MethodName: "GetNobleOrderStat",
			Handler:    _NobleMember_GetNobleOrderStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.vip/noble_member.proto",
}
