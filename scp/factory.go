package scp

import (
	"github.com/paashzj/gl/ssh"
	"github.com/pkg/sftp"
)

func NewScpClient(host, user, password string, port uint) (*sftp.Client, error) {
	client, err := ssh.NewPasswordClient(host, user, password, port)
	if err != nil {
		return nil, err
	}
	return client.NewSftp()
}
