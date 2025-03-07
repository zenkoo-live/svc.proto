// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.vip/noble_salarys.proto

package vip

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

// Api Endpoints for NobleSalary service

func NewNobleSalaryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for NobleSalary service

type NobleSalaryService interface {
	// 发放俸禄(任务系统调用/每天)
	Distribute(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*DistributeSalaryResp, error)
	// 领取俸禄
	Receive(ctx context.Context, in *ReceiveSalaryReq, opts ...client.CallOption) (*ReceiveSalaryResp, error)
	// 查询俸禄列表(按发放时间倒序)
	List(ctx context.Context, in *NobleSalaryListReq, opts ...client.CallOption) (*NobleSalaryListResp, error)
	// 查询俸禄领信息(金额、状态等)
	GetReceiveInfo(ctx context.Context, in *ReceiveSalaryReq, opts ...client.CallOption) (*NobleSalaryInfoResp, error)
}

type nobleSalaryService struct {
	c    client.Client
	name string
}

func NewNobleSalaryService(name string, c client.Client) NobleSalaryService {
	return &nobleSalaryService{
		c:    c,
		name: name,
	}
}

func (c *nobleSalaryService) Distribute(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*DistributeSalaryResp, error) {
	req := c.c.NewRequest(c.name, "NobleSalary.Distribute", in)
	out := new(DistributeSalaryResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleSalaryService) Receive(ctx context.Context, in *ReceiveSalaryReq, opts ...client.CallOption) (*ReceiveSalaryResp, error) {
	req := c.c.NewRequest(c.name, "NobleSalary.Receive", in)
	out := new(ReceiveSalaryResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleSalaryService) List(ctx context.Context, in *NobleSalaryListReq, opts ...client.CallOption) (*NobleSalaryListResp, error) {
	req := c.c.NewRequest(c.name, "NobleSalary.List", in)
	out := new(NobleSalaryListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nobleSalaryService) GetReceiveInfo(ctx context.Context, in *ReceiveSalaryReq, opts ...client.CallOption) (*NobleSalaryInfoResp, error) {
	req := c.c.NewRequest(c.name, "NobleSalary.GetReceiveInfo", in)
	out := new(NobleSalaryInfoResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NobleSalary service

type NobleSalaryHandler interface {
	// 发放俸禄(任务系统调用/每天)
	Distribute(context.Context, *EmptyRequest, *DistributeSalaryResp) error
	// 领取俸禄
	Receive(context.Context, *ReceiveSalaryReq, *ReceiveSalaryResp) error
	// 查询俸禄列表(按发放时间倒序)
	List(context.Context, *NobleSalaryListReq, *NobleSalaryListResp) error
	// 查询俸禄领信息(金额、状态等)
	GetReceiveInfo(context.Context, *ReceiveSalaryReq, *NobleSalaryInfoResp) error
}

func RegisterNobleSalaryHandler(s server.Server, hdlr NobleSalaryHandler, opts ...server.HandlerOption) error {
	type nobleSalary interface {
		Distribute(ctx context.Context, in *EmptyRequest, out *DistributeSalaryResp) error
		Receive(ctx context.Context, in *ReceiveSalaryReq, out *ReceiveSalaryResp) error
		List(ctx context.Context, in *NobleSalaryListReq, out *NobleSalaryListResp) error
		GetReceiveInfo(ctx context.Context, in *ReceiveSalaryReq, out *NobleSalaryInfoResp) error
	}
	type NobleSalary struct {
		nobleSalary
	}
	h := &nobleSalaryHandler{hdlr}
	return s.Handle(s.NewHandler(&NobleSalary{h}, opts...))
}

type nobleSalaryHandler struct {
	NobleSalaryHandler
}

func (h *nobleSalaryHandler) Distribute(ctx context.Context, in *EmptyRequest, out *DistributeSalaryResp) error {
	return h.NobleSalaryHandler.Distribute(ctx, in, out)
}

func (h *nobleSalaryHandler) Receive(ctx context.Context, in *ReceiveSalaryReq, out *ReceiveSalaryResp) error {
	return h.NobleSalaryHandler.Receive(ctx, in, out)
}

func (h *nobleSalaryHandler) List(ctx context.Context, in *NobleSalaryListReq, out *NobleSalaryListResp) error {
	return h.NobleSalaryHandler.List(ctx, in, out)
}

func (h *nobleSalaryHandler) GetReceiveInfo(ctx context.Context, in *ReceiveSalaryReq, out *NobleSalaryInfoResp) error {
	return h.NobleSalaryHandler.GetReceiveInfo(ctx, in, out)
}
