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

type State struct {
	ultBonus float64
	idMap    map[key.TargetID]bool
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

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit:   onBeforeHit,
			OnAfterAction: onAfterAction,
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
		State:  State{ultBonus: ultDmg, idMap: make(map[key.TargetID]bool)},
	})
}

// increase ult damage
func onBeforeHit(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	state := mod.State().(State)
	_, hasId := state.idMap[e.Hit.Defender.ID()]

	if e.Hit.AttackType == model.AttackType_ULT && !hasId {
		state.idMap[e.Hit.Defender.ID()] = true
		e.Hit.DamageValue += state.ultBonus * e.Hit.Attacker.DEF()
	}
}

// remove modifier so next ult deals ult dmg + only 1x bonus from this lc
func onAfterAction(mod *modifier.ModifierInstance, e event.ActionEvent) {
	state := mod.State().(State)
	for k := range state.idMap {
		delete(state.idMap, k)
	}
}
