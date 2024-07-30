package firesmith

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	check = "firesmith-of-lava-forging"
	buff  = "firesmith-of-lava-forging-buff"
)

// 2pc: Increases Fire DMG by 10%.
// 4pc: Increases the wearer's Skill DMG by 12%.
//      After unleashing Ultimate, increases the wearer's Fire DMG by 12% for the next attack.

func init() {
	relic.Register(key.FiresmithOfLavaForging, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.FireDamagePercent: 0.1},
			},
			{
				MinCount: 4,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   check,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitAll,
			OnAfterAction:  onAfterUltimate,
		},
	})
	modifier.Register(buff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAfterAttack: removeBuff,
		},
	})
}

func onBeforeHitAll(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_SKILL {
		e.Hit.Attacker.AddProperty(check, prop.AllDamagePercent, 0.12)
	}
}

func onAfterUltimate(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType == model.AttackType_ULT {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   buff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.FireDamagePercent: 0.12},
		})
	}
}

func removeBuff(mod *modifier.Instance, e event.AttackEnd) {
	mod.RemoveSelf()
}
