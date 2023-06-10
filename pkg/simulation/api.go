package simulation

import (
	"math/rand"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (sim *simulation) Events() *event.System {
	return sim.event
}

func (sim *simulation) Rand() *rand.Rand {
	return sim.rand
}

func (sim *simulation) AddModifier(target key.TargetID, instance info.Modifier) (bool, error) {
	return sim.modManager.AddModifier(target, instance)
}

func (sim *simulation) RemoveModifier(target key.TargetID, modifier key.Modifier) {
	sim.modManager.RemoveModifier(target, modifier)
}

func (sim *simulation) RemoveModifierFromSource(target key.TargetID, source key.TargetID, modifier key.Modifier) {
	sim.modManager.RemoveModifierFromSource(target, source, modifier)
}

func (sim *simulation) ExtendModifierDuration(target key.TargetID, modifier key.Modifier, amt int) {
	sim.modManager.ExtendDuration(target, modifier, amt)
}

func (sim *simulation) ExtendModifierCount(target key.TargetID, modifier key.Modifier, amt float64) {
	sim.modManager.ExtendCount(target, modifier, amt)
}

func (sim *simulation) HasModifier(target key.TargetID, modifier key.Modifier) bool {
	return sim.modManager.HasModifier(target, modifier)
}

func (sim *simulation) ModifierCount(target key.TargetID, statusType model.StatusType) int {
	state := sim.modManager.EvalModifiers(target)
	return state.Counts[statusType]
}

func (sim *simulation) HasBehaviorFlag(target key.TargetID, flags ...model.BehaviorFlag) bool {
	return sim.modManager.HasFlag(target, flags...)
}

func (sim *simulation) Stats(target key.TargetID) *info.Stats {
	return sim.attributeService.Stats(target)
}

func (sim *simulation) CharacterInfo(target key.TargetID) (info.Character, error) {
	return sim.charManager.Info(target)
}

func (sim *simulation) EnemyInfo(target key.TargetID) (info.Enemy, error) {
	panic("not implemented") // TODO: Implement
}

func (sim *simulation) SetGauge(target key.TargetID, amt float64) error {
	return sim.turnManager.SetGauge(target, amt)
}

func (sim *simulation) ModifyGaugeNormalized(target key.TargetID, amt float64) error {
	return sim.turnManager.ModifyGaugeNormalized(target, amt)
}

func (sim *simulation) ModifyGaugeAV(target key.TargetID, amt float64) error {
	return sim.turnManager.ModifyGaugeAV(target, amt)
}

func (sim *simulation) SetCurrentGaugeCost(amt float64) {
	sim.turnManager.SetCurrentGaugeCost(amt)
}

func (sim *simulation) ModifyCurrentGaugeCost(amt float64) {
	sim.turnManager.ModifyCurrentGaugeCost(amt)
}

func (sim *simulation) Attack(atk info.Attack) {
	// TODO:
	sim.combatManager.Attack(atk, model.SkillEffect_INVALID_SKILL_EFFECT)
}

func (sim *simulation) Heal(heal info.Heal) {
	sim.combatManager.Heal(heal)
}

func (sim *simulation) AddShield() {
	panic("not implemented") // TODO: Implement
}

func (sim *simulation) RemoveShield() {
	panic("not implemented") // TODO: Implement
}

func (sim *simulation) SetHP(target key.TargetID, source key.TargetID, amt float64) error {
	return sim.attributeService.SetHP(target, source, amt)
}

func (sim *simulation) ModifyHPByRatio(target key.TargetID, source key.TargetID, data info.ModifyHPByRatio) error {
	return sim.attributeService.ModifyHPByRatio(target, source, data)
}

func (sim *simulation) ModifyHPByAmount(target key.TargetID, source key.TargetID, amt float64) error {
	return sim.attributeService.ModifyHPByAmount(target, source, amt)
}

func (sim *simulation) ModifyStance(target key.TargetID, source key.TargetID, amt float64) error {
	return sim.attributeService.ModifyStance(target, source, amt)
}

func (sim *simulation) ModifyEnergy(target key.TargetID, amt float64) error {
	return sim.attributeService.ModifyEnergy(target, amt)
}

func (sim *simulation) ModifyEnergyFixed(target key.TargetID, amt float64) error {
	return sim.attributeService.ModifyEnergyFixed(target, amt)
}
