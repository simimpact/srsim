package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Attack key.Attack = "jingliu-attack"
)

var attackRatio = []float64{0.3, 0.7}

// Attack should be banned when enter enhanced mode
// TODO: create issue later
func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range attackRatio {
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
	state.EndAttack()
}
