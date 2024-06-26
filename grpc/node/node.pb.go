// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: node.proto

package node

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReplicateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename      string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	OtherNodeIp   string `protobuf:"bytes,3,opt,name=other_node_ip,json=otherNodeIp,proto3" json:"other_node_ip,omitempty"`
	OtherNodePort uint32 `protobuf:"varint,4,opt,name=other_node_port,json=otherNodePort,proto3" json:"other_node_port,omitempty"`
}

func (x *ReplicateRequest) Reset() {
	*x = ReplicateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicateRequest) ProtoMessage() {}

func (x *ReplicateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicateRequest.ProtoReflect.Descriptor instead.
func (*ReplicateRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *ReplicateRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ReplicateRequest) GetOtherNodeIp() string {
	if x != nil {
		return x.OtherNodeIp
	}
	return ""
}

func (x *ReplicateRequest) GetOtherNodePort() uint32 {
	if x != nil {
		return x.OtherNodePort
	}
	return 0
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7a, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x32, 0x43, 0x0a, 0x04, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x23, 0x5a, 0x21, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_node_proto_goTypes = []interface{}{
	(*ReplicateRequest)(nil), // 0: node.ReplicateRequest
	(*emptypb.Empty)(nil),    // 1: google.protobuf.Empty
}
var file_node_proto_depIdxs = []int32{
	0, // 0: node.Node.Replicate:input_type -> node.ReplicateRequest
	1, // 1: node.Node.Replicate:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicateRequest); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}
