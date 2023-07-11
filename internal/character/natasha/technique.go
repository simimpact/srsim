package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	technique                key.Modifier = "natasha-technique"
	techniqueScalePercentage float64      = 0.80
)

func init() {
	modifier.Register(technique, modifier.Config{
		StatusType: model.StatusType_STATUS_DEBUFF,
		Duration:   1,
		TickMoment: modifier.ModifierPhase1End,
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	targets := c.engine.Enemies()
	c.engine.Attack(info.Attack{
		Targets:    targets,
		Source:     c.id,
		AttackType: model.AttackType_MAZE,
		DamageType: model.DamageType_PHYSICAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: techniqueScalePercentage,
		},
	})
	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name:   technique,
			Source: c.id,
			Chance: 1,
			Stats:  info.PropMap{prop.AllDamageReduce: 0.30},
		})
	}
}
