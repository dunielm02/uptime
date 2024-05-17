package ssh

import (
	"context"
	"io"
	"lifeChecker/config"
	"log"
	"net"
)

type PortForward struct {
	shhProps   sshConnectionProps
	localAddrs string
}

func GetProtFrowardFromConfig(cfg config.PortForwardConfig) PortForward {
	return PortForward{
		localAddrs: cfg.LocalAddress,
		shhProps: sshConnectionProps{
			username:         cfg.Username,
			password:         cfg.Password,
			serverAddrString: cfg.ServerAddress,
			remoteAddrString: cfg.RemoteAddress,
		},
	}
}

func (p *PortForward) ForwardPort(ctx context.Context) {
	localListener, err := net.Listen("tcp", p.localAddrs)
	if err != nil {
		log.Fatalf("net.Listen failed: %v", err)
	}
	for {
		select {
		case <-ctx.Done():
			localListener.Close()
			return
		default:
			localConn, err := localListener.Accept()
			if err != nil {
				log.Fatalf("listen.Accept failed: %v", err)
			}
			go p.connectWithRemoteServer(localConn, ctx)
		}
	}
}

func (p *PortForward) connectWithRemoteServer(localConn net.Conn, ctx context.Context) {
	sshConn, err := p.shhProps.NewConnection()
	if err != nil {
		log.Println("Error ocurred while connecting via ssh: ", err)
	}

	go func() {
		io.Copy(sshConn, localConn)
	}()

	go func() {
		io.Copy(localConn, sshConn)
	}()

	go func() {
		<-ctx.Done()
		localConn.Close()
		sshConn.Close()
	}()
}
