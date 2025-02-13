// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: svc.biz.vip/level.proto

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

type MemberLevelInfo struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	MemberId        string                 `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Level           int32                  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Exp             int64                  `protobuf:"varint,3,opt,name=exp,proto3" json:"exp,omitempty"`
	ExpCurrentLevel int64                  `protobuf:"varint,4,opt,name=exp_current_level,json=expCurrentLevel,proto3" json:"exp_current_level,omitempty"` // 当前等级所需经验
	ExpNextLevel    int64                  `protobuf:"varint,5,opt,name=exp_next_level,json=expNextLevel,proto3" json:"exp_next_level,omitempty"`          // 下一等级所需经验
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *MemberLevelInfo) Reset() {
	*x = MemberLevelInfo{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberLevelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberLevelInfo) ProtoMessage() {}

func (x *MemberLevelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberLevelInfo.ProtoReflect.Descriptor instead.
func (*MemberLevelInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{0}
}

func (x *MemberLevelInfo) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *MemberLevelInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *MemberLevelInfo) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *MemberLevelInfo) GetExpCurrentLevel() int64 {
	if x != nil {
		return x.ExpCurrentLevel
	}
	return 0
}

func (x *MemberLevelInfo) GetExpNextLevel() int64 {
	if x != nil {
		return x.ExpNextLevel
	}
	return 0
}

type GetMemberLevelReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MemberId      string                 `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMemberLevelReq) Reset() {
	*x = GetMemberLevelReq{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemberLevelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberLevelReq) ProtoMessage() {}

func (x *GetMemberLevelReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberLevelReq.ProtoReflect.Descriptor instead.
func (*GetMemberLevelReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{1}
}

func (x *GetMemberLevelReq) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type GetMemberLevelResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LevelInfo     *MemberLevelInfo       `protobuf:"bytes,1,opt,name=level_info,json=levelInfo,proto3" json:"level_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMemberLevelResp) Reset() {
	*x = GetMemberLevelResp{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemberLevelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberLevelResp) ProtoMessage() {}

func (x *GetMemberLevelResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberLevelResp.ProtoReflect.Descriptor instead.
func (*GetMemberLevelResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{2}
}

func (x *GetMemberLevelResp) GetLevelInfo() *MemberLevelInfo {
	if x != nil {
		return x.LevelInfo
	}
	return nil
}

type MGetMemberLevelReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MemberIds     []string               `protobuf:"bytes,1,rep,name=member_ids,json=memberIds,proto3" json:"member_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MGetMemberLevelReq) Reset() {
	*x = MGetMemberLevelReq{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MGetMemberLevelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetMemberLevelReq) ProtoMessage() {}

func (x *MGetMemberLevelReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetMemberLevelReq.ProtoReflect.Descriptor instead.
func (*MGetMemberLevelReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{3}
}

func (x *MGetMemberLevelReq) GetMemberIds() []string {
	if x != nil {
		return x.MemberIds
	}
	return nil
}

type MGetMemberLevelResp struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Items         map[string]*MemberLevelInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MGetMemberLevelResp) Reset() {
	*x = MGetMemberLevelResp{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MGetMemberLevelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetMemberLevelResp) ProtoMessage() {}

func (x *MGetMemberLevelResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetMemberLevelResp.ProtoReflect.Descriptor instead.
func (*MGetMemberLevelResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{4}
}

func (x *MGetMemberLevelResp) GetItems() map[string]*MemberLevelInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type LevelInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Level         int32                  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	MaxExp        int64                  `protobuf:"varint,2,opt,name=max_exp,json=maxExp,proto3" json:"max_exp,omitempty"`
	Icon          string                 `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Color         string                 `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 创建时间
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,102,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // 更新时间
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LevelInfo) Reset() {
	*x = LevelInfo{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LevelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelInfo) ProtoMessage() {}

func (x *LevelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelInfo.ProtoReflect.Descriptor instead.
func (*LevelInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{5}
}

func (x *LevelInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *LevelInfo) GetMaxExp() int64 {
	if x != nil {
		return x.MaxExp
	}
	return 0
}

func (x *LevelInfo) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *LevelInfo) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *LevelInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *LevelInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetAllLevelListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllLevelListReq) Reset() {
	*x = GetAllLevelListReq{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllLevelListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllLevelListReq) ProtoMessage() {}

func (x *GetAllLevelListReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllLevelListReq.ProtoReflect.Descriptor instead.
func (*GetAllLevelListReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{6}
}

type GetLevelListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                            // 页数
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`                          // 条数
	WithTotal     bool                   `protobuf:"varint,3,opt,name=with_total,json=withTotal,proto3" json:"with_total,omitempty"` // 是否返回总数
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLevelListReq) Reset() {
	*x = GetLevelListReq{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLevelListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLevelListReq) ProtoMessage() {}

func (x *GetLevelListReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLevelListReq.ProtoReflect.Descriptor instead.
func (*GetLevelListReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{7}
}

func (x *GetLevelListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetLevelListReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetLevelListReq) GetWithTotal() bool {
	if x != nil {
		return x.WithTotal
	}
	return false
}

type GetLevelListResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*LevelInfo           `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total         int64                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLevelListResp) Reset() {
	*x = GetLevelListResp{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLevelListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLevelListResp) ProtoMessage() {}

func (x *GetLevelListResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLevelListResp.ProtoReflect.Descriptor instead.
func (*GetLevelListResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{8}
}

func (x *GetLevelListResp) GetItems() []*LevelInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GetLevelListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AddLevelReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LevelInfo     *LevelInfo             `protobuf:"bytes,1,opt,name=level_info,json=levelInfo,proto3" json:"level_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddLevelReq) Reset() {
	*x = AddLevelReq{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddLevelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLevelReq) ProtoMessage() {}

func (x *AddLevelReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLevelReq.ProtoReflect.Descriptor instead.
func (*AddLevelReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{9}
}

func (x *AddLevelReq) GetLevelInfo() *LevelInfo {
	if x != nil {
		return x.LevelInfo
	}
	return nil
}

type AddLevelResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddLevelResp) Reset() {
	*x = AddLevelResp{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddLevelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLevelResp) ProtoMessage() {}

func (x *AddLevelResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLevelResp.ProtoReflect.Descriptor instead.
func (*AddLevelResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{10}
}

type UpdateLevelReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Level         int32                  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	LevelInfo     *LevelInfo             `protobuf:"bytes,2,opt,name=level_info,json=levelInfo,proto3" json:"level_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateLevelReq) Reset() {
	*x = UpdateLevelReq{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateLevelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLevelReq) ProtoMessage() {}

func (x *UpdateLevelReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLevelReq.ProtoReflect.Descriptor instead.
func (*UpdateLevelReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateLevelReq) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *UpdateLevelReq) GetLevelInfo() *LevelInfo {
	if x != nil {
		return x.LevelInfo
	}
	return nil
}

type UpdateLevelReqResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateLevelReqResp) Reset() {
	*x = UpdateLevelReqResp{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateLevelReqResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLevelReqResp) ProtoMessage() {}

func (x *UpdateLevelReqResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLevelReqResp.ProtoReflect.Descriptor instead.
func (*UpdateLevelReqResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{12}
}

type DelLevelReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Level         int32                  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelLevelReq) Reset() {
	*x = DelLevelReq{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelLevelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelLevelReq) ProtoMessage() {}

func (x *DelLevelReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelLevelReq.ProtoReflect.Descriptor instead.
func (*DelLevelReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{13}
}

func (x *DelLevelReq) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type DelLevelResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DelLevelResp) Reset() {
	*x = DelLevelResp{}
	mi := &file_svc_biz_vip_level_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelLevelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelLevelResp) ProtoMessage() {}

func (x *DelLevelResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_level_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelLevelResp.ProtoReflect.Descriptor instead.
func (*DelLevelResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_level_proto_rawDescGZIP(), []int{14}
}

var File_svc_biz_vip_level_proto protoreflect.FileDescriptor

var file_svc_biz_vip_level_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70,
	0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x65, 0x78, 0x70,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0e,
	0x65, 0x78, 0x70, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x4e, 0x65, 0x78, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0x30, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0a, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x33, 0x0a, 0x12, 0x4d, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0xb0, 0x01, 0x0a,
	0x13, 0x4d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69,
	0x70, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x56, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xda, 0x01, 0x0a, 0x09, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x45, 0x78, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x14, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x22, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x77, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x56,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x44, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x0e, 0x0a, 0x0c,
	0x41, 0x64, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5d, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x23, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x0e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x32, 0xad, 0x04, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x4d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x53, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1f, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x2e,
	0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x41, 0x64, 0x64, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x1b, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69,
	0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x1a, 0x1f, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x18, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x44, 0x65,
	0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e,
	0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x3b, 0x76, 0x69, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_vip_level_proto_rawDescOnce sync.Once
	file_svc_biz_vip_level_proto_rawDescData = file_svc_biz_vip_level_proto_rawDesc
)

func file_svc_biz_vip_level_proto_rawDescGZIP() []byte {
	file_svc_biz_vip_level_proto_rawDescOnce.Do(func() {
		file_svc_biz_vip_level_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_vip_level_proto_rawDescData)
	})
	return file_svc_biz_vip_level_proto_rawDescData
}

var file_svc_biz_vip_level_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_svc_biz_vip_level_proto_goTypes = []any{
	(*MemberLevelInfo)(nil),       // 0: svc.biz.vip.MemberLevelInfo
	(*GetMemberLevelReq)(nil),     // 1: svc.biz.vip.GetMemberLevelReq
	(*GetMemberLevelResp)(nil),    // 2: svc.biz.vip.GetMemberLevelResp
	(*MGetMemberLevelReq)(nil),    // 3: svc.biz.vip.MGetMemberLevelReq
	(*MGetMemberLevelResp)(nil),   // 4: svc.biz.vip.MGetMemberLevelResp
	(*LevelInfo)(nil),             // 5: svc.biz.vip.LevelInfo
	(*GetAllLevelListReq)(nil),    // 6: svc.biz.vip.GetAllLevelListReq
	(*GetLevelListReq)(nil),       // 7: svc.biz.vip.GetLevelListReq
	(*GetLevelListResp)(nil),      // 8: svc.biz.vip.GetLevelListResp
	(*AddLevelReq)(nil),           // 9: svc.biz.vip.AddLevelReq
	(*AddLevelResp)(nil),          // 10: svc.biz.vip.AddLevelResp
	(*UpdateLevelReq)(nil),        // 11: svc.biz.vip.UpdateLevelReq
	(*UpdateLevelReqResp)(nil),    // 12: svc.biz.vip.UpdateLevelReqResp
	(*DelLevelReq)(nil),           // 13: svc.biz.vip.DelLevelReq
	(*DelLevelResp)(nil),          // 14: svc.biz.vip.DelLevelResp
	nil,                           // 15: svc.biz.vip.MGetMemberLevelResp.ItemsEntry
	(*timestamppb.Timestamp)(nil), // 16: google.protobuf.Timestamp
}
var file_svc_biz_vip_level_proto_depIdxs = []int32{
	0,  // 0: svc.biz.vip.GetMemberLevelResp.level_info:type_name -> svc.biz.vip.MemberLevelInfo
	15, // 1: svc.biz.vip.MGetMemberLevelResp.items:type_name -> svc.biz.vip.MGetMemberLevelResp.ItemsEntry
	16, // 2: svc.biz.vip.LevelInfo.created_at:type_name -> google.protobuf.Timestamp
	16, // 3: svc.biz.vip.LevelInfo.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 4: svc.biz.vip.GetLevelListResp.items:type_name -> svc.biz.vip.LevelInfo
	5,  // 5: svc.biz.vip.AddLevelReq.level_info:type_name -> svc.biz.vip.LevelInfo
	5,  // 6: svc.biz.vip.UpdateLevelReq.level_info:type_name -> svc.biz.vip.LevelInfo
	0,  // 7: svc.biz.vip.MGetMemberLevelResp.ItemsEntry.value:type_name -> svc.biz.vip.MemberLevelInfo
	1,  // 8: svc.biz.vip.Level.GetMemberLevel:input_type -> svc.biz.vip.GetMemberLevelReq
	3,  // 9: svc.biz.vip.Level.MGetMemberLevel:input_type -> svc.biz.vip.MGetMemberLevelReq
	6,  // 10: svc.biz.vip.Level.GetAllLevelList:input_type -> svc.biz.vip.GetAllLevelListReq
	7,  // 11: svc.biz.vip.Level.GetLevelList:input_type -> svc.biz.vip.GetLevelListReq
	9,  // 12: svc.biz.vip.Level.AddLevel:input_type -> svc.biz.vip.AddLevelReq
	11, // 13: svc.biz.vip.Level.UpdateLevel:input_type -> svc.biz.vip.UpdateLevelReq
	13, // 14: svc.biz.vip.Level.DelLevel:input_type -> svc.biz.vip.DelLevelReq
	2,  // 15: svc.biz.vip.Level.GetMemberLevel:output_type -> svc.biz.vip.GetMemberLevelResp
	4,  // 16: svc.biz.vip.Level.MGetMemberLevel:output_type -> svc.biz.vip.MGetMemberLevelResp
	8,  // 17: svc.biz.vip.Level.GetAllLevelList:output_type -> svc.biz.vip.GetLevelListResp
	8,  // 18: svc.biz.vip.Level.GetLevelList:output_type -> svc.biz.vip.GetLevelListResp
	10, // 19: svc.biz.vip.Level.AddLevel:output_type -> svc.biz.vip.AddLevelResp
	12, // 20: svc.biz.vip.Level.UpdateLevel:output_type -> svc.biz.vip.UpdateLevelReqResp
	14, // 21: svc.biz.vip.Level.DelLevel:output_type -> svc.biz.vip.DelLevelResp
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_svc_biz_vip_level_proto_init() }
func file_svc_biz_vip_level_proto_init() {
	if File_svc_biz_vip_level_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_biz_vip_level_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_vip_level_proto_goTypes,
		DependencyIndexes: file_svc_biz_vip_level_proto_depIdxs,
		MessageInfos:      file_svc_biz_vip_level_proto_msgTypes,
	}.Build()
	File_svc_biz_vip_level_proto = out.File
	file_svc_biz_vip_level_proto_rawDesc = nil
	file_svc_biz_vip_level_proto_goTypes = nil
	file_svc_biz_vip_level_proto_depIdxs = nil
}
