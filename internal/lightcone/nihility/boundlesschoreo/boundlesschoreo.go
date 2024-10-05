package boundlesschoreo

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
	Choreo = "boundless-choreo"
	Check  = "boundless-choreo-cdmg"
)

// Increase the wearer's CRIT Rate by 8/10/12/14/16%.
// The wearer deals 24/30/36/42/48% more CRIT DMG to enemies that are currently Slowed or have reduced DEF.

func init() {
	lightcone.Register(key.BoundlessChoreo, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: applyCdmg,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	crAmt := 0.06 + 0.02*float64(lc.Imposition)
	cdmgAmt := 0.18 + 0.06*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		Stats:  info.PropMap{prop.CritChance: crAmt},
		State:  cdmgAmt,
	})
}

func applyCdmg(mod *modifier.Instance, e event.HitStart) {
	// bypass if any check fails
	if !mod.Engine().HasBehaviorFlag(e.Defender, model.BehaviorFlag_STAT_SPEED_DOWN) {
		return
	}
	if !mod.Engine().HasBehaviorFlag(e.Defender, model.BehaviorFlag_STAT_DEF_DOWN) {
		return
	}
	e.Hit.Attacker.AddProperty(Check, prop.CritDMG, mod.State().(float64))
}
