package checkers

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpCheckLife(t *testing.T) {
	go setUpTcpListener()

	t.Run("Accepted tcp connection", func(t *testing.T) {
		service := TcpService{
			hostName: "localhost",
			port:     8000,
		}

		_, err := service.CheckLife()

		assert.Nil(t, err, err)
	})

	t.Run("Refused tcp connection", func(t *testing.T) {
		service := TcpService{
			hostName: "localhost",
			port:     8001,
		}

		_, err := service.CheckLife()

		assert.NotNil(t, err, err)
	})

}

func setUpTcpListener() {
	l, err := net.Listen("tcp", ":8000")

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
