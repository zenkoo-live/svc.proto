// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.org/org.proto

package org

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

// Api Endpoints for Org service

func NewOrgEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Org service

type OrgService interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error)
	GetDepartment(ctx context.Context, in *GetDepartmentReq, opts ...client.CallOption) (*GetDepartmentResp, error)
	ListDepartments(ctx context.Context, in *ListDepartmentsReq, opts ...client.CallOption) (*ListDepartmentsResp, error)
	FilterDepartments(ctx context.Context, in *FilterDepartmentsReq, opts ...client.CallOption) (*FilterDepartmentsResp, error)
	AddDepartment(ctx context.Context, in *AddDepartmentReq, opts ...client.CallOption) (*AddDepartmentResp, error)
	UpdateDepartment(ctx context.Context, in *UpdateDepartmentReq, opts ...client.CallOption) (*UpdateDepartmentResp, error)
	DeleteDepartment(ctx context.Context, in *DeleteDepartmentReq, opts ...client.CallOption) (*DeleteDepartmentResp, error)
	TotalDepartments(ctx context.Context, in *TotalDepartmentsReq, opts ...client.CallOption) (*TotalDepartmentsResp, error)
	DepartmentAdditionsSet(ctx context.Context, in *DepartmentAdditionsSetReq, opts ...client.CallOption) (*DepartmentAdditionsSetResp, error)
	DepartmentAdditionsGet(ctx context.Context, in *DepartmentAdditionsGetReq, opts ...client.CallOption) (*DepartmentAdditionsGetResp, error)
	DepartmentAdditionsDump(ctx context.Context, in *DepartmentAdditionsDumpReq, opts ...client.CallOption) (*DepartmentAdditionsDumpResp, error)
	GetMerchant(ctx context.Context, in *GetMerchantReq, opts ...client.CallOption) (*GetMerchantResp, error)
	ListMerchants(ctx context.Context, in *ListMerchantsReq, opts ...client.CallOption) (*ListMerchantsResp, error)
	FilterMerchants(ctx context.Context, in *FilterMerchantsReq, opts ...client.CallOption) (*FilterMerchantsResp, error)
	AddMerchant(ctx context.Context, in *AddMerchantReq, opts ...client.CallOption) (*AddMerchantResp, error)
	UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...client.CallOption) (*UpdateMerchantResp, error)
	DeleteMerchant(ctx context.Context, in *DeleteMerchantReq, opts ...client.CallOption) (*DeleteMerchantResp, error)
	TotalMerchants(ctx context.Context, in *TotalMerchantsReq, opts ...client.CallOption) (*TotalMerchantsResp, error)
	MerchantAdditionsSet(ctx context.Context, in *MerchantAdditionsSetReq, opts ...client.CallOption) (*MerchantAdditionsSetResp, error)
	MerchantAdditionsGet(ctx context.Context, in *MerchantAdditionsGetReq, opts ...client.CallOption) (*MerchantAdditionsGetResp, error)
	MerchantAdditionsDump(ctx context.Context, in *MerchantAdditionsDumpReq, opts ...client.CallOption) (*MerchantAdditionsDumpResp, error)
	GetUnion(ctx context.Context, in *GetUnionReq, opts ...client.CallOption) (*GetUnionResp, error)
	ListUnions(ctx context.Context, in *ListUnionsReq, opts ...client.CallOption) (*ListUnionsResp, error)
	FilterUnions(ctx context.Context, in *FilterUnionsReq, opts ...client.CallOption) (*FilterUnionsResp, error)
	AddUnion(ctx context.Context, in *AddUnionReq, opts ...client.CallOption) (*AddUnionResp, error)
	UpdateUnion(ctx context.Context, in *UpdateUnionReq, opts ...client.CallOption) (*UpdateUnionResp, error)
	DeleteUnion(ctx context.Context, in *DeleteUnionReq, opts ...client.CallOption) (*DeleteUnionResp, error)
	TotalUnions(ctx context.Context, in *TotalUnionsReq, opts ...client.CallOption) (*TotalUnionsResp, error)
	UnionAdditionsSet(ctx context.Context, in *UnionAdditionsSetReq, opts ...client.CallOption) (*UnionAdditionsSetResp, error)
	UnionAdditionsGet(ctx context.Context, in *UnionAdditionsGetReq, opts ...client.CallOption) (*UnionAdditionsGetResp, error)
	UnionAdditionsDump(ctx context.Context, in *UnionAdditionsDumpReq, opts ...client.CallOption) (*UnionAdditionsDumpResp, error)
}

