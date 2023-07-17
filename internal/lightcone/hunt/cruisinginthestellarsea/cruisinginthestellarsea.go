package cruisinginthestellarsea

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
	CruisingintheStellarSea                     = "cruising-in-the-stellar-sea"
	CruisingintheStellarSeaATKBuff key.Modifier = "cruising-in-the-stellar-sea-atk-buff"
)

type Amts struct {
	cr  float64
	atk float64
}

// Increases CRIT rate by 8/10/12/14/16%
// Increases CRIT rate against enemies w/ HP <= 50% by an extra 8/10/12/14/16%
// On enemy defeat, ATK is increased by 20/25/30/35/40% for 2 turn(s)
func init() {
	lightcone.Register(key.CruisingintheStellarSea, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(CruisingintheStellarSea, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitAll,
			OnTriggerDeath: onTriggerDeath,
		},
	})

	modifier.Register(CruisingintheStellarSeaATKBuff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	crAmt := 0.06 + 0.02*float64(lc.Imposition)
	atkAmt := 0.15 + 0.05*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   CruisingintheStellarSea,
		Source: owner,
		Stats:  info.PropMap{prop.CritChance: crAmt},
		State:  Amts{cr: crAmt, atk: atkAmt},
	})
}

func onBeforeHitAll(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.Defender.CurrentHPRatio() <= 0.5 {
		e.Hit.Attacker.AddProperty(CruisingintheStellarSea, prop.CritChance, mod.State().(Amts).cr)
	}
}

func onTriggerDeath(mod *modifier.Instance, target key.TargetID) {
	amt := mod.State().(Amts).atk

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:     CruisingintheStellarSeaATKBuff,
		Source:   mod.Owner(),
		Duration: 2,
		Stats:    info.PropMap{prop.ATKPercent: amt},
	})
}
