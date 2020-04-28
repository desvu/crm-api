// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        (unknown)
// source: pkg/grpc/proto/game.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FindGamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameID string `protobuf:"bytes,1,opt,name=gameID,proto3" json:"gameID,omitempty"`
	Slug   string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Genre  string `protobuf:"bytes,3,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *FindGamesRequest) Reset() {
	*x = FindGamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGamesRequest) ProtoMessage() {}

func (x *FindGamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGamesRequest.ProtoReflect.Descriptor instead.
func (*FindGamesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{0}
}

func (x *FindGamesRequest) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

func (x *FindGamesRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *FindGamesRequest) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type FindGamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games []*Game `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
}

func (x *FindGamesResponse) Reset() {
	*x = FindGamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGamesResponse) ProtoMessage() {}

func (x *FindGamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGamesResponse.ProtoReflect.Descriptor instead.
func (*FindGamesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{1}
}

func (x *FindGamesResponse) GetGames() []*Game {
	if x != nil {
		return x.Games
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameID     string `protobuf:"bytes,1,opt,name=gameID,proto3" json:"gameID,omitempty"`
	RevisionID uint64 `protobuf:"varint,2,opt,name=revisionID,proto3" json:"revisionID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

func (x *Request) GetRevisionID() uint64 {
	if x != nil {
		return x.RevisionID
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string       `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Title       string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type        string       `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	RevisionID  uint64       `protobuf:"varint,4,opt,name=revisionID,proto3" json:"revisionID,omitempty"`
	Summary     string       `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	Description string       `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Slug        string       `protobuf:"bytes,7,opt,name=slug,proto3" json:"slug,omitempty"`
	License     string       `protobuf:"bytes,8,opt,name=license,proto3" json:"license,omitempty"`
	Platforms   []string     `protobuf:"bytes,9,rep,name=platforms,proto3" json:"platforms,omitempty"`
	Tags        []*Tag       `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	Features    []*Feature   `protobuf:"bytes,11,rep,name=features,proto3" json:"features,omitempty"`
	Developers  []*Developer `protobuf:"bytes,12,rep,name=developers,proto3" json:"developers,omitempty"`
	Genres      []*Genre     `protobuf:"bytes,13,rep,name=genres,proto3" json:"genres,omitempty"`
	Publishers  []*Publisher `protobuf:"bytes,14,rep,name=publishers,proto3" json:"publishers,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{4}
}

func (x *Game) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Game) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Game) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Game) GetRevisionID() uint64 {
	if x != nil {
		return x.RevisionID
	}
	return 0
}

func (x *Game) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Game) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Game) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Game) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *Game) GetPlatforms() []string {
	if x != nil {
		return x.Platforms
	}
	return nil
}

func (x *Game) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Game) GetFeatures() []*Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *Game) GetDevelopers() []*Developer {
	if x != nil {
		return x.Developers
	}
	return nil
}

func (x *Game) GetGenres() []*Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Game) GetPublishers() []*Publisher {
	if x != nil {
		return x.Publishers
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{5}
}

func (x *Tag) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Developer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Developer) Reset() {
	*x = Developer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Developer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Developer) ProtoMessage() {}

func (x *Developer) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Developer.ProtoReflect.Descriptor instead.
func (*Developer) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{6}
}

func (x *Developer) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Developer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{7}
}

func (x *Feature) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Genre) Reset() {
	*x = Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{8}
}

func (x *Genre) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Genre) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Publisher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Publisher) Reset() {
	*x = Publisher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_game_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Publisher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Publisher) ProtoMessage() {}

func (x *Publisher) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_game_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Publisher.ProtoReflect.Descriptor instead.
func (*Publisher) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_game_proto_rawDescGZIP(), []int{9}
}

func (x *Publisher) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Publisher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_pkg_grpc_proto_game_proto protoreflect.FileDescriptor

var file_pkg_grpc_proto_game_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x54, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x22, 0x36, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73,
	0x22, 0x41, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x61, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d,
	0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x22, 0x2b, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65,
	0x22, 0xbe, 0x03, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x73, 0x22, 0x29, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x09,
	0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a,
	0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x05,
	0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x09, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x8a, 0x01, 0x0a, 0x0b, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_grpc_proto_game_proto_rawDescOnce sync.Once
	file_pkg_grpc_proto_game_proto_rawDescData = file_pkg_grpc_proto_game_proto_rawDesc
)

func file_pkg_grpc_proto_game_proto_rawDescGZIP() []byte {
	file_pkg_grpc_proto_game_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_proto_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_proto_game_proto_rawDescData)
	})
	return file_pkg_grpc_proto_game_proto_rawDescData
}

var file_pkg_grpc_proto_game_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_grpc_proto_game_proto_goTypes = []interface{}{
	(*FindGamesRequest)(nil),  // 0: proto.FindGamesRequest
	(*FindGamesResponse)(nil), // 1: proto.FindGamesResponse
	(*Request)(nil),           // 2: proto.Request
	(*Response)(nil),          // 3: proto.Response
	(*Game)(nil),              // 4: proto.Game
	(*Tag)(nil),               // 5: proto.Tag
	(*Developer)(nil),         // 6: proto.Developer
	(*Feature)(nil),           // 7: proto.Feature
	(*Genre)(nil),             // 8: proto.Genre
	(*Publisher)(nil),         // 9: proto.Publisher
}
var file_pkg_grpc_proto_game_proto_depIdxs = []int32{
	4, // 0: proto.FindGamesResponse.games:type_name -> proto.Game
	4, // 1: proto.Response.game:type_name -> proto.Game
	5, // 2: proto.Game.tags:type_name -> proto.Tag
	7, // 3: proto.Game.features:type_name -> proto.Feature
	6, // 4: proto.Game.developers:type_name -> proto.Developer
	8, // 5: proto.Game.genres:type_name -> proto.Genre
	9, // 6: proto.Game.publishers:type_name -> proto.Publisher
	2, // 7: proto.GameService.GetByIDAndRevisionID:input_type -> proto.Request
	0, // 8: proto.GameService.FindGames:input_type -> proto.FindGamesRequest
	3, // 9: proto.GameService.GetByIDAndRevisionID:output_type -> proto.Response
	1, // 10: proto.GameService.FindGames:output_type -> proto.FindGamesResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_grpc_proto_game_proto_init() }
func file_pkg_grpc_proto_game_proto_init() {
	if File_pkg_grpc_proto_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_proto_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGamesRequest); i {
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
		file_pkg_grpc_proto_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGamesResponse); i {
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
		file_pkg_grpc_proto_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_pkg_grpc_proto_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_pkg_grpc_proto_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_pkg_grpc_proto_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_pkg_grpc_proto_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Developer); i {
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
		file_pkg_grpc_proto_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_pkg_grpc_proto_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre); i {
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
		file_pkg_grpc_proto_game_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Publisher); i {
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
			RawDescriptor: file_pkg_grpc_proto_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_proto_game_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_proto_game_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_proto_game_proto_msgTypes,
	}.Build()
	File_pkg_grpc_proto_game_proto = out.File
	file_pkg_grpc_proto_game_proto_rawDesc = nil
	file_pkg_grpc_proto_game_proto_goTypes = nil
	file_pkg_grpc_proto_game_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceClient interface {
	GetByIDAndRevisionID(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// FindGames returns list of latest published games filtered by request parameters
	FindGames(ctx context.Context, in *FindGamesRequest, opts ...grpc.CallOption) (*FindGamesResponse, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) GetByIDAndRevisionID(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.GameService/GetByIDAndRevisionID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) FindGames(ctx context.Context, in *FindGamesRequest, opts ...grpc.CallOption) (*FindGamesResponse, error) {
	out := new(FindGamesResponse)
	err := c.cc.Invoke(ctx, "/proto.GameService/FindGames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
type GameServiceServer interface {
	GetByIDAndRevisionID(context.Context, *Request) (*Response, error)
	// FindGames returns list of latest published games filtered by request parameters
	FindGames(context.Context, *FindGamesRequest) (*FindGamesResponse, error)
}

// UnimplementedGameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (*UnimplementedGameServiceServer) GetByIDAndRevisionID(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIDAndRevisionID not implemented")
}
func (*UnimplementedGameServiceServer) FindGames(context.Context, *FindGamesRequest) (*FindGamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGames not implemented")
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_GetByIDAndRevisionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetByIDAndRevisionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameService/GetByIDAndRevisionID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetByIDAndRevisionID(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_FindGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindGamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).FindGames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameService/FindGames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).FindGames(ctx, req.(*FindGamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByIDAndRevisionID",
			Handler:    _GameService_GetByIDAndRevisionID_Handler,
		},
		{
			MethodName: "FindGames",
			Handler:    _GameService_FindGames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/proto/game.proto",
}
