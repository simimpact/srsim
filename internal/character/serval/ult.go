package serval

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Ult key.Attack = "serval-ult"

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	if c.info.Eidolon >= 4 {
		for _, trg := range c.engine.Enemies() {
			c.engine.AddModifier(trg, info.Modifier{
				Name: common.Shock,
				State: &common.ShockState{
					DamagePercentage: skillDot[c.info.SkillLevelIndex()],
				},
				Source:   c.id,
				Chance:   1,
				Duration: 2,
			})
		}
	}
	c.engine.Attack(info.Attack{
		Key:        Ult,
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_THUNDER,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		StanceDamage: 60.0,
		EnergyGain:   5,
	})
	for _, trg := range c.engine.Enemies() {
		c.engine.ExtendModifierDuration(trg, common.Shock, 2)
	}
	state.EndAttack()
}
