// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: feed.proto

package feed

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "offer_tiktok/biz/model/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DouyinFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime *int64  `protobuf:"varint,1,opt,name=latest_time,json=latestTime" json:"latest_time,omitempty" form:"latest_time" query:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty" form:"token" query:"token"`                                          // 可选参数，登录用户设置
}

func (x *DouyinFeedRequest) Reset() {
	*x = DouyinFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFeedRequest) ProtoMessage() {}

func (x *DouyinFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFeedRequest.ProtoReflect.Descriptor instead.
func (*DouyinFeedRequest) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinFeedRequest) GetLatestTime() int64 {
	if x != nil && x.LatestTime != nil {
		return *x.LatestTime
	}
	return 0
}

func (x *DouyinFeedRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type DouyinFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list" form:"video_list" query:"video_list"`                                  // 视频列表
	NextTime   *int64   `protobuf:"varint,4,opt,name=next_time,json=nextTime" json:"next_time,omitempty" form:"next_time" query:"next_time"`                            // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

func (x *DouyinFeedResponse) Reset() {
	*x = DouyinFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFeedResponse) ProtoMessage() {}

func (x *DouyinFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFeedResponse.ProtoReflect.Descriptor instead.
func (*DouyinFeedResponse) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinFeedResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinFeedResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *DouyinFeedResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *DouyinFeedResponse) GetNextTime() int64 {
	if x != nil && x.NextTime != nil {
		return *x.NextTime
	}
	return 0
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                                    // 视频唯一标识
	Author        *User   `protobuf:"bytes,2,req,name=author" json:"author,required" form:"author,required" query:"author,required"`                                                     // 视频作者信息
	PlayUrl       *string `protobuf:"bytes,3,req,name=play_url,json=playUrl" json:"play_url,required" form:"play_url,required" query:"play_url,required"`                                // 视频播放地址
	CoverUrl      *string `protobuf:"bytes,4,req,name=cover_url,json=coverUrl" json:"cover_url,required" form:"cover_url,required" query:"cover_url,required"`                           // 视频封面地址
	FavoriteCount *int64  `protobuf:"varint,5,req,name=favorite_count,json=favoriteCount" json:"favorite_count,required" form:"favorite_count,required" query:"favorite_count,required"` // 视频的点赞总数
	CommentCount  *int64  `protobuf:"varint,6,req,name=comment_count,json=commentCount" json:"comment_count,required" form:"comment_count,required" query:"comment_count,required"`      // 视频的评论总数
	IsFavorite    *bool   `protobuf:"varint,7,req,name=is_favorite,json=isFavorite" json:"is_favorite,required" form:"is_favorite,required" query:"is_favorite,required"`                // true-已点赞，false-未点赞
	Title         *string `protobuf:"bytes,8,req,name=title" json:"title,required" form:"title,required" query:"title,required"`                                                         // 视频标题
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{2}
}

func (x *Video) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil && x.PlayUrl != nil {
		return *x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil && x.CoverUrl != nil {
		return *x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil && x.FavoriteCount != nil {
		return *x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil && x.CommentCount != nil {
		return *x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil && x.IsFavorite != nil {
		return *x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                   // 用户id
	Name          *string `protobuf:"bytes,2,req,name=name" json:"name,required" form:"name,required" query:"name,required"`                                            // 用户名称
	FollowCount   *int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count,omitempty" form:"follow_count" query:"follow_count"`           // 关注总数
	FollowerCount *int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count,omitempty" form:"follower_count" query:"follower_count"` // 粉丝总数
	IsFollow      *bool   `protobuf:"varint,5,req,name=is_follow,json=isFollow" json:"is_follow,required" form:"is_follow,required" query:"is_follow,required"`         // true-已关注，false-未关注
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[3]
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
	return file_feed_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
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
	if x != nil && x.IsFollow != nil {
		return *x.IsFollow
	}
	return false
}

var File_feed_proto protoreflect.FileDescriptor

var file_feed_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x65,
	0x65, 0x64, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a,
	0x13, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9f, 0x01, 0x0a, 0x14,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xf6, 0x01,
	0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x02, 0x28,
	0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07,
	0x20, 0x02, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x02, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x32, 0x60, 0x0a, 0x0c, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x66, 0x65,
	0x65, 0x64, 0x12, 0x19, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0xca, 0xc1, 0x18, 0x0d, 0x2f,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x42, 0x23, 0x5a, 0x21,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2f, 0x62, 0x69, 0x7a,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x66, 0x65, 0x65,
	0x64,
}

var (
	file_feed_proto_rawDescOnce sync.Once
	file_feed_proto_rawDescData = file_feed_proto_rawDesc
)

func file_feed_proto_rawDescGZIP() []byte {
	file_feed_proto_rawDescOnce.Do(func() {
		file_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_feed_proto_rawDescData)
	})
	return file_feed_proto_rawDescData
}

var file_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_feed_proto_goTypes = []interface{}{
	(*DouyinFeedRequest)(nil),  // 0: feed.douyin_feed_request
	(*DouyinFeedResponse)(nil), // 1: feed.douyin_feed_response
	(*Video)(nil),              // 2: feed.Video
	(*User)(nil),               // 3: feed.User
}
var file_feed_proto_depIdxs = []int32{
	2, // 0: feed.douyin_feed_response.video_list:type_name -> feed.Video
	3, // 1: feed.Video.author:type_name -> feed.User
	0, // 2: feed.feed_service.feed:input_type -> feed.douyin_feed_request
	1, // 3: feed.feed_service.feed:output_type -> feed.douyin_feed_response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_feed_proto_init() }
func file_feed_proto_init() {
	if File_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFeedRequest); i {
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
		file_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFeedResponse); i {
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
		file_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feed_proto_goTypes,
		DependencyIndexes: file_feed_proto_depIdxs,
		MessageInfos:      file_feed_proto_msgTypes,
	}.Build()
	File_feed_proto = out.File
	file_feed_proto_rawDesc = nil
	file_feed_proto_goTypes = nil
	file_feed_proto_depIdxs = nil
}
