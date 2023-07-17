package fermata

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
	name = "fermata"
)

// Increases the Break Effect dealt by the wearer by 16%/20%/24%/28%/32%, and increases their DMG
// to enemies afflicted with Shock or Wind Shear by 16%/20%/24%/28%/32%. This also applies to DoT.
func init() {
	lightcone.Register(key.Fermata, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})

	modifier.Register(name, modifier.Config{
		CanModifySnapshot: true,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitAll,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.12 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   name,
		Source: owner,
		Stats:  info.PropMap{prop.BreakEffect: amt},
		State:  amt,
	})
}

var triggerFlags = []model.BehaviorFlag{
	model.BehaviorFlag_STAT_DOT_ELECTRIC,
	model.BehaviorFlag_STAT_DOT_POISON,
}

func onBeforeHitAll(mod *modifier.Instance, e event.HitStart) {
	amt := mod.State().(float64)

	if mod.Engine().HasBehaviorFlag(e.Defender, triggerFlags...) {
		e.Hit.Attacker.AddProperty(name, prop.AllDamagePercent, amt)
	}
}
