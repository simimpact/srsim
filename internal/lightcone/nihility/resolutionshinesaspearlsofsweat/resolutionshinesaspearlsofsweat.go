package resolutionshinesaspearlsofsweat

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
	sweat    key.Modifier = "resolution-shines-as-pearls-of-sweat"
	ensnared key.Modifier = "resolution-shines-as-pearls-of-sweat-ensnared"
)

type state struct {
	debuffChance, defDownAmt float64
}

// When the wearer hits an enemy and if the hit enemy is not already Ensnared,
// then there is a 60% base chance to Ensnare the hit enemy.
// Ensnared enemies' DEF decreases by 12% for 1 turn(s).

func init() {
	lightcone.Register(key.ResolutionShinesAsPearlsofSweat, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(sweat, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: applyEnsnared,
		},
	})
	modifier.Register(ensnared, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DEF_DOWN,
		},
		StatusType: model.StatusType_STATUS_DEBUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	debuffChance := 0.5 + 0.1*float64(lc.Imposition)
	defDownAmt := 0.11 + 0.01*float64(lc.Imposition)

	modState := state{
		debuffChance: debuffChance,
		defDownAmt:   defDownAmt,
	}

	engine.AddModifier(owner, info.Modifier{
		Name:   sweat,
		Source: owner,
		State:  &modState,
	})
}

func applyEnsnared(mod *modifier.Instance, e event.HitStart) {
	state := mod.State().(*state)
	// only apply ensnared if target not yet ensnared.
	if mod.Engine().HasModifier(e.Defender, ensnared) {
		return
	}
	mod.Engine().AddModifier(e.Defender, info.Modifier{
		Name:     ensnared,
		Source:   mod.Owner(),
		Stats:    info.PropMap{prop.DEFPercent: -state.defDownAmt},
		Chance:   state.debuffChance,
		Duration: 1,
	})
}
