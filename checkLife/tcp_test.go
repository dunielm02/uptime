package checklife

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTcpCheckLife(t *testing.T) {
	go setUpTcpListener()

	t.Run("Accepted tcp connection", func(t *testing.T) {
		service := TcpService{
			hostName: "localhost",
			port:     8000,
		}

		err := service.CheckLife()

		assert.Nil(t, err, err)
	})

	t.Run("Refused tcp connection", func(t *testing.T) {
		service := TcpService{
			hostName: "localhost",
			port:     8001,
		}

		err := service.CheckLife()

		assert.NotNil(t, err, err)
	})

}

func setUpTcpListener() {
	l, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("error creating listening: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("error Accepting the incoming connection")
		}

		go func() {
			time.Sleep(2 * time.Second)
			conn.Close()
		}()
	}
}
