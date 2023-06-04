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
		Params: map[string]float64{
			amt: 0.24 + 0.04*float64(lc.Ascension),
		},
	})
}

func onAdd(mod *modifier.ModifierInstance) {
	mod.AddProperty(prop.AggroPercent, 2.0)
	mod.AddProperty(prop.DEFPercent, mod.Params()[amt])
	mod.AddProperty(prop.EffectHitRate, mod.Params()[amt])
}

func onPhase2(mod *modifier.ModifierInstance) {
	mod.SetProperty(prop.DEFPercent, mod.Params()[amt])
}

func onAfterBeingAttacked(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	mod.SetProperty(prop.DEFPercent, 2.0*mod.Params()[amt])
}
