package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (sim *Simulation) AddModifier(target key.TargetID, instance info.Modifier) (bool, error) {
	sim.ActionTargets[target] = true
	return sim.Modifier.AddModifier(target, instance)
}

func (sim *Simulation) RemoveModifier(target key.TargetID, modifier key.Modifier) {
	sim.ActionTargets[target] = true
	sim.Modifier.RemoveModifier(target, modifier)
}

func (sim *Simulation) RemoveModifierFromSource(target, source key.TargetID, modifier key.Modifier) {
	sim.ActionTargets[target] = true
	sim.Modifier.RemoveModifierFromSource(target, source, modifier)
}

func (sim *Simulation) DispelStatus(target key.TargetID, dispel info.Dispel) {
	sim.ActionTargets[target] = true
	sim.Modifier.DispelStatus(target, dispel)
}

func (sim *Simulation) ExtendModifierDuration(target key.TargetID, modifier key.Modifier, amt int) {
	sim.ActionTargets[target] = true
	sim.Modifier.ExtendDuration(target, modifier, amt)
}

func (sim *Simulation) ExtendModifierCount(target key.TargetID, modifier key.Modifier, amt float64) {
	sim.ActionTargets[target] = true
	sim.Modifier.ExtendCount(target, modifier, amt)
}

func (sim *Simulation) HasModifier(target key.TargetID, modifier key.Modifier) bool {
	return sim.Modifier.HasModifier(target, modifier)
}

func (sim *Simulation) HasModifierFromSource(target, source key.TargetID, modifier key.Modifier) bool {
	return sim.Modifier.HasModifierFromSource(target, source, modifier)
}

func (sim *Simulation) ModifierStatusCount(target key.TargetID, statusType model.StatusType) int {
	state := sim.Modifier.EvalModifiers(target)
	return state.Counts[statusType]
}

func (sim *Simulation) HasBehaviorFlag(target key.TargetID, flags ...model.BehaviorFlag) bool {
	return sim.Modifier.HasFlag(target, flags...)
}

func (sim *Simulation) GetModifiers(target key.TargetID, modifier key.Modifier) []info.Modifier {
	return sim.Modifier.GetModifiers(target, modifier)
}

func (sim *Simulation) ModifierStackCount(target, source key.TargetID, modifier key.Modifier) float64 {
	return sim.Modifier.ModifierStackCount(target, source, modifier)
}
