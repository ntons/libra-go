// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: libra/v1/chart.proto

package v1

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

// chart key {app_id:name}[:suffix]
// 一般情况下，suffix不需要提供，redis集群会以{app_id:name}
// 的一致性哈希结果对请求进行路由。
// 但有一种情况需要考虑使用suffix，需要construct_from但时候，
// 需要保证被复制对象和新生成对象在同一台节点上，这时候就
// 需要suffix，实现内部会生成{app_id:name}:suffix这样的key，
// 括号内为redis的hash tag，只用hash tag来进行一致性哈希，
// 从而保证所有相关对象都在同一节点上。
type ChartKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//string app_id   = 1;
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3" json:"suffix,omitempty"`
}

func (x *ChartKey) Reset() {
	*x = ChartKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_chart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartKey) ProtoMessage() {}

func (x *ChartKey) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_chart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartKey.ProtoReflect.Descriptor instead.
func (*ChartKey) Descriptor() ([]byte, []int) {
	return file_libra_v1_chart_proto_rawDescGZIP(), []int{0}
}

func (x *ChartKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChartKey) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

// chart entry
type ChartEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank  int32  `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Info  string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	Score int64  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ChartEntry) Reset() {
	*x = ChartEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_chart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartEntry) ProtoMessage() {}

func (x *ChartEntry) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_chart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartEntry.ProtoReflect.Descriptor instead.
func (*ChartEntry) Descriptor() ([]byte, []int) {
	return file_libra_v1_chart_proto_rawDescGZIP(), []int{1}
}

func (x *ChartEntry) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *ChartEntry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChartEntry) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *ChartEntry) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

// chart options
// NOTE:
//  if options not set, the existed chart keep it's current options
//  if field of options not set, the existed chart SET THIS OPTION TO DEFAULT
type ChartOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// maximal entries in the chart
	Capacity int32 `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	// duplicate from this name IF CHART NOT EXISTS
	// can only duplicate from chart which has the same name
	ConstructFrom *ChartKey `protobuf:"bytes,2,opt,name=construct_from,json=constructFrom,proto3" json:"construct_from,omitempty"`
	// chart expire at timestamp
	ExpireAt int64 `protobuf:"varint,3,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	// chart expire after duration, if expire_at set, this option was ignored
	IdleExpire int64 `protobuf:"varint,4,opt,name=idle_expire,json=idleExpire,proto3" json:"idle_expire,omitempty"`
}

func (x *ChartOptions) Reset() {
	*x = ChartOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_chart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartOptions) ProtoMessage() {}

func (x *ChartOptions) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_chart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartOptions.ProtoReflect.Descriptor instead.
func (*ChartOptions) Descriptor() ([]byte, []int) {
	return file_libra_v1_chart_proto_rawDescGZIP(), []int{2}
}

func (x *ChartOptions) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *ChartOptions) GetConstructFrom() *ChartKey {
	if x != nil {
		return x.ConstructFrom
	}
	return nil
}

func (x *ChartOptions) GetExpireAt() int64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *ChartOptions) GetIdleExpire() int64 {
	if x != nil {
		return x.IdleExpire
	}
	return 0
}

var File_libra_v1_chart_proto protoreflect.FileDescriptor

var file_libra_v1_chart_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x22, 0x36, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0x5a, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x39, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x64, 0x6c,
	0x65, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x69, 0x64, 0x6c, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x6f, 0x6e, 0x73, 0x2f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_libra_v1_chart_proto_rawDescOnce sync.Once
	file_libra_v1_chart_proto_rawDescData = file_libra_v1_chart_proto_rawDesc
)

func file_libra_v1_chart_proto_rawDescGZIP() []byte {
	file_libra_v1_chart_proto_rawDescOnce.Do(func() {
		file_libra_v1_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_libra_v1_chart_proto_rawDescData)
	})
	return file_libra_v1_chart_proto_rawDescData
}

var file_libra_v1_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_libra_v1_chart_proto_goTypes = []interface{}{
	(*ChartKey)(nil),     // 0: libra.v1.ChartKey
	(*ChartEntry)(nil),   // 1: libra.v1.ChartEntry
	(*ChartOptions)(nil), // 2: libra.v1.ChartOptions
}
var file_libra_v1_chart_proto_depIdxs = []int32{
	0, // 0: libra.v1.ChartOptions.construct_from:type_name -> libra.v1.ChartKey
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_libra_v1_chart_proto_init() }
func file_libra_v1_chart_proto_init() {
	if File_libra_v1_chart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_libra_v1_chart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartKey); i {
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
		file_libra_v1_chart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartEntry); i {
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
		file_libra_v1_chart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartOptions); i {
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
			RawDescriptor: file_libra_v1_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_libra_v1_chart_proto_goTypes,
		DependencyIndexes: file_libra_v1_chart_proto_depIdxs,
		MessageInfos:      file_libra_v1_chart_proto_msgTypes,
	}.Build()
	File_libra_v1_chart_proto = out.File
	file_libra_v1_chart_proto_rawDesc = nil
	file_libra_v1_chart_proto_goTypes = nil
	file_libra_v1_chart_proto_depIdxs = nil
}