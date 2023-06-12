package danheng

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var attackHits = []float64{0.45, 0.55}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for _, hitRatio := range attackHits {
		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_WIND,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: atk[c.info.AbilityLevel.Attack],
			},
			StanceDamage: 30.0,
			EnergyGain:   20.0,
			HitRatio:     hitRatio,
		})
	}

	// end attack + attempt A4
	state.EndAttack()
	c.a4()
}
