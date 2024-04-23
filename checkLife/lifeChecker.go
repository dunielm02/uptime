package checklife

type LifeChecker interface {
	GetName() string
	CheckLife() (bool, error)
}
