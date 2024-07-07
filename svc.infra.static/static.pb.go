// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: svc.infra.static/static.proto

package static

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitDBResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InitDBResp) Reset() {
	*x = InitDBResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_static_static_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitDBResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDBResp) ProtoMessage() {}

func (x *InitDBResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitDBResp.ProtoReflect.Descriptor instead.
func (*InitDBResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{0}
}

func (x *InitDBResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type S3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppSecret string `protobuf:"bytes,2,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
	Region    string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Bucket    string `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Queue     string `protobuf:"bytes,5,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *S3) Reset() {
	*x = S3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_static_static_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3) ProtoMessage() {}

func (x *S3) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3.ProtoReflect.Descriptor instead.
func (*S3) Descriptor() ([]byte, []int) {
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{1}
}

func (x *S3) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *S3) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

func (x *S3) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *S3) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

type ConfigurationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S3         *S3      `protobuf:"bytes,1,opt,name=s3,proto3" json:"s3,omitempty"`                                     // s3配置
	Domains    []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`                           // 静态域名
	MerchantId string   `protobuf:"bytes,255,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"` // 商户id
}

func (x *ConfigurationMessage) Reset() {
	*x = ConfigurationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_static_static_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurationMessage) ProtoMessage() {}

func (x *ConfigurationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurationMessage.ProtoReflect.Descriptor instead.
func (*ConfigurationMessage) Descriptor() ([]byte, []int) {
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigurationMessage) GetS3() *S3 {
	if x != nil {
		return x.S3
	}
	return nil
}

func (x *ConfigurationMessage) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *ConfigurationMessage) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type ConfigurationResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ConfigurationResponseMessage) Reset() {
	*x = ConfigurationResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_static_static_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurationResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurationResponseMessage) ProtoMessage() {}

func (x *ConfigurationResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurationResponseMessage.ProtoReflect.Descriptor instead.
func (*ConfigurationResponseMessage) Descriptor() ([]byte, []int) {
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigurationResponseMessage) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type DomainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domains []string `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (x *DomainsResponse) Reset() {
	*x = DomainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_static_static_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainsResponse) ProtoMessage() {}

func (x *DomainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainsResponse.ProtoReflect.Descriptor instead.
func (*DomainsResponse) Descriptor() ([]byte, []int) {
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{4}
}

func (x *DomainsResponse) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

type UploadRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                               // 为空将会自动生成
	Binary     []byte `protobuf:"bytes,2,opt,name=binary,proto3" json:"binary,omitempty"`                           // 文件内容
	IsPreMode  bool   `protobuf:"varint,3,opt,name=is_pre_mode,json=isPreMode,proto3" json:"is_pre_mode,omitempty"` // 预写模式, 返回token客户端根据这个直接上传
	Bucket     string `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`                           // 自定义桶
	UserId     string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // 用户id
	MerchantId string `protobuf:"bytes,6,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"` // 商户id
}

func (x *UploadRequestMessage) Reset() {
	*x = UploadRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_static_static_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequestMessage) ProtoMessage() {}

func (x *UploadRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequestMessage.ProtoReflect.Descriptor instead.
func (*UploadRequestMessage) Descriptor() ([]byte, []int) {
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{5}
}

func (x *UploadRequestMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadRequestMessage) GetBinary() []byte {
	if x != nil {
		return x.Binary
	}
	return nil
}

func (x *UploadRequestMessage) GetIsPreMode() bool {
	if x != nil {
		return x.IsPreMode
	}
	return false
}

func (x *UploadRequestMessage) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *UploadRequestMessage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UploadRequestMessage) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type StreamRequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                               // 为空将会自动生成
	Bucket     string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`                           // 自定义桶
	UserId     string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // 用户id
	MerchantId string `protobuf:"bytes,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"` // 商户id
}

func (x *StreamRequestInfo) Reset() {
	*x = StreamRequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_static_static_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequestInfo) ProtoMessage() {}

