package ssh

import (
	"github.com/melbahja/goph"
)

func NewPasswordClient(host, user, password string, port uint) (c *goph.Client, err error) {
	auth := goph.Password(password)
	return goph.NewConn(&goph.Config{
		Auth:     auth,
		User:     user,
		Addr:     host,
		Port:     port,
		Timeout:  0,
		Callback: trustHostCallback,
	})

}