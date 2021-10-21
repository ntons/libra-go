// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: libra/v1/mailbox.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data *anypb.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Mail) Reset() {
	*x = Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_mailbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail) ProtoMessage() {}

func (x *Mail) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_mailbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mail.ProtoReflect.Descriptor instead.
func (*Mail) Descriptor() ([]byte, []int) {
	return file_libra_v1_mailbox_proto_rawDescGZIP(), []int{0}
}

func (x *Mail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Mail) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type MailboxListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *EntryKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *MailboxListRequest) Reset() {
	*x = MailboxListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_mailbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxListRequest) ProtoMessage() {}

func (x *MailboxListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_mailbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxListRequest.ProtoReflect.Descriptor instead.
func (*MailboxListRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_mailbox_proto_rawDescGZIP(), []int{1}
}

func (x *MailboxListRequest) GetKey() *EntryKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type MailboxListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mails []*Mail `protobuf:"bytes,1,rep,name=mails,proto3" json:"mails,omitempty"`
}

func (x *MailboxListResponse) Reset() {
	*x = MailboxListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_mailbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxListResponse) ProtoMessage() {}

func (x *MailboxListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_mailbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxListResponse.ProtoReflect.Descriptor instead.
func (*MailboxListResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_mailbox_proto_rawDescGZIP(), []int{2}
}

func (x *MailboxListResponse) GetMails() []*Mail {
	if x != nil {
		return x.Mails
	}
	return nil
}

type MailboxPushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      *EntryKey  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data     *anypb.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Capacity int32      `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"` // maximum mails to keep
}

func (x *MailboxPushRequest) Reset() {
	*x = MailboxPushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_mailbox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxPushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxPushRequest) ProtoMessage() {}

func (x *MailboxPushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_mailbox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxPushRequest.ProtoReflect.Descriptor instead.
func (*MailboxPushRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_mailbox_proto_rawDescGZIP(), []int{3}
}

func (x *MailboxPushRequest) GetKey() *EntryKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *MailboxPushRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MailboxPushRequest) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

type MailboxPushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MailboxPushResponse) Reset() {
	*x = MailboxPushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_mailbox_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxPushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxPushResponse) ProtoMessage() {}

