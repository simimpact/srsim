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
)

var ultHits = []float64{0.3, 0.3, 0.4}

func init() {
	modifier.Register(Point, modifier.Config{
		StatusType:        model.StatusType_STATUS_BUFF,
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
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()*7/15]},
			StanceDamage: 60,
			HitRatio:     hitRatio,
		})
		c.AddTalent()
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Point,
		Source: c.id,
	})
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Point,
		Source: c.id,
	})
	if c.info.Eidolon >= 2 {
		c.engine.InsertAction(c.id)
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   Point,
			Source: c.id,
		})
	}
}
