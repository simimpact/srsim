package serval

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Normal key.Attack = "serval-normal"

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Normal,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_THUNDER,
		AttackType: model.AttackType_NORMAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30.0,
		EnergyGain:   20.0,
	})
	if c.info.Eidolon >= 1 {
		targets := c.engine.AdjacentTo(target)
		random_index := c.engine.Rand().Intn(len(targets))
		c.engine.Attack(info.Attack{
			Key:        SkillAdjacent,
			Source:     c.id,
			Targets:    []key.TargetID{targets[random_index]},
			DamageType: model.DamageType_THUNDER,
			AttackType: model.AttackType_PURSUED,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: 0.6,
			},
			StanceDamage: 30.0,
			EnergyGain:   0.0,
		})
	}
	state.EndAttack()
}
