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
	mod    key.Modifier = "this-is-me"
	modUlt key.Modifier = "this-is-me-ult"
)

type State struct {
	ultBonus float64
}

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

	modifier.Register(modUlt, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: onBeforeHit,
			OnAfterHit:  onAfterHit,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.12 + 0.04*float64(lc.Imposition)
	ultDmg := 0.45 + 0.15*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.DEFPercent: amt},
		State:  State{ultDmg},
	})
}

// increase ult damage
func onBeforeHit(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	state := mod.State().(State)
	if e.Hit.AttackType == model.AttackType_ULT {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   modUlt,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.DEFPercent: state.ultBonus},
		})
	}
}

// remove modifier so next ult deals ult dmg + only 1x bonus from this lc
func onAfterHit(mod *modifier.ModifierInstance, e event.HitEndEvent) {
	mod.RemoveSelf()
}
