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

func (sim *Simulation) ModifySP(data info.ModifySP) error {
	return sim.Attr.ModifySP(data)
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

func (sim *Simulation) MaxStance(target key.TargetID) float64 {
	return sim.Attr.MaxStance(target)
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

func (sim *Simulation) TurnOrder() []key.TargetID {
	return sim.Turn.TurnOrder()
}

func (sim *Simulation) GetActiveTarget() key.TargetID {
	return sim.Turn.GetActiveTarget()
}

func (sim *Simulation) SetGauge(data info.ModifyAttribute) error {
	sim.ActionTargets[data.Target] = true
	return sim.Turn.SetGauge(data)
}

func (sim *Simulation) ModifyGaugeNormalized(data info.ModifyAttribute) error {
	sim.ActionTargets[data.Target] = true
	return sim.Turn.ModifyGaugeNormalized(data)
}

func (sim *Simulation) ModifyGaugeAV(data info.ModifyAttribute) error {
	sim.ActionTargets[data.Target] = true
	return sim.Turn.ModifyGaugeAV(data)
}

func (sim *Simulation) SetCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
	sim.Turn.SetCurrentGaugeCost(data)
}

func (sim *Simulation) ModifyCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
	sim.Turn.ModifyCurrentGaugeCost(data)
}
