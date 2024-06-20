// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: svc.infra.link/trace.proto

package link

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

type StartTraceRequest_TracerType int32

const (
	StartTraceRequest_TRACER_ACCOUNT StartTraceRequest_TracerType = 0
	StartTraceRequest_TRACER_SESSION StartTraceRequest_TracerType = 1
	StartTraceRequest_TRACER_DEVICE  StartTraceRequest_TracerType = 2
	StartTraceRequest_TRACER_REMOTE  StartTraceRequest_TracerType = 3
	StartTraceRequest_TRACER_GLOBAL  StartTraceRequest_TracerType = 4
)

// Enum value maps for StartTraceRequest_TracerType.
var (
	StartTraceRequest_TracerType_name = map[int32]string{
		0: "TRACER_ACCOUNT",
		1: "TRACER_SESSION",
		2: "TRACER_DEVICE",
		3: "TRACER_REMOTE",
		4: "TRACER_GLOBAL",
	}
	StartTraceRequest_TracerType_value = map[string]int32{
		"TRACER_ACCOUNT": 0,
		"TRACER_SESSION": 1,
		"TRACER_DEVICE":  2,
		"TRACER_REMOTE":  3,
		"TRACER_GLOBAL":  4,
	}
)

func (x StartTraceRequest_TracerType) Enum() *StartTraceRequest_TracerType {
	p := new(StartTraceRequest_TracerType)
	*p = x
	return p
}

func (x StartTraceRequest_TracerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StartTraceRequest_TracerType) Descriptor() protoreflect.EnumDescriptor {
	return file_svc_infra_link_trace_proto_enumTypes[0].Descriptor()
}

func (StartTraceRequest_TracerType) Type() protoreflect.EnumType {
	return &file_svc_infra_link_trace_proto_enumTypes[0]
}

func (x StartTraceRequest_TracerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StartTraceRequest_TracerType.Descriptor instead.
func (StartTraceRequest_TracerType) EnumDescriptor() ([]byte, []int) {
	return file_svc_infra_link_trace_proto_rawDescGZIP(), []int{0, 0}
}

type StartTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TracerType StartTraceRequest_TracerType `protobuf:"varint,1,opt,name=tracer_type,json=tracerType,proto3,enum=svc.infra.link.StartTraceRequest_TracerType" json:"tracer_type,omitempty"`
	SessionId  string                       `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`    // 连接ID
	Device     string                       `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`                           // 设备标识
	AccountId  string                       `protobuf:"bytes,4,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`    // 账号ID
	RemoteAddr string                       `protobuf:"bytes,5,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"` // 远程地址
	Duration   int64                        `protobuf:"varint,127,opt,name=duration,proto3" json:"duration,omitempty"`                    // 持续时长
}

func (x *StartTraceRequest) Reset() {
	*x = StartTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_trace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTraceRequest) ProtoMessage() {}

func (x *StartTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_trace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTraceRequest.ProtoReflect.Descriptor instead.
func (*StartTraceRequest) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_trace_proto_rawDescGZIP(), []int{0}
}

func (x *StartTraceRequest) GetTracerType() StartTraceRequest_TracerType {
	if x != nil {
		return x.TracerType
	}
	return StartTraceRequest_TRACER_ACCOUNT
}

func (x *StartTraceRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *StartTraceRequest) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *StartTraceRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *StartTraceRequest) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *StartTraceRequest) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type StartTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`                           // 状态
	TraceId string `protobuf:"bytes,127,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"` // Trace标识
}

func (x *StartTraceResponse) Reset() {
	*x = StartTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_trace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTraceResponse) ProtoMessage() {}

func (x *StartTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_trace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTraceResponse.ProtoReflect.Descriptor instead.
func (*StartTraceResponse) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_trace_proto_rawDescGZIP(), []int{1}
}

func (x *StartTraceResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *StartTraceResponse) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

type StopTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopTraceRequest) Reset() {
	*x = StopTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_trace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTraceRequest) ProtoMessage() {}

func (x *StopTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_trace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTraceRequest.ProtoReflect.Descriptor instead.
func (*StopTraceRequest) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_trace_proto_rawDescGZIP(), []int{2}
}

type StopTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"` // 状态
}

