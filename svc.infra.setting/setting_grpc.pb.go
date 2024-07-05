// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: svc.infra.setting/setting.proto

package setting

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Setting_InitDB_FullMethodName              = "/svc.infra.setting.Setting/InitDB"
	Setting_GetConfiguration_FullMethodName    = "/svc.infra.setting.Setting/GetConfiguration"
	Setting_AddConfiguration_FullMethodName    = "/svc.infra.setting.Setting/AddConfiguration"
	Setting_UpdateConfiguration_FullMethodName = "/svc.infra.setting.Setting/UpdateConfiguration"
	Setting_DeleteConfiguration_FullMethodName = "/svc.infra.setting.Setting/DeleteConfiguration"
	Setting_Greeting_FullMethodName            = "/svc.infra.setting.Setting/Greeting"
)

// SettingClient is the client API for Setting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingClient interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error)
	GetConfiguration(ctx context.Context, in *GetConfigurationReq, opts ...grpc.CallOption) (*GetConfigurationResp, error)
	AddConfiguration(ctx context.Context, in *AddConfigurationReq, opts ...grpc.CallOption) (*AddConfigurationResp, error)
	UpdateConfiguration(ctx context.Context, in *UpdateConfigurationReq, opts ...grpc.CallOption) (*UpdateConfigurationResp, error)
	DeleteConfiguration(ctx context.Context, in *DeleteConfigurationReq, opts ...grpc.CallOption) (*DeleteConfigurationResp, error)
	Greeting(ctx context.Context, in *SettingGreetingReq, opts ...grpc.CallOption) (*SettingGreetingResp, error)
}

type settingClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingClient(cc grpc.ClientConnInterface) SettingClient {
	return &settingClient{cc}
}

func (c *settingClient) InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitDBResp)
	err := c.cc.Invoke(ctx, Setting_InitDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) GetConfiguration(ctx context.Context, in *GetConfigurationReq, opts ...grpc.CallOption) (*GetConfigurationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConfigurationResp)
	err := c.cc.Invoke(ctx, Setting_GetConfiguration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) AddConfiguration(ctx context.Context, in *AddConfigurationReq, opts ...grpc.CallOption) (*AddConfigurationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddConfigurationResp)
	err := c.cc.Invoke(ctx, Setting_AddConfiguration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) UpdateConfiguration(ctx context.Context, in *UpdateConfigurationReq, opts ...grpc.CallOption) (*UpdateConfigurationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateConfigurationResp)
	err := c.cc.Invoke(ctx, Setting_UpdateConfiguration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) DeleteConfiguration(ctx context.Context, in *DeleteConfigurationReq, opts ...grpc.CallOption) (*DeleteConfigurationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteConfigurationResp)
	err := c.cc.Invoke(ctx, Setting_DeleteConfiguration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingClient) Greeting(ctx context.Context, in *SettingGreetingReq, opts ...grpc.CallOption) (*SettingGreetingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SettingGreetingResp)
	err := c.cc.Invoke(ctx, Setting_Greeting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingServer is the server API for Setting service.
// All implementations must embed UnimplementedSettingServer
// for forward compatibility
type SettingServer interface {
	InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error)
	GetConfiguration(context.Context, *GetConfigurationReq) (*GetConfigurationResp, error)
	AddConfiguration(context.Context, *AddConfigurationReq) (*AddConfigurationResp, error)
	UpdateConfiguration(context.Context, *UpdateConfigurationReq) (*UpdateConfigurationResp, error)
	DeleteConfiguration(context.Context, *DeleteConfigurationReq) (*DeleteConfigurationResp, error)
	Greeting(context.Context, *SettingGreetingReq) (*SettingGreetingResp, error)
	mustEmbedUnimplementedSettingServer()
}

// UnimplementedSettingServer must be embedded to have forward compatible implementations.
type UnimplementedSettingServer struct {
}

func (UnimplementedSettingServer) InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDB not implemented")
}
func (UnimplementedSettingServer) GetConfiguration(context.Context, *GetConfigurationReq) (*GetConfigurationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedSettingServer) AddConfiguration(context.Context, *AddConfigurationReq) (*AddConfigurationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConfiguration not implemented")
}
func (UnimplementedSettingServer) UpdateConfiguration(context.Context, *UpdateConfigurationReq) (*UpdateConfigurationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfiguration not implemented")
}
func (UnimplementedSettingServer) DeleteConfiguration(context.Context, *DeleteConfigurationReq) (*DeleteConfigurationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfiguration not implemented")
}
func (UnimplementedSettingServer) Greeting(context.Context, *SettingGreetingReq) (*SettingGreetingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}
func (UnimplementedSettingServer) mustEmbedUnimplementedSettingServer() {}

// UnsafeSettingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingServer will
// result in compilation errors.
type UnsafeSettingServer interface {
	mustEmbedUnimplementedSettingServer()
}

func RegisterSettingServer(s grpc.ServiceRegistrar, srv SettingServer) {
	s.RegisterService(&Setting_ServiceDesc, srv)
}

func _Setting_InitDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).InitDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Setting_InitDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).InitDB(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Setting_GetConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).GetConfiguration(ctx, req.(*GetConfigurationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_AddConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddConfigurationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).AddConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Setting_AddConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).AddConfiguration(ctx, req.(*AddConfigurationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_UpdateConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigurationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).UpdateConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Setting_UpdateConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).UpdateConfiguration(ctx, req.(*UpdateConfigurationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_DeleteConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigurationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).DeleteConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Setting_DeleteConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).DeleteConfiguration(ctx, req.(*DeleteConfigurationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setting_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettingGreetingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Setting_Greeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).Greeting(ctx, req.(*SettingGreetingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Setting_ServiceDesc is the grpc.ServiceDesc for Setting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Setting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.infra.setting.Setting",
	HandlerType: (*SettingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDB",
			Handler:    _Setting_InitDB_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _Setting_GetConfiguration_Handler,
		},
		{
			MethodName: "AddConfiguration",
			Handler:    _Setting_AddConfiguration_Handler,
		},
		{
			MethodName: "UpdateConfiguration",
			Handler:    _Setting_UpdateConfiguration_Handler,
		},
		{
			MethodName: "DeleteConfiguration",
			Handler:    _Setting_DeleteConfiguration_Handler,
		},
		{
			MethodName: "Greeting",
			Handler:    _Setting_Greeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.infra.setting/setting.proto",
}
