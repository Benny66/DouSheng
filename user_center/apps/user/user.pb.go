// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: apps/user/pb/user.proto

package user

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

// 用户注册 & 登录 的接口请求 model
type LoginAndRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名【注册的话，最长32个字符】
	// @gotags: json:"username" form:"username" binding:"required,max=32" validate:"required,max=32"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" form:"username" binding:"required,max=32" validate:"required,max=32"`
	// 密码
	// @gotags: json:"password" form:"password" binding:"required,max=32" validate:"required,max=32"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password" form:"password" binding:"required,max=32" validate:"required,max=32"`
}

func (x *LoginAndRegisterRequest) Reset() {
	*x = LoginAndRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAndRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAndRegisterRequest) ProtoMessage() {}

func (x *LoginAndRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAndRegisterRequest.ProtoReflect.Descriptor instead.
func (*LoginAndRegisterRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginAndRegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginAndRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 用户注册 & 登录 的接口响应 model
type TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	// @gotags: json:"user_id"
	UserId int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id"`
	// 用户鉴权Token
	// @gotags: json:"token"
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token"`
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *TokenResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 获取用户信息 的接口请求 model
type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	// @gotags: json:"user_id" form:"user_id" validate:"required" binding:"required"
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id" form:"user_id" validate:"required" binding:"required"`
	// 用户鉴权Token
	// @gotags: json:"token" form:"token" binding:"required"
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token" form:"token" binding:"required"`
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 获取用户信息 的接口响应 model
type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户信息
	// @gotags: json:"user"
	User *User `protobuf:"bytes,3,opt,name=user,proto3" json:"user"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfoResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// 调用用户信息 时返回的User
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	// @gotags: json:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 用户名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// 关注总数
	// @gotags: json:"follow_count"
	FollowCount *int64 `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3,oneof" json:"follow_count"`
	// 粉丝总数
	// @gotags: json:"follower_count"
	FollowerCount *int64 `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3,oneof" json:"follower_count"`
	// 已关注-true，false-未关注
	IsFollow bool `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[4]
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
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil && x.FollowCount != nil {
		return *x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil && x.FollowerCount != nil {
		return *x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

// 与数据库对应的 PO 对象
type UserPo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	// @gotags: json:"id" gorm:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" gorm:"id"`
	// 用户名称
	// @gotags: json:"username" gorm:"username"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username" gorm:"username"`
	// 用户名称
	// @gotags: json:"password" gorm:"password"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password" gorm:"password"`
}

func (x *UserPo) Reset() {
	*x = UserPo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPo) ProtoMessage() {}

func (x *UserPo) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPo.ProtoReflect.Descriptor instead.
func (*UserPo) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserPo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserPo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserPo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 用户IDs
type UserMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID 列表
	// @gotags: json:"user_ids"
	UserIds []int64 `protobuf:"varint,1,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids"`
}

func (x *UserMapRequest) Reset() {
	*x = UserMapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMapRequest) ProtoMessage() {}

