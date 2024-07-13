// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
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
	Trade_SendGiftInRoom_FullMethodName        = "/svc.biz.trade.Trade/SendGiftInRoom"
	Trade_BuyRoomTicket_FullMethodName         = "/svc.biz.trade.Trade/BuyRoomTicket"
	Trade_PayRoomDurationFee_FullMethodName    = "/svc.biz.trade.Trade/PayRoomDurationFee"
	Trade_JoinStreamerFansGroup_FullMethodName = "/svc.biz.trade.Trade/JoinStreamerFansGroup"
	Trade_PayBulletChat_FullMethodName         = "/svc.biz.trade.Trade/PayBulletChat"
	Trade_VipActivate_FullMethodName           = "/svc.biz.trade.Trade/VipActivate"
	Trade_VipExtend_FullMethodName             = "/svc.biz.trade.Trade/VipExtend"
	Trade_VipUpgrade_FullMethodName            = "/svc.biz.trade.Trade/VipUpgrade"
	Trade_BuyRide_FullMethodName               = "/svc.biz.trade.Trade/BuyRide"
	Trade_BuyLuckyId_FullMethodName            = "/svc.biz.trade.Trade/BuyLuckyId"
	Trade_MoneyRecharge_FullMethodName         = "/svc.biz.trade.Trade/MoneyRecharge"
	Trade_MoneyWithdraw_FullMethodName         = "/svc.biz.trade.Trade/MoneyWithdraw"
	Trade_MoneyExchangeCoin_FullMethodName     = "/svc.biz.trade.Trade/MoneyExchangeCoin"
)

// TradeClient is the client API for Trade service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradeClient interface {
	SendGiftInRoom(ctx context.Context, in *SendGiftInRoomReq, opts ...grpc.CallOption) (*SendGiftInRoomResp, error)
	BuyRoomTicket(ctx context.Context, in *BuyRoomTicketReq, opts ...grpc.CallOption) (*BuyRoomTicketResp, error)
	PayRoomDurationFee(ctx context.Context, in *PayRoomDurationFeeReq, opts ...grpc.CallOption) (*PayRoomDurationFeeResp, error)
	JoinStreamerFansGroup(ctx context.Context, in *JoinStreamerFansGroupReq, opts ...grpc.CallOption) (*JoinStreamerFansGroupResp, error)
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

func (c *tradeClient) SendGiftInRoom(ctx context.Context, in *SendGiftInRoomReq, opts ...grpc.CallOption) (*SendGiftInRoomResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendGiftInRoomResp)
	err := c.cc.Invoke(ctx, Trade_SendGiftInRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) BuyRoomTicket(ctx context.Context, in *BuyRoomTicketReq, opts ...grpc.CallOption) (*BuyRoomTicketResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BuyRoomTicketResp)
	err := c.cc.Invoke(ctx, Trade_BuyRoomTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) PayRoomDurationFee(ctx context.Context, in *PayRoomDurationFeeReq, opts ...grpc.CallOption) (*PayRoomDurationFeeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayRoomDurationFeeResp)
	err := c.cc.Invoke(ctx, Trade_PayRoomDurationFee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeClient) JoinStreamerFansGroup(ctx context.Context, in *JoinStreamerFansGroupReq, opts ...grpc.CallOption) (*JoinStreamerFansGroupResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinStreamerFansGroupResp)
	err := c.cc.Invoke(ctx, Trade_JoinStreamerFansGroup_FullMethodName, in, out, cOpts...)
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
	SendGiftInRoom(context.Context, *SendGiftInRoomReq) (*SendGiftInRoomResp, error)
	BuyRoomTicket(context.Context, *BuyRoomTicketReq) (*BuyRoomTicketResp, error)
	PayRoomDurationFee(context.Context, *PayRoomDurationFeeReq) (*PayRoomDurationFeeResp, error)
	JoinStreamerFansGroup(context.Context, *JoinStreamerFansGroupReq) (*JoinStreamerFansGroupResp, error)
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

func (UnimplementedTradeServer) SendGiftInRoom(context.Context, *SendGiftInRoomReq) (*SendGiftInRoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGiftInRoom not implemented")
}
func (UnimplementedTradeServer) BuyRoomTicket(context.Context, *BuyRoomTicketReq) (*BuyRoomTicketResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyRoomTicket not implemented")
}
func (UnimplementedTradeServer) PayRoomDurationFee(context.Context, *PayRoomDurationFeeReq) (*PayRoomDurationFeeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayRoomDurationFee not implemented")
}
func (UnimplementedTradeServer) JoinStreamerFansGroup(context.Context, *JoinStreamerFansGroupReq) (*JoinStreamerFansGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinStreamerFansGroup not implemented")
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

func _Trade_SendGiftInRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGiftInRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).SendGiftInRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_SendGiftInRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).SendGiftInRoom(ctx, req.(*SendGiftInRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_BuyRoomTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyRoomTicketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).BuyRoomTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_BuyRoomTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).BuyRoomTicket(ctx, req.(*BuyRoomTicketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_PayRoomDurationFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRoomDurationFeeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).PayRoomDurationFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_PayRoomDurationFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).PayRoomDurationFee(ctx, req.(*PayRoomDurationFeeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trade_JoinStreamerFansGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinStreamerFansGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeServer).JoinStreamerFansGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trade_JoinStreamerFansGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeServer).JoinStreamerFansGroup(ctx, req.(*JoinStreamerFansGroupReq))
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
			MethodName: "SendGiftInRoom",
			Handler:    _Trade_SendGiftInRoom_Handler,
		},
		{
			MethodName: "BuyRoomTicket",
			Handler:    _Trade_BuyRoomTicket_Handler,
		},
		{
			MethodName: "PayRoomDurationFee",
			Handler:    _Trade_PayRoomDurationFee_Handler,
		},
		{
			MethodName: "JoinStreamerFansGroup",
			Handler:    _Trade_JoinStreamerFansGroup_Handler,
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
