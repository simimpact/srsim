package perfecttiming

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

// Increases the wearer's Effect RES by 16% and increases Outgoing Healing by an amount
// that is equal to 33% of Effect RES. Outgoing Healing can be increased this way by up to 15%.

const (
	PTEffRes    key.Modifier = "perfect-timing"
	PTHealBoost key.Modifier = "perfect-timing-heal-boost"
)

func init() {
	lightcone.Register(key.PerfectTiming, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	modifier.Register(PTEffRes, modifier.Config{})

	modifier.Register(PTHealBoost, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// OnAdd -> give eff. RES
	effResAmt := 0.12 + 0.04*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   PTEffRes,
		Source: owner,
		Stats:  info.PropMap{prop.EffectRES: effResAmt},
	})
	// OnBattleStart -> giveHealBoost to owner. track last healBoostAmt through func returns
	healBoostAmt := 0.0
	engine.Events().BattleStart.Subscribe(func(e event.BattleStart) {
		for _, char := range e.CharStats {
			if char.ID() == owner {
				healBoostAmt = giveHealBoost(engine, char, lc.Imposition, healBoostAmt)
			}
		}
	})
	// OnActionEnd -> recalc HealBoost.
	engine.Events().ActionEnd.Subscribe(func(e event.ActionEnd) {
		healBoostAmt = giveHealBoost(engine, engine.Stats(owner), lc.Imposition, healBoostAmt)
	})
}

func giveHealBoost(engine engine.Engine, owner *info.Stats, imposition int, prevBuff float64) float64 {
	// calc healBoostAmt from current effRES value
	currEffRes := engine.Stats(owner.ID()).EffectRES()
	healBoostAmt := currEffRes * (0.30 + 0.03*float64(imposition))
	// set max healboost amt
	maxhealBoostAmt := 0.12 + 0.03*float64(imposition)
	if healBoostAmt > maxhealBoostAmt {
		healBoostAmt = maxhealBoostAmt
	}
	// update modifier if healBoostAmt changes
	if healBoostAmt != prevBuff {
		engine.AddModifier(owner.ID(), info.Modifier{
			Name:   PTHealBoost,
			Source: owner.ID(),
			Stats:  info.PropMap{prop.HealBoostConvert: healBoostAmt},
		})
		return healBoostAmt
	}
	return prevBuff
}