func (x *StreamRequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequestInfo.ProtoReflect.Descriptor instead.
func (*StreamRequestInfo) Descriptor() ([]byte, []int) {
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{6}
}

func (x *StreamRequestInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamRequestInfo) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *StreamRequestInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *StreamRequestInfo) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type UploadStreamRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UploadStreamRequestMessage_Info
	//	*UploadStreamRequestMessage_Binary
	Data isUploadStreamRequestMessage_Data `protobuf_oneof:"data"`
}

func (x *UploadStreamRequestMessage) Reset() {
	*x = UploadStreamRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_static_static_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadStreamRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStreamRequestMessage) ProtoMessage() {}

func (x *UploadStreamRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStreamRequestMessage.ProtoReflect.Descriptor instead.
func (*UploadStreamRequestMessage) Descriptor() ([]byte, []int) {
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{7}
}

func (m *UploadStreamRequestMessage) GetData() isUploadStreamRequestMessage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadStreamRequestMessage) GetInfo() *StreamRequestInfo {
	if x, ok := x.GetData().(*UploadStreamRequestMessage_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadStreamRequestMessage) GetBinary() []byte {
	if x, ok := x.GetData().(*UploadStreamRequestMessage_Binary); ok {
		return x.Binary
	}
	return nil
}

type isUploadStreamRequestMessage_Data interface {
	isUploadStreamRequestMessage_Data()
}

type UploadStreamRequestMessage_Info struct {
	Info *StreamRequestInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadStreamRequestMessage_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*UploadStreamRequestMessage_Info) isUploadStreamRequestMessage_Data() {}

func (*UploadStreamRequestMessage_Binary) isUploadStreamRequestMessage_Data() {}

type UploadResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`               // 返回文件路径 ，如果是预写模式该值就是token 或者地址
	Domain      string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`           // 域
	Provider    string `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`       // oss 提供商
	Certificate string `protobuf:"bytes,4,opt,name=certificate,proto3" json:"certificate,omitempty"` // 预写凭证
	Bucket      string `protobuf:"bytes,5,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Region      string `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *UploadResponseMessage) Reset() {
	*x = UploadResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_static_static_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponseMessage) ProtoMessage() {}

func (x *UploadResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponseMessage.ProtoReflect.Descriptor instead.
func (*UploadResponseMessage) Descriptor() ([]byte, []int) {
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{8}
}

func (x *UploadResponseMessage) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UploadResponseMessage) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *UploadResponseMessage) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *UploadResponseMessage) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *UploadResponseMessage) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *UploadResponseMessage) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

var File_svc_infra_static_static_proto protoreflect.FileDescriptor

var file_svc_infra_static_static_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24,
	0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x02, 0x53, 0x33, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x78, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x24, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x53,
	0x33, 0x52, 0x02, 0x73, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0xff,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x36, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0b, 0x69,
	0x73, 0x5f, 0x70, 0x72, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x79, 0x0a,
	0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x1a, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xb1, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x32, 0xa4, 0x07, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x12, 0x3e, 0x0a, 0x06, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x44, 0x0a, 0x07, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x2e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x6b, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x42, 0x1b,
	0x5a, 0x19, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_svc_infra_static_static_proto_rawDescOnce sync.Once
	file_svc_infra_static_static_proto_rawDescData = file_svc_infra_static_static_proto_rawDesc
)

func file_svc_infra_static_static_proto_rawDescGZIP() []byte {
	file_svc_infra_static_static_proto_rawDescOnce.Do(func() {
		file_svc_infra_static_static_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_infra_static_static_proto_rawDescData)
	})
	return file_svc_infra_static_static_proto_rawDescData
}

