package v1

import (
	"context"
	"io"
	"io/fs"
	"log"
	"net"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	bpb "iceline-hosting.com/backend/proto/backupmanager"
	fpb "iceline-hosting.com/backend/proto/fsmanager"
	sfpb "iceline-hosting.com/backend/proto/fsmanager"
)

type fsUsecase interface {
	GetFileStat(path string) (fs.FileInfo, error)
	ListFilePath(path string) ([]fs.DirEntry, error)
	CreateDir(path, name string) error
	DownloadFile(path, name string) (int64, io.ReadCloser, error)
	CreateFile(path, name string) (io.WriteCloser, error)
	UploadFile(ctx context.Context, path, name string, file io.Reader) error
	DeleteFile(path, name string) error
	RecoverFile(path, name string) error
	MoveFile(path, name, newPath, newName string) error
	CopyFile(path, name, newPath, newName string) error
	CompressFile(ctx context.Context, path, name string) error
	BulkCompressFile(ctx context.Context, path string, filenames []string) (string, error)
	GetFileData(path string) (string, error)
	SetFileData(path string, data []byte) error
}

type backupUsecase interface {
	BackupFile(ctx context.Context, path string) (uint64, string, error)
	RestoreFile(ctx context.Context, path, name string) error
}

type sftpUseCase interface {
	ConnectToSFTP(ctx context.Context, host, password string, port int32, username string) (string, error)
}

type usecase interface {
	fsUsecase
	backupUsecase
	sftpUseCase
}

type controller struct {
	use  usecase
	port string

	log    *otelzap.Logger
	server *grpc.Server
	fpb.UnimplementedFsManagerServer
	bpb.UnimplementedBackupManagerServer
	sfpb.UnimplementedSftpManagerServer
}

type settings func(*controller)

func WithPort(port string) settings {
	return func(c *controller) {
		c.port = port
	}
}

const (
	defaultPort = "50000"
)

func NewController(logger *otelzap.Logger, use usecase, opts ...settings) *controller {
	c := &controller{
		log:  logger,
		use:  use,
		port: defaultPort,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.server = grpc.NewServer()
	return c
}

func (c *controller) Run() error {
	listener, err := net.Listen("tcp", ":"+c.port)
	if err != nil {
		log.Println("err at 81*******************************************************")
		return err
	}

	fpb.RegisterFsManagerServer(c.server, c)
	c.log.Info("registered fs manager server")

	bpb.RegisterBackupManagerServer(c.server, c)
	c.log.Info("registered backup manager server")

	sfpb.RegisterSftpManagerServer(c.server, c)
	c.log.Info("registered sftp manager server")

	c.log.Info("started listening", zap.String("port", c.port))
	return c.server.Serve(listener)
}

func (c *controller) Stop() {
	c.server.Stop()
}
