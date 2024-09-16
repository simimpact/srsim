package guinaifen

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Technique = "guinaifen-technique"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	bounceCount := 4
	for i := 0; i < bounceCount; i++ {
		target := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Enemies(),
			Filter: func(target key.TargetID) bool {
				return c.engine.HPRatio(target) > 0
			},
			Max: 1,
		})[0]
		c.engine.Attack(info.Attack{
			Key:        Skill,
			Targets:    []key.TargetID{target},
			Source:     c.id,
			AttackType: model.AttackType_MAZE,
			DamageType: model.DamageType_FIRE,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: 0.5,
			},
		})

		if bounceCount == 4 {
			c.engine.EndAttack()
		}

		maxStacks := 3
		if c.info.Eidolon >= 6 {
			maxStacks = 4
		}

		// apply Firekiss
		c.engine.AddModifier(target, info.Modifier{
			Name:              Firekiss,
			Source:            c.id,
			Chance:            1,
			Duration:          3,
			MaxCount:          float64(maxStacks),
			CountAddWhenStack: 1,
		})
	}
}
