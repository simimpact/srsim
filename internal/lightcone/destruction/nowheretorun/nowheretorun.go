package nowheretorun

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	NowheretoRun key.Modifier = "nowhere-to-run"
)

// Increases the wearer's ATK by 24%/30%/36%/42%/48%.
// Whenever the wearer defeats an enemy, they restore HP equal to 12%/15%/18%/21%/24% of their ATK.

func init() {
	lightcone.Register(key.NowheretoRun, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(NowheretoRun, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: onTriggerDeath,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amtATKPercent := 0.18 + 0.06*float64(lc.Imposition)
	amtHeal := 0.09 + 0.03*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   NowheretoRun,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: amtATKPercent},
		State:  amtHeal,
	})
}

func onTriggerDeath(mod *modifier.ModifierInstance, target key.TargetID) {
	mod.Engine().Heal(info.Heal{
		Targets:  []key.TargetID{mod.Owner()},
		Source:   mod.Owner(),
		BaseHeal: info.HealMap{model.HealFormula_BY_HEALER_ATK: mod.State().(float64)},
	})
}
