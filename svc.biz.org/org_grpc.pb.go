// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: svc.biz.org/org.proto

package org

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Org_InitDB_FullMethodName           = "/svc.biz.org.Org/InitDB"
	Org_GetDepartment_FullMethodName    = "/svc.biz.org.Org/GetDepartment"
	Org_ListDepartments_FullMethodName  = "/svc.biz.org.Org/ListDepartments"
	Org_AddDepartment_FullMethodName    = "/svc.biz.org.Org/AddDepartment"
	Org_UpdateDepartment_FullMethodName = "/svc.biz.org.Org/UpdateDepartment"
	Org_DeleteDepartment_FullMethodName = "/svc.biz.org.Org/DeleteDepartment"
	Org_GetMerchant_FullMethodName      = "/svc.biz.org.Org/GetMerchant"
	Org_ListMerchants_FullMethodName    = "/svc.biz.org.Org/ListMerchants"
	Org_AddMerchant_FullMethodName      = "/svc.biz.org.Org/AddMerchant"
	Org_UpdateMerchant_FullMethodName   = "/svc.biz.org.Org/UpdateMerchant"
	Org_DeleteMerchant_FullMethodName   = "/svc.biz.org.Org/DeleteMerchant"
	Org_GetUnion_FullMethodName         = "/svc.biz.org.Org/GetUnion"
	Org_ListUnions_FullMethodName       = "/svc.biz.org.Org/ListUnions"
	Org_AddUnion_FullMethodName         = "/svc.biz.org.Org/AddUnion"
	Org_UpdateUnion_FullMethodName      = "/svc.biz.org.Org/UpdateUnion"
	Org_DeleteUnion_FullMethodName      = "/svc.biz.org.Org/DeleteUnion"
)

// OrgClient is the client API for Org service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrgClient interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error)
	GetDepartment(ctx context.Context, in *GetDepartmentReq, opts ...grpc.CallOption) (*GetDepartmentResp, error)
	ListDepartments(ctx context.Context, in *ListDepartmentsReq, opts ...grpc.CallOption) (*ListDepartmentsResp, error)
	AddDepartment(ctx context.Context, in *AddDepartmentReq, opts ...grpc.CallOption) (*AddDepartmentResp, error)
	UpdateDepartment(ctx context.Context, in *UpdateDepartmentReq, opts ...grpc.CallOption) (*UpdateDepartmentResp, error)
	DeleteDepartment(ctx context.Context, in *DeleteDepartmentReq, opts ...grpc.CallOption) (*DeleteDepartmentResp, error)
	GetMerchant(ctx context.Context, in *GetMerchantReq, opts ...grpc.CallOption) (*GetMerchantResp, error)
	ListMerchants(ctx context.Context, in *ListMerchantsReq, opts ...grpc.CallOption) (*ListMerchantsResp, error)
	AddMerchant(ctx context.Context, in *AddMerchantReq, opts ...grpc.CallOption) (*AddMerchantResp, error)
	UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...grpc.CallOption) (*UpdateMerchantResp, error)
	DeleteMerchant(ctx context.Context, in *DeleteMerchantReq, opts ...grpc.CallOption) (*DeleteMerchantResp, error)
	GetUnion(ctx context.Context, in *GetUnionReq, opts ...grpc.CallOption) (*GetUnionResp, error)
	ListUnions(ctx context.Context, in *ListUnionsReq, opts ...grpc.CallOption) (*ListUnionsResp, error)
	AddUnion(ctx context.Context, in *AddUnionReq, opts ...grpc.CallOption) (*AddUnionResp, error)
	UpdateUnion(ctx context.Context, in *UpdateUnionReq, opts ...grpc.CallOption) (*UpdateUnionResp, error)
	DeleteUnion(ctx context.Context, in *DeleteUnionReq, opts ...grpc.CallOption) (*DeleteUnionResp, error)
}

type orgClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgClient(cc grpc.ClientConnInterface) OrgClient {
	return &orgClient{cc}
}

func (c *orgClient) InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitDBResp)
	err := c.cc.Invoke(ctx, Org_InitDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) GetDepartment(ctx context.Context, in *GetDepartmentReq, opts ...grpc.CallOption) (*GetDepartmentResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDepartmentResp)
	err := c.cc.Invoke(ctx, Org_GetDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) ListDepartments(ctx context.Context, in *ListDepartmentsReq, opts ...grpc.CallOption) (*ListDepartmentsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDepartmentsResp)
	err := c.cc.Invoke(ctx, Org_ListDepartments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) AddDepartment(ctx context.Context, in *AddDepartmentReq, opts ...grpc.CallOption) (*AddDepartmentResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddDepartmentResp)
	err := c.cc.Invoke(ctx, Org_AddDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) UpdateDepartment(ctx context.Context, in *UpdateDepartmentReq, opts ...grpc.CallOption) (*UpdateDepartmentResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDepartmentResp)
	err := c.cc.Invoke(ctx, Org_UpdateDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) DeleteDepartment(ctx context.Context, in *DeleteDepartmentReq, opts ...grpc.CallOption) (*DeleteDepartmentResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDepartmentResp)
	err := c.cc.Invoke(ctx, Org_DeleteDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) GetMerchant(ctx context.Context, in *GetMerchantReq, opts ...grpc.CallOption) (*GetMerchantResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMerchantResp)
	err := c.cc.Invoke(ctx, Org_GetMerchant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) ListMerchants(ctx context.Context, in *ListMerchantsReq, opts ...grpc.CallOption) (*ListMerchantsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMerchantsResp)
	err := c.cc.Invoke(ctx, Org_ListMerchants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) AddMerchant(ctx context.Context, in *AddMerchantReq, opts ...grpc.CallOption) (*AddMerchantResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMerchantResp)
	err := c.cc.Invoke(ctx, Org_AddMerchant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...grpc.CallOption) (*UpdateMerchantResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMerchantResp)
	err := c.cc.Invoke(ctx, Org_UpdateMerchant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) DeleteMerchant(ctx context.Context, in *DeleteMerchantReq, opts ...grpc.CallOption) (*DeleteMerchantResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMerchantResp)
	err := c.cc.Invoke(ctx, Org_DeleteMerchant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) GetUnion(ctx context.Context, in *GetUnionReq, opts ...grpc.CallOption) (*GetUnionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUnionResp)
	err := c.cc.Invoke(ctx, Org_GetUnion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) ListUnions(ctx context.Context, in *ListUnionsReq, opts ...grpc.CallOption) (*ListUnionsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUnionsResp)
	err := c.cc.Invoke(ctx, Org_ListUnions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) AddUnion(ctx context.Context, in *AddUnionReq, opts ...grpc.CallOption) (*AddUnionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUnionResp)
	err := c.cc.Invoke(ctx, Org_AddUnion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) UpdateUnion(ctx context.Context, in *UpdateUnionReq, opts ...grpc.CallOption) (*UpdateUnionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUnionResp)
	err := c.cc.Invoke(ctx, Org_UpdateUnion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgClient) DeleteUnion(ctx context.Context, in *DeleteUnionReq, opts ...grpc.CallOption) (*DeleteUnionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUnionResp)
	err := c.cc.Invoke(ctx, Org_DeleteUnion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgServer is the server API for Org service.
// All implementations must embed UnimplementedOrgServer
// for forward compatibility.
type OrgServer interface {
	InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error)
	GetDepartment(context.Context, *GetDepartmentReq) (*GetDepartmentResp, error)
	ListDepartments(context.Context, *ListDepartmentsReq) (*ListDepartmentsResp, error)
	AddDepartment(context.Context, *AddDepartmentReq) (*AddDepartmentResp, error)
	UpdateDepartment(context.Context, *UpdateDepartmentReq) (*UpdateDepartmentResp, error)
	DeleteDepartment(context.Context, *DeleteDepartmentReq) (*DeleteDepartmentResp, error)
	GetMerchant(context.Context, *GetMerchantReq) (*GetMerchantResp, error)
	ListMerchants(context.Context, *ListMerchantsReq) (*ListMerchantsResp, error)
	AddMerchant(context.Context, *AddMerchantReq) (*AddMerchantResp, error)
	UpdateMerchant(context.Context, *UpdateMerchantReq) (*UpdateMerchantResp, error)
	DeleteMerchant(context.Context, *DeleteMerchantReq) (*DeleteMerchantResp, error)
	GetUnion(context.Context, *GetUnionReq) (*GetUnionResp, error)
	ListUnions(context.Context, *ListUnionsReq) (*ListUnionsResp, error)
	AddUnion(context.Context, *AddUnionReq) (*AddUnionResp, error)
	UpdateUnion(context.Context, *UpdateUnionReq) (*UpdateUnionResp, error)
	DeleteUnion(context.Context, *DeleteUnionReq) (*DeleteUnionResp, error)
	mustEmbedUnimplementedOrgServer()
}

// UnimplementedOrgServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrgServer struct{}

func (UnimplementedOrgServer) InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDB not implemented")
}
func (UnimplementedOrgServer) GetDepartment(context.Context, *GetDepartmentReq) (*GetDepartmentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartment not implemented")
}
func (UnimplementedOrgServer) ListDepartments(context.Context, *ListDepartmentsReq) (*ListDepartmentsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepartments not implemented")
}
func (UnimplementedOrgServer) AddDepartment(context.Context, *AddDepartmentReq) (*AddDepartmentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDepartment not implemented")
}
func (UnimplementedOrgServer) UpdateDepartment(context.Context, *UpdateDepartmentReq) (*UpdateDepartmentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDepartment not implemented")
}
func (UnimplementedOrgServer) DeleteDepartment(context.Context, *DeleteDepartmentReq) (*DeleteDepartmentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDepartment not implemented")
}
func (UnimplementedOrgServer) GetMerchant(context.Context, *GetMerchantReq) (*GetMerchantResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchant not implemented")
}
func (UnimplementedOrgServer) ListMerchants(context.Context, *ListMerchantsReq) (*ListMerchantsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMerchants not implemented")
}
func (UnimplementedOrgServer) AddMerchant(context.Context, *AddMerchantReq) (*AddMerchantResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMerchant not implemented")
}
func (UnimplementedOrgServer) UpdateMerchant(context.Context, *UpdateMerchantReq) (*UpdateMerchantResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMerchant not implemented")
}
func (UnimplementedOrgServer) DeleteMerchant(context.Context, *DeleteMerchantReq) (*DeleteMerchantResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMerchant not implemented")
}
func (UnimplementedOrgServer) GetUnion(context.Context, *GetUnionReq) (*GetUnionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnion not implemented")
}
func (UnimplementedOrgServer) ListUnions(context.Context, *ListUnionsReq) (*ListUnionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUnions not implemented")
}
func (UnimplementedOrgServer) AddUnion(context.Context, *AddUnionReq) (*AddUnionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUnion not implemented")
}
func (UnimplementedOrgServer) UpdateUnion(context.Context, *UpdateUnionReq) (*UpdateUnionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUnion not implemented")
}
func (UnimplementedOrgServer) DeleteUnion(context.Context, *DeleteUnionReq) (*DeleteUnionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUnion not implemented")
}
func (UnimplementedOrgServer) mustEmbedUnimplementedOrgServer() {}
func (UnimplementedOrgServer) testEmbeddedByValue()             {}

// UnsafeOrgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgServer will
// result in compilation errors.
type UnsafeOrgServer interface {
	mustEmbedUnimplementedOrgServer()
}

