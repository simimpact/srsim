package kafka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Normal key.Attack = "kafka-normal"

var hitSplit = []float64{0.5, 0.5}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range hitSplit {
		c.engine.Attack(info.Attack{
			Key:        Normal,
			HitIndex:   i,
			Targets:    []key.TargetID{target},
			Source:     c.id,
			AttackType: model.AttackType_NORMAL,
			DamageType: model.DamageType_THUNDER,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
			},
			EnergyGain:   20,
			StanceDamage: 30,
			HitRatio:     hitRatio,
		})
	}

	c.engine.EndAttack()
}
