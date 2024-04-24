package checklife

type LifeChecker interface {
	IsInverted() bool
	GetName() string
	CheckLife() error
}
