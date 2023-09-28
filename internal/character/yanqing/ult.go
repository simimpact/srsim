package yanqing

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	UltEffect            = "yanqing-ult-effect"
	Ult       key.Attack = "yanqing-ult"
)

func init() {
	modifier.Register(UltEffect, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   1,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   UltEffect,
		Source: c.id,
		Stats:  info.PropMap{prop.CritChance: 0.6},
	})
	c.engine.Attack(info.Attack{
		Key:          Ult,
		Source:       c.id,
		Targets:      []key.TargetID{target},
		DamageType:   model.DamageType_ICE,
		AttackType:   model.AttackType_ULT,
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: ultRate[c.info.UltLevelIndex()]},
		StanceDamage: 90,
		EnergyGain:   5,
	})
	c.tryFollow(target)
}
