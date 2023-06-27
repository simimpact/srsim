package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (sim *Simulation) CharacterInstance(id key.TargetID) (info.CharInstance, error) {
	return sim.Char.Get(id)
}

func (sim *Simulation) CharacterInfo(id key.TargetID) (info.Character, error) {
	return sim.Char.Info(id)
}

func (sim *Simulation) EnemyInfo(id key.TargetID) (info.Enemy, error) {
	return sim.Enemy.Info(id)
}
