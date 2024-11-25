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
	MutualDemiseCheck key.Modifier = "mutual-demise-check"
	MutualDemiseBuff  key.Modifier = "mutual-demise-buff"
)

func init() {
	lightcone.Register(key.MutualDemise, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(MutualDemiseCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: adjustCritRate,
			OnHPChange: func(mod *modifier.Instance, e event.HPChange) {
				adjustCritRate(mod)
			},
		},
	})

	modifier.Register(MutualDemiseBuff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		CanDispel:  true,
	})
}

// If the wearer's HP is less than 80%, CRIT Rate increases by 12%/15%/18%/21%/24%
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   MutualDemiseCheck,
		Source: owner,
		State:  0.09 + 0.03*float64(lc.Imposition),
	})
}

func adjustCritRate(mod *modifier.Instance) {
	if mod.Engine().HPRatio(mod.Owner()) < 0.8 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   MutualDemiseBuff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.CritChance: mod.State().(float64)},
		})
	} else {
		mod.Engine().RemoveModifier(mod.Owner(), MutualDemiseBuff)
	}
}
