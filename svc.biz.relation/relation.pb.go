// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: svc.biz.relation/relation.proto

package relation

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

type RelationType int32

const (
	RelationType_RelationTypeUnknown         RelationType = 0 // 未知
	RelationType_RelationTypeFollow          RelationType = 1 // 关注主播
	RelationType_RelationTypeHistory         RelationType = 2 // 观看历史
	RelationType_RelationTypeMuzzle          RelationType = 3 // 禁言
	RelationType_RelationTypeBlacklistIP     RelationType = 4 // ip黑名单
	RelationType_RelationTypeBlacklistDevice RelationType = 5 // 设备黑名单
)

// Enum value maps for RelationType.
var (
	RelationType_name = map[int32]string{
		0: "RelationTypeUnknown",
		1: "RelationTypeFollow",
		2: "RelationTypeHistory",
		3: "RelationTypeMuzzle",
		4: "RelationTypeBlacklistIP",
		5: "RelationTypeBlacklistDevice",
	}
	RelationType_value = map[string]int32{
		"RelationTypeUnknown":         0,
		"RelationTypeFollow":          1,
		"RelationTypeHistory":         2,
		"RelationTypeMuzzle":          3,
		"RelationTypeBlacklistIP":     4,
		"RelationTypeBlacklistDevice": 5,
	}
)

func (x RelationType) Enum() *RelationType {
	p := new(RelationType)
	*p = x
	return p
}

func (x RelationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationType) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_biz_relation_relation_proto_enumTypes[0].Descriptor()
}

func (RelationType) Type() protoreflect.EnumType {
	return &file_svc_biz_relation_relation_proto_enumTypes[0]
}

func (x RelationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationType.Descriptor instead.
func (RelationType) EnumDescriptor() ([]byte, []int) {
	return file_svc_biz_relation_relation_proto_rawDescGZIP(), []int{0}
}

type RelationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationType   RelationType           `protobuf:"varint,1,opt,name=relation_type,json=relationType,proto3,enum=svc.biz.relation.RelationType" json:"relation_type,omitempty"` // 关系类型
	Member         string                 `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`                                                                     // 成员（名单属于谁）
	RelationMember string                 `protobuf:"bytes,3,opt,name=relation_member,json=relationMember,proto3" json:"relation_member,omitempty"`                               // 产生关系的成员（名单内有谁）
	ExpireAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`                                                 // 过期时间（可无，为空则永久有效）
	Remark         string                 `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`                                                                     // 备注
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                                            // 创建时间
}

func (x *RelationInfo) Reset() {
	*x = RelationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_relation_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationInfo) ProtoMessage() {}

func (x *RelationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_relation_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationInfo.ProtoReflect.Descriptor instead.
func (*RelationInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_relation_relation_proto_rawDescGZIP(), []int{0}
}

func (x *RelationInfo) GetRelationType() RelationType {
	if x != nil {
		return x.RelationType
	}
	return RelationType_RelationTypeUnknown
}

func (x *RelationInfo) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *RelationInfo) GetRelationMember() string {
	if x != nil {
		return x.RelationMember
	}
	return ""
}

func (x *RelationInfo) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

func (x *RelationInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *RelationInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type RelationAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationInfo *RelationInfo `protobuf:"bytes,1,opt,name=relation_info,json=relationInfo,proto3" json:"relation_info,omitempty"`
}

func (x *RelationAddReq) Reset() {
	*x = RelationAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_relation_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationAddReq) ProtoMessage() {}

func (x *RelationAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_relation_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationAddReq.ProtoReflect.Descriptor instead.
func (*RelationAddReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_relation_relation_proto_rawDescGZIP(), []int{1}
}

func (x *RelationAddReq) GetRelationInfo() *RelationInfo {
	if x != nil {
		return x.RelationInfo
	}
	return nil
}

type RelationDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationType   RelationType `protobuf:"varint,1,opt,name=relation_type,json=relationType,proto3,enum=svc.biz.relation.RelationType" json:"relation_type,omitempty"`
	Member         string       `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`                                       // 成员
	RelationMember string       `protobuf:"bytes,3,opt,name=relation_member,json=relationMember,proto3" json:"relation_member,omitempty"` // 产生关系的成员
}

func (x *RelationDelReq) Reset() {
	*x = RelationDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_relation_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationDelReq) ProtoMessage() {}

func (x *RelationDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_relation_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationDelReq.ProtoReflect.Descriptor instead.
func (*RelationDelReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_relation_relation_proto_rawDescGZIP(), []int{2}
}

func (x *RelationDelReq) GetRelationType() RelationType {
	if x != nil {
		return x.RelationType
	}
	return RelationType_RelationTypeUnknown
}

func (x *RelationDelReq) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *RelationDelReq) GetRelationMember() string {
	if x != nil {
		return x.RelationMember
	}
	return ""
}

type RelationHasReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationType   RelationType `protobuf:"varint,1,opt,name=relation_type,json=relationType,proto3,enum=svc.biz.relation.RelationType" json:"relation_type,omitempty"`
	Member         string       `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`                                       // 成员
	RelationMember string       `protobuf:"bytes,3,opt,name=relation_member,json=relationMember,proto3" json:"relation_member,omitempty"` // 产生关系的成员
}

func (x *RelationHasReq) Reset() {
	*x = RelationHasReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_relation_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationHasReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationHasReq) ProtoMessage() {}

func (x *RelationHasReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_relation_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationHasReq.ProtoReflect.Descriptor instead.
func (*RelationHasReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_relation_relation_proto_rawDescGZIP(), []int{3}
}

func (x *RelationHasReq) GetRelationType() RelationType {
	if x != nil {
		return x.RelationType
	}
	return RelationType_RelationTypeUnknown
}

func (x *RelationHasReq) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *RelationHasReq) GetRelationMember() string {
	if x != nil {
		return x.RelationMember
	}
	return ""
}

type RelationHasResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RelationHasResp) Reset() {
	*x = RelationHasResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_relation_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationHasResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationHasResp) ProtoMessage() {}

func (x *RelationHasResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_relation_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationHasResp.ProtoReflect.Descriptor instead.
func (*RelationHasResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_relation_relation_proto_rawDescGZIP(), []int{4}
}

func (x *RelationHasResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type GetRelationListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationType RelationType `protobuf:"varint,1,opt,name=relation_type,json=relationType,proto3,enum=svc.biz.relation.RelationType" json:"relation_type,omitempty"`
	Member       string       `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	Page         int64        `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit        int64        `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetRelationListReq) Reset() {
	*x = GetRelationListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_relation_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationListReq) ProtoMessage() {}

