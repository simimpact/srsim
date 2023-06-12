package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (sim *simulation) AddShield(id key.Shield, shield info.Shield) {
	sim.actionTargets[shield.Target] = true
	sim.shield.AddShield(id, shield)
}

func (sim *simulation) RemoveShield(id key.Shield, target key.TargetID) {
	sim.actionTargets[target] = true
	sim.shield.RemoveShield(id, target)
}

func (sim *simulation) HasShield(target key.TargetID, shield key.Shield) bool {
	return sim.shield.HasShield(target, shield)
}

func (sim *simulation) IsShielded(target key.TargetID) bool {
	return sim.shield.IsShielded(target)
}
