package checkers

import (
	"lifeChecker/config"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
)

type TcpService struct {
	name        string
	dialer      net.Dialer
	waitingTime time.Duration
	inverted    bool
	TcpServiceSpec
}

type TcpServiceSpec struct {
	HostName string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

func getTcpServiceFromConfig(cfg config.ServiceConfig) *TcpService {
	var spec TcpServiceSpec

	err := mapstructure.Decode(cfg.Spec, &spec)
	if err != nil {
		log.Fatal("error converting tcp spec: ", err)
	}

	dialer := net.Dialer{
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}

	return &TcpService{
		name:           cfg.Name,
		waitingTime:    time.Duration(cfg.WaitingTime) * time.Second,
		dialer:         dialer,
		inverted:       cfg.Inverted,
		TcpServiceSpec: spec,
	}
}

func (service *TcpService) CheckLife() (time.Duration, error) {
	initDial := time.Now()
	conn, err := service.dialer.Dial("tcp", service.HostName+":"+strconv.Itoa(service.Port))
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
