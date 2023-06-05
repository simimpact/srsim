package simulation

import "github.com/simimpact/srsim/pkg/key"

type TargetType int

const (
	TargetInvalid TargetType = iota
	TargetCharacter
	TargetEnemy
	TargetNeutral
)

// TODO: AddTarget
func (sim *simulation) AddTarget() key.TargetID {
	panic("not implemented") // TODO: Implement
}

func (sim *simulation) IsValid(target key.TargetID) bool {
	if _, ok := sim.targets[target]; ok {
		return true
	}
	return false
}

func (sim *simulation) IsCharacter(target key.TargetID) bool {
	if targetType, ok := sim.targets[target]; ok {
		return targetType == TargetCharacter
	}
	return false
}

func (sim *simulation) IsEnemy(target key.TargetID) bool {
	if targetType, ok := sim.targets[target]; ok {
		return targetType == TargetEnemy
	}
	return false
}

func (sim *simulation) IsNeutral(target key.TargetID) bool {
	if targetType, ok := sim.targets[target]; ok {
		return targetType == TargetNeutral
	}
	return false
}

func (sim *simulation) AdjacentTo(target key.TargetID) []key.TargetID {
	var targets []key.TargetID

	switch sim.targets[target] {
	case TargetCharacter:
		targets = sim.characters
	case TargetEnemy:
		targets = sim.enemies
	case TargetNeutral:
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

func (sim *simulation) Characters() []key.TargetID {
	return sim.characters
}

func (sim *simulation) Enemies() []key.TargetID {
	return sim.enemies
}

func (sim *simulation) Neutrals() []key.TargetID {
	return sim.neutrals
}
