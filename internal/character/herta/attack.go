package herta

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Normal key.Attack = "herta-normal"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	canDoE1 := false
	if c.info.Eidolon >= 1 {
		if c.engine.HPRatio(target) <= 0.5 {
			canDoE1 = true
		}
	}

	c.engine.Attack(info.Attack{
		Key:          Normal,
		Targets:      []key.TargetID{target},
		Source:       c.id,
		AttackType:   model.AttackType_NORMAL,
		DamageType:   model.DamageType_ICE,
		EnergyGain:   20,
		StanceDamage: 30,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
		},
	})

	if canDoE1 {
		c.engine.Attack(info.Attack{
			Key:        e1,
			Targets:    []key.TargetID{target},
			Source:     c.id,
			DamageType: model.DamageType_ICE,
			AttackType: model.AttackType_PURSUED,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: 0.4,
			},
		})
	}

	c.engine.EndAttack()
}
