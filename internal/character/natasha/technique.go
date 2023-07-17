package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Technique                        = "natasha-technique"
	techniqueScalePercentage float64 = 0.80
)

func init() {
	modifier.Register(Technique, modifier.Config{
		StatusType:    model.StatusType_STATUS_DEBUFF,
		Duration:      1,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_FATIGUE},
		Stacking:      modifier.ReplaceBySource,
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	targets := c.engine.Enemies()

	c.engine.Attack(info.Attack{
		Key:        Technique,
		Targets:    c.engine.Retarget(info.Retarget{Targets: targets, Max: 1}),
		Source:     c.id,
		AttackType: model.AttackType_MAZE,
		DamageType: model.DamageType_PHYSICAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: techniqueScalePercentage,
		},
	})
	c.engine.EndAttack()

	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name:   Technique,
			Source: c.id,
			Chance: 1,
			Stats:  info.PropMap{prop.Fatigue: 0.30},
		})
	}
}
