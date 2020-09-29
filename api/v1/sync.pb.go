// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: v1/sync.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	any "github.com/golang/protobuf/ptypes/any"
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

type SyncLockOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncLockOptions) Reset() {
	*x = SyncLockOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncLockOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncLockOptions) ProtoMessage() {}

func (x *SyncLockOptions) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncLockOptions.ProtoReflect.Descriptor instead.
func (*SyncLockOptions) Descriptor() ([]byte, []int) {
	return file_v1_sync_proto_rawDescGZIP(), []int{0}
}

type SyncUnlockOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 即使操作失败，依然释放锁
	EvenOnFailure bool `protobuf:"varint,1,opt,name=even_on_failure,json=evenOnFailure,proto3" json:"even_on_failure,omitempty"`
}

func (x *SyncUnlockOptions) Reset() {
	*x = SyncUnlockOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncUnlockOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncUnlockOptions) ProtoMessage() {}

func (x *SyncUnlockOptions) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncUnlockOptions.ProtoReflect.Descriptor instead.
func (*SyncUnlockOptions) Descriptor() ([]byte, []int) {
	return file_v1_sync_proto_rawDescGZIP(), []int{1}
}

func (x *SyncUnlockOptions) GetEvenOnFailure() bool {
	if x != nil {
		return x.EvenOnFailure
	}
	return false
}

type SyncLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         *EntityKey       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	LockOptions *SyncLockOptions `protobuf:"bytes,11,opt,name=lock_options,json=lockOptions,proto3" json:"lock_options,omitempty"`
}

func (x *SyncLockRequest) Reset() {
	*x = SyncLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncLockRequest) ProtoMessage() {}

func (x *SyncLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncLockRequest.ProtoReflect.Descriptor instead.
func (*SyncLockRequest) Descriptor() ([]byte, []int) {
	return file_v1_sync_proto_rawDescGZIP(), []int{2}
}

func (x *SyncLockRequest) GetKey() *EntityKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SyncLockRequest) GetLockOptions() *SyncLockOptions {
	if x != nil {
		return x.LockOptions
	}
	return nil
}

type SyncLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 获取到的互斥锁，你不需要知道这是个什么，在解锁时候传回来就行了
	Lock *any.Any `protobuf:"bytes,10,opt,name=lock,proto3" json:"lock,omitempty"`
}

func (x *SyncLockResponse) Reset() {
	*x = SyncLockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncLockResponse) ProtoMessage() {}

func (x *SyncLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncLockResponse.ProtoReflect.Descriptor instead.
func (*SyncLockResponse) Descriptor() ([]byte, []int) {
	return file_v1_sync_proto_rawDescGZIP(), []int{3}
}

func (x *SyncLockResponse) GetLock() *any.Any {
	if x != nil {
		return x.Lock
	}
	return nil
}

type SyncUnlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lock          *any.Any           `protobuf:"bytes,10,opt,name=lock,proto3" json:"lock,omitempty"`
	UnlockOptions *SyncUnlockOptions `protobuf:"bytes,11,opt,name=unlock_options,json=unlockOptions,proto3" json:"unlock_options,omitempty"`
}

func (x *SyncUnlockRequest) Reset() {
	*x = SyncUnlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sync_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncUnlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncUnlockRequest) ProtoMessage() {}

func (x *SyncUnlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sync_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncUnlockRequest.ProtoReflect.Descriptor instead.
func (*SyncUnlockRequest) Descriptor() ([]byte, []int) {
	return file_v1_sync_proto_rawDescGZIP(), []int{4}
}

func (x *SyncUnlockRequest) GetLock() *any.Any {
	if x != nil {
		return x.Lock
	}
	return nil
}

func (x *SyncUnlockRequest) GetUnlockOptions() *SyncUnlockOptions {
	if x != nil {
		return x.UnlockOptions
	}
	return nil
}

type SyncUnlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncUnlockResponse) Reset() {
	*x = SyncUnlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_sync_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncUnlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncUnlockResponse) ProtoMessage() {}

