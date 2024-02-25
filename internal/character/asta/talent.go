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
	asta_talent            = "asta-talent"
	asta_charging_stacks   = "asta-charging-stacks"
	asta_teammate_charging = "asta-charging-teammate"
)

func init() {
	modifier.Register(asta_talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1:      phase1TalentListener,
			OnAfterAttack: talentAttackListener,
		},
	})

	modifier.Register(asta_charging_stacks, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		MaxCount:   5,
		TickMoment: modifier.ModifierPhase1End,
		Listeners: modifier.Listeners{
			OnRemove: talentRemoveOnDeath,
			OnAdd:    addedStacksListener,
		},
	})

	modifier.Register(asta_teammate_charging, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   asta_talent,
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
		newStackCount := mod.Engine().ModifierStackCount(mod.Owner(), mod.Owner(), asta_charging_stacks) - float64(stacksToRemove)
		if newStackCount < 0 {
			newStackCount = 0
		}
		if newStackCount <= 2 {
			mod.Engine().RemoveModifier(mod.Owner(), e4)
		}
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   asta_charging_stacks,
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

	newStackCount := mod.Engine().ModifierStackCount(mod.Owner(), mod.Source(), asta_charging_stacks) + float64(stacksToAdd)
	if newStackCount > 5 {
		newStackCount = 5
	}

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   asta_charging_stacks,
		Source: mod.Source(),
		Count:  newStackCount,
	})

}

func talentRemoveOnDeath(mod *modifier.Instance) {
	for _, ally := range mod.Engine().Characters() {
		mod.Engine().RemoveModifier(ally, asta_teammate_charging)
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
			Name:   asta_teammate_charging,
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
