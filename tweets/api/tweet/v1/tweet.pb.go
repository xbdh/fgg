// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api/tweet/v1/tweet.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TweetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId  uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	//  uint64 likesCount=4;
	//  uint64 commentsCount=5;
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *TweetInfo) Reset() {
	*x = TweetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tweet_v1_tweet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TweetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TweetInfo) ProtoMessage() {}

func (x *TweetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_tweet_v1_tweet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TweetInfo.ProtoReflect.Descriptor instead.
func (*TweetInfo) Descriptor() ([]byte, []int) {
	return file_api_tweet_v1_tweet_proto_rawDescGZIP(), []int{0}
}

func (x *TweetInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TweetInfo) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TweetInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TweetInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tweet_v1_tweet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tweet_v1_tweet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_api_tweet_v1_tweet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetInfo *TweetInfo `protobuf:"bytes,1,opt,name=tweetInfo,proto3" json:"tweetInfo,omitempty"`
}

func (x *CreateReply) Reset() {
	*x = CreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tweet_v1_tweet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReply) ProtoMessage() {}

func (x *CreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_tweet_v1_tweet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReply.ProtoReflect.Descriptor instead.
func (*CreateReply) Descriptor() ([]byte, []int) {
	return file_api_tweet_v1_tweet_proto_rawDescGZIP(), []int{2}
}

func (x *CreateReply) GetTweetInfo() *TweetInfo {
	if x != nil {
		return x.TweetInfo
	}
	return nil
}

type UserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserIdRequest) Reset() {
	*x = UserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tweet_v1_tweet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdRequest) ProtoMessage() {}

func (x *UserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tweet_v1_tweet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdRequest.ProtoReflect.Descriptor instead.
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return file_api_tweet_v1_tweet_proto_rawDescGZIP(), []int{3}
}

func (x *UserIdRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetInfo []*TweetInfo `protobuf:"bytes,1,rep,name=tweetInfo,proto3" json:"tweetInfo,omitempty"`
}

func (x *UserIdReply) Reset() {
	*x = UserIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tweet_v1_tweet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdReply) ProtoMessage() {}

func (x *UserIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_tweet_v1_tweet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdReply.ProtoReflect.Descriptor instead.
func (*UserIdReply) Descriptor() ([]byte, []int) {
	return file_api_tweet_v1_tweet_proto_rawDescGZIP(), []int{4}
}

func (x *UserIdReply) GetTweetInfo() []*TweetInfo {
	if x != nil {
		return x.TweetInfo
	}
	return nil
}

type TweetIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetId uint64 `protobuf:"varint,1,opt,name=tweetId,proto3" json:"tweetId,omitempty"`
}

func (x *TweetIdRequest) Reset() {
	*x = TweetIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tweet_v1_tweet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TweetIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TweetIdRequest) ProtoMessage() {}

func (x *TweetIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tweet_v1_tweet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TweetIdRequest.ProtoReflect.Descriptor instead.
func (*TweetIdRequest) Descriptor() ([]byte, []int) {
	return file_api_tweet_v1_tweet_proto_rawDescGZIP(), []int{5}
}

func (x *TweetIdRequest) GetTweetId() uint64 {
	if x != nil {
		return x.TweetId
	}
	return 0
}

type TweetIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetInfo *TweetInfo `protobuf:"bytes,1,opt,name=tweetInfo,proto3" json:"tweetInfo,omitempty"`
}

func (x *TweetIdReply) Reset() {
	*x = TweetIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tweet_v1_tweet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TweetIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TweetIdReply) ProtoMessage() {}

