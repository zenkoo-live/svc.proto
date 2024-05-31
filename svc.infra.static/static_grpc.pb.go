// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StaticClient is the client API for Static service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaticClient interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error)
	UploadAvatar(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
	UploadCover(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
	UploadVideo(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
	UploadImage(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error)
}

type staticClient struct {
	cc grpc.ClientConnInterface
}

func NewStaticClient(cc grpc.ClientConnInterface) StaticClient {
	return &staticClient{cc}
}

func (c *staticClient) InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error) {
	out := new(InitDBResp)
	err := c.cc.Invoke(ctx, "/svc.infra.static.Static/InitDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadAvatar(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, "/svc.infra.static.Static/UploadAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadCover(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, "/svc.infra.static.Static/UploadCover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadVideo(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, "/svc.infra.static.Static/UploadVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticClient) UploadImage(ctx context.Context, in *UploadRequestMessage, opts ...grpc.CallOption) (*UploadResponseMessage, error) {
	out := new(UploadResponseMessage)
	err := c.cc.Invoke(ctx, "/svc.infra.static.Static/UploadImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaticServer is the server API for Static service.
// All implementations must embed UnimplementedStaticServer
// for forward compatibility
type StaticServer interface {
	InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error)
	UploadAvatar(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	UploadCover(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	UploadVideo(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	UploadImage(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error)
	mustEmbedUnimplementedStaticServer()
}

// UnimplementedStaticServer must be embedded to have forward compatible implementations.
type UnimplementedStaticServer struct {
}

func (UnimplementedStaticServer) InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDB not implemented")
}
func (UnimplementedStaticServer) UploadAvatar(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedStaticServer) UploadCover(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCover not implemented")
}
func (UnimplementedStaticServer) UploadVideo(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedStaticServer) UploadImage(context.Context, *UploadRequestMessage) (*UploadResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
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
		FullMethod: "/svc.infra.static.Static/InitDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).InitDB(ctx, req.(*emptypb.Empty))
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
		FullMethod: "/svc.infra.static.Static/UploadAvatar",
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
		FullMethod: "/svc.infra.static.Static/UploadCover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).UploadCover(ctx, req.(*UploadRequestMessage))
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
		FullMethod: "/svc.infra.static.Static/UploadVideo",
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
		FullMethod: "/svc.infra.static.Static/UploadImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).UploadImage(ctx, req.(*UploadRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "UploadAvatar",
			Handler:    _Static_UploadAvatar_Handler,
		},
		{
			MethodName: "UploadCover",
			Handler:    _Static_UploadCover_Handler,
		},
		{
			MethodName: "UploadVideo",
			Handler:    _Static_UploadVideo_Handler,
		},
		{
			MethodName: "UploadImage",
			Handler:    _Static_UploadImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.infra.static/static.proto",
}