func (x *StopTraceResponse) Reset() {
	*x = StopTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_infra_link_trace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTraceResponse) ProtoMessage() {}

func (x *StopTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_link_trace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTraceResponse.ProtoReflect.Descriptor instead.
func (*StopTraceResponse) Descriptor() ([]byte, []int) {
	return file_svc_infra_link_trace_proto_rawDescGZIP(), []int{3}
}

func (x *StopTraceResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_svc_infra_link_trace_proto protoreflect.FileDescriptor

var file_svc_infra_link_trace_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x76,
	0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xe4, 0x02, 0x0a,
	0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x7f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x52, 0x41, 0x43, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x52, 0x41, 0x43, 0x45, 0x52,
	0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x52,
	0x41, 0x43, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x54, 0x52, 0x41, 0x43, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x10, 0x03,
	0x12, 0x11, 0x0a, 0x0d, 0x54, 0x52, 0x41, 0x43, 0x45, 0x52, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41,
	0x4c, 0x10, 0x04, 0x22, 0x3f, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x7f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x70,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xac, 0x01,
	0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a,
	0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x20, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15,
	0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x3b, 0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_infra_link_trace_proto_rawDescOnce sync.Once
	file_svc_infra_link_trace_proto_rawDescData = file_svc_infra_link_trace_proto_rawDesc
)

func file_svc_infra_link_trace_proto_rawDescGZIP() []byte {
	file_svc_infra_link_trace_proto_rawDescOnce.Do(func() {
		file_svc_infra_link_trace_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_infra_link_trace_proto_rawDescData)
	})
	return file_svc_infra_link_trace_proto_rawDescData
}

var file_svc_infra_link_trace_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_svc_infra_link_trace_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_svc_infra_link_trace_proto_goTypes = []any{
	(StartTraceRequest_TracerType)(0), // 0: svc.infra.link.StartTraceRequest.TracerType
	(*StartTraceRequest)(nil),         // 1: svc.infra.link.StartTraceRequest
	(*StartTraceResponse)(nil),        // 2: svc.infra.link.StartTraceResponse
	(*StopTraceRequest)(nil),          // 3: svc.infra.link.StopTraceRequest
	(*StopTraceResponse)(nil),         // 4: svc.infra.link.StopTraceResponse
}
var file_svc_infra_link_trace_proto_depIdxs = []int32{
	0, // 0: svc.infra.link.StartTraceRequest.tracer_type:type_name -> svc.infra.link.StartTraceRequest.TracerType
	1, // 1: svc.infra.link.LinkTrace.Start:input_type -> svc.infra.link.StartTraceRequest
	3, // 2: svc.infra.link.LinkTrace.Stop:input_type -> svc.infra.link.StopTraceRequest
	2, // 3: svc.infra.link.LinkTrace.Start:output_type -> svc.infra.link.StartTraceResponse
	4, // 4: svc.infra.link.LinkTrace.Stop:output_type -> svc.infra.link.StopTraceResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_svc_infra_link_trace_proto_init() }
func file_svc_infra_link_trace_proto_init() {
	if File_svc_infra_link_trace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_infra_link_trace_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StartTraceRequest); i {
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
		file_svc_infra_link_trace_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StartTraceResponse); i {
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
		file_svc_infra_link_trace_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StopTraceRequest); i {
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
		file_svc_infra_link_trace_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StopTraceResponse); i {
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
			RawDescriptor: file_svc_infra_link_trace_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_infra_link_trace_proto_goTypes,
		DependencyIndexes: file_svc_infra_link_trace_proto_depIdxs,
		EnumInfos:         file_svc_infra_link_trace_proto_enumTypes,
		MessageInfos:      file_svc_infra_link_trace_proto_msgTypes,
	}.Build()
	File_svc_infra_link_trace_proto = out.File
	file_svc_infra_link_trace_proto_rawDesc = nil
	file_svc_infra_link_trace_proto_goTypes = nil
	file_svc_infra_link_trace_proto_depIdxs = nil
}
