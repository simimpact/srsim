package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	UltPrimary  key.Attack = "danhengimbibitorlunae-ult-primary"
	UltAdjacent key.Attack = "danhengimbibitorlunae-ult-adjacent"
	E2          key.Reason = "danhengimbibitorlunae-e2"
)

var ultHits = []float64{0.3, 0.3, 0.4}

func init() {
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range ultHits {
		c.engine.Attack(info.Attack{
			Key:          UltPrimary,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_IMAGINARY,
			AttackType:   model.AttackType_ULT,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()]},
			StanceDamage: 60,
			EnergyGain:   5,
			HitRatio:     hitRatio,
		})
		c.engine.Attack(info.Attack{
			Key:          UltAdjacent,
			HitIndex:     i,
			Source:       c.id,
			Targets:      c.engine.AdjacentTo(target),
			DamageType:   model.DamageType_IMAGINARY,
			AttackType:   model.AttackType_ULT,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()] * 7 / 15},
			StanceDamage: 60,
			HitRatio:     hitRatio,
		})
		c.AddTalent()
	}
	state.EndAttack()
	c.point += 2
	// if E2,advanced forward 100% and 1 more point after ult
	if c.info.Eidolon >= 2 {
		c.point++
		c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    E2,
			Source: c.id,
			Target: c.id,
			Amount: -1,
		})
	}
	if c.point > 3 {
		c.point = 3
	}
}
