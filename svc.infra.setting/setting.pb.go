// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: svc.infra.setting/setting.proto

package setting

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
		mi := &file_svc_infra_setting_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitDBResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDBResp) ProtoMessage() {}

func (x *InitDBResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[0]
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
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{0}
}

func (x *InitDBResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type SettingGreetingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SettingGreetingReq) Reset() {
	*x = SettingGreetingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingGreetingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingGreetingReq) ProtoMessage() {}

func (x *SettingGreetingReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingGreetingReq.ProtoReflect.Descriptor instead.
func (*SettingGreetingReq) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{1}
}

func (x *SettingGreetingReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SettingGreetingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SettingGreetingResp) Reset() {
	*x = SettingGreetingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingGreetingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingGreetingResp) ProtoMessage() {}

func (x *SettingGreetingResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingGreetingResp.ProtoReflect.Descriptor instead.
func (*SettingGreetingResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{2}
}

func (x *SettingGreetingResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`                                //配置键
	Value     string                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`                            // 配置json数据字符串值
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,255,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 创建时间
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{3}
}

func (x *Configuration) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Configuration) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Configuration) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetConfigurationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetConfigurationReq) Reset() {
	*x = GetConfigurationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigurationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationReq) ProtoMessage() {}

func (x *GetConfigurationReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationReq.ProtoReflect.Descriptor instead.
func (*GetConfigurationReq) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{4}
}

func (x *GetConfigurationReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetConfigurationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetConfigurationResp) Reset() {
	*x = GetConfigurationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigurationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationResp) ProtoMessage() {}

func (x *GetConfigurationResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationResp.ProtoReflect.Descriptor instead.
func (*GetConfigurationResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{5}
}

func (x *GetConfigurationResp) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type AddConfigurationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Force bool   `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *AddConfigurationReq) Reset() {
	*x = AddConfigurationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddConfigurationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddConfigurationReq) ProtoMessage() {}

func (x *AddConfigurationReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddConfigurationReq.ProtoReflect.Descriptor instead.
func (*AddConfigurationReq) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{6}
}

func (x *AddConfigurationReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AddConfigurationReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AddConfigurationReq) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type AddConfigurationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configuration *Configuration `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Result        bool           `protobuf:"varint,255,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddConfigurationResp) Reset() {
	*x = AddConfigurationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddConfigurationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddConfigurationResp) ProtoMessage() {}

func (x *AddConfigurationResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddConfigurationResp.ProtoReflect.Descriptor instead.
func (*AddConfigurationResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{7}
}

func (x *AddConfigurationResp) GetConfiguration() *Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *AddConfigurationResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type UpdateConfigurationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateConfigurationReq) Reset() {
	*x = UpdateConfigurationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigurationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigurationReq) ProtoMessage() {}

func (x *UpdateConfigurationReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigurationReq.ProtoReflect.Descriptor instead.
func (*UpdateConfigurationReq) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateConfigurationReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateConfigurationReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpdateConfigurationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated int64 `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *UpdateConfigurationResp) Reset() {
	*x = UpdateConfigurationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigurationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigurationResp) ProtoMessage() {}

func (x *UpdateConfigurationResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigurationResp.ProtoReflect.Descriptor instead.
func (*UpdateConfigurationResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateConfigurationResp) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type DeleteConfigurationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition *Configuration `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *DeleteConfigurationReq) Reset() {
	*x = DeleteConfigurationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfigurationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfigurationReq) ProtoMessage() {}

func (x *DeleteConfigurationReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfigurationReq.ProtoReflect.Descriptor instead.
func (*DeleteConfigurationReq) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteConfigurationReq) GetCondition() *Configuration {
	if x != nil {
		return x.Condition
	}
	return nil
}

type DeleteConfigurationResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted int64 `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteConfigurationResp) Reset() {
	*x = DeleteConfigurationResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_setting_setting_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfigurationResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfigurationResp) ProtoMessage() {}

func (x *DeleteConfigurationResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_setting_setting_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfigurationResp.ProtoReflect.Descriptor instead.
func (*DeleteConfigurationResp) Descriptor() ([]byte, []int) {
	return file_svc_infra_setting_setting_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteConfigurationResp) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

var File_svc_infra_setting_setting_proto protoreflect.FileDescriptor

var file_svc_infra_setting_setting_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x28, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x73, 0x0a,
	0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2c, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x53, 0x0a, 0x13, 0x41, 0x64, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x77,
	0x0a, 0x14, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x40, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x33, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x58,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x3e, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xcd, 0x04,
	0x0a, 0x07, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x0a, 0x06, 0x49, 0x6e, 0x69,
	0x74, 0x44, 0x42, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x12, 0x63, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x63, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x6c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x6c, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x5b, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x1d, 0x5a,
	0x1b, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x3b, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_infra_setting_setting_proto_rawDescOnce sync.Once
	file_svc_infra_setting_setting_proto_rawDescData = file_svc_infra_setting_setting_proto_rawDesc
)

