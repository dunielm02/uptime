package ssh

import (
	"io"
	"log"
	"net"
)

type PortForward struct {
	shhProps   SshConnectionProps
	localAddrs string
}

func (p *PortForward) ForwardPort() {
	localListener, err := net.Listen("tcp", p.localAddrs)
	if err != nil {
		log.Fatalf("net.Listen failed: %v", err)
	}

	for {
		localConn, err := localListener.Accept()
		if err != nil {
			log.Fatalf("listen.Accept failed: %v", err)
		}
		go func(localConn net.Conn) {
			sshConn, err := p.shhProps.NewConnection()
			if err != nil {
				log.Println("Error ocurred while connecting via ssh: ", err)
			}

			go func() {
				_, err = io.Copy(sshConn, localConn)
				if err != nil {
					log.Fatalf("io.Copy failed: %v", err)
				}
			}()

			go func() {
				_, err = io.Copy(localConn, sshConn)
				if err != nil {
					log.Fatalf("io.Copy failed: %v", err)
				}
			}()
		}(localConn)
	}
}
