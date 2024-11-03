package xueyi

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent    = "xueyi-talent"
	TalentFua = "xueyi-talent-fua"
)

/*
*
Xueyi talent fua should be able to trigger on the following:

OnEnterBattle
OnListenAfterAttack
OnCustomEvent (?)
OnListenModifierRemove

OnEnterBattle, OnListenAfterAttack should both be covered by a subscription to stance change. If Xueyi ever gains stacks from entering battle it
will be from techniques, initial overworld attacsks which should emit stance changes, same for AfterAttack.
OnCustomEvent SHOULD also be covered by this, as this is presumably whenever karma stacks are added to.
OnListenModifierRemove should be covered by having talent loop fua triggers until it can't anymore.
*/
func (c *char) initTalent() {
	if c.info.Eidolon >= 6 {
		c.stackReq = 6
	}
	c.engine.Events().StanceChange.Subscribe(c.handleStanceChange)
}

func (c *char) handleStanceChange(e event.StanceChange) {
	if e.Key == TalentFua {
		return
	}
	if e.Source != c.id {
		if c.engine.IsCharacter(e.Source) {
			c.incrementTalentStacks(1)
		}
		return
	}
	diff := e.OldStance - e.NewStance
	increment := 1
	if diff > 30 {
		increment = int(diff / 30)
	}
	c.incrementTalentStacks(increment)
}

func (c *char) incrementTalentStacks(amt int) {
	if amt > 8 {
		amt = 8
	}
	c.curStacks += amt
	overflow := 0
	if c.info.Traces["103"] {
		overflow = 6
	}
	if c.curStacks > c.stackReq+overflow {
		c.curStacks = c.stackReq + overflow
	}

	// This should be correct behavior?
	for c.curStacks >= c.stackReq {
		c.curStacks -= c.stackReq

		// Activate talent fua
		c.engine.InsertAbility(info.Insert{
			Key:        TalentFua,
			Source:     c.id,
			AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
			Priority:   info.CharInsertAttackSelf,
			Execute:    c.talentFua,
		})
	}
}

func (c *char) talentFua() {
	if c.info.Eidolon >= 2 {
		c.engine.Heal(info.Heal{
			Source:  c.id,
			Targets: []key.TargetID{c.id},
			Key:     E2,
			BaseHeal: info.HealMap{
				model.HealFormula_BY_HEALER_MAX_HP: 0.05,
			},
		})
	}

	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E1,
			Source: c.id,
		})
	}
	toughness := 15.0
	if c.info.Eidolon >= 2 {
		toughness = 0.0
	}
	// Do 3 fuas
	for i := 0; i < 3; i++ {
		target := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Enemies(),
			Filter: func(target key.TargetID) bool {
				return c.engine.HPRatio(target) > 0
			},
			Max:          1,
			IncludeLimbo: true,
		})[0]
		c.engine.Attack(info.Attack{
			Key:        TalentFua,
			Source:     c.id,
			Targets:    []key.TargetID{target},
			AttackType: model.AttackType_INSERT,
			DamageType: model.DamageType_QUANTUM,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: talent[c.info.TalentLevelIndex()],
			},
			StanceDamage: toughness,
			EnergyGain:   2,
		})
		if c.info.Eidolon >= 2 {
			c.engine.ModifyStance(info.ModifyAttribute{
				Key:    TalentFua,
				Source: c.id,
				Target: target,
				Amount: 15.0,
			})
		}
	}
	c.engine.RemoveModifier(c.id, E2)
	c.engine.RemoveModifier(c.id, E1)
}
