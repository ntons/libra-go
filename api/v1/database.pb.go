// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.1
// source: v1/database.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DatabaseRegisterSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proto file descriptor set
	DescriptorSet *descriptorpb.FileDescriptorSet `protobuf:"bytes,1,opt,name=descriptor_set,json=descriptorSet,proto3" json:"descriptor_set,omitempty"`
	// proto message name
	MessageName string `protobuf:"bytes,2,opt,name=message_name,json=messageName,proto3" json:"message_name,omitempty"`
}

func (x *DatabaseRegisterSchemaRequest) Reset() {
	*x = DatabaseRegisterSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseRegisterSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseRegisterSchemaRequest) ProtoMessage() {}

func (x *DatabaseRegisterSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseRegisterSchemaRequest.ProtoReflect.Descriptor instead.
func (*DatabaseRegisterSchemaRequest) Descriptor() ([]byte, []int) {
	return file_v1_database_proto_rawDescGZIP(), []int{0}
}

func (x *DatabaseRegisterSchemaRequest) GetDescriptorSet() *descriptorpb.FileDescriptorSet {
	if x != nil {
		return x.DescriptorSet
	}
	return nil
}

func (x *DatabaseRegisterSchemaRequest) GetMessageName() string {
	if x != nil {
		return x.MessageName
	}
	return ""
}

type DatabaseRegisterSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique identifier for registered proto
	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *DatabaseRegisterSchemaResponse) Reset() {
	*x = DatabaseRegisterSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseRegisterSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseRegisterSchemaResponse) ProtoMessage() {}

func (x *DatabaseRegisterSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseRegisterSchemaResponse.ProtoReflect.Descriptor instead.
func (*DatabaseRegisterSchemaResponse) Descriptor() ([]byte, []int) {
	return file_v1_database_proto_rawDescGZIP(), []int{1}
}

func (x *DatabaseRegisterSchemaResponse) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type DatabaseGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key           *EntryKey  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	AddIfNotFound *anypb.Any `protobuf:"bytes,9,opt,name=add_if_not_found,json=addIfNotFound,proto3" json:"add_if_not_found,omitempty"`
	// 如果设置lock_options会在获取数据之前尝试获取id对应的互斥锁
	// 如果上锁失败，数据将不会获取
	LockOptions *DistlockLockOptions `protobuf:"bytes,11,opt,name=lock_options,json=lockOptions,proto3" json:"lock_options,omitempty"`
}

func (x *DatabaseGetRequest) Reset() {
	*x = DatabaseGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseGetRequest) ProtoMessage() {}

func (x *DatabaseGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseGetRequest.ProtoReflect.Descriptor instead.
func (*DatabaseGetRequest) Descriptor() ([]byte, []int) {
	return file_v1_database_proto_rawDescGZIP(), []int{2}
}

func (x *DatabaseGetRequest) GetKey() *EntryKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *DatabaseGetRequest) GetAddIfNotFound() *anypb.Any {
	if x != nil {
		return x.AddIfNotFound
	}
	return nil
}

func (x *DatabaseGetRequest) GetLockOptions() *DistlockLockOptions {
	if x != nil {
		return x.LockOptions
	}
	return nil
}

type DatabaseGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision  int64      `protobuf:"varint,8,opt,name=revision,proto3" json:"revision,omitempty"` // data revision
	Data      *anypb.Any `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
	LockToken *anypb.Any `protobuf:"bytes,10,opt,name=lock_token,json=lockToken,proto3" json:"lock_token,omitempty"`
}

func (x *DatabaseGetResponse) Reset() {
	*x = DatabaseGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseGetResponse) ProtoMessage() {}

func (x *DatabaseGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseGetResponse.ProtoReflect.Descriptor instead.
func (*DatabaseGetResponse) Descriptor() ([]byte, []int) {
	return file_v1_database_proto_rawDescGZIP(), []int{3}
}

func (x *DatabaseGetResponse) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *DatabaseGetResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DatabaseGetResponse) GetLockToken() *anypb.Any {
	if x != nil {
		return x.LockToken
	}
	return nil
}

type DatabaseSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  *EntryKey  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data *anypb.Any `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
	// 获取到的互斥锁，更新数据前会校验锁是否有效，如果锁无效则不会更新数据。
	// 如果不设置，数据则会被强制更新。
	// 每次更新都会重置锁的TTL，即使更新失败。
	LockToken *anypb.Any `protobuf:"bytes,10,opt,name=lock_token,json=lockToken,proto3" json:"lock_token,omitempty"`
	// 设置unlock_options会在操作完成时解锁，要求lock_token有效
	UnlockOptions *DistlockUnlockOptions `protobuf:"bytes,11,opt,name=unlock_options,json=unlockOptions,proto3" json:"unlock_options,omitempty"`
}

func (x *DatabaseSetRequest) Reset() {
	*x = DatabaseSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseSetRequest) ProtoMessage() {}

func (x *DatabaseSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseSetRequest.ProtoReflect.Descriptor instead.
func (*DatabaseSetRequest) Descriptor() ([]byte, []int) {
	return file_v1_database_proto_rawDescGZIP(), []int{4}
}

func (x *DatabaseSetRequest) GetKey() *EntryKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *DatabaseSetRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DatabaseSetRequest) GetLockToken() *anypb.Any {
	if x != nil {
		return x.LockToken
	}
	return nil
}

func (x *DatabaseSetRequest) GetUnlockOptions() *DistlockUnlockOptions {
	if x != nil {
		return x.UnlockOptions
	}
	return nil
}

type DatabaseSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision int64 `protobuf:"varint,8,opt,name=revision,proto3" json:"revision,omitempty"` // data revision after setting
}

