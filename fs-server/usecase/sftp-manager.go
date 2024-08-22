package usecase

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

func (fm *FileManager) ConnectToSFTP(ctx context.Context, host, password string, port int32, username string) (string, error) {

	if fm.sftpConnected {
		return "SFTP connection already established", nil
	}
	// Connect to the SFTP server
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// Enable SSH debugging
		// LogLevel: ssh.LogLevelDebug,
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		fm.log.Error("Failed to connect to SFTP server", zap.Error(err))
		return "", err
	}
	// defer conn.Close()

	// Open an SFTP session over the existing SSH connection
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		fm.log.Error("Failed to initialize SFTP session", zap.Error(err))
		return "", err
	}
	fm.sftpClient = sftpClient
	fm.sftpConnected = true
	// defer sftpClient.Close()
	fm.log.Debug("Successfully connected to sftp server")
	return "Success", nil

}

// func (fm *FileManager) DeleteFileFromSFTP(path, filename string) error {
// 	sftpClient, err := fm.GetSFTPClient()
// 	if err != nil {
// 		return fmt.Errorf("SFTP client not connected: %w", err)
// 	}
// 	fullPath:=filepath.Join(fm.VolumeDir,path,filename)

// }

func (fm *FileManager) DisconnectFromSFTP() {
	if fm.sftpClient != nil {
		fm.sftpClient.Close()
		fm.sftpClient = nil
	}
}

func (fm *FileManager) GetSFTPClient() (*sftp.Client, error) {
	if !fm.sftpConnected {
		return nil, errors.New("SFTP client not connected")
	}
	return fm.sftpClient, nil
}

// func (fm *FileManager) UploadFileToSFTP(ctx context.Context, host string, port int32, username, password, filename, remotepath string) (string, error) {
// 	// Connect to the SFTP server
// 	config := &ssh.ClientConfig{
// 		User: username,
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password(password),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 		// Enable SSH debugging
// 		// LogLevel: ssh.LogLevelDebug,
// 	}
// 	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
// 	if err != nil {
// 		fm.log.Error("Failed to connect to SFTP server", zap.Error(err))
// 		return "", err
// 	}
// 	defer conn.Close()

// 	// Open an SFTP session over the existing SSH connection
// 	sftpClient, err := sftp.NewClient(conn)
// 	if err != nil {
// 		fm.log.Error("Failed to initialize SFTP session", zap.Error(err))
// 		return "", err
// 	}
// 	defer sftpClient.Close()

// }

func (fm *FileManager) DeleteFileFromSFTP(path, filename string) error {
	sftpClient, err := fm.GetSFTPClient()
	if err != nil {
		fm.log.Error("error in getting sftp client", zap.Error(err))
		return err
	}
	fullPath := filepath.Join("incoming", path, filename)

	// TODO:implement file path validator
	deletedPath := filepath.Join("incoming", path, ".trash-"+filename)

	err = sftpClient.Rename(fullPath, deletedPath)
	if err != nil {
		return fmt.Errorf("could not move to trash: %w", err)
	}
	return nil
}