func (x *MailboxPushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_mailbox_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxPushResponse.ProtoReflect.Descriptor instead.
func (*MailboxPushResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_mailbox_proto_rawDescGZIP(), []int{4}
}

func (x *MailboxPushResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MailboxPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *EntryKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Ids []string  `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *MailboxPullRequest) Reset() {
	*x = MailboxPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_mailbox_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxPullRequest) ProtoMessage() {}

func (x *MailboxPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_mailbox_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxPullRequest.ProtoReflect.Descriptor instead.
func (*MailboxPullRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_mailbox_proto_rawDescGZIP(), []int{5}
}

func (x *MailboxPullRequest) GetKey() *EntryKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *MailboxPullRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type MailboxPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PulledIds []string `protobuf:"bytes,1,rep,name=pulled_ids,json=pulledIds,proto3" json:"pulled_ids,omitempty"` // pulled id list
}

func (x *MailboxPullResponse) Reset() {
	*x = MailboxPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_mailbox_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxPullResponse) ProtoMessage() {}

func (x *MailboxPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_mailbox_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxPullResponse.ProtoReflect.Descriptor instead.
func (*MailboxPullResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_mailbox_proto_rawDescGZIP(), []int{6}
}

func (x *MailboxPullResponse) GetPulledIds() []string {
	if x != nil {
		return x.PulledIds
	}
	return nil
}

var File_libra_v1_mailbox_proto protoreflect.FileDescriptor

var file_libra_v1_mailbox_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6b, 0x65,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x12, 0x4d, 0x61, 0x69,
	0x6c, 0x62, 0x6f, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3b, 0x0a, 0x13, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x6d, 0x61, 0x69,
	0x6c, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x50, 0x75,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x25, 0x0a, 0x13, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x12,
	0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x34, 0x0a, 0x13, 0x4d, 0x61,
	0x69, 0x6c, 0x62, 0x6f, 0x78, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x6c, 0x6c, 0x65, 0x64, 0x49, 0x64, 0x73,
	0x32, 0xd8, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x12, 0x43, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x69, 0x6c, 0x62, 0x6f, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x43, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x50, 0x75, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x1c,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f,
	0x78, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x50,
	0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x6f, 0x6e, 0x73, 0x2f,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_libra_v1_mailbox_proto_rawDescOnce sync.Once
	file_libra_v1_mailbox_proto_rawDescData = file_libra_v1_mailbox_proto_rawDesc
)

func file_libra_v1_mailbox_proto_rawDescGZIP() []byte {
	file_libra_v1_mailbox_proto_rawDescOnce.Do(func() {
		file_libra_v1_mailbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_libra_v1_mailbox_proto_rawDescData)
	})
	return file_libra_v1_mailbox_proto_rawDescData
}

var file_libra_v1_mailbox_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_libra_v1_mailbox_proto_goTypes = []interface{}{
	(*Mail)(nil),                // 0: libra.v1.Mail
	(*MailboxListRequest)(nil),  // 1: libra.v1.MailboxListRequest
	(*MailboxListResponse)(nil), // 2: libra.v1.MailboxListResponse
	(*MailboxPushRequest)(nil),  // 3: libra.v1.MailboxPushRequest
	(*MailboxPushResponse)(nil), // 4: libra.v1.MailboxPushResponse
	(*MailboxPullRequest)(nil),  // 5: libra.v1.MailboxPullRequest
	(*MailboxPullResponse)(nil), // 6: libra.v1.MailboxPullResponse
	(*anypb.Any)(nil),           // 7: google.protobuf.Any
	(*EntryKey)(nil),            // 8: libra.v1.EntryKey
}
var file_libra_v1_mailbox_proto_depIdxs = []int32{
	7, // 0: libra.v1.Mail.data:type_name -> google.protobuf.Any
	8, // 1: libra.v1.MailboxListRequest.key:type_name -> libra.v1.EntryKey
	0, // 2: libra.v1.MailboxListResponse.mails:type_name -> libra.v1.Mail
	8, // 3: libra.v1.MailboxPushRequest.key:type_name -> libra.v1.EntryKey
	7, // 4: libra.v1.MailboxPushRequest.data:type_name -> google.protobuf.Any
	8, // 5: libra.v1.MailboxPullRequest.key:type_name -> libra.v1.EntryKey
	1, // 6: libra.v1.Mailbox.List:input_type -> libra.v1.MailboxListRequest
	3, // 7: libra.v1.Mailbox.Push:input_type -> libra.v1.MailboxPushRequest
	5, // 8: libra.v1.Mailbox.Pull:input_type -> libra.v1.MailboxPullRequest
	2, // 9: libra.v1.Mailbox.List:output_type -> libra.v1.MailboxListResponse
	4, // 10: libra.v1.Mailbox.Push:output_type -> libra.v1.MailboxPushResponse
	6, // 11: libra.v1.Mailbox.Pull:output_type -> libra.v1.MailboxPullResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_libra_v1_mailbox_proto_init() }
func file_libra_v1_mailbox_proto_init() {
	if File_libra_v1_mailbox_proto != nil {
		return
	}
	file_libra_v1_entry_key_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_libra_v1_mailbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mail); i {
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
		file_libra_v1_mailbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxListRequest); i {
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
		file_libra_v1_mailbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxListResponse); i {
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
		file_libra_v1_mailbox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxPushRequest); i {
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
		file_libra_v1_mailbox_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxPushResponse); i {
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
		file_libra_v1_mailbox_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxPullRequest); i {
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
		file_libra_v1_mailbox_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxPullResponse); i {
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
			RawDescriptor: file_libra_v1_mailbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_libra_v1_mailbox_proto_goTypes,
		DependencyIndexes: file_libra_v1_mailbox_proto_depIdxs,
		MessageInfos:      file_libra_v1_mailbox_proto_msgTypes,
	}.Build()
	File_libra_v1_mailbox_proto = out.File
	file_libra_v1_mailbox_proto_rawDesc = nil
	file_libra_v1_mailbox_proto_goTypes = nil
	file_libra_v1_mailbox_proto_depIdxs = nil
}
