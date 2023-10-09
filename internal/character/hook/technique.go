package hook

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Technique = "hook-technique"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	techniqueTarget := c.engine.Retarget(info.Retarget{
		Max:          1,
		IncludeLimbo: false,
		Targets:      c.engine.Enemies(),
	})

	c.engine.Attack(info.Attack{
		Key:        Technique,
		Targets:    techniqueTarget,
		Source:     c.id,
		AttackType: model.AttackType_MAZE,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.5,
		},
	})

	c.talentProc(techniqueTarget[0])

	for _, t := range c.engine.Enemies() {
		c.engine.AddModifier(t, info.Modifier{
			Name:   common.Burn,
			Source: c.id,
			State: &common.BurnState{
				DamagePercentage: 0.5,
				DamageValue:      0,
			},
			Chance:   1,
			Duration: 3,
		})
	}

}
