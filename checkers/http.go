package checkers

import (
	"bytes"
	"fmt"
	"lifeChecker/config"
	"log"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
)

// HttpService represents a service that we will prove
// is it is alive via http(s)
type HttpService struct {
	name                 string
	client               *http.Client
	inverted             bool
	waitingTime          time.Duration
	notificationChannels []string
	HttpServiceSpec
}

type HttpServiceSpec struct {
	Url                string            `mapstructure:"url"`
	Method             string            `mapstructure:"method"`
	RequestBody        string            `mapstructure:"body"`
	RequestHeaders     map[string]string `mapstructure:"headers"`
	ExpectedStatusCode int               `mapstructure:"expected-status"`
}

func getHttpServiceFromConfig(cfg config.ServiceConfig) *HttpService {
	var spec HttpServiceSpec

	fmt.Println(cfg.Spec)

	err := mapstructure.Decode(cfg.Spec, &spec)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(spec)

	client := &http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}

	return &HttpService{
		name:                 cfg.Name,
		waitingTime:          time.Duration(cfg.WaitingTime) * time.Second,
		client:               client,
		inverted:             cfg.Inverted,
		notificationChannels: cfg.NotificationChannels,
		HttpServiceSpec:      spec,
	}
}

func (service *HttpService) CheckLife() (time.Duration, error) {
	req, err := http.NewRequest(service.Method, service.Url, bytes.NewReader([]byte(service.RequestBody)))
	if err != nil {
		return 0, fmt.Errorf("error initializing the request: %s", err.Error())
	}

	for key, val := range service.RequestHeaders {
		req.Header.Set(key, val)
	}
	initReq := time.Now()
	res, err := service.client.Do(req)
	duration := time.Since(initReq)
	if err != nil {
		return 0, fmt.Errorf("error sending the request: %s", err.Error())
	}

	if res.StatusCode == service.ExpectedStatusCode {
		return duration, nil
	}

	return 0, fmt.Errorf("unexpected status code: %d", res.StatusCode)
}

func (service *HttpService) GetNotificationChannelsNames() []string {
	return service.notificationChannels
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
