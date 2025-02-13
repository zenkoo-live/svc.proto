// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: svc.infra.static/static.proto

package static

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type InitDBResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        bool                   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitDBResp) Reset() {
	*x = InitDBResp{}
	mi := &file_svc_infra_static_static_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitDBResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDBResp) ProtoMessage() {}

func (x *InitDBResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[0]
	if x != nil {
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

type UploadRequestMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Binary        []byte                 `protobuf:"bytes,2,opt,name=binary,proto3" json:"binary,omitempty"`
	PreToken      bool                   `protobuf:"varint,3,opt,name=pre_token,json=preToken,proto3" json:"pre_token,omitempty"`
	Bucket        string                 `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	UserId        string                 `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadRequestMessage) Reset() {
	*x = UploadRequestMessage{}
	mi := &file_svc_infra_static_static_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequestMessage) ProtoMessage() {}

func (x *UploadRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[1]
	if x != nil {
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
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{1}
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

func (x *UploadRequestMessage) GetPreToken() bool {
	if x != nil {
		return x.PreToken
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

type UploadResponseMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Path          string                 `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Domain        string                 `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadResponseMessage) Reset() {
	*x = UploadResponseMessage{}
	mi := &file_svc_infra_static_static_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponseMessage) ProtoMessage() {}

func (x *UploadResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_svc_infra_static_static_proto_msgTypes[2]
	if x != nil {
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
	return file_svc_infra_static_static_proto_rawDescGZIP(), []int{2}
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

var File_svc_infra_static_static_proto protoreflect.FileDescriptor

var file_svc_infra_static_static_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24,
	0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x72,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x15, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x32, 0xc9, 0x03, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x3e, 0x0a, 0x06, 0x49, 0x6e, 0x69, 0x74, 0x44,
	0x42, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5f, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x73, 0x76,
	0x63, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x3b, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_svc_infra_static_static_proto_rawDescOnce sync.Once
	file_svc_infra_static_static_proto_rawDescData []byte
)

func file_svc_infra_static_static_proto_rawDescGZIP() []byte {
	file_svc_infra_static_static_proto_rawDescOnce.Do(func() {
		file_svc_infra_static_static_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_svc_infra_static_static_proto_rawDesc), len(file_svc_infra_static_static_proto_rawDesc)))
	})
	return file_svc_infra_static_static_proto_rawDescData
}

var file_svc_infra_static_static_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_svc_infra_static_static_proto_goTypes = []any{
	(*InitDBResp)(nil),            // 0: svc.infra.static.InitDBResp
	(*UploadRequestMessage)(nil),  // 1: svc.infra.static.UploadRequestMessage
	(*UploadResponseMessage)(nil), // 2: svc.infra.static.UploadResponseMessage
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_svc_infra_static_static_proto_depIdxs = []int32{
	3, // 0: svc.infra.static.Static.InitDB:input_type -> google.protobuf.Empty
	1, // 1: svc.infra.static.Static.UploadAvatar:input_type -> svc.infra.static.UploadRequestMessage
	1, // 2: svc.infra.static.Static.UploadCover:input_type -> svc.infra.static.UploadRequestMessage
	1, // 3: svc.infra.static.Static.UploadVideo:input_type -> svc.infra.static.UploadRequestMessage
	1, // 4: svc.infra.static.Static.UploadImage:input_type -> svc.infra.static.UploadRequestMessage
	0, // 5: svc.infra.static.Static.InitDB:output_type -> svc.infra.static.InitDBResp
	2, // 6: svc.infra.static.Static.UploadAvatar:output_type -> svc.infra.static.UploadResponseMessage
	2, // 7: svc.infra.static.Static.UploadCover:output_type -> svc.infra.static.UploadResponseMessage
	2, // 8: svc.infra.static.Static.UploadVideo:output_type -> svc.infra.static.UploadResponseMessage
	2, // 9: svc.infra.static.Static.UploadImage:output_type -> svc.infra.static.UploadResponseMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_infra_static_static_proto_init() }
func file_svc_infra_static_static_proto_init() {
	if File_svc_infra_static_static_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_svc_infra_static_static_proto_rawDesc), len(file_svc_infra_static_static_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_infra_static_static_proto_goTypes,
		DependencyIndexes: file_svc_infra_static_static_proto_depIdxs,
		MessageInfos:      file_svc_infra_static_static_proto_msgTypes,
	}.Build()
	File_svc_infra_static_static_proto = out.File
	file_svc_infra_static_static_proto_goTypes = nil
	file_svc_infra_static_static_proto_depIdxs = nil
}