var file_svc_infra_static_static_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_svc_infra_static_static_proto_goTypes = []any{
	(*InitDBResp)(nil),                   // 0: svc.infra.static.InitDBResp
	(*S3)(nil),                           // 1: svc.infra.static.S3
	(*ConfigurationMessage)(nil),         // 2: svc.infra.static.ConfigurationMessage
	(*ConfigurationResponseMessage)(nil), // 3: svc.infra.static.ConfigurationResponseMessage
	(*DomainsResponse)(nil),              // 4: svc.infra.static.DomainsResponse
	(*UploadRequestMessage)(nil),         // 5: svc.infra.static.UploadRequestMessage
	(*StreamRequestInfo)(nil),            // 6: svc.infra.static.StreamRequestInfo
	(*UploadStreamRequestMessage)(nil),   // 7: svc.infra.static.UploadStreamRequestMessage
	(*UploadResponseMessage)(nil),        // 8: svc.infra.static.UploadResponseMessage
	(*emptypb.Empty)(nil),                // 9: google.protobuf.Empty
}
var file_svc_infra_static_static_proto_depIdxs = []int32{
	1,  // 0: svc.infra.static.ConfigurationMessage.s3:type_name -> svc.infra.static.S3
	6,  // 1: svc.infra.static.UploadStreamRequestMessage.info:type_name -> svc.infra.static.StreamRequestInfo
	9,  // 2: svc.infra.static.Static.InitDB:input_type -> google.protobuf.Empty
	9,  // 3: svc.infra.static.Static.Domains:input_type -> google.protobuf.Empty
	2,  // 4: svc.infra.static.Static.Configuration:input_type -> svc.infra.static.ConfigurationMessage
	5,  // 5: svc.infra.static.Static.UploadAvatar:input_type -> svc.infra.static.UploadRequestMessage
	5,  // 6: svc.infra.static.Static.UploadCover:input_type -> svc.infra.static.UploadRequestMessage
	5,  // 7: svc.infra.static.Static.UploadAudio:input_type -> svc.infra.static.UploadRequestMessage
	5,  // 8: svc.infra.static.Static.UploadVideo:input_type -> svc.infra.static.UploadRequestMessage
	5,  // 9: svc.infra.static.Static.UploadImage:input_type -> svc.infra.static.UploadRequestMessage
	5,  // 10: svc.infra.static.Static.UploadFile:input_type -> svc.infra.static.UploadRequestMessage
	7,  // 11: svc.infra.static.Static.UploadStreamFile:input_type -> svc.infra.static.UploadStreamRequestMessage
	0,  // 12: svc.infra.static.Static.InitDB:output_type -> svc.infra.static.InitDBResp
	4,  // 13: svc.infra.static.Static.Domains:output_type -> svc.infra.static.DomainsResponse
	3,  // 14: svc.infra.static.Static.Configuration:output_type -> svc.infra.static.ConfigurationResponseMessage
	8,  // 15: svc.infra.static.Static.UploadAvatar:output_type -> svc.infra.static.UploadResponseMessage
	8,  // 16: svc.infra.static.Static.UploadCover:output_type -> svc.infra.static.UploadResponseMessage
	8,  // 17: svc.infra.static.Static.UploadAudio:output_type -> svc.infra.static.UploadResponseMessage
	8,  // 18: svc.infra.static.Static.UploadVideo:output_type -> svc.infra.static.UploadResponseMessage
	8,  // 19: svc.infra.static.Static.UploadImage:output_type -> svc.infra.static.UploadResponseMessage
	8,  // 20: svc.infra.static.Static.UploadFile:output_type -> svc.infra.static.UploadResponseMessage
	8,  // 21: svc.infra.static.Static.UploadStreamFile:output_type -> svc.infra.static.UploadResponseMessage
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_svc_infra_static_static_proto_init() }
func file_svc_infra_static_static_proto_init() {
	if File_svc_infra_static_static_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_infra_static_static_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InitDBResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_static_static_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*S3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_static_static_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ConfigurationMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_static_static_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ConfigurationResponseMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_static_static_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DomainsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_static_static_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UploadRequestMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_static_static_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*StreamRequestInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_static_static_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UploadStreamRequestMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_svc_infra_static_static_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UploadResponseMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_svc_infra_static_static_proto_msgTypes[7].OneofWrappers = []any{
		(*UploadStreamRequestMessage_Info)(nil),
		(*UploadStreamRequestMessage_Binary)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_infra_static_static_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_infra_static_static_proto_goTypes,
		DependencyIndexes: file_svc_infra_static_static_proto_depIdxs,
		MessageInfos:      file_svc_infra_static_static_proto_msgTypes,
	}.Build()
	File_svc_infra_static_static_proto = out.File
	file_svc_infra_static_static_proto_rawDesc = nil
	file_svc_infra_static_static_proto_goTypes = nil
	file_svc_infra_static_static_proto_depIdxs = nil
}
