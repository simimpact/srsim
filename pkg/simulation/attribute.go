package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

// TODO: move this to attr service?
func (sim *simulation) ModifySP(amt int) int {
	old := sim.sp
	sim.sp += amt
	if sim.sp > 5 {
		sim.sp = 5
	}

	if old != sim.sp {
		sim.event.SPChange.Emit(event.SPChangeEvent{
			OldSP: old,
			NewSP: sim.sp,
		})
	}
	return sim.sp
}

func (sim *simulation) SP() int {
	return sim.sp
}

func (sim *simulation) Stats(target key.TargetID) *info.Stats {
	return sim.attr.Stats(target)
}

func (sim *simulation) Stance(target key.TargetID) float64 {
	return sim.attr.Stance(target)
}

func (sim *simulation) Energy(target key.TargetID) float64 {
	return sim.attr.Energy(target)
}

func (sim *simulation) HPRatio(target key.TargetID) float64 {
	return sim.attr.HPRatio(target)
}

func (sim *simulation) SetHP(target key.TargetID, source key.TargetID, amt float64) error {
	return sim.attr.SetHP(target, source, amt, false)
}

func (sim *simulation) ModifyHPByRatio(target key.TargetID, source key.TargetID, data info.ModifyHPByRatio) error {
	return sim.attr.ModifyHPByRatio(target, source, data, false)
}

func (sim *simulation) ModifyHPByAmount(target key.TargetID, source key.TargetID, amt float64) error {
	return sim.attr.ModifyHPByAmount(target, source, amt, false)
}

func (sim *simulation) ModifyStance(target key.TargetID, source key.TargetID, amt float64) error {
	return sim.attr.ModifyStance(target, source, amt)
}

func (sim *simulation) ModifyEnergy(target key.TargetID, amt float64) error {
	return sim.attr.ModifyEnergy(target, amt)
}

func (sim *simulation) ModifyEnergyFixed(target key.TargetID, amt float64) error {
	return sim.attr.ModifyEnergyFixed(target, amt)
}

func (sim *simulation) SetGauge(target key.TargetID, amt float64) error {
	return sim.turn.SetGauge(target, amt)
}

func (sim *simulation) ModifyGaugeNormalized(target key.TargetID, amt float64) error {
	return sim.turn.ModifyGaugeNormalized(target, amt)
}

func (sim *simulation) ModifyGaugeAV(target key.TargetID, amt float64) error {
	return sim.turn.ModifyGaugeAV(target, amt)
}

func (sim *simulation) SetCurrentGaugeCost(amt float64) {
	sim.turn.SetCurrentGaugeCost(amt)
}

func (sim *simulation) ModifyCurrentGaugeCost(amt float64) {
	sim.turn.ModifyCurrentGaugeCost(amt)
}
