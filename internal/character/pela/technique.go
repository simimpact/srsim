package pela

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TechniqueDefDown key.Modifier = "pela-technique-def-down"
)

func init() {
	modifier.Register(TechniqueDefDown, modifier.Config{
		TickMoment:    modifier.ModifierPhase1End,
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_DEF_DOWN},
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	targets := c.engine.Enemies()
	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{targets[c.engine.Rand().Intn(len(targets))]},
		DamageType: model.DamageType_ICE,
		AttackType: model.AttackType_MAZE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.8,
		},
		StanceDamage: 60.0,
		EnergyGain:   0.0,
	})

	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name:     TechniqueDefDown,
			Source:   c.id,
			Chance:   1,
			Duration: 2,
			Stats:    info.PropMap{prop.DEFPercent: -0.2},
		})
	}
}