func (x *TweetIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_tweet_v1_tweet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TweetIdReply.ProtoReflect.Descriptor instead.
func (*TweetIdReply) Descriptor() ([]byte, []int) {
	return file_api_tweet_v1_tweet_proto_rawDescGZIP(), []int{6}
}

func (x *TweetIdReply) GetTweetInfo() *TweetInfo {
	if x != nil {
		return x.TweetInfo
	}
	return nil
}

var File_api_tweet_v1_tweet_proto protoreflect.FileDescriptor

var file_api_tweet_v1_tweet_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x77, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x77, 0x65, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x41, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x40, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x31, 0x0a, 0x09, 0x74, 0x77, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x77, 0x65, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x27, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x09,
	0x74, 0x77, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x77, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x2a, 0x0a, 0x0e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x74, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0c, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x09, 0x74,
	0x77, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x77, 0x65, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xd9,
	0x01, 0x0a, 0x06, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x74, 0x77, 0x65, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x42, 0x79, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x77, 0x65, 0x65,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x46, 0x0a, 0x17, 0x64, 0x65,
	0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x77, 0x65,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x16, 0x74, 0x77, 0x65, 0x65,
	0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_tweet_v1_tweet_proto_rawDescOnce sync.Once
	file_api_tweet_v1_tweet_proto_rawDescData = file_api_tweet_v1_tweet_proto_rawDesc
)

func file_api_tweet_v1_tweet_proto_rawDescGZIP() []byte {
	file_api_tweet_v1_tweet_proto_rawDescOnce.Do(func() {
		file_api_tweet_v1_tweet_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_tweet_v1_tweet_proto_rawDescData)
	})
	return file_api_tweet_v1_tweet_proto_rawDescData
}

var file_api_tweet_v1_tweet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_tweet_v1_tweet_proto_goTypes = []interface{}{
	(*TweetInfo)(nil),             // 0: tweet.v1.TweetInfo
	(*CreateRequest)(nil),         // 1: tweet.v1.CreateRequest
	(*CreateReply)(nil),           // 2: tweet.v1.CreateReply
	(*UserIdRequest)(nil),         // 3: tweet.v1.UserIdRequest
	(*UserIdReply)(nil),           // 4: tweet.v1.UserIdReply
	(*TweetIdRequest)(nil),        // 5: tweet.v1.TweetIdRequest
	(*TweetIdReply)(nil),          // 6: tweet.v1.TweetIdReply
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_api_tweet_v1_tweet_proto_depIdxs = []int32{
	7, // 0: tweet.v1.TweetInfo.createdAt:type_name -> google.protobuf.Timestamp
	0, // 1: tweet.v1.CreateReply.tweetInfo:type_name -> tweet.v1.TweetInfo
	0, // 2: tweet.v1.UserIdReply.tweetInfo:type_name -> tweet.v1.TweetInfo
	0, // 3: tweet.v1.TweetIdReply.tweetInfo:type_name -> tweet.v1.TweetInfo
	1, // 4: tweet.v1.Tweets.CreateTweets:input_type -> tweet.v1.CreateRequest
	3, // 5: tweet.v1.Tweets.GetTweetByUserId:input_type -> tweet.v1.UserIdRequest
	5, // 6: tweet.v1.Tweets.GetTweetByTweetId:input_type -> tweet.v1.TweetIdRequest
	2, // 7: tweet.v1.Tweets.CreateTweets:output_type -> tweet.v1.CreateReply
	4, // 8: tweet.v1.Tweets.GetTweetByUserId:output_type -> tweet.v1.UserIdReply
	6, // 9: tweet.v1.Tweets.GetTweetByTweetId:output_type -> tweet.v1.TweetIdReply
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_tweet_v1_tweet_proto_init() }
func file_api_tweet_v1_tweet_proto_init() {
	if File_api_tweet_v1_tweet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_tweet_v1_tweet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TweetInfo); i {
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
		file_api_tweet_v1_tweet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_api_tweet_v1_tweet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReply); i {
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
		file_api_tweet_v1_tweet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdRequest); i {
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
		file_api_tweet_v1_tweet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdReply); i {
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
		file_api_tweet_v1_tweet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TweetIdRequest); i {
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
		file_api_tweet_v1_tweet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TweetIdReply); i {
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
			RawDescriptor: file_api_tweet_v1_tweet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_tweet_v1_tweet_proto_goTypes,
		DependencyIndexes: file_api_tweet_v1_tweet_proto_depIdxs,
		MessageInfos:      file_api_tweet_v1_tweet_proto_msgTypes,
	}.Build()
	File_api_tweet_v1_tweet_proto = out.File
	file_api_tweet_v1_tweet_proto_rawDesc = nil
	file_api_tweet_v1_tweet_proto_goTypes = nil
	file_api_tweet_v1_tweet_proto_depIdxs = nil
}
