package underthebluesky

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
	Check key.Modifier = "under-the-blue-sky"
	Buff  key.Modifier = "under-the-blue-sky-buff"
)

func init() {
	lightcone.Register(key.UndertheBlueSky, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: onTriggerDeath,
		},
	})

	modifier.Register(Buff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   3,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

	atkPercent := 0.12 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: atkPercent},
		State:  0.09 + 0.03*float64(lc.Imposition),
	})
}

func onTriggerDeath(mod *modifier.ModifierInstance, target key.TargetID) {

	critChance := mod.State().(float64)

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   Buff,
		Source: mod.Owner(),
		Stats:  info.PropMap{prop.CritChance: critChance},
	})
}
