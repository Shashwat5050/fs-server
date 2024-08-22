package v1

import (
	"context"
	"fmt"
	"io"

	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"google.golang.org/protobuf/types/known/emptypb"

	sfp "iceline-hosting.com/backend/proto/sftpmanager"
)

func (c *controller) ConnectToSFTP(ctx context.Context, req *sfp.SFTPConnectRequest) (*sfp.SFTPConnectResponse, error) {
	c.log.Info("Connect from outside project", zap.String("user", req.Username))

	msg, err := c.use.ConnectToSFTP(ctx, req.Host, req.Password, req.Port, req.Username)
	if err != nil {
		return nil, err
	}

	return &sfp.SFTPConnectResponse{
		Success:      true,
		ErrorMessage: msg,
	}, nil
}

func (c *controller) SFTPUploadFile(req sfp.SftpManager_SFTPUploadFileServer) error {
	c.log.Debug("UploadFileToSFTP")

	var (
		filename, path string
		size           uint32
	)

	// Establish connection to SFTP server
	// data, err := req.Recv()
	// if err != nil {
	// 	return err
	// }
	// sftpClient, err := GetSftpClient(context.Background(), data.Host, data.Password, data.Port, data.Username)
	sftpClient, err := c.use.GetSFTPClient()
	if err != nil {
		c.log.Error("error in getting sftp client", zap.Error(err))
		return err
	}
	// defer sftpClient.Close()

	for {
		if err := req.Context().Err(); err != nil {
			c.log.Error("error in request context", zap.Error(err))
			return err
		}
		data, err := req.Recv()
		if err == io.EOF {
			c.log.Error("error is eof", zap.Error(err))
			break
		} else if err != nil {
			c.log.Error("error while request receiving", zap.Error(err))
			return err
		}
		if filename == "" {
			filename = data.GetFileName()
			path = data.GetRemotePath()
		}
		err = sftpClient.MkdirAll("incoming/" + path)
		if err != nil {
			c.log.Error("error creating directories", zap.Error(err))
			return err
		}
		// Open a file on the SFTP server for writing
		c.log.Debug("remote path", zap.Any("remotepath", path))
		newpath := "incoming/" + path
		dstFile, err := sftpClient.Create(newpath + "/" + filename)
		// dstFile, err := sftpClient.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)

		if err != nil {
			c.log.Error("error occured in creating file", zap.Error(err))
			return err
		}

		// Write the file data to the SFTP server
		n, err := dstFile.Write(data.GetData())
		if err != nil {
			c.log.Error("error in writing in destination file", zap.Error(err))
			return err
		}

		size += uint32(n)

		// Close the destination file
		err = dstFile.Close()
		if err != nil {
			c.log.Error("error in closing destination file", zap.Error(err))
			return err
		}
		err = req.SendAndClose(&sfp.SFTPUploadFileResponse{
			FileName: filename,
			Path:     path,
			Size:     size,
		})
		if err != nil {
			c.log.Error("error in send and close request", zap.Error(err))
			return err
		}
		return nil
		// Receive next chunk of data

	}

	c.log.Info("File uploaded successfully", zap.String("fileName", filename), zap.String("path", path), zap.Uint32("size", size))
	return nil
}

func (c *controller) SFTPDeleteFile(ctx context.Context, req *sfp.SFTPDeleteFileRequest) (*emptypb.Empty, error) {
	c.log.Info("Delete file from sftp", zap.String("path", req.Path), zap.String("name", req.Name))

	err := c.use.DeleteFileFromSFTP(req.Path, req.Name)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// func(c *controller)SFTPDeleteFile(ctx context.Context,req *sfp.DeleteFileRequest)(*emptypb.Empty,error){
// 	c.log.Info("DeleteFile",zap.String("path",req.Path),zap.String("name",req.Name))

// 	err:=c.use.DeleteFileFromSFTP(req.Path,req.Name)
// 	if err!=nil{
// 		return nil,err
// 	}
// 	return &emptypb.Empty{},nil
// }

func GetSftpClient(ctx context.Context, host, password string, port int32, username string) (*sftp.Client, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return nil, err
	}

	sftpClient, err := sftp.NewClient(sshConn)
	if err != nil {
		return nil, err
	}

	return sftpClient, nil
}