func (x *DatabaseSetResponse) Reset() {
	*x = DatabaseSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseSetResponse) ProtoMessage() {}

func (x *DatabaseSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseSetResponse.ProtoReflect.Descriptor instead.
func (*DatabaseSetResponse) Descriptor() ([]byte, []int) {
	return file_v1_database_proto_rawDescGZIP(), []int{5}
}

func (x *DatabaseSetResponse) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

var File_v1_database_proto protoreflect.FileDescriptor

var file_v1_database_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x76,
	0x31, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x01, 0x0a, 0x1d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x49, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x0d,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x38, 0x0a, 0x1e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0xbb, 0x01, 0x0a, 0x12, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x5f, 0x69,
	0x66, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x49, 0x66, 0x4e, 0x6f,
	0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x40, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x6c, 0x6f, 0x63, 0x6b,
	0x4c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x6c, 0x6f, 0x63,
	0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xe1, 0x01, 0x0a, 0x12,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x33, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x46, 0x0a, 0x0e, 0x75, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x6c,
	0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x0d, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x31, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x32, 0xf9, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x65, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x27, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x03, 0x53, 0x65,
	0x74, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_database_proto_rawDescOnce sync.Once
	file_v1_database_proto_rawDescData = file_v1_database_proto_rawDesc
)

func file_v1_database_proto_rawDescGZIP() []byte {
	file_v1_database_proto_rawDescOnce.Do(func() {
		file_v1_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_database_proto_rawDescData)
	})
	return file_v1_database_proto_rawDescData
}

var file_v1_database_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_database_proto_goTypes = []interface{}{
	(*DatabaseRegisterSchemaRequest)(nil),  // 0: libra.v1.DatabaseRegisterSchemaRequest
	(*DatabaseRegisterSchemaResponse)(nil), // 1: libra.v1.DatabaseRegisterSchemaResponse
	(*DatabaseGetRequest)(nil),             // 2: libra.v1.DatabaseGetRequest
	(*DatabaseGetResponse)(nil),            // 3: libra.v1.DatabaseGetResponse
	(*DatabaseSetRequest)(nil),             // 4: libra.v1.DatabaseSetRequest
	(*DatabaseSetResponse)(nil),            // 5: libra.v1.DatabaseSetResponse
	(*descriptorpb.FileDescriptorSet)(nil), // 6: google.protobuf.FileDescriptorSet
	(*EntryKey)(nil),                       // 7: libra.v1.EntryKey
	(*anypb.Any)(nil),                      // 8: google.protobuf.Any
	(*DistlockLockOptions)(nil),            // 9: libra.v1.DistlockLockOptions
	(*DistlockUnlockOptions)(nil),          // 10: libra.v1.DistlockUnlockOptions
}
var file_v1_database_proto_depIdxs = []int32{
	6,  // 0: libra.v1.DatabaseRegisterSchemaRequest.descriptor_set:type_name -> google.protobuf.FileDescriptorSet
	7,  // 1: libra.v1.DatabaseGetRequest.key:type_name -> libra.v1.EntryKey
	8,  // 2: libra.v1.DatabaseGetRequest.add_if_not_found:type_name -> google.protobuf.Any
	9,  // 3: libra.v1.DatabaseGetRequest.lock_options:type_name -> libra.v1.DistlockLockOptions
	8,  // 4: libra.v1.DatabaseGetResponse.data:type_name -> google.protobuf.Any
	8,  // 5: libra.v1.DatabaseGetResponse.lock_token:type_name -> google.protobuf.Any
	7,  // 6: libra.v1.DatabaseSetRequest.key:type_name -> libra.v1.EntryKey
	8,  // 7: libra.v1.DatabaseSetRequest.data:type_name -> google.protobuf.Any
	8,  // 8: libra.v1.DatabaseSetRequest.lock_token:type_name -> google.protobuf.Any
	10, // 9: libra.v1.DatabaseSetRequest.unlock_options:type_name -> libra.v1.DistlockUnlockOptions
	0,  // 10: libra.v1.Database.RegisterSchema:input_type -> libra.v1.DatabaseRegisterSchemaRequest
	2,  // 11: libra.v1.Database.Get:input_type -> libra.v1.DatabaseGetRequest
	4,  // 12: libra.v1.Database.Set:input_type -> libra.v1.DatabaseSetRequest
	1,  // 13: libra.v1.Database.RegisterSchema:output_type -> libra.v1.DatabaseRegisterSchemaResponse
	3,  // 14: libra.v1.Database.Get:output_type -> libra.v1.DatabaseGetResponse
	5,  // 15: libra.v1.Database.Set:output_type -> libra.v1.DatabaseSetResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_v1_database_proto_init() }
func file_v1_database_proto_init() {
	if File_v1_database_proto != nil {
		return
	}
	file_v1_entrykey_proto_init()
	file_v1_distlock_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseRegisterSchemaRequest); i {
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
		file_v1_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseRegisterSchemaResponse); i {
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
		file_v1_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseGetRequest); i {
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
		file_v1_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseGetResponse); i {
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
		file_v1_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseSetRequest); i {
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
		file_v1_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseSetResponse); i {
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
			RawDescriptor: file_v1_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_database_proto_goTypes,
		DependencyIndexes: file_v1_database_proto_depIdxs,
		MessageInfos:      file_v1_database_proto_msgTypes,
	}.Build()
	File_v1_database_proto = out.File
	file_v1_database_proto_rawDesc = nil
	file_v1_database_proto_goTypes = nil
	file_v1_database_proto_depIdxs = nil
}
