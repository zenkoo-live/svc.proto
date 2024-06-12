// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: svc.infra.center/instruction.proto

package center

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RemoveSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session string `protobuf:"bytes,15,opt,name=session,proto3" json:"session,omitempty"` // 连接Session
}

func (x *RemoveSessionRequest) Reset() {
	*x = RemoveSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_center_instruction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSessionRequest) ProtoMessage() {}

func (x *RemoveSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_center_instruction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSessionRequest.ProtoReflect.Descriptor instead.
func (*RemoveSessionRequest) Descriptor() ([]byte, []int) {
	return file_svc_infra_center_instruction_proto_rawDescGZIP(), []int{0}
}

func (x *RemoveSessionRequest) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type RemoveSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"` // 是否成功
}

func (x *RemoveSessionResponse) Reset() {
	*x = RemoveSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_center_instruction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSessionResponse) ProtoMessage() {}

func (x *RemoveSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_center_instruction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSessionResponse.ProtoReflect.Descriptor instead.
func (*RemoveSessionResponse) Descriptor() ([]byte, []int) {
	return file_svc_infra_center_instruction_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveSessionResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type RemoveAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId int64 `protobuf:"varint,16,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` // 账号ID
	GroupId   int64 `protobuf:"varint,17,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`       // 群组ID（可留空）
}

func (x *RemoveAccountRequest) Reset() {
	*x = RemoveAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_center_instruction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAccountRequest) ProtoMessage() {}

func (x *RemoveAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_center_instruction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAccountRequest.ProtoReflect.Descriptor instead.
func (*RemoveAccountRequest) Descriptor() ([]byte, []int) {
	return file_svc_infra_center_instruction_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveAccountRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *RemoveAccountRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type RemoveAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"` // 是否成功（不代表结果）
}

func (x *RemoveAccountResponse) Reset() {
	*x = RemoveAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_center_instruction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAccountResponse) ProtoMessage() {}

func (x *RemoveAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_center_instruction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAccountResponse.ProtoReflect.Descriptor instead.
func (*RemoveAccountResponse) Descriptor() ([]byte, []int) {
	return file_svc_infra_center_instruction_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveAccountResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type RemoveDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device string `protobuf:"bytes,18,opt,name=device,proto3" json:"device,omitempty"` // 设备标识（指纹）
}

func (x *RemoveDeviceRequest) Reset() {
	*x = RemoveDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_center_instruction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDeviceRequest) ProtoMessage() {}

func (x *RemoveDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_center_instruction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDeviceRequest.ProtoReflect.Descriptor instead.
func (*RemoveDeviceRequest) Descriptor() ([]byte, []int) {
	return file_svc_infra_center_instruction_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveDeviceRequest) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

type RemoveDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"` // 是否成功
}

func (x *RemoveDeviceResponse) Reset() {
	*x = RemoveDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_center_instruction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDeviceResponse) ProtoMessage() {}

func (x *RemoveDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_center_instruction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDeviceResponse.ProtoReflect.Descriptor instead.
func (*RemoveDeviceResponse) Descriptor() ([]byte, []int) {
	return file_svc_infra_center_instruction_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveDeviceResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_svc_infra_center_instruction_proto protoreflect.FileDescriptor

var file_svc_infra_center_instruction_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6b, 0x22, 0x50, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x2d, 0x0a, 0x13,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x32, 0xba, 0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0d, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5f, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x25, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_infra_center_instruction_proto_rawDescOnce sync.Once
	file_svc_infra_center_instruction_proto_rawDescData = file_svc_infra_center_instruction_proto_rawDesc
)

func file_svc_infra_center_instruction_proto_rawDescGZIP() []byte {
	file_svc_infra_center_instruction_proto_rawDescOnce.Do(func() {
		file_svc_infra_center_instruction_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_infra_center_instruction_proto_rawDescData)
	})
	return file_svc_infra_center_instruction_proto_rawDescData
}

var file_svc_infra_center_instruction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_svc_infra_center_instruction_proto_goTypes = []interface{}{
	(*RemoveSessionRequest)(nil),  // 0: svc.infra.center.RemoveSessionRequest
	(*RemoveSessionResponse)(nil), // 1: svc.infra.center.RemoveSessionResponse
	(*RemoveAccountRequest)(nil),  // 2: svc.infra.center.RemoveAccountRequest
	(*RemoveAccountResponse)(nil), // 3: svc.infra.center.RemoveAccountResponse
	(*RemoveDeviceRequest)(nil),   // 4: svc.infra.center.RemoveDeviceRequest
	(*RemoveDeviceResponse)(nil),  // 5: svc.infra.center.RemoveDeviceResponse
}
var file_svc_infra_center_instruction_proto_depIdxs = []int32{
	0, // 0: svc.infra.center.LinkInstruction.RemoveSession:input_type -> svc.infra.center.RemoveSessionRequest
	2, // 1: svc.infra.center.LinkInstruction.RemoveAccount:input_type -> svc.infra.center.RemoveAccountRequest
	4, // 2: svc.infra.center.LinkInstruction.RemoveDevice:input_type -> svc.infra.center.RemoveDeviceRequest
	1, // 3: svc.infra.center.LinkInstruction.RemoveSession:output_type -> svc.infra.center.RemoveSessionResponse
	3, // 4: svc.infra.center.LinkInstruction.RemoveAccount:output_type -> svc.infra.center.RemoveAccountResponse
	5, // 5: svc.infra.center.LinkInstruction.RemoveDevice:output_type -> svc.infra.center.RemoveDeviceResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_infra_center_instruction_proto_init() }
func file_svc_infra_center_instruction_proto_init() {
	if File_svc_infra_center_instruction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_infra_center_instruction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSessionRequest); i {
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
		file_svc_infra_center_instruction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSessionResponse); i {
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
		file_svc_infra_center_instruction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAccountRequest); i {
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
		file_svc_infra_center_instruction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAccountResponse); i {
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
		file_svc_infra_center_instruction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDeviceRequest); i {
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
		file_svc_infra_center_instruction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDeviceResponse); i {
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
			RawDescriptor: file_svc_infra_center_instruction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_infra_center_instruction_proto_goTypes,
		DependencyIndexes: file_svc_infra_center_instruction_proto_depIdxs,
		MessageInfos:      file_svc_infra_center_instruction_proto_msgTypes,
	}.Build()
	File_svc_infra_center_instruction_proto = out.File
	file_svc_infra_center_instruction_proto_rawDesc = nil
	file_svc_infra_center_instruction_proto_goTypes = nil
	file_svc_infra_center_instruction_proto_depIdxs = nil
}