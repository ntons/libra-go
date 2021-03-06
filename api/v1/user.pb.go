// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.0
// source: v1/user.proto

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

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AcctId   []string          `protobuf:"bytes,4,rep,name=acct_id,json=acctId,proto3" json:"acct_id,omitempty"`
	Metadata map[string]string `protobuf:"bytes,15,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserData) GetAcctId() []string {
	if x != nil {
		return x.AcctId
	}
	return nil
}

func (x *UserData) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type UserLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string     `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	State *anypb.Any `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"` // depending on implementation
}

func (x *UserLoginRequest) Reset() {
	*x = UserLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRequest) ProtoMessage() {}

func (x *UserLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRequest.ProtoReflect.Descriptor instead.
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoginRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *UserLoginRequest) GetState() *anypb.Any {
	if x != nil {
		return x.State
	}
	return nil
}

type UserLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserData `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserLoginResponse) Reset() {
	*x = UserLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResponse) ProtoMessage() {}

func (x *UserLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResponse.ProtoReflect.Descriptor instead.
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserLoginResponse) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

type UserLogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserLogoutRequest) Reset() {
	*x = UserLogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogoutRequest) ProtoMessage() {}

func (x *UserLogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogoutRequest.ProtoReflect.Descriptor instead.
func (*UserLogoutRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{3}
}

type UserLogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserLogoutResponse) Reset() {
	*x = UserLogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogoutResponse) ProtoMessage() {}

func (x *UserLogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogoutResponse.ProtoReflect.Descriptor instead.
func (*UserLogoutResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{4}
}

type UserBindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcctId string `protobuf:"bytes,1,opt,name=acct_id,json=acctId,proto3" json:"acct_id,omitempty"`
}

func (x *UserBindRequest) Reset() {
	*x = UserBindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBindRequest) ProtoMessage() {}

func (x *UserBindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBindRequest.ProtoReflect.Descriptor instead.
func (*UserBindRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserBindRequest) GetAcctId() string {
	if x != nil {
		return x.AcctId
	}
	return ""
}

type UserBindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserBindResponse) Reset() {
	*x = UserBindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBindResponse) ProtoMessage() {}

func (x *UserBindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBindResponse.ProtoReflect.Descriptor instead.
func (*UserBindResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{6}
}

type UserSetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserSetMetadataRequest) Reset() {
	*x = UserSetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSetMetadataRequest) ProtoMessage() {}

func (x *UserSetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSetMetadataRequest.ProtoReflect.Descriptor instead.
func (*UserSetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserSetMetadataRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type UserSetMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserSetMetadataResponse) Reset() {
	*x = UserSetMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSetMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSetMetadataResponse) ProtoMessage() {}

func (x *UserSetMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSetMetadataResponse.ProtoReflect.Descriptor instead.
func (*UserSetMetadataResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{8}
}

// login state for development, no security mechanism required
type DevLoginState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *DevLoginState) Reset() {
	*x = DevLoginState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevLoginState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevLoginState) ProtoMessage() {}

func (x *DevLoginState) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevLoginState.ProtoReflect.Descriptor instead.
func (*DevLoginState) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{9}
}

func (x *DevLoginState) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// Librad not implement any account (check) system, the real account state
// should be verified via a third-party server, then send a login request
// to librad to fetch passport(token).
type UniformLoginState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User which matchs one of the acct_id will be selected,
	// then all of the acct_id will be bind to this user.
	AcctId []string `protobuf:"bytes,1,rep,name=acct_id,json=acctId,proto3" json:"acct_id,omitempty"`
	// timestamp when signature signed
	Timestamp int64 `protobuf:"varint,13,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// random string max length to 32 bytes
	// 16 bytes length at least is recommended
	Nonce string `protobuf:"bytes,14,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// the signature should be signed at SERVER SIDE,
	// no matter the login request is sent from client or server.
	Signature string `protobuf:"bytes,15,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *UniformLoginState) Reset() {
	*x = UniformLoginState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniformLoginState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniformLoginState) ProtoMessage() {}

func (x *UniformLoginState) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniformLoginState.ProtoReflect.Descriptor instead.
func (*UniformLoginState) Descriptor() ([]byte, []int) {
	return file_v1_user_proto_rawDescGZIP(), []int{10}
}

