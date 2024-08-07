// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: svc.infra.static/static.proto

package static

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
	Static_InitDB_FullMethodName           = "/svc.infra.static.Static/InitDB"
	Static_Domains_FullMethodName          = "/svc.infra.static.Static/Domains"
	Static_Configuration_FullMethodName    = "/svc.infra.static.Static/Configuration"
	Static_UploadAvatar_FullMethodName     = "/svc.infra.static.Static/UploadAvatar"
	Static_UploadCover_FullMethodName      = "/svc.infra.static.Static/UploadCover"
	Static_UploadAudio_FullMethodName      = "/svc.infra.static.Static/UploadAudio"
	Static_UploadVideo_FullMethodName      = "/svc.infra.static.Static/UploadVideo"
	Static_UploadImage_FullMethodName      = "/svc.infra.static.Static/UploadImage"
	Static_UploadFile_FullMethodName       = "/svc.infra.static.Static/UploadFile"
	Static_UploadStreamFile_FullMethodName = "/svc.infra.static.Static/UploadStreamFile"
)

// StaticClient is the client API for Static service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaticClient interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error)
	Domains(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DomainsResponse, error)
	Configuration(ctx context.Context, in *ConfigurationMessage, opts ...grpc.CallOption) (*ConfigurationResponseMessage, error)
	UploadAvatar(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
	UploadCover(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
	UploadAudio(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
	UploadVideo(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
	UploadImage(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
	UploadFile(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
	UploadStreamFile(ctx context.Context, opts ...grpc.CallOption) (Static_UploadStreamFileClient, error)
}

type staticClient struct {
	cc grpc.ClientConnInterface
}

func NewStaticClient(cc grpc.ClientConnInterface) StaticClient {
	return &staticClient{cc}
}

func (c *staticClient) InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitDBResp)
	err := c.cc.Invoke(ctx, Static_InitDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) Domains(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DomainsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DomainsResponse)
	err := c.cc.Invoke(ctx, Static_Domains_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) Configuration(ctx context.Context, in *ConfigurationMessage, opts ...grpc.CallOption) (*ConfigurationResponseMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfigurationResponseMessage)
	err := c.cc.Invoke(ctx, Static_Configuration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadAvatar(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, Static_UploadAvatar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadCover(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, Static_UploadCover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadAudio(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, Static_UploadAudio_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadVideo(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, Static_UploadVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadImage(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, Static_UploadImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadFile(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, Static_UploadFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadStreamFile(ctx context.Context, opts ...grpc.CallOption) (Static_UploadStreamFileClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Static_ServiceDesc.Streams[0], Static_UploadStreamFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &staticUploadStreamFileClient{ClientStream: stream}
	return x, nil
}

type Static_UploadStreamFileClient interface {
	Send(*UploadStreamRequestMessage) error
	CloseAndRecv() (*UploadResponseMessage, error)
	grpc.ClientStream
}

type staticUploadStreamFileClient struct {
	grpc.ClientStream
}

func (x *staticUploadStreamFileClient) Send(m *UploadStreamRequestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *staticUploadStreamFileClient) CloseAndRecv() (*UploadResponseMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponseMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StaticServer is the server API for Static service.
// All implementations must embed UnimplementedStaticServer
// for forward compatibility
type StaticServer interface {
	InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error)
	Domains(context.Context, *emptypb.Empty) (*DomainsResponse, error)
	Configuration(context.Context, *ConfigurationMessage) (*ConfigurationResponseMessage, error)
	UploadAvatar(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	UploadCover(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	UploadAudio(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	UploadVideo(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	UploadImage(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	UploadFile(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	UploadStreamFile(Static_UploadStreamFileServer) error
	mustEmbedUnimplementedStaticServer()
}

// UnimplementedStaticServer must be embedded to have forward compatible implementations.
type UnimplementedStaticServer struct {
}

func (UnimplementedStaticServer) InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDB not implemented")
}
func (UnimplementedStaticServer) Domains(context.Context, *emptypb.Empty) (*DomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Domains not implemented")
}
func (UnimplementedStaticServer) Configuration(context.Context, *ConfigurationMessage) (*ConfigurationResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configuration not implemented")
}
func (UnimplementedStaticServer) UploadAvatar(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedStaticServer) UploadCover(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCover not implemented")
}
func (UnimplementedStaticServer) UploadAudio(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAudio not implemented")
}
func (UnimplementedStaticServer) UploadVideo(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedStaticServer) UploadImage(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedStaticServer) UploadFile(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedStaticServer) UploadStreamFile(Static_UploadStreamFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadStreamFile not implemented")
}
func (UnimplementedStaticServer) mustEmbedUnimplementedStaticServer() {}

// UnsafeStaticServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaticServer will
// result in compilation errors.
type UnsafeStaticServer interface {
	mustEmbedUnimplementedStaticServer()
}

func RegisterStaticServer(s grpc.ServiceRegistrar, srv StaticServer) {
	s.RegisterService(&Static_ServiceDesc, srv)
}

func _Static_InitDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).InitDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_InitDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).InitDB(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_Domains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).Domains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_Domains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).Domains(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_Configuration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigurationMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).Configuration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_Configuration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).Configuration(ctx, req.(*ConfigurationMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_UploadAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).UploadAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_UploadAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).UploadAvatar(ctx, req.(*UploadRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_UploadCover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).UploadCover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_UploadCover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).UploadCover(ctx, req.(*UploadRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_UploadAudio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).UploadAudio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_UploadAudio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).UploadAudio(ctx, req.(*UploadRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_UploadVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).UploadVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_UploadVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).UploadVideo(ctx, req.(*UploadRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_UploadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).UploadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_UploadImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).UploadImage(ctx, req.(*UploadRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Static_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).UploadFile(ctx, req.(*UploadRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Static_UploadStreamFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StaticServer).UploadStreamFile(&staticUploadStreamFileServer{ServerStream: stream})
}

type Static_UploadStreamFileServer interface {
	SendAndClose(*UploadResponseMessage) error
	Recv() (*UploadStreamRequestMessage, error)
	grpc.ServerStream
}

type staticUploadStreamFileServer struct {
	grpc.ServerStream
}

func (x *staticUploadStreamFileServer) SendAndClose(m *UploadResponseMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *staticUploadStreamFileServer) Recv() (*UploadStreamRequestMessage, error) {
	m := new(UploadStreamRequestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Static_ServiceDesc is the grpc.ServiceDesc for Static service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Static_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.infra.static.Static",
	HandlerType: (*StaticServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDB",
			Handler:    _Static_InitDB_Handler,
		},
		{
			MethodName: "Domains",
			Handler:    _Static_Domains_Handler,
		},
		{
			MethodName: "Configuration",
			Handler:    _Static_Configuration_Handler,
		},
		{
			MethodName: "UploadAvatar",
			Handler:    _Static_UploadAvatar_Handler,
		},
		{
			MethodName: "UploadCover",
			Handler:    _Static_UploadCover_Handler,
		},
		{
			MethodName: "UploadAudio",
			Handler:    _Static_UploadAudio_Handler,
		},
		{
			MethodName: "UploadVideo",
			Handler:    _Static_UploadVideo_Handler,
		},
		{
			MethodName: "UploadImage",
			Handler:    _Static_UploadImage_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _Static_UploadFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadStreamFile",
			Handler:       _Static_UploadStreamFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "svc.infra.static/static.proto",
}
