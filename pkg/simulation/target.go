package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target"
	"github.com/simimpact/srsim/pkg/key"
)

// TODO: AddTarget
func (sim *Simulation) AddNeutralTarget() key.TargetID {
	panic("not implemented") // TODO: Implement
}

// TODO: AddTarget
func (sim *Simulation) RemoveNeutralTarget(id key.TargetID) {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) IsValid(target key.TargetID) bool {
	if _, ok := sim.Targets[target]; ok {
		return true
	}
	return false
}

func (sim *Simulation) IsCharacter(target key.TargetID) bool {
	if targetType, ok := sim.Targets[target]; ok {
		return targetType == info.ClassCharacter
	}
	return false
}

func (sim *Simulation) IsEnemy(target key.TargetID) bool {
	if targetType, ok := sim.Targets[target]; ok {
		return targetType == info.ClassEnemy
	}
	return false
}

func (sim *Simulation) IsNeutral(target key.TargetID) bool {
	if targetType, ok := sim.Targets[target]; ok {
		return targetType == info.ClassNeutral
	}
	return false
}

func (sim *Simulation) AdjacentTo(t key.TargetID) []key.TargetID {
	var targets []key.TargetID

	switch sim.Targets[t] {
	case info.ClassCharacter:
		targets = sim.characters
	case info.ClassEnemy:
		targets = sim.enemies
	case info.ClassNeutral:
		targets = sim.neutrals
	default:
		targets = nil
	}

	return target.AdjacentTo(targets, t)
}

func (sim *Simulation) Characters() []key.TargetID {
	out := make([]key.TargetID, len(sim.characters))
	copy(out, sim.characters)
	return out
}

func (sim *Simulation) Enemies() []key.TargetID {
	out := make([]key.TargetID, len(sim.enemies))
	copy(out, sim.enemies)
	return out
}

func (sim *Simulation) Neutrals() []key.TargetID {
	out := make([]key.TargetID, len(sim.neutrals))
	copy(out, sim.neutrals)
	return out
}

func (sim *Simulation) Retarget(data info.Retarget) []key.TargetID {
	return target.Retarget(sim.Rand(), sim.Attr, data)
}
