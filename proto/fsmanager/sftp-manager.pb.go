// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: sftp-manager.proto

package fsmanager

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

type SFTPConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SFTPConnectRequest) Reset() {
	*x = SFTPConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sftp_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SFTPConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SFTPConnectRequest) ProtoMessage() {}

func (x *SFTPConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sftp_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SFTPConnectRequest.ProtoReflect.Descriptor instead.
func (*SFTPConnectRequest) Descriptor() ([]byte, []int) {
	return file_sftp_manager_proto_rawDescGZIP(), []int{0}
}

func (x *SFTPConnectRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SFTPConnectRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *SFTPConnectRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SFTPConnectRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SFTPConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *SFTPConnectResponse) Reset() {
	*x = SFTPConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sftp_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SFTPConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SFTPConnectResponse) ProtoMessage() {}

func (x *SFTPConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sftp_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SFTPConnectResponse.ProtoReflect.Descriptor instead.
func (*SFTPConnectResponse) Descriptor() ([]byte, []int) {
	return file_sftp_manager_proto_rawDescGZIP(), []int{1}
}

func (x *SFTPConnectResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SFTPConnectResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type SFTPUploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemotePath string `protobuf:"bytes,1,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
	FileName   string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Data       []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SFTPUploadFileRequest) Reset() {
	*x = SFTPUploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sftp_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SFTPUploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SFTPUploadFileRequest) ProtoMessage() {}

func (x *SFTPUploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sftp_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SFTPUploadFileRequest.ProtoReflect.Descriptor instead.
func (*SFTPUploadFileRequest) Descriptor() ([]byte, []int) {
	return file_sftp_manager_proto_rawDescGZIP(), []int{2}
}

func (x *SFTPUploadFileRequest) GetRemotePath() string {
	if x != nil {
		return x.RemotePath
	}
	return ""
}

func (x *SFTPUploadFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *SFTPUploadFileRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SFTPUploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *SFTPUploadFileResponse) Reset() {
	*x = SFTPUploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sftp_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SFTPUploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SFTPUploadFileResponse) ProtoMessage() {}

func (x *SFTPUploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sftp_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SFTPUploadFileResponse.ProtoReflect.Descriptor instead.
func (*SFTPUploadFileResponse) Descriptor() ([]byte, []int) {
	return file_sftp_manager_proto_rawDescGZIP(), []int{3}
}

func (x *SFTPUploadFileResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SFTPUploadFileResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type SFTPDownloadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemotePath string `protobuf:"bytes,1,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
	FileName   string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *SFTPDownloadFileRequest) Reset() {
	*x = SFTPDownloadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sftp_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SFTPDownloadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SFTPDownloadFileRequest) ProtoMessage() {}

func (x *SFTPDownloadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sftp_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SFTPDownloadFileRequest.ProtoReflect.Descriptor instead.
func (*SFTPDownloadFileRequest) Descriptor() ([]byte, []int) {
	return file_sftp_manager_proto_rawDescGZIP(), []int{4}
}

func (x *SFTPDownloadFileRequest) GetRemotePath() string {
	if x != nil {
		return x.RemotePath
	}
	return ""
}

func (x *SFTPDownloadFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type SFTPDownloadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SFTPDownloadFileResponse) Reset() {
	*x = SFTPDownloadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sftp_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SFTPDownloadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SFTPDownloadFileResponse) ProtoMessage() {}

func (x *SFTPDownloadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sftp_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SFTPDownloadFileResponse.ProtoReflect.Descriptor instead.
func (*SFTPDownloadFileResponse) Descriptor() ([]byte, []int) {
	return file_sftp_manager_proto_rawDescGZIP(), []int{5}
}

func (x *SFTPDownloadFileResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_sftp_manager_proto protoreflect.FileDescriptor

var file_sftp_manager_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x66, 0x74, 0x70, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22,
	0x74, 0x0a, 0x12, 0x53, 0x46, 0x54, 0x50, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x54, 0x0a, 0x13, 0x53, 0x46, 0x54, 0x50, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x15, 0x53,
	0x46, 0x54, 0x50, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x16, 0x53, 0x46, 0x54, 0x50, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x57, 0x0a, 0x17, 0x53, 0x46, 0x54, 0x50, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x18, 0x53, 0x46, 0x54, 0x50,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x9b, 0x02, 0x0a, 0x0b, 0x53, 0x66, 0x74,
	0x70, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x54, 0x6f, 0x53, 0x46, 0x54, 0x50, 0x12, 0x1d, 0x2e, 0x66, 0x73, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x46, 0x54, 0x50, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x73, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x46, 0x54, 0x50, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0e, 0x53, 0x46,
	0x54, 0x50, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x66,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x46, 0x54, 0x50, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x66, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x46, 0x54, 0x50, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x5f, 0x0a, 0x10, 0x53, 0x46, 0x54, 0x50, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x66, 0x73, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x46, 0x54, 0x50, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x66, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x46, 0x54, 0x50, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x69, 0x63, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2d, 0x68, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x73, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sftp_manager_proto_rawDescOnce sync.Once
	file_sftp_manager_proto_rawDescData = file_sftp_manager_proto_rawDesc
)

func file_sftp_manager_proto_rawDescGZIP() []byte {
	file_sftp_manager_proto_rawDescOnce.Do(func() {
		file_sftp_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_sftp_manager_proto_rawDescData)
	})
	return file_sftp_manager_proto_rawDescData
}

var file_sftp_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sftp_manager_proto_goTypes = []any{
	(*SFTPConnectRequest)(nil),       // 0: fsmanager.SFTPConnectRequest
	(*SFTPConnectResponse)(nil),      // 1: fsmanager.SFTPConnectResponse
	(*SFTPUploadFileRequest)(nil),    // 2: fsmanager.SFTPUploadFileRequest
	(*SFTPUploadFileResponse)(nil),   // 3: fsmanager.SFTPUploadFileResponse
	(*SFTPDownloadFileRequest)(nil),  // 4: fsmanager.SFTPDownloadFileRequest
	(*SFTPDownloadFileResponse)(nil), // 5: fsmanager.SFTPDownloadFileResponse
}
var file_sftp_manager_proto_depIdxs = []int32{
	0, // 0: fsmanager.SftpManager.ConnectToSFTP:input_type -> fsmanager.SFTPConnectRequest
	2, // 1: fsmanager.SftpManager.SFTPUploadFile:input_type -> fsmanager.SFTPUploadFileRequest
	4, // 2: fsmanager.SftpManager.SFTPDownloadFile:input_type -> fsmanager.SFTPDownloadFileRequest
	1, // 3: fsmanager.SftpManager.ConnectToSFTP:output_type -> fsmanager.SFTPConnectResponse
	3, // 4: fsmanager.SftpManager.SFTPUploadFile:output_type -> fsmanager.SFTPUploadFileResponse
	5, // 5: fsmanager.SftpManager.SFTPDownloadFile:output_type -> fsmanager.SFTPDownloadFileResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sftp_manager_proto_init() }
func file_sftp_manager_proto_init() {
	if File_sftp_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sftp_manager_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SFTPConnectRequest); i {
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
		file_sftp_manager_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SFTPConnectResponse); i {
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
		file_sftp_manager_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SFTPUploadFileRequest); i {
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
		file_sftp_manager_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SFTPUploadFileResponse); i {
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
		file_sftp_manager_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SFTPDownloadFileRequest); i {
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
		file_sftp_manager_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SFTPDownloadFileResponse); i {
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
			RawDescriptor: file_sftp_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sftp_manager_proto_goTypes,
		DependencyIndexes: file_sftp_manager_proto_depIdxs,
		MessageInfos:      file_sftp_manager_proto_msgTypes,
	}.Build()
	File_sftp_manager_proto = out.File
	file_sftp_manager_proto_rawDesc = nil
	file_sftp_manager_proto_goTypes = nil
	file_sftp_manager_proto_depIdxs = nil
}