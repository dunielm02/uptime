package checklife

import "time"

type LifeChecker interface {
	GetName() string
	CheckLife() (time.Duration, error)
	IsInverted() bool
}
