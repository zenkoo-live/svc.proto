// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.vip/noble_member.proto

package vip

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
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

// Api Endpoints for NobleMember service

func NewNobleMemberEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for NobleMember service

type NobleMemberService interface {
	// JoinNoble 加入贵族
	JoinNoble(ctx context.Context, in *JoinNobleReq, opts ...client.CallOption) (*JoinNobleResp, error)
	// GetNobleMember 获取成员贵族信息
	GetNobleMember(ctx context.Context, in *GetNobleMemberReq, opts ...client.CallOption) (*GetNobleMemberResp, error)
	// GetOnlineNobleMemberByStreamerID 获取主播贵族在线成员列表
	GetOnlineNobleMemberByStreamerID(ctx context.Context, in *GetOnlineNobleMemberByStreamerIDReq, opts ...client.CallOption) (*GetOnlineNobleMemberByStreamerIDResp, error)
}

type nobleMemberService struct {
	c    client.Client
	name string
}

func NewNobleMemberService(name string, c client.Client) NobleMemberService {
	return &nobleMemberService{
		c:    c,
		name: name,
	}
}

func (c *nobleMemberService) JoinNoble(ctx context.Context, in *JoinNobleReq, opts ...client.CallOption) (*JoinNobleResp, error) {
	req := c.c.NewRequest(c.name, "NobleMember.JoinNoble", in)
	out := new(JoinNobleResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberService) GetNobleMember(ctx context.Context, in *GetNobleMemberReq, opts ...client.CallOption) (*GetNobleMemberResp, error) {
	req := c.c.NewRequest(c.name, "NobleMember.GetNobleMember", in)
	out := new(GetNobleMemberResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleMemberService) GetOnlineNobleMemberByStreamerID(ctx context.Context, in *GetOnlineNobleMemberByStreamerIDReq, opts ...client.CallOption) (*GetOnlineNobleMemberByStreamerIDResp, error) {
	req := c.c.NewRequest(c.name, "NobleMember.GetOnlineNobleMemberByStreamerID", in)
	out := new(GetOnlineNobleMemberByStreamerIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NobleMember service

type NobleMemberHandler interface {
	// JoinNoble 加入贵族
	JoinNoble(context.Context, *JoinNobleReq, *JoinNobleResp) error
	// GetNobleMember 获取成员贵族信息
	GetNobleMember(context.Context, *GetNobleMemberReq, *GetNobleMemberResp) error
	// GetOnlineNobleMemberByStreamerID 获取主播贵族在线成员列表
	GetOnlineNobleMemberByStreamerID(context.Context, *GetOnlineNobleMemberByStreamerIDReq, *GetOnlineNobleMemberByStreamerIDResp) error
}

func RegisterNobleMemberHandler(s server.Server, hdlr NobleMemberHandler, opts ...server.HandlerOption) error {
	type nobleMember interface {
		JoinNoble(ctx context.Context, in *JoinNobleReq, out *JoinNobleResp) error
		GetNobleMember(ctx context.Context, in *GetNobleMemberReq, out *GetNobleMemberResp) error
		GetOnlineNobleMemberByStreamerID(ctx context.Context, in *GetOnlineNobleMemberByStreamerIDReq, out *GetOnlineNobleMemberByStreamerIDResp) error
	}
	type NobleMember struct {
		nobleMember
	}
	h := &nobleMemberHandler{hdlr}
	return s.Handle(s.NewHandler(&NobleMember{h}, opts...))
}

type nobleMemberHandler struct {
	NobleMemberHandler
}

func (h *nobleMemberHandler) JoinNoble(ctx context.Context, in *JoinNobleReq, out *JoinNobleResp) error {
	return h.NobleMemberHandler.JoinNoble(ctx, in, out)
}

func (h *nobleMemberHandler) GetNobleMember(ctx context.Context, in *GetNobleMemberReq, out *GetNobleMemberResp) error {
	return h.NobleMemberHandler.GetNobleMember(ctx, in, out)
}

func (h *nobleMemberHandler) GetOnlineNobleMemberByStreamerID(ctx context.Context, in *GetOnlineNobleMemberByStreamerIDReq, out *GetOnlineNobleMemberByStreamerIDResp) error {
	return h.NobleMemberHandler.GetOnlineNobleMemberByStreamerID(ctx, in, out)
}