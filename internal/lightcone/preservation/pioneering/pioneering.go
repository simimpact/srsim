package pioneering

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	mod key.Modifier = "pioneering"
)

// When the wearer Breaks an enemy's Weakness, the wearer restores HP by
// 12%/14%/16%/18%/20% of their Max HP.
func init() {
	lightcone.Register(key.Pioneering, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerBreak: onTriggerBreak,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	healAmt := 0.1 + 0.02*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		State:  healAmt,
	})
}

func onTriggerBreak(mod *modifier.Instance, target key.TargetID) {
	healAmt := mod.State().(float64)
	mod.Engine().Heal(info.Heal{
		Targets:  []key.TargetID{mod.Owner()},
		Source:   mod.Owner(),
		BaseHeal: info.HealMap{model.HealFormula_BY_TARGET_MAX_HP: healAmt},
	})
}
