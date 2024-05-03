package checkers

import (
	"fmt"
	"lifeChecker/config"
	"time"

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
	host        string
	pingCount   int
	mostReceive int
}

func getPingServiceFromConfig(cfg config.ServiceConfig) *PingService {
	spec := cfg.Spec.(PingServiceSpec)

	return &PingService{
		name:            cfg.Name,
		timeout:         cfg.Timeout,
		waitingTime:     time.Duration(cfg.WaitingTime) * time.Second,
		inverted:        cfg.Inverted,
		PingServiceSpec: spec,
	}
}

func (service *PingService) CheckLife() (time.Duration, error) {
	pinger, err := probing.NewPinger(service.host)
	if err != nil {
		return 0, fmt.Errorf("error creating the pinger: %v", err)
	}

	pinger.SetPrivileged(true)
	pinger.Timeout = time.Duration(service.timeout) * time.Second
	pinger.Count = service.pingCount

	err = pinger.Run()
	if err != nil {
		return 0, fmt.Errorf("error running the pinger: %v", err)
	}

	if pinger.Statistics().PacketsRecv < service.mostReceive {
		return 0, fmt.Errorf("only %d packets were received from %d", service.pingCount, pinger.Statistics().PacketsRecv)
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
