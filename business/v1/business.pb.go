// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.0--rc2
// source: api/business/v1/business.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 回复评价的请求
type ReplyReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewID      int64                  `protobuf:"varint,1,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	StoreID       int64                  `protobuf:"varint,2,opt,name=storeID,proto3" json:"storeID,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	PicInfo       string                 `protobuf:"bytes,4,opt,name=picInfo,proto3" json:"picInfo,omitempty"`
	VideoInfo     string                 `protobuf:"bytes,5,opt,name=videoInfo,proto3" json:"videoInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplyReviewRequest) Reset() {
	*x = ReplyReviewRequest{}
	mi := &file_api_business_v1_business_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplyReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyReviewRequest) ProtoMessage() {}

func (x *ReplyReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_business_v1_business_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyReviewRequest.ProtoReflect.Descriptor instead.
func (*ReplyReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_business_v1_business_proto_rawDescGZIP(), []int{0}
}

func (x *ReplyReviewRequest) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *ReplyReviewRequest) GetStoreID() int64 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

func (x *ReplyReviewRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ReplyReviewRequest) GetPicInfo() string {
	if x != nil {
		return x.PicInfo
	}
	return ""
}

func (x *ReplyReviewRequest) GetVideoInfo() string {
	if x != nil {
		return x.VideoInfo
	}
	return ""
}

// 回复评价的返回值
type ReplyReviewReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReplyID       int64                  `protobuf:"varint,1,opt,name=replyID,proto3" json:"replyID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplyReviewReply) Reset() {
	*x = ReplyReviewReply{}
	mi := &file_api_business_v1_business_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplyReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyReviewReply) ProtoMessage() {}

func (x *ReplyReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_business_v1_business_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyReviewReply.ProtoReflect.Descriptor instead.
func (*ReplyReviewReply) Descriptor() ([]byte, []int) {
	return file_api_business_v1_business_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyReviewReply) GetReplyID() int64 {
	if x != nil {
		return x.ReplyID
	}
	return 0
}

var File_api_business_v1_business_proto protoreflect.FileDescriptor

var file_api_business_v1_business_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10,
	0x02, 0x18, 0xc8, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x49, 0x44, 0x32, 0x86, 0x01, 0x0a, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x12, 0x7a, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a,
	0x01, 0x2a, 0x22, 0x18, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x30, 0x0a, 0x0f,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x1b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_business_v1_business_proto_rawDescOnce sync.Once
	file_api_business_v1_business_proto_rawDescData = file_api_business_v1_business_proto_rawDesc
)

func file_api_business_v1_business_proto_rawDescGZIP() []byte {
	file_api_business_v1_business_proto_rawDescOnce.Do(func() {
		file_api_business_v1_business_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_business_v1_business_proto_rawDescData)
	})
	return file_api_business_v1_business_proto_rawDescData
}

var file_api_business_v1_business_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_business_v1_business_proto_goTypes = []any{
	(*ReplyReviewRequest)(nil), // 0: api.business.v1.ReplyReviewRequest
	(*ReplyReviewReply)(nil),   // 1: api.business.v1.ReplyReviewReply
}
var file_api_business_v1_business_proto_depIdxs = []int32{
	0, // 0: api.business.v1.Business.ReplyReview:input_type -> api.business.v1.ReplyReviewRequest
	1, // 1: api.business.v1.Business.ReplyReview:output_type -> api.business.v1.ReplyReviewReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_business_v1_business_proto_init() }
func file_api_business_v1_business_proto_init() {
	if File_api_business_v1_business_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_business_v1_business_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_business_v1_business_proto_goTypes,
		DependencyIndexes: file_api_business_v1_business_proto_depIdxs,
		MessageInfos:      file_api_business_v1_business_proto_msgTypes,
	}.Build()
	File_api_business_v1_business_proto = out.File
	file_api_business_v1_business_proto_rawDesc = nil
	file_api_business_v1_business_proto_goTypes = nil
	file_api_business_v1_business_proto_depIdxs = nil
}
