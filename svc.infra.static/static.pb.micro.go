// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.infra.static/static.proto

package static

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Api Endpoints for Static service

func NewStaticEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Static service

type StaticService interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error)
}

type staticService struct {
	c    client.Client
	name string
}

func NewStaticService(name string, c client.Client) StaticService {
	return &staticService{
		c:    c,
		name: name,
	}
}

func (c *staticService) InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error) {
	req := c.c.NewRequest(c.name, "Static.InitDB", in)
	out := new(InitDBResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Static service

type StaticHandler interface {
	InitDB(context.Context, *emptypb.Empty, *InitDBResp) error
}

func RegisterStaticHandler(s server.Server, hdlr StaticHandler, opts ...server.HandlerOption) error {
	type static interface {
		InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error
	}
	type Static struct {
		static
	}
	h := &staticHandler{hdlr}
	return s.Handle(s.NewHandler(&Static{h}, opts...))
}

type staticHandler struct {
	StaticHandler
}

func (h *staticHandler) InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error {
	return h.StaticHandler.InitDB(ctx, in, out)
}