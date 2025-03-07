// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: svc.biz.vip/noble_salarys.proto

package vip

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NOBLE_SALARY_STATUS int32

const (
	NOBLE_SALARY_STATUS_NOBLE_SALARY_STATUS_UNKNOWN  NOBLE_SALARY_STATUS = 0 // 未知
	NOBLE_SALARY_STATUS_NOBLE_SALARY_STATUS_PENDING  NOBLE_SALARY_STATUS = 1 // 待领取
	NOBLE_SALARY_STATUS_NOBLE_SALARY_STATUS_EXPIRED  NOBLE_SALARY_STATUS = 2 // 已过期
	NOBLE_SALARY_STATUS_NOBLE_SALARY_STATUS_RECEIVED NOBLE_SALARY_STATUS = 3 // 已领取
)

// Enum value maps for NOBLE_SALARY_STATUS.
var (
	NOBLE_SALARY_STATUS_name = map[int32]string{
		0: "NOBLE_SALARY_STATUS_UNKNOWN",
		1: "NOBLE_SALARY_STATUS_PENDING",
		2: "NOBLE_SALARY_STATUS_EXPIRED",
		3: "NOBLE_SALARY_STATUS_RECEIVED",
	}
	NOBLE_SALARY_STATUS_value = map[string]int32{
		"NOBLE_SALARY_STATUS_UNKNOWN":  0,
		"NOBLE_SALARY_STATUS_PENDING":  1,
		"NOBLE_SALARY_STATUS_EXPIRED":  2,
		"NOBLE_SALARY_STATUS_RECEIVED": 3,
	}
)

func (x NOBLE_SALARY_STATUS) Enum() *NOBLE_SALARY_STATUS {
	p := new(NOBLE_SALARY_STATUS)
	*p = x
	return p
}

func (x NOBLE_SALARY_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NOBLE_SALARY_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_biz_vip_noble_salarys_proto_enumTypes[0].Descriptor()
}

func (NOBLE_SALARY_STATUS) Type() protoreflect.EnumType {
	return &file_svc_biz_vip_noble_salarys_proto_enumTypes[0]
}

func (x NOBLE_SALARY_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NOBLE_SALARY_STATUS.Descriptor instead.
func (NOBLE_SALARY_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{0}
}

type NobleSalaryInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                               // 俸禄id
	Amount        int32                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`                                      // 俸禄金额
	Status        NOBLE_SALARY_STATUS    `protobuf:"varint,3,opt,name=status,proto3,enum=svc.biz.vip.NOBLE_SALARY_STATUS" json:"status,omitempty"` // 领取状态
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                // 发放时间
	ReceivedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`             // 领取时间
	ExpiredAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`                // 过期时间
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NobleSalaryInfo) Reset() {
	*x = NobleSalaryInfo{}
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NobleSalaryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NobleSalaryInfo) ProtoMessage() {}

func (x *NobleSalaryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NobleSalaryInfo.ProtoReflect.Descriptor instead.
func (*NobleSalaryInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{0}
}

func (x *NobleSalaryInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NobleSalaryInfo) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *NobleSalaryInfo) GetStatus() NOBLE_SALARY_STATUS {
	if x != nil {
		return x.Status
	}
	return NOBLE_SALARY_STATUS_NOBLE_SALARY_STATUS_UNKNOWN
}

func (x *NobleSalaryInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *NobleSalaryInfo) GetReceivedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedAt
	}
	return nil
}

func (x *NobleSalaryInfo) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

// 空请求类型
type EmptyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{1}
}

// 空响应类型
type EmptyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{2}
}

type NobleSalaryInfoResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Info          *NobleSalaryInfo       `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NobleSalaryInfoResp) Reset() {
	*x = NobleSalaryInfoResp{}
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NobleSalaryInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NobleSalaryInfoResp) ProtoMessage() {}