func (x *GetRelationListReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_relation_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationListReq.ProtoReflect.Descriptor instead.
func (*GetRelationListReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_relation_relation_proto_rawDescGZIP(), []int{5}
}

func (x *GetRelationListReq) GetRelationType() RelationType {
	if x != nil {
		return x.RelationType
	}
	return RelationType_RelationTypeUnknown
}

func (x *GetRelationListReq) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *GetRelationListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetRelationListReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetRelationListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*RelationInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetRelationListResp) Reset() {
	*x = GetRelationListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_relation_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRelationListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRelationListResp) ProtoMessage() {}

func (x *GetRelationListResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_relation_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRelationListResp.ProtoReflect.Descriptor instead.
func (*GetRelationListResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_relation_relation_proto_rawDescGZIP(), []int{6}
}

func (x *GetRelationListResp) GetItems() []*RelationInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_svc_biz_relation_relation_proto protoreflect.FileDescriptor

var file_svc_biz_relation_relation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa0, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x43, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x55, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x43, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x0e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x43,
	0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x96, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x61, 0x73, 0x52, 0x65, 0x71, 0x12, 0x43, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x29, 0x0a,
	0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x43, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x2a, 0xae, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x75,
	0x7a, 0x7a, 0x6c, 0x65, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x49,
	0x50, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x10, 0x05, 0x32, 0xd0, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x47, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64,
	0x12, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0b, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x12, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x61, 0x73, 0x12, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x25, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x73, 0x76, 0x63,
	0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_relation_relation_proto_rawDescOnce sync.Once
	file_svc_biz_relation_relation_proto_rawDescData = file_svc_biz_relation_relation_proto_rawDesc
)

