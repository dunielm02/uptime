package checkers

import (
	"lifeChecker/config"
	"log"
	"time"
)

type State byte

const (
	NoStatus State = iota
	StatusAlive
	StatusDead
)

type LifeChecker interface {
	GetName() string
	CheckLife() (time.Duration, error)
	IsInverted() bool
	GetQueueTime() time.Duration
	GetNotificationChannelsNames() []string
	GetState() State
}

func GetFromConfig(cfg config.ServiceConfig) LifeChecker {
	switch cfg.Type {
	case "http":
		return getHttpServiceFromConfig(cfg)
	case "tcp":
		return getTcpServiceFromConfig(cfg)
	case "ping":
		return getPingServiceFromConfig(cfg)
	default:
		log.Printf("The type: %s is not recognized.\n", cfg.Type)
	}

	return nil
}
