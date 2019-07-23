// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/postpb/post.proto

package postpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 投稿の定義
type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Usecase              string   `protobuf:"bytes,6,opt,name=usecase,proto3" json:"usecase,omitempty"`
	Genre                string   `protobuf:"bytes,7,opt,name=genre,proto3" json:"genre,omitempty"`
	UserId               string   `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Post) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Post) GetUsecase() string {
	if m != nil {
		return m.Usecase
	}
	return ""
}

func (m *Post) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func (m *Post) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// 投稿の作成
type CreatePostRequest struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{1}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type CreatePostResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostResponse) Reset()         { *m = CreatePostResponse{} }
func (m *CreatePostResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePostResponse) ProtoMessage()    {}
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{2}
}

func (m *CreatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostResponse.Unmarshal(m, b)
}
func (m *CreatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostResponse.Marshal(b, m, deterministic)
}
func (m *CreatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostResponse.Merge(m, src)
}
func (m *CreatePostResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePostResponse.Size(m)
}
func (m *CreatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostResponse proto.InternalMessageInfo

func (m *CreatePostResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 投稿の削除
type DeletePostRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostRequest) Reset()         { *m = DeletePostRequest{} }
func (m *DeletePostRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePostRequest) ProtoMessage()    {}
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{3}
}

func (m *DeletePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostRequest.Unmarshal(m, b)
}
func (m *DeletePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostRequest.Marshal(b, m, deterministic)
}
func (m *DeletePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostRequest.Merge(m, src)
}
func (m *DeletePostRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePostRequest.Size(m)
}
func (m *DeletePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostRequest proto.InternalMessageInfo

func (m *DeletePostRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeletePostResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostResponse) Reset()         { *m = DeletePostResponse{} }
func (m *DeletePostResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePostResponse) ProtoMessage()    {}
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{4}
}

func (m *DeletePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostResponse.Unmarshal(m, b)
}
func (m *DeletePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostResponse.Marshal(b, m, deterministic)
}
func (m *DeletePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostResponse.Merge(m, src)
}
func (m *DeletePostResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePostResponse.Size(m)
}
func (m *DeletePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostResponse proto.InternalMessageInfo

func (m *DeletePostResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// ユーザーIDを紐付いたすべての投稿を取得
type GetAllPostsByUserIDRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllPostsByUserIDRequest) Reset()         { *m = GetAllPostsByUserIDRequest{} }
func (m *GetAllPostsByUserIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllPostsByUserIDRequest) ProtoMessage()    {}
func (*GetAllPostsByUserIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{5}
}

func (m *GetAllPostsByUserIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPostsByUserIDRequest.Unmarshal(m, b)
}
func (m *GetAllPostsByUserIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPostsByUserIDRequest.Marshal(b, m, deterministic)
}
func (m *GetAllPostsByUserIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPostsByUserIDRequest.Merge(m, src)
}
func (m *GetAllPostsByUserIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllPostsByUserIDRequest.Size(m)
}
func (m *GetAllPostsByUserIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPostsByUserIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPostsByUserIDRequest proto.InternalMessageInfo

func (m *GetAllPostsByUserIDRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetAllPostsByUserIDResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllPostsByUserIDResponse) Reset()         { *m = GetAllPostsByUserIDResponse{} }
func (m *GetAllPostsByUserIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllPostsByUserIDResponse) ProtoMessage()    {}
func (*GetAllPostsByUserIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{6}
}

func (m *GetAllPostsByUserIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPostsByUserIDResponse.Unmarshal(m, b)
}
func (m *GetAllPostsByUserIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPostsByUserIDResponse.Marshal(b, m, deterministic)
}
func (m *GetAllPostsByUserIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPostsByUserIDResponse.Merge(m, src)
}
func (m *GetAllPostsByUserIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllPostsByUserIDResponse.Size(m)
}
func (m *GetAllPostsByUserIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPostsByUserIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPostsByUserIDResponse proto.InternalMessageInfo

func (m *GetAllPostsByUserIDResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

// タイトル検索
type SearchPostsByTitleRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPostsByTitleRequest) Reset()         { *m = SearchPostsByTitleRequest{} }
func (m *SearchPostsByTitleRequest) String() string { return proto.CompactTextString(m) }
func (*SearchPostsByTitleRequest) ProtoMessage()    {}
func (*SearchPostsByTitleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{7}
}

func (m *SearchPostsByTitleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPostsByTitleRequest.Unmarshal(m, b)
}
func (m *SearchPostsByTitleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPostsByTitleRequest.Marshal(b, m, deterministic)
}
func (m *SearchPostsByTitleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPostsByTitleRequest.Merge(m, src)
}
func (m *SearchPostsByTitleRequest) XXX_Size() int {
	return xxx_messageInfo_SearchPostsByTitleRequest.Size(m)
}
func (m *SearchPostsByTitleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPostsByTitleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPostsByTitleRequest proto.InternalMessageInfo

func (m *SearchPostsByTitleRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SearchPostsByTitleRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type SearchPostsByTitleResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPostsByTitleResponse) Reset()         { *m = SearchPostsByTitleResponse{} }
func (m *SearchPostsByTitleResponse) String() string { return proto.CompactTextString(m) }
func (*SearchPostsByTitleResponse) ProtoMessage()    {}
func (*SearchPostsByTitleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{8}
}

func (m *SearchPostsByTitleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPostsByTitleResponse.Unmarshal(m, b)
}
func (m *SearchPostsByTitleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPostsByTitleResponse.Marshal(b, m, deterministic)
}
func (m *SearchPostsByTitleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPostsByTitleResponse.Merge(m, src)
}
func (m *SearchPostsByTitleResponse) XXX_Size() int {
	return xxx_messageInfo_SearchPostsByTitleResponse.Size(m)
}
func (m *SearchPostsByTitleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPostsByTitleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPostsByTitleResponse proto.InternalMessageInfo

func (m *SearchPostsByTitleResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

// ユースケース検索
type SearchPostsByUsecaseRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Usecase              string   `protobuf:"bytes,2,opt,name=usecase,proto3" json:"usecase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPostsByUsecaseRequest) Reset()         { *m = SearchPostsByUsecaseRequest{} }
func (m *SearchPostsByUsecaseRequest) String() string { return proto.CompactTextString(m) }
func (*SearchPostsByUsecaseRequest) ProtoMessage()    {}
func (*SearchPostsByUsecaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{9}
}

