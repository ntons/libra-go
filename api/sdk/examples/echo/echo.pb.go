// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: sdk/examples/echo/echo.proto

package echo

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_examples_echo_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_examples_echo_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_sdk_examples_echo_echo_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_examples_echo_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_examples_echo_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_sdk_examples_echo_echo_proto_rawDescGZIP(), []int{1}
}

func (x *SendResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_sdk_examples_echo_echo_proto protoreflect.FileDescriptor

var file_sdk_examples_echo_echo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65,
	0x63, 0x68, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x28, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x5b, 0x0a, 0x04, 0x45, 0x63,
	0x68, 0x6f, 0x12, 0x53, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x24, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sdk_examples_echo_echo_proto_rawDescOnce sync.Once
	file_sdk_examples_echo_echo_proto_rawDescData = file_sdk_examples_echo_echo_proto_rawDesc
)

func file_sdk_examples_echo_echo_proto_rawDescGZIP() []byte {
	file_sdk_examples_echo_echo_proto_rawDescOnce.Do(func() {
		file_sdk_examples_echo_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_examples_echo_echo_proto_rawDescData)
	})
	return file_sdk_examples_echo_echo_proto_rawDescData
}

var file_sdk_examples_echo_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sdk_examples_echo_echo_proto_goTypes = []interface{}{
	(*SendRequest)(nil),  // 0: libra.sdk.examples.echo.SendRequest
	(*SendResponse)(nil), // 1: libra.sdk.examples.echo.SendResponse
}
var file_sdk_examples_echo_echo_proto_depIdxs = []int32{
	0, // 0: libra.sdk.examples.echo.Echo.Send:input_type -> libra.sdk.examples.echo.SendRequest
	1, // 1: libra.sdk.examples.echo.Echo.Send:output_type -> libra.sdk.examples.echo.SendResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sdk_examples_echo_echo_proto_init() }
func file_sdk_examples_echo_echo_proto_init() {
	if File_sdk_examples_echo_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_examples_echo_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_sdk_examples_echo_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
			RawDescriptor: file_sdk_examples_echo_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sdk_examples_echo_echo_proto_goTypes,
		DependencyIndexes: file_sdk_examples_echo_echo_proto_depIdxs,
		MessageInfos:      file_sdk_examples_echo_echo_proto_msgTypes,
	}.Build()
	File_sdk_examples_echo_echo_proto = out.File
	file_sdk_examples_echo_echo_proto_rawDesc = nil
	file_sdk_examples_echo_echo_proto_goTypes = nil
	file_sdk_examples_echo_echo_proto_depIdxs = nil
}
