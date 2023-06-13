package mutualdemise

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
	MutualDemise key.Modifier = "mutual-demise"
)

func init() {
	lightcone.Register(key.MutualDemise, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(key.MutualDemise, modifier.Config{
		Listeners: modifier.Listeners{
			OnHPChange: onHPChange,
			OnAdd:      onAdd,
		},
	})
}

//If the wearer's HP is less than 80%, CRIT Rate increases by 12%/15%/18%/21%/24%

func adjustCritRate(mod *modifier.ModifierInstance) {
	if mod.Engine().HPRatio(mod.Owner()) < 0.8 {
		mod.SetProperty(prop.CritChance, mod.State().(float64))
	} else {
		mod.SetProperty(prop.CritChance, 0)
	}
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   MutualDemise,
		Source: owner,
		State:  0.09 + 0.03*float64(lc.Imposition),
	})
}

func onAdd(mod *modifier.ModifierInstance) {
	adjustCritRate(mod)
}

func onHPChange(mod *modifier.ModifierInstance, e event.HPChangeEvent) {
	adjustCritRate(mod)
}
