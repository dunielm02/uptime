package ssh

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSsh(t *testing.T) {
	forward := PortForward{
		shhProps: sshConnectionProps{
			username:         "test",
			password:         "test",
			serverAddrString: "localhost:3030",
			remoteAddrString: "localhost:8000",
		},
		localAddrs: "localhost:8000",
	}

	ctx, cancel := context.WithCancel(context.Background())

	go forward.ForwardPort(ctx)

	time.Sleep(time.Second)

	res, err := http.DefaultClient.Get("http://localhost:8000/hello")
	assert.Nil(t, err)

	bts, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, "Hello, World from a container", string(bts))

	cancel()

	_, err = http.DefaultClient.Get("http://localhost:8000/hello")
	assert.NotNil(t, err)
}
