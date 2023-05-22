package dummy

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
)

func init() {
	// simulation.RegisterEnemyFunc(key.DummyEnemy, New)
}

type enemy struct {
}

func New(engine *engine.Engine, id key.TargetID) (interface{}, error) {
	e := &enemy{}

	return e, nil
}

func (c *enemy) Exec(key.ActionType) {
	//all enemy logic goes here i guess?

}
