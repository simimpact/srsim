package hook

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	//no in-game name for this, so idk what to call this
	SkillEnhancement            = "SkillEnhancement"
	Ult                         = "hook-ultimate"
	A6               key.Reason = "hook-a4"
)

var hitSplitRatio = []float64{0.3, 0.7}

func (c *char) Ult(target key.TargetID, state info.ActionState) {

	for hitIndex, hitRatio := range hitSplitRatio {
		c.engine.Attack(info.Attack{
			Key:        Ult,
			Source:     c.id,
			HitIndex:   hitIndex,
			HitRatio:   hitRatio,
			Targets:    []key.TargetID{target},
			AttackType: model.AttackType_ULT,
			DamageType: model.DamageType_FIRE,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: ultimate[c.info.UltLevelIndex()],
			},
			EnergyGain:   5,
			StanceDamage: 90,
		})
	}

	//A6
	if c.info.Traces["103"] {
		c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    A6,
			Target: c.id,
			Source: c.id,
			Amount: -0.2,
		})

		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    A6,
			Target: c.id,
			Source: c.id,
			Amount: 5,
		})
	}

	c.engine.AddModifier(c.id, info.Modifier{
		Name:   SkillEnhancement,
		Source: c.id,
	})

	c.engine.EndAttack()
}
