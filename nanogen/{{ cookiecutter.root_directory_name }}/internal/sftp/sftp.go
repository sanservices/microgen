package sftp

import (
	"context"
	"fmt"
	"function/internal/config"
	"log"

	"github.com/pkg/sftp"
	"github.com/sanservices/apilogger/v2"
	"golang.org/x/crypto/ssh"
)

func New(ctx context.Context, config *config.Settings) (*sftp.Client, *ssh.Client, error) {

	apilogger.Info(ctx, apilogger.LogCatUncategorized, "Connecting to sftp server...")

	var auths []ssh.AuthMethod
	auths = append(auths, ssh.Password(config.Sftp.Password))

	cc := ssh.ClientConfig{
		User:            config.Sftp.Username,
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%d", config.Sftp.Host, config.Sftp.Port)

	// Establish an SSH connection
	sshClient, err := ssh.Dial("tcp", addr, &cc)
	if err != nil {
		log.Printf("Failed to connect to [%s]: %v\n", config.Sftp.Host, err)

		return nil, nil, err
	}

	// Create new SFTP client
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Printf("Unable to start sftp subsystem: %v\n", err)

		return nil, nil, err
	}

	apilogger.Info(ctx, apilogger.LogCatUncategorized, "Connected to sftp server")

	return sftpClient, sshClient, nil
}
