// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: libra/v1/gateway.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GatewayMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic        string            `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Id           string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ProducerName string            `protobuf:"bytes,3,opt,name=producer_name,json=producerName,proto3" json:"producer_name,omitempty"`
	EventTime    int64             `protobuf:"varint,4,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	PublishTime  int64             `protobuf:"varint,5,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	Properties   map[string]string `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value        *anypb.Any        `protobuf:"bytes,10,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GatewayMessage) Reset() {
	*x = GatewayMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMessage) ProtoMessage() {}

func (x *GatewayMessage) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMessage.ProtoReflect.Descriptor instead.
func (*GatewayMessage) Descriptor() ([]byte, []int) {
	return file_libra_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayMessage) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *GatewayMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GatewayMessage) GetProducerName() string {
	if x != nil {
		return x.ProducerName
	}
	return ""
}

func (x *GatewayMessage) GetEventTime() int64 {
	if x != nil {
		return x.EventTime
	}
	return 0
}

func (x *GatewayMessage) GetPublishTime() int64 {
	if x != nil {
		return x.PublishTime
	}
	return 0
}

func (x *GatewayMessage) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *GatewayMessage) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type GatewayWatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic  string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *GatewayWatchRequest) Reset() {
	*x = GatewayWatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayWatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayWatchRequest) ProtoMessage() {}

func (x *GatewayWatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayWatchRequest.ProtoReflect.Descriptor instead.
func (*GatewayWatchRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayWatchRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *GatewayWatchRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type GatewaySendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxLen     int64           `protobuf:"varint,1,opt,name=max_len,json=maxLen,proto3" json:"max_len,omitempty"`
	AutoCreate bool            `protobuf:"varint,2,opt,name=auto_create,json=autoCreate,proto3" json:"auto_create,omitempty"`
	Msg        *GatewayMessage `protobuf:"bytes,10,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GatewaySendRequest) Reset() {
	*x = GatewaySendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewaySendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewaySendRequest) ProtoMessage() {}

func (x *GatewaySendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewaySendRequest.ProtoReflect.Descriptor instead.
func (*GatewaySendRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *GatewaySendRequest) GetMaxLen() int64 {
	if x != nil {
		return x.MaxLen
	}
	return 0
}

func (x *GatewaySendRequest) GetAutoCreate() bool {
	if x != nil {
		return x.AutoCreate
	}
	return false
}

func (x *GatewaySendRequest) GetMsg() *GatewayMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

type GatewaySendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GatewaySendResponse) Reset() {
	*x = GatewaySendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewaySendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewaySendResponse) ProtoMessage() {}

func (x *GatewaySendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewaySendResponse.ProtoReflect.Descriptor instead.
func (*GatewaySendResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_gateway_proto_rawDescGZIP(), []int{3}
}

var File_libra_v1_gateway_proto protoreflect.FileDescriptor

var file_libra_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02,
	0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x7a, 0x0a, 0x12, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74,
	0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x94, 0x01, 0x0a, 0x07, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x44, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x04,
	0x53, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x74, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_libra_v1_gateway_proto_rawDescOnce sync.Once
	file_libra_v1_gateway_proto_rawDescData = file_libra_v1_gateway_proto_rawDesc
)

func file_libra_v1_gateway_proto_rawDescGZIP() []byte {
	file_libra_v1_gateway_proto_rawDescOnce.Do(func() {
		file_libra_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_libra_v1_gateway_proto_rawDescData)
	})
	return file_libra_v1_gateway_proto_rawDescData
}

var file_libra_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_libra_v1_gateway_proto_goTypes = []interface{}{
	(*GatewayMessage)(nil),      // 0: libra.v1.GatewayMessage
	(*GatewayWatchRequest)(nil), // 1: libra.v1.GatewayWatchRequest
	(*GatewaySendRequest)(nil),  // 2: libra.v1.GatewaySendRequest
	(*GatewaySendResponse)(nil), // 3: libra.v1.GatewaySendResponse
	nil,                         // 4: libra.v1.GatewayMessage.PropertiesEntry
	(*anypb.Any)(nil),           // 5: google.protobuf.Any
}
var file_libra_v1_gateway_proto_depIdxs = []int32{
	4, // 0: libra.v1.GatewayMessage.properties:type_name -> libra.v1.GatewayMessage.PropertiesEntry
	5, // 1: libra.v1.GatewayMessage.value:type_name -> google.protobuf.Any
	0, // 2: libra.v1.GatewaySendRequest.msg:type_name -> libra.v1.GatewayMessage
	1, // 3: libra.v1.Gateway.Watch:input_type -> libra.v1.GatewayWatchRequest
	2, // 4: libra.v1.Gateway.Send:input_type -> libra.v1.GatewaySendRequest
	0, // 5: libra.v1.Gateway.Watch:output_type -> libra.v1.GatewayMessage
	3, // 6: libra.v1.Gateway.Send:output_type -> libra.v1.GatewaySendResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_libra_v1_gateway_proto_init() }
func file_libra_v1_gateway_proto_init() {
	if File_libra_v1_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_libra_v1_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMessage); i {
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
		file_libra_v1_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayWatchRequest); i {
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
		file_libra_v1_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewaySendRequest); i {
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
		file_libra_v1_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewaySendResponse); i {
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
			RawDescriptor: file_libra_v1_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_libra_v1_gateway_proto_goTypes,
		DependencyIndexes: file_libra_v1_gateway_proto_depIdxs,
		MessageInfos:      file_libra_v1_gateway_proto_msgTypes,
	}.Build()
	File_libra_v1_gateway_proto = out.File
	file_libra_v1_gateway_proto_rawDesc = nil
	file_libra_v1_gateway_proto_goTypes = nil
	file_libra_v1_gateway_proto_depIdxs = nil
}
