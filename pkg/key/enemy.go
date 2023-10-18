package key

type Enemy string

const (
	DummyEnemy Enemy = "dummy"
)

func (e Enemy) String() string {
	return string(e)
}
