package defense

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	mod key.Modifier = "defense-lc" // incase defense keyword is too universal
)

// When the wearer unleashes their Ultimate, they restore HP by
// 18%/21%/24%/27%/30% of their Max HP.
func init() {
	lightcone.Register(key.Defense, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: onBeforeAction,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	healAmt := 0.15 + 0.03*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		State:  healAmt,
	})
}

func onBeforeAction(mod *modifier.Instance, e event.ActionStartEvent) {
	healAmt := mod.State().(float64)

	if e.AttackType == model.AttackType_ULT {
		mod.Engine().Heal(info.Heal{
			Targets:  []key.TargetID{mod.Owner()},
			Source:   mod.Owner(),
			BaseHeal: info.HealMap{model.HealFormula_BY_HEALER_MAX_HP: healAmt},
		})
	}
}
