package yanqing

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SoulsteelUlt            = "yanqing-soulsteelsync-ult"
	UltEffect               = "yanqing-ult-effect"
	Ult          key.Attack = "yanqing-ult"
)

func init() {
	modifier.Register(UltEffect, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.Replace,
		Duration:   1,
		Listeners: modifier.Listeners{
			OnTriggerDeath: E6Listener,
		},
	})
	modifier.Register(SoulsteelUlt, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.Replace,
		Duration:   1,
		Listeners: modifier.Listeners{
			OnAfterBeingHitAll: OnHitRemove,
			OnTriggerDeath:     E6Listener,
		},
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	if !c.engine.HasModifier(c.id, UltEffect) {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   UltEffect,
			Source: c.id,
			Stats:  info.PropMap{prop.CritChance: 0.6},
		})
	}
	if c.engine.HasModifier(c.id, Soulsteel) {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   SoulsteelUlt,
			Source: c.id,
			Stats:  info.PropMap{prop.CritDMG: ultCritDmg[c.info.UltLevelIndex()]},
		})
	}
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
