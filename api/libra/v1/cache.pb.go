// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: libra/v1/cache.proto

package v1

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

type CacheOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeoutMilliseconds int32 `protobuf:"varint,1,opt,name=timeout_milliseconds,json=timeoutMilliseconds,proto3" json:"timeout_milliseconds,omitempty"`
}

func (x *CacheOptions) Reset() {
	*x = CacheOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheOptions) ProtoMessage() {}

func (x *CacheOptions) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheOptions.ProtoReflect.Descriptor instead.
func (*CacheOptions) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{0}
}

func (x *CacheOptions) GetTimeoutMilliseconds() int32 {
	if x != nil {
		return x.TimeoutMilliseconds
	}
	return 0
}

type CacheGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string                   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Options *CacheGetRequest_Options `protobuf:"bytes,10,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *CacheGetRequest) Reset() {
	*x = CacheGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheGetRequest) ProtoMessage() {}

func (x *CacheGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheGetRequest.ProtoReflect.Descriptor instead.
func (*CacheGetRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{1}
}

func (x *CacheGetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CacheGetRequest) GetOptions() *CacheGetRequest_Options {
	if x != nil {
		return x.Options
	}
	return nil
}

type CacheGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CacheGetResponse) Reset() {
	*x = CacheGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheGetResponse) ProtoMessage() {}

func (x *CacheGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheGetResponse.ProtoReflect.Descriptor instead.
func (*CacheGetResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{2}
}

func (x *CacheGetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type CacheSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string                   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   []byte                   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Options *CacheSetRequest_Options `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *CacheSetRequest) Reset() {
	*x = CacheSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheSetRequest) ProtoMessage() {}

func (x *CacheSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheSetRequest.ProtoReflect.Descriptor instead.
func (*CacheSetRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{3}
}

func (x *CacheSetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CacheSetRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *CacheSetRequest) GetOptions() *CacheSetRequest_Options {
	if x != nil {
		return x.Options
	}
	return nil
}

type CacheSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CacheSetResponse) Reset() {
	*x = CacheSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheSetResponse) ProtoMessage() {}

