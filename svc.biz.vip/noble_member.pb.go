// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: svc.biz.vip/noble_member.proto

package vip

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type NobleMemberInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level      NobleLevel             `protobuf:"varint,1,opt,name=level,proto3,enum=svc.biz.vip.NobleLevel" json:"level,omitempty"` // 贵族等级
	MemberId   string                 `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`        // 用户id
	StreamerId string                 `protobuf:"bytes,3,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`  // 主播id（可为空）
	JoinTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=join_time,json=joinTime,proto3" json:"join_time,omitempty"`        // 加入时间
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`  // 过期时间
}

func (x *NobleMemberInfo) Reset() {
	*x = NobleMemberInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_noble_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NobleMemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NobleMemberInfo) ProtoMessage() {}

func (x *NobleMemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NobleMemberInfo.ProtoReflect.Descriptor instead.
func (*NobleMemberInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_member_proto_rawDescGZIP(), []int{0}
}

func (x *NobleMemberInfo) GetLevel() NobleLevel {
	if x != nil {
		return x.Level
	}
	return NobleLevel_NobleLevelUnknown
}

func (x *NobleMemberInfo) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *NobleMemberInfo) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

func (x *NobleMemberInfo) GetJoinTime() *timestamppb.Timestamp {
	if x != nil {
		return x.JoinTime
	}
	return nil
}

func (x *NobleMemberInfo) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

type JoinNobleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level      NobleLevel `protobuf:"varint,1,opt,name=level,proto3,enum=svc.biz.vip.NobleLevel" json:"level,omitempty"` // 贵族等级
	MemberId   string     `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`        // 用户id
	StreamerId string     `protobuf:"bytes,3,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"`  // 主播id（可为空；贵族可直接开通，也可在某个直播间开通）
}

func (x *JoinNobleReq) Reset() {
	*x = JoinNobleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_noble_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinNobleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinNobleReq) ProtoMessage() {}

func (x *JoinNobleReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinNobleReq.ProtoReflect.Descriptor instead.
func (*JoinNobleReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_member_proto_rawDescGZIP(), []int{1}
}

func (x *JoinNobleReq) GetLevel() NobleLevel {
	if x != nil {
		return x.Level
	}
	return NobleLevel_NobleLevelUnknown
}

func (x *JoinNobleReq) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *JoinNobleReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

type JoinNobleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinNobleResp) Reset() {
	*x = JoinNobleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_noble_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinNobleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinNobleResp) ProtoMessage() {}

func (x *JoinNobleResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinNobleResp.ProtoReflect.Descriptor instead.
func (*JoinNobleResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_member_proto_rawDescGZIP(), []int{2}
}

type GetNobleMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"` // 用户id
}

func (x *GetNobleMemberReq) Reset() {
	*x = GetNobleMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_noble_member_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNobleMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNobleMemberReq) ProtoMessage() {}

func (x *GetNobleMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_member_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNobleMemberReq.ProtoReflect.Descriptor instead.
func (*GetNobleMemberReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_member_proto_rawDescGZIP(), []int{3}
}

func (x *GetNobleMemberReq) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type GetNobleMemberResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NobleMember *NobleMemberInfo `protobuf:"bytes,1,opt,name=noble_member,json=nobleMember,proto3" json:"noble_member,omitempty"`
}

func (x *GetNobleMemberResp) Reset() {
	*x = GetNobleMemberResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_noble_member_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNobleMemberResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNobleMemberResp) ProtoMessage() {}

func (x *GetNobleMemberResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_member_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNobleMemberResp.ProtoReflect.Descriptor instead.
func (*GetNobleMemberResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_member_proto_rawDescGZIP(), []int{4}
}

func (x *GetNobleMemberResp) GetNobleMember() *NobleMemberInfo {
	if x != nil {
		return x.NobleMember
	}
	return nil
}

type GetOnlineNobleMemberByStreamerIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamerId string `protobuf:"bytes,1,opt,name=streamer_id,json=streamerId,proto3" json:"streamer_id,omitempty"` // 主播id
}

func (x *GetOnlineNobleMemberByStreamerIDReq) Reset() {
	*x = GetOnlineNobleMemberByStreamerIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_noble_member_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOnlineNobleMemberByStreamerIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOnlineNobleMemberByStreamerIDReq) ProtoMessage() {}

func (x *GetOnlineNobleMemberByStreamerIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_member_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOnlineNobleMemberByStreamerIDReq.ProtoReflect.Descriptor instead.
func (*GetOnlineNobleMemberByStreamerIDReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_member_proto_rawDescGZIP(), []int{5}
}

func (x *GetOnlineNobleMemberByStreamerIDReq) GetStreamerId() string {
	if x != nil {
		return x.StreamerId
	}
	return ""
}

type GetOnlineNobleMemberByStreamerIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*NobleMemberInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetOnlineNobleMemberByStreamerIDResp) Reset() {
	*x = GetOnlineNobleMemberByStreamerIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_vip_noble_member_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOnlineNobleMemberByStreamerIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOnlineNobleMemberByStreamerIDResp) ProtoMessage() {}

func (x *GetOnlineNobleMemberByStreamerIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_member_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOnlineNobleMemberByStreamerIDResp.ProtoReflect.Descriptor instead.
func (*GetOnlineNobleMemberByStreamerIDResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_member_proto_rawDescGZIP(), []int{6}
}

func (x *GetOnlineNobleMemberByStreamerIDResp) GetItems() []*NobleMemberInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_svc_biz_vip_noble_member_proto protoreflect.FileDescriptor

var file_svc_biz_vip_noble_member_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2f, 0x6e, 0x6f,
	0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2f, 0x6e, 0x6f, 0x62, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x62, 0x6c,
	0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x7b,
	0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x2d,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x62, 0x6c,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x4a,
	0x6f, 0x69, 0x6e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x30, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x55,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x0c, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6e, 0x6f, 0x62, 0x6c, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a,
	0x24, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76,
	0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xb4, 0x02, 0x0a, 0x0b, 0x4e, 0x6f,
	0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x09, 0x4a, 0x6f, 0x69,
	0x6e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e,
	0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x1f, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x89, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x30, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69,
	0x70, 0x3b, 0x76, 0x69, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_vip_noble_member_proto_rawDescOnce sync.Once
	file_svc_biz_vip_noble_member_proto_rawDescData = file_svc_biz_vip_noble_member_proto_rawDesc
)

func file_svc_biz_vip_noble_member_proto_rawDescGZIP() []byte {
	file_svc_biz_vip_noble_member_proto_rawDescOnce.Do(func() {
		file_svc_biz_vip_noble_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_vip_noble_member_proto_rawDescData)
	})
	return file_svc_biz_vip_noble_member_proto_rawDescData
}

var file_svc_biz_vip_noble_member_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_svc_biz_vip_noble_member_proto_goTypes = []any{
	(*NobleMemberInfo)(nil),                      // 0: svc.biz.vip.NobleMemberInfo
	(*JoinNobleReq)(nil),                         // 1: svc.biz.vip.JoinNobleReq
	(*JoinNobleResp)(nil),                        // 2: svc.biz.vip.JoinNobleResp
	(*GetNobleMemberReq)(nil),                    // 3: svc.biz.vip.GetNobleMemberReq
	(*GetNobleMemberResp)(nil),                   // 4: svc.biz.vip.GetNobleMemberResp
	(*GetOnlineNobleMemberByStreamerIDReq)(nil),  // 5: svc.biz.vip.GetOnlineNobleMemberByStreamerIDReq
	(*GetOnlineNobleMemberByStreamerIDResp)(nil), // 6: svc.biz.vip.GetOnlineNobleMemberByStreamerIDResp
	(NobleLevel)(0),                              // 7: svc.biz.vip.NobleLevel
	(*timestamppb.Timestamp)(nil),                // 8: google.protobuf.Timestamp
}
var file_svc_biz_vip_noble_member_proto_depIdxs = []int32{
	7, // 0: svc.biz.vip.NobleMemberInfo.level:type_name -> svc.biz.vip.NobleLevel
	8, // 1: svc.biz.vip.NobleMemberInfo.join_time:type_name -> google.protobuf.Timestamp
	8, // 2: svc.biz.vip.NobleMemberInfo.expire_time:type_name -> google.protobuf.Timestamp
	7, // 3: svc.biz.vip.JoinNobleReq.level:type_name -> svc.biz.vip.NobleLevel
	0, // 4: svc.biz.vip.GetNobleMemberResp.noble_member:type_name -> svc.biz.vip.NobleMemberInfo
	0, // 5: svc.biz.vip.GetOnlineNobleMemberByStreamerIDResp.items:type_name -> svc.biz.vip.NobleMemberInfo
	1, // 6: svc.biz.vip.NobleMember.JoinNoble:input_type -> svc.biz.vip.JoinNobleReq
	3, // 7: svc.biz.vip.NobleMember.GetNobleMember:input_type -> svc.biz.vip.GetNobleMemberReq
	5, // 8: svc.biz.vip.NobleMember.GetOnlineNobleMemberByStreamerID:input_type -> svc.biz.vip.GetOnlineNobleMemberByStreamerIDReq
	2, // 9: svc.biz.vip.NobleMember.JoinNoble:output_type -> svc.biz.vip.JoinNobleResp
	4, // 10: svc.biz.vip.NobleMember.GetNobleMember:output_type -> svc.biz.vip.GetNobleMemberResp
	6, // 11: svc.biz.vip.NobleMember.GetOnlineNobleMemberByStreamerID:output_type -> svc.biz.vip.GetOnlineNobleMemberByStreamerIDResp
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_svc_biz_vip_noble_member_proto_init() }
func file_svc_biz_vip_noble_member_proto_init() {
	if File_svc_biz_vip_noble_member_proto != nil {
		return
	}
	file_svc_biz_vip_noble_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_vip_noble_member_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NobleMemberInfo); i {
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
		file_svc_biz_vip_noble_member_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*JoinNobleReq); i {
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
		file_svc_biz_vip_noble_member_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*JoinNobleResp); i {
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
		file_svc_biz_vip_noble_member_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetNobleMemberReq); i {
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
		file_svc_biz_vip_noble_member_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetNobleMemberResp); i {
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
		file_svc_biz_vip_noble_member_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetOnlineNobleMemberByStreamerIDReq); i {
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
		file_svc_biz_vip_noble_member_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetOnlineNobleMemberByStreamerIDResp); i {
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
			RawDescriptor: file_svc_biz_vip_noble_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_vip_noble_member_proto_goTypes,
		DependencyIndexes: file_svc_biz_vip_noble_member_proto_depIdxs,
		MessageInfos:      file_svc_biz_vip_noble_member_proto_msgTypes,
	}.Build()
	File_svc_biz_vip_noble_member_proto = out.File
	file_svc_biz_vip_noble_member_proto_rawDesc = nil
	file_svc_biz_vip_noble_member_proto_goTypes = nil
	file_svc_biz_vip_noble_member_proto_depIdxs = nil
}
