package checkers

import (
	"net"
	"strconv"
	"time"
)

type TcpService struct {
	name     string
	hostName string
	port     int
	dialer   net.Dialer
	inverted bool `default:"false"`
}

func (service *TcpService) CheckLife() (time.Duration, error) {
	initDial := time.Now()
	conn, err := service.dialer.Dial("tcp", service.hostName+":"+strconv.Itoa(service.port))
	duration := time.Since(initDial)
	if err != nil {
		return 0, err
	}

	err = conn.Close()

	return duration, err
}

func (service *TcpService) GetName() string {
	return service.name
}

func (service *TcpService) IsInverted() bool {
	return service.inverted
}
