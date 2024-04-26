package portForward

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

type Tunel struct {
	username         string
	password         string
	serverAddrString string
	// localAddrString  string
	remoteAddrString string
}

func (t *Tunel) Start() (net.Conn, error) {
	// Setup SSH config (type *ssh.ClientConfig)
	config := &ssh.ClientConfig{
		User: t.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(t.password),
		},
	}

	// Setup sshClientConn (type *ssh.ClientConn)
	sshClientConn, err := ssh.Dial("tcp", t.serverAddrString, config)
	if err != nil {
		return nil, fmt.Errorf("ssh.Dial failed: %s", err)
	}

	// Setup sshConn (type net.Conn)
	sshConn, err := sshClientConn.Dial("tcp", t.remoteAddrString)
	if err != nil {
		return nil, fmt.Errorf("dial tcp filed by ssh: %s", err)
	}

	return sshConn, nil
}