func (x *SyncUnlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sync_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncUnlockResponse.ProtoReflect.Descriptor instead.
func (*SyncUnlockResponse) Descriptor() ([]byte, []int) {
	return file_v1_sync_proto_rawDescGZIP(), []int{5}
}

var File_v1_sync_proto protoreflect.FileDescriptor

var file_v1_sync_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x4c,
	0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3b, 0x0a, 0x11, 0x53, 0x79,
	0x6e, 0x63, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x4f, 0x6e,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x22, 0x76, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x4c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x3c, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x81, 0x01,
	0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x42, 0x0a,
	0x0e, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0d, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8a, 0x01, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63,
	0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x43, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2d, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_sync_proto_rawDescOnce sync.Once
	file_v1_sync_proto_rawDescData = file_v1_sync_proto_rawDesc
)

func file_v1_sync_proto_rawDescGZIP() []byte {
	file_v1_sync_proto_rawDescOnce.Do(func() {
		file_v1_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_sync_proto_rawDescData)
	})
	return file_v1_sync_proto_rawDescData
}

var file_v1_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_sync_proto_goTypes = []interface{}{
	(*SyncLockOptions)(nil),    // 0: libra.v1.SyncLockOptions
	(*SyncUnlockOptions)(nil),  // 1: libra.v1.SyncUnlockOptions
	(*SyncLockRequest)(nil),    // 2: libra.v1.SyncLockRequest
	(*SyncLockResponse)(nil),   // 3: libra.v1.SyncLockResponse
	(*SyncUnlockRequest)(nil),  // 4: libra.v1.SyncUnlockRequest
	(*SyncUnlockResponse)(nil), // 5: libra.v1.SyncUnlockResponse
	(*EntityKey)(nil),          // 6: libra.v1.EntityKey
	(*any.Any)(nil),            // 7: google.protobuf.Any
}
var file_v1_sync_proto_depIdxs = []int32{
	6, // 0: libra.v1.SyncLockRequest.key:type_name -> libra.v1.EntityKey
	0, // 1: libra.v1.SyncLockRequest.lock_options:type_name -> libra.v1.SyncLockOptions
	7, // 2: libra.v1.SyncLockResponse.lock:type_name -> google.protobuf.Any
	7, // 3: libra.v1.SyncUnlockRequest.lock:type_name -> google.protobuf.Any
	1, // 4: libra.v1.SyncUnlockRequest.unlock_options:type_name -> libra.v1.SyncUnlockOptions
	2, // 5: libra.v1.Sync.Lock:input_type -> libra.v1.SyncLockRequest
	4, // 6: libra.v1.Sync.Unlock:input_type -> libra.v1.SyncUnlockRequest
	3, // 7: libra.v1.Sync.Lock:output_type -> libra.v1.SyncLockResponse
	5, // 8: libra.v1.Sync.Unlock:output_type -> libra.v1.SyncUnlockResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_sync_proto_init() }
func file_v1_sync_proto_init() {
	if File_v1_sync_proto != nil {
		return
	}
	file_v1_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncLockOptions); i {
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
		file_v1_sync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncUnlockOptions); i {
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
		file_v1_sync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncLockRequest); i {
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
		file_v1_sync_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncLockResponse); i {
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
		file_v1_sync_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncUnlockRequest); i {
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
		file_v1_sync_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncUnlockResponse); i {
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
			RawDescriptor: file_v1_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_sync_proto_goTypes,
		DependencyIndexes: file_v1_sync_proto_depIdxs,
		MessageInfos:      file_v1_sync_proto_msgTypes,
	}.Build()
	File_v1_sync_proto = out.File
	file_v1_sync_proto_rawDesc = nil
	file_v1_sync_proto_goTypes = nil
	file_v1_sync_proto_depIdxs = nil
}