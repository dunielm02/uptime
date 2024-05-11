package mocks

import (
	"log"
	"net"
)

func StartTcpServiceMock() {
	l, err := net.Listen("tcp", ":3201")

	defer func() {
		err := l.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	if err != nil {
		log.Fatalf("error creating listening: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("error Accepting the incoming connection")
		}

		go func() {
			conn.Close()
		}()
	}
}
