package dummy

import (
	"github.com/simimpact/srsim/pkg/engine/simulation"
	"github.com/simimpact/srsim/pkg/engine/system"
	"github.com/simimpact/srsim/pkg/key"
)

func init() {
	simulation.RegisterEnemyFunc(key.DummyEnemy, New)
}

type enemy struct {

}

func New(sys *system.EnemyServices, id key.TargetID) (simulation.Target, error) {
	e := &enemy{}

	return e, nil
}


func (c *enemy) Exec(key.ActionType) {
	//all enemy logic goes here i guess?

}
