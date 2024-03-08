package v1

import (
	"context"

	"go.uber.org/zap"
	bpb "iceline-hosting.com/backend/proto/backupmanager"
)

func (c *controller) CreateBackup(ctx context.Context, req *bpb.CreateBackupRequest) (*bpb.CreateBackupResponse, error) {
	c.log.Info("CreateBackup", zap.String("path", req.Path))

	size, name, err := c.use.BackupFile(ctx, req.Path)
	if err != nil {
		c.log.Error("CreateBackup", zap.Error(err))
		return nil, err
	}

	return &bpb.CreateBackupResponse{
		Size: size,
		Name: name,
	}, nil
}

func (c *controller) RestoreBackup(ctx context.Context, req *bpb.RestoreBackupRequest) (*bpb.RestoreBackupResponse, error) {
	c.log.Info("RestoreBackup", zap.String("path", req.Path))

	if err := c.use.RestoreFile(ctx, req.Path, req.Name); err != nil {
		c.log.Error("RestoreBackup", zap.Error(err))
		return nil, err
	}

	return &bpb.RestoreBackupResponse{}, nil
}
