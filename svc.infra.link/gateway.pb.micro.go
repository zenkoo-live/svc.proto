// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.infra.link/gateway.proto

package link

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
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

// Api Endpoints for LinkGateway service

func NewLinkGatewayEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LinkGateway service

type LinkGatewayService interface {
	// 网关列表
	List(ctx context.Context, in *ListGatewayRequest, opts ...client.CallOption) (*ListGatewayResponse, error)
	// 获取网关信息
	Get(ctx context.Context, in *GetGatewayRequest, opts ...client.CallOption) (*GetGatewayResponse, error)
	// 分配网关
	Select(ctx context.Context, in *SelectGatewayRequest, opts ...client.CallOption) (*SelectGatewayResponse, error)
}

type linkGatewayService struct {
	c    client.Client
	name string
}

func NewLinkGatewayService(name string, c client.Client) LinkGatewayService {
	return &linkGatewayService{
		c:    c,
		name: name,
	}
}

func (c *linkGatewayService) List(ctx context.Context, in *ListGatewayRequest, opts ...client.CallOption) (*ListGatewayResponse, error) {
	req := c.c.NewRequest(c.name, "LinkGateway.List", in)
	out := new(ListGatewayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkGatewayService) Get(ctx context.Context, in *GetGatewayRequest, opts ...client.CallOption) (*GetGatewayResponse, error) {
	req := c.c.NewRequest(c.name, "LinkGateway.Get", in)
	out := new(GetGatewayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkGatewayService) Select(ctx context.Context, in *SelectGatewayRequest, opts ...client.CallOption) (*SelectGatewayResponse, error) {
	req := c.c.NewRequest(c.name, "LinkGateway.Select", in)
	out := new(SelectGatewayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LinkGateway service

type LinkGatewayHandler interface {
	// 网关列表
	List(context.Context, *ListGatewayRequest, *ListGatewayResponse) error
	// 获取网关信息
	Get(context.Context, *GetGatewayRequest, *GetGatewayResponse) error
	// 分配网关
	Select(context.Context, *SelectGatewayRequest, *SelectGatewayResponse) error
}

func RegisterLinkGatewayHandler(s server.Server, hdlr LinkGatewayHandler, opts ...server.HandlerOption) error {
	type linkGateway interface {
		List(ctx context.Context, in *ListGatewayRequest, out *ListGatewayResponse) error
		Get(ctx context.Context, in *GetGatewayRequest, out *GetGatewayResponse) error
		Select(ctx context.Context, in *SelectGatewayRequest, out *SelectGatewayResponse) error
	}
	type LinkGateway struct {
		linkGateway
	}
	h := &linkGatewayHandler{hdlr}
	return s.Handle(s.NewHandler(&LinkGateway{h}, opts...))
}

type linkGatewayHandler struct {
	LinkGatewayHandler
}

func (h *linkGatewayHandler) List(ctx context.Context, in *ListGatewayRequest, out *ListGatewayResponse) error {
	return h.LinkGatewayHandler.List(ctx, in, out)
}

func (h *linkGatewayHandler) Get(ctx context.Context, in *GetGatewayRequest, out *GetGatewayResponse) error {
	return h.LinkGatewayHandler.Get(ctx, in, out)
}

func (h *linkGatewayHandler) Select(ctx context.Context, in *SelectGatewayRequest, out *SelectGatewayResponse) error {
	return h.LinkGatewayHandler.Select(ctx, in, out)
}
