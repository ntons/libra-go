// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: libra/v1/useradmin.proto

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

type UserAdminSetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string            `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserAdminSetMetadataRequest) Reset() {
	*x = UserAdminSetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminSetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminSetMetadataRequest) ProtoMessage() {}

func (x *UserAdminSetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminSetMetadataRequest.ProtoReflect.Descriptor instead.
func (*UserAdminSetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{0}
}

func (x *UserAdminSetMetadataRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserAdminSetMetadataRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type UserAdminSetMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserAdminSetMetadataResponse) Reset() {
	*x = UserAdminSetMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminSetMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminSetMetadataResponse) ProtoMessage() {}

func (x *UserAdminSetMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminSetMetadataResponse.ProtoReflect.Descriptor instead.
func (*UserAdminSetMetadataResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{1}
}

type UserAdminGetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserAdminGetMetadataRequest) Reset() {
	*x = UserAdminGetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminGetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminGetMetadataRequest) ProtoMessage() {}

func (x *UserAdminGetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminGetMetadataRequest.ProtoReflect.Descriptor instead.
func (*UserAdminGetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{2}
}

func (x *UserAdminGetMetadataRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserAdminGetMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserAdminGetMetadataResponse) Reset() {
	*x = UserAdminGetMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminGetMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminGetMetadataResponse) ProtoMessage() {}

func (x *UserAdminGetMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminGetMetadataResponse.ProtoReflect.Descriptor instead.
func (*UserAdminGetMetadataResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{3}
}

func (x *UserAdminGetMetadataResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type UserAdminGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids     []string                     `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Options *UserAdminGetRequest_Options `protobuf:"bytes,10,opt,name=options,proto3" json:"options,omitempty"` // get options
}

func (x *UserAdminGetRequest) Reset() {
	*x = UserAdminGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminGetRequest) ProtoMessage() {}

func (x *UserAdminGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminGetRequest.ProtoReflect.Descriptor instead.
func (*UserAdminGetRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{4}
}

func (x *UserAdminGetRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *UserAdminGetRequest) GetOptions() *UserAdminGetRequest_Options {
	if x != nil {
		return x.Options
	}
	return nil
}

type UserAdminGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserData `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Roles []*RoleData `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *UserAdminGetResponse) Reset() {
	*x = UserAdminGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminGetResponse) ProtoMessage() {}

func (x *UserAdminGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminGetResponse.ProtoReflect.Descriptor instead.
func (*UserAdminGetResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{5}
}

func (x *UserAdminGetResponse) GetUsers() []*UserData {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *UserAdminGetResponse) GetRoles() []*RoleData {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UserBanState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BanTo  int64  `protobuf:"varint,2,opt,name=ban_to,json=banTo,proto3" json:"ban_to,omitempty"`
	BanFor string `protobuf:"bytes,3,opt,name=ban_for,json=banFor,proto3" json:"ban_for,omitempty"`
}

func (x *UserBanState) Reset() {
	*x = UserBanState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBanState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBanState) ProtoMessage() {}

func (x *UserBanState) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBanState.ProtoReflect.Descriptor instead.
func (*UserBanState) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{6}
}

func (x *UserBanState) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserBanState) GetBanTo() int64 {
	if x != nil {
		return x.BanTo
	}
	return 0
}

func (x *UserBanState) GetBanFor() string {
	if x != nil {
		return x.BanFor
	}
	return ""
}

type UserAdminBanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	Seconds int64    `protobuf:"varint,2,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Reason  string   `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *UserAdminBanRequest) Reset() {
	*x = UserAdminBanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminBanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminBanRequest) ProtoMessage() {}

func (x *UserAdminBanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminBanRequest.ProtoReflect.Descriptor instead.
func (*UserAdminBanRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{7}
}

func (x *UserAdminBanRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *UserAdminBanRequest) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *UserAdminBanRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type UserAdminBanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States []*UserBanState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *UserAdminBanResponse) Reset() {
	*x = UserAdminBanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminBanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminBanResponse) ProtoMessage() {}

