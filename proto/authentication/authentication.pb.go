// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: authentication.proto

package authentication

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	jwt "iceline-hosting.com/backend/proto/jwt"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VerifiedState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State bool `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *VerifiedState) Reset() {
	*x = VerifiedState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifiedState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifiedState) ProtoMessage() {}

func (x *VerifiedState) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifiedState.ProtoReflect.Descriptor instead.
func (*VerifiedState) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *VerifiedState) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

var File_authentication_proto protoreflect.FileDescriptor

var file_authentication_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0x74, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x47, 0x72, 0x70, 0x63, 0x12, 0x08, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x4a,
	0x57, 0x54, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x4a, 0x77, 0x74, 0x47, 0x72, 0x70, 0x63, 0x12, 0x08, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x4a,
	0x57, 0x54, 0x1a, 0x08, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x4a, 0x57, 0x54, 0x22, 0x00, 0x42, 0x32,
	0x5a, 0x30, 0x69, 0x63, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authentication_proto_rawDescOnce sync.Once
	file_authentication_proto_rawDescData = file_authentication_proto_rawDesc
)

func file_authentication_proto_rawDescGZIP() []byte {
	file_authentication_proto_rawDescOnce.Do(func() {
		file_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_authentication_proto_rawDescData)
	})
	return file_authentication_proto_rawDescData
}

var file_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_authentication_proto_goTypes = []any{
	(*VerifiedState)(nil), // 0: authentication.VerifiedState
	(*jwt.JWT)(nil),       // 1: jwt.JWT
}
var file_authentication_proto_depIdxs = []int32{
	1, // 0: authentication.Authentication.VerifyGrpc:input_type -> jwt.JWT
	1, // 1: authentication.Authentication.RegenerateJwtGrpc:input_type -> jwt.JWT
	0, // 2: authentication.Authentication.VerifyGrpc:output_type -> authentication.VerifiedState
	1, // 3: authentication.Authentication.RegenerateJwtGrpc:output_type -> jwt.JWT
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authentication_proto_init() }
func file_authentication_proto_init() {
	if File_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authentication_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*VerifiedState); i {
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
			RawDescriptor: file_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authentication_proto_goTypes,
		DependencyIndexes: file_authentication_proto_depIdxs,
		MessageInfos:      file_authentication_proto_msgTypes,
	}.Build()
	File_authentication_proto = out.File
	file_authentication_proto_rawDesc = nil
	file_authentication_proto_goTypes = nil
	file_authentication_proto_depIdxs = nil
}