func RegisterOrgServer(s grpc.ServiceRegistrar, srv OrgServer) {
	// If the following call pancis, it indicates UnimplementedOrgServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Org_ServiceDesc, srv)
}

func _Org_InitDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).InitDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_InitDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).InitDB(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_GetDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepartmentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).GetDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_GetDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).GetDepartment(ctx, req.(*GetDepartmentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_ListDepartments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDepartmentsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).ListDepartments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_ListDepartments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).ListDepartments(ctx, req.(*ListDepartmentsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_AddDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDepartmentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).AddDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_AddDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).AddDepartment(ctx, req.(*AddDepartmentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_UpdateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDepartmentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).UpdateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_UpdateDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).UpdateDepartment(ctx, req.(*UpdateDepartmentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_DeleteDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDepartmentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).DeleteDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_DeleteDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).DeleteDepartment(ctx, req.(*DeleteDepartmentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_GetMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).GetMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_GetMerchant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).GetMerchant(ctx, req.(*GetMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_ListMerchants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMerchantsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).ListMerchants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_ListMerchants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).ListMerchants(ctx, req.(*ListMerchantsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_AddMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).AddMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_AddMerchant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).AddMerchant(ctx, req.(*AddMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_UpdateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).UpdateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_UpdateMerchant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).UpdateMerchant(ctx, req.(*UpdateMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_DeleteMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).DeleteMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_DeleteMerchant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).DeleteMerchant(ctx, req.(*DeleteMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_GetUnion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).GetUnion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_GetUnion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).GetUnion(ctx, req.(*GetUnionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_ListUnions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUnionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).ListUnions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_ListUnions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).ListUnions(ctx, req.(*ListUnionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_AddUnion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUnionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).AddUnion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_AddUnion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).AddUnion(ctx, req.(*AddUnionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_UpdateUnion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUnionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).UpdateUnion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_UpdateUnion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).UpdateUnion(ctx, req.(*UpdateUnionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Org_DeleteUnion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUnionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServer).DeleteUnion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Org_DeleteUnion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServer).DeleteUnion(ctx, req.(*DeleteUnionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Org_ServiceDesc is the grpc.ServiceDesc for Org service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Org_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.biz.org.Org",
	HandlerType: (*OrgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDB",
			Handler:    _Org_InitDB_Handler,
		},
		{
			MethodName: "GetDepartment",
			Handler:    _Org_GetDepartment_Handler,
		},
		{
			MethodName: "ListDepartments",
			Handler:    _Org_ListDepartments_Handler,
		},
		{
			MethodName: "AddDepartment",
			Handler:    _Org_AddDepartment_Handler,
		},
		{
			MethodName: "UpdateDepartment",
			Handler:    _Org_UpdateDepartment_Handler,
		},
		{
			MethodName: "DeleteDepartment",
			Handler:    _Org_DeleteDepartment_Handler,
		},
		{
			MethodName: "GetMerchant",
			Handler:    _Org_GetMerchant_Handler,
		},
		{
			MethodName: "ListMerchants",
			Handler:    _Org_ListMerchants_Handler,
		},
		{
			MethodName: "AddMerchant",
			Handler:    _Org_AddMerchant_Handler,
		},
		{
			MethodName: "UpdateMerchant",
			Handler:    _Org_UpdateMerchant_Handler,
		},
		{
			MethodName: "DeleteMerchant",
			Handler:    _Org_DeleteMerchant_Handler,
		},
		{
			MethodName: "GetUnion",
			Handler:    _Org_GetUnion_Handler,
		},
		{
			MethodName: "ListUnions",
			Handler:    _Org_ListUnions_Handler,
		},
		{
			MethodName: "AddUnion",
			Handler:    _Org_AddUnion_Handler,
		},
		{
			MethodName: "UpdateUnion",
			Handler:    _Org_UpdateUnion_Handler,
		},
		{
			MethodName: "DeleteUnion",
			Handler:    _Org_DeleteUnion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.biz.org/org.proto",
}
