package serval

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TechniqueAtk   key.Attack   = "serval-technique"
	TechniqueShock key.Modifier = "serval-technique-shock"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	targets := c.engine.Enemies()
	c.engine.Attack(info.Attack{
		Key:        TechniqueAtk,
		Source:     c.id,
		Targets:    []key.TargetID{targets[c.engine.Rand().Intn(len(targets))]},
		DamageType: model.DamageType_THUNDER,
		AttackType: model.AttackType_MAZE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.5,
		},
		StanceDamage: 60.0,
		EnergyGain:   0.0,
	})

	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name: common.Shock,
			State: &common.ShockState{
				DamagePercentage: 0.5,
				DamageValue:      0,
			},
			Source:   c.id,
			Chance:   1,
			Duration: 2,
		})
	}
}
