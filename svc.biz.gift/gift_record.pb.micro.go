// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.gift/gift_record.proto

package gift

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	client "go-micro.dev/v5/client"
	server "go-micro.dev/v5/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for GiftRecord service

type GiftRecordService interface {
	// GetSendRecordList 送礼记录
	GetSendRecordList(ctx context.Context, in *GetSendRecordListReq, opts ...client.CallOption) (*GetSendRecordListResp, error)
	// GetGetRecordList 收礼记录
	GetGetRecordList(ctx context.Context, in *GetGetRecordListReq, opts ...client.CallOption) (*GetGetRecordListResp, error)
	// GetStat 礼物统计
	GetStat(ctx context.Context, in *GetStatReq, opts ...client.CallOption) (*GetStatResp, error)
}

type giftRecordService struct {
	c    client.Client
	name string
}

func NewGiftRecordService(name string, c client.Client) GiftRecordService {
	return &giftRecordService{
		c:    c,
		name: name,
	}
}

func (c *giftRecordService) GetSendRecordList(ctx context.Context, in *GetSendRecordListReq, opts ...client.CallOption) (*GetSendRecordListResp, error) {
	req := c.c.NewRequest(c.name, "GiftRecord.GetSendRecordList", in)
	out := new(GetSendRecordListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftRecordService) GetGetRecordList(ctx context.Context, in *GetGetRecordListReq, opts ...client.CallOption) (*GetGetRecordListResp, error) {
	req := c.c.NewRequest(c.name, "GiftRecord.GetGetRecordList", in)
	out := new(GetGetRecordListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftRecordService) GetStat(ctx context.Context, in *GetStatReq, opts ...client.CallOption) (*GetStatResp, error) {
	req := c.c.NewRequest(c.name, "GiftRecord.GetStat", in)
	out := new(GetStatResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GiftRecord service

type GiftRecordHandler interface {
	// GetSendRecordList 送礼记录
	GetSendRecordList(context.Context, *GetSendRecordListReq, *GetSendRecordListResp) error
	// GetGetRecordList 收礼记录
	GetGetRecordList(context.Context, *GetGetRecordListReq, *GetGetRecordListResp) error
	// GetStat 礼物统计
	GetStat(context.Context, *GetStatReq, *GetStatResp) error
}

func RegisterGiftRecordHandler(s server.Server, hdlr GiftRecordHandler, opts ...server.HandlerOption) error {
	type giftRecord interface {
		GetSendRecordList(ctx context.Context, in *GetSendRecordListReq, out *GetSendRecordListResp) error
		GetGetRecordList(ctx context.Context, in *GetGetRecordListReq, out *GetGetRecordListResp) error
		GetStat(ctx context.Context, in *GetStatReq, out *GetStatResp) error
	}
	type GiftRecord struct {
		giftRecord
	}
	h := &giftRecordHandler{hdlr}
	return s.Handle(s.NewHandler(&GiftRecord{h}, opts...))
}

type giftRecordHandler struct {
	GiftRecordHandler
}

func (h *giftRecordHandler) GetSendRecordList(ctx context.Context, in *GetSendRecordListReq, out *GetSendRecordListResp) error {
	return h.GiftRecordHandler.GetSendRecordList(ctx, in, out)
}

func (h *giftRecordHandler) GetGetRecordList(ctx context.Context, in *GetGetRecordListReq, out *GetGetRecordListResp) error {
	return h.GiftRecordHandler.GetGetRecordList(ctx, in, out)
}

func (h *giftRecordHandler) GetStat(ctx context.Context, in *GetStatReq, out *GetStatResp) error {
	return h.GiftRecordHandler.GetStat(ctx, in, out)
}