func file_svc_biz_relation_relation_proto_rawDescGZIP() []byte {
	file_svc_biz_relation_relation_proto_rawDescOnce.Do(func() {
		file_svc_biz_relation_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_relation_relation_proto_rawDescData)
	})
	return file_svc_biz_relation_relation_proto_rawDescData
}

var file_svc_biz_relation_relation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_svc_biz_relation_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_svc_biz_relation_relation_proto_goTypes = []interface{}{
	(RelationType)(0),             // 0: svc.biz.relation.RelationType
	(*RelationInfo)(nil),          // 1: svc.biz.relation.RelationInfo
	(*RelationAddReq)(nil),        // 2: svc.biz.relation.RelationAddReq
	(*RelationDelReq)(nil),        // 3: svc.biz.relation.RelationDelReq
	(*RelationHasReq)(nil),        // 4: svc.biz.relation.RelationHasReq
	(*RelationHasResp)(nil),       // 5: svc.biz.relation.RelationHasResp
	(*GetRelationListReq)(nil),    // 6: svc.biz.relation.GetRelationListReq
	(*GetRelationListResp)(nil),   // 7: svc.biz.relation.GetRelationListResp
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_svc_biz_relation_relation_proto_depIdxs = []int32{
	0,  // 0: svc.biz.relation.RelationInfo.relation_type:type_name -> svc.biz.relation.RelationType
	8,  // 1: svc.biz.relation.RelationInfo.expire_at:type_name -> google.protobuf.Timestamp
	8,  // 2: svc.biz.relation.RelationInfo.created_at:type_name -> google.protobuf.Timestamp
	1,  // 3: svc.biz.relation.RelationAddReq.relation_info:type_name -> svc.biz.relation.RelationInfo
	0,  // 4: svc.biz.relation.RelationDelReq.relation_type:type_name -> svc.biz.relation.RelationType
	0,  // 5: svc.biz.relation.RelationHasReq.relation_type:type_name -> svc.biz.relation.RelationType
	0,  // 6: svc.biz.relation.GetRelationListReq.relation_type:type_name -> svc.biz.relation.RelationType
	1,  // 7: svc.biz.relation.GetRelationListResp.items:type_name -> svc.biz.relation.RelationInfo
	2,  // 8: svc.biz.relation.Relation.RelationAdd:input_type -> svc.biz.relation.RelationAddReq
	3,  // 9: svc.biz.relation.Relation.RelationDel:input_type -> svc.biz.relation.RelationDelReq
	4,  // 10: svc.biz.relation.Relation.RelationHas:input_type -> svc.biz.relation.RelationHasReq
	6,  // 11: svc.biz.relation.Relation.GetRelationList:input_type -> svc.biz.relation.GetRelationListReq
	9,  // 12: svc.biz.relation.Relation.RelationAdd:output_type -> google.protobuf.Empty
	9,  // 13: svc.biz.relation.Relation.RelationDel:output_type -> google.protobuf.Empty
	5,  // 14: svc.biz.relation.Relation.RelationHas:output_type -> svc.biz.relation.RelationHasResp
	7,  // 15: svc.biz.relation.Relation.GetRelationList:output_type -> svc.biz.relation.GetRelationListResp
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_svc_biz_relation_relation_proto_init() }
func file_svc_biz_relation_relation_proto_init() {
	if File_svc_biz_relation_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_relation_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationInfo); i {
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
		file_svc_biz_relation_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationAddReq); i {
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
		file_svc_biz_relation_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationDelReq); i {
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
		file_svc_biz_relation_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationHasReq); i {
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
		file_svc_biz_relation_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationHasResp); i {
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
		file_svc_biz_relation_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationListReq); i {
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
		file_svc_biz_relation_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRelationListResp); i {
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
			RawDescriptor: file_svc_biz_relation_relation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_relation_relation_proto_goTypes,
		DependencyIndexes: file_svc_biz_relation_relation_proto_depIdxs,
		EnumInfos:         file_svc_biz_relation_relation_proto_enumTypes,
		MessageInfos:      file_svc_biz_relation_relation_proto_msgTypes,
	}.Build()
	File_svc_biz_relation_relation_proto = out.File
	file_svc_biz_relation_relation_proto_rawDesc = nil
	file_svc_biz_relation_relation_proto_goTypes = nil
	file_svc_biz_relation_relation_proto_depIdxs = nil
}
