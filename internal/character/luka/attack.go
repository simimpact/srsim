package luka

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Normal         key.Attack = "luka-normal"
	EnhancedNormal key.Attack = "luka-enhanced-normal"
	DirectHit      key.Attack = "luka-direct-hit"
	RisingUppercut key.Attack = "luka-rising-uppercut"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	c.e1Check(target)
	if c.fightingSpirit < 2 {
		c.basicAttack(target, state)
	} else {
		c.enhancedBasic(target, state)
	}
}

var hitIndices = []float64{0.4, 0.3, 0.3}

func (c *char) enhancedBasic(target key.TargetID, state info.ActionState) {
	c.fightingSpirit -= 2

	punchCount := 3
	extraPunchCount := 0
	for i := 0; i < punchCount; i++ {

		c.engine.Attack(info.Attack{
			Key:        DirectHit,
			Targets:    []key.TargetID{target},
			Source:     c.id,
			AttackType: model.AttackType_NORMAL,
			DamageType: model.DamageType_PHYSICAL,
			HitIndex:   i,
			HitRatio:   hitIndices[i],
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: enhancedBasicDirectPunch[c.info.AttackLevelIndex()],
			},
			// The exact calc they use in the dm
			StanceDamage: 60 * 0.5,
		})
		if c.info.Traces["103"] && c.engine.Rand().Float64() > 0.5 {
			extraPunchCount += 1
			c.engine.Attack(info.Attack{
				Key:        DirectHit,
				Targets:    []key.TargetID{target},
				Source:     c.id,
				AttackType: model.AttackType_NORMAL,
				DamageType: model.DamageType_PHYSICAL,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: enhancedBasicDirectPunch[c.info.AttackLevelIndex()],
				},
			})
		}
	}

	c.risingUppercut(target)

	for _, dot := range c.engine.GetModifersByBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_BLEED) {
		dot.State.(common.TriggerableDot).TriggerDot(dot, talentRatio[c.info.TalentLevelIndex()], c.engine, target)
	}

	if c.info.Eidolon >= 6 {
		for _, dot := range c.engine.GetModifersByBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_BLEED) {
			for k := 0; k < (punchCount + extraPunchCount); k++ {
				dot.State.(common.TriggerableDot).TriggerDot(dot, 0.08, c.engine, target)
			}
		}
	}

	state.EndAttack()
}

func (c *char) risingUppercut(target key.TargetID) {
	c.engine.Attack(
		info.Attack{
			Key:        RisingUppercut,
			Targets:    []key.TargetID{target},
			Source:     c.id,
			DamageType: model.DamageType_PHYSICAL,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: enhancedBasicRisingUppercut[c.info.AttackLevelIndex()],
			},
			StanceDamage: 60 * 0.5,
		},
	)
}

func (c *char) basicAttack(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Normal,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_PHYSICAL,
		Source:     c.id,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30,
		EnergyGain:   20,
	})

	state.EndAttack()
}
