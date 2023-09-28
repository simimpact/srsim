package yanqing

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Attack key.Attack = "yanqing-attack"
)

var attackHits = []float64{0.5, 0.25, 0.25}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range attackHits {
		c.engine.Attack(info.Attack{
			Key:          Attack,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_ICE,
			AttackType:   model.AttackType_NORMAL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()]},
			StanceDamage: 30,
			EnergyGain:   20,
			HitRatio:     hitRatio,
		})
	}
	c.tryFollow(target)
}
