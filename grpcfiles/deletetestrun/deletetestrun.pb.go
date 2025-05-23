// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.27.0
// source: deletetestrun.proto

package deletetestrun

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

// Request and Response messages
type DeleteTestRunRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Test run ID as a string
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTestRunRequest) Reset() {
	*x = DeleteTestRunRequest{}
	mi := &file_deletetestrun_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTestRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestRunRequest) ProtoMessage() {}

func (x *DeleteTestRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deletetestrun_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestRunRequest.ProtoReflect.Descriptor instead.
func (*DeleteTestRunRequest) Descriptor() ([]byte, []int) {
	return file_deletetestrun_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteTestRunRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTestRunResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTestRunResponse) Reset() {
	*x = DeleteTestRunResponse{}
	mi := &file_deletetestrun_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTestRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestRunResponse) ProtoMessage() {}

func (x *DeleteTestRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deletetestrun_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestRunResponse.ProtoReflect.Descriptor instead.
func (*DeleteTestRunResponse) Descriptor() ([]byte, []int) {
	return file_deletetestrun_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteTestRunResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteTestRunResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_deletetestrun_proto protoreflect.FileDescriptor

var file_deletetestrun_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x72, 0x75, 0x6e, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x6c, 0x0a, 0x0e, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x23, 0x2e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75,
	0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x3b, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x74, 0x65, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_deletetestrun_proto_rawDescOnce sync.Once
	file_deletetestrun_proto_rawDescData = file_deletetestrun_proto_rawDesc
)

func file_deletetestrun_proto_rawDescGZIP() []byte {
	file_deletetestrun_proto_rawDescOnce.Do(func() {
		file_deletetestrun_proto_rawDescData = protoimpl.X.CompressGZIP(file_deletetestrun_proto_rawDescData)
	})
	return file_deletetestrun_proto_rawDescData
}

var file_deletetestrun_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_deletetestrun_proto_goTypes = []any{
	(*DeleteTestRunRequest)(nil),  // 0: deletetestrun.DeleteTestRunRequest
	(*DeleteTestRunResponse)(nil), // 1: deletetestrun.DeleteTestRunResponse
}
var file_deletetestrun_proto_depIdxs = []int32{
	0, // 0: deletetestrun.TestRunService.DeleteTestRun:input_type -> deletetestrun.DeleteTestRunRequest
	1, // 1: deletetestrun.TestRunService.DeleteTestRun:output_type -> deletetestrun.DeleteTestRunResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deletetestrun_proto_init() }
func file_deletetestrun_proto_init() {
	if File_deletetestrun_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deletetestrun_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deletetestrun_proto_goTypes,
		DependencyIndexes: file_deletetestrun_proto_depIdxs,
		MessageInfos:      file_deletetestrun_proto_msgTypes,
	}.Build()
	File_deletetestrun_proto = out.File
	file_deletetestrun_proto_rawDesc = nil
	file_deletetestrun_proto_goTypes = nil
	file_deletetestrun_proto_depIdxs = nil
}
