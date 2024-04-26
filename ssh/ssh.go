package ssh

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

type SshConnection struct {
	username         string
	password         string
	serverAddrString string
	remoteAddrString string
	sshClientConn    *ssh.Client
	sshConn          net.Conn
}

func (c *SshConnection) Close() error {
	err := c.sshConn.Close()
	if err != nil {
		return err
	}

	err = c.sshClientConn.Close()
	if err != nil {
		return err
	}

	c.sshClientConn = nil
	c.sshConn = nil

	return nil
}

func (c *SshConnection) Start() (net.Conn, error) {
	config := &ssh.ClientConfig{
		User: c.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.password),
		},
	}

	var err error

	c.sshClientConn, err = ssh.Dial("tcp", c.serverAddrString, config)
	if err != nil {
		return nil, fmt.Errorf("ssh.Dial failed: %s", err)
	}

	c.sshConn, err = c.sshClientConn.Dial("tcp", c.remoteAddrString)
	if err != nil {
		return nil, fmt.Errorf("dial tcp filed by ssh: %s", err)
	}

	return c.sshConn, nil
}
