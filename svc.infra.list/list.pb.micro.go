// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.infra.list/list.proto

package list

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

// Api Endpoints for List service

func NewListEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for List service

type ListService interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error)
	// Methods
	GetItem(ctx context.Context, in *GetItemReq, opts ...client.CallOption) (*GetItemResp, error)
	AddItem(ctx context.Context, in *AddItemReq, opts ...client.CallOption) (*AddItemResp, error)
	DeleteItem(ctx context.Context, in *DeleteItemReq, opts ...client.CallOption) (*DeleteItemResp, error)
	UpdateItem(ctx context.Context, in *UpdateItemReq, opts ...client.CallOption) (*UpdateItemResp, error)
	ListItems(ctx context.Context, in *ListItemsReq, opts ...client.CallOption) (*ListItemsResp, error)
	TotalItems(ctx context.Context, in *TotalItemsReq, opts ...client.CallOption) (*TotalItemsResp, error)
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

func (c *listService) GetItem(ctx context.Context, in *GetItemReq, opts ...client.CallOption) (*GetItemResp, error) {
	req := c.c.NewRequest(c.name, "List.GetItem", in)
	out := new(GetItemResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listService) AddItem(ctx context.Context, in *AddItemReq, opts ...client.CallOption) (*AddItemResp, error) {
	req := c.c.NewRequest(c.name, "List.AddItem", in)
	out := new(AddItemResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listService) DeleteItem(ctx context.Context, in *DeleteItemReq, opts ...client.CallOption) (*DeleteItemResp, error) {
	req := c.c.NewRequest(c.name, "List.DeleteItem", in)
	out := new(DeleteItemResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listService) UpdateItem(ctx context.Context, in *UpdateItemReq, opts ...client.CallOption) (*UpdateItemResp, error) {
	req := c.c.NewRequest(c.name, "List.UpdateItem", in)
	out := new(UpdateItemResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listService) ListItems(ctx context.Context, in *ListItemsReq, opts ...client.CallOption) (*ListItemsResp, error) {
	req := c.c.NewRequest(c.name, "List.ListItems", in)
	out := new(ListItemsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listService) TotalItems(ctx context.Context, in *TotalItemsReq, opts ...client.CallOption) (*TotalItemsResp, error) {
	req := c.c.NewRequest(c.name, "List.TotalItems", in)
	out := new(TotalItemsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for List service

type ListHandler interface {
	InitDB(context.Context, *emptypb.Empty, *InitDBResp) error
	// Methods
	GetItem(context.Context, *GetItemReq, *GetItemResp) error
	AddItem(context.Context, *AddItemReq, *AddItemResp) error
	DeleteItem(context.Context, *DeleteItemReq, *DeleteItemResp) error
	UpdateItem(context.Context, *UpdateItemReq, *UpdateItemResp) error
	ListItems(context.Context, *ListItemsReq, *ListItemsResp) error
	TotalItems(context.Context, *TotalItemsReq, *TotalItemsResp) error
}

func RegisterListHandler(s server.Server, hdlr ListHandler, opts ...server.HandlerOption) error {
	type list interface {
		InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error
		GetItem(ctx context.Context, in *GetItemReq, out *GetItemResp) error
		AddItem(ctx context.Context, in *AddItemReq, out *AddItemResp) error
		DeleteItem(ctx context.Context, in *DeleteItemReq, out *DeleteItemResp) error
		UpdateItem(ctx context.Context, in *UpdateItemReq, out *UpdateItemResp) error
		ListItems(ctx context.Context, in *ListItemsReq, out *ListItemsResp) error
		TotalItems(ctx context.Context, in *TotalItemsReq, out *TotalItemsResp) error
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

func (h *listHandler) GetItem(ctx context.Context, in *GetItemReq, out *GetItemResp) error {
	return h.ListHandler.GetItem(ctx, in, out)
}

func (h *listHandler) AddItem(ctx context.Context, in *AddItemReq, out *AddItemResp) error {
	return h.ListHandler.AddItem(ctx, in, out)
}

func (h *listHandler) DeleteItem(ctx context.Context, in *DeleteItemReq, out *DeleteItemResp) error {
	return h.ListHandler.DeleteItem(ctx, in, out)
}

func (h *listHandler) UpdateItem(ctx context.Context, in *UpdateItemReq, out *UpdateItemResp) error {
	return h.ListHandler.UpdateItem(ctx, in, out)
}

func (h *listHandler) ListItems(ctx context.Context, in *ListItemsReq, out *ListItemsResp) error {
	return h.ListHandler.ListItems(ctx, in, out)
}

func (h *listHandler) TotalItems(ctx context.Context, in *TotalItemsReq, out *TotalItemsResp) error {
	return h.ListHandler.TotalItems(ctx, in, out)
}
