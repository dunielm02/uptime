package checkers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Add a ping receiver for make the test process better.

func TestPingCheckLife(t *testing.T) {
	t.Run("Accepted ping", func(t *testing.T) {
		service := PingService{
			PingServiceSpec: PingServiceSpec{
				Host:        "127.0.0.1",
				PingCount:   4,
				MustReceive: 4,
			},
			timeout: 5,
		}

		_, err := service.CheckLife()

		assert.Nil(t, err, err)
	})

	t.Run("Failing Ping: Host Does not Exist", func(t *testing.T) {
		service := PingService{
			PingServiceSpec: PingServiceSpec{
				Host:        "192.168.255.255",
				PingCount:   4,
				MustReceive: 4,
			},
			timeout: 5,
		}

		_, err := service.CheckLife()

		assert.NotNil(t, err, err)
	})
}
