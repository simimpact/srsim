package riverflowsinspring

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
	RiverFlowsinSpring     key.Modifier = "river_flows_in_spring"
	RiverFlowsinSpringBuff key.Modifier = "river_flows_in_spring_buff"
)

type Amts struct {
	spd float64
	dmg float64
}

// After entering battle, increases the wearer's SPD by 8/9/10/11/12% and DMG
// by 12/15/18/21/24%. When the wearer takes DMG, this effect will disappear.
// This effect will resume after the end of the wearer's next turn.
func init() {
	lightcone.Register(key.RiverFlowsinSpring, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(RiverFlowsinSpring, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase2:           onPhase2,
			OnAfterBeingHitAll: onAfterBeingHitAll,
		},
	})

	modifier.Register(RiverFlowsinSpringBuff, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	spd_amt := 0.07 + 0.01*float64(lc.Imposition)
	dmg_amt := 0.09 + 0.03*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   RiverFlowsinSpring,
		Source: owner,
		State:  Amts{spd: spd_amt, dmg: dmg_amt},
	})

	engine.AddModifier(owner, info.Modifier{
		Name:   RiverFlowsinSpringBuff,
		Source: owner,
		Stats: info.PropMap{
			prop.SPDPercent:       spd_amt,
			prop.AllDamagePercent: dmg_amt,
		},
	})
}

func onPhase2(mod *modifier.ModifierInstance) {
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   RiverFlowsinSpringBuff,
		Source: mod.Owner(),
		Stats: info.PropMap{
			prop.SPDPercent:       mod.State().(Amts).spd,
			prop.AllDamagePercent: mod.State().(Amts).dmg,
		},
	})
}

func onAfterBeingHitAll(mod *modifier.ModifierInstance, e event.HitEndEvent) {
	if e.HPDamage > 0 {
		mod.Engine().RemoveModifier(mod.Owner(), RiverFlowsinSpringBuff)
	}
}
