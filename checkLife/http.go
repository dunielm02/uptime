package checklife

import (
	"bytes"
	"fmt"
	"net/http"
)

// HttpService represents a service that we will prove
// is it is alive via http(s)
type HttpService struct {
	name               string
	url                string
	method             string
	client             *http.Client
	requestBody        []byte
	expectedStatusCode int
}

func (service *HttpService) CheckLife() error {
	req, err := http.NewRequest(service.method, service.url, bytes.NewReader(service.requestBody))
	if err != nil {
		return fmt.Errorf("error initializing the request: %s", err.Error())
	}

	res, err := service.client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending the request: %s", err.Error())
	}

	if res.StatusCode == service.expectedStatusCode {
		return nil
	}

	return fmt.Errorf("unexpected status code: %d", res.StatusCode)
}

func (service *HttpService) GetName() string {
	return service.name
}