func (x *NobleSalaryInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NobleSalaryInfoResp.ProtoReflect.Descriptor instead.
func (*NobleSalaryInfoResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{3}
}

func (x *NobleSalaryInfoResp) GetInfo() *NobleSalaryInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type DistributeSalaryResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int32                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`                     // 发放数
	RunTime       float32                `protobuf:"fixed32,2,opt,name=run_time,json=runTime,proto3" json:"run_time,omitempty"` // 执行时间
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DistributeSalaryResp) Reset() {
	*x = DistributeSalaryResp{}
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DistributeSalaryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributeSalaryResp) ProtoMessage() {}

func (x *DistributeSalaryResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributeSalaryResp.ProtoReflect.Descriptor instead.
func (*DistributeSalaryResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{4}
}

func (x *DistributeSalaryResp) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DistributeSalaryResp) GetRunTime() float32 {
	if x != nil {
		return x.RunTime
	}
	return 0
}

type ReceiveSalaryReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReceiveSalaryReq) Reset() {
	*x = ReceiveSalaryReq{}
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReceiveSalaryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveSalaryReq) ProtoMessage() {}

func (x *ReceiveSalaryReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveSalaryReq.ProtoReflect.Descriptor instead.
func (*ReceiveSalaryReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{5}
}

func (x *ReceiveSalaryReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReceiveSalaryResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        NOBLE_SALARY_STATUS    `protobuf:"varint,2,opt,name=status,proto3,enum=svc.biz.vip.NOBLE_SALARY_STATUS" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReceiveSalaryResp) Reset() {
	*x = ReceiveSalaryResp{}
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReceiveSalaryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveSalaryResp) ProtoMessage() {}

func (x *ReceiveSalaryResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveSalaryResp.ProtoReflect.Descriptor instead.
func (*ReceiveSalaryResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{6}
}

func (x *ReceiveSalaryResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReceiveSalaryResp) GetStatus() NOBLE_SALARY_STATUS {
	if x != nil {
		return x.Status
	}
	return NOBLE_SALARY_STATUS_NOBLE_SALARY_STATUS_UNKNOWN
}

type NobleSalaryListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ViewerId      string                 `protobuf:"bytes,1,opt,name=viewer_id,json=viewerId,proto3" json:"viewer_id,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NobleSalaryListReq) Reset() {
	*x = NobleSalaryListReq{}
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NobleSalaryListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NobleSalaryListReq) ProtoMessage() {}

func (x *NobleSalaryListReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NobleSalaryListReq.ProtoReflect.Descriptor instead.
func (*NobleSalaryListReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{7}
}

func (x *NobleSalaryListReq) GetViewerId() string {
	if x != nil {
		return x.ViewerId
	}
	return ""
}

func (x *NobleSalaryListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *NobleSalaryListReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type NobleSalaryListResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*NobleSalaryInfo     `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NobleSalaryListResp) Reset() {
	*x = NobleSalaryListResp{}
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NobleSalaryListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NobleSalaryListResp) ProtoMessage() {}

func (x *NobleSalaryListResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_vip_noble_salarys_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NobleSalaryListResp.ProtoReflect.Descriptor instead.
func (*NobleSalaryListResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_vip_noble_salarys_proto_rawDescGZIP(), []int{8}
}

func (x *NobleSalaryListResp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *NobleSalaryListResp) GetItems() []*NobleSalaryInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_svc_biz_vip_noble_salarys_proto protoreflect.FileDescriptor

var file_svc_biz_vip_noble_salarys_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2f, 0x6e, 0x6f,
	0x62, 0x6c, 0x65, 0x5f, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa6, 0x02, 0x0a, 0x0f, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4e, 0x4f, 0x42, 0x4c, 0x45, 0x5f,
	0x53, 0x41, 0x4c, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x4e, 0x6f, 0x62,
	0x6c, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x30, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x62,
	0x6c, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x47, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5d, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76,
	0x69, 0x70, 0x2e, 0x4e, 0x4f, 0x42, 0x4c, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x41, 0x52, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5b,
	0x0a, 0x12, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5f, 0x0a, 0x13, 0x4e,
	0x6f, 0x62, 0x6c, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x9a, 0x01, 0x0a,
	0x13, 0x4e, 0x4f, 0x42, 0x4c, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x4f, 0x42, 0x4c, 0x45, 0x5f, 0x53, 0x41,
	0x4c, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x4f, 0x42, 0x4c, 0x45, 0x5f, 0x53,
	0x41, 0x4c, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x4f, 0x42, 0x4c, 0x45, 0x5f,
	0x53, 0x41, 0x4c, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x58,
	0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x4e, 0x4f, 0x42, 0x4c, 0x45,
	0x5f, 0x53, 0x41, 0x4c, 0x41, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52,
	0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0x03, 0x32, 0xc9, 0x02, 0x0a, 0x0b, 0x4e, 0x6f,
	0x62, 0x6c, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x4c, 0x0a, 0x0a, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69,
	0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70,
	0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x53,
	0x61, 0x6c, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x62, 0x6c, 0x65,
	0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x2e,
	0x4e, 0x6f, 0x62, 0x6c, 0x65, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x76, 0x69, 0x70, 0x3b, 0x76, 0x69, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_svc_biz_vip_noble_salarys_proto_rawDescOnce sync.Once
	file_svc_biz_vip_noble_salarys_proto_rawDescData []byte
)

