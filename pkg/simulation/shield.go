package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (sim *Simulation) AddShield(id key.Shield, shield info.Shield) {
	sim.ActionTargets[shield.Target] = true
	sim.Shield.AddShield(id, shield)
}

func (sim *Simulation) RemoveShield(id key.Shield, target key.TargetID) {
	sim.ActionTargets[target] = true
	sim.Shield.RemoveShield(id, target)
}

func (sim *Simulation) HasShield(target key.TargetID, shield key.Shield) bool {
	return sim.Shield.HasShield(target, shield)
}

func (sim *Simulation) IsShielded(target key.TargetID) bool {
	return sim.Shield.IsShielded(target)
}
