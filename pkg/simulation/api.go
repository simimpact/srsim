package simulation

import (
	"math/rand"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (sim *Simulation) Events() *event.System {
	return sim.event
}

func (sim *Simulation) Rand() *rand.Rand {
	return sim.rand
}

func (sim *Simulation) AddModifier(target key.TargetID, instance info.Modifier) error {
	return sim.modManager.AddModifier(target, instance)
}

func (sim *Simulation) RemoveModifier(target key.TargetID, modifier key.Modifier) {
	sim.modManager.RemoveModifier(target, modifier)
}

func (sim *Simulation) RemoveModifierFromSource(target key.TargetID, source key.TargetID, modifier key.Modifier) {
	sim.modManager.RemoveModifierFromSource(target, source, modifier)
}

func (sim *Simulation) ExtendModifierDuration(target key.TargetID, modifier key.Modifier, amt int) {
	sim.modManager.ExtendDuration(target, modifier, amt)
}

func (sim *Simulation) ExtendModifierCount(target key.TargetID, modifier key.Modifier, amt float64) {
	sim.modManager.ExtendCount(target, modifier, amt)
}

func (sim *Simulation) HasModifier(target key.TargetID, modifier key.Modifier) bool {
	return sim.modManager.HasModifier(target, modifier)
}

func (sim *Simulation) ModifierCount(target key.TargetID, statusType model.StatusType) int {
	state := sim.modManager.EvalModifiers(target)
	return state.Counts[statusType]
}

func (sim *Simulation) HasBehaviorFlag(target key.TargetID, flags ...model.BehaviorFlag) bool {
	return sim.modManager.HasFlag(target, flags...)
}

func (sim *Simulation) Stats(target key.TargetID) *info.Stats {
	return sim.attributeService.Stats(target)
}

func (sim *Simulation) CharacterInfo(target key.TargetID) (info.Character, error) {
	return sim.charManager.Info(target)
}

func (sim *Simulation) EnemyInfo(target key.TargetID) (model.Enemy, error) {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) IsCharacter(target key.TargetID) bool {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) IsEnemy(target key.TargetID) bool {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) AdjacentTo(target key.TargetID) []key.TargetID {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) Characters() []key.TargetID {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) Enemies() []key.TargetID {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) Neutrals() []key.TargetID {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) IsValid(target key.TargetID) bool {
	// TODO: Implement
	return target != 0
}

func (sim *Simulation) ModifyGauge(target key.TargetID, modifyType model.ModifyGauge, amt float64) {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) SetGauge(target key.TargetID, amt float64) {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) Attack(atk info.Attack) {
	// TODO:
	sim.combatManager.Attack(atk, model.AttackEffect_INVALID_ATTACK_EFFECT)
}

func (sim *Simulation) Heal(heal info.Heal) {
	sim.combatManager.Heal(heal)
}

func (sim *Simulation) AddShield() {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) RemoveShield() {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) AddTarget() key.TargetID {
	panic("not implemented") // TODO: Implement
}

func (sim *Simulation) SetHP(target key.TargetID, source key.TargetID, amt float64) error {
	return sim.attributeService.SetHP(target, source, amt)
}

func (sim *Simulation) ModifyHPByRatio(target key.TargetID, source key.TargetID, data info.ModifyHPByRatio) error {
	return sim.attributeService.ModifyHPByRatio(target, source, data)
}

func (sim *Simulation) ModifyHPByAmount(target key.TargetID, source key.TargetID, amt float64) error {
	return sim.attributeService.ModifyHPByAmount(target, source, amt)
}

func (sim *Simulation) ModifyStance(target key.TargetID, source key.TargetID, amt float64) error {
	return sim.attributeService.ModifyStance(target, source, amt)
}

func (sim *Simulation) ModifyEnergy(target key.TargetID, amt float64) error {
	return sim.attributeService.ModifyEnergy(target, amt)
}

func (sim *Simulation) ModifyEnergyFixed(target key.TargetID, amt float64) error {
	return sim.attributeService.ModifyEnergyFixed(target, amt)
}
