// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: sum.proto

package proto

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

type SummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	LastNumber  int32 `protobuf:"varint,2,opt,name=last_number,json=lastNumber,proto3" json:"last_number,omitempty"`
}

func (x *SummaryRequest) Reset() {
	*x = SummaryRequest{}
	mi := &file_sum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryRequest) ProtoMessage() {}

func (x *SummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryRequest.ProtoReflect.Descriptor instead.
func (*SummaryRequest) Descriptor() ([]byte, []int) {
	return file_sum_proto_rawDescGZIP(), []int{0}
}

func (x *SummaryRequest) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *SummaryRequest) GetLastNumber() int32 {
	if x != nil {
		return x.LastNumber
	}
	return 0
}

type SummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SummaryResponse) Reset() {
	*x = SummaryResponse{}
	mi := &file_sum_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummaryResponse) ProtoMessage() {}

func (x *SummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sum_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummaryResponse.ProtoReflect.Descriptor instead.
func (*SummaryResponse) Descriptor() ([]byte, []int) {
	return file_sum_proto_rawDescGZIP(), []int{1}
}

func (x *SummaryResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_sum_proto protoreflect.FileDescriptor

var file_sum_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x22, 0x54, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0f, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x4e, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x17, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x68, 0x64, 0x41, 0x6d, 0x69, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sum_proto_rawDescOnce sync.Once
	file_sum_proto_rawDescData = file_sum_proto_rawDesc
)

func file_sum_proto_rawDescGZIP() []byte {
	file_sum_proto_rawDescOnce.Do(func() {
		file_sum_proto_rawDescData = protoimpl.X.CompressGZIP(file_sum_proto_rawDescData)
	})
	return file_sum_proto_rawDescData
}

var file_sum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sum_proto_goTypes = []any{
	(*SummaryRequest)(nil),  // 0: summary.SummaryRequest
	(*SummaryResponse)(nil), // 1: summary.SummaryResponse
}
var file_sum_proto_depIdxs = []int32{
	0, // 0: summary.SummaryService.Summary:input_type -> summary.SummaryRequest
	1, // 1: summary.SummaryService.Summary:output_type -> summary.SummaryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sum_proto_init() }
func file_sum_proto_init() {
	if File_sum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sum_proto_goTypes,
		DependencyIndexes: file_sum_proto_depIdxs,
		MessageInfos:      file_sum_proto_msgTypes,
	}.Build()
	File_sum_proto = out.File
	file_sum_proto_rawDesc = nil
	file_sum_proto_goTypes = nil
	file_sum_proto_depIdxs = nil
}