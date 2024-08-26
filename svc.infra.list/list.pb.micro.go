// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.infra.list/list.proto

package list

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

// Api Endpoints for List service

func NewListEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for List service

type ListService interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error)
}

type listService struct {
	c    client.Client
	name string
}

func NewListService(name string, c client.Client) ListService {
	return &listService{
		c:    c,
		name: name,
	}
}

func (c *listService) InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error) {
	req := c.c.NewRequest(c.name, "List.InitDB", in)
	out := new(InitDBResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for List service

type ListHandler interface {
	InitDB(context.Context, *emptypb.Empty, *InitDBResp) error
}

func RegisterListHandler(s server.Server, hdlr ListHandler, opts ...server.HandlerOption) error {
	type list interface {
		InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error
	}
	type List struct {
		list
	}
	h := &listHandler{hdlr}
	return s.Handle(s.NewHandler(&List{h}, opts...))
}

type listHandler struct {
	ListHandler
}

func (h *listHandler) InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error {
	return h.ListHandler.InitDB(ctx, in, out)
}
