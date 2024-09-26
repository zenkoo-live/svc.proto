// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: svc.biz.log/log.proto

package log

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

// LogInfo 日志详情
type LogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogId       string                 `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Object      string                 `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`                              // 操作对象
	ObjectId    string                 `protobuf:"bytes,3,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`          // 操作对象uuid
	Action      string                 `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`                              // 操作行为
	Operator    string                 `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`                          // 操作人
	OperateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=operate_time,json=operateTime,proto3" json:"operate_time,omitempty"` // 操作时间
	Extra       string                 `protobuf:"bytes,7,opt,name=extra,proto3" json:"extra,omitempty"`                                // 扩展信息,爱存啥存啥
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,255,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *LogInfo) Reset() {
	*x = LogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_log_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInfo) ProtoMessage() {}

func (x *LogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_log_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInfo.ProtoReflect.Descriptor instead.
func (*LogInfo) Descriptor() ([]byte, []int) {
	return file_svc_biz_log_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogInfo) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *LogInfo) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *LogInfo) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *LogInfo) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *LogInfo) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *LogInfo) GetOperateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.OperateTime
	}
	return nil
}

func (x *LogInfo) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *LogInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type MGetLastLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items map[string]*LogInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MGetLastLogResp) Reset() {
	*x = MGetLastLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_log_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MGetLastLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetLastLogResp) ProtoMessage() {}

func (x *MGetLastLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_log_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetLastLogResp.ProtoReflect.Descriptor instead.
func (*MGetLastLogResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_log_log_proto_rawDescGZIP(), []int{1}
}

func (x *MGetLastLogResp) GetItems() map[string]*LogInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type AddLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogInfo *LogInfo `protobuf:"bytes,1,opt,name=log_info,json=logInfo,proto3" json:"log_info,omitempty"`
}

func (x *AddLogReq) Reset() {
	*x = AddLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_log_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLogReq) ProtoMessage() {}

func (x *AddLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_log_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLogReq.ProtoReflect.Descriptor instead.
func (*AddLogReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_log_log_proto_rawDescGZIP(), []int{2}
}

func (x *AddLogReq) GetLogInfo() *LogInfo {
	if x != nil {
		return x.LogInfo
	}
	return nil
}

type AddLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddLogResp) Reset() {
	*x = AddLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_log_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLogResp) ProtoMessage() {}

func (x *AddLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_log_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLogResp.ProtoReflect.Descriptor instead.
func (*AddLogResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_log_log_proto_rawDescGZIP(), []int{3}
}

type MGetLastLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object    string   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`                        // 操作对象
	ObjectIds []string `protobuf:"bytes,2,rep,name=object_ids,json=objectIds,proto3" json:"object_ids,omitempty"` // 操作对象uuid
}

func (x *MGetLastLogReq) Reset() {
	*x = MGetLastLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_biz_log_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MGetLastLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetLastLogReq) ProtoMessage() {}

func (x *MGetLastLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_log_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetLastLogReq.ProtoReflect.Descriptor instead.
func (*MGetLastLogReq) Descriptor() ([]byte, []int) {
	return file_svc_biz_log_log_proto_rawDescGZIP(), []int{4}
}

func (x *MGetLastLogReq) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *MGetLastLogReq) GetObjectIds() []string {
	if x != nil {
		return x.ObjectIds
	}
	return nil
}

var File_svc_biz_log_log_proto protoreflect.FileDescriptor

var file_svc_biz_log_log_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a,
	0x2e, 0x6c, 0x6f, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x0f, 0x4d, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e,
	0x6c, 0x6f, 0x67, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x4e, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x12, 0x2f, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x47, 0x0a, 0x0e, 0x4d, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x32, 0x8e, 0x01, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x12, 0x3b, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x0b, 0x4d, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x1b,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4d, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x73, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e,
	0x2f, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x6c, 0x6f, 0x67, 0x3b, 0x6c, 0x6f, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_log_log_proto_rawDescOnce sync.Once
	file_svc_biz_log_log_proto_rawDescData = file_svc_biz_log_log_proto_rawDesc
)

func file_svc_biz_log_log_proto_rawDescGZIP() []byte {
	file_svc_biz_log_log_proto_rawDescOnce.Do(func() {
		file_svc_biz_log_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_log_log_proto_rawDescData)
	})
	return file_svc_biz_log_log_proto_rawDescData
}

var file_svc_biz_log_log_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_svc_biz_log_log_proto_goTypes = []any{
	(*LogInfo)(nil),               // 0: svc.biz.log.LogInfo
	(*MGetLastLogResp)(nil),       // 1: svc.biz.log.MGetLastLogResp
	(*AddLogReq)(nil),             // 2: svc.biz.log.AddLogReq
	(*AddLogResp)(nil),            // 3: svc.biz.log.AddLogResp
	(*MGetLastLogReq)(nil),        // 4: svc.biz.log.MGetLastLogReq
	nil,                           // 5: svc.biz.log.MGetLastLogResp.ItemsEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_svc_biz_log_log_proto_depIdxs = []int32{
	6, // 0: svc.biz.log.LogInfo.operate_time:type_name -> google.protobuf.Timestamp
	6, // 1: svc.biz.log.LogInfo.created_at:type_name -> google.protobuf.Timestamp
	5, // 2: svc.biz.log.MGetLastLogResp.items:type_name -> svc.biz.log.MGetLastLogResp.ItemsEntry
	0, // 3: svc.biz.log.AddLogReq.log_info:type_name -> svc.biz.log.LogInfo
	0, // 4: svc.biz.log.MGetLastLogResp.ItemsEntry.value:type_name -> svc.biz.log.LogInfo
	2, // 5: svc.biz.log.Log.AddLog:input_type -> svc.biz.log.AddLogReq
	4, // 6: svc.biz.log.Log.MGetLastLog:input_type -> svc.biz.log.MGetLastLogReq
	3, // 7: svc.biz.log.Log.AddLog:output_type -> svc.biz.log.AddLogResp
	1, // 8: svc.biz.log.Log.MGetLastLog:output_type -> svc.biz.log.MGetLastLogResp
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_svc_biz_log_log_proto_init() }
func file_svc_biz_log_log_proto_init() {
	if File_svc_biz_log_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_biz_log_log_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LogInfo); i {
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
		file_svc_biz_log_log_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MGetLastLogResp); i {
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
		file_svc_biz_log_log_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddLogReq); i {
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
		file_svc_biz_log_log_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddLogResp); i {
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
		file_svc_biz_log_log_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MGetLastLogReq); i {
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
			RawDescriptor: file_svc_biz_log_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_log_log_proto_goTypes,
		DependencyIndexes: file_svc_biz_log_log_proto_depIdxs,
		MessageInfos:      file_svc_biz_log_log_proto_msgTypes,
	}.Build()
	File_svc_biz_log_log_proto = out.File
	file_svc_biz_log_log_proto_rawDesc = nil
	file_svc_biz_log_log_proto_goTypes = nil
	file_svc_biz_log_log_proto_depIdxs = nil
}
