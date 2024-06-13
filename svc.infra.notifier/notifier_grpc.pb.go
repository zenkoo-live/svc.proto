// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: svc.infra.notifier/notifier.proto

package notifier

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

const (
	Notifier_InitDB_FullMethodName              = "/svc.infra.notifier.Notifier/InitDB"
	Notifier_GetSmsChannelList_FullMethodName   = "/svc.infra.notifier.Notifier/GetSmsChannelList"
	Notifier_CreatedSmsChannel_FullMethodName   = "/svc.infra.notifier.Notifier/CreatedSmsChannel"
	Notifier_UpdatedSmsChannel_FullMethodName   = "/svc.infra.notifier.Notifier/UpdatedSmsChannel"
	Notifier_DeletedSmsChannel_FullMethodName   = "/svc.infra.notifier.Notifier/DeletedSmsChannel"
	Notifier_GetSmsTemplateList_FullMethodName  = "/svc.infra.notifier.Notifier/GetSmsTemplateList"
	Notifier_CreatedSmsTemplate_FullMethodName  = "/svc.infra.notifier.Notifier/CreatedSmsTemplate"
	Notifier_UpdatedSmsTemplate_FullMethodName  = "/svc.infra.notifier.Notifier/UpdatedSmsTemplate"
	Notifier_DeletedSmsTemplate_FullMethodName  = "/svc.infra.notifier.Notifier/DeletedSmsTemplate"
	Notifier_GetSmsBizLogList_FullMethodName    = "/svc.infra.notifier.Notifier/GetSmsBizLogList"
	Notifier_CreatedSmsBizLog_FullMethodName    = "/svc.infra.notifier.Notifier/CreatedSmsBizLog"
	Notifier_CreatedSmsSend_FullMethodName      = "/svc.infra.notifier.Notifier/CreatedSmsSend"
	Notifier_CreatedSmsVerify_FullMethodName    = "/svc.infra.notifier.Notifier/CreatedSmsVerify"
	Notifier_GetCloudSmsTemplate_FullMethodName = "/svc.infra.notifier.Notifier/GetCloudSmsTemplate"
	Notifier_CreatedSmsCodeBind_FullMethodName  = "/svc.infra.notifier.Notifier/CreatedSmsCodeBind"
	Notifier_GetCloudSmsSign_FullMethodName     = "/svc.infra.notifier.Notifier/GetCloudSmsSign"
)

