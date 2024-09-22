package guard

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	guard  = "guard-of-wuthering-snow"
	heal   = "guard-of-wuthering-snow-heal"
	energy = "guard-of-wuthering-snow-energy"
)

// 2pc: Reduces DMG taken by 8%.
// 4pc: At the beginning of the turn, if the wearer's HP is equal to or less than 50%,
//      restores HP equal to 8% of their Max HP and regenerates 5 Energy.

func init() {
	relic.Register(key.GuardOfWutheringSnow, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.AllDamageReduce: 0.08},
			},
			{
				MinCount: 4,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   guard,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(guard, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: onPhase1,
		},
	})
}

func onPhase1(mod *modifier.Instance) {
	// if above 50% HP, bypass
	if mod.Engine().HPRatio(mod.Owner()) > 0.5 {
		return
	}
	mod.Engine().Heal(info.Heal{
		Key:      heal,
		Targets:  []key.TargetID{mod.Owner()},
		Source:   mod.Owner(),
		BaseHeal: info.HealMap{model.HealFormula_BY_TARGET_MAX_HP: 0.08},
	})
	mod.Engine().ModifyEnergy(info.ModifyAttribute{
		Key:    energy,
		Target: mod.Owner(),
		Source: mod.Owner(),
		Amount: 5.0,
	})
}
