package v1

import (
	"context"
	"io"
	"path/filepath"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	fpb "iceline-hosting.com/backend/proto/fsmanager"
)

func (c *controller) GetFileStat(ctx context.Context, req *fpb.GetFileStatRequest) (*fpb.GetFileStatResponse, error) {
	c.log.Info("GetFileStat", zap.String("path", req.Path))

	fileStat, err := c.use.GetFileStat(req.Path)
	if err != nil {
		return nil, err
	}

	return &fpb.GetFileStatResponse{
		Stat: &fpb.FileInfo{
			Name:    fileStat.Name(),
			Size:    fileStat.Size(),
			Mode:    uint32(fileStat.Mode()),
			ModTime: fileStat.ModTime().Unix(),
			IsDir:   fileStat.IsDir(),
		},
	}, nil
}

func (c *controller) ListFilePath(ctx context.Context, req *fpb.ListFilePathRequest) (*fpb.ListFilePathResponse, error) {
	c.log.Info("ListFilePath", zap.String("path", req.Path))

	filePaths, err := c.use.ListFilePath(req.Path)
	if err != nil {
		return nil, err
	}

	var fileInfos []*fpb.DirEntry
	for _, filePath := range filePaths {
		info, err := filePath.Info()
		if err != nil {
			return nil, err
		}

		fileInfos = append(fileInfos, &fpb.DirEntry{
			Name:  filePath.Name(),
			IsDir: filePath.IsDir(),
			Type:  uint32(filePath.Type()),
			Info: &fpb.FileInfo{
				Name:    info.Name(),
				Size:    info.Size(),
				Mode:    uint32(info.Mode()),
				ModTime: info.ModTime().Unix(),
				IsDir:   info.IsDir(),
			},
		})
	}

	return &fpb.ListFilePathResponse{
		FileList: fileInfos,
	}, nil
}

func (c *controller) CreatePath(ctx context.Context, req *fpb.CreatePathRequest) (*emptypb.Empty, error) {
	c.log.Info("CreatePath", zap.String("path", req.Path), zap.String("name", req.Name))

	if req.IsDir {
		err := c.use.CreateDir(req.Path, req.Name)
		if err != nil {
			return nil, err
		}
	} else {
		w, err := c.use.CreateFile(req.Path, req.Name)
		if err != nil {
			return nil, err
		}

		err = w.Close()
		if err != nil {
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}

func (c *controller) DownloadFile(req *fpb.DownloadFileRequest, stream fpb.FsManager_DownloadFileServer) error {
	c.log.Info("DownloadFile", zap.String("path", req.Path), zap.String("name", req.Name))

	n, body, err := c.use.DownloadFile(req.Path, req.Name)
	if err != nil {
		return err
	}

	i := 0
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			if i >= int(n) {
				return nil
			}

			buf := make([]byte, 1024)
			readN, err := body.Read(buf)
			if err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}

			err = stream.Send(&fpb.DownloadFileResponse{
				Data: buf,
			})
			if err != nil {
				return err
			}

			i += readN
		}
	}
}

func (c *controller) BulkDownloadFile(req *fpb.BulkDownloadFileRequest, stream fpb.FsManager_BulkDownloadFileServer) error {
	c.log.Info("BulkDownloadFile", zap.String("path", req.Path), zap.Strings("file names", req.FileNameList))

	outputPath, err := c.use.BulkCompressFile(stream.Context(), req.Path, req.FileNameList)
	if err != nil {
		return err
	}

	outputPath, fileName := filepath.Split(outputPath)

	n, body, err := c.use.DownloadFile(outputPath, fileName)
	if err != nil {
		return err
	}

	i := 0
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			if i >= int(n) {
				return nil
			}

			buf := make([]byte, 1024)
			readN, err := body.Read(buf)
			if err == io.EOF {
				return nil
			} else if err != nil {
				return err
			}

			err = stream.Send(&fpb.DownloadFileResponse{
				Data: buf,
			})
			if err != nil {
				return err
			}

			i += readN
		}
	}
}

func (c *controller) UploadFile(req fpb.FsManager_UploadFileServer) error {
	c.log.Info("UploadFile")

	var (
		filename, path string
		fd             io.WriteCloser
		size           uint32
	)

	for {
		if err := req.Context().Err(); err != nil {
			return err
		}

		data, err := req.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if filename == "" {
			filename = data.GetFileName()
			path = data.GetPath()
			fd, err = c.use.CreateFile(path, filename)
			if err != nil {
				return err
			}
		}

		n, err := fd.Write(data.GetData())
		if err != nil {
			return err
		}

		size += uint32(n)
	}

	err := fd.Close()
	if err != nil {
		return err
	}

	err = req.SendAndClose(&fpb.UploadFileResponse{
		FileName: filename,
		Path:     path,
		Size:     size,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *controller) DeleteFile(ctx context.Context, req *fpb.DeleteFileRequest) (*emptypb.Empty, error) {
	c.log.Info("DeleteFile", zap.String("path", req.Path), zap.String("name", req.Name))

	err := c.use.DeleteFile(req.Path, req.Name)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (c *controller) RecoverFile(ctx context.Context, req *fpb.RecoverFileRequest) (*emptypb.Empty, error) {
	c.log.Info("RecoverFile", zap.String("path", req.Path), zap.String("name", req.Name))

	err := c.use.RecoverFile(req.Path, req.Name)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (c *controller) MoveFile(ctx context.Context, req *fpb.TwoFileRequest) (*emptypb.Empty, error) {
	c.log.Info("MoveFile", zap.String("path", req.Path), zap.String("name", req.Name), zap.String("newPath", req.NewPath))

	err := c.use.MoveFile(req.Path, req.Name, req.NewPath, req.NewName)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (c *controller) CopyFile(ctx context.Context, req *fpb.TwoFileRequest) (*emptypb.Empty, error) {
	c.log.Info("CopyFile", zap.String("path", req.Path), zap.String("name", req.Name), zap.String("newPath", req.NewPath))

	err := c.use.CopyFile(req.Path, req.Name, req.NewPath, req.NewName)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (c *controller) CompressFile(ctx context.Context, req *fpb.CompressFileRequest) (*emptypb.Empty, error) {
	c.log.Info("CompressFile", zap.String("path", req.Path), zap.String("name", req.Name))

	err := c.use.CompressFile(ctx, req.Path, req.Name)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (c *controller) BulkCompressFile(ctx context.Context, req *fpb.BulkCompressFileRequest) (*emptypb.Empty, error) {
	c.log.Info("BulkCompressFile", zap.String("path", req.Path), zap.Strings("file names", req.FileNameList))

	_, err := c.use.BulkCompressFile(ctx, req.Path, req.FileNameList)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
