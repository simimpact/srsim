package thisisme

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
	mod key.Modifier = "this-is-me"
)

// Increases the wearer's DEF by 16%/20%/24%/28%/32%. Increases the DMG of the
// wearer when they use their Ultimate by 60%/75%/90%/105%/120% of the wearer's
// DEF. This effect only applies 1 time per enemy target during each use of the
// wearer's Ultimate.
func init() {
	lightcone.Register(key.ThisIsMe, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: func(mod *modifier.ModifierInstance, e event.HitStartEvent) {},
			OnAfterHit:  func(mod *modifier.ModifierInstance, e event.HitEndEvent) {},
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.12 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.DEFPercent: amt},
	})
}

func onBeforeHit(mod *modifier.ModifierInstance, e event.HitStartEvent) {
}

// remove modifier so next ult deals ult dmg + only 1x bonus from this lc
func onAfterHit(mod *modifier.ModifierInstance, e event.HitEndEvent) {
	mod.RemoveSelf()
}
