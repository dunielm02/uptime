package checkers

import (
	"bytes"
	"fmt"
	"lifeChecker/config"
	"net/http"
	"time"
)

// HttpService represents a service that we will prove
// is it is alive via http(s)
type HttpService struct {
	name        string
	client      *http.Client
	inverted    bool
	waitingTime time.Duration
	HttpServiceSpec
}

type HttpServiceSpec struct {
	url                string
	method             string
	requestBody        []byte
	requestHeaders     map[string]string
	expectedStatusCode int
}

func getHttpServiceFromConfig(cfg config.ServiceConfig) *HttpService {
	spec := cfg.Spec.(HttpServiceSpec)

	client := &http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}

	return &HttpService{
		name:            cfg.Name,
		HttpServiceSpec: spec,
		inverted:        cfg.Inverted,
		waitingTime:     time.Duration(cfg.WaitingTime) * time.Second,
		client:          client,
	}
}

func (service *HttpService) CheckLife() (time.Duration, error) {
	req, err := http.NewRequest(service.method, service.url, bytes.NewReader(service.requestBody))
	if err != nil {
		return 0, fmt.Errorf("error initializing the request: %s", err.Error())
	}

	for key, val := range service.requestHeaders {
		req.Header.Set(key, val)
	}
	initReq := time.Now()
	res, err := service.client.Do(req)
	duration := time.Since(initReq)
	if err != nil {
		return 0, fmt.Errorf("error sending the request: %s", err.Error())
	}

	if res.StatusCode == service.expectedStatusCode {
		return duration, nil
	}

	return 0, fmt.Errorf("unexpected status code: %d", res.StatusCode)
}

func (service *HttpService) GetName() string {
	return service.name
}

func (service *HttpService) IsInverted() bool {
	return service.inverted
}

func (service *HttpService) GetQueueTime() time.Duration {
	return service.waitingTime
}
