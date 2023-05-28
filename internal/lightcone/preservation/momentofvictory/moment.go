package momentofvictory

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
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
			OnAdd: onAdd,
			// TODO: add OnPhase2 & OnAfterBeingAttacked listeners once implemented
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
	mod.AddProperty(model.Property_AGGRO_PERCENT, 2.0)
	mod.AddProperty(model.Property_DEF_PERCENT, mod.Params()[amt])
	mod.AddProperty(model.Property_EFFECT_HIT_RATE, mod.Params()[amt])
}

func onPhase2(mod *modifier.ModifierInstance) {
	mod.SetProperty(model.Property_DEF_PERCENT, mod.Params()[amt])
}

func onAfterBeingAttacked(mod *modifier.ModifierInstance) {
	mod.SetProperty(model.Property_DEF_PERCENT, 2.0*mod.Params()[amt])
}
