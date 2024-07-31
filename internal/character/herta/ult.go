package herta

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult = "herta-ult"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	if c.info.Traces["103"] {
		frozenEnemies := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Enemies(),
			Filter: func(target key.TargetID) bool {
				return c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_CTRL_FROZEN)
			},
		})

		for _, enemy := range frozenEnemies {
			c.engine.AddModifier(enemy, info.Modifier{
				Name:   a6,
				Source: c.id,
			})
		}
	}

	c.engine.Attack(info.Attack{
		Key:        Ult,
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_ICE,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		StanceDamage: 60,
		EnergyGain:   5,
	})

	if c.info.Eidolon >= 6 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   e6,
			Source: c.id,
			Stats: info.PropMap{
				prop.ATKPercent: 0.25,
			},
			Duration:        1,
			TickImmediately: true,
		})
	}

	c.engine.RemoveModifier(c.id, skillHPCheck)

	c.engine.EndAttack()
}
