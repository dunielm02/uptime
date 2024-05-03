package checkers

import (
	"lifeChecker/config"
	"net"
	"strconv"
	"time"
)

type TcpService struct {
	name        string
	dialer      net.Dialer
	waitingTime time.Duration
	inverted    bool
	TcpServiceSpec
}

type TcpServiceSpec struct {
	hostName string
	port     int
}

func getTcpServiceFromConfig(cfg config.ServiceConfig) *TcpService {
	dialer := net.Dialer{
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}

	return &TcpService{
		name:        cfg.Name,
		waitingTime: time.Duration(cfg.WaitingTime) * time.Second,
		dialer:      dialer,
		inverted:    cfg.Inverted,
	}
}

func (service *TcpService) CheckLife() (time.Duration, error) {
	initDial := time.Now()
	conn, err := service.dialer.Dial("tcp", service.hostName+":"+strconv.Itoa(service.port))
	duration := time.Since(initDial)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	return duration, err
}

func (service *TcpService) GetQueueTime() time.Duration {
	return service.waitingTime
}

func (service *TcpService) GetName() string {
	return service.name
}

func (service *TcpService) IsInverted() bool {
	return service.inverted
}
