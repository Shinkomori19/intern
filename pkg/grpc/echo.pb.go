// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: echo.proto

package helloworld

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

// The request message from client.
type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqMessage string `protobuf:"bytes,1,opt,name=reqMessage,proto3" json:"reqMessage,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetReqMessage() string {
	if x != nil {
		return x.ReqMessage
	}
	return ""
}

// The response message from server.
type EchoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResMessage string `protobuf:"bytes,1,opt,name=resMessage,proto3" json:"resMessage,omitempty"`
}

func (x *EchoReply) Reset() {
	*x = EchoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoReply) ProtoMessage() {}

func (x *EchoReply) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoReply.ProtoReflect.Descriptor instead.
func (*EchoReply) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoReply) GetResMessage() string {
	if x != nil {
		return x.ResMessage
	}
	return ""
}

var File_echo_proto protoreflect.FileDescriptor

var file_echo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x2d, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x09, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x3f, 0x0a, 0x0b,
	0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x67, 0x0a,
	0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x0f, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echo_proto_rawDescOnce sync.Once
	file_echo_proto_rawDescData = file_echo_proto_rawDesc
)

func file_echo_proto_rawDescGZIP() []byte {
	file_echo_proto_rawDescOnce.Do(func() {
		file_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_proto_rawDescData)
	})
	return file_echo_proto_rawDescData
}

var file_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_echo_proto_goTypes = []interface{}{
	(*EchoRequest)(nil), // 0: helloworld.EchoRequest
	(*EchoReply)(nil),   // 1: helloworld.EchoReply
}
var file_echo_proto_depIdxs = []int32{
	0, // 0: helloworld.Echo.EchoMessage:input_type -> helloworld.EchoRequest
	1, // 1: helloworld.Echo.EchoMessage:output_type -> helloworld.EchoReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_echo_proto_init() }
func file_echo_proto_init() {
	if File_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoReply); i {
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
			RawDescriptor: file_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_proto_goTypes,
		DependencyIndexes: file_echo_proto_depIdxs,
		MessageInfos:      file_echo_proto_msgTypes,
	}.Build()
	File_echo_proto = out.File
	file_echo_proto_rawDesc = nil
	file_echo_proto_goTypes = nil
	file_echo_proto_depIdxs = nil
}