// NotifierClient is the client API for Notifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifierClient interface {
	InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error)
	// sms channel
	GetSmsChannelList(ctx context.Context, in *SmsChannelListRequest, opts ...grpc.CallOption) (*SmsChannelListResponse, error)
	CreatedSmsChannel(ctx context.Context, in *CreatedSmsChannelRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdatedSmsChannel(ctx context.Context, in *UpdatedSmsChannelRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	DeletedSmsChannel(ctx context.Context, in *DeletedSmsChannelRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// sms template
	GetSmsTemplateList(ctx context.Context, in *SmsTemplateListRequest, opts ...grpc.CallOption) (*SmsTemplateListResponse, error)
	CreatedSmsTemplate(ctx context.Context, in *CreatedSmsTemplateRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdatedSmsTemplate(ctx context.Context, in *UpdatedSmsTemplateRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	DeletedSmsTemplate(ctx context.Context, in *DeletedSmsTemplateRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// sms biz send
	GetSmsBizLogList(ctx context.Context, in *SmsBizSendLogListRequest, opts ...grpc.CallOption) (*SmsBizSendLogListResponse, error)
	CreatedSmsBizLog(ctx context.Context, in *CreatedSmsBizSendLogRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// send sms operation
	CreatedSmsSend(ctx context.Context, in *CreatedSmsSendRequest, opts ...grpc.CallOption) (*CreatedSmsSendResponse, error)
	CreatedSmsVerify(ctx context.Context, in *CreatedSmsVerifyRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	GetCloudSmsTemplate(ctx context.Context, in *GetCloudSmsTemplateRequest, opts ...grpc.CallOption) (*GetCloudSmsTemplateResponse, error)
	CreatedSmsCodeBind(ctx context.Context, in *CreatedSmsCodeBindRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// 获取签名列表
	GetCloudSmsSign(ctx context.Context, in *GetCloudSmsSignRequest, opts ...grpc.CallOption) (*GetCloudSmsSignResponse, error)
}

type notifierClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifierClient(cc grpc.ClientConnInterface) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) InitDB(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InitDBResp, error) {
	out := new(InitDBResp)
	err := c.cc.Invoke(ctx, Notifier_InitDB_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) GetSmsChannelList(ctx context.Context, in *SmsChannelListRequest, opts ...grpc.CallOption) (*SmsChannelListResponse, error) {
	out := new(SmsChannelListResponse)
	err := c.cc.Invoke(ctx, Notifier_GetSmsChannelList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) CreatedSmsChannel(ctx context.Context, in *CreatedSmsChannelRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Notifier_CreatedSmsChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) UpdatedSmsChannel(ctx context.Context, in *UpdatedSmsChannelRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Notifier_UpdatedSmsChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) DeletedSmsChannel(ctx context.Context, in *DeletedSmsChannelRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Notifier_DeletedSmsChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) GetSmsTemplateList(ctx context.Context, in *SmsTemplateListRequest, opts ...grpc.CallOption) (*SmsTemplateListResponse, error) {
	out := new(SmsTemplateListResponse)
	err := c.cc.Invoke(ctx, Notifier_GetSmsTemplateList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) CreatedSmsTemplate(ctx context.Context, in *CreatedSmsTemplateRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Notifier_CreatedSmsTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) UpdatedSmsTemplate(ctx context.Context, in *UpdatedSmsTemplateRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Notifier_UpdatedSmsTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) DeletedSmsTemplate(ctx context.Context, in *DeletedSmsTemplateRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Notifier_DeletedSmsTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) GetSmsBizLogList(ctx context.Context, in *SmsBizSendLogListRequest, opts ...grpc.CallOption) (*SmsBizSendLogListResponse, error) {
	out := new(SmsBizSendLogListResponse)
	err := c.cc.Invoke(ctx, Notifier_GetSmsBizLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) CreatedSmsBizLog(ctx context.Context, in *CreatedSmsBizSendLogRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Notifier_CreatedSmsBizLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) CreatedSmsSend(ctx context.Context, in *CreatedSmsSendRequest, opts ...grpc.CallOption) (*CreatedSmsSendResponse, error) {
	out := new(CreatedSmsSendResponse)
	err := c.cc.Invoke(ctx, Notifier_CreatedSmsSend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) CreatedSmsVerify(ctx context.Context, in *CreatedSmsVerifyRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Notifier_CreatedSmsVerify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) GetCloudSmsTemplate(ctx context.Context, in *GetCloudSmsTemplateRequest, opts ...grpc.CallOption) (*GetCloudSmsTemplateResponse, error) {
	out := new(GetCloudSmsTemplateResponse)
	err := c.cc.Invoke(ctx, Notifier_GetCloudSmsTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) CreatedSmsCodeBind(ctx context.Context, in *CreatedSmsCodeBindRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Notifier_CreatedSmsCodeBind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) GetCloudSmsSign(ctx context.Context, in *GetCloudSmsSignRequest, opts ...grpc.CallOption) (*GetCloudSmsSignResponse, error) {
	out := new(GetCloudSmsSignResponse)
	err := c.cc.Invoke(ctx, Notifier_GetCloudSmsSign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifierServer is the server API for Notifier service.
// All implementations must embed UnimplementedNotifierServer
// for forward compatibility
type NotifierServer interface {
	InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error)
	// sms channel
	GetSmsChannelList(context.Context, *SmsChannelListRequest) (*SmsChannelListResponse, error)
	CreatedSmsChannel(context.Context, *CreatedSmsChannelRequest) (*CommonResponse, error)
	UpdatedSmsChannel(context.Context, *UpdatedSmsChannelRequest) (*CommonResponse, error)
	DeletedSmsChannel(context.Context, *DeletedSmsChannelRequest) (*CommonResponse, error)
	// sms template
	GetSmsTemplateList(context.Context, *SmsTemplateListRequest) (*SmsTemplateListResponse, error)
	CreatedSmsTemplate(context.Context, *CreatedSmsTemplateRequest) (*CommonResponse, error)
	UpdatedSmsTemplate(context.Context, *UpdatedSmsTemplateRequest) (*CommonResponse, error)
	DeletedSmsTemplate(context.Context, *DeletedSmsTemplateRequest) (*CommonResponse, error)
	// sms biz send
	GetSmsBizLogList(context.Context, *SmsBizSendLogListRequest) (*SmsBizSendLogListResponse, error)
	CreatedSmsBizLog(context.Context, *CreatedSmsBizSendLogRequest) (*CommonResponse, error)
	// send sms operation
	CreatedSmsSend(context.Context, *CreatedSmsSendRequest) (*CreatedSmsSendResponse, error)
	CreatedSmsVerify(context.Context, *CreatedSmsVerifyRequest) (*CommonResponse, error)
	GetCloudSmsTemplate(context.Context, *GetCloudSmsTemplateRequest) (*GetCloudSmsTemplateResponse, error)
	CreatedSmsCodeBind(context.Context, *CreatedSmsCodeBindRequest) (*CommonResponse, error)
	// 获取签名列表
	GetCloudSmsSign(context.Context, *GetCloudSmsSignRequest) (*GetCloudSmsSignResponse, error)
	mustEmbedUnimplementedNotifierServer()
}

// UnimplementedNotifierServer must be embedded to have forward compatible implementations.
type UnimplementedNotifierServer struct {
}

func (UnimplementedNotifierServer) InitDB(context.Context, *emptypb.Empty) (*InitDBResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDB not implemented")
}
func (UnimplementedNotifierServer) GetSmsChannelList(context.Context, *SmsChannelListRequest) (*SmsChannelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmsChannelList not implemented")
}
func (UnimplementedNotifierServer) CreatedSmsChannel(context.Context, *CreatedSmsChannelRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedSmsChannel not implemented")
}
func (UnimplementedNotifierServer) UpdatedSmsChannel(context.Context, *UpdatedSmsChannelRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatedSmsChannel not implemented")
}
func (UnimplementedNotifierServer) DeletedSmsChannel(context.Context, *DeletedSmsChannelRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletedSmsChannel not implemented")
}
func (UnimplementedNotifierServer) GetSmsTemplateList(context.Context, *SmsTemplateListRequest) (*SmsTemplateListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmsTemplateList not implemented")
}
func (UnimplementedNotifierServer) CreatedSmsTemplate(context.Context, *CreatedSmsTemplateRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedSmsTemplate not implemented")
}
func (UnimplementedNotifierServer) UpdatedSmsTemplate(context.Context, *UpdatedSmsTemplateRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatedSmsTemplate not implemented")
}
func (UnimplementedNotifierServer) DeletedSmsTemplate(context.Context, *DeletedSmsTemplateRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletedSmsTemplate not implemented")
}
func (UnimplementedNotifierServer) GetSmsBizLogList(context.Context, *SmsBizSendLogListRequest) (*SmsBizSendLogListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmsBizLogList not implemented")
}
func (UnimplementedNotifierServer) CreatedSmsBizLog(context.Context, *CreatedSmsBizSendLogRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedSmsBizLog not implemented")
}
func (UnimplementedNotifierServer) CreatedSmsSend(context.Context, *CreatedSmsSendRequest) (*CreatedSmsSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedSmsSend not implemented")
}
func (UnimplementedNotifierServer) CreatedSmsVerify(context.Context, *CreatedSmsVerifyRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedSmsVerify not implemented")
}
func (UnimplementedNotifierServer) GetCloudSmsTemplate(context.Context, *GetCloudSmsTemplateRequest) (*GetCloudSmsTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCloudSmsTemplate not implemented")
}
func (UnimplementedNotifierServer) CreatedSmsCodeBind(context.Context, *CreatedSmsCodeBindRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedSmsCodeBind not implemented")
}
func (UnimplementedNotifierServer) GetCloudSmsSign(context.Context, *GetCloudSmsSignRequest) (*GetCloudSmsSignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCloudSmsSign not implemented")
}
func (UnimplementedNotifierServer) mustEmbedUnimplementedNotifierServer() {}

// UnsafeNotifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifierServer will
// result in compilation errors.
type UnsafeNotifierServer interface {
	mustEmbedUnimplementedNotifierServer()
}

func RegisterNotifierServer(s grpc.ServiceRegistrar, srv NotifierServer) {
	s.RegisterService(&Notifier_ServiceDesc, srv)
}

func _Notifier_InitDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).InitDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_InitDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).InitDB(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_GetSmsChannelList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).GetSmsChannelList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_GetSmsChannelList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).GetSmsChannelList(ctx, req.(*SmsChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_CreatedSmsChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedSmsChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).CreatedSmsChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_CreatedSmsChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).CreatedSmsChannel(ctx, req.(*CreatedSmsChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_UpdatedSmsChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatedSmsChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).UpdatedSmsChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_UpdatedSmsChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).UpdatedSmsChannel(ctx, req.(*UpdatedSmsChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_DeletedSmsChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletedSmsChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).DeletedSmsChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_DeletedSmsChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).DeletedSmsChannel(ctx, req.(*DeletedSmsChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_GetSmsTemplateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsTemplateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).GetSmsTemplateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_GetSmsTemplateList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).GetSmsTemplateList(ctx, req.(*SmsTemplateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_CreatedSmsTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedSmsTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).CreatedSmsTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_CreatedSmsTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).CreatedSmsTemplate(ctx, req.(*CreatedSmsTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_UpdatedSmsTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatedSmsTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).UpdatedSmsTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_UpdatedSmsTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).UpdatedSmsTemplate(ctx, req.(*UpdatedSmsTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_DeletedSmsTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletedSmsTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).DeletedSmsTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_DeletedSmsTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).DeletedSmsTemplate(ctx, req.(*DeletedSmsTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_GetSmsBizLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsBizSendLogListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).GetSmsBizLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_GetSmsBizLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).GetSmsBizLogList(ctx, req.(*SmsBizSendLogListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_CreatedSmsBizLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedSmsBizSendLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).CreatedSmsBizLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_CreatedSmsBizLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).CreatedSmsBizLog(ctx, req.(*CreatedSmsBizSendLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_CreatedSmsSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedSmsSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).CreatedSmsSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_CreatedSmsSend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).CreatedSmsSend(ctx, req.(*CreatedSmsSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_CreatedSmsVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedSmsVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).CreatedSmsVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_CreatedSmsVerify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).CreatedSmsVerify(ctx, req.(*CreatedSmsVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_GetCloudSmsTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudSmsTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).GetCloudSmsTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_GetCloudSmsTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).GetCloudSmsTemplate(ctx, req.(*GetCloudSmsTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_CreatedSmsCodeBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedSmsCodeBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).CreatedSmsCodeBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_CreatedSmsCodeBind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).CreatedSmsCodeBind(ctx, req.(*CreatedSmsCodeBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_GetCloudSmsSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudSmsSignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).GetCloudSmsSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifier_GetCloudSmsSign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).GetCloudSmsSign(ctx, req.(*GetCloudSmsSignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notifier_ServiceDesc is the grpc.ServiceDesc for Notifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc.infra.notifier.Notifier",
	HandlerType: (*NotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDB",
			Handler:    _Notifier_InitDB_Handler,
		},
		{
			MethodName: "GetSmsChannelList",
			Handler:    _Notifier_GetSmsChannelList_Handler,
		},
		{
			MethodName: "CreatedSmsChannel",
			Handler:    _Notifier_CreatedSmsChannel_Handler,
		},
		{
			MethodName: "UpdatedSmsChannel",
			Handler:    _Notifier_UpdatedSmsChannel_Handler,
		},
		{
			MethodName: "DeletedSmsChannel",
			Handler:    _Notifier_DeletedSmsChannel_Handler,
		},
		{
			MethodName: "GetSmsTemplateList",
			Handler:    _Notifier_GetSmsTemplateList_Handler,
		},
		{
			MethodName: "CreatedSmsTemplate",
			Handler:    _Notifier_CreatedSmsTemplate_Handler,
		},
		{
			MethodName: "UpdatedSmsTemplate",
			Handler:    _Notifier_UpdatedSmsTemplate_Handler,
		},
		{
			MethodName: "DeletedSmsTemplate",
			Handler:    _Notifier_DeletedSmsTemplate_Handler,
		},
		{
			MethodName: "GetSmsBizLogList",
			Handler:    _Notifier_GetSmsBizLogList_Handler,
		},
		{
			MethodName: "CreatedSmsBizLog",
			Handler:    _Notifier_CreatedSmsBizLog_Handler,
		},
		{
			MethodName: "CreatedSmsSend",
			Handler:    _Notifier_CreatedSmsSend_Handler,
		},
		{
			MethodName: "CreatedSmsVerify",
			Handler:    _Notifier_CreatedSmsVerify_Handler,
		},
		{
			MethodName: "GetCloudSmsTemplate",
			Handler:    _Notifier_GetCloudSmsTemplate_Handler,
		},
		{
			MethodName: "CreatedSmsCodeBind",
			Handler:    _Notifier_CreatedSmsCodeBind_Handler,
		},
		{
			MethodName: "GetCloudSmsSign",
			Handler:    _Notifier_GetCloudSmsSign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.infra.notifier/notifier.proto",
}
