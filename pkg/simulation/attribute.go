package simulation

import (
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

func (sim *Simulation) ModifySP(reason key.Reason, amt int) error {
	return sim.Attr.ModifySP(reason, amt)
}

func (sim *Simulation) SP() int {
	return sim.Attr.SP()
}

func (sim *Simulation) Stats(target key.TargetID) *info.Stats {
	return sim.Attr.Stats(target)
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

func (sim *Simulation) SetHP(data info.ModifyAttribute) error {
	sim.ActionTargets[data.Target] = true
	return sim.Attr.SetHP(data, false)
}

func (sim *Simulation) ModifyHPByRatio(data info.ModifyHPByRatio) error {
	sim.ActionTargets[data.Target] = true
	return sim.Attr.ModifyHPByRatio(data, false)
}

func (sim *Simulation) ModifyStance(data info.ModifyAttribute) error {
	sim.ActionTargets[data.Target] = true
	return sim.Attr.ModifyStance(data)
}

func (sim *Simulation) ModifyEnergy(data info.ModifyAttribute) error {
	sim.ActionTargets[data.Target] = true
	return sim.Attr.ModifyEnergy(data)
}

func (sim *Simulation) ModifyEnergyFixed(data info.ModifyAttribute) error {
	sim.ActionTargets[data.Target] = true
	return sim.Attr.ModifyEnergyFixed(data)
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
