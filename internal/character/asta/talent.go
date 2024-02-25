package asta

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	astaTalent           = "asta-talent"
	astaChargingStacks   = "asta-charging-stacks"
	astaTeammateCharging = "asta-charging-teammate"
)

// Scuffed
func init() {
	modifier.Register(astaTalent, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1:      phase1TalentListener,
			OnAfterAttack: talentAttackListener,
		},
	})

	modifier.Register(astaChargingStacks, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		MaxCount:   5,
		TickMoment: modifier.ModifierPhase1End,
		Listeners: modifier.Listeners{
			OnRemove: talentRemoveOnDeath,
			OnAdd:    addedStacksListener,
		},
	})

	modifier.Register(astaTeammateCharging, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   astaTalent,
		Source: c.id,
	})
}

func phase1TalentListener(mod *modifier.Instance) {
	asta, _ := mod.Engine().CharacterInfo(mod.Owner())
	stacksToRemove := 2
	if asta.Eidolon >= 6 {
		stacksToRemove = 1
	}

	if !mod.Engine().HasModifier(mod.Owner(), e2) {
		newStackCount := mod.Engine().ModifierStackCount(mod.Owner(), mod.Owner(), astaChargingStacks) - float64(stacksToRemove)
		if newStackCount < 0 {
			newStackCount = 0
		}
		if newStackCount <= 2 {
			mod.Engine().RemoveModifier(mod.Owner(), e4)
		}
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   astaChargingStacks,
			Source: mod.Owner(),
			Count:  newStackCount,
		},
		)
	}
}

func talentAttackListener(mod *modifier.Instance, e event.AttackEnd) {
	targeted := mod.Engine().Retarget(info.Retarget{
		Targets: e.Targets,
		Filter: func(target key.TargetID) bool {
			return true
		},
		IncludeLimbo: true,
	})

	stacksToAdd := 1
	for _, t := range targeted {
		if mod.Engine().Stats(t).IsWeakTo(model.DamageType_FIRE) {
			stacksToAdd = 2
			break
		}
	}

	newStackCount := mod.Engine().ModifierStackCount(mod.Owner(), mod.Source(), astaChargingStacks) + float64(stacksToAdd)
	if newStackCount > 5 {
		newStackCount = 5
	}

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   astaChargingStacks,
		Source: mod.Source(),
		Count:  newStackCount,
	})
}

func talentRemoveOnDeath(mod *modifier.Instance) {
	for _, ally := range mod.Engine().Characters() {
		mod.Engine().RemoveModifier(ally, astaTeammateCharging)
	}
	asta, _ := mod.Engine().CharacterInfo(mod.Owner())
	if asta.Eidolon >= 4 {
		mod.Engine().RemoveModifier(mod.Owner(), e4)
	}
}

func addedStacksListener(mod *modifier.Instance) {
	a6DefRatio := 0.0
	asta, _ := mod.Engine().CharacterInfo(mod.Owner())
	if asta.Traces["103"] {
		a6DefRatio = 0.06
	}

	mod.AddProperty(prop.ATKPercent, talent[asta.TalentLevelIndex()]*mod.Count())
	mod.AddProperty(prop.DEFPercent, a6DefRatio*mod.Count())

	for _, ally := range mod.Engine().Characters() {
		mod.Engine().AddModifier(ally, info.Modifier{
			Name:   astaTeammateCharging,
			Source: mod.Owner(),
			Stats: info.PropMap{
				prop.ATKPercent: talent[asta.TalentLevelIndex()] * mod.Count(),
			},
		})
	}

	if asta.Eidolon >= 4 {
		if mod.Count() >= 2 {
			mod.Engine().AddModifier(mod.Owner(), info.Modifier{
				Name:   e4,
				Source: mod.Owner(),
				Stats: info.PropMap{
					prop.EnergyRegen: 0.15,
				},
			})
		}
	}
}
