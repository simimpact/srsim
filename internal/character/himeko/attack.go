package himeko

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Normal key.Attack = "himeko-normal"

var hitSplit = []float64{0.4, 0.6}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for hitIndex, split := range hitSplit {
		c.engine.Attack(info.Attack{
			Key:          Normal,
			HitIndex:     hitIndex,
			HitRatio:     split,
			Targets:      []key.TargetID{target},
			Source:       c.id,
			AttackType:   model.AttackType_NORMAL,
			DamageType:   model.DamageType_FIRE,
			EnergyGain:   20,
			StanceDamage: 30,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
			},
		})
	}

	c.engine.EndAttack()
}
