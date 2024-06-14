// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.vip/fanbase.proto

package vip

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Api Endpoints for Fanbase service

func NewFanbaseEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Fanbase service

type FanbaseService interface {
	// CreateFanbase 创建粉丝团
	CreateFanbase(ctx context.Context, in *CreateFanbaseReq, opts ...client.CallOption) (*CreateFanbaseResp, error)
	// GetFanbaseByStreamerID 获取粉丝团
	GetFanbaseByStreamerID(ctx context.Context, in *GetFanbaseByStreamerIDResp, opts ...client.CallOption) (*GetFanbaseResp, error)
	// GetFanbaseByName 通过名称获取粉丝团
	GetFanbaseByName(ctx context.Context, in *GetFanbaseByNameReq, opts ...client.CallOption) (*GetFanbaseResp, error)
	// UpdateFanbaseByStreamerID 更新粉丝团
	UpdateFanbaseByStreamerID(ctx context.Context, in *UpdateFanbaseByStreamerIDReq, opts ...client.CallOption) (*emptypb.Empty, error)
}

type fanbaseService struct {
	c    client.Client
	name string
}

func NewFanbaseService(name string, c client.Client) FanbaseService {
	return &fanbaseService{
		c:    c,
		name: name,
	}
}

func (c *fanbaseService) CreateFanbase(ctx context.Context, in *CreateFanbaseReq, opts ...client.CallOption) (*CreateFanbaseResp, error) {
	req := c.c.NewRequest(c.name, "Fanbase.CreateFanbase", in)
	out := new(CreateFanbaseResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseService) GetFanbaseByStreamerID(ctx context.Context, in *GetFanbaseByStreamerIDResp, opts ...client.CallOption) (*GetFanbaseResp, error) {
	req := c.c.NewRequest(c.name, "Fanbase.GetFanbaseByStreamerID", in)
	out := new(GetFanbaseResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseService) GetFanbaseByName(ctx context.Context, in *GetFanbaseByNameReq, opts ...client.CallOption) (*GetFanbaseResp, error) {
	req := c.c.NewRequest(c.name, "Fanbase.GetFanbaseByName", in)
	out := new(GetFanbaseResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanbaseService) UpdateFanbaseByStreamerID(ctx context.Context, in *UpdateFanbaseByStreamerIDReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Fanbase.UpdateFanbaseByStreamerID", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fanbase service

type FanbaseHandler interface {
	// CreateFanbase 创建粉丝团
	CreateFanbase(context.Context, *CreateFanbaseReq, *CreateFanbaseResp) error
	// GetFanbaseByStreamerID 获取粉丝团
	GetFanbaseByStreamerID(context.Context, *GetFanbaseByStreamerIDResp, *GetFanbaseResp) error
	// GetFanbaseByName 通过名称获取粉丝团
	GetFanbaseByName(context.Context, *GetFanbaseByNameReq, *GetFanbaseResp) error
	// UpdateFanbaseByStreamerID 更新粉丝团
	UpdateFanbaseByStreamerID(context.Context, *UpdateFanbaseByStreamerIDReq, *emptypb.Empty) error
}

func RegisterFanbaseHandler(s server.Server, hdlr FanbaseHandler, opts ...server.HandlerOption) error {
	type fanbase interface {
		CreateFanbase(ctx context.Context, in *CreateFanbaseReq, out *CreateFanbaseResp) error
		GetFanbaseByStreamerID(ctx context.Context, in *GetFanbaseByStreamerIDResp, out *GetFanbaseResp) error
		GetFanbaseByName(ctx context.Context, in *GetFanbaseByNameReq, out *GetFanbaseResp) error
		UpdateFanbaseByStreamerID(ctx context.Context, in *UpdateFanbaseByStreamerIDReq, out *emptypb.Empty) error
	}
	type Fanbase struct {
		fanbase
	}
	h := &fanbaseHandler{hdlr}
	return s.Handle(s.NewHandler(&Fanbase{h}, opts...))
}

type fanbaseHandler struct {
	FanbaseHandler
}

func (h *fanbaseHandler) CreateFanbase(ctx context.Context, in *CreateFanbaseReq, out *CreateFanbaseResp) error {
	return h.FanbaseHandler.CreateFanbase(ctx, in, out)
}

func (h *fanbaseHandler) GetFanbaseByStreamerID(ctx context.Context, in *GetFanbaseByStreamerIDResp, out *GetFanbaseResp) error {
	return h.FanbaseHandler.GetFanbaseByStreamerID(ctx, in, out)
}

func (h *fanbaseHandler) GetFanbaseByName(ctx context.Context, in *GetFanbaseByNameReq, out *GetFanbaseResp) error {
	return h.FanbaseHandler.GetFanbaseByName(ctx, in, out)
}

func (h *fanbaseHandler) UpdateFanbaseByStreamerID(ctx context.Context, in *UpdateFanbaseByStreamerIDReq, out *emptypb.Empty) error {
	return h.FanbaseHandler.UpdateFanbaseByStreamerID(ctx, in, out)
}