func file_svc_biz_vip_noble_salarys_proto_rawDescGZIP() []byte {
	file_svc_biz_vip_noble_salarys_proto_rawDescOnce.Do(func() {
		file_svc_biz_vip_noble_salarys_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_svc_biz_vip_noble_salarys_proto_rawDesc), len(file_svc_biz_vip_noble_salarys_proto_rawDesc)))
	})
	return file_svc_biz_vip_noble_salarys_proto_rawDescData
}

var file_svc_biz_vip_noble_salarys_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_svc_biz_vip_noble_salarys_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_svc_biz_vip_noble_salarys_proto_goTypes = []any{
	(NOBLE_SALARY_STATUS)(0),      // 0: svc.biz.vip.NOBLE_SALARY_STATUS
	(*NobleSalaryInfo)(nil),       // 1: svc.biz.vip.NobleSalaryInfo
	(*EmptyRequest)(nil),          // 2: svc.biz.vip.EmptyRequest
	(*EmptyResponse)(nil),         // 3: svc.biz.vip.EmptyResponse
	(*NobleSalaryInfoResp)(nil),   // 4: svc.biz.vip.NobleSalaryInfoResp
	(*DistributeSalaryResp)(nil),  // 5: svc.biz.vip.DistributeSalaryResp
	(*ReceiveSalaryReq)(nil),      // 6: svc.biz.vip.ReceiveSalaryReq
	(*ReceiveSalaryResp)(nil),     // 7: svc.biz.vip.ReceiveSalaryResp
	(*NobleSalaryListReq)(nil),    // 8: svc.biz.vip.NobleSalaryListReq
	(*NobleSalaryListResp)(nil),   // 9: svc.biz.vip.NobleSalaryListResp
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_svc_biz_vip_noble_salarys_proto_depIdxs = []int32{
	0,  // 0: svc.biz.vip.NobleSalaryInfo.status:type_name -> svc.biz.vip.NOBLE_SALARY_STATUS
	10, // 1: svc.biz.vip.NobleSalaryInfo.created_at:type_name -> google.protobuf.Timestamp
	10, // 2: svc.biz.vip.NobleSalaryInfo.received_at:type_name -> google.protobuf.Timestamp
	10, // 3: svc.biz.vip.NobleSalaryInfo.expired_at:type_name -> google.protobuf.Timestamp
	1,  // 4: svc.biz.vip.NobleSalaryInfoResp.info:type_name -> svc.biz.vip.NobleSalaryInfo
	0,  // 5: svc.biz.vip.ReceiveSalaryResp.status:type_name -> svc.biz.vip.NOBLE_SALARY_STATUS
	1,  // 6: svc.biz.vip.NobleSalaryListResp.items:type_name -> svc.biz.vip.NobleSalaryInfo
	2,  // 7: svc.biz.vip.NobleSalary.Distribute:input_type -> svc.biz.vip.EmptyRequest
	6,  // 8: svc.biz.vip.NobleSalary.Receive:input_type -> svc.biz.vip.ReceiveSalaryReq
	8,  // 9: svc.biz.vip.NobleSalary.List:input_type -> svc.biz.vip.NobleSalaryListReq
	6,  // 10: svc.biz.vip.NobleSalary.GetReceiveInfo:input_type -> svc.biz.vip.ReceiveSalaryReq
	5,  // 11: svc.biz.vip.NobleSalary.Distribute:output_type -> svc.biz.vip.DistributeSalaryResp
	7,  // 12: svc.biz.vip.NobleSalary.Receive:output_type -> svc.biz.vip.ReceiveSalaryResp
	9,  // 13: svc.biz.vip.NobleSalary.List:output_type -> svc.biz.vip.NobleSalaryListResp
	4,  // 14: svc.biz.vip.NobleSalary.GetReceiveInfo:output_type -> svc.biz.vip.NobleSalaryInfoResp
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_svc_biz_vip_noble_salarys_proto_init() }
func file_svc_biz_vip_noble_salarys_proto_init() {
	if File_svc_biz_vip_noble_salarys_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_svc_biz_vip_noble_salarys_proto_rawDesc), len(file_svc_biz_vip_noble_salarys_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_vip_noble_salarys_proto_goTypes,
		DependencyIndexes: file_svc_biz_vip_noble_salarys_proto_depIdxs,
		EnumInfos:         file_svc_biz_vip_noble_salarys_proto_enumTypes,
		MessageInfos:      file_svc_biz_vip_noble_salarys_proto_msgTypes,
	}.Build()
	File_svc_biz_vip_noble_salarys_proto = out.File
	file_svc_biz_vip_noble_salarys_proto_goTypes = nil
	file_svc_biz_vip_noble_salarys_proto_depIdxs = nil
}
