package checkers

import (
	"lifeChecker/config"
	"log"
	"time"
)

type LifeChecker interface {
	GetName() string
	CheckLife() (time.Duration, error)
	IsInverted() bool
	GetQueueTime() time.Duration
}

func GetFromConfig(cfg config.ServiceConfig) LifeChecker {
	switch cfg.Type {
	case "http":
		return getHttpServiceFromConfig(cfg)
	case "tcp":
	case "ping":
	default:
		log.Printf("The type: %s is not recognized.\n", cfg.Type)
	}

	return nil
}