func (x *UniformLoginState) GetAcctId() []string {
	if x != nil {
		return x.AcctId
	}
	return nil
}

func (x *UniformLoginState) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *UniformLoginState) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *UniformLoginState) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_v1_user_proto protoreflect.FileDescriptor

var file_v1_user_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x55, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14,
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x12, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x7e, 0x0a, 0x11, 0x55, 0x6e, 0x69, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x32, 0xa0, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x04, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x74, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x2d, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_user_proto_rawDescOnce sync.Once
	file_v1_user_proto_rawDescData = file_v1_user_proto_rawDesc
)

func file_v1_user_proto_rawDescGZIP() []byte {
	file_v1_user_proto_rawDescOnce.Do(func() {
		file_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_user_proto_rawDescData)
	})
	return file_v1_user_proto_rawDescData
}

var file_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_v1_user_proto_goTypes = []interface{}{
	(*UserData)(nil),                // 0: libra.v1.UserData
	(*UserLoginRequest)(nil),        // 1: libra.v1.UserLoginRequest
	(*UserLoginResponse)(nil),       // 2: libra.v1.UserLoginResponse
	(*UserLogoutRequest)(nil),       // 3: libra.v1.UserLogoutRequest
	(*UserLogoutResponse)(nil),      // 4: libra.v1.UserLogoutResponse
	(*UserBindRequest)(nil),         // 5: libra.v1.UserBindRequest
	(*UserBindResponse)(nil),        // 6: libra.v1.UserBindResponse
	(*UserSetMetadataRequest)(nil),  // 7: libra.v1.UserSetMetadataRequest
	(*UserSetMetadataResponse)(nil), // 8: libra.v1.UserSetMetadataResponse
	(*DevLoginState)(nil),           // 9: libra.v1.DevLoginState
	(*UniformLoginState)(nil),       // 10: libra.v1.UniformLoginState
	nil,                             // 11: libra.v1.UserData.MetadataEntry
	nil,                             // 12: libra.v1.UserSetMetadataRequest.MetadataEntry
	(*anypb.Any)(nil),               // 13: google.protobuf.Any
}
var file_v1_user_proto_depIdxs = []int32{
	11, // 0: libra.v1.UserData.metadata:type_name -> libra.v1.UserData.MetadataEntry
	13, // 1: libra.v1.UserLoginRequest.state:type_name -> google.protobuf.Any
	0,  // 2: libra.v1.UserLoginResponse.user:type_name -> libra.v1.UserData
	12, // 3: libra.v1.UserSetMetadataRequest.metadata:type_name -> libra.v1.UserSetMetadataRequest.MetadataEntry
	1,  // 4: libra.v1.User.Login:input_type -> libra.v1.UserLoginRequest
	3,  // 5: libra.v1.User.Logout:input_type -> libra.v1.UserLogoutRequest
	5,  // 6: libra.v1.User.Bind:input_type -> libra.v1.UserBindRequest
	7,  // 7: libra.v1.User.SetMetadata:input_type -> libra.v1.UserSetMetadataRequest
	2,  // 8: libra.v1.User.Login:output_type -> libra.v1.UserLoginResponse
	4,  // 9: libra.v1.User.Logout:output_type -> libra.v1.UserLogoutResponse
	6,  // 10: libra.v1.User.Bind:output_type -> libra.v1.UserBindResponse
	8,  // 11: libra.v1.User.SetMetadata:output_type -> libra.v1.UserSetMetadataResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_v1_user_proto_init() }
func file_v1_user_proto_init() {
	if File_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRequest); i {
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
		file_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginResponse); i {
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
		file_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLogoutRequest); i {
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
		file_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLogoutResponse); i {
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
		file_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBindRequest); i {
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
		file_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBindResponse); i {
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
		file_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSetMetadataRequest); i {
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
		file_v1_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSetMetadataResponse); i {
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
		file_v1_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevLoginState); i {
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
		file_v1_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniformLoginState); i {
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
			RawDescriptor: file_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_user_proto_goTypes,
		DependencyIndexes: file_v1_user_proto_depIdxs,
		MessageInfos:      file_v1_user_proto_msgTypes,
	}.Build()
	File_v1_user_proto = out.File
	file_v1_user_proto_rawDesc = nil
	file_v1_user_proto_goTypes = nil
	file_v1_user_proto_depIdxs = nil
}
