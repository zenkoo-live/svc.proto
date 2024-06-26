// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.infra.link/instruction.proto

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

// Api Endpoints for LinkInstruction service

func NewLinkInstructionEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LinkInstruction service

type LinkInstructionService interface {
	// 移除（踢）连接
	RemoveSession(ctx context.Context, in *RemoveSessionRequest, opts ...client.CallOption) (*RemoveSessionResponse, error)
	// 移除（踢）账号
	RemoveAccount(ctx context.Context, in *RemoveAccountRequest, opts ...client.CallOption) (*RemoveAccountResponse, error)
	// 移除（踢）设备
	RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...client.CallOption) (*RemoveDeviceResponse, error)
}

type linkInstructionService struct {
	c    client.Client
	name string
}

func NewLinkInstructionService(name string, c client.Client) LinkInstructionService {
	return &linkInstructionService{
		c:    c,
		name: name,
	}
}

func (c *linkInstructionService) RemoveSession(ctx context.Context, in *RemoveSessionRequest, opts ...client.CallOption) (*RemoveSessionResponse, error) {
	req := c.c.NewRequest(c.name, "LinkInstruction.RemoveSession", in)
	out := new(RemoveSessionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkInstructionService) RemoveAccount(ctx context.Context, in *RemoveAccountRequest, opts ...client.CallOption) (*RemoveAccountResponse, error) {
	req := c.c.NewRequest(c.name, "LinkInstruction.RemoveAccount", in)
	out := new(RemoveAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkInstructionService) RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...client.CallOption) (*RemoveDeviceResponse, error) {
	req := c.c.NewRequest(c.name, "LinkInstruction.RemoveDevice", in)
	out := new(RemoveDeviceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LinkInstruction service

type LinkInstructionHandler interface {
	// 移除（踢）连接
	RemoveSession(context.Context, *RemoveSessionRequest, *RemoveSessionResponse) error
	// 移除（踢）账号
	RemoveAccount(context.Context, *RemoveAccountRequest, *RemoveAccountResponse) error
	// 移除（踢）设备
	RemoveDevice(context.Context, *RemoveDeviceRequest, *RemoveDeviceResponse) error
}

func RegisterLinkInstructionHandler(s server.Server, hdlr LinkInstructionHandler, opts ...server.HandlerOption) error {
	type linkInstruction interface {
		RemoveSession(ctx context.Context, in *RemoveSessionRequest, out *RemoveSessionResponse) error
		RemoveAccount(ctx context.Context, in *RemoveAccountRequest, out *RemoveAccountResponse) error
		RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, out *RemoveDeviceResponse) error
	}
	type LinkInstruction struct {
		linkInstruction
	}
	h := &linkInstructionHandler{hdlr}
	return s.Handle(s.NewHandler(&LinkInstruction{h}, opts...))
}

type linkInstructionHandler struct {
	LinkInstructionHandler
}

func (h *linkInstructionHandler) RemoveSession(ctx context.Context, in *RemoveSessionRequest, out *RemoveSessionResponse) error {
	return h.LinkInstructionHandler.RemoveSession(ctx, in, out)
}

func (h *linkInstructionHandler) RemoveAccount(ctx context.Context, in *RemoveAccountRequest, out *RemoveAccountResponse) error {
	return h.LinkInstructionHandler.RemoveAccount(ctx, in, out)
}

func (h *linkInstructionHandler) RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, out *RemoveDeviceResponse) error {
	return h.LinkInstructionHandler.RemoveDevice(ctx, in, out)
}
