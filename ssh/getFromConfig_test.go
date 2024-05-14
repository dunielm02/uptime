package ssh

import (
	"lifeChecker/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFromConfig(t *testing.T) {
	config := config.GetConfigFromYamlFile("../uptime-config.yml")

	PortForward := GetProtFrowardFromConfig(config.PortForwards[0])

	assert.Equal(t, PortForward.localAddrs, "localhost:8000")
	assert.Equal(t, PortForward.shhProps.serverAddrString, "www.host.com")
	assert.Equal(t, PortForward.shhProps.password, "<YourPassword>")
	assert.Equal(t, PortForward.shhProps.username, "<YourUsername>")
	assert.Equal(t, PortForward.shhProps.remoteAddrString, "localhost:5432")
}
