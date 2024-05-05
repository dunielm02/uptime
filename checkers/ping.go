package checkers

import (
	"fmt"
	"lifeChecker/config"
	"log"
	"time"

	"github.com/mitchellh/mapstructure"
	probing "github.com/prometheus-community/pro-bing"
)

//ADD IPV6 PING

type PingService struct {
	name        string
	timeout     int
	waitingTime time.Duration
	inverted    bool
	PingServiceSpec
}

type PingServiceSpec struct {
	Host        string `mapstructure:"host"`
	PingCount   int    `mapstructure:"ping-count"`
	MostReceive int    `mapstructure:"most-receive"`
}

func getPingServiceFromConfig(cfg config.ServiceConfig) *PingService {
	var spec PingServiceSpec

	err := mapstructure.Decode(cfg.Spec, &spec)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(spec)

	return &PingService{
		name:            cfg.Name,
		timeout:         cfg.Timeout,
		waitingTime:     time.Duration(cfg.WaitingTime) * time.Second,
		inverted:        cfg.Inverted,
		PingServiceSpec: spec,
	}
}

func (service *PingService) CheckLife() (time.Duration, error) {
	pinger, err := probing.NewPinger(service.Host)
	if err != nil {
		return 0, fmt.Errorf("error creating the pinger: %v", err)
	}

	pinger.SetPrivileged(true)
	pinger.Timeout = time.Duration(service.timeout) * time.Second
	pinger.Count = service.PingCount

	err = pinger.Run()
	if err != nil {
		return 0, fmt.Errorf("error running the pinger: %v", err)
	}

	if pinger.Statistics().PacketsRecv < service.MostReceive {
		return 0, fmt.Errorf("only %d packets were received from %d", service.PingCount, pinger.Statistics().PacketsRecv)
	}

	return pinger.Statistics().AvgRtt, nil
}

func (service *PingService) GetQueueTime() time.Duration {
	return service.waitingTime
}

func (service *PingService) IsInverted() bool {
	return service.inverted
}

func (service *PingService) GetName() string {
	return service.name
}
