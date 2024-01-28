package aninstantbeforeagaze

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
	AnInstantBeforeAGaze              = "an-instant-before-a-gaze"
	ultBuff              key.Modifier = "an-instant-before-a-gaze-ult-buff"
)

// Increases the wearer's CRIT DMG by 36%.
// When the wearer uses Ultimate, increases the wearer's Ultimate DMG based on their Max Energy.
// Each point of Energy increases the Ultimate DMG by 0.36%, up to 180 points of Energy.
func init() {
	lightcone.Register(key.AnInstantBeforeAGaze, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
	modifier.Register(AnInstantBeforeAGaze, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: addUltBuff,
			OnAfterAction:  removeUltBuff,
		},
	})
	modifier.Register(ultBuff, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: onBeforeHit,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	maxenergy := engine.MaxEnergy(owner)
	if maxenergy > 180 {
		maxenergy = 180
	}
	amt := 0.30 + 0.06*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   AnInstantBeforeAGaze,
		Source: owner,
		Stats:  info.PropMap{prop.CritDMG: amt},
		State:  maxenergy * amt / 100,
	})
}

func addUltBuff(mod *modifier.Instance, e event.ActionStart) {
	if e.AttackType == model.AttackType_ULT {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   ultBuff,
			Source: mod.Owner(),
		})
	}
}

func onBeforeHit(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_ULT {
		state := mod.State().(float64)
		e.Hit.Attacker.AddProperty(AnInstantBeforeAGaze, prop.AllDamagePercent, state)
	}
}

func removeUltBuff(mod *modifier.Instance, e event.ActionEnd) {
	mod.Engine().RemoveModifier(mod.Owner(), ultBuff)
}
