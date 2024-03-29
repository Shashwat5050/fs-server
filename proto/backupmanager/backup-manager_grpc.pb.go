// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: backup-manager.proto

package backupmanager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BackupManagerClient is the client API for BackupManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackupManagerClient interface {
	CreateBackup(ctx context.Context, in *CreateBackupRequest, opts ...grpc.CallOption) (*CreateBackupResponse, error)
	RestoreBackup(ctx context.Context, in *RestoreBackupRequest, opts ...grpc.CallOption) (*RestoreBackupResponse, error)
}

type backupManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupManagerClient(cc grpc.ClientConnInterface) BackupManagerClient {
	return &backupManagerClient{cc}
}

func (c *backupManagerClient) CreateBackup(ctx context.Context, in *CreateBackupRequest, opts ...grpc.CallOption) (*CreateBackupResponse, error) {
	out := new(CreateBackupResponse)
	err := c.cc.Invoke(ctx, "/backupmanager.BackupManager/CreateBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupManagerClient) RestoreBackup(ctx context.Context, in *RestoreBackupRequest, opts ...grpc.CallOption) (*RestoreBackupResponse, error) {
	out := new(RestoreBackupResponse)
	err := c.cc.Invoke(ctx, "/backupmanager.BackupManager/RestoreBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupManagerServer is the server API for BackupManager service.
// All implementations must embed UnimplementedBackupManagerServer
// for forward compatibility
type BackupManagerServer interface {
	CreateBackup(context.Context, *CreateBackupRequest) (*CreateBackupResponse, error)
	RestoreBackup(context.Context, *RestoreBackupRequest) (*RestoreBackupResponse, error)
	mustEmbedUnimplementedBackupManagerServer()
}

// UnimplementedBackupManagerServer must be embedded to have forward compatible implementations.
type UnimplementedBackupManagerServer struct {
}

func (UnimplementedBackupManagerServer) CreateBackup(context.Context, *CreateBackupRequest) (*CreateBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBackup not implemented")
}
func (UnimplementedBackupManagerServer) RestoreBackup(context.Context, *RestoreBackupRequest) (*RestoreBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreBackup not implemented")
}
func (UnimplementedBackupManagerServer) mustEmbedUnimplementedBackupManagerServer() {}

// UnsafeBackupManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackupManagerServer will
// result in compilation errors.
type UnsafeBackupManagerServer interface {
	mustEmbedUnimplementedBackupManagerServer()
}

func RegisterBackupManagerServer(s grpc.ServiceRegistrar, srv BackupManagerServer) {
	s.RegisterService(&BackupManager_ServiceDesc, srv)
}

func _BackupManager_CreateBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupManagerServer).CreateBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backupmanager.BackupManager/CreateBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupManagerServer).CreateBackup(ctx, req.(*CreateBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupManager_RestoreBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupManagerServer).RestoreBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backupmanager.BackupManager/RestoreBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupManagerServer).RestoreBackup(ctx, req.(*RestoreBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BackupManager_ServiceDesc is the grpc.ServiceDesc for BackupManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackupManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backupmanager.BackupManager",
	HandlerType: (*BackupManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBackup",
			Handler:    _BackupManager_CreateBackup_Handler,
		},
		{
			MethodName: "RestoreBackup",
			Handler:    _BackupManager_RestoreBackup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backup-manager.proto",
}
