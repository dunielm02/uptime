package ssh

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type sshConnectionProps struct {
	username         string
	password         string
	serverAddrString string
	remoteAddrString string
}

func (c *sshConnectionProps) NewConnection() (net.Conn, error) {
	config := &ssh.ClientConfig{
		HostKeyCallback: trustedHostKeyCallback(""),
		User:            c.username,
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

func keyString(k ssh.PublicKey) string {
	return k.Type() + " " + base64.StdEncoding.EncodeToString(k.Marshal()) // e.g. "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTY...."
}

func trustedHostKeyCallback(trustedKey string) ssh.HostKeyCallback {

	if trustedKey == "" {
		return func(_ string, _ net.Addr, k ssh.PublicKey) error {
			log.Printf("WARNING: SSH-key verification is *NOT* in effect: to fix, add this trustedKey: %q", keyString(k))
			return nil
		}
	}

	return func(_ string, _ net.Addr, k ssh.PublicKey) error {
		ks := keyString(k)
		if trustedKey != ks {
			return fmt.Errorf("SSH-key verification: expected %q but got %q", trustedKey, ks)
		}

		return nil
	}
}
