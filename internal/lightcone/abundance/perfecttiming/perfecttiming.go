package perfecttiming

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Increases the wearer's Effect RES by 16% and increases Outgoing Healing by an amount
// that is equal to 33% of Effect RES. Outgoing Healing can be increased this way by up to 15%.

const (
	PTEffRes    key.Modifier = "perfect_timing_eff_res"
	PTHealBoost key.Modifier = "perfect_timing_heal_boost"
)

func init() {
	lightcone.Register(key.PerfectTiming, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	modifier.Register(PTEffRes, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: giveHealBuff,
		},
	})
	modifier.Register(PTHealBoost, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	effResAmt := 0.12 + 0.04*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   PTEffRes,
		Source: owner,
		Stats:  info.PropMap{prop.EffectRES: effResAmt},
		State:  float64(lc.Imposition), // 2 calcs outside depends on lc imp. level.
	})
}

func giveHealBuff(mod *modifier.Instance) {
	// take user's eff res value post-gear buff.
	currEffRes := mod.Engine().Stats(mod.Owner()).EffectRES()
	// out. heal buff depend on eff res post-buff. can't take in base stats.
	healBuffAmt := currEffRes * (0.30 + 0.03*mod.State().(float64))
	maxHealBuffAmt := 0.12 + 0.03*mod.State().(float64)
	if healBuffAmt > maxHealBuffAmt {
		healBuffAmt = maxHealBuffAmt
	}
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   PTHealBoost,
		Source: mod.Owner(),
		Stats:  info.PropMap{prop.HealBoost: healBuffAmt},
	})
}
