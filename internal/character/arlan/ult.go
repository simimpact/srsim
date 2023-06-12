package arlan

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var ultHits = []float64{0.3, 0.1, 0.6}

func (c *char) Ult(target key.TargetID, state info.ActionState) {

	c.e2()

	for _, hitRatio := range ultHits {

		// Primary Target
		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_THUNDER,
			AttackType: model.AttackType_ULT,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: ultDMG[c.info.AbilityLevel.Ult],
			},
			StanceDamage: 60.0,
			EnergyGain:   5.0,
			HitRatio:     hitRatio,
		})

		// Adjacent Targets
		additionalMod := 0.5

		if c.info.Eidolon >= 6 {
			additionalMod = 1
		}

		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    c.engine.AdjacentTo(target),
			DamageType: model.DamageType_THUNDER,
			AttackType: model.AttackType_ULT,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: additionalMod * ultDMG[c.info.AbilityLevel.Ult],
			},
			StanceDamage: 60.0,
			EnergyGain:   0.0,
			HitRatio:     hitRatio,
		})
	}

	state.EndAttack()
}
