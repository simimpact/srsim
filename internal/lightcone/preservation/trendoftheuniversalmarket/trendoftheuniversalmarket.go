package trendoftheuniversalmarket

import (
	"github.com/simimpact/srsim/internal/global/common"
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
	mod key.Modifier = "trend-of-the-universal-market"
)

type State struct {
	dotChance float64
	dotDmg    float64
}

// Increases the wearer's DEF by 16%/20%/24%/28%/32%. When the wearer is
// attacked, there is a 100%/105%/110%/115%/120% base chance to Burn the enemy.
// For each turn, the wearer deals DoT that is equal to 40%/50%/60%/70%/80% of
// the wearer's DEF for 2 turn(s).
func init() {
	lightcone.Register(key.TrendoftheUniversalMarket, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterBeingAttacked: onAfterBeingAttacked,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// wearer DEF%
	amt := 0.12 + 0.04*float64(lc.Imposition)
	// wearer dot chance
	dotChance := 0.95 + 0.05*float64(lc.Imposition)
	dotDmg := 0.3 + 0.1*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.DEFPercent: amt},
		State:  State{dotChance, dotDmg},
	})

}

// chance to DoT the attacker
func onAfterBeingAttacked(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	state := mod.State().(State)

	mod.Engine().AddModifier(e.Attacker, info.Modifier{
		Name:   common.Burn,
		Source: mod.Owner(),
		State: common.BurnState{
			DEFDamagePercentage: state.dotDmg,
		},
		Chance:   state.dotChance,
		Duration: 2,
	})
}