func file_svc_infra_setting_setting_proto_rawDescGZIP() []byte {
	file_svc_infra_setting_setting_proto_rawDescOnce.Do(func() {
		file_svc_infra_setting_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_infra_setting_setting_proto_rawDescData)
	})
	return file_svc_infra_setting_setting_proto_rawDescData
}

var file_svc_infra_setting_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_svc_infra_setting_setting_proto_goTypes = []any{
	(*InitDBResp)(nil),              // 0: svc.infra.setting.InitDBResp
	(*SettingGreetingReq)(nil),      // 1: svc.infra.setting.SettingGreetingReq
	(*SettingGreetingResp)(nil),     // 2: svc.infra.setting.SettingGreetingResp
	(*Configuration)(nil),           // 3: svc.infra.setting.Configuration
	(*GetConfigurationReq)(nil),     // 4: svc.infra.setting.GetConfigurationReq
	(*GetConfigurationResp)(nil),    // 5: svc.infra.setting.GetConfigurationResp
	(*AddConfigurationReq)(nil),     // 6: svc.infra.setting.AddConfigurationReq
	(*AddConfigurationResp)(nil),    // 7: svc.infra.setting.AddConfigurationResp
	(*UpdateConfigurationReq)(nil),  // 8: svc.infra.setting.UpdateConfigurationReq
	(*UpdateConfigurationResp)(nil), // 9: svc.infra.setting.UpdateConfigurationResp
	(*DeleteConfigurationReq)(nil),  // 10: svc.infra.setting.DeleteConfigurationReq
	(*DeleteConfigurationResp)(nil), // 11: svc.infra.setting.DeleteConfigurationResp
	(*timestamppb.Timestamp)(nil),   // 12: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),           // 13: google.protobuf.Empty
}
var file_svc_infra_setting_setting_proto_depIdxs = []int32{
	12, // 0: svc.infra.setting.Configuration.created_at:type_name -> google.protobuf.Timestamp
	3,  // 1: svc.infra.setting.AddConfigurationResp.configuration:type_name -> svc.infra.setting.Configuration
	3,  // 2: svc.infra.setting.DeleteConfigurationReq.condition:type_name -> svc.infra.setting.Configuration
	13, // 3: svc.infra.setting.Setting.InitDB:input_type -> google.protobuf.Empty
	4,  // 4: svc.infra.setting.Setting.GetConfiguration:input_type -> svc.infra.setting.GetConfigurationReq
	6,  // 5: svc.infra.setting.Setting.AddConfiguration:input_type -> svc.infra.setting.AddConfigurationReq
	8,  // 6: svc.infra.setting.Setting.UpdateConfiguration:input_type -> svc.infra.setting.UpdateConfigurationReq
	10, // 7: svc.infra.setting.Setting.DeleteConfiguration:input_type -> svc.infra.setting.DeleteConfigurationReq
	1,  // 8: svc.infra.setting.Setting.Greeting:input_type -> svc.infra.setting.SettingGreetingReq
	0,  // 9: svc.infra.setting.Setting.InitDB:output_type -> svc.infra.setting.InitDBResp
	5,  // 10: svc.infra.setting.Setting.GetConfiguration:output_type -> svc.infra.setting.GetConfigurationResp
	7,  // 11: svc.infra.setting.Setting.AddConfiguration:output_type -> svc.infra.setting.AddConfigurationResp
	9,  // 12: svc.infra.setting.Setting.UpdateConfiguration:output_type -> svc.infra.setting.UpdateConfigurationResp
	11, // 13: svc.infra.setting.Setting.DeleteConfiguration:output_type -> svc.infra.setting.DeleteConfigurationResp
	2,  // 14: svc.infra.setting.Setting.Greeting:output_type -> svc.infra.setting.SettingGreetingResp
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_svc_infra_setting_setting_proto_init() }
func file_svc_infra_setting_setting_proto_init() {
	if File_svc_infra_setting_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_infra_setting_setting_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_svc_infra_setting_setting_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SettingGreetingReq); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SettingGreetingResp); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Configuration); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetConfigurationReq); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetConfigurationResp); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AddConfigurationReq); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*AddConfigurationResp); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateConfigurationReq); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateConfigurationResp); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteConfigurationReq); i {
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
		file_svc_infra_setting_setting_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteConfigurationResp); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_infra_setting_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_infra_setting_setting_proto_goTypes,
		DependencyIndexes: file_svc_infra_setting_setting_proto_depIdxs,
		MessageInfos:      file_svc_infra_setting_setting_proto_msgTypes,
	}.Build()
	File_svc_infra_setting_setting_proto = out.File
	file_svc_infra_setting_setting_proto_rawDesc = nil
	file_svc_infra_setting_setting_proto_goTypes = nil
	file_svc_infra_setting_setting_proto_depIdxs = nil
}
