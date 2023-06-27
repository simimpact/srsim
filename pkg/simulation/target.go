package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/info"
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

func (sim *Simulation) AdjacentTo(target key.TargetID) []key.TargetID {
	var targets []key.TargetID

	switch sim.Targets[target] {
	case info.ClassCharacter:
		targets = sim.characters
	case info.ClassEnemy:
		targets = sim.enemies
	case info.ClassNeutral:
		targets = sim.neutrals
	default:
		targets = nil
	}

	for i, t := range targets {
		if t != target {
			continue
		}

		out := make([]key.TargetID, 0, 3)
		if i != 0 {
			out = append(out, targets[i-1])
		}
		out = append(out, t)
		if i != len(targets)-1 {
			out = append(out, targets[i+1])
		}
		return out
	}
	return nil
}

func (sim *Simulation) Characters() []key.TargetID {
	return sim.characters
}

func (sim *Simulation) Enemies() []key.TargetID {
	return sim.enemies
}

func (sim *Simulation) Neutrals() []key.TargetID {
	return sim.neutrals
}
