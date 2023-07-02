package amber

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
	amber     key.Modifier = "amber"
	amberbuff key.Modifier = "amber-buff"
	amt       string       = "amount"
)

// Increases the wearer's DEF by 16%/20%/24%/28%/32%. If the wearer's current
// HP is lower than 50%, increases their DEF by a further
// 16%/20%/24%/28%/32%.
func init() {
	lightcone.Register(key.Amber, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})
	modifier.Register(amber, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: onLowerHalfHp,
			OnHPChange: func(mod *modifier.Instance, e event.HPChange) {
				onLowerHalfHp(mod)
			},
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.12 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   amber,
		Source: owner,
		Stats:  info.PropMap{prop.DEFPercent: amt},
		State:  amt,
	})
}

// DEF increases by another 16%/20%/24%/28%/32%
func onLowerHalfHp(mod *modifier.Instance) {
	if mod.Engine().HPRatio(mod.Owner()) < 0.5 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   amberbuff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.DEFPercent: mod.State().(float64)},
		})
	} else {
		mod.Engine().RemoveModifier(mod.Owner(), amberbuff)
	}
}
