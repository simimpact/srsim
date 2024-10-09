package whatisreal

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
	WhatIsReal = "what-is-real"
)

func init() {
	lightcone.Register(key.WhatIsReal, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	modifier.Register(WhatIsReal, modifier.Config{
		Stacking: modifier.Unique,
		Listeners: modifier.Listeners{
			OnAfterAction: afterBasic,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.18 + 0.06*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   WhatIsReal,
		Source: owner,
		Stats:  info.PropMap{prop.BreakEffect: amt},
		State:  float64(lc.Imposition),
	})
}

func afterBasic(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType != model.AttackType_NORMAL {
		return
	}

	amt := 0.015 + 0.005*mod.State().(float64)
	mod.Engine().Heal(info.Heal{
		Key:     WhatIsReal,
		Targets: []key.TargetID{mod.Owner()},
		Source:  mod.Owner(),
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: amt,
		},
		HealValue: 800,
	})
}
