package subscribeformore

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
	SubscribeforMore key.Modifier = "subscribe_for_more"
)

// Increases the DMG of the wearer's Basic ATK and Skill by 24/30/36/42/48%.
// This effect increases by an extra 24/30/36/42/48% when the wearer's current
// Energy reaches its max level.
func init() {
	lightcone.Register(key.SubscribeforMore, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(SubscribeforMore, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: onBeforeHit,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   SubscribeforMore,
		Source: owner,
		State:  0.18 + .06*float64(lc.Imposition),
	})
}

func onBeforeHit(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_NORMAL || e.Hit.AttackType == model.AttackType_SKILL {
		amt := mod.State().(float64)
		if e.Hit.Attacker.Energy() == e.Hit.Attacker.MaxEnergy() {
			amt *= 2
		}
		e.Hit.Attacker.AddProperty(prop.AllDamagePercent, amt)
	}
}
