// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: svc.biz.vip/fanbase_member.proto

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FanbaseMember_JoinFanbase_FullMethodName                        = "/svc.biz.vip.FanbaseMember/JoinFanbase"
	FanbaseMember_LeaveFanbase_FullMethodName                       = "/svc.biz.vip.FanbaseMember/LeaveFanbase"
	FanbaseMember_GetFanbaseMember_FullMethodName                   = "/svc.biz.vip.FanbaseMember/GetFanbaseMember"
	FanbaseMember_MGetFanbaseMember_FullMethodName                  = "/svc.biz.vip.FanbaseMember/MGetFanbaseMember"
	FanbaseMember_GetFanbaseMemberByStreamerID_FullMethodName       = "/svc.biz.vip.FanbaseMember/GetFanbaseMemberByStreamerID"
	FanbaseMember_CountFanbaseMemberByStreamerID_FullMethodName     = "/svc.biz.vip.FanbaseMember/CountFanbaseMemberByStreamerID"
	FanbaseMember_GetOnlineFanbaseMemberByStreamerID_FullMethodName = "/svc.biz.vip.FanbaseMember/GetOnlineFanbaseMemberByStreamerID"
	FanbaseMember_GetFanbaseMembertByMemberID_FullMethodName        = "/svc.biz.vip.FanbaseMember/GetFanbaseMembertByMemberID"
	FanbaseMember_CountFanbaseMembertByMemberID_FullMethodName      = "/svc.biz.vip.FanbaseMember/CountFanbaseMembertByMemberID"
	FanbaseMember_GetFanbaseOrders_FullMethodName                   = "/svc.biz.vip.FanbaseMember/GetFanbaseOrders"
	FanbaseMember_GetFanbaseOrderStat_FullMethodName                = "/svc.biz.vip.FanbaseMember/GetFanbaseOrderStat"
)

