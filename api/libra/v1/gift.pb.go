// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: libra/v1/gift.proto

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

type GiftData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 过期时间
	ExpireAt int64 `protobuf:"varint,2,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	// 礼包内容
	Payload *anypb.Any `protobuf:"bytes,10,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *GiftData) Reset() {
	*x = GiftData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftData) ProtoMessage() {}

func (x *GiftData) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftData.ProtoReflect.Descriptor instead.
func (*GiftData) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{0}
}

func (x *GiftData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GiftData) GetExpireAt() int64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *GiftData) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

type GiftCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *GiftData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Codes []string  `protobuf:"bytes,10,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *GiftCreateRequest) Reset() {
	*x = GiftCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftCreateRequest) ProtoMessage() {}

func (x *GiftCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftCreateRequest.ProtoReflect.Descriptor instead.
func (*GiftCreateRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{1}
}

func (x *GiftCreateRequest) GetData() *GiftData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GiftCreateRequest) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

type GiftCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GiftCreateResponse) Reset() {
	*x = GiftCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftCreateResponse) ProtoMessage() {}

func (x *GiftCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftCreateResponse.ProtoReflect.Descriptor instead.
func (*GiftCreateResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{2}
}

type GiftRevokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GiftRevokeRequest) Reset() {
	*x = GiftRevokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftRevokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftRevokeRequest) ProtoMessage() {}

func (x *GiftRevokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftRevokeRequest.ProtoReflect.Descriptor instead.
func (*GiftRevokeRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{3}
}

func (x *GiftRevokeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GiftRevokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GiftRevokeResponse) Reset() {
	*x = GiftRevokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftRevokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftRevokeResponse) ProtoMessage() {}

func (x *GiftRevokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftRevokeResponse.ProtoReflect.Descriptor instead.
func (*GiftRevokeResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{4}
}

type GiftUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// data中的id被忽略，以上面的id为准
	Data       *GiftData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	CodesToAdd []string  `protobuf:"bytes,13,rep,name=codes_to_add,json=codesToAdd,proto3" json:"codes_to_add,omitempty"`
	CodesToDel []string  `protobuf:"bytes,14,rep,name=codes_to_del,json=codesToDel,proto3" json:"codes_to_del,omitempty"`
}

func (x *GiftUpdateRequest) Reset() {
	*x = GiftUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftUpdateRequest) ProtoMessage() {}

func (x *GiftUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftUpdateRequest.ProtoReflect.Descriptor instead.
func (*GiftUpdateRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{5}
}

func (x *GiftUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GiftUpdateRequest) GetData() *GiftData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GiftUpdateRequest) GetCodesToAdd() []string {
	if x != nil {
		return x.CodesToAdd
	}
	return nil
}

func (x *GiftUpdateRequest) GetCodesToDel() []string {
	if x != nil {
		return x.CodesToDel
	}
	return nil
}

type GiftUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GiftUpdateResponse) Reset() {
	*x = GiftUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftUpdateResponse) ProtoMessage() {}

func (x *GiftUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftUpdateResponse.ProtoReflect.Descriptor instead.
func (*GiftUpdateResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{6}
}

type GiftVerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GiftVerifyRequest) Reset() {
	*x = GiftVerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftVerifyRequest) ProtoMessage() {}

func (x *GiftVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftVerifyRequest.ProtoReflect.Descriptor instead.
func (*GiftVerifyRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{7}
}

func (x *GiftVerifyRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GiftVerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GiftData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GiftVerifyResponse) Reset() {
	*x = GiftVerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftVerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftVerifyResponse) ProtoMessage() {}

func (x *GiftVerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftVerifyResponse.ProtoReflect.Descriptor instead.
func (*GiftVerifyResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{8}
}

func (x *GiftVerifyResponse) GetData() *GiftData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GiftRedeemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GiftRedeemRequest) Reset() {
	*x = GiftRedeemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftRedeemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftRedeemRequest) ProtoMessage() {}

func (x *GiftRedeemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftRedeemRequest.ProtoReflect.Descriptor instead.
func (*GiftRedeemRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{9}
}

func (x *GiftRedeemRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GiftRedeemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GiftData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GiftRedeemResponse) Reset() {
	*x = GiftRedeemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_gift_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiftRedeemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiftRedeemResponse) ProtoMessage() {}

func (x *GiftRedeemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_gift_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiftRedeemResponse.ProtoReflect.Descriptor instead.
func (*GiftRedeemResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_gift_proto_rawDescGZIP(), []int{10}
}

func (x *GiftRedeemResponse) GetData() *GiftData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_libra_v1_gift_proto protoreflect.FileDescriptor

var file_libra_v1_gift_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x66, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x08, 0x47, 0x69,
	0x66, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x51, 0x0a, 0x11, 0x47, 0x69, 0x66, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x69, 0x66, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11,
	0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x47, 0x69, 0x66, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x5f, 0x74,
	0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x6c, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x54, 0x6f, 0x44, 0x65, 0x6c, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x69, 0x66,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x27, 0x0a, 0x11, 0x47, 0x69, 0x66, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3c, 0x0a, 0x12, 0x47, 0x69, 0x66, 0x74,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65,
	0x64, 0x65, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x3c, 0x0a, 0x12, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x69, 0x66, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xdf, 0x02,
	0x0a, 0x04, 0x47, 0x69, 0x66, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69,
	0x66, 0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12,
	0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66, 0x74, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x52, 0x65,
	0x64, 0x65, 0x65, 0x6d, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x66,
	0x74, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74,
	0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_libra_v1_gift_proto_rawDescOnce sync.Once
	file_libra_v1_gift_proto_rawDescData = file_libra_v1_gift_proto_rawDesc
)

func file_libra_v1_gift_proto_rawDescGZIP() []byte {
	file_libra_v1_gift_proto_rawDescOnce.Do(func() {
		file_libra_v1_gift_proto_rawDescData = protoimpl.X.CompressGZIP(file_libra_v1_gift_proto_rawDescData)
	})
	return file_libra_v1_gift_proto_rawDescData
}

var file_libra_v1_gift_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_libra_v1_gift_proto_goTypes = []interface{}{
	(*GiftData)(nil),           // 0: libra.v1.GiftData
	(*GiftCreateRequest)(nil),  // 1: libra.v1.GiftCreateRequest
	(*GiftCreateResponse)(nil), // 2: libra.v1.GiftCreateResponse
	(*GiftRevokeRequest)(nil),  // 3: libra.v1.GiftRevokeRequest
	(*GiftRevokeResponse)(nil), // 4: libra.v1.GiftRevokeResponse
	(*GiftUpdateRequest)(nil),  // 5: libra.v1.GiftUpdateRequest
	(*GiftUpdateResponse)(nil), // 6: libra.v1.GiftUpdateResponse
	(*GiftVerifyRequest)(nil),  // 7: libra.v1.GiftVerifyRequest
	(*GiftVerifyResponse)(nil), // 8: libra.v1.GiftVerifyResponse
	(*GiftRedeemRequest)(nil),  // 9: libra.v1.GiftRedeemRequest
	(*GiftRedeemResponse)(nil), // 10: libra.v1.GiftRedeemResponse
	(*anypb.Any)(nil),          // 11: google.protobuf.Any
}
var file_libra_v1_gift_proto_depIdxs = []int32{
	11, // 0: libra.v1.GiftData.payload:type_name -> google.protobuf.Any
	0,  // 1: libra.v1.GiftCreateRequest.data:type_name -> libra.v1.GiftData
	0,  // 2: libra.v1.GiftUpdateRequest.data:type_name -> libra.v1.GiftData
	0,  // 3: libra.v1.GiftVerifyResponse.data:type_name -> libra.v1.GiftData
	0,  // 4: libra.v1.GiftRedeemResponse.data:type_name -> libra.v1.GiftData
	1,  // 5: libra.v1.Gift.Create:input_type -> libra.v1.GiftCreateRequest
	3,  // 6: libra.v1.Gift.Revoke:input_type -> libra.v1.GiftRevokeRequest
	5,  // 7: libra.v1.Gift.Update:input_type -> libra.v1.GiftUpdateRequest
	7,  // 8: libra.v1.Gift.Verify:input_type -> libra.v1.GiftVerifyRequest
	9,  // 9: libra.v1.Gift.Redeem:input_type -> libra.v1.GiftRedeemRequest
	2,  // 10: libra.v1.Gift.Create:output_type -> libra.v1.GiftCreateResponse
	4,  // 11: libra.v1.Gift.Revoke:output_type -> libra.v1.GiftRevokeResponse
	6,  // 12: libra.v1.Gift.Update:output_type -> libra.v1.GiftUpdateResponse
	8,  // 13: libra.v1.Gift.Verify:output_type -> libra.v1.GiftVerifyResponse
	10, // 14: libra.v1.Gift.Redeem:output_type -> libra.v1.GiftRedeemResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_libra_v1_gift_proto_init() }
func file_libra_v1_gift_proto_init() {
	if File_libra_v1_gift_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_libra_v1_gift_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftData); i {
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
		file_libra_v1_gift_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftCreateRequest); i {
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
		file_libra_v1_gift_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftCreateResponse); i {
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
		file_libra_v1_gift_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftRevokeRequest); i {
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
		file_libra_v1_gift_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftRevokeResponse); i {
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
		file_libra_v1_gift_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftUpdateRequest); i {
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
		file_libra_v1_gift_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftUpdateResponse); i {
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
		file_libra_v1_gift_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftVerifyRequest); i {
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
		file_libra_v1_gift_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftVerifyResponse); i {
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
		file_libra_v1_gift_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftRedeemRequest); i {
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
		file_libra_v1_gift_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiftRedeemResponse); i {
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
			RawDescriptor: file_libra_v1_gift_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_libra_v1_gift_proto_goTypes,
		DependencyIndexes: file_libra_v1_gift_proto_depIdxs,
		MessageInfos:      file_libra_v1_gift_proto_msgTypes,
	}.Build()
	File_libra_v1_gift_proto = out.File
	file_libra_v1_gift_proto_rawDesc = nil
	file_libra_v1_gift_proto_goTypes = nil
	file_libra_v1_gift_proto_depIdxs = nil
}
