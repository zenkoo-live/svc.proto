// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.vip/fanbase_member.proto

package vip

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for FanbaseMember service

func NewFanbaseMemberEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FanbaseMember service

type FanbaseMemberService interface {
	// JoinFanbase 加入粉丝团
	JoinFanbase(ctx context.Context, in *JoinFanbaseReq, opts ...client.CallOption) (*emptypb.Empty, error)
	// LeaveFanbase 离开粉丝团
	LeaveFanbase(ctx context.Context, in *LeaveFanbaseReq, opts ...client.CallOption) (*emptypb.Empty, error)
	// GetFanbaseMember 获取主播的某个粉丝团成员信息
	GetFanbaseMember(ctx context.Context, in *GetFanbaseMemberReq, opts ...client.CallOption) (*GetFanbaseMemberResp, error)
	// MGetFanbaseMember 批量获取主播的粉丝团成员信息
	MGetFanbaseMember(ctx context.Context, in *MGetFanbaseMemberReq, opts ...client.CallOption) (*MGetFanbaseMemberResp, error)
	// GetFanbaseMemberByStreamerID 获取主播粉丝团成员列表
	GetFanbaseMemberByStreamerID(ctx context.Context, in *GetFanbaseMemberByStreamerIDReq, opts ...client.CallOption) (*GetFanbaseMemberByStreamerIDResp, error)
	// CountFanbaseMemberByStreamerID 获取主播粉丝团成员总数
	CountFanbaseMemberByStreamerID(ctx context.Context, in *CountFanbaseMemberByStreamerIDReq, opts ...client.CallOption) (*CountFanbaseMemberByStreamerIDResp, error)
	// GetOnlineFanbaseMemberByStreamerID 获取主播粉丝团在线成员列表
	GetOnlineFanbaseMemberByStreamerID(ctx context.Context, in *GetOnlineFanbaseMemberByStreamerIDReq, opts ...client.CallOption) (*GetOnlineFanbaseMemberByStreamerIDResp, error)
	// GetFanbaseMembertByMemberID 获取用户加入的粉丝团列表
	GetFanbaseMembertByMemberID(ctx context.Context, in *GetFanbaseMembertByMemberIDReq, opts ...client.CallOption) (*GetFanbaseMembertByMemberIDResp, error)
	// CountFanbaseMembertByMemberID 获取用户加入的粉丝团数量
	CountFanbaseMembertByMemberID(ctx context.Context, in *CountFanbaseMembertByMemberIDReq, opts ...client.CallOption) (*CountFanbaseMembertByMemberIDResp, error)
	// GetFanbaseOrders 获取粉丝团订单列表
	GetFanbaseOrders(ctx context.Context, in *GetFanbaseOrdersReq, opts ...client.CallOption) (*GetFanbaseOrdersResp, error)
	// GetFanbaseOrderStat 获取粉丝团订单统计
	GetFanbaseOrderStat(ctx context.Context, in *GetFanbaseOrderStatReq, opts ...client.CallOption) (*GetFanbaseOrderStatResp, error)
}

type fanbaseMemberService struct {
	c    client.Client
	name string
}

func NewFanbaseMemberService(name string, c client.Client) FanbaseMemberService {
	return &fanbaseMemberService{
		c:    c,
		name: name,
	}
}

