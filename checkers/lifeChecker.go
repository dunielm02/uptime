package checkers

import "time"

type LifeChecker interface {
	GetName() string
	CheckLife() (time.Duration, error)
	IsInverted() bool
	GetQueueTime() time.Duration
}
