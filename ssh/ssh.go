package ssh

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

type SshConnectionProps struct {
	username         string
	password         string
	serverAddrString string
	remoteAddrString string
}

func (c *SshConnectionProps) NewConnection() (net.Conn, error) {
	config := &ssh.ClientConfig{
		User: c.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.password),
		},
	}

	sshClientConn, err := ssh.Dial("tcp", c.serverAddrString, config)
	if err != nil {
		return nil, fmt.Errorf("ssh.Dial failed: %s", err)
	}

	sshConn, err := sshClientConn.Dial("tcp", c.remoteAddrString)
	if err != nil {
		return nil, fmt.Errorf("dial tcp filed by ssh: %s", err)
	}

	return sshConn, nil
}
