// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: v1/mailing.proto

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

type Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Mail) Reset() {
	*x = Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mailing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail) ProtoMessage() {}

func (x *Mail) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mailing_proto_msgTypes[0]
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
	return file_v1_mailing_proto_rawDescGZIP(), []int{0}
}

func (x *Mail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Mail) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MailingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *EntityKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *MailingListRequest) Reset() {
	*x = MailingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mailing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailingListRequest) ProtoMessage() {}

func (x *MailingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mailing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailingListRequest.ProtoReflect.Descriptor instead.
func (*MailingListRequest) Descriptor() ([]byte, []int) {
	return file_v1_mailing_proto_rawDescGZIP(), []int{1}
}

func (x *MailingListRequest) GetKey() *EntityKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type MailingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mails []*Mail `protobuf:"bytes,1,rep,name=mails,proto3" json:"mails,omitempty"`
}

func (x *MailingListResponse) Reset() {
	*x = MailingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mailing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailingListResponse) ProtoMessage() {}

func (x *MailingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mailing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailingListResponse.ProtoReflect.Descriptor instead.
func (*MailingListResponse) Descriptor() ([]byte, []int) {
	return file_v1_mailing_proto_rawDescGZIP(), []int{2}
}

func (x *MailingListResponse) GetMails() []*Mail {
	if x != nil {
		return x.Mails
	}
	return nil
}

type MailingPushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      *EntityKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Content  string     `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Capacity int32      `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"` // maximum mails to keep
}

func (x *MailingPushRequest) Reset() {
	*x = MailingPushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mailing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailingPushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailingPushRequest) ProtoMessage() {}

func (x *MailingPushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mailing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailingPushRequest.ProtoReflect.Descriptor instead.
func (*MailingPushRequest) Descriptor() ([]byte, []int) {
	return file_v1_mailing_proto_rawDescGZIP(), []int{3}
}

func (x *MailingPushRequest) GetKey() *EntityKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *MailingPushRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MailingPushRequest) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

type MailingPushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailId string `protobuf:"bytes,1,opt,name=mail_id,json=mailId,proto3" json:"mail_id,omitempty"`
}

func (x *MailingPushResponse) Reset() {
	*x = MailingPushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mailing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailingPushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailingPushResponse) ProtoMessage() {}

func (x *MailingPushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mailing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailingPushResponse.ProtoReflect.Descriptor instead.
func (*MailingPushResponse) Descriptor() ([]byte, []int) {
	return file_v1_mailing_proto_rawDescGZIP(), []int{4}
}

func (x *MailingPushResponse) GetMailId() string {
	if x != nil {
		return x.MailId
	}
	return ""
}

type MailingPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *EntityKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Ids []string   `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *MailingPullRequest) Reset() {
	*x = MailingPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mailing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailingPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailingPullRequest) ProtoMessage() {}

func (x *MailingPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mailing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailingPullRequest.ProtoReflect.Descriptor instead.
func (*MailingPullRequest) Descriptor() ([]byte, []int) {
	return file_v1_mailing_proto_rawDescGZIP(), []int{5}
}

func (x *MailingPullRequest) GetKey() *EntityKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *MailingPullRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type MailingPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PulledIds []string `protobuf:"bytes,1,rep,name=pulled_ids,json=pulledIds,proto3" json:"pulled_ids,omitempty"` // pulled id list
}

func (x *MailingPullResponse) Reset() {
	*x = MailingPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mailing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailingPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailingPullResponse) ProtoMessage() {}

func (x *MailingPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mailing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailingPullResponse.ProtoReflect.Descriptor instead.
func (*MailingPullResponse) Descriptor() ([]byte, []int) {
	return file_v1_mailing_proto_rawDescGZIP(), []int{6}
}

func (x *MailingPullResponse) GetPulledIds() []string {
	if x != nil {
		return x.PulledIds
	}
	return nil
}

var File_v1_mailing_proto protoreflect.FileDescriptor

var file_v1_mailing_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x0f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a,
	0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x3b, 0x0a, 0x12, 0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3b, 0x0a, 0x13,
	0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x05, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x71, 0x0a, 0x12, 0x4d, 0x61, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65,
	0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x2e, 0x0a, 0x13,
	0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x12,
	0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x34, 0x0a, 0x13, 0x4d,
	0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x6c, 0x6c, 0x65, 0x64, 0x49, 0x64,
	0x73, 0x32, 0xd8, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x43, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x75, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12,
	0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x69,
	0x6e, 0x67, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x6f, 0x6e, 0x73,
	0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_mailing_proto_rawDescOnce sync.Once
	file_v1_mailing_proto_rawDescData = file_v1_mailing_proto_rawDesc
)

func file_v1_mailing_proto_rawDescGZIP() []byte {
	file_v1_mailing_proto_rawDescOnce.Do(func() {
		file_v1_mailing_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_mailing_proto_rawDescData)
	})
	return file_v1_mailing_proto_rawDescData
}

var file_v1_mailing_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_mailing_proto_goTypes = []interface{}{
	(*Mail)(nil),                // 0: libra.v1.Mail
	(*MailingListRequest)(nil),  // 1: libra.v1.MailingListRequest
	(*MailingListResponse)(nil), // 2: libra.v1.MailingListResponse
	(*MailingPushRequest)(nil),  // 3: libra.v1.MailingPushRequest
	(*MailingPushResponse)(nil), // 4: libra.v1.MailingPushResponse
	(*MailingPullRequest)(nil),  // 5: libra.v1.MailingPullRequest
	(*MailingPullResponse)(nil), // 6: libra.v1.MailingPullResponse
	(*EntityKey)(nil),           // 7: libra.v1.EntityKey
}
var file_v1_mailing_proto_depIdxs = []int32{
	7, // 0: libra.v1.MailingListRequest.key:type_name -> libra.v1.EntityKey
	0, // 1: libra.v1.MailingListResponse.mails:type_name -> libra.v1.Mail
	7, // 2: libra.v1.MailingPushRequest.key:type_name -> libra.v1.EntityKey
	7, // 3: libra.v1.MailingPullRequest.key:type_name -> libra.v1.EntityKey
	1, // 4: libra.v1.Mailing.List:input_type -> libra.v1.MailingListRequest
	3, // 5: libra.v1.Mailing.Push:input_type -> libra.v1.MailingPushRequest
	5, // 6: libra.v1.Mailing.Pull:input_type -> libra.v1.MailingPullRequest
	2, // 7: libra.v1.Mailing.List:output_type -> libra.v1.MailingListResponse
	4, // 8: libra.v1.Mailing.Push:output_type -> libra.v1.MailingPushResponse
	6, // 9: libra.v1.Mailing.Pull:output_type -> libra.v1.MailingPullResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_mailing_proto_init() }
func file_v1_mailing_proto_init() {
	if File_v1_mailing_proto != nil {
		return
	}
	file_v1_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_mailing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_mailing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailingListRequest); i {
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
		file_v1_mailing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailingListResponse); i {
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
		file_v1_mailing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailingPushRequest); i {
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
		file_v1_mailing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailingPushResponse); i {
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
		file_v1_mailing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailingPullRequest); i {
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
		file_v1_mailing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailingPullResponse); i {
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
			RawDescriptor: file_v1_mailing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_mailing_proto_goTypes,
		DependencyIndexes: file_v1_mailing_proto_depIdxs,
		MessageInfos:      file_v1_mailing_proto_msgTypes,
	}.Build()
	File_v1_mailing_proto = out.File
	file_v1_mailing_proto_rawDesc = nil
	file_v1_mailing_proto_goTypes = nil
	file_v1_mailing_proto_depIdxs = nil
}