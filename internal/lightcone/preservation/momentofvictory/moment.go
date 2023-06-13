package momentofvictory

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
	mod key.Modifier = "moment-of-victory"
	amt string       = "amount"
)

// Increases the wearer's DEF by 24% and Effect Hit Rate by 24%. Increases the chance for the
// wearer to be attacked by enemies. When the wearer is attacked, increase their DEF by
// an extra 24% until the end of the wearer's turn.
func init() {
	lightcone.Register(key.MomentOfVictory, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(mod, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnAdd:                onAdd,
			OnPhase2:             onPhase2,
			OnAfterBeingAttacked: onAfterBeingAttacked,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		State:  0.24 + 0.04*float64(lc.Rank),
	})
}

func onAdd(mod *modifier.ModifierInstance) {
	amount := mod.State().(float64)
	mod.AddProperty(prop.AggroPercent, 2.0)
	mod.AddProperty(prop.DEFPercent, amount)
	mod.AddProperty(prop.EffectHitRate, amount)
}

// reset back to 1x amount at end of turn
func onPhase2(mod *modifier.ModifierInstance) {
	mod.SetProperty(prop.DEFPercent, mod.State().(float64))
}

// after attack, double the DEF
func onAfterBeingAttacked(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	mod.SetProperty(prop.DEFPercent, 2.0*mod.State().(float64))
}
