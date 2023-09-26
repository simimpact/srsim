package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Point            = "danhengimbibitorlunae-point"
	Ult   key.Attack = "danhengimbibitorlunae-ult"
	E2    key.Reason = "danhengimbibitorlunae-e2"
)

var ultHits = []float64{0.3, 0.3, 0.4}

func init() {
	modifier.Register(Point, modifier.Config{
		StatusType:        model.StatusType_UNKNOWN_STATUS,
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          3,
		CountAddWhenStack: 1,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range ultHits {
		c.engine.Attack(info.Attack{
			Key:          Ult,
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
			Key:          Ult,
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
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Point,
		Count:  2,
		Source: c.id,
	})
	// if E2,advanced forward 100% and 1 more point after ult
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   Point,
			Source: c.id,
		})
		c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    E2,
			Source: c.id,
			Target: c.id,
			Amount: -1,
		})
	}
}
