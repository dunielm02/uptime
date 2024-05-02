package checkers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// HttpService represents a service that we will prove
// is it is alive via http(s)
type HttpService struct {
	name               string
	url                string
	method             string
	client             *http.Client
	requestBody        []byte
	requestHeaders     map[string]string
	expectedStatusCode int
	timeout            int
	inverted           bool
}

func (service *HttpService) IsInverted() bool {
	return service.inverted
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
