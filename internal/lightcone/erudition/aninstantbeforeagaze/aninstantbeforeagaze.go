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
	AnInstantBeforeAGaze = "an-instant-before-a-gaze"
)

func init() {
	lightcone.Register(key.AnInstantBeforeAGaze, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})

	modifier.Register(AnInstantBeforeAGaze, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: onBeforeHit,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.30 + 0.06*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   AnInstantBeforeAGaze,
		Source: owner,
		Stats:  info.PropMap{prop.CritDMG: amt},
		State:  float64(lc.Imposition),
	})
}

func onBeforeHit(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType != model.AttackType_ULT {
		return
	}

	maxenergy := e.Hit.Attacker.MaxEnergy()
	if maxenergy > 180 {
		maxenergy = 180
	}

	dmgAmt := 0.36 * 0.06 * mod.State().(float64)
	e.Hit.Attacker.AddProperty(AnInstantBeforeAGaze, prop.AllDamagePercent, dmgAmt*maxenergy)
}
