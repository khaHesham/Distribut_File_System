// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: master_client.proto

package master_client

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

var File_master_client_proto protoreflect.FileDescriptor

var file_master_client_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x32, 0x16, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x2c, 0x5a, 0x2a,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_master_client_proto_goTypes = []interface{}{}
var file_master_client_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_master_client_proto_init() }
func file_master_client_proto_init() {
	if File_master_client_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_master_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_master_client_proto_goTypes,
		DependencyIndexes: file_master_client_proto_depIdxs,
	}.Build()
	File_master_client_proto = out.File
	file_master_client_proto_rawDesc = nil
	file_master_client_proto_goTypes = nil
	file_master_client_proto_depIdxs = nil
}
