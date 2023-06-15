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
	CruisingintheStellarSeaCR       key.Modifier = "cruising_in_the_stellar_sea_cr"
	CruisingintheStellarSeaATKCheck key.Modifier = "cruising_in_the_stellar_sea_atk_check"
	CruisingintheStellarSeaATKBuff  key.Modifier = "cruising_in_the_stellar_sea_atk_buff"
)

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

	modifier.Register(CruisingintheStellarSeaCR, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitAll,
		},
	})

	modifier.Register(CruisingintheStellarSeaATKCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: onTriggerDeath,
		},
	})

	modifier.Register(CruisingintheStellarSeaATKBuff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	cr_amt := 0.06 + 0.02*float64(lc.Imposition)
	atk_amt := 0.15 + 0.05*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   CruisingintheStellarSeaCR,
		Source: owner,
		Stats:  info.PropMap{prop.CritChance: cr_amt},
		State:  cr_amt,
	})

	engine.AddModifier(owner, info.Modifier{
		Name:   CruisingintheStellarSeaATKCheck,
		Source: owner,
		State:  atk_amt,
	})
}

func onBeforeHitAll(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	if e.Hit.Defender.CurrentHPRatio() <= 0.5 {
		e.Hit.Attacker.AddProperty(prop.CritChance, mod.State().(float64))
	}
}

func onTriggerDeath(mod *modifier.ModifierInstance, target key.TargetID) {
	amt := mod.State().(float64)

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:     CruisingintheStellarSeaATKBuff,
		Source:   mod.Owner(),
		Duration: 2,
		Stats:    info.PropMap{prop.ATKPercent: amt},
	})
}
