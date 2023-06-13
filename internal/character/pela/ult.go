package pela

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	UltDefShred key.Modifier = "pela-ult-def-shred"
)

func init() {
	modifier.Register(UltDefShred, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_DEF_DOWN},
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	targets := c.engine.Enemies()

	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name:     UltDefShred,
			Source:   c.id,
			Chance:   1,
			Duration: 2,
			Stats:    info.PropMap{prop.DEFPercent: -ultDefShred[c.info.AbilityLevel.Ult-1]},
		})
	}

	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    targets,
		DamageType: model.DamageType_ICE,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.AbilityLevel.Ult-1],
		},
		StanceDamage: 60.0,
		EnergyGain:   5,
	})

	state.EndAttack()
}
