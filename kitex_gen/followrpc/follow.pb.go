// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: follow.proto

package followrpc

import (
	context "context"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FollowCount   int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	FollowerCount int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	IsFollow      bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

type CreateFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`       // 用户id
	FollowId int64 `protobuf:"varint,2,opt,name=follow_id,json=followId,proto3" json:"follow_id,omitempty"` // 关注用户的id
}

func (x *CreateFollowRequest) Reset() {
	*x = CreateFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFollowRequest) ProtoMessage() {}

func (x *CreateFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFollowRequest.ProtoReflect.Descriptor instead.
func (*CreateFollowRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateFollowRequest) GetFollowId() int64 {
	if x != nil {
		return x.FollowId
	}
	return 0
}

type CreateFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
}

func (x *CreateFollowResponse) Reset() {
	*x = CreateFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFollowResponse) ProtoMessage() {}

func (x *CreateFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFollowResponse.ProtoReflect.Descriptor instead.
func (*CreateFollowResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFollowResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreateFollowResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

type DeleteFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FollowId int64 `protobuf:"varint,2,opt,name=follow_id,json=followId,proto3" json:"follow_id,omitempty"`
}

func (x *DeleteFollowRequest) Reset() {
	*x = DeleteFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFollowRequest) ProtoMessage() {}

func (x *DeleteFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFollowRequest.ProtoReflect.Descriptor instead.
func (*DeleteFollowRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteFollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteFollowRequest) GetFollowId() int64 {
	if x != nil {
		return x.FollowId
	}
	return 0
}

type DeleteFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
}

func (x *DeleteFollowResponse) Reset() {
	*x = DeleteFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFollowResponse) ProtoMessage() {}

func (x *DeleteFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFollowResponse.ProtoReflect.Descriptor instead.
func (*DeleteFollowResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFollowResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DeleteFollowResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

type QueryFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *QueryFollowRequest) Reset() {
	*x = QueryFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFollowRequest) ProtoMessage() {}

func (x *QueryFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFollowRequest.ProtoReflect.Descriptor instead.
func (*QueryFollowRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{5}
}

func (x *QueryFollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type QueryFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string  `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	Users         []*User `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *QueryFollowResponse) Reset() {
	*x = QueryFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFollowResponse) ProtoMessage() {}

func (x *QueryFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFollowResponse.ProtoReflect.Descriptor instead.
func (*QueryFollowResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{6}
}

func (x *QueryFollowResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *QueryFollowResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *QueryFollowResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type QueryFollowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *QueryFollowerRequest) Reset() {
	*x = QueryFollowerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFollowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFollowerRequest) ProtoMessage() {}

func (x *QueryFollowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFollowerRequest.ProtoReflect.Descriptor instead.
func (*QueryFollowerRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{7}
}

func (x *QueryFollowerRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type QueryFollowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string  `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	Users         []*User `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *QueryFollowerResponse) Reset() {
	*x = QueryFollowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFollowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFollowerResponse) ProtoMessage() {}

func (x *QueryFollowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFollowerResponse.ProtoReflect.Descriptor instead.
func (*QueryFollowerResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{8}
}

func (x *QueryFollowerResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *QueryFollowerResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *QueryFollowerResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FollowId int64 `protobuf:"varint,2,opt,name=follow_id,json=followId,proto3" json:"follow_id,omitempty"`
}

func (x *GetFollowRequest) Reset() {
	*x = GetFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowRequest) ProtoMessage() {}

func (x *GetFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowRequest.ProtoReflect.Descriptor instead.
func (*GetFollowRequest) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{9}
}

func (x *GetFollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetFollowRequest) GetFollowId() int64 {
	if x != nil {
		return x.FollowId
	}
	return 0
}

type GetFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
}

func (x *GetFollowResponse) Reset() {
	*x = GetFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowResponse) ProtoMessage() {}

func (x *GetFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follow_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowResponse.ProtoReflect.Descriptor instead.
func (*GetFollowResponse) Descriptor() ([]byte, []int) {
	return file_follow_proto_rawDescGZIP(), []int{10}
}

func (x *GetFollowResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetFollowResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

var File_follow_proto protoreflect.FileDescriptor

var file_follow_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0xa2, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0x4b, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x2f, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x15, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x48, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x87, 0x03, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x1a, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x18, 0x2e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x68, 0x30, 0x32, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x2d, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follow_proto_rawDescOnce sync.Once
	file_follow_proto_rawDescData = file_follow_proto_rawDesc
)

func file_follow_proto_rawDescGZIP() []byte {
	file_follow_proto_rawDescOnce.Do(func() {
		file_follow_proto_rawDescData = protoimpl.X.CompressGZIP(file_follow_proto_rawDescData)
	})
	return file_follow_proto_rawDescData
}

var file_follow_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_follow_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: follow.User
	(*CreateFollowRequest)(nil),   // 1: follow.CreateFollowRequest
	(*CreateFollowResponse)(nil),  // 2: follow.CreateFollowResponse
	(*DeleteFollowRequest)(nil),   // 3: follow.DeleteFollowRequest
	(*DeleteFollowResponse)(nil),  // 4: follow.DeleteFollowResponse
	(*QueryFollowRequest)(nil),    // 5: follow.QueryFollowRequest
	(*QueryFollowResponse)(nil),   // 6: follow.QueryFollowResponse
	(*QueryFollowerRequest)(nil),  // 7: follow.QueryFollowerRequest
	(*QueryFollowerResponse)(nil), // 8: follow.QueryFollowerResponse
	(*GetFollowRequest)(nil),      // 9: follow.GetFollowRequest
	(*GetFollowResponse)(nil),     // 10: follow.GetFollowResponse
}
var file_follow_proto_depIdxs = []int32{
	0,  // 0: follow.QueryFollowResponse.users:type_name -> follow.User
	0,  // 1: follow.QueryFollowerResponse.users:type_name -> follow.User
	1,  // 2: follow.FollowService.CreateFollow:input_type -> follow.CreateFollowRequest
	3,  // 3: follow.FollowService.DeleteFollow:input_type -> follow.DeleteFollowRequest
	5,  // 4: follow.FollowService.QueryFollow:input_type -> follow.QueryFollowRequest
	7,  // 5: follow.FollowService.QueryFollower:input_type -> follow.QueryFollowerRequest
	9,  // 6: follow.FollowService.GetFollow:input_type -> follow.GetFollowRequest
	2,  // 7: follow.FollowService.CreateFollow:output_type -> follow.CreateFollowResponse
	4,  // 8: follow.FollowService.DeleteFollow:output_type -> follow.DeleteFollowResponse
	6,  // 9: follow.FollowService.QueryFollow:output_type -> follow.QueryFollowResponse
	8,  // 10: follow.FollowService.QueryFollower:output_type -> follow.QueryFollowerResponse
	10, // 11: follow.FollowService.GetFollow:output_type -> follow.GetFollowResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_follow_proto_init() }
func file_follow_proto_init() {
	if File_follow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_follow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFollowRequest); i {
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
		file_follow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFollowResponse); i {
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
		file_follow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFollowRequest); i {
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
		file_follow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFollowResponse); i {
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
		file_follow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFollowRequest); i {
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
		file_follow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFollowResponse); i {
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
		file_follow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFollowerRequest); i {
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
		file_follow_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFollowerResponse); i {
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
		file_follow_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowRequest); i {
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
		file_follow_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowResponse); i {
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
			RawDescriptor: file_follow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follow_proto_goTypes,
		DependencyIndexes: file_follow_proto_depIdxs,
		MessageInfos:      file_follow_proto_msgTypes,
	}.Build()
	File_follow_proto = out.File
	file_follow_proto_rawDesc = nil
	file_follow_proto_goTypes = nil
	file_follow_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.3.1. DO NOT EDIT.

type FollowService interface {
	CreateFollow(ctx context.Context, req *CreateFollowRequest) (res *CreateFollowResponse, err error)
	DeleteFollow(ctx context.Context, req *DeleteFollowRequest) (res *DeleteFollowResponse, err error)
	QueryFollow(ctx context.Context, req *QueryFollowRequest) (res *QueryFollowResponse, err error)
	QueryFollower(ctx context.Context, req *QueryFollowerRequest) (res *QueryFollowerResponse, err error)
	GetFollow(ctx context.Context, req *GetFollowRequest) (res *GetFollowResponse, err error)
}
