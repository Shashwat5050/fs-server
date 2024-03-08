package usecase

import (
	"context"
	"fmt"

	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

func (fm *FileManager) ConnectToSFTP(ctx context.Context, host, password string, port int32, username string) (string, error) {

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
	defer conn.Close()

	// Open an SFTP session over the existing SSH connection
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		fm.log.Error("Failed to initialize SFTP session", zap.Error(err))
		return "", err
	}
	defer sftpClient.Close()

	fm.log.Debug("Successfully connected to sftp server")
	return "Success", nil

}
