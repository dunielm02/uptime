package checklife

import "time"

type LifeChecker interface {
	IsInverted() bool
	GetName() string
	CheckLife() (time.Duration, error)
}