// FanbaseMemberClient is the client API for FanbaseMember service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 粉丝团
type FanbaseMemberClient interface {
	// JoinFanbase 加入粉丝团
	JoinFanbase(ctx context.Context, in *JoinFanbaseReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// LeaveFanbase 离开粉丝团
	LeaveFanbase(ctx context.Context, in *LeaveFanbaseReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetFanbaseMember 获取主播的某个粉丝团成员信息
	GetFanbaseMember(ctx context.Context, in *GetFanbaseMemberReq, opts ...grpc.CallOption) (*GetFanbaseMemberResp, error)
	// MGetFanbaseMember 批量获取主播的粉丝团成员信息
	MGetFanbaseMember(ctx context.Context, in *MGetFanbaseMemberReq, opts ...grpc.CallOption) (*MGetFanbaseMemberResp, error)
	// GetFanbaseMemberByStreamerID 获取主播粉丝团成员列表
	GetFanbaseMemberByStreamerID(ctx context.Context, in *GetFanbaseMemberByStreamerIDReq, opts ...grpc.CallOption) (*GetFanbaseMemberByStreamerIDResp, error)
	// CountFanbaseMemberByStreamerID 获取主播粉丝团成员总数
	CountFanbaseMemberByStreamerID(ctx context.Context, in *CountFanbaseMemberByStreamerIDReq, opts ...grpc.CallOption) (*CountFanbaseMemberByStreamerIDResp, error)
	// GetOnlineFanbaseMemberByStreamerID 获取主播粉丝团在线成员列表
	GetOnlineFanbaseMemberByStreamerID(ctx context.Context, in *GetOnlineFanbaseMemberByStreamerIDReq, opts ...grpc.CallOption) (*GetOnlineFanbaseMemberByStreamerIDResp, error)
	// GetFanbaseMembertByMemberID 获取用户加入的粉丝团列表
	GetFanbaseMembertByMemberID(ctx context.Context, in *GetFanbaseMembertByMemberIDReq, opts ...grpc.CallOption) (*GetFanbaseMembertByMemberIDResp, error)
	// CountFanbaseMembertByMemberID 获取用户加入的粉丝团数量
	CountFanbaseMembertByMemberID(ctx context.Context, in *CountFanbaseMembertByMemberIDReq, opts ...grpc.CallOption) (*CountFanbaseMembertByMemberIDResp, error)
	// GetFanbaseOrders 获取粉丝团订单列表
	GetFanbaseOrders(ctx context.Context, in *GetFanbaseOrdersReq, opts ...grpc.CallOption) (*GetFanbaseOrdersResp, error)
	// GetFanbaseOrderStat 获取粉丝团订单统计
	GetFanbaseOrderStat(ctx context.Context, in *GetFanbaseOrderStatReq, opts ...grpc.CallOption) (*GetFanbaseOrderStatResp, error)
}

type fanbaseMemberClient struct {
	cc grpc.ClientConnInterface
}

func NewFanbaseMemberClient(cc grpc.ClientConnInterface) FanbaseMemberClient {
	return &fanbaseMemberClient{cc}
}

func (c *fanbaseMemberClient) JoinFanbase(ctx context.Context, in *JoinFanbaseReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FanbaseMember_JoinFanbase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) LeaveFanbase(ctx context.Context, in *LeaveFanbaseReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FanbaseMember_LeaveFanbase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) GetFanbaseMember(ctx context.Context, in *GetFanbaseMemberReq, opts ...grpc.CallOption) (*GetFanbaseMemberResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFanbaseMemberResp)
	err := c.cc.Invoke(ctx, FanbaseMember_GetFanbaseMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) MGetFanbaseMember(ctx context.Context, in *MGetFanbaseMemberReq, opts ...grpc.CallOption) (*MGetFanbaseMemberResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MGetFanbaseMemberResp)
	err := c.cc.Invoke(ctx, FanbaseMember_MGetFanbaseMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) GetFanbaseMemberByStreamerID(ctx context.Context, in *GetFanbaseMemberByStreamerIDReq, opts ...grpc.CallOption) (*GetFanbaseMemberByStreamerIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFanbaseMemberByStreamerIDResp)
	err := c.cc.Invoke(ctx, FanbaseMember_GetFanbaseMemberByStreamerID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) CountFanbaseMemberByStreamerID(ctx context.Context, in *CountFanbaseMemberByStreamerIDReq, opts ...grpc.CallOption) (*CountFanbaseMemberByStreamerIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountFanbaseMemberByStreamerIDResp)
	err := c.cc.Invoke(ctx, FanbaseMember_CountFanbaseMemberByStreamerID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) GetOnlineFanbaseMemberByStreamerID(ctx context.Context, in *GetOnlineFanbaseMemberByStreamerIDReq, opts ...grpc.CallOption) (*GetOnlineFanbaseMemberByStreamerIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOnlineFanbaseMemberByStreamerIDResp)
	err := c.cc.Invoke(ctx, FanbaseMember_GetOnlineFanbaseMemberByStreamerID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) GetFanbaseMembertByMemberID(ctx context.Context, in *GetFanbaseMembertByMemberIDReq, opts ...grpc.CallOption) (*GetFanbaseMembertByMemberIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFanbaseMembertByMemberIDResp)
	err := c.cc.Invoke(ctx, FanbaseMember_GetFanbaseMembertByMemberID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) CountFanbaseMembertByMemberID(ctx context.Context, in *CountFanbaseMembertByMemberIDReq, opts ...grpc.CallOption) (*CountFanbaseMembertByMemberIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountFanbaseMembertByMemberIDResp)
	err := c.cc.Invoke(ctx, FanbaseMember_CountFanbaseMembertByMemberID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) GetFanbaseOrders(ctx context.Context, in *GetFanbaseOrdersReq, opts ...grpc.CallOption) (*GetFanbaseOrdersResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFanbaseOrdersResp)
	err := c.cc.Invoke(ctx, FanbaseMember_GetFanbaseOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberClient) GetFanbaseOrderStat(ctx context.Context, in *GetFanbaseOrderStatReq, opts ...grpc.CallOption) (*GetFanbaseOrderStatResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFanbaseOrderStatResp)
	err := c.cc.Invoke(ctx, FanbaseMember_GetFanbaseOrderStat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FanbaseMemberServer is the server API for FanbaseMember service.
// All implementations must embed UnimplementedFanbaseMemberServer
// for forward compatibility.
//
// 粉丝团
type FanbaseMemberServer interface {
	// JoinFanbase 加入粉丝团
	JoinFanbase(context.Context, *JoinFanbaseReq) (*emptypb.Empty, error)
	// LeaveFanbase 离开粉丝团
	LeaveFanbase(context.Context, *LeaveFanbaseReq) (*emptypb.Empty, error)
	// GetFanbaseMember 获取主播的某个粉丝团成员信息
	GetFanbaseMember(context.Context, *GetFanbaseMemberReq) (*GetFanbaseMemberResp, error)
	// MGetFanbaseMember 批量获取主播的粉丝团成员信息
	MGetFanbaseMember(context.Context, *MGetFanbaseMemberReq) (*MGetFanbaseMemberResp, error)
	// GetFanbaseMemberByStreamerID 获取主播粉丝团成员列表
	GetFanbaseMemberByStreamerID(context.Context, *GetFanbaseMemberByStreamerIDReq) (*GetFanbaseMemberByStreamerIDResp, error)
	// CountFanbaseMemberByStreamerID 获取主播粉丝团成员总数
	CountFanbaseMemberByStreamerID(context.Context, *CountFanbaseMemberByStreamerIDReq) (*CountFanbaseMemberByStreamerIDResp, error)
	// GetOnlineFanbaseMemberByStreamerID 获取主播粉丝团在线成员列表
	GetOnlineFanbaseMemberByStreamerID(context.Context, *GetOnlineFanbaseMemberByStreamerIDReq) (*GetOnlineFanbaseMemberByStreamerIDResp, error)
	// GetFanbaseMembertByMemberID 获取用户加入的粉丝团列表
	GetFanbaseMembertByMemberID(context.Context, *GetFanbaseMembertByMemberIDReq) (*GetFanbaseMembertByMemberIDResp, error)
	// CountFanbaseMembertByMemberID 获取用户加入的粉丝团数量
	CountFanbaseMembertByMemberID(context.Context, *CountFanbaseMembertByMemberIDReq) (*CountFanbaseMembertByMemberIDResp, error)
	// GetFanbaseOrders 获取粉丝团订单列表
	GetFanbaseOrders(context.Context, *GetFanbaseOrdersReq) (*GetFanbaseOrdersResp, error)
	// GetFanbaseOrderStat 获取粉丝团订单统计
	GetFanbaseOrderStat(context.Context, *GetFanbaseOrderStatReq) (*GetFanbaseOrderStatResp, error)
	mustEmbedUnimplementedFanbaseMemberServer()
}

// UnimplementedFanbaseMemberServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFanbaseMemberServer struct{}

func (UnimplementedFanbaseMemberServer) JoinFanbase(context.Context, *JoinFanbaseReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinFanbase not implemented")
}
func (UnimplementedFanbaseMemberServer) LeaveFanbase(context.Context, *LeaveFanbaseReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveFanbase not implemented")
}
func (UnimplementedFanbaseMemberServer) GetFanbaseMember(context.Context, *GetFanbaseMemberReq) (*GetFanbaseMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFanbaseMember not implemented")
}
func (UnimplementedFanbaseMemberServer) MGetFanbaseMember(context.Context, *MGetFanbaseMemberReq) (*MGetFanbaseMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetFanbaseMember not implemented")
}
func (UnimplementedFanbaseMemberServer) GetFanbaseMemberByStreamerID(context.Context, *GetFanbaseMemberByStreamerIDReq) (*GetFanbaseMemberByStreamerIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFanbaseMemberByStreamerID not implemented")
}
func (UnimplementedFanbaseMemberServer) CountFanbaseMemberByStreamerID(context.Context, *CountFanbaseMemberByStreamerIDReq) (*CountFanbaseMemberByStreamerIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFanbaseMemberByStreamerID not implemented")
}
func (UnimplementedFanbaseMemberServer) GetOnlineFanbaseMemberByStreamerID(context.Context, *GetOnlineFanbaseMemberByStreamerIDReq) (*GetOnlineFanbaseMemberByStreamerIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineFanbaseMemberByStreamerID not implemented")
}
func (UnimplementedFanbaseMemberServer) GetFanbaseMembertByMemberID(context.Context, *GetFanbaseMembertByMemberIDReq) (*GetFanbaseMembertByMemberIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFanbaseMembertByMemberID not implemented")
}
func (UnimplementedFanbaseMemberServer) CountFanbaseMembertByMemberID(context.Context, *CountFanbaseMembertByMemberIDReq) (*CountFanbaseMembertByMemberIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFanbaseMembertByMemberID not implemented")
}
func (UnimplementedFanbaseMemberServer) GetFanbaseOrders(context.Context, *GetFanbaseOrdersReq) (*GetFanbaseOrdersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFanbaseOrders not implemented")
}
func (UnimplementedFanbaseMemberServer) GetFanbaseOrderStat(context.Context, *GetFanbaseOrderStatReq) (*GetFanbaseOrderStatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFanbaseOrderStat not implemented")
}
func (UnimplementedFanbaseMemberServer) mustEmbedUnimplementedFanbaseMemberServer() {}
func (UnimplementedFanbaseMemberServer) testEmbeddedByValue()                       {}

// UnsafeFanbaseMemberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FanbaseMemberServer will
// result in compilation errors.
type UnsafeFanbaseMemberServer interface {
	mustEmbedUnimplementedFanbaseMemberServer()
}

func RegisterFanbaseMemberServer(s grpc.ServiceRegistrar, srv FanbaseMemberServer) {
	// If the following call pancis, it indicates UnimplementedFanbaseMemberServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FanbaseMember_ServiceDesc, srv)
}

func _FanbaseMember_JoinFanbase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinFanbaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).JoinFanbase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_JoinFanbase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).JoinFanbase(ctx, req.(*JoinFanbaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_LeaveFanbase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveFanbaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).LeaveFanbase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_LeaveFanbase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).LeaveFanbase(ctx, req.(*LeaveFanbaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_GetFanbaseMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFanbaseMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).GetFanbaseMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_GetFanbaseMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).GetFanbaseMember(ctx, req.(*GetFanbaseMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_MGetFanbaseMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetFanbaseMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).MGetFanbaseMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_MGetFanbaseMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).MGetFanbaseMember(ctx, req.(*MGetFanbaseMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_GetFanbaseMemberByStreamerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFanbaseMemberByStreamerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).GetFanbaseMemberByStreamerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_GetFanbaseMemberByStreamerID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).GetFanbaseMemberByStreamerID(ctx, req.(*GetFanbaseMemberByStreamerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_CountFanbaseMemberByStreamerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFanbaseMemberByStreamerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).CountFanbaseMemberByStreamerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_CountFanbaseMemberByStreamerID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).CountFanbaseMemberByStreamerID(ctx, req.(*CountFanbaseMemberByStreamerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_GetOnlineFanbaseMemberByStreamerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOnlineFanbaseMemberByStreamerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).GetOnlineFanbaseMemberByStreamerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_GetOnlineFanbaseMemberByStreamerID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).GetOnlineFanbaseMemberByStreamerID(ctx, req.(*GetOnlineFanbaseMemberByStreamerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_GetFanbaseMembertByMemberID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFanbaseMembertByMemberIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).GetFanbaseMembertByMemberID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_GetFanbaseMembertByMemberID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).GetFanbaseMembertByMemberID(ctx, req.(*GetFanbaseMembertByMemberIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_CountFanbaseMembertByMemberID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFanbaseMembertByMemberIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).CountFanbaseMembertByMemberID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_CountFanbaseMembertByMemberID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).CountFanbaseMembertByMemberID(ctx, req.(*CountFanbaseMembertByMemberIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_GetFanbaseOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFanbaseOrdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).GetFanbaseOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_GetFanbaseOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).GetFanbaseOrders(ctx, req.(*GetFanbaseOrdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FanbaseMember_GetFanbaseOrderStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFanbaseOrderStatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanbaseMemberServer).GetFanbaseOrderStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FanbaseMember_GetFanbaseOrderStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanbaseMemberServer).GetFanbaseOrderStat(ctx, req.(*GetFanbaseOrderStatReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FanbaseMember_ServiceDesc is the grpc.ServiceDesc for FanbaseMember service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FanbaseMember_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.vip.FanbaseMember",
	HandlerType: (*FanbaseMemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinFanbase",
			Handler:    _FanbaseMember_JoinFanbase_Handler,
		},
		{
			MethodName: "LeaveFanbase",
			Handler:    _FanbaseMember_LeaveFanbase_Handler,
		},
		{
			MethodName: "GetFanbaseMember",
			Handler:    _FanbaseMember_GetFanbaseMember_Handler,
		},
		{
			MethodName: "MGetFanbaseMember",
			Handler:    _FanbaseMember_MGetFanbaseMember_Handler,
		},
		{
			MethodName: "GetFanbaseMemberByStreamerID",
			Handler:    _FanbaseMember_GetFanbaseMemberByStreamerID_Handler,
		},
		{
			MethodName: "CountFanbaseMemberByStreamerID",
			Handler:    _FanbaseMember_CountFanbaseMemberByStreamerID_Handler,
		},
		{
			MethodName: "GetOnlineFanbaseMemberByStreamerID",
			Handler:    _FanbaseMember_GetOnlineFanbaseMemberByStreamerID_Handler,
		},
		{
			MethodName: "GetFanbaseMembertByMemberID",
			Handler:    _FanbaseMember_GetFanbaseMembertByMemberID_Handler,
		},
		{
			MethodName: "CountFanbaseMembertByMemberID",
			Handler:    _FanbaseMember_CountFanbaseMembertByMemberID_Handler,
		},
		{
			MethodName: "GetFanbaseOrders",
			Handler:    _FanbaseMember_GetFanbaseOrders_Handler,
		},
		{
			MethodName: "GetFanbaseOrderStat",
			Handler:    _FanbaseMember_GetFanbaseOrderStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.vip/fanbase_member.proto",
}
