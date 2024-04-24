package checklife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Add a ping receiver for make the test process better.

func TestPingCheckLife(t *testing.T) {
	t.Run("Accepted ping", func(t *testing.T) {
		service := PingService{
			host:        "127.0.0.1",
			timeout:     5,
			pingCount:   4,
			mostReceive: 4,
		}

		err := service.CheckLife()

		assert.Nil(t, err, err)
	})

	t.Run("Failing Ping: Host Does not Exist", func(t *testing.T) {
		service := PingService{
			host:        "192.168.255.255",
			timeout:     5,
			pingCount:   4,
			mostReceive: 4,
		}

		err := service.CheckLife()

		assert.NotNil(t, err, err)
	})
}
