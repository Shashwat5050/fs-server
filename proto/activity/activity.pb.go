// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: activity.proto

package activity

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type RegisterUserActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string               `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ActivityType string               `protobuf:"bytes,2,opt,name=activity_type,json=activityType,proto3" json:"activity_type,omitempty"`
	LogTime      *timestamp.Timestamp `protobuf:"bytes,3,opt,name=log_time,json=logTime,proto3" json:"log_time,omitempty"`
}

func (x *RegisterUserActivityRequest) Reset() {
	*x = RegisterUserActivityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserActivityRequest) ProtoMessage() {}

func (x *RegisterUserActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserActivityRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserActivityRequest) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterUserActivityRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterUserActivityRequest) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

func (x *RegisterUserActivityRequest) GetLogTime() *timestamp.Timestamp {
	if x != nil {
		return x.LogTime
	}
	return nil
}

type RegisterGsActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string               `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ServerName   string               `protobuf:"bytes,2,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	ActivityType string               `protobuf:"bytes,3,opt,name=activity_type,json=activityType,proto3" json:"activity_type,omitempty"`
	LogTime      *timestamp.Timestamp `protobuf:"bytes,4,opt,name=log_time,json=logTime,proto3" json:"log_time,omitempty"`
}

func (x *RegisterGsActivityRequest) Reset() {
	*x = RegisterGsActivityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterGsActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterGsActivityRequest) ProtoMessage() {}

func (x *RegisterGsActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterGsActivityRequest.ProtoReflect.Descriptor instead.
func (*RegisterGsActivityRequest) Descriptor() ([]byte, []int) {
	return file_activity_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterGsActivityRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterGsActivityRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *RegisterGsActivityRequest) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

func (x *RegisterGsActivityRequest) GetLogTime() *timestamp.Timestamp {
	if x != nil {
		return x.LogTime
	}
	return nil
}

var File_activity_proto protoreflect.FileDescriptor

var file_activity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x1b, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb1, 0x01,
	0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f,
	0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x32, 0xbf, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x25, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x53,
	0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x12, 0x23, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x69, 0x63, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x68,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_proto_rawDescOnce sync.Once
	file_activity_proto_rawDescData = file_activity_proto_rawDesc
)

func file_activity_proto_rawDescGZIP() []byte {
	file_activity_proto_rawDescOnce.Do(func() {
		file_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_proto_rawDescData)
	})
	return file_activity_proto_rawDescData
}

var file_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_activity_proto_goTypes = []any{
	(*RegisterUserActivityRequest)(nil), // 0: activity.RegisterUserActivityRequest
	(*RegisterGsActivityRequest)(nil),   // 1: activity.RegisterGsActivityRequest
	(*timestamp.Timestamp)(nil),         // 2: google.protobuf.Timestamp
	(*empty.Empty)(nil),                 // 3: google.protobuf.Empty
}
var file_activity_proto_depIdxs = []int32{
	2, // 0: activity.RegisterUserActivityRequest.log_time:type_name -> google.protobuf.Timestamp
	2, // 1: activity.RegisterGsActivityRequest.log_time:type_name -> google.protobuf.Timestamp
	0, // 2: activity.ActivityManager.RegisterUserActivity:input_type -> activity.RegisterUserActivityRequest
	1, // 3: activity.ActivityManager.RegisterGsActivity:input_type -> activity.RegisterGsActivityRequest
	3, // 4: activity.ActivityManager.RegisterUserActivity:output_type -> google.protobuf.Empty
	3, // 5: activity.ActivityManager.RegisterGsActivity:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_activity_proto_init() }
func file_activity_proto_init() {
	if File_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_activity_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterUserActivityRequest); i {
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
		file_activity_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterGsActivityRequest); i {
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
			RawDescriptor: file_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activity_proto_goTypes,
		DependencyIndexes: file_activity_proto_depIdxs,
		MessageInfos:      file_activity_proto_msgTypes,
	}.Build()
	File_activity_proto = out.File
	file_activity_proto_rawDesc = nil
	file_activity_proto_goTypes = nil
	file_activity_proto_depIdxs = nil
}
