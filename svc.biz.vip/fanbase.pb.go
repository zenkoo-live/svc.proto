// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: svc.biz.vip/fanbase.proto

package vip

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

type FanbaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FanbaseId  string                 `protobuf:"bytes,1,opt,name=fanbase_id,json=fanbaseId,proto3" json:"fanbase_id,omitempty"`
	StreamerId string                 `protobuf:"bytes,2,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
	Name       string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 创建时间
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,102,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // 更新时间
}

func (x *FanbaseInfo) Reset() {
	*x = FanbaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanbaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanbaseInfo) ProtoMessage() {}

func (x *FanbaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanbaseInfo.ProtoReflect.Descriptor instead.
func (*FanbaseInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_proto_rawDescGZIP(), []int{0}
}

func (x *FanbaseInfo) GetFanbaseId() string {
	if x != nil {
		return x.FanbaseId
	}
	return ""
}

func (x *FanbaseInfo) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *FanbaseInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FanbaseInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FanbaseInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateFanbaseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fanbase *FanbaseInfo `protobuf:"bytes,1,opt,name=fanbase,proto3" json:"fanbase,omitempty"`
}

func (x *CreateFanbaseReq) Reset() {
	*x = CreateFanbaseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFanbaseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFanbaseReq) ProtoMessage() {}

func (x *CreateFanbaseReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFanbaseReq.ProtoReflect.Descriptor instead.
func (*CreateFanbaseReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFanbaseReq) GetFanbase() *FanbaseInfo {
	if x != nil {
		return x.Fanbase
	}
	return nil
}

type CreateFanbaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fanbase *FanbaseInfo `protobuf:"bytes,1,opt,name=fanbase,proto3" json:"fanbase,omitempty"`
}

func (x *CreateFanbaseResp) Reset() {
	*x = CreateFanbaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFanbaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFanbaseResp) ProtoMessage() {}

func (x *CreateFanbaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFanbaseResp.ProtoReflect.Descriptor instead.
func (*CreateFanbaseResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFanbaseResp) GetFanbase() *FanbaseInfo {
	if x != nil {
		return x.Fanbase
	}
	return nil
}

type GetFanbaseByStreamerIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamerId string `protobuf:"bytes,1,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
}

func (x *GetFanbaseByStreamerIDResp) Reset() {
	*x = GetFanbaseByStreamerIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFanbaseByStreamerIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFanbaseByStreamerIDResp) ProtoMessage() {}

func (x *GetFanbaseByStreamerIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFanbaseByStreamerIDResp.ProtoReflect.Descriptor instead.
func (*GetFanbaseByStreamerIDResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_proto_rawDescGZIP(), []int{3}
}

func (x *GetFanbaseByStreamerIDResp) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

type GetFanbaseByNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetFanbaseByNameReq) Reset() {
	*x = GetFanbaseByNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFanbaseByNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFanbaseByNameReq) ProtoMessage() {}

