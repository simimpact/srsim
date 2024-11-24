package whereaboutsshoulddreamsrest

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
	routed                 = "whereabouts-should-dreams-rest-routed"
	whereaboutsBreakEffect = "whereabouts-should-dreams-rest-break-effect"
)

// Increases the wearer's Break Effect by 60%.
// When the wearer deals Break DMG to an enemy target, inflicts Routed on the enemy, lasting for 2 turn(s).
// Targets afflicted with Routed receive 24% increased Break DMG from the wearer, and their SPD is lowered by 20%.
// Effects of the same type cannot be stacked.
func init() {
	lightcone.Register(key.WhereaboutsShouldDreamsRest, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})
	modifier.Register(routed, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_DEBUFF,
		CanDispel:  true,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitAll,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.5 + 0.1*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   whereaboutsBreakEffect,
		Source: owner,
		Stats:  info.PropMap{prop.BreakEffect: amt},
		State:  float64(lc.Imposition),
	})
}

func onBeforeHitAll(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType != model.AttackType_ELEMENT_DAMAGE {
		return
	}

	mod.Engine().AddModifier(e.Hit.Defender.ID(), info.Modifier{
		Name:     routed,
		Source:   mod.Owner(),
		Duration: 2,
		Stats:    info.PropMap{prop.SPDPercent: -0.2},
	})

	if mod.Engine().HasModifier(e.Hit.Defender.ID(), routed) {
		amt := 0.2 + 0.04*mod.State().(float64)
		e.Hit.Defender.AddProperty(routed, prop.AllDamageTaken, amt)
	}
}
