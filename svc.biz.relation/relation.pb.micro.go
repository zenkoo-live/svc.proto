// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.relation/relation.proto

package relation

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Api Endpoints for Relation service

func NewRelationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Relation service

type RelationService interface {
	// RelationAdd 新建关系
	RelationAdd(ctx context.Context, in *RelationAddReq, opts ...client.CallOption) (*emptypb.Empty, error)
	// RelationGet 获取关系
	RelationGet(ctx context.Context, in *RelationGetReq, opts ...client.CallOption) (*RelationGetResp, error)
	// RelationDel 删除关系
	RelationDel(ctx context.Context, in *RelationDelReq, opts ...client.CallOption) (*emptypb.Empty, error)
	// RelationCheck 关系检测
	RelationCheck(ctx context.Context, in *RelationCheckReq, opts ...client.CallOption) (*RelationCheckResp, error)
	// RelationMCheck 批量关系检测
	RelationMCheck(ctx context.Context, in *RelationMCheckReq, opts ...client.CallOption) (*RelationMCheckResp, error)
	// GetRelationCount 获取关系数量
	GetRelationCount(ctx context.Context, in *GetRelationCountReq, opts ...client.CallOption) (*GetRelationCountResp, error)
	// GetRelationList 获取关系列表
	GetRelationList(ctx context.Context, in *GetRelationListReq, opts ...client.CallOption) (*GetRelationListResp, error)
}

type relationService struct {
	c    client.Client
	name string
}

func NewRelationService(name string, c client.Client) RelationService {
	return &relationService{
		c:    c,
		name: name,
	}
}

func (c *relationService) RelationAdd(ctx context.Context, in *RelationAddReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Relation.RelationAdd", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationService) RelationGet(ctx context.Context, in *RelationGetReq, opts ...client.CallOption) (*RelationGetResp, error) {
	req := c.c.NewRequest(c.name, "Relation.RelationGet", in)
	out := new(RelationGetResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationService) RelationDel(ctx context.Context, in *RelationDelReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Relation.RelationDel", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationService) RelationCheck(ctx context.Context, in *RelationCheckReq, opts ...client.CallOption) (*RelationCheckResp, error) {
	req := c.c.NewRequest(c.name, "Relation.RelationCheck", in)
	out := new(RelationCheckResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationService) RelationMCheck(ctx context.Context, in *RelationMCheckReq, opts ...client.CallOption) (*RelationMCheckResp, error) {
	req := c.c.NewRequest(c.name, "Relation.RelationMCheck", in)
	out := new(RelationMCheckResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationService) GetRelationCount(ctx context.Context, in *GetRelationCountReq, opts ...client.CallOption) (*GetRelationCountResp, error) {
	req := c.c.NewRequest(c.name, "Relation.GetRelationCount", in)
	out := new(GetRelationCountResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationService) GetRelationList(ctx context.Context, in *GetRelationListReq, opts ...client.CallOption) (*GetRelationListResp, error) {
	req := c.c.NewRequest(c.name, "Relation.GetRelationList", in)
	out := new(GetRelationListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Relation service

type RelationHandler interface {
	// RelationAdd 新建关系
	RelationAdd(context.Context, *RelationAddReq, *emptypb.Empty) error
	// RelationGet 获取关系
	RelationGet(context.Context, *RelationGetReq, *RelationGetResp) error
	// RelationDel 删除关系
	RelationDel(context.Context, *RelationDelReq, *emptypb.Empty) error
	// RelationCheck 关系检测
	RelationCheck(context.Context, *RelationCheckReq, *RelationCheckResp) error
	// RelationMCheck 批量关系检测
	RelationMCheck(context.Context, *RelationMCheckReq, *RelationMCheckResp) error
	// GetRelationCount 获取关系数量
	GetRelationCount(context.Context, *GetRelationCountReq, *GetRelationCountResp) error
	// GetRelationList 获取关系列表
	GetRelationList(context.Context, *GetRelationListReq, *GetRelationListResp) error
}

func RegisterRelationHandler(s server.Server, hdlr RelationHandler, opts ...server.HandlerOption) error {
	type relation interface {
		RelationAdd(ctx context.Context, in *RelationAddReq, out *emptypb.Empty) error
		RelationGet(ctx context.Context, in *RelationGetReq, out *RelationGetResp) error
		RelationDel(ctx context.Context, in *RelationDelReq, out *emptypb.Empty) error
		RelationCheck(ctx context.Context, in *RelationCheckReq, out *RelationCheckResp) error
		RelationMCheck(ctx context.Context, in *RelationMCheckReq, out *RelationMCheckResp) error
		GetRelationCount(ctx context.Context, in *GetRelationCountReq, out *GetRelationCountResp) error
		GetRelationList(ctx context.Context, in *GetRelationListReq, out *GetRelationListResp) error
	}
	type Relation struct {
		relation
	}
	h := &relationHandler{hdlr}
	return s.Handle(s.NewHandler(&Relation{h}, opts...))
}

type relationHandler struct {
	RelationHandler
}

func (h *relationHandler) RelationAdd(ctx context.Context, in *RelationAddReq, out *emptypb.Empty) error {
	return h.RelationHandler.RelationAdd(ctx, in, out)
}

func (h *relationHandler) RelationGet(ctx context.Context, in *RelationGetReq, out *RelationGetResp) error {
	return h.RelationHandler.RelationGet(ctx, in, out)
}

func (h *relationHandler) RelationDel(ctx context.Context, in *RelationDelReq, out *emptypb.Empty) error {
	return h.RelationHandler.RelationDel(ctx, in, out)
}

func (h *relationHandler) RelationCheck(ctx context.Context, in *RelationCheckReq, out *RelationCheckResp) error {
	return h.RelationHandler.RelationCheck(ctx, in, out)
}

func (h *relationHandler) RelationMCheck(ctx context.Context, in *RelationMCheckReq, out *RelationMCheckResp) error {
	return h.RelationHandler.RelationMCheck(ctx, in, out)
}

func (h *relationHandler) GetRelationCount(ctx context.Context, in *GetRelationCountReq, out *GetRelationCountResp) error {
	return h.RelationHandler.GetRelationCount(ctx, in, out)
}

func (h *relationHandler) GetRelationList(ctx context.Context, in *GetRelationListReq, out *GetRelationListResp) error {
	return h.RelationHandler.GetRelationList(ctx, in, out)
}
