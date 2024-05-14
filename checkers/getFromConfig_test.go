package checkers

import (
	"lifeChecker/config"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetFromConfig(t *testing.T) {
	config := config.GetConfigFromYamlFile("../uptime-config.yml")

	for _, i := range config.Services {
		if i.Type == "http" {
			service := getHttpServiceFromConfig(i)
			compHttpResult(t, service)
		} else if i.Type == "tcp" {
			service := getTcpServiceFromConfig(i)
			compTcpResult(t, service)
		} else if i.Type == "ping" {
			service := getPingServiceFromConfig(i)
			compPingResult(t, service)
		}
	}
}

func compHttpResult(t *testing.T, service *HttpService) {
	var http_result = HttpService{
		name:                 "My Http Service",
		inverted:             false,
		waitingTime:          60 * time.Second,
		notificationChannels: []string{"telegram"},
		state:                NoStatus,
		HttpServiceSpec: HttpServiceSpec{
			Url:    "google.com",
			Method: "POST",
			RequestHeaders: map[string]string{
				"content-type":   "application/json",
				"authentication": "bearer",
			},
			ExpectedStatusCode: 201,
			RequestBody:        "{\"ID\": 4524}",
		},
	}
	assert.Equal(t, service.state, http_result.state)
	assert.Equal(t, service.name, http_result.name)
	assert.Equal(t, service.inverted, http_result.inverted)
	assert.Equal(t, service.waitingTime, http_result.waitingTime)
	assert.Equal(t, service.Url, http_result.Url)
	assert.Equal(t, service.Method, http_result.Method)
	assert.Equal(t, service.ExpectedStatusCode, http_result.ExpectedStatusCode)
	assert.Equal(t, service.RequestBody, http_result.RequestBody)

	for i := 0; i < len(service.notificationChannels); i++ {
		assert.Equal(t, service.notificationChannels[i], http_result.notificationChannels[i])
	}

	for k := range service.RequestHeaders {
		assert.Equal(t, service.RequestHeaders[k], http_result.RequestHeaders[k])
	}
}

func compTcpResult(t *testing.T, service *TcpService) {
	var tcp_result = TcpService{
		name:                 "My Tcp Service",
		inverted:             false,
		notificationChannels: []string{"telegram"},
		waitingTime:          60 * time.Second,
		state:                NoStatus,
		TcpServiceSpec: TcpServiceSpec{
			HostName: "www.google.com",
			Port:     3000,
		},
	}
	assert.Equal(t, service.state, tcp_result.state)
	assert.Equal(t, service.name, tcp_result.name)
	assert.Equal(t, service.inverted, tcp_result.inverted)
	assert.Equal(t, service.waitingTime, tcp_result.waitingTime)
	assert.Equal(t, service.HostName, tcp_result.HostName)
	assert.Equal(t, service.Port, tcp_result.Port)

	for i := 0; i < len(service.notificationChannels); i++ {
		assert.Equal(t, service.notificationChannels[i], tcp_result.notificationChannels[i])
	}
}

func compPingResult(t *testing.T, service *PingService) {
	var ping_result = PingService{
		name:                 "My Ping Service",
		inverted:             false,
		notificationChannels: []string{"telegram"},
		waitingTime:          60 * time.Second,
		state:                NoStatus,
		PingServiceSpec: PingServiceSpec{
			Host:        "www.google.com",
			PingCount:   4,
			MustReceive: 4,
		},
	}
	assert.Equal(t, service.state, ping_result.state)
	assert.Equal(t, service.name, ping_result.name)
	assert.Equal(t, service.inverted, ping_result.inverted)
	assert.Equal(t, service.waitingTime, ping_result.waitingTime)
	assert.Equal(t, service.Host, ping_result.Host)
	assert.Equal(t, service.PingCount, ping_result.PingCount)
	assert.Equal(t, service.MustReceive, ping_result.MustReceive)

	for i := 0; i < len(service.notificationChannels); i++ {
		assert.Equal(t, service.notificationChannels[i], ping_result.notificationChannels[i])
	}
}
