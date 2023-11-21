package huohuo

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Attack key.Attack = "huohuo-attack"
)

var attackHits = []float64{0.2, 0.2, 0.2, 0.4}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range attackHits {
		c.engine.Attack(info.Attack{
			Key:          Attack,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_WIND,
			AttackType:   model.AttackType_NORMAL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_MAX_HP: atk[c.info.AttackLevelIndex()]},
			StanceDamage: 30,
			EnergyGain:   20,
			HitRatio:     hitRatio,
		})
	}
	state.EndAttack()
}
