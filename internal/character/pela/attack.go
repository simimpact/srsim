package pela

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var attackHits = []float64{0.5, 0.5}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for _, hitRatio := range attackHits {
		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_ICE,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
			},
			StanceDamage: 30.0,
			EnergyGain:   20.0,
			HitRatio:     hitRatio,
		})
	}

	state.EndAttack()
}
