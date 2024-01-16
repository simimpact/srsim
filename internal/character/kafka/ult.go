package kafka

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Ult = "kafka-ult"

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Ult,
		Targets:    c.engine.Enemies(),
		Source:     c.id,
		AttackType: model.AttackType_ULT,
		DamageType: model.DamageType_THUNDER,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		EnergyGain:   5,
		StanceDamage: 60,
	})

	c.applyShock(c.engine.Enemies())

	dots := []key.Modifier{common.Shock, common.BreakShock}
	if c.info.Traces["101"] {
		dots = allTriggerableDots
	}

	// kinda ugly, but itll do
	for _, t := range c.engine.Enemies() {
		for _, triggerable := range dots {
			for _, dot := range c.engine.GetModifiers(t, triggerable) {
				dot.State.(common.TriggerableDot).TriggerDot(dot, ultDotDetonate[c.info.UltLevelIndex()], c.engine, t)
			}
		}
	}

	c.engine.EndAttack()
}

func (c *char) applyShock(targets []key.TargetID) {
	shockDur := 2
	e6AdditionalMultiplier := 0.0
	if c.info.Eidolon >= 6 {
		shockDur += 1
		e6AdditionalMultiplier += 1.56
	}

	shockChance := 1.0
	if c.info.Traces["103"] {
		shockChance += 0.3
	}

	for _, t := range targets {
		c.engine.AddModifier(t, info.Modifier{
			Name:     common.Shock,
			Source:   c.id,
			Chance:   shockChance,
			Duration: shockDur,
			State: common.ShockState{
				DamagePercentage: ultDotValue[c.info.UltLevelIndex()] + e6AdditionalMultiplier,
				DamageValue:      0,
			},
		})
	}
}
