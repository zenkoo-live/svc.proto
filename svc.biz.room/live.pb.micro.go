// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.room/live.proto

package room

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

// Api Endpoints for Live service

func NewLiveEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Live service

type LiveService interface {
	// GetLive 查询直播信息
	GetLive(ctx context.Context, in *GetLiveReq, opts ...client.CallOption) (*GetLiveResp, error)
	// MGetLive 批量获取直播信息
	MGetLive(ctx context.Context, in *MGetLiveReq, opts ...client.CallOption) (*MGetLiveResp, error)
	// ListLive 获取直播列表
	ListLive(ctx context.Context, in *ListLiveReq, opts ...client.CallOption) (*ListLiveResp, error)
	// StatLive 获取直播统计信息
	StatLive(ctx context.Context, in *StatLiveReq, opts ...client.CallOption) (*StatLiveResp, error)
}

type liveService struct {
	c    client.Client
	name string
}

func NewLiveService(name string, c client.Client) LiveService {
	return &liveService{
		c:    c,
		name: name,
	}
}

func (c *liveService) GetLive(ctx context.Context, in *GetLiveReq, opts ...client.CallOption) (*GetLiveResp, error) {
	req := c.c.NewRequest(c.name, "Live.GetLive", in)
	out := new(GetLiveResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveService) MGetLive(ctx context.Context, in *MGetLiveReq, opts ...client.CallOption) (*MGetLiveResp, error) {
	req := c.c.NewRequest(c.name, "Live.MGetLive", in)
	out := new(MGetLiveResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveService) ListLive(ctx context.Context, in *ListLiveReq, opts ...client.CallOption) (*ListLiveResp, error) {
	req := c.c.NewRequest(c.name, "Live.ListLive", in)
	out := new(ListLiveResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveService) StatLive(ctx context.Context, in *StatLiveReq, opts ...client.CallOption) (*StatLiveResp, error) {
	req := c.c.NewRequest(c.name, "Live.StatLive", in)
	out := new(StatLiveResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Live service

type LiveHandler interface {
	// GetLive 查询直播信息
	GetLive(context.Context, *GetLiveReq, *GetLiveResp) error
	// MGetLive 批量获取直播信息
	MGetLive(context.Context, *MGetLiveReq, *MGetLiveResp) error
	// ListLive 获取直播列表
	ListLive(context.Context, *ListLiveReq, *ListLiveResp) error
	// StatLive 获取直播统计信息
	StatLive(context.Context, *StatLiveReq, *StatLiveResp) error
}

func RegisterLiveHandler(s server.Server, hdlr LiveHandler, opts ...server.HandlerOption) error {
	type live interface {
		GetLive(ctx context.Context, in *GetLiveReq, out *GetLiveResp) error
		MGetLive(ctx context.Context, in *MGetLiveReq, out *MGetLiveResp) error
		ListLive(ctx context.Context, in *ListLiveReq, out *ListLiveResp) error
		StatLive(ctx context.Context, in *StatLiveReq, out *StatLiveResp) error
	}
	type Live struct {
		live
	}
	h := &liveHandler{hdlr}
	return s.Handle(s.NewHandler(&Live{h}, opts...))
}

type liveHandler struct {
	LiveHandler
}

func (h *liveHandler) GetLive(ctx context.Context, in *GetLiveReq, out *GetLiveResp) error {
	return h.LiveHandler.GetLive(ctx, in, out)
}

func (h *liveHandler) MGetLive(ctx context.Context, in *MGetLiveReq, out *MGetLiveResp) error {
	return h.LiveHandler.MGetLive(ctx, in, out)
}

func (h *liveHandler) ListLive(ctx context.Context, in *ListLiveReq, out *ListLiveResp) error {
	return h.LiveHandler.ListLive(ctx, in, out)
}

func (h *liveHandler) StatLive(ctx context.Context, in *StatLiveReq, out *StatLiveResp) error {
	return h.LiveHandler.StatLive(ctx, in, out)
}