type orgService struct {
	c    client.Client
	name string
}

func NewOrgService(name string, c client.Client) OrgService {
	return &orgService{
		c:    c,
		name: name,
	}
}

func (c *orgService) InitDB(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*InitDBResp, error) {
	req := c.c.NewRequest(c.name, "Org.InitDB", in)
	out := new(InitDBResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) GetDepartment(ctx context.Context, in *GetDepartmentReq, opts ...client.CallOption) (*GetDepartmentResp, error) {
	req := c.c.NewRequest(c.name, "Org.GetDepartment", in)
	out := new(GetDepartmentResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) ListDepartments(ctx context.Context, in *ListDepartmentsReq, opts ...client.CallOption) (*ListDepartmentsResp, error) {
	req := c.c.NewRequest(c.name, "Org.ListDepartments", in)
	out := new(ListDepartmentsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) FilterDepartments(ctx context.Context, in *FilterDepartmentsReq, opts ...client.CallOption) (*FilterDepartmentsResp, error) {
	req := c.c.NewRequest(c.name, "Org.FilterDepartments", in)
	out := new(FilterDepartmentsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) AddDepartment(ctx context.Context, in *AddDepartmentReq, opts ...client.CallOption) (*AddDepartmentResp, error) {
	req := c.c.NewRequest(c.name, "Org.AddDepartment", in)
	out := new(AddDepartmentResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) UpdateDepartment(ctx context.Context, in *UpdateDepartmentReq, opts ...client.CallOption) (*UpdateDepartmentResp, error) {
	req := c.c.NewRequest(c.name, "Org.UpdateDepartment", in)
	out := new(UpdateDepartmentResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) DeleteDepartment(ctx context.Context, in *DeleteDepartmentReq, opts ...client.CallOption) (*DeleteDepartmentResp, error) {
	req := c.c.NewRequest(c.name, "Org.DeleteDepartment", in)
	out := new(DeleteDepartmentResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) TotalDepartments(ctx context.Context, in *TotalDepartmentsReq, opts ...client.CallOption) (*TotalDepartmentsResp, error) {
	req := c.c.NewRequest(c.name, "Org.TotalDepartments", in)
	out := new(TotalDepartmentsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) DepartmentAdditionsSet(ctx context.Context, in *DepartmentAdditionsSetReq, opts ...client.CallOption) (*DepartmentAdditionsSetResp, error) {
	req := c.c.NewRequest(c.name, "Org.DepartmentAdditionsSet", in)
	out := new(DepartmentAdditionsSetResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) DepartmentAdditionsGet(ctx context.Context, in *DepartmentAdditionsGetReq, opts ...client.CallOption) (*DepartmentAdditionsGetResp, error) {
	req := c.c.NewRequest(c.name, "Org.DepartmentAdditionsGet", in)
	out := new(DepartmentAdditionsGetResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) DepartmentAdditionsDump(ctx context.Context, in *DepartmentAdditionsDumpReq, opts ...client.CallOption) (*DepartmentAdditionsDumpResp, error) {
	req := c.c.NewRequest(c.name, "Org.DepartmentAdditionsDump", in)
	out := new(DepartmentAdditionsDumpResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) GetMerchant(ctx context.Context, in *GetMerchantReq, opts ...client.CallOption) (*GetMerchantResp, error) {
	req := c.c.NewRequest(c.name, "Org.GetMerchant", in)
	out := new(GetMerchantResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) ListMerchants(ctx context.Context, in *ListMerchantsReq, opts ...client.CallOption) (*ListMerchantsResp, error) {
	req := c.c.NewRequest(c.name, "Org.ListMerchants", in)
	out := new(ListMerchantsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) FilterMerchants(ctx context.Context, in *FilterMerchantsReq, opts ...client.CallOption) (*FilterMerchantsResp, error) {
	req := c.c.NewRequest(c.name, "Org.FilterMerchants", in)
	out := new(FilterMerchantsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) AddMerchant(ctx context.Context, in *AddMerchantReq, opts ...client.CallOption) (*AddMerchantResp, error) {
	req := c.c.NewRequest(c.name, "Org.AddMerchant", in)
	out := new(AddMerchantResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...client.CallOption) (*UpdateMerchantResp, error) {
	req := c.c.NewRequest(c.name, "Org.UpdateMerchant", in)
	out := new(UpdateMerchantResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) DeleteMerchant(ctx context.Context, in *DeleteMerchantReq, opts ...client.CallOption) (*DeleteMerchantResp, error) {
	req := c.c.NewRequest(c.name, "Org.DeleteMerchant", in)
	out := new(DeleteMerchantResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) TotalMerchants(ctx context.Context, in *TotalMerchantsReq, opts ...client.CallOption) (*TotalMerchantsResp, error) {
	req := c.c.NewRequest(c.name, "Org.TotalMerchants", in)
	out := new(TotalMerchantsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) MerchantAdditionsSet(ctx context.Context, in *MerchantAdditionsSetReq, opts ...client.CallOption) (*MerchantAdditionsSetResp, error) {
	req := c.c.NewRequest(c.name, "Org.MerchantAdditionsSet", in)
	out := new(MerchantAdditionsSetResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) MerchantAdditionsGet(ctx context.Context, in *MerchantAdditionsGetReq, opts ...client.CallOption) (*MerchantAdditionsGetResp, error) {
	req := c.c.NewRequest(c.name, "Org.MerchantAdditionsGet", in)
	out := new(MerchantAdditionsGetResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) MerchantAdditionsDump(ctx context.Context, in *MerchantAdditionsDumpReq, opts ...client.CallOption) (*MerchantAdditionsDumpResp, error) {
	req := c.c.NewRequest(c.name, "Org.MerchantAdditionsDump", in)
	out := new(MerchantAdditionsDumpResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) GetUnion(ctx context.Context, in *GetUnionReq, opts ...client.CallOption) (*GetUnionResp, error) {
	req := c.c.NewRequest(c.name, "Org.GetUnion", in)
	out := new(GetUnionResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) ListUnions(ctx context.Context, in *ListUnionsReq, opts ...client.CallOption) (*ListUnionsResp, error) {
	req := c.c.NewRequest(c.name, "Org.ListUnions", in)
	out := new(ListUnionsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) FilterUnions(ctx context.Context, in *FilterUnionsReq, opts ...client.CallOption) (*FilterUnionsResp, error) {
	req := c.c.NewRequest(c.name, "Org.FilterUnions", in)
	out := new(FilterUnionsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) AddUnion(ctx context.Context, in *AddUnionReq, opts ...client.CallOption) (*AddUnionResp, error) {
	req := c.c.NewRequest(c.name, "Org.AddUnion", in)
	out := new(AddUnionResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) UpdateUnion(ctx context.Context, in *UpdateUnionReq, opts ...client.CallOption) (*UpdateUnionResp, error) {
	req := c.c.NewRequest(c.name, "Org.UpdateUnion", in)
	out := new(UpdateUnionResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) DeleteUnion(ctx context.Context, in *DeleteUnionReq, opts ...client.CallOption) (*DeleteUnionResp, error) {
	req := c.c.NewRequest(c.name, "Org.DeleteUnion", in)
	out := new(DeleteUnionResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) TotalUnions(ctx context.Context, in *TotalUnionsReq, opts ...client.CallOption) (*TotalUnionsResp, error) {
	req := c.c.NewRequest(c.name, "Org.TotalUnions", in)
	out := new(TotalUnionsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) UnionAdditionsSet(ctx context.Context, in *UnionAdditionsSetReq, opts ...client.CallOption) (*UnionAdditionsSetResp, error) {
	req := c.c.NewRequest(c.name, "Org.UnionAdditionsSet", in)
	out := new(UnionAdditionsSetResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) UnionAdditionsGet(ctx context.Context, in *UnionAdditionsGetReq, opts ...client.CallOption) (*UnionAdditionsGetResp, error) {
	req := c.c.NewRequest(c.name, "Org.UnionAdditionsGet", in)
	out := new(UnionAdditionsGetResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) UnionAdditionsDump(ctx context.Context, in *UnionAdditionsDumpReq, opts ...client.CallOption) (*UnionAdditionsDumpResp, error) {
	req := c.c.NewRequest(c.name, "Org.UnionAdditionsDump", in)
	out := new(UnionAdditionsDumpResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Org service

type OrgHandler interface {
	InitDB(context.Context, *emptypb.Empty, *InitDBResp) error
	GetDepartment(context.Context, *GetDepartmentReq, *GetDepartmentResp) error
	ListDepartments(context.Context, *ListDepartmentsReq, *ListDepartmentsResp) error
	FilterDepartments(context.Context, *FilterDepartmentsReq, *FilterDepartmentsResp) error
	AddDepartment(context.Context, *AddDepartmentReq, *AddDepartmentResp) error
	UpdateDepartment(context.Context, *UpdateDepartmentReq, *UpdateDepartmentResp) error
	DeleteDepartment(context.Context, *DeleteDepartmentReq, *DeleteDepartmentResp) error
	TotalDepartments(context.Context, *TotalDepartmentsReq, *TotalDepartmentsResp) error
	DepartmentAdditionsSet(context.Context, *DepartmentAdditionsSetReq, *DepartmentAdditionsSetResp) error
	DepartmentAdditionsGet(context.Context, *DepartmentAdditionsGetReq, *DepartmentAdditionsGetResp) error
	DepartmentAdditionsDump(context.Context, *DepartmentAdditionsDumpReq, *DepartmentAdditionsDumpResp) error
	GetMerchant(context.Context, *GetMerchantReq, *GetMerchantResp) error
	ListMerchants(context.Context, *ListMerchantsReq, *ListMerchantsResp) error
	FilterMerchants(context.Context, *FilterMerchantsReq, *FilterMerchantsResp) error
	AddMerchant(context.Context, *AddMerchantReq, *AddMerchantResp) error
	UpdateMerchant(context.Context, *UpdateMerchantReq, *UpdateMerchantResp) error
	DeleteMerchant(context.Context, *DeleteMerchantReq, *DeleteMerchantResp) error
	TotalMerchants(context.Context, *TotalMerchantsReq, *TotalMerchantsResp) error
	MerchantAdditionsSet(context.Context, *MerchantAdditionsSetReq, *MerchantAdditionsSetResp) error
	MerchantAdditionsGet(context.Context, *MerchantAdditionsGetReq, *MerchantAdditionsGetResp) error
	MerchantAdditionsDump(context.Context, *MerchantAdditionsDumpReq, *MerchantAdditionsDumpResp) error
	GetUnion(context.Context, *GetUnionReq, *GetUnionResp) error
	ListUnions(context.Context, *ListUnionsReq, *ListUnionsResp) error
	FilterUnions(context.Context, *FilterUnionsReq, *FilterUnionsResp) error
	AddUnion(context.Context, *AddUnionReq, *AddUnionResp) error
	UpdateUnion(context.Context, *UpdateUnionReq, *UpdateUnionResp) error
	DeleteUnion(context.Context, *DeleteUnionReq, *DeleteUnionResp) error
	TotalUnions(context.Context, *TotalUnionsReq, *TotalUnionsResp) error
	UnionAdditionsSet(context.Context, *UnionAdditionsSetReq, *UnionAdditionsSetResp) error
	UnionAdditionsGet(context.Context, *UnionAdditionsGetReq, *UnionAdditionsGetResp) error
	UnionAdditionsDump(context.Context, *UnionAdditionsDumpReq, *UnionAdditionsDumpResp) error
}

func RegisterOrgHandler(s server.Server, hdlr OrgHandler, opts ...server.HandlerOption) error {
	type org interface {
		InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error
		GetDepartment(ctx context.Context, in *GetDepartmentReq, out *GetDepartmentResp) error
		ListDepartments(ctx context.Context, in *ListDepartmentsReq, out *ListDepartmentsResp) error
		FilterDepartments(ctx context.Context, in *FilterDepartmentsReq, out *FilterDepartmentsResp) error
		AddDepartment(ctx context.Context, in *AddDepartmentReq, out *AddDepartmentResp) error
		UpdateDepartment(ctx context.Context, in *UpdateDepartmentReq, out *UpdateDepartmentResp) error
		DeleteDepartment(ctx context.Context, in *DeleteDepartmentReq, out *DeleteDepartmentResp) error
		TotalDepartments(ctx context.Context, in *TotalDepartmentsReq, out *TotalDepartmentsResp) error
		DepartmentAdditionsSet(ctx context.Context, in *DepartmentAdditionsSetReq, out *DepartmentAdditionsSetResp) error
		DepartmentAdditionsGet(ctx context.Context, in *DepartmentAdditionsGetReq, out *DepartmentAdditionsGetResp) error
		DepartmentAdditionsDump(ctx context.Context, in *DepartmentAdditionsDumpReq, out *DepartmentAdditionsDumpResp) error
		GetMerchant(ctx context.Context, in *GetMerchantReq, out *GetMerchantResp) error
		ListMerchants(ctx context.Context, in *ListMerchantsReq, out *ListMerchantsResp) error
		FilterMerchants(ctx context.Context, in *FilterMerchantsReq, out *FilterMerchantsResp) error
		AddMerchant(ctx context.Context, in *AddMerchantReq, out *AddMerchantResp) error
		UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, out *UpdateMerchantResp) error
		DeleteMerchant(ctx context.Context, in *DeleteMerchantReq, out *DeleteMerchantResp) error
		TotalMerchants(ctx context.Context, in *TotalMerchantsReq, out *TotalMerchantsResp) error
		MerchantAdditionsSet(ctx context.Context, in *MerchantAdditionsSetReq, out *MerchantAdditionsSetResp) error
		MerchantAdditionsGet(ctx context.Context, in *MerchantAdditionsGetReq, out *MerchantAdditionsGetResp) error
		MerchantAdditionsDump(ctx context.Context, in *MerchantAdditionsDumpReq, out *MerchantAdditionsDumpResp) error
		GetUnion(ctx context.Context, in *GetUnionReq, out *GetUnionResp) error
		ListUnions(ctx context.Context, in *ListUnionsReq, out *ListUnionsResp) error
		FilterUnions(ctx context.Context, in *FilterUnionsReq, out *FilterUnionsResp) error
		AddUnion(ctx context.Context, in *AddUnionReq, out *AddUnionResp) error
		UpdateUnion(ctx context.Context, in *UpdateUnionReq, out *UpdateUnionResp) error
		DeleteUnion(ctx context.Context, in *DeleteUnionReq, out *DeleteUnionResp) error
		TotalUnions(ctx context.Context, in *TotalUnionsReq, out *TotalUnionsResp) error
		UnionAdditionsSet(ctx context.Context, in *UnionAdditionsSetReq, out *UnionAdditionsSetResp) error
		UnionAdditionsGet(ctx context.Context, in *UnionAdditionsGetReq, out *UnionAdditionsGetResp) error
		UnionAdditionsDump(ctx context.Context, in *UnionAdditionsDumpReq, out *UnionAdditionsDumpResp) error
	}
	type Org struct {
		org
	}
	h := &orgHandler{hdlr}
	return s.Handle(s.NewHandler(&Org{h}, opts...))
}

type orgHandler struct {
	OrgHandler
}

func (h *orgHandler) InitDB(ctx context.Context, in *emptypb.Empty, out *InitDBResp) error {
	return h.OrgHandler.InitDB(ctx, in, out)
}

func (h *orgHandler) GetDepartment(ctx context.Context, in *GetDepartmentReq, out *GetDepartmentResp) error {
	return h.OrgHandler.GetDepartment(ctx, in, out)
}

func (h *orgHandler) ListDepartments(ctx context.Context, in *ListDepartmentsReq, out *ListDepartmentsResp) error {
	return h.OrgHandler.ListDepartments(ctx, in, out)
}

func (h *orgHandler) FilterDepartments(ctx context.Context, in *FilterDepartmentsReq, out *FilterDepartmentsResp) error {
	return h.OrgHandler.FilterDepartments(ctx, in, out)
}

func (h *orgHandler) AddDepartment(ctx context.Context, in *AddDepartmentReq, out *AddDepartmentResp) error {
	return h.OrgHandler.AddDepartment(ctx, in, out)
}

func (h *orgHandler) UpdateDepartment(ctx context.Context, in *UpdateDepartmentReq, out *UpdateDepartmentResp) error {
	return h.OrgHandler.UpdateDepartment(ctx, in, out)
}

func (h *orgHandler) DeleteDepartment(ctx context.Context, in *DeleteDepartmentReq, out *DeleteDepartmentResp) error {
	return h.OrgHandler.DeleteDepartment(ctx, in, out)
}

func (h *orgHandler) TotalDepartments(ctx context.Context, in *TotalDepartmentsReq, out *TotalDepartmentsResp) error {
	return h.OrgHandler.TotalDepartments(ctx, in, out)
}

func (h *orgHandler) DepartmentAdditionsSet(ctx context.Context, in *DepartmentAdditionsSetReq, out *DepartmentAdditionsSetResp) error {
	return h.OrgHandler.DepartmentAdditionsSet(ctx, in, out)
}

func (h *orgHandler) DepartmentAdditionsGet(ctx context.Context, in *DepartmentAdditionsGetReq, out *DepartmentAdditionsGetResp) error {
	return h.OrgHandler.DepartmentAdditionsGet(ctx, in, out)
}

func (h *orgHandler) DepartmentAdditionsDump(ctx context.Context, in *DepartmentAdditionsDumpReq, out *DepartmentAdditionsDumpResp) error {
	return h.OrgHandler.DepartmentAdditionsDump(ctx, in, out)
}

func (h *orgHandler) GetMerchant(ctx context.Context, in *GetMerchantReq, out *GetMerchantResp) error {
	return h.OrgHandler.GetMerchant(ctx, in, out)
}

func (h *orgHandler) ListMerchants(ctx context.Context, in *ListMerchantsReq, out *ListMerchantsResp) error {
	return h.OrgHandler.ListMerchants(ctx, in, out)
}

func (h *orgHandler) FilterMerchants(ctx context.Context, in *FilterMerchantsReq, out *FilterMerchantsResp) error {
	return h.OrgHandler.FilterMerchants(ctx, in, out)
}

func (h *orgHandler) AddMerchant(ctx context.Context, in *AddMerchantReq, out *AddMerchantResp) error {
	return h.OrgHandler.AddMerchant(ctx, in, out)
}

func (h *orgHandler) UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, out *UpdateMerchantResp) error {
	return h.OrgHandler.UpdateMerchant(ctx, in, out)
}

func (h *orgHandler) DeleteMerchant(ctx context.Context, in *DeleteMerchantReq, out *DeleteMerchantResp) error {
	return h.OrgHandler.DeleteMerchant(ctx, in, out)
}

func (h *orgHandler) TotalMerchants(ctx context.Context, in *TotalMerchantsReq, out *TotalMerchantsResp) error {
	return h.OrgHandler.TotalMerchants(ctx, in, out)
}

func (h *orgHandler) MerchantAdditionsSet(ctx context.Context, in *MerchantAdditionsSetReq, out *MerchantAdditionsSetResp) error {
	return h.OrgHandler.MerchantAdditionsSet(ctx, in, out)
}

func (h *orgHandler) MerchantAdditionsGet(ctx context.Context, in *MerchantAdditionsGetReq, out *MerchantAdditionsGetResp) error {
	return h.OrgHandler.MerchantAdditionsGet(ctx, in, out)
}

func (h *orgHandler) MerchantAdditionsDump(ctx context.Context, in *MerchantAdditionsDumpReq, out *MerchantAdditionsDumpResp) error {
	return h.OrgHandler.MerchantAdditionsDump(ctx, in, out)
}

func (h *orgHandler) GetUnion(ctx context.Context, in *GetUnionReq, out *GetUnionResp) error {
	return h.OrgHandler.GetUnion(ctx, in, out)
}

func (h *orgHandler) ListUnions(ctx context.Context, in *ListUnionsReq, out *ListUnionsResp) error {
	return h.OrgHandler.ListUnions(ctx, in, out)
}

func (h *orgHandler) FilterUnions(ctx context.Context, in *FilterUnionsReq, out *FilterUnionsResp) error {
	return h.OrgHandler.FilterUnions(ctx, in, out)
}

func (h *orgHandler) AddUnion(ctx context.Context, in *AddUnionReq, out *AddUnionResp) error {
	return h.OrgHandler.AddUnion(ctx, in, out)
}

func (h *orgHandler) UpdateUnion(ctx context.Context, in *UpdateUnionReq, out *UpdateUnionResp) error {
	return h.OrgHandler.UpdateUnion(ctx, in, out)
}

func (h *orgHandler) DeleteUnion(ctx context.Context, in *DeleteUnionReq, out *DeleteUnionResp) error {
	return h.OrgHandler.DeleteUnion(ctx, in, out)
}

func (h *orgHandler) TotalUnions(ctx context.Context, in *TotalUnionsReq, out *TotalUnionsResp) error {
	return h.OrgHandler.TotalUnions(ctx, in, out)
}

func (h *orgHandler) UnionAdditionsSet(ctx context.Context, in *UnionAdditionsSetReq, out *UnionAdditionsSetResp) error {
	return h.OrgHandler.UnionAdditionsSet(ctx, in, out)
}

func (h *orgHandler) UnionAdditionsGet(ctx context.Context, in *UnionAdditionsGetReq, out *UnionAdditionsGetResp) error {
	return h.OrgHandler.UnionAdditionsGet(ctx, in, out)
}

func (h *orgHandler) UnionAdditionsDump(ctx context.Context, in *UnionAdditionsDumpReq, out *UnionAdditionsDumpResp) error {
	return h.OrgHandler.UnionAdditionsDump(ctx, in, out)
}
