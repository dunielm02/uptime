package mocks

import (
	"lifeChecker/checkers"
	"time"
)

type Mock_checkLifeService struct {
	Name                      string
	Inverted                  bool
	Status                    checkers.State
	NotificationChannelsNames []string
	Err                       error
}

func (s *Mock_checkLifeService) GetName() string                   { return s.Name }
func (s *Mock_checkLifeService) CheckLife() (time.Duration, error) { return time.Duration(100), s.Err }
func (s *Mock_checkLifeService) IsInverted() bool                  { return s.Inverted }
func (s *Mock_checkLifeService) GetQueueTime() time.Duration       { return time.Duration(1) * time.Second }
func (s *Mock_checkLifeService) GetNotificationChannelsNames() []string {
	return s.NotificationChannelsNames
}
func (s *Mock_checkLifeService) GetState() checkers.State { return s.Status }
