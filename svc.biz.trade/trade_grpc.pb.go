// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: svc.biz.trade/trade.proto

package trade

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
	Trade_SendGiftInLive_FullMethodName      = "/svc.biz.trade.Trade/SendGiftInLive"
	Trade_BuyLiveTicket_FullMethodName       = "/svc.biz.trade.Trade/BuyLiveTicket"
	Trade_PayLiveDurationFee_FullMethodName  = "/svc.biz.trade.Trade/PayLiveDurationFee"
	Trade_JoinAnchorFansGroup_FullMethodName = "/svc.biz.trade.Trade/JoinAnchorFansGroup"
	Trade_PayBulletChat_FullMethodName       = "/svc.biz.trade.Trade/PayBulletChat"
	Trade_VipActivate_FullMethodName         = "/svc.biz.trade.Trade/VipActivate"
	Trade_VipExtend_FullMethodName           = "/svc.biz.trade.Trade/VipExtend"
	Trade_VipUpgrade_FullMethodName          = "/svc.biz.trade.Trade/VipUpgrade"
	Trade_BuyRide_FullMethodName             = "/svc.biz.trade.Trade/BuyRide"
	Trade_BuyLuckyId_FullMethodName          = "/svc.biz.trade.Trade/BuyLuckyId"
	Trade_MoneyRecharge_FullMethodName       = "/svc.biz.trade.Trade/MoneyRecharge"
	Trade_MoneyWithdraw_FullMethodName       = "/svc.biz.trade.Trade/MoneyWithdraw"
	Trade_MoneyExchangeCoin_FullMethodName   = "/svc.biz.trade.Trade/MoneyExchangeCoin"
)

// TradeClient is the client API for Trade service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradeClient interface {
	SendGiftInLive(ctx context.Context, in *SendGiftInLiveReq, opts ...grpc.CallOption) (*SendGiftInLiveResp, error)
	BuyLiveTicket(ctx context.Context, in *BuyLiveTicketReq, opts ...grpc.CallOption) (*BuyLiveTicketResp, error)
	PayLiveDurationFee(ctx context.Context, in *PayLiveDurationFeeReq, opts ...grpc.CallOption) (*PayLiveDurationFeeResp, error)
	JoinAnchorFansGroup(ctx context.Context, in *JoinAnchorFansGroupReq, opts ...grpc.CallOption) (*JoinAnchorFansGroupResp, error)
	PayBulletChat(ctx context.Context, in *PayBulletChatReq, opts ...grpc.CallOption) (*PayBulletChatResp, error)
	VipActivate(ctx context.Context, in *VipActivateReq, opts ...grpc.CallOption) (*VipActivateResp, error)
	VipExtend(ctx context.Context, in *VipExtendReq, opts ...grpc.CallOption) (*VipExtendResp, error)
	VipUpgrade(ctx context.Context, in *VipUpgradeReq, opts ...grpc.CallOption) (*VipUpgradeResp, error)
	BuyRide(ctx context.Context, in *BuyRideReq, opts ...grpc.CallOption) (*BuyRideResp, error)
	BuyLuckyId(ctx context.Context, in *BuyLuckyIdReq, opts ...grpc.CallOption) (*BuyLuckyIdResp, error)
	MoneyRecharge(ctx context.Context, in *MoneyRechargeReq, opts ...grpc.CallOption) (*MoneyRechargeResp, error)
	MoneyWithdraw(ctx context.Context, in *MoneyWithdrawReq, opts ...grpc.CallOption) (*MoneyWithdrawResp, error)
	MoneyExchangeCoin(ctx context.Context, in *MoneyExchangeCoinReq, opts ...grpc.CallOption) (*MoneyExchangeCoinResp, error)
}

type tradeClient struct {
	cc grpc.ClientConnInterface
}

func NewTradeClient(cc grpc.ClientConnInterface) TradeClient {
	return &tradeClient{cc}
}

