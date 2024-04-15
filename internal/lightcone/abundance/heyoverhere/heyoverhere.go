package heyoverhere

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

// Increases the wearer's Max HP by 8%.
// When the wearer uses their Skill, increases Outgoing Healing by 16%, lasting for 2 turn(s).
const (
	HeyOverHere     key.Modifier = "hey-over-here"
	HeyOverHereBuff key.Modifier = "hey-over-here-heal-buff"
)

func init() {
	lightcone.Register(key.HeyOverHere, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})

	modifier.Register(HeyOverHere, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: addHealModifier,
		},
	})

	modifier.Register(HeyOverHereBuff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	hpBuff := 0.07 + 0.01*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   HeyOverHere,
		Source: owner,
		Stats:  info.PropMap{prop.HPPercent: hpBuff},
		State:  0.13 + 0.03*float64(lc.Imposition), // healing bonus
	})
}

func addHealModifier(mod *modifier.Instance, e event.ActionStart) {
	amount := mod.State().(float64)

	if e.AttackType == model.AttackType_SKILL {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     HeyOverHereBuff,
			Source:   mod.Owner(),
			Stats:    info.PropMap{prop.HealBoost: amount},
			Duration: 2,
		})
	}
}
