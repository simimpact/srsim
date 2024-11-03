package hanya

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill    = "hanya-skill"
	Burden   = "hanya-burden"
	Sanction = "hanya-sanction"
)

type BurdenState struct {
	atkCount          int
	triggersRemaining int
}

func init() {
	modifier.Register(Burden, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeBeingAttacked: BurdenCallbackBuff,
			OnAfterBeingAttacked:  BurdenCallbackSP,
			OnBeforeDying:         BurdenAboutToDie,
		},
	})

	modifier.Register(Sanction, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Skill,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_PHYSICAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		StanceDamage: 60,
		EnergyGain:   30,
	})

	if c.engine.HPRatio(target) > 0 {
		c.engine.AddModifier(target, info.Modifier{
			Name:   Burden,
			Source: c.id,
			State: BurdenState{
				atkCount:          0,
				triggersRemaining: 2,
			},
		})
	}
}

func BurdenCallbackBuff(mod *modifier.Instance, e event.AttackStart) {
	if e.AttackType != model.AttackType_SKILL && e.AttackType != model.AttackType_NORMAL && e.AttackType != model.AttackType_ULT {
		return
	}
	if mod.Engine().IsCharacter(e.Attacker) {
		hanya, _ := mod.Engine().CharacterInfo(mod.Source())
		damageBuff := talent[hanya.TalentLevelIndex()]
		if hanya.Eidolon >= 6 {
			damageBuff += 0.1
		}
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     Sanction,
			Source:   mod.Source(),
			Duration: 2,
			Stats: info.PropMap{
				prop.AllDamagePercent: damageBuff,
			},
		})
	}
}

func BurdenCallbackSP(mod *modifier.Instance, e event.AttackEnd) {
	if e.AttackType != model.AttackType_SKILL && e.AttackType != model.AttackType_NORMAL && e.AttackType != model.AttackType_ULT {
		return
	}
	state := mod.State().(*BurdenState)
	state.atkCount += 1
	// It shouldn't ever actually be greater than 2
	if state.atkCount >= 2 {
		state.atkCount = 0
		state.triggersRemaining -= 1
		mod.Engine().ModifySP(info.ModifySP{
			Key:    Burden,
			Source: mod.Source(),
			Amount: 1,
		})
		mod.Engine().AddModifier(e.Attacker, info.Modifier{
			Name:     A2,
			Source:   mod.Source(),
			Duration: 1,
		})
		hanya, _ := mod.Engine().CharacterInfo(mod.Source())
		// A6
		if hanya.Traces["103"] {
			mod.Engine().ModifyEnergy(info.ModifyAttribute{
				Key:    A6,
				Source: mod.Source(),
				Target: mod.Source(),
				Amount: 2,
			})
		}
		if state.triggersRemaining < 1 {
			mod.RemoveSelf()
		}
	}
}

func BurdenAboutToDie(mod *modifier.Instance) {
	hanya, _ := mod.Engine().CharacterInfo(mod.Source())
	// A4
	if hanya.Traces["102"] && mod.State().(BurdenState).atkCount <= 1 {
		mod.Engine().ModifySP(info.ModifySP{
			Key:    A4,
			Source: mod.Source(),
			Amount: 1,
		})
	}
}
