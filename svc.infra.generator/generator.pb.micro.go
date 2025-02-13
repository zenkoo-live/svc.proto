// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.infra.generator/generator.proto

package generator

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

// Api Endpoints for Generator service

func NewGeneratorEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Generator service

type GeneratorService interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error)
	AddID(ctx context.Context, in *AddIDReq, opts ...client.CallOption) (*AddIDResp, error)
	OrdianID(ctx context.Context, in *OrdianIDReq, opts ...client.CallOption) (*OrdianIDResp, error)
	NextID(ctx context.Context, in *NextIDReq, opts ...client.CallOption) (*NextIDResp, error)
	IsPrettyID(ctx context.Context, in *IsPrettyIDReq, opts ...client.CallOption) (*IsPrettyIDResp, error)
}

type generatorService struct {
	c    client.Client
	name string
}

func NewGeneratorService(name string, c client.Client) GeneratorService {
	return &generatorService{
		c:    c,
		name: name,
	}
}

func (c *generatorService) InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error) {
	req := c.c.NewRequest(c.name, "Generator.InitDB", in)
	out := new(InitDBResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorService) AddID(ctx context.Context, in *AddIDReq, opts ...client.CallOption) (*AddIDResp, error) {
	req := c.c.NewRequest(c.name, "Generator.AddID", in)
	out := new(AddIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorService) OrdianID(ctx context.Context, in *OrdianIDReq, opts ...client.CallOption) (*OrdianIDResp, error) {
	req := c.c.NewRequest(c.name, "Generator.OrdianID", in)
	out := new(OrdianIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorService) NextID(ctx context.Context, in *NextIDReq, opts ...client.CallOption) (*NextIDResp, error) {
	req := c.c.NewRequest(c.name, "Generator.NextID", in)
	out := new(NextIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorService) IsPrettyID(ctx context.Context, in *IsPrettyIDReq, opts ...client.CallOption) (*IsPrettyIDResp, error) {
	req := c.c.NewRequest(c.name, "Generator.IsPrettyID", in)
	out := new(IsPrettyIDResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Generator service

type GeneratorHandler interface {
	InitDB(context.Context, *emptypb.Empty, *InitDBResp) error
	AddID(context.Context, *AddIDReq, *AddIDResp) error
	OrdianID(context.Context, *OrdianIDReq, *OrdianIDResp) error
	NextID(context.Context, *NextIDReq, *NextIDResp) error
	IsPrettyID(context.Context, *IsPrettyIDReq, *IsPrettyIDResp) error
}

func RegisterGeneratorHandler(s server.Server, hdlr GeneratorHandler, opts ...server.HandlerOption) error {
	type generator interface {
		InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error
		AddID(ctx context.Context, in *AddIDReq, out *AddIDResp) error
		OrdianID(ctx context.Context, in *OrdianIDReq, out *OrdianIDResp) error
		NextID(ctx context.Context, in *NextIDReq, out *NextIDResp) error
		IsPrettyID(ctx context.Context, in *IsPrettyIDReq, out *IsPrettyIDResp) error
	}
	type Generator struct {
		generator
	}
	h := &generatorHandler{hdlr}
	return s.Handle(s.NewHandler(&Generator{h}, opts...))
}

type generatorHandler struct {
	GeneratorHandler
}

func (h *generatorHandler) InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error {
	return h.GeneratorHandler.InitDB(ctx, in, out)
}

func (h *generatorHandler) AddID(ctx context.Context, in *AddIDReq, out *AddIDResp) error {
	return h.GeneratorHandler.AddID(ctx, in, out)
}

func (h *generatorHandler) OrdianID(ctx context.Context, in *OrdianIDReq, out *OrdianIDResp) error {
	return h.GeneratorHandler.OrdianID(ctx, in, out)
}

func (h *generatorHandler) NextID(ctx context.Context, in *NextIDReq, out *NextIDResp) error {
	return h.GeneratorHandler.NextID(ctx, in, out)
}

func (h *generatorHandler) IsPrettyID(ctx context.Context, in *IsPrettyIDReq, out *IsPrettyIDResp) error {
	return h.GeneratorHandler.IsPrettyID(ctx, in, out)
}