func (x *UserAdminBanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminBanResponse.ProtoReflect.Descriptor instead.
func (*UserAdminBanResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{8}
}

func (x *UserAdminBanResponse) GetStates() []*UserBanState {
	if x != nil {
		return x.States
	}
	return nil
}

type UserAdminBindAcctIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId                     string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`    // 目标用户
	AcctIds                    []string `protobuf:"bytes,2,rep,name=acct_ids,json=acctIds,proto3" json:"acct_ids,omitempty"` // 要转移的账号
	TakeOverAcctIdIfDuplicated bool     `protobuf:"varint,8,opt,name=take_over_acct_id_if_duplicated,json=takeOverAcctIdIfDuplicated,proto3" json:"take_over_acct_id_if_duplicated,omitempty"`
}

func (x *UserAdminBindAcctIdRequest) Reset() {
	*x = UserAdminBindAcctIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminBindAcctIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminBindAcctIdRequest) ProtoMessage() {}

func (x *UserAdminBindAcctIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminBindAcctIdRequest.ProtoReflect.Descriptor instead.
func (*UserAdminBindAcctIdRequest) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{9}
}

func (x *UserAdminBindAcctIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserAdminBindAcctIdRequest) GetAcctIds() []string {
	if x != nil {
		return x.AcctIds
	}
	return nil
}

func (x *UserAdminBindAcctIdRequest) GetTakeOverAcctIdIfDuplicated() bool {
	if x != nil {
		return x.TakeOverAcctIdIfDuplicated
	}
	return false
}

type UserAdminBindAcctIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserAdminBindAcctIdResponse) Reset() {
	*x = UserAdminBindAcctIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminBindAcctIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminBindAcctIdResponse) ProtoMessage() {}

func (x *UserAdminBindAcctIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminBindAcctIdResponse.ProtoReflect.Descriptor instead.
func (*UserAdminBindAcctIdResponse) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{10}
}

type UserAdminGetRequest_Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FuzzyId   bool `protobuf:"varint,1,opt,name=fuzzy_id,json=fuzzyId,proto3" json:"fuzzy_id,omitempty"`       // id may be user or role
	WithRoles bool `protobuf:"varint,2,opt,name=with_roles,json=withRoles,proto3" json:"with_roles,omitempty"` // get user and related roles
}

func (x *UserAdminGetRequest_Options) Reset() {
	*x = UserAdminGetRequest_Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libra_v1_useradmin_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAdminGetRequest_Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAdminGetRequest_Options) ProtoMessage() {}

func (x *UserAdminGetRequest_Options) ProtoReflect() protoreflect.Message {
	mi := &file_libra_v1_useradmin_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAdminGetRequest_Options.ProtoReflect.Descriptor instead.
func (*UserAdminGetRequest_Options) Descriptor() ([]byte, []int) {
	return file_libra_v1_useradmin_proto_rawDescGZIP(), []int{4, 0}
}

func (x *UserAdminGetRequest_Options) GetFuzzyId() bool {
	if x != nil {
		return x.FuzzyId
	}
	return false
}

func (x *UserAdminGetRequest_Options) GetWithRoles() bool {
	if x != nil {
		return x.WithRoles
	}
	return false
}

var File_libra_v1_useradmin_proto protoreflect.FileDescriptor

var file_libra_v1_useradmin_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4,
	0x01, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1e, 0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xad, 0x01,
	0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xad, 0x01,
	0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x43, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x75, 0x7a, 0x7a, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x75, 0x7a, 0x7a, 0x79, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x77, 0x69, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x6a, 0x0a,
	0x14, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x28, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x61, 0x6e,
	0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x61, 0x6e, 0x54, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x46, 0x6f, 0x72, 0x22, 0x62, 0x0a, 0x13, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x46, 0x0a,
	0x14, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x63, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x43, 0x0a, 0x1f, 0x74, 0x61, 0x6b, 0x65,
	0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x69, 0x66,
	0x5f, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x1a, 0x74, 0x61, 0x6b, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x41, 0x63, 0x63, 0x74, 0x49,
	0x64, 0x49, 0x66, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x22, 0x1d, 0x0a,
	0x1b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x63,
	0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xae, 0x03, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x5c, 0x0a, 0x0b, 0x53, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x03,
	0x42, 0x61, 0x6e, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x24, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x41,
	0x63, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74, 0x6f, 0x6e,
	0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_libra_v1_useradmin_proto_rawDescOnce sync.Once
	file_libra_v1_useradmin_proto_rawDescData = file_libra_v1_useradmin_proto_rawDesc
)

func file_libra_v1_useradmin_proto_rawDescGZIP() []byte {
	file_libra_v1_useradmin_proto_rawDescOnce.Do(func() {
		file_libra_v1_useradmin_proto_rawDescData = protoimpl.X.CompressGZIP(file_libra_v1_useradmin_proto_rawDescData)
	})
	return file_libra_v1_useradmin_proto_rawDescData
}

var file_libra_v1_useradmin_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_libra_v1_useradmin_proto_goTypes = []interface{}{
	(*UserAdminSetMetadataRequest)(nil),  // 0: libra.v1.UserAdminSetMetadataRequest
	(*UserAdminSetMetadataResponse)(nil), // 1: libra.v1.UserAdminSetMetadataResponse
	(*UserAdminGetMetadataRequest)(nil),  // 2: libra.v1.UserAdminGetMetadataRequest
	(*UserAdminGetMetadataResponse)(nil), // 3: libra.v1.UserAdminGetMetadataResponse
	(*UserAdminGetRequest)(nil),          // 4: libra.v1.UserAdminGetRequest
	(*UserAdminGetResponse)(nil),         // 5: libra.v1.UserAdminGetResponse
	(*UserBanState)(nil),                 // 6: libra.v1.UserBanState
	(*UserAdminBanRequest)(nil),          // 7: libra.v1.UserAdminBanRequest
	(*UserAdminBanResponse)(nil),         // 8: libra.v1.UserAdminBanResponse
	(*UserAdminBindAcctIdRequest)(nil),   // 9: libra.v1.UserAdminBindAcctIdRequest
	(*UserAdminBindAcctIdResponse)(nil),  // 10: libra.v1.UserAdminBindAcctIdResponse
	nil,                                  // 11: libra.v1.UserAdminSetMetadataRequest.MetadataEntry
	nil,                                  // 12: libra.v1.UserAdminGetMetadataResponse.MetadataEntry
	(*UserAdminGetRequest_Options)(nil),  // 13: libra.v1.UserAdminGetRequest.Options
	(*UserData)(nil),                     // 14: libra.v1.UserData
	(*RoleData)(nil),                     // 15: libra.v1.RoleData
}
var file_libra_v1_useradmin_proto_depIdxs = []int32{
	11, // 0: libra.v1.UserAdminSetMetadataRequest.metadata:type_name -> libra.v1.UserAdminSetMetadataRequest.MetadataEntry
	12, // 1: libra.v1.UserAdminGetMetadataResponse.metadata:type_name -> libra.v1.UserAdminGetMetadataResponse.MetadataEntry
	13, // 2: libra.v1.UserAdminGetRequest.options:type_name -> libra.v1.UserAdminGetRequest.Options
	14, // 3: libra.v1.UserAdminGetResponse.users:type_name -> libra.v1.UserData
	15, // 4: libra.v1.UserAdminGetResponse.roles:type_name -> libra.v1.RoleData
	6,  // 5: libra.v1.UserAdminBanResponse.states:type_name -> libra.v1.UserBanState
	0,  // 6: libra.v1.UserAdmin.SetMetadata:input_type -> libra.v1.UserAdminSetMetadataRequest
	2,  // 7: libra.v1.UserAdmin.GetMetadata:input_type -> libra.v1.UserAdminGetMetadataRequest
	4,  // 8: libra.v1.UserAdmin.Get:input_type -> libra.v1.UserAdminGetRequest
	7,  // 9: libra.v1.UserAdmin.Ban:input_type -> libra.v1.UserAdminBanRequest
	9,  // 10: libra.v1.UserAdmin.BindAcctId:input_type -> libra.v1.UserAdminBindAcctIdRequest
	1,  // 11: libra.v1.UserAdmin.SetMetadata:output_type -> libra.v1.UserAdminSetMetadataResponse
	3,  // 12: libra.v1.UserAdmin.GetMetadata:output_type -> libra.v1.UserAdminGetMetadataResponse
	5,  // 13: libra.v1.UserAdmin.Get:output_type -> libra.v1.UserAdminGetResponse
	8,  // 14: libra.v1.UserAdmin.Ban:output_type -> libra.v1.UserAdminBanResponse
	10, // 15: libra.v1.UserAdmin.BindAcctId:output_type -> libra.v1.UserAdminBindAcctIdResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_libra_v1_useradmin_proto_init() }
func file_libra_v1_useradmin_proto_init() {
	if File_libra_v1_useradmin_proto != nil {
		return
	}
	file_libra_v1_user_proto_init()
	file_libra_v1_role_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_libra_v1_useradmin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminSetMetadataRequest); i {
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
		file_libra_v1_useradmin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminSetMetadataResponse); i {
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
		file_libra_v1_useradmin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminGetMetadataRequest); i {
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
		file_libra_v1_useradmin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminGetMetadataResponse); i {
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
		file_libra_v1_useradmin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminGetRequest); i {
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
		file_libra_v1_useradmin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminGetResponse); i {
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
		file_libra_v1_useradmin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBanState); i {
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
		file_libra_v1_useradmin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminBanRequest); i {
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
		file_libra_v1_useradmin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminBanResponse); i {
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
		file_libra_v1_useradmin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminBindAcctIdRequest); i {
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
		file_libra_v1_useradmin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminBindAcctIdResponse); i {
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
		file_libra_v1_useradmin_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAdminGetRequest_Options); i {
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
			RawDescriptor: file_libra_v1_useradmin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_libra_v1_useradmin_proto_goTypes,
		DependencyIndexes: file_libra_v1_useradmin_proto_depIdxs,
		MessageInfos:      file_libra_v1_useradmin_proto_msgTypes,
	}.Build()
	File_libra_v1_useradmin_proto = out.File
	file_libra_v1_useradmin_proto_rawDesc = nil
	file_libra_v1_useradmin_proto_goTypes = nil
	file_libra_v1_useradmin_proto_depIdxs = nil
}
