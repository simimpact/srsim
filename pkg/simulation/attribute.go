package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type snapshot struct {
	characters []*info.Stats
	enemies    []*info.Stats
	neutrals   []*info.Stats
}

func (sim *Simulation) createSnapshot() snapshot {
	charStats := make([]*info.Stats, len(sim.characters))
	for i, t := range sim.characters {
		charStats[i] = sim.Attr.Stats(t)
	}
	enemyStats := make([]*info.Stats, len(sim.enemies))
	for i, t := range sim.enemies {
		enemyStats[i] = sim.Attr.Stats(t)
	}
	neutralStats := make([]*info.Stats, len(sim.neutrals))
	for i, t := range sim.neutrals {
		neutralStats[i] = sim.Attr.Stats(t)
	}
	return snapshot{
		characters: charStats,
		enemies:    enemyStats,
		neutrals:   neutralStats,
	}
}

// TODO: move this to attr service?
func (sim *Simulation) ModifySP(amt int) int {
	old := sim.Sp
	sim.Sp += amt
	if sim.Sp > 5 {
		sim.Sp = 5
	}

	if old != sim.Sp {
		sim.Event.SPChange.Emit(event.SPChange{
			OldSP: old,
			NewSP: sim.Sp,
		})
	}
	return sim.Sp
}

func (sim *Simulation) SP() int {
	return sim.Sp
}

func (sim *Simulation) Stats(target key.TargetID) *info.Stats {
	return sim.Attr.Stats(target)
}

func (sim *Simulation) IsAlive(target key.TargetID) bool {
	return sim.Attr.State(target) == info.Alive
}

func (sim *Simulation) Stance(target key.TargetID) float64 {
	return sim.Attr.Stance(target)
}

func (sim *Simulation) Energy(target key.TargetID) float64 {
	return sim.Attr.Energy(target)
}

func (sim *Simulation) MaxEnergy(target key.TargetID) float64 {
	return sim.Attr.MaxEnergy(target)
}

func (sim *Simulation) EnergyRatio(target key.TargetID) float64 {
	return sim.Attr.EnergyRatio(target)
}

func (sim *Simulation) HPRatio(target key.TargetID) float64 {
	return sim.Attr.HPRatio(target)
}

func (sim *Simulation) SetHP(target, source key.TargetID, amt float64) error {
	sim.ActionTargets[target] = true
	return sim.Attr.SetHP(target, source, amt, false)
}

func (sim *Simulation) ModifyHPByRatio(target, source key.TargetID, data info.ModifyHPByRatio) error {
	sim.ActionTargets[target] = true
	return sim.Attr.ModifyHPByRatio(target, source, data, false)
}

func (sim *Simulation) ModifyHPByAmount(target, source key.TargetID, amt float64) error {
	sim.ActionTargets[target] = true
	return sim.Attr.ModifyHPByAmount(target, source, amt, false)
}

func (sim *Simulation) ModifyStance(target, source key.TargetID, amt float64) error {
	sim.ActionTargets[target] = true
	return sim.Attr.ModifyStance(target, source, amt)
}

func (sim *Simulation) ModifyEnergy(target key.TargetID, amt float64) error {
	sim.ActionTargets[target] = true
	return sim.Attr.ModifyEnergy(target, amt)
}

func (sim *Simulation) ModifyEnergyFixed(target key.TargetID, amt float64) error {
	sim.ActionTargets[target] = true
	return sim.Attr.ModifyEnergyFixed(target, amt)
}

func (sim *Simulation) SetGauge(target key.TargetID, amt float64) error {
	sim.ActionTargets[target] = true
	return sim.Turn.SetGauge(target, amt)
}

func (sim *Simulation) ModifyGaugeNormalized(target key.TargetID, amt float64) error {
	sim.ActionTargets[target] = true
	return sim.Turn.ModifyGaugeNormalized(target, amt)
}

func (sim *Simulation) ModifyGaugeAV(target key.TargetID, amt float64) error {
	sim.ActionTargets[target] = true
	return sim.Turn.ModifyGaugeAV(target, amt)
}

func (sim *Simulation) SetCurrentGaugeCost(amt float64) {
	sim.Turn.SetCurrentGaugeCost(amt)
}

func (sim *Simulation) ModifyCurrentGaugeCost(amt float64) {
	sim.Turn.ModifyCurrentGaugeCost(amt)
}
