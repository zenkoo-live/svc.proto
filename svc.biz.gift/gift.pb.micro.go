// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.gift/gift.proto

package gift

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Api Endpoints for Gift service

func NewGiftEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Gift service

type GiftService interface {
	Add(ctx context.Context, in *GiftAddReq, opts ...client.CallOption) (*GiftAddResp, error)
	Get(ctx context.Context, in *GiftGetReq, opts ...client.CallOption) (*GiftGetResp, error)
	Update(ctx context.Context, in *GiftUpdateReq, opts ...client.CallOption) (*GiftUpdateResp, error)
	ListAdmin(ctx context.Context, in *GiftListAdminReq, opts ...client.CallOption) (*GiftListAdminResp, error)
	List(ctx context.Context, in *GiftListReq, opts ...client.CallOption) (*GiftListResp, error)
	Send(ctx context.Context, in *GiftSendReq, opts ...client.CallOption) (*GiftSendResp, error)
	SendRecord(ctx context.Context, in *GiftSendRecordReq, opts ...client.CallOption) (*GiftSendRecordResp, error)
	GetRecord(ctx context.Context, in *GiftGetRecordReq, opts ...client.CallOption) (*GiftGetRecordResp, error)
	LiveStat(ctx context.Context, in *LiveStatReq, opts ...client.CallOption) (*LiveStatResp, error)
}

type giftService struct {
	c    client.Client
	name string
}

func NewGiftService(name string, c client.Client) GiftService {
	return &giftService{
		c:    c,
		name: name,
	}
}

func (c *giftService) Add(ctx context.Context, in *GiftAddReq, opts ...client.CallOption) (*GiftAddResp, error) {
	req := c.c.NewRequest(c.name, "Gift.Add", in)
	out := new(GiftAddResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftService) Get(ctx context.Context, in *GiftGetReq, opts ...client.CallOption) (*GiftGetResp, error) {
	req := c.c.NewRequest(c.name, "Gift.Get", in)
	out := new(GiftGetResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftService) Update(ctx context.Context, in *GiftUpdateReq, opts ...client.CallOption) (*GiftUpdateResp, error) {
	req := c.c.NewRequest(c.name, "Gift.Update", in)
	out := new(GiftUpdateResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftService) ListAdmin(ctx context.Context, in *GiftListAdminReq, opts ...client.CallOption) (*GiftListAdminResp, error) {
	req := c.c.NewRequest(c.name, "Gift.ListAdmin", in)
	out := new(GiftListAdminResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftService) List(ctx context.Context, in *GiftListReq, opts ...client.CallOption) (*GiftListResp, error) {
	req := c.c.NewRequest(c.name, "Gift.List", in)
	out := new(GiftListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftService) Send(ctx context.Context, in *GiftSendReq, opts ...client.CallOption) (*GiftSendResp, error) {
	req := c.c.NewRequest(c.name, "Gift.Send", in)
	out := new(GiftSendResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftService) SendRecord(ctx context.Context, in *GiftSendRecordReq, opts ...client.CallOption) (*GiftSendRecordResp, error) {
	req := c.c.NewRequest(c.name, "Gift.SendRecord", in)
	out := new(GiftSendRecordResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftService) GetRecord(ctx context.Context, in *GiftGetRecordReq, opts ...client.CallOption) (*GiftGetRecordResp, error) {
	req := c.c.NewRequest(c.name, "Gift.GetRecord", in)
	out := new(GiftGetRecordResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftService) LiveStat(ctx context.Context, in *LiveStatReq, opts ...client.CallOption) (*LiveStatResp, error) {
	req := c.c.NewRequest(c.name, "Gift.LiveStat", in)
	out := new(LiveStatResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gift service

type GiftHandler interface {
	Add(context.Context, *GiftAddReq, *GiftAddResp) error
	Get(context.Context, *GiftGetReq, *GiftGetResp) error
	Update(context.Context, *GiftUpdateReq, *GiftUpdateResp) error
	ListAdmin(context.Context, *GiftListAdminReq, *GiftListAdminResp) error
	List(context.Context, *GiftListReq, *GiftListResp) error
	Send(context.Context, *GiftSendReq, *GiftSendResp) error
	SendRecord(context.Context, *GiftSendRecordReq, *GiftSendRecordResp) error
	GetRecord(context.Context, *GiftGetRecordReq, *GiftGetRecordResp) error
	LiveStat(context.Context, *LiveStatReq, *LiveStatResp) error
}

func RegisterGiftHandler(s server.Server, hdlr GiftHandler, opts ...server.HandlerOption) error {
	type gift interface {
		Add(ctx context.Context, in *GiftAddReq, out *GiftAddResp) error
		Get(ctx context.Context, in *GiftGetReq, out *GiftGetResp) error
		Update(ctx context.Context, in *GiftUpdateReq, out *GiftUpdateResp) error
		ListAdmin(ctx context.Context, in *GiftListAdminReq, out *GiftListAdminResp) error
		List(ctx context.Context, in *GiftListReq, out *GiftListResp) error
		Send(ctx context.Context, in *GiftSendReq, out *GiftSendResp) error
		SendRecord(ctx context.Context, in *GiftSendRecordReq, out *GiftSendRecordResp) error
		GetRecord(ctx context.Context, in *GiftGetRecordReq, out *GiftGetRecordResp) error
		LiveStat(ctx context.Context, in *LiveStatReq, out *LiveStatResp) error
	}
	type Gift struct {
		gift
	}
	h := &giftHandler{hdlr}
	return s.Handle(s.NewHandler(&Gift{h}, opts...))
}

type giftHandler struct {
	GiftHandler
}

func (h *giftHandler) Add(ctx context.Context, in *GiftAddReq, out *GiftAddResp) error {
	return h.GiftHandler.Add(ctx, in, out)
}

func (h *giftHandler) Get(ctx context.Context, in *GiftGetReq, out *GiftGetResp) error {
	return h.GiftHandler.Get(ctx, in, out)
}

func (h *giftHandler) Update(ctx context.Context, in *GiftUpdateReq, out *GiftUpdateResp) error {
	return h.GiftHandler.Update(ctx, in, out)
}

func (h *giftHandler) ListAdmin(ctx context.Context, in *GiftListAdminReq, out *GiftListAdminResp) error {
	return h.GiftHandler.ListAdmin(ctx, in, out)
}

func (h *giftHandler) List(ctx context.Context, in *GiftListReq, out *GiftListResp) error {
	return h.GiftHandler.List(ctx, in, out)
}

func (h *giftHandler) Send(ctx context.Context, in *GiftSendReq, out *GiftSendResp) error {
	return h.GiftHandler.Send(ctx, in, out)
}

func (h *giftHandler) SendRecord(ctx context.Context, in *GiftSendRecordReq, out *GiftSendRecordResp) error {
	return h.GiftHandler.SendRecord(ctx, in, out)
}

func (h *giftHandler) GetRecord(ctx context.Context, in *GiftGetRecordReq, out *GiftGetRecordResp) error {
	return h.GiftHandler.GetRecord(ctx, in, out)
}

func (h *giftHandler) LiveStat(ctx context.Context, in *LiveStatReq, out *LiveStatResp) error {
	return h.GiftHandler.LiveStat(ctx, in, out)
}