func (x *CacheSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheSetResponse.ProtoReflect.Descriptor instead.
func (*CacheSetResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{4}
}

type CacheAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string                   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   []byte                   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Options *CacheAddRequest_Options `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *CacheAddRequest) Reset() {
	*x = CacheAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheAddRequest) ProtoMessage() {}

func (x *CacheAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheAddRequest.ProtoReflect.Descriptor instead.
func (*CacheAddRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{5}
}

func (x *CacheAddRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CacheAddRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *CacheAddRequest) GetOptions() *CacheAddRequest_Options {
	if x != nil {
		return x.Options
	}
	return nil
}

type CacheAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CacheAddResponse) Reset() {
	*x = CacheAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheAddResponse) ProtoMessage() {}

func (x *CacheAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheAddResponse.ProtoReflect.Descriptor instead.
func (*CacheAddResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{6}
}

type CacheGetRequest_Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Polling in C# cause NotFound exception which lead to high latency
	RegardNotFoundAsEmpty bool `protobuf:"varint,1,opt,name=regard_not_found_as_empty,json=regardNotFoundAsEmpty,proto3" json:"regard_not_found_as_empty,omitempty"`
}

func (x *CacheGetRequest_Options) Reset() {
	*x = CacheGetRequest_Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheGetRequest_Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheGetRequest_Options) ProtoMessage() {}

func (x *CacheGetRequest_Options) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheGetRequest_Options.ProtoReflect.Descriptor instead.
func (*CacheGetRequest_Options) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CacheGetRequest_Options) GetRegardNotFoundAsEmpty() bool {
	if x != nil {
		return x.RegardNotFoundAsEmpty
	}
	return false
}

type CacheSetRequest_Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeoutMilliseconds int32 `protobuf:"varint,1,opt,name=timeout_milliseconds,json=timeoutMilliseconds,proto3" json:"timeout_milliseconds,omitempty"`
}

func (x *CacheSetRequest_Options) Reset() {
	*x = CacheSetRequest_Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheSetRequest_Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheSetRequest_Options) ProtoMessage() {}

func (x *CacheSetRequest_Options) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheSetRequest_Options.ProtoReflect.Descriptor instead.
func (*CacheSetRequest_Options) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CacheSetRequest_Options) GetTimeoutMilliseconds() int32 {
	if x != nil {
		return x.TimeoutMilliseconds
	}
	return 0
}

type CacheAddRequest_Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeoutMilliseconds int32 `protobuf:"varint,1,opt,name=timeout_milliseconds,json=timeoutMilliseconds,proto3" json:"timeout_milliseconds,omitempty"`
}

func (x *CacheAddRequest_Options) Reset() {
	*x = CacheAddRequest_Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_cache_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheAddRequest_Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheAddRequest_Options) ProtoMessage() {}

func (x *CacheAddRequest_Options) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_cache_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheAddRequest_Options.ProtoReflect.Descriptor instead.
func (*CacheAddRequest_Options) Descriptor() ([]byte, []int) {
	return file_libra_v1_cache_proto_rawDescGZIP(), []int{5, 0}
}

func (x *CacheAddRequest_Options) GetTimeoutMilliseconds() int32 {
	if x != nil {
		return x.TimeoutMilliseconds
	}
	return 0
}

var File_libra_v1_cache_proto protoreflect.FileDescriptor

var file_libra_v1_cache_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x22, 0x41, 0x0a, 0x0c, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x31, 0x0a, 0x14, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x43, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x38, 0x0a, 0x19, 0x72, 0x65, 0x67, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x73, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x72, 0x65, 0x67, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x74, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x41, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x0a, 0x10, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3c,
	0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x12, 0x0a, 0x10,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xb4, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3c, 0x0a, 0x07, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f,
	0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x13, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc1, 0x01, 0x0a, 0x05,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74,
	0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_libra_v1_cache_proto_rawDescOnce sync.Once
	file_libra_v1_cache_proto_rawDescData = file_libra_v1_cache_proto_rawDesc
)

func file_libra_v1_cache_proto_rawDescGZIP() []byte {
	file_libra_v1_cache_proto_rawDescOnce.Do(func() {
		file_libra_v1_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_libra_v1_cache_proto_rawDescData)
	})
	return file_libra_v1_cache_proto_rawDescData
}

var file_libra_v1_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_libra_v1_cache_proto_goTypes = []interface{}{
	(*CacheOptions)(nil),            // 0: libra.v1.CacheOptions
	(*CacheGetRequest)(nil),         // 1: libra.v1.CacheGetRequest
	(*CacheGetResponse)(nil),        // 2: libra.v1.CacheGetResponse
	(*CacheSetRequest)(nil),         // 3: libra.v1.CacheSetRequest
	(*CacheSetResponse)(nil),        // 4: libra.v1.CacheSetResponse
	(*CacheAddRequest)(nil),         // 5: libra.v1.CacheAddRequest
	(*CacheAddResponse)(nil),        // 6: libra.v1.CacheAddResponse
	(*CacheGetRequest_Options)(nil), // 7: libra.v1.CacheGetRequest.Options
	(*CacheSetRequest_Options)(nil), // 8: libra.v1.CacheSetRequest.Options
	(*CacheAddRequest_Options)(nil), // 9: libra.v1.CacheAddRequest.Options
}
var file_libra_v1_cache_proto_depIdxs = []int32{
	7, // 0: libra.v1.CacheGetRequest.options:type_name -> libra.v1.CacheGetRequest.Options
	8, // 1: libra.v1.CacheSetRequest.options:type_name -> libra.v1.CacheSetRequest.Options
	9, // 2: libra.v1.CacheAddRequest.options:type_name -> libra.v1.CacheAddRequest.Options
	1, // 3: libra.v1.Cache.Get:input_type -> libra.v1.CacheGetRequest
	3, // 4: libra.v1.Cache.Set:input_type -> libra.v1.CacheSetRequest
	5, // 5: libra.v1.Cache.Add:input_type -> libra.v1.CacheAddRequest
	2, // 6: libra.v1.Cache.Get:output_type -> libra.v1.CacheGetResponse
	4, // 7: libra.v1.Cache.Set:output_type -> libra.v1.CacheSetResponse
	6, // 8: libra.v1.Cache.Add:output_type -> libra.v1.CacheAddResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_libra_v1_cache_proto_init() }
func file_libra_v1_cache_proto_init() {
	if File_libra_v1_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_libra_v1_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheOptions); i {
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
		file_libra_v1_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheGetRequest); i {
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
		file_libra_v1_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheGetResponse); i {
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
		file_libra_v1_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheSetRequest); i {
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
		file_libra_v1_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheSetResponse); i {
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
		file_libra_v1_cache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheAddRequest); i {
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
		file_libra_v1_cache_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheAddResponse); i {
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
		file_libra_v1_cache_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheGetRequest_Options); i {
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
		file_libra_v1_cache_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheSetRequest_Options); i {
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
		file_libra_v1_cache_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheAddRequest_Options); i {
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
			RawDescriptor: file_libra_v1_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_libra_v1_cache_proto_goTypes,
		DependencyIndexes: file_libra_v1_cache_proto_depIdxs,
		MessageInfos:      file_libra_v1_cache_proto_msgTypes,
	}.Build()
	File_libra_v1_cache_proto = out.File
	file_libra_v1_cache_proto_rawDesc = nil
	file_libra_v1_cache_proto_goTypes = nil
	file_libra_v1_cache_proto_depIdxs = nil
}
