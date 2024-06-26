// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.infra.link/message.proto

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

// Api Endpoints for LinkMessage service

func NewLinkMessageEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LinkMessage service

type LinkMessageService interface {
	// 发送全局消息
	SendGlobal(ctx context.Context, in *SendGlobalRequest, opts ...client.CallOption) (*SendGlobalResponse, error)
	// 发送群组消息
	SendGroup(ctx context.Context, in *SendGroupRequest, opts ...client.CallOption) (*SendGroupResponse, error)
	// 发送账号消息（对单人)
	SendAccount(ctx context.Context, in *SendAccountRequest, opts ...client.CallOption) (*SendAccountResponse, error)
	// 发送设备消息
	SendDevice(ctx context.Context, in *SendDeviceRequest, opts ...client.CallOption) (*SendDeviceResponse, error)
	// 发送连接消息
	SendSession(ctx context.Context, in *SendSessionRequest, opts ...client.CallOption) (*SendSessionResponse, error)
}

type linkMessageService struct {
	c    client.Client
	name string
}

func NewLinkMessageService(name string, c client.Client) LinkMessageService {
	return &linkMessageService{
		c:    c,
		name: name,
	}
}

func (c *linkMessageService) SendGlobal(ctx context.Context, in *SendGlobalRequest, opts ...client.CallOption) (*SendGlobalResponse, error) {
	req := c.c.NewRequest(c.name, "LinkMessage.SendGlobal", in)
	out := new(SendGlobalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkMessageService) SendGroup(ctx context.Context, in *SendGroupRequest, opts ...client.CallOption) (*SendGroupResponse, error) {
	req := c.c.NewRequest(c.name, "LinkMessage.SendGroup", in)
	out := new(SendGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkMessageService) SendAccount(ctx context.Context, in *SendAccountRequest, opts ...client.CallOption) (*SendAccountResponse, error) {
	req := c.c.NewRequest(c.name, "LinkMessage.SendAccount", in)
	out := new(SendAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkMessageService) SendDevice(ctx context.Context, in *SendDeviceRequest, opts ...client.CallOption) (*SendDeviceResponse, error) {
	req := c.c.NewRequest(c.name, "LinkMessage.SendDevice", in)
	out := new(SendDeviceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkMessageService) SendSession(ctx context.Context, in *SendSessionRequest, opts ...client.CallOption) (*SendSessionResponse, error) {
	req := c.c.NewRequest(c.name, "LinkMessage.SendSession", in)
	out := new(SendSessionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LinkMessage service

type LinkMessageHandler interface {
	// 发送全局消息
	SendGlobal(context.Context, *SendGlobalRequest, *SendGlobalResponse) error
	// 发送群组消息
	SendGroup(context.Context, *SendGroupRequest, *SendGroupResponse) error
	// 发送账号消息（对单人)
	SendAccount(context.Context, *SendAccountRequest, *SendAccountResponse) error
	// 发送设备消息
	SendDevice(context.Context, *SendDeviceRequest, *SendDeviceResponse) error
	// 发送连接消息
	SendSession(context.Context, *SendSessionRequest, *SendSessionResponse) error
}

func RegisterLinkMessageHandler(s server.Server, hdlr LinkMessageHandler, opts ...server.HandlerOption) error {
	type linkMessage interface {
		SendGlobal(ctx context.Context, in *SendGlobalRequest, out *SendGlobalResponse) error
		SendGroup(ctx context.Context, in *SendGroupRequest, out *SendGroupResponse) error
		SendAccount(ctx context.Context, in *SendAccountRequest, out *SendAccountResponse) error
		SendDevice(ctx context.Context, in *SendDeviceRequest, out *SendDeviceResponse) error
		SendSession(ctx context.Context, in *SendSessionRequest, out *SendSessionResponse) error
	}
	type LinkMessage struct {
		linkMessage
	}
	h := &linkMessageHandler{hdlr}
	return s.Handle(s.NewHandler(&LinkMessage{h}, opts...))
}

type linkMessageHandler struct {
	LinkMessageHandler
}

func (h *linkMessageHandler) SendGlobal(ctx context.Context, in *SendGlobalRequest, out *SendGlobalResponse) error {
	return h.LinkMessageHandler.SendGlobal(ctx, in, out)
}

func (h *linkMessageHandler) SendGroup(ctx context.Context, in *SendGroupRequest, out *SendGroupResponse) error {
	return h.LinkMessageHandler.SendGroup(ctx, in, out)
}

func (h *linkMessageHandler) SendAccount(ctx context.Context, in *SendAccountRequest, out *SendAccountResponse) error {
	return h.LinkMessageHandler.SendAccount(ctx, in, out)
}

func (h *linkMessageHandler) SendDevice(ctx context.Context, in *SendDeviceRequest, out *SendDeviceResponse) error {
	return h.LinkMessageHandler.SendDevice(ctx, in, out)
}

func (h *linkMessageHandler) SendSession(ctx context.Context, in *SendSessionRequest, out *SendSessionResponse) error {
	return h.LinkMessageHandler.SendSession(ctx, in, out)
}
