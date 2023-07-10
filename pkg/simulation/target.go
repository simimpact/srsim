package simulation

import (
	"math/rand"

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
	// first filter out all <= 0 HP targets IF includeLimbo is FALSE or default(which is also FALSE).
	targetList := data.Targets
	if !data.IncludeLimbo {
		for _, target := range targetList {
			if sim.HPRatio(target) <= 0 {
				targetList = remove(targetList, target)
			}
		}
	}

	// filter targetList. assign to validTargets. if not provided, bypass
	var validTargets []key.TargetID
	if data.Filter != nil {
		validTargets = data.Filter(targetList)
	} else {
		validTargets = targetList
	}

	// NOTE : handle cases when len(validTargets) == 0 (target list empty after filter).
	if len(validTargets) == 0 {
		return validTargets
	}

	// choose Max amount of targets from valid candidates.
	// deal with edge cases
	if data.Max == 0 {
		data.Max = 1
	} else if data.Max > len(validTargets) {
		data.Max = len(validTargets)
	}

	// shuffle list indexes randomly
	rand.Shuffle(len(validTargets), func(i, j int) {
		validTargets[i], validTargets[j] = validTargets[j], validTargets[i]
	})

	// choose max amt of targets
	chosenOnes := validTargets[0:data.Max]

	// return filtered list
	return chosenOnes
}
