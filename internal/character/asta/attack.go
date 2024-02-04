package asta

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2     = "asta-a2"
	Normal = "asta-normal"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	if c.info.Traces["102"] {
		c.engine.AddModifier(target, info.Modifier{
			Name:   common.Burn,
			Source: c.id,
			State: &common.BurnState{
				DamagePercentage:    0.5,
				DamageValue:         0,
				DEFDamagePercentage: 0,
			},
			Chance:   0.8,
			Duration: 3,
		})
	}

	c.engine.Attack(info.Attack{
		Key:        Normal,
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
		},
		Targets:      []key.TargetID{target},
		Source:       c.id,
		EnergyGain:   20,
		StanceDamage: 30,
	})
}
