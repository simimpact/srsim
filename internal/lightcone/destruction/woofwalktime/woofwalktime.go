package woofwalktime

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	mod key.Modifier = "woof-walk-time"
)

// Increases the wearer's ATK by 10%, and increases their DMG to enemies
// afflicted with Burn or Bleed by 16%. This also applies to DoT.

// impl note : literally same as fermata. diff DoT check and initial buff

func init() {
	lightcone.Register(key.WoofWalkTime, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})
	modifier.Register(mod, modifier.Config{
		CanModifySnapshot: true,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: dmgBoostOnBurnBleed,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	atkAmt := 0.075 + 0.025*float64(lc.Imposition)
	dmgBoostAmt := 0.12 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: atkAmt},
		State:  dmgBoostAmt,
	})
}

var triggerFlags = []model.BehaviorFlag{
	model.BehaviorFlag_STAT_DOT_BURN,
	model.BehaviorFlag_STAT_DOT_BLEED,
}

func dmgBoostOnBurnBleed(mod *modifier.Instance, e event.HitStart) {
	dmgBoostAmt := mod.State().(float64)

	if mod.Engine().HasBehaviorFlag(e.Defender, triggerFlags...) {
		e.Hit.Attacker.AddProperty(prop.AllDamagePercent, dmgBoostAmt)
	}
}
