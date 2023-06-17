package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (sim *simulation) CharacterInstance(id key.TargetID) (info.CharInstance, error) {
	return sim.char.Get(id)
}

func (sim *simulation) CharacterInfo(id key.TargetID) (info.Character, error) {
	return sim.char.Info(id)
}

func (sim *simulation) EnemyInfo(id key.TargetID) (info.Enemy, error) {
	return sim.enemy.Info(id)
}