func (m *SearchPostsByUsecaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPostsByUsecaseRequest.Unmarshal(m, b)
}
func (m *SearchPostsByUsecaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPostsByUsecaseRequest.Marshal(b, m, deterministic)
}
func (m *SearchPostsByUsecaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPostsByUsecaseRequest.Merge(m, src)
}
func (m *SearchPostsByUsecaseRequest) XXX_Size() int {
	return xxx_messageInfo_SearchPostsByUsecaseRequest.Size(m)
}
func (m *SearchPostsByUsecaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPostsByUsecaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPostsByUsecaseRequest proto.InternalMessageInfo

func (m *SearchPostsByUsecaseRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SearchPostsByUsecaseRequest) GetUsecase() string {
	if m != nil {
		return m.Usecase
	}
	return ""
}

type SearchPostsByUsecaseResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPostsByUsecaseResponse) Reset()         { *m = SearchPostsByUsecaseResponse{} }
func (m *SearchPostsByUsecaseResponse) String() string { return proto.CompactTextString(m) }
func (*SearchPostsByUsecaseResponse) ProtoMessage()    {}
func (*SearchPostsByUsecaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{10}
}

func (m *SearchPostsByUsecaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPostsByUsecaseResponse.Unmarshal(m, b)
}
func (m *SearchPostsByUsecaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPostsByUsecaseResponse.Marshal(b, m, deterministic)
}
func (m *SearchPostsByUsecaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPostsByUsecaseResponse.Merge(m, src)
}
func (m *SearchPostsByUsecaseResponse) XXX_Size() int {
	return xxx_messageInfo_SearchPostsByUsecaseResponse.Size(m)
}
func (m *SearchPostsByUsecaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPostsByUsecaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPostsByUsecaseResponse proto.InternalMessageInfo

func (m *SearchPostsByUsecaseResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

// ジャンル検索
type SearchPostsByGenreRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Genre                string   `protobuf:"bytes,2,opt,name=genre,proto3" json:"genre,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPostsByGenreRequest) Reset()         { *m = SearchPostsByGenreRequest{} }
func (m *SearchPostsByGenreRequest) String() string { return proto.CompactTextString(m) }
func (*SearchPostsByGenreRequest) ProtoMessage()    {}
func (*SearchPostsByGenreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{11}
}

func (m *SearchPostsByGenreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPostsByGenreRequest.Unmarshal(m, b)
}
func (m *SearchPostsByGenreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPostsByGenreRequest.Marshal(b, m, deterministic)
}
func (m *SearchPostsByGenreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPostsByGenreRequest.Merge(m, src)
}
func (m *SearchPostsByGenreRequest) XXX_Size() int {
	return xxx_messageInfo_SearchPostsByGenreRequest.Size(m)
}
func (m *SearchPostsByGenreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPostsByGenreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPostsByGenreRequest proto.InternalMessageInfo

func (m *SearchPostsByGenreRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SearchPostsByGenreRequest) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

type SearchPostsByGenreResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchPostsByGenreResponse) Reset()         { *m = SearchPostsByGenreResponse{} }
func (m *SearchPostsByGenreResponse) String() string { return proto.CompactTextString(m) }
func (*SearchPostsByGenreResponse) ProtoMessage()    {}
func (*SearchPostsByGenreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{12}
}

func (m *SearchPostsByGenreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchPostsByGenreResponse.Unmarshal(m, b)
}
func (m *SearchPostsByGenreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchPostsByGenreResponse.Marshal(b, m, deterministic)
}
func (m *SearchPostsByGenreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchPostsByGenreResponse.Merge(m, src)
}
func (m *SearchPostsByGenreResponse) XXX_Size() int {
	return xxx_messageInfo_SearchPostsByGenreResponse.Size(m)
}
func (m *SearchPostsByGenreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchPostsByGenreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchPostsByGenreResponse proto.InternalMessageInfo

func (m *SearchPostsByGenreResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

// URLからスクレイピングしてくれる
type PostURLRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostURLRequest) Reset()         { *m = PostURLRequest{} }
func (m *PostURLRequest) String() string { return proto.CompactTextString(m) }
func (*PostURLRequest) ProtoMessage()    {}
func (*PostURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{13}
}

func (m *PostURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostURLRequest.Unmarshal(m, b)
}
func (m *PostURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostURLRequest.Marshal(b, m, deterministic)
}
func (m *PostURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostURLRequest.Merge(m, src)
}
func (m *PostURLRequest) XXX_Size() int {
	return xxx_messageInfo_PostURLRequest.Size(m)
}
func (m *PostURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostURLRequest proto.InternalMessageInfo

func (m *PostURLRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type PostResponse struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image                string   `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResponse) Reset()         { *m = PostResponse{} }
func (m *PostResponse) String() string { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()    {}
func (*PostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a98b095e97a36b, []int{14}
}

func (m *PostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResponse.Unmarshal(m, b)
}
func (m *PostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResponse.Marshal(b, m, deterministic)
}
func (m *PostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResponse.Merge(m, src)
}
func (m *PostResponse) XXX_Size() int {
	return xxx_messageInfo_PostResponse.Size(m)
}
func (m *PostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostResponse proto.InternalMessageInfo

func (m *PostResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PostResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PostResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PostResponse) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func init() {
	proto.RegisterType((*Post)(nil), "postpb.Post")
	proto.RegisterType((*CreatePostRequest)(nil), "postpb.CreatePostRequest")
	proto.RegisterType((*CreatePostResponse)(nil), "postpb.CreatePostResponse")
	proto.RegisterType((*DeletePostRequest)(nil), "postpb.DeletePostRequest")
	proto.RegisterType((*DeletePostResponse)(nil), "postpb.DeletePostResponse")
	proto.RegisterType((*GetAllPostsByUserIDRequest)(nil), "postpb.GetAllPostsByUserIDRequest")
	proto.RegisterType((*GetAllPostsByUserIDResponse)(nil), "postpb.GetAllPostsByUserIDResponse")
	proto.RegisterType((*SearchPostsByTitleRequest)(nil), "postpb.SearchPostsByTitleRequest")
	proto.RegisterType((*SearchPostsByTitleResponse)(nil), "postpb.SearchPostsByTitleResponse")
	proto.RegisterType((*SearchPostsByUsecaseRequest)(nil), "postpb.SearchPostsByUsecaseRequest")
	proto.RegisterType((*SearchPostsByUsecaseResponse)(nil), "postpb.SearchPostsByUsecaseResponse")
	proto.RegisterType((*SearchPostsByGenreRequest)(nil), "postpb.SearchPostsByGenreRequest")
	proto.RegisterType((*SearchPostsByGenreResponse)(nil), "postpb.SearchPostsByGenreResponse")
	proto.RegisterType((*PostURLRequest)(nil), "postpb.PostURLRequest")
	proto.RegisterType((*PostResponse)(nil), "postpb.PostResponse")
}

func init() { proto.RegisterFile("proto/postpb/post.proto", fileDescriptor_12a98b095e97a36b) }

var fileDescriptor_12a98b095e97a36b = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x9d, 0x8f, 0xc2, 0x24, 0x44, 0x74, 0x88, 0xe8, 0xd6, 0xe5, 0x10, 0x36, 0x1c, 0x38,
	0x05, 0xa9, 0xa8, 0x67, 0x68, 0x08, 0x8a, 0x40, 0x1c, 0xaa, 0x94, 0x5c, 0x90, 0x10, 0xb8, 0xf6,
	0xa8, 0x58, 0x72, 0x63, 0xe3, 0xdd, 0x20, 0xf1, 0xeb, 0xf8, 0x2d, 0xfc, 0x13, 0xb4, 0x6b, 0x3b,
	0x5e, 0xd7, 0x76, 0x42, 0xd5, 0x53, 0x3b, 0xb3, 0x6f, 0xde, 0xce, 0x5b, 0xbf, 0xa7, 0xc0, 0x51,
	0x9c, 0x44, 0x32, 0x7a, 0x15, 0x47, 0x42, 0xc6, 0x57, 0xfa, 0xcf, 0x54, 0x77, 0xb0, 0x97, 0xb6,
	0xf8, 0x1f, 0x0b, 0x3a, 0x17, 0x91, 0x90, 0x38, 0x04, 0x3b, 0xf0, 0x99, 0x35, 0xb6, 0x5e, 0x3e,
	0x5c, 0xda, 0x81, 0x8f, 0x8f, 0xa1, 0xbd, 0x49, 0x42, 0x66, 0xeb, 0x86, 0xfa, 0x17, 0x47, 0xd0,
	0x95, 0x81, 0x0c, 0x89, 0xb5, 0x75, 0x2f, 0x2d, 0x70, 0x0c, 0x7d, 0x9f, 0x84, 0x97, 0x04, 0xb1,
	0x0c, 0xa2, 0x35, 0xeb, 0xe8, 0x33, 0xb3, 0xa5, 0xe6, 0x82, 0x1b, 0xf7, 0x9a, 0x58, 0x37, 0x9d,
	0xd3, 0x05, 0x32, 0x38, 0xd8, 0x08, 0xf2, 0x5c, 0x41, 0xac, 0xa7, 0xfb, 0x79, 0xa9, 0xf0, 0xd7,
	0xb4, 0x4e, 0x88, 0x1d, 0xa4, 0x78, 0x5d, 0xe0, 0x91, 0xc6, 0x27, 0xdf, 0x02, 0x9f, 0x3d, 0xd0,
	0xfd, 0x9e, 0x2a, 0x3f, 0xf8, 0xfc, 0x0c, 0x0e, 0xdf, 0x25, 0xe4, 0x4a, 0x52, 0x32, 0x96, 0xf4,
	0x73, 0x43, 0x42, 0xe2, 0x18, 0x3a, 0x4a, 0xa0, 0xd6, 0xd3, 0x3f, 0x1d, 0x4c, 0x53, 0xb5, 0x53,
	0x0d, 0xd1, 0x27, 0x7c, 0x0a, 0x68, 0x8e, 0x89, 0x38, 0x5a, 0x0b, 0xbd, 0xd5, 0x0d, 0x09, 0xa1,
	0xb6, 0x4d, 0x9f, 0x22, 0x2f, 0xf9, 0x04, 0x0e, 0xe7, 0x14, 0x52, 0xf9, 0x9a, 0x5b, 0x8f, 0xa6,
	0x48, 0x4d, 0xd0, 0x5e, 0xd2, 0x33, 0x70, 0x16, 0x24, 0xcf, 0xc3, 0x50, 0xe1, 0xc5, 0xec, 0xf7,
	0x4a, 0x49, 0x9a, 0xe7, 0xec, 0x86, 0x64, 0xab, 0x24, 0xf9, 0x1c, 0x4e, 0x6a, 0xc7, 0xb2, 0xfb,
	0x38, 0x74, 0x95, 0x44, 0xc1, 0xac, 0x71, 0xbb, 0xa2, 0x3e, 0x3d, 0xe2, 0x1f, 0xe1, 0xf8, 0x92,
	0xdc, 0xc4, 0xfb, 0x91, 0x51, 0x7c, 0x56, 0x1f, 0x73, 0xdf, 0xc5, 0x85, 0x05, 0x6c, 0xc3, 0x02,
	0xfc, 0x2d, 0x38, 0x75, 0x5c, 0x77, 0xd8, 0xe6, 0x02, 0x4e, 0x4a, 0x0c, 0xab, 0xd4, 0x0a, 0x7b,
	0xf7, 0x31, 0x4c, 0x64, 0x97, 0x4c, 0xc4, 0x67, 0xf0, 0xac, 0x9e, 0xf1, 0x1e, 0x6f, 0xb4, 0x50,
	0x46, 0xfc, 0x9f, 0x37, 0x4a, 0xed, 0x6b, 0x1b, 0xf6, 0xad, 0xbc, 0x51, 0xc6, 0x75, 0x87, 0x6d,
	0x38, 0x0c, 0x55, 0xb9, 0x5a, 0x7e, 0xca, 0x57, 0xc8, 0x22, 0x6a, 0x6d, 0x23, 0xca, 0xd7, 0x30,
	0x28, 0x39, 0xaf, 0x82, 0xa8, 0xff, 0x82, 0xb7, 0x43, 0xdc, 0xde, 0x11, 0xe2, 0x8e, 0x11, 0xe2,
	0xd3, 0xbf, 0x1d, 0xe8, 0xab, 0x0b, 0x2f, 0x29, 0xf9, 0x15, 0x78, 0x84, 0x6f, 0xe0, 0xd1, 0x82,
	0xa4, 0xea, 0xcc, 0x49, 0xba, 0x41, 0x88, 0x4f, 0x4d, 0x25, 0xc5, 0xea, 0xce, 0xa8, 0xa4, 0x30,
	0x5b, 0x97, 0xb7, 0xf0, 0x3d, 0x40, 0x91, 0x4a, 0x3c, 0xce, 0x51, 0x95, 0x80, 0x3b, 0x4e, 0xdd,
	0x91, 0x49, 0x53, 0xe4, 0xb0, 0xa0, 0xa9, 0x04, 0xb8, 0xa0, 0xa9, 0xc6, 0x96, 0xb7, 0xf0, 0x3b,
	0x3c, 0xa9, 0xc9, 0x19, 0xf2, 0x7c, 0xa8, 0x39, 0xbb, 0xce, 0x64, 0x27, 0x66, 0x7b, 0xc3, 0x57,
	0xc0, 0x6a, 0x74, 0xf0, 0x79, 0x3e, 0xdc, 0x18, 0x51, 0x87, 0xef, 0x82, 0x6c, 0xe9, 0x3d, 0x18,
	0xd5, 0xa5, 0x00, 0x27, 0xb5, 0xd3, 0xe5, 0xd4, 0x39, 0x2f, 0x76, 0x83, 0x1a, 0x35, 0x68, 0x6b,
	0x37, 0x68, 0x30, 0x23, 0xd4, 0xa0, 0xa1, 0x94, 0x0c, 0xde, 0x9a, 0x0d, 0xbf, 0x0c, 0xcc, 0x1f,
	0xb1, 0xab, 0x9e, 0xae, 0x5e, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x9e, 0x8e, 0x26, 0xdb,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostServiceClient interface {
	// URLからスクレイピングしてくれる
	GetPostDetail(ctx context.Context, in *PostURLRequest, opts ...grpc.CallOption) (*PostResponse, error)
	// 投稿の作成
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	// 投稿の削除
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	// ユーザーIDを紐付いたすべての投稿を取得
	GetAllPostsByUserID(ctx context.Context, in *GetAllPostsByUserIDRequest, opts ...grpc.CallOption) (*GetAllPostsByUserIDResponse, error)
	// タイトル検索
	SearchPostsByTitle(ctx context.Context, in *SearchPostsByTitleRequest, opts ...grpc.CallOption) (*SearchPostsByTitleResponse, error)
	// ユースケース検索
	SearchPostsByUsecase(ctx context.Context, in *SearchPostsByUsecaseRequest, opts ...grpc.CallOption) (*SearchPostsByUsecaseResponse, error)
	// ジャンル検索
	SearchPostsByGenre(ctx context.Context, in *SearchPostsByGenreRequest, opts ...grpc.CallOption) (*SearchPostsByGenreResponse, error)
}

type postServiceClient struct {
	cc *grpc.ClientConn
}

func NewPostServiceClient(cc *grpc.ClientConn) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) GetPostDetail(ctx context.Context, in *PostURLRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/GetPostDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllPostsByUserID(ctx context.Context, in *GetAllPostsByUserIDRequest, opts ...grpc.CallOption) (*GetAllPostsByUserIDResponse, error) {
	out := new(GetAllPostsByUserIDResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/GetAllPostsByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) SearchPostsByTitle(ctx context.Context, in *SearchPostsByTitleRequest, opts ...grpc.CallOption) (*SearchPostsByTitleResponse, error) {
	out := new(SearchPostsByTitleResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/SearchPostsByTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) SearchPostsByUsecase(ctx context.Context, in *SearchPostsByUsecaseRequest, opts ...grpc.CallOption) (*SearchPostsByUsecaseResponse, error) {
	out := new(SearchPostsByUsecaseResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/SearchPostsByUsecase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) SearchPostsByGenre(ctx context.Context, in *SearchPostsByGenreRequest, opts ...grpc.CallOption) (*SearchPostsByGenreResponse, error) {
	out := new(SearchPostsByGenreResponse)
	err := c.cc.Invoke(ctx, "/postpb.PostService/SearchPostsByGenre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
type PostServiceServer interface {
	// URLからスクレイピングしてくれる
	GetPostDetail(context.Context, *PostURLRequest) (*PostResponse, error)
	// 投稿の作成
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	// 投稿の削除
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	// ユーザーIDを紐付いたすべての投稿を取得
	GetAllPostsByUserID(context.Context, *GetAllPostsByUserIDRequest) (*GetAllPostsByUserIDResponse, error)
	// タイトル検索
	SearchPostsByTitle(context.Context, *SearchPostsByTitleRequest) (*SearchPostsByTitleResponse, error)
	// ユースケース検索
	SearchPostsByUsecase(context.Context, *SearchPostsByUsecaseRequest) (*SearchPostsByUsecaseResponse, error)
	// ジャンル検索
	SearchPostsByGenre(context.Context, *SearchPostsByGenreRequest) (*SearchPostsByGenreResponse, error)
}

func RegisterPostServiceServer(s *grpc.Server, srv PostServiceServer) {
	s.RegisterService(&_PostService_serviceDesc, srv)
}

func _PostService_GetPostDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/GetPostDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostDetail(ctx, req.(*PostURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllPostsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPostsByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllPostsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/GetAllPostsByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllPostsByUserID(ctx, req.(*GetAllPostsByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_SearchPostsByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostsByTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).SearchPostsByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/SearchPostsByTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).SearchPostsByTitle(ctx, req.(*SearchPostsByTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_SearchPostsByUsecase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostsByUsecaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).SearchPostsByUsecase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/SearchPostsByUsecase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).SearchPostsByUsecase(ctx, req.(*SearchPostsByUsecaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_SearchPostsByGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostsByGenreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).SearchPostsByGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/postpb.PostService/SearchPostsByGenre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).SearchPostsByGenre(ctx, req.(*SearchPostsByGenreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "postpb.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPostDetail",
			Handler:    _PostService_GetPostDetail_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "GetAllPostsByUserID",
			Handler:    _PostService_GetAllPostsByUserID_Handler,
		},
		{
			MethodName: "SearchPostsByTitle",
			Handler:    _PostService_SearchPostsByTitle_Handler,
		},
		{
			MethodName: "SearchPostsByUsecase",
			Handler:    _PostService_SearchPostsByUsecase_Handler,
		},
		{
			MethodName: "SearchPostsByGenre",
			Handler:    _PostService_SearchPostsByGenre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/postpb/post.proto",
}