func (x *UserMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMapRequest.ProtoReflect.Descriptor instead.
func (*UserMapRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserMapRequest) GetUserIds() []int64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// 用户列表：map[userId] = User
type UserMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户列表：map[userId] = User
	// @gotags: json:"user_map"
	UserMap map[int64]*User `protobuf:"bytes,1,rep,name=user_map,json=userMap,proto3" json:"user_map" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserMapResponse) Reset() {
	*x = UserMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMapResponse) ProtoMessage() {}

func (x *UserMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMapResponse.ProtoReflect.Descriptor instead.
func (*UserMapResponse) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserMapResponse) GetUserMap() map[int64]*User {
	if x != nil {
		return x.UserMap
	}
	return nil
}

var File_apps_user_pb_user_proto protoreflect.FileDescriptor

var file_apps_user_pb_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x6f, 0x75, 0x73, 0x68,
	0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x17, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3e, 0x0a, 0x0d, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x40, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3b, 0x0a,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xbf, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x2a, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b,
	0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x1a, 0x4f, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68,
	0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xc1, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x26, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68,
	0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x26, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65,
	0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1e, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x1d, 0x2e,
	0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64,
	0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x2d, 0x54, 0x6f,
	0x2d, 0x42, 0x79, 0x74, 0x65, 0x2f, 0x44, 0x6f, 0x75, 0x53, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_user_pb_user_proto_rawDescOnce sync.Once
	file_apps_user_pb_user_proto_rawDescData = file_apps_user_pb_user_proto_rawDesc
)

func file_apps_user_pb_user_proto_rawDescGZIP() []byte {
	file_apps_user_pb_user_proto_rawDescOnce.Do(func() {
		file_apps_user_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_user_pb_user_proto_rawDescData)
	})
	return file_apps_user_pb_user_proto_rawDescData
}

var file_apps_user_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_user_pb_user_proto_goTypes = []interface{}{
	(*LoginAndRegisterRequest)(nil), // 0: dousheng.user.LoginAndRegisterRequest
	(*TokenResponse)(nil),           // 1: dousheng.user.TokenResponse
	(*UserInfoRequest)(nil),         // 2: dousheng.user.UserInfoRequest
	(*UserInfoResponse)(nil),        // 3: dousheng.user.UserInfoResponse
	(*User)(nil),                    // 4: dousheng.user.User
	(*UserPo)(nil),                  // 5: dousheng.user.UserPo
	(*UserMapRequest)(nil),          // 6: dousheng.user.UserMapRequest
	(*UserMapResponse)(nil),         // 7: dousheng.user.UserMapResponse
	nil,                             // 8: dousheng.user.UserMapResponse.UserMapEntry
}
var file_apps_user_pb_user_proto_depIdxs = []int32{
	4, // 0: dousheng.user.UserInfoResponse.user:type_name -> dousheng.user.User
	8, // 1: dousheng.user.UserMapResponse.user_map:type_name -> dousheng.user.UserMapResponse.UserMapEntry
	4, // 2: dousheng.user.UserMapResponse.UserMapEntry.value:type_name -> dousheng.user.User
	0, // 3: dousheng.user.Service.Register:input_type -> dousheng.user.LoginAndRegisterRequest
	0, // 4: dousheng.user.Service.Login:input_type -> dousheng.user.LoginAndRegisterRequest
	2, // 5: dousheng.user.Service.UserInfo:input_type -> dousheng.user.UserInfoRequest
	6, // 6: dousheng.user.Service.UserMap:input_type -> dousheng.user.UserMapRequest
	1, // 7: dousheng.user.Service.Register:output_type -> dousheng.user.TokenResponse
	1, // 8: dousheng.user.Service.Login:output_type -> dousheng.user.TokenResponse
	3, // 9: dousheng.user.Service.UserInfo:output_type -> dousheng.user.UserInfoResponse
	7, // 10: dousheng.user.Service.UserMap:output_type -> dousheng.user.UserMapResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_user_pb_user_proto_init() }
func file_apps_user_pb_user_proto_init() {
	if File_apps_user_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_user_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAndRegisterRequest); i {
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
		file_apps_user_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse); i {
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
		file_apps_user_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRequest); i {
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
		file_apps_user_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
		file_apps_user_pb_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_user_pb_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPo); i {
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
		file_apps_user_pb_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMapRequest); i {
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
		file_apps_user_pb_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMapResponse); i {
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
	file_apps_user_pb_user_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_user_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_user_pb_user_proto_goTypes,
		DependencyIndexes: file_apps_user_pb_user_proto_depIdxs,
		MessageInfos:      file_apps_user_pb_user_proto_msgTypes,
	}.Build()
	File_apps_user_pb_user_proto = out.File
	file_apps_user_pb_user_proto_rawDesc = nil
	file_apps_user_pb_user_proto_goTypes = nil
	file_apps_user_pb_user_proto_depIdxs = nil
}
