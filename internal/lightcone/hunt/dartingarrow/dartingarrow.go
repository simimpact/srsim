package dartingarrow

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
	DartingArrowCheck key.Modifier = "darting-arrow-check"
	DartingArrowBuff  key.Modifier = "darting-arrow-buff"
)

// When the wearer defeats an enemy, increases ATK by 24%/30%/36%/42%/48% for 3 turn(s).
func init() {
	lightcone.Register(key.DartingArrow, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(DartingArrowCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: onTriggerDeath,
		},
	})

	modifier.Register(DartingArrowBuff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   DartingArrowCheck,
		Source: owner,
		State:  0.18 + 0.06*float64(lc.Imposition),
	})
}

func onTriggerDeath(mod *modifier.Instance, target key.TargetID) {
	amt := mod.State().(float64)

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:     DartingArrowBuff,
		Source:   mod.Owner(),
		Duration: 3,
		Stats:    info.PropMap{prop.ATKPercent: amt},
	})
}
