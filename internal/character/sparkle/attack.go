package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SparkleBasic key.Attack = "sparkle-basic"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	// A2
	if c.info.Traces["101"] {
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    A2,
			Amount: 10,
			Target: c.id,
			Source: c.id,
		})
	}

	c.engine.Attack(info.Attack{
		Key:        SparkleBasic,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_QUANTUM,
		AttackType: model.AttackType_NORMAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
		},
		EnergyGain:   20,
		StanceDamage: 30,
	})

	state.EndAttack()
}
