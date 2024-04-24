package checklife

import (
	"fmt"

	probing "github.com/prometheus-community/pro-bing"
)

type PingService struct {
	name      string
	host      string
	pingCount int  `default:"4"`
	inverted  bool `default:"false"`
}

func (service *PingService) CheckLife() error {
	pinger, err := probing.NewPinger(service.host)
	if err != nil {
		return fmt.Errorf("error creating the pinger: %v", err)
	}

	pinger.Count = service.pingCount
	err = pinger.Run()
	if err != nil {
		return fmt.Errorf("error running the pinger: %v", err)
	}
	
	return nil
}

func (service *PingService) IsInverted() bool {
	return service.inverted
}

func (service *PingService) GetName() string {
	return service.name
}
