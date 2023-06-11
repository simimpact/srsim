package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (sim *simulation) AddModifier(target key.TargetID, instance info.Modifier) (bool, error) {
	return sim.modifier.AddModifier(target, instance)
}

func (sim *simulation) RemoveModifier(target key.TargetID, modifier key.Modifier) {
	sim.modifier.RemoveModifier(target, modifier)
}

func (sim *simulation) RemoveModifierFromSource(target key.TargetID, source key.TargetID, modifier key.Modifier) {
	sim.modifier.RemoveModifierFromSource(target, source, modifier)
}

func (sim *simulation) ExtendModifierDuration(target key.TargetID, modifier key.Modifier, amt int) {
	sim.modifier.ExtendDuration(target, modifier, amt)
}

func (sim *simulation) ExtendModifierCount(target key.TargetID, modifier key.Modifier, amt float64) {
	sim.modifier.ExtendCount(target, modifier, amt)
}

func (sim *simulation) HasModifier(target key.TargetID, modifier key.Modifier) bool {
	return sim.modifier.HasModifier(target, modifier)
}

func (sim *simulation) ModifierCount(target key.TargetID, statusType model.StatusType) int {
	state := sim.modifier.EvalModifiers(target)
	return state.Counts[statusType]
}

func (sim *simulation) HasBehaviorFlag(target key.TargetID, flags ...model.BehaviorFlag) bool {
	return sim.modifier.HasFlag(target, flags...)
}