func (c *tradeClient) SendGiftInLive(ctx context.Context, in *SendGiftInLiveReq, opts ...grpc.CallOption) (*SendGiftInLiveResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendGiftInLiveResp)
	err := c.cc.Invoke(ctx, Trade_SendGiftInLive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) BuyLiveTicket(ctx context.Context, in *BuyLiveTicketReq, opts ...grpc.CallOption) (*BuyLiveTicketResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BuyLiveTicketResp)
	err := c.cc.Invoke(ctx, Trade_BuyLiveTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) PayLiveDurationFee(ctx context.Context, in *PayLiveDurationFeeReq, opts ...grpc.CallOption) (*PayLiveDurationFeeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayLiveDurationFeeResp)
	err := c.cc.Invoke(ctx, Trade_PayLiveDurationFee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) JoinAnchorFansGroup(ctx context.Context, in *JoinAnchorFansGroupReq, opts ...grpc.CallOption) (*JoinAnchorFansGroupResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinAnchorFansGroupResp)
	err := c.cc.Invoke(ctx, Trade_JoinAnchorFansGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) PayBulletChat(ctx context.Context, in *PayBulletChatReq, opts ...grpc.CallOption) (*PayBulletChatResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayBulletChatResp)
	err := c.cc.Invoke(ctx, Trade_PayBulletChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) VipActivate(ctx context.Context, in *VipActivateReq, opts ...grpc.CallOption) (*VipActivateResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VipActivateResp)
	err := c.cc.Invoke(ctx, Trade_VipActivate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) VipExtend(ctx context.Context, in *VipExtendReq, opts ...grpc.CallOption) (*VipExtendResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VipExtendResp)
	err := c.cc.Invoke(ctx, Trade_VipExtend_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) VipUpgrade(ctx context.Context, in *VipUpgradeReq, opts ...grpc.CallOption) (*VipUpgradeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VipUpgradeResp)
	err := c.cc.Invoke(ctx, Trade_VipUpgrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) BuyRide(ctx context.Context, in *BuyRideReq, opts ...grpc.CallOption) (*BuyRideResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BuyRideResp)
	err := c.cc.Invoke(ctx, Trade_BuyRide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) BuyLuckyId(ctx context.Context, in *BuyLuckyIdReq, opts ...grpc.CallOption) (*BuyLuckyIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BuyLuckyIdResp)
	err := c.cc.Invoke(ctx, Trade_BuyLuckyId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) MoneyRecharge(ctx context.Context, in *MoneyRechargeReq, opts ...grpc.CallOption) (*MoneyRechargeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MoneyRechargeResp)
	err := c.cc.Invoke(ctx, Trade_MoneyRecharge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) MoneyWithdraw(ctx context.Context, in *MoneyWithdrawReq, opts ...grpc.CallOption) (*MoneyWithdrawResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MoneyWithdrawResp)
	err := c.cc.Invoke(ctx, Trade_MoneyWithdraw_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) MoneyExchangeCoin(ctx context.Context, in *MoneyExchangeCoinReq, opts ...grpc.CallOption) (*MoneyExchangeCoinResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MoneyExchangeCoinResp)
	err := c.cc.Invoke(ctx, Trade_MoneyExchangeCoin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradeServer is the server API for Trade service.
// All implementations must embed UnimplementedTradeServer
// for forward compatibility
type TradeServer interface {
	SendGiftInLive(context.Context, *SendGiftInLiveReq) (*SendGiftInLiveResp, error)
	BuyLiveTicket(context.Context, *BuyLiveTicketReq) (*BuyLiveTicketResp, error)
	PayLiveDurationFee(context.Context, *PayLiveDurationFeeReq) (*PayLiveDurationFeeResp, error)
	JoinAnchorFansGroup(context.Context, *JoinAnchorFansGroupReq) (*JoinAnchorFansGroupResp, error)
	PayBulletChat(context.Context, *PayBulletChatReq) (*PayBulletChatResp, error)
	VipActivate(context.Context, *VipActivateReq) (*VipActivateResp, error)
	VipExtend(context.Context, *VipExtendReq) (*VipExtendResp, error)
	VipUpgrade(context.Context, *VipUpgradeReq) (*VipUpgradeResp, error)
	BuyRide(context.Context, *BuyRideReq) (*BuyRideResp, error)
	BuyLuckyId(context.Context, *BuyLuckyIdReq) (*BuyLuckyIdResp, error)
	MoneyRecharge(context.Context, *MoneyRechargeReq) (*MoneyRechargeResp, error)
	MoneyWithdraw(context.Context, *MoneyWithdrawReq) (*MoneyWithdrawResp, error)
	MoneyExchangeCoin(context.Context, *MoneyExchangeCoinReq) (*MoneyExchangeCoinResp, error)
	mustEmbedUnimplementedTradeServer()
}

// UnimplementedTradeServer must be embedded to have forward compatible implementations.
type UnimplementedTradeServer struct {
}

func (UnimplementedTradeServer) SendGiftInLive(context.Context, *SendGiftInLiveReq) (*SendGiftInLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGiftInLive not implemented")
}
func (UnimplementedTradeServer) BuyLiveTicket(context.Context, *BuyLiveTicketReq) (*BuyLiveTicketResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyLiveTicket not implemented")
}
func (UnimplementedTradeServer) PayLiveDurationFee(context.Context, *PayLiveDurationFeeReq) (*PayLiveDurationFeeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayLiveDurationFee not implemented")
}
func (UnimplementedTradeServer) JoinAnchorFansGroup(context.Context, *JoinAnchorFansGroupReq) (*JoinAnchorFansGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinAnchorFansGroup not implemented")
}
func (UnimplementedTradeServer) PayBulletChat(context.Context, *PayBulletChatReq) (*PayBulletChatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayBulletChat not implemented")
}
func (UnimplementedTradeServer) VipActivate(context.Context, *VipActivateReq) (*VipActivateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VipActivate not implemented")
}
func (UnimplementedTradeServer) VipExtend(context.Context, *VipExtendReq) (*VipExtendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VipExtend not implemented")
}
func (UnimplementedTradeServer) VipUpgrade(context.Context, *VipUpgradeReq) (*VipUpgradeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VipUpgrade not implemented")
}
func (UnimplementedTradeServer) BuyRide(context.Context, *BuyRideReq) (*BuyRideResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyRide not implemented")
}
func (UnimplementedTradeServer) BuyLuckyId(context.Context, *BuyLuckyIdReq) (*BuyLuckyIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyLuckyId not implemented")
}
func (UnimplementedTradeServer) MoneyRecharge(context.Context, *MoneyRechargeReq) (*MoneyRechargeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoneyRecharge not implemented")
}
func (UnimplementedTradeServer) MoneyWithdraw(context.Context, *MoneyWithdrawReq) (*MoneyWithdrawResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoneyWithdraw not implemented")
}
func (UnimplementedTradeServer) MoneyExchangeCoin(context.Context, *MoneyExchangeCoinReq) (*MoneyExchangeCoinResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoneyExchangeCoin not implemented")
}
func (UnimplementedTradeServer) mustEmbedUnimplementedTradeServer() {}

// UnsafeTradeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradeServer will
// result in compilation errors.
type UnsafeTradeServer interface {
	mustEmbedUnimplementedTradeServer()
}

func RegisterTradeServer(s grpc.ServiceRegistrar, srv TradeServer) {
	s.RegisterService(&Trade_ServiceDesc, srv)
}

func _Trade_SendGiftInLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGiftInLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).SendGiftInLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_SendGiftInLive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).SendGiftInLive(ctx, req.(*SendGiftInLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_BuyLiveTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyLiveTicketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).BuyLiveTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_BuyLiveTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).BuyLiveTicket(ctx, req.(*BuyLiveTicketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_PayLiveDurationFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayLiveDurationFeeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).PayLiveDurationFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_PayLiveDurationFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).PayLiveDurationFee(ctx, req.(*PayLiveDurationFeeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_JoinAnchorFansGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinAnchorFansGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).JoinAnchorFansGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_JoinAnchorFansGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).JoinAnchorFansGroup(ctx, req.(*JoinAnchorFansGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_PayBulletChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayBulletChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).PayBulletChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_PayBulletChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).PayBulletChat(ctx, req.(*PayBulletChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_VipActivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipActivateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).VipActivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_VipActivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).VipActivate(ctx, req.(*VipActivateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_VipExtend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipExtendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).VipExtend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_VipExtend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).VipExtend(ctx, req.(*VipExtendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_VipUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipUpgradeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).VipUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_VipUpgrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).VipUpgrade(ctx, req.(*VipUpgradeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_BuyRide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyRideReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).BuyRide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_BuyRide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).BuyRide(ctx, req.(*BuyRideReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_BuyLuckyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyLuckyIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).BuyLuckyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_BuyLuckyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).BuyLuckyId(ctx, req.(*BuyLuckyIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_MoneyRecharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoneyRechargeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).MoneyRecharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_MoneyRecharge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).MoneyRecharge(ctx, req.(*MoneyRechargeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_MoneyWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoneyWithdrawReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).MoneyWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_MoneyWithdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).MoneyWithdraw(ctx, req.(*MoneyWithdrawReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_MoneyExchangeCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoneyExchangeCoinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).MoneyExchangeCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_MoneyExchangeCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).MoneyExchangeCoin(ctx, req.(*MoneyExchangeCoinReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Trade_ServiceDesc is the grpc.ServiceDesc for Trade service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Trade_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.trade.Trade",
	HandlerType: (*TradeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendGiftInLive",
			Handler:    _Trade_SendGiftInLive_Handler,
		},
		{
			MethodName: "BuyLiveTicket",
			Handler:    _Trade_BuyLiveTicket_Handler,
		},
		{
			MethodName: "PayLiveDurationFee",
			Handler:    _Trade_PayLiveDurationFee_Handler,
		},
		{
			MethodName: "JoinAnchorFansGroup",
			Handler:    _Trade_JoinAnchorFansGroup_Handler,
		},
		{
			MethodName: "PayBulletChat",
			Handler:    _Trade_PayBulletChat_Handler,
		},
		{
			MethodName: "VipActivate",
			Handler:    _Trade_VipActivate_Handler,
		},
		{
			MethodName: "VipExtend",
			Handler:    _Trade_VipExtend_Handler,
		},
		{
			MethodName: "VipUpgrade",
			Handler:    _Trade_VipUpgrade_Handler,
		},
		{
			MethodName: "BuyRide",
			Handler:    _Trade_BuyRide_Handler,
		},
		{
			MethodName: "BuyLuckyId",
			Handler:    _Trade_BuyLuckyId_Handler,
		},
		{
			MethodName: "MoneyRecharge",
			Handler:    _Trade_MoneyRecharge_Handler,
		},
		{
			MethodName: "MoneyWithdraw",
			Handler:    _Trade_MoneyWithdraw_Handler,
		},
		{
			MethodName: "MoneyExchangeCoin",
			Handler:    _Trade_MoneyExchangeCoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.trade/trade.proto",
}
