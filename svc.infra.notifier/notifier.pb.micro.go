// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.infra.notifier/notifier.proto

package notifier

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

// Api Endpoints for Notifier service

func NewNotifierEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Notifier service

type NotifierService interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error)
}

type notifierService struct {
	c    client.Client
	name string
}

func NewNotifierService(name string, c client.Client) NotifierService {
	return &notifierService{
		c:    c,
		name: name,
	}
}

func (c *notifierService) InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error) {
	req := c.c.NewRequest(c.name, "Notifier.InitDB", in)
	out := new(InitDBResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notifier service

type NotifierHandler interface {
	InitDB(context.Context, *emptypb.Empty, *InitDBResp) error
}

func RegisterNotifierHandler(s server.Server, hdlr NotifierHandler, opts ...server.HandlerOption) error {
	type notifier interface {
		InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error
	}
	type Notifier struct {
		notifier
	}
	h := &notifierHandler{hdlr}
	return s.Handle(s.NewHandler(&Notifier{h}, opts...))
}

type notifierHandler struct {
	NotifierHandler
}

func (h *notifierHandler) InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error {
	return h.NotifierHandler.InitDB(ctx, in, out)
}