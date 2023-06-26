package todayisanotherpeacefulday

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
	mod key.Modifier = "today-is-another-peaceful-day"
)

// After entering battle, increases the wearer's DMG based on their Max Energy.
// DMG increases by 0.2% per point of Energy, up to 160 Energy.
func init() {
	lightcone.Register(key.TodayIsAnotherPeacefulDay, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: onBeforeHit,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	maxenergy := engine.Stats(owner).MaxEnergy()
	if maxenergy > 160 {
		maxenergy = 160
	}
	amount := 0.15 + 0.05*float64(lc.Ascension)

	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.AllDamagePercent: amount * maxenergy},
		State:  amount * maxenergy,
	})
}

func onBeforeHit(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	amount := mod.State().(float64)
	e.Hit.Attacker.AddProperty(prop.AllDamagePercent, amount)
}
