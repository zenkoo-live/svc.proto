// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: svc.web.dashboard/dashboard.proto

package dashboard

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_svc_web_dashboard_dashboard_proto protoreflect.FileDescriptor

var file_svc_web_dashboard_dashboard_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x76, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x76, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x32, 0x0b, 0x0a, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x77, 0x65, 0x62,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x3b, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_svc_web_dashboard_dashboard_proto_goTypes = []any{}
var file_svc_web_dashboard_dashboard_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_web_dashboard_dashboard_proto_init() }
func file_svc_web_dashboard_dashboard_proto_init() {
	if File_svc_web_dashboard_dashboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_web_dashboard_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_web_dashboard_dashboard_proto_goTypes,
		DependencyIndexes: file_svc_web_dashboard_dashboard_proto_depIdxs,
	}.Build()
	File_svc_web_dashboard_dashboard_proto = out.File
	file_svc_web_dashboard_dashboard_proto_rawDesc = nil
	file_svc_web_dashboard_dashboard_proto_goTypes = nil
	file_svc_web_dashboard_dashboard_proto_depIdxs = nil
}