func (x *GetFanbaseByNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFanbaseByNameReq.ProtoReflect.Descriptor instead.
func (*GetFanbaseByNameReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_proto_rawDescGZIP(), []int{4}
}

func (x *GetFanbaseByNameReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetFanbaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fanbase *FanbaseInfo `protobuf:"bytes,1,opt,name=fanbase,proto3" json:"fanbase,omitempty"`
}

func (x *GetFanbaseResp) Reset() {
	*x = GetFanbaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFanbaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFanbaseResp) ProtoMessage() {}

func (x *GetFanbaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFanbaseResp.ProtoReflect.Descriptor instead.
func (*GetFanbaseResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_proto_rawDescGZIP(), []int{5}
}

func (x *GetFanbaseResp) GetFanbase() *FanbaseInfo {
	if x != nil {
		return x.Fanbase
	}
	return nil
}

type UpdateFanbaseByStreamerIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamerId string                 `protobuf:"bytes,1,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`
	Fanbase    *FanbaseInfo           `protobuf:"bytes,2,opt,name=fanbase,proto3" json:"fanbase,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateFanbaseByStreamerIDReq) Reset() {
	*x = UpdateFanbaseByStreamerIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_fanbase_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFanbaseByStreamerIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFanbaseByStreamerIDReq) ProtoMessage() {}

func (x *UpdateFanbaseByStreamerIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_fanbase_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFanbaseByStreamerIDReq.ProtoReflect.Descriptor instead.
func (*UpdateFanbaseByStreamerIDReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_fanbase_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFanbaseByStreamerIDReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *UpdateFanbaseByStreamerIDReq) GetFanbase() *FanbaseInfo {
	if x != nil {
		return x.Fanbase
	}
	return nil
}

func (x *UpdateFanbaseByStreamerIDReq) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_svc_biz_vip_fanbase_proto protoreflect.FileDescriptor

var file_svc_biz_vip_fanbase_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2f, 0x66, 0x61,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x0b, 0x46, 0x61, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x61,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x46, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a, 0x07, 0x66, 0x61, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x66, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x32, 0x0a, 0x07, 0x66, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x46,
	0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x66, 0x61, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x32, 0x0a, 0x07, 0x66, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x46,
	0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x66, 0x61, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x07, 0x66, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x76, 0x69, 0x70, 0x2e, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x66, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0xf4, 0x02, 0x0a, 0x07, 0x46, 0x61, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69,
	0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x27,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65,
	0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x1a, 0x1b, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x19, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x42, 0x79, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x29, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x13, 0x5a,
	0x11, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x3b, 0x76,
	0x69, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_vip_fanbase_proto_rawDescOnce sync.Once
	file_svc_biz_vip_fanbase_proto_rawDescData = file_svc_biz_vip_fanbase_proto_rawDesc
)

func file_svc_biz_vip_fanbase_proto_rawDescGZIP() []byte {
	file_svc_biz_vip_fanbase_proto_rawDescOnce.Do(func() {
		file_svc_biz_vip_fanbase_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_vip_fanbase_proto_rawDescData)
	})
	return file_svc_biz_vip_fanbase_proto_rawDescData
}

var file_svc_biz_vip_fanbase_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_svc_biz_vip_fanbase_proto_goTypes = []any{
	(*FanbaseInfo)(nil),                  // 0: svc.biz.vip.FanbaseInfo
	(*CreateFanbaseReq)(nil),             // 1: svc.biz.vip.CreateFanbaseReq
	(*CreateFanbaseResp)(nil),            // 2: svc.biz.vip.CreateFanbaseResp
	(*GetFanbaseByStreamerIDResp)(nil),   // 3: svc.biz.vip.GetFanbaseByStreamerIDResp
	(*GetFanbaseByNameReq)(nil),          // 4: svc.biz.vip.GetFanbaseByNameReq
	(*GetFanbaseResp)(nil),               // 5: svc.biz.vip.GetFanbaseResp
	(*UpdateFanbaseByStreamerIDReq)(nil), // 6: svc.biz.vip.UpdateFanbaseByStreamerIDReq
	(*timestamppb.Timestamp)(nil),        // 7: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil),        // 8: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),                // 9: google.protobuf.Empty
}
var file_svc_biz_vip_fanbase_proto_depIdxs = []int32{
	7,  // 0: svc.biz.vip.FanbaseInfo.created_at:type_name -> google.protobuf.Timestamp
	7,  // 1: svc.biz.vip.FanbaseInfo.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: svc.biz.vip.CreateFanbaseReq.fanbase:type_name -> svc.biz.vip.FanbaseInfo
	0,  // 3: svc.biz.vip.CreateFanbaseResp.fanbase:type_name -> svc.biz.vip.FanbaseInfo
	0,  // 4: svc.biz.vip.GetFanbaseResp.fanbase:type_name -> svc.biz.vip.FanbaseInfo
	0,  // 5: svc.biz.vip.UpdateFanbaseByStreamerIDReq.fanbase:type_name -> svc.biz.vip.FanbaseInfo
	8,  // 6: svc.biz.vip.UpdateFanbaseByStreamerIDReq.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 7: svc.biz.vip.Fanbase.CreateFanbase:input_type -> svc.biz.vip.CreateFanbaseReq
	3,  // 8: svc.biz.vip.Fanbase.GetFanbaseByStreamerID:input_type -> svc.biz.vip.GetFanbaseByStreamerIDResp
	4,  // 9: svc.biz.vip.Fanbase.GetFanbaseByName:input_type -> svc.biz.vip.GetFanbaseByNameReq
	6,  // 10: svc.biz.vip.Fanbase.UpdateFanbaseByStreamerID:input_type -> svc.biz.vip.UpdateFanbaseByStreamerIDReq
	2,  // 11: svc.biz.vip.Fanbase.CreateFanbase:output_type -> svc.biz.vip.CreateFanbaseResp
	5,  // 12: svc.biz.vip.Fanbase.GetFanbaseByStreamerID:output_type -> svc.biz.vip.GetFanbaseResp
	5,  // 13: svc.biz.vip.Fanbase.GetFanbaseByName:output_type -> svc.biz.vip.GetFanbaseResp
	9,  // 14: svc.biz.vip.Fanbase.UpdateFanbaseByStreamerID:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_svc_biz_vip_fanbase_proto_init() }
func file_svc_biz_vip_fanbase_proto_init() {
	if File_svc_biz_vip_fanbase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_vip_fanbase_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FanbaseInfo); i {
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
		file_svc_biz_vip_fanbase_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFanbaseReq); i {
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
		file_svc_biz_vip_fanbase_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFanbaseResp); i {
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
		file_svc_biz_vip_fanbase_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetFanbaseByStreamerIDResp); i {
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
		file_svc_biz_vip_fanbase_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetFanbaseByNameReq); i {
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
		file_svc_biz_vip_fanbase_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetFanbaseResp); i {
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
		file_svc_biz_vip_fanbase_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateFanbaseByStreamerIDReq); i {
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
			RawDescriptor: file_svc_biz_vip_fanbase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_vip_fanbase_proto_goTypes,
		DependencyIndexes: file_svc_biz_vip_fanbase_proto_depIdxs,
		MessageInfos:      file_svc_biz_vip_fanbase_proto_msgTypes,
	}.Build()
	File_svc_biz_vip_fanbase_proto = out.File
	file_svc_biz_vip_fanbase_proto_rawDesc = nil
	file_svc_biz_vip_fanbase_proto_goTypes = nil
	file_svc_biz_vip_fanbase_proto_depIdxs = nil
}
