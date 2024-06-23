// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.vip/noble.proto

package vip

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
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

// Api Endpoints for Noble service

func NewNobleEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Noble service

type NobleService interface {
	// CreateNoble 创建贵族
	CreateNoble(ctx context.Context, in *CreateNobleReq, opts ...client.CallOption) (*CreateNobleResp, error)
	// GetNoble 创建贵族
	GetNoble(ctx context.Context, in *GetNobleReq, opts ...client.CallOption) (*GetNobleResp, error)
	// CreateNoble 创建贵族
	GetNobleList(ctx context.Context, in *GetNobleListReq, opts ...client.CallOption) (*GetNobleListResp, error)
	// UpdateNoble 更新贵族
	UpdateNoble(ctx context.Context, in *UpdateNobleReq, opts ...client.CallOption) (*UpdateNobleResp, error)
}

type nobleService struct {
	c    client.Client
	name string
}

func NewNobleService(name string, c client.Client) NobleService {
	return &nobleService{
		c:    c,
		name: name,
	}
}

func (c *nobleService) CreateNoble(ctx context.Context, in *CreateNobleReq, opts ...client.CallOption) (*CreateNobleResp, error) {
	req := c.c.NewRequest(c.name, "Noble.CreateNoble", in)
	out := new(CreateNobleResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleService) GetNoble(ctx context.Context, in *GetNobleReq, opts ...client.CallOption) (*GetNobleResp, error) {
	req := c.c.NewRequest(c.name, "Noble.GetNoble", in)
	out := new(GetNobleResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleService) GetNobleList(ctx context.Context, in *GetNobleListReq, opts ...client.CallOption) (*GetNobleListResp, error) {
	req := c.c.NewRequest(c.name, "Noble.GetNobleList", in)
	out := new(GetNobleListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleService) UpdateNoble(ctx context.Context, in *UpdateNobleReq, opts ...client.CallOption) (*UpdateNobleResp, error) {
	req := c.c.NewRequest(c.name, "Noble.UpdateNoble", in)
	out := new(UpdateNobleResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Noble service

type NobleHandler interface {
	// CreateNoble 创建贵族
	CreateNoble(context.Context, *CreateNobleReq, *CreateNobleResp) error
	// GetNoble 创建贵族
	GetNoble(context.Context, *GetNobleReq, *GetNobleResp) error
	// CreateNoble 创建贵族
	GetNobleList(context.Context, *GetNobleListReq, *GetNobleListResp) error
	// UpdateNoble 更新贵族
	UpdateNoble(context.Context, *UpdateNobleReq, *UpdateNobleResp) error
}

func RegisterNobleHandler(s server.Server, hdlr NobleHandler, opts ...server.HandlerOption) error {
	type noble interface {
		CreateNoble(ctx context.Context, in *CreateNobleReq, out *CreateNobleResp) error
		GetNoble(ctx context.Context, in *GetNobleReq, out *GetNobleResp) error
		GetNobleList(ctx context.Context, in *GetNobleListReq, out *GetNobleListResp) error
		UpdateNoble(ctx context.Context, in *UpdateNobleReq, out *UpdateNobleResp) error
	}
	type Noble struct {
		noble
	}
	h := &nobleHandler{hdlr}
	return s.Handle(s.NewHandler(&Noble{h}, opts...))
}

type nobleHandler struct {
	NobleHandler
}

func (h *nobleHandler) CreateNoble(ctx context.Context, in *CreateNobleReq, out *CreateNobleResp) error {
	return h.NobleHandler.CreateNoble(ctx, in, out)
}

func (h *nobleHandler) GetNoble(ctx context.Context, in *GetNobleReq, out *GetNobleResp) error {
	return h.NobleHandler.GetNoble(ctx, in, out)
}

func (h *nobleHandler) GetNobleList(ctx context.Context, in *GetNobleListReq, out *GetNobleListResp) error {
	return h.NobleHandler.GetNobleList(ctx, in, out)
}

func (h *nobleHandler) UpdateNoble(ctx context.Context, in *UpdateNobleReq, out *UpdateNobleResp) error {
	return h.NobleHandler.UpdateNoble(ctx, in, out)
}