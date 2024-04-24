package checklife

import (
	"fmt"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

//ADD IPV6 PING

type PingService struct {
	name        string
	host        string
	timeout     int
	pingCount   int
	mostReceive int
	inverted    bool `default:"false"`
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

func (service *PingService) IsInverted() bool {
	return service.inverted
}

func (service *PingService) GetName() string {
	return service.name
}
