package checklife

import (
	"net"
	"strconv"
)

type TcpService struct {
	name     string
	hostName string
	port     int
	dialer   net.Dialer
	inverted bool `default:"false"`
}

func (service *TcpService) CheckLife() error {
	conn, err := service.dialer.Dial("tcp", service.hostName+":"+strconv.Itoa(service.port))

	if err != nil {
		return err
	}

	err = conn.Close()

	return err
}

func (service *TcpService) GetName() string {
	return service.name
}

func (service *TcpService) IsInverted() bool {
	return service.inverted
}