func (c *fanbaseMemberService) JoinFanbase(ctx context.Context, in *JoinFanbaseReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.JoinFanbase", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) LeaveFanbase(ctx context.Context, in *LeaveFanbaseReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.LeaveFanbase", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) GetFanbaseMember(ctx context.Context, in *GetFanbaseMemberReq, opts ...client.CallOption) (*GetFanbaseMemberResp, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.GetFanbaseMember", in)
	out := new(GetFanbaseMemberResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) MGetFanbaseMember(ctx context.Context, in *MGetFanbaseMemberReq, opts ...client.CallOption) (*MGetFanbaseMemberResp, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.MGetFanbaseMember", in)
	out := new(MGetFanbaseMemberResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) GetFanbaseMemberByStreamerID(ctx context.Context, in *GetFanbaseMemberByStreamerIDReq, opts ...client.CallOption) (*GetFanbaseMemberByStreamerIDResp, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.GetFanbaseMemberByStreamerID", in)
	out := new(GetFanbaseMemberByStreamerIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) CountFanbaseMemberByStreamerID(ctx context.Context, in *CountFanbaseMemberByStreamerIDReq, opts ...client.CallOption) (*CountFanbaseMemberByStreamerIDResp, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.CountFanbaseMemberByStreamerID", in)
	out := new(CountFanbaseMemberByStreamerIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) GetOnlineFanbaseMemberByStreamerID(ctx context.Context, in *GetOnlineFanbaseMemberByStreamerIDReq, opts ...client.CallOption) (*GetOnlineFanbaseMemberByStreamerIDResp, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.GetOnlineFanbaseMemberByStreamerID", in)
	out := new(GetOnlineFanbaseMemberByStreamerIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) GetFanbaseMembertByMemberID(ctx context.Context, in *GetFanbaseMembertByMemberIDReq, opts ...client.CallOption) (*GetFanbaseMembertByMemberIDResp, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.GetFanbaseMembertByMemberID", in)
	out := new(GetFanbaseMembertByMemberIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) CountFanbaseMembertByMemberID(ctx context.Context, in *CountFanbaseMembertByMemberIDReq, opts ...client.CallOption) (*CountFanbaseMembertByMemberIDResp, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.CountFanbaseMembertByMemberID", in)
	out := new(CountFanbaseMembertByMemberIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) GetFanbaseOrders(ctx context.Context, in *GetFanbaseOrdersReq, opts ...client.CallOption) (*GetFanbaseOrdersResp, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.GetFanbaseOrders", in)
	out := new(GetFanbaseOrdersResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseMemberService) GetFanbaseOrderStat(ctx context.Context, in *GetFanbaseOrderStatReq, opts ...client.CallOption) (*GetFanbaseOrderStatResp, error) {
	req := c.c.NewRequest(c.name, "FanbaseMember.GetFanbaseOrderStat", in)
	out := new(GetFanbaseOrderStatResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FanbaseMember service

type FanbaseMemberHandler interface {
	// JoinFanbase 加入粉丝团
	JoinFanbase(context.Context, *JoinFanbaseReq, *emptypb.Empty) error
	// LeaveFanbase 离开粉丝团
	LeaveFanbase(context.Context, *LeaveFanbaseReq, *emptypb.Empty) error
	// GetFanbaseMember 获取主播的某个粉丝团成员信息
	GetFanbaseMember(context.Context, *GetFanbaseMemberReq, *GetFanbaseMemberResp) error
	// MGetFanbaseMember 批量获取主播的粉丝团成员信息
	MGetFanbaseMember(context.Context, *MGetFanbaseMemberReq, *MGetFanbaseMemberResp) error
	// GetFanbaseMemberByStreamerID 获取主播粉丝团成员列表
	GetFanbaseMemberByStreamerID(context.Context, *GetFanbaseMemberByStreamerIDReq, *GetFanbaseMemberByStreamerIDResp) error
	// CountFanbaseMemberByStreamerID 获取主播粉丝团成员总数
	CountFanbaseMemberByStreamerID(context.Context, *CountFanbaseMemberByStreamerIDReq, *CountFanbaseMemberByStreamerIDResp) error
	// GetOnlineFanbaseMemberByStreamerID 获取主播粉丝团在线成员列表
	GetOnlineFanbaseMemberByStreamerID(context.Context, *GetOnlineFanbaseMemberByStreamerIDReq, *GetOnlineFanbaseMemberByStreamerIDResp) error
	// GetFanbaseMembertByMemberID 获取用户加入的粉丝团列表
	GetFanbaseMembertByMemberID(context.Context, *GetFanbaseMembertByMemberIDReq, *GetFanbaseMembertByMemberIDResp) error
	// CountFanbaseMembertByMemberID 获取用户加入的粉丝团数量
	CountFanbaseMembertByMemberID(context.Context, *CountFanbaseMembertByMemberIDReq, *CountFanbaseMembertByMemberIDResp) error
	// GetFanbaseOrders 获取粉丝团订单列表
	GetFanbaseOrders(context.Context, *GetFanbaseOrdersReq, *GetFanbaseOrdersResp) error
	// GetFanbaseOrderStat 获取粉丝团订单统计
	GetFanbaseOrderStat(context.Context, *GetFanbaseOrderStatReq, *GetFanbaseOrderStatResp) error
}

func RegisterFanbaseMemberHandler(s server.Server, hdlr FanbaseMemberHandler, opts ...server.HandlerOption) error {
	type fanbaseMember interface {
		JoinFanbase(ctx context.Context, in *JoinFanbaseReq, out *emptypb.Empty) error
		LeaveFanbase(ctx context.Context, in *LeaveFanbaseReq, out *emptypb.Empty) error
		GetFanbaseMember(ctx context.Context, in *GetFanbaseMemberReq, out *GetFanbaseMemberResp) error
		MGetFanbaseMember(ctx context.Context, in *MGetFanbaseMemberReq, out *MGetFanbaseMemberResp) error
		GetFanbaseMemberByStreamerID(ctx context.Context, in *GetFanbaseMemberByStreamerIDReq, out *GetFanbaseMemberByStreamerIDResp) error
		CountFanbaseMemberByStreamerID(ctx context.Context, in *CountFanbaseMemberByStreamerIDReq, out *CountFanbaseMemberByStreamerIDResp) error
		GetOnlineFanbaseMemberByStreamerID(ctx context.Context, in *GetOnlineFanbaseMemberByStreamerIDReq, out *GetOnlineFanbaseMemberByStreamerIDResp) error
		GetFanbaseMembertByMemberID(ctx context.Context, in *GetFanbaseMembertByMemberIDReq, out *GetFanbaseMembertByMemberIDResp) error
		CountFanbaseMembertByMemberID(ctx context.Context, in *CountFanbaseMembertByMemberIDReq, out *CountFanbaseMembertByMemberIDResp) error
		GetFanbaseOrders(ctx context.Context, in *GetFanbaseOrdersReq, out *GetFanbaseOrdersResp) error
		GetFanbaseOrderStat(ctx context.Context, in *GetFanbaseOrderStatReq, out *GetFanbaseOrderStatResp) error
	}
	type FanbaseMember struct {
		fanbaseMember
	}
	h := &fanbaseMemberHandler{hdlr}
	return s.Handle(s.NewHandler(&FanbaseMember{h}, opts...))
}

type fanbaseMemberHandler struct {
	FanbaseMemberHandler
}

func (h *fanbaseMemberHandler) JoinFanbase(ctx context.Context, in *JoinFanbaseReq, out *emptypb.Empty) error {
	return h.FanbaseMemberHandler.JoinFanbase(ctx, in, out)
}

func (h *fanbaseMemberHandler) LeaveFanbase(ctx context.Context, in *LeaveFanbaseReq, out *emptypb.Empty) error {
	return h.FanbaseMemberHandler.LeaveFanbase(ctx, in, out)
}

func (h *fanbaseMemberHandler) GetFanbaseMember(ctx context.Context, in *GetFanbaseMemberReq, out *GetFanbaseMemberResp) error {
	return h.FanbaseMemberHandler.GetFanbaseMember(ctx, in, out)
}

func (h *fanbaseMemberHandler) MGetFanbaseMember(ctx context.Context, in *MGetFanbaseMemberReq, out *MGetFanbaseMemberResp) error {
	return h.FanbaseMemberHandler.MGetFanbaseMember(ctx, in, out)
}

func (h *fanbaseMemberHandler) GetFanbaseMemberByStreamerID(ctx context.Context, in *GetFanbaseMemberByStreamerIDReq, out *GetFanbaseMemberByStreamerIDResp) error {
	return h.FanbaseMemberHandler.GetFanbaseMemberByStreamerID(ctx, in, out)
}

func (h *fanbaseMemberHandler) CountFanbaseMemberByStreamerID(ctx context.Context, in *CountFanbaseMemberByStreamerIDReq, out *CountFanbaseMemberByStreamerIDResp) error {
	return h.FanbaseMemberHandler.CountFanbaseMemberByStreamerID(ctx, in, out)
}

func (h *fanbaseMemberHandler) GetOnlineFanbaseMemberByStreamerID(ctx context.Context, in *GetOnlineFanbaseMemberByStreamerIDReq, out *GetOnlineFanbaseMemberByStreamerIDResp) error {
	return h.FanbaseMemberHandler.GetOnlineFanbaseMemberByStreamerID(ctx, in, out)
}

func (h *fanbaseMemberHandler) GetFanbaseMembertByMemberID(ctx context.Context, in *GetFanbaseMembertByMemberIDReq, out *GetFanbaseMembertByMemberIDResp) error {
	return h.FanbaseMemberHandler.GetFanbaseMembertByMemberID(ctx, in, out)
}

func (h *fanbaseMemberHandler) CountFanbaseMembertByMemberID(ctx context.Context, in *CountFanbaseMembertByMemberIDReq, out *CountFanbaseMembertByMemberIDResp) error {
	return h.FanbaseMemberHandler.CountFanbaseMembertByMemberID(ctx, in, out)
}

func (h *fanbaseMemberHandler) GetFanbaseOrders(ctx context.Context, in *GetFanbaseOrdersReq, out *GetFanbaseOrdersResp) error {
	return h.FanbaseMemberHandler.GetFanbaseOrders(ctx, in, out)
}

func (h *fanbaseMemberHandler) GetFanbaseOrderStat(ctx context.Context, in *GetFanbaseOrderStatReq, out *GetFanbaseOrderStatResp) error {
	return h.FanbaseMemberHandler.GetFanbaseOrderStat(ctx, in, out)
}
