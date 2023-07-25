package adversarial

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
	AdversarialCheck key.Modifier = "adversarial_check"
	AdversarialBuff  key.Modifier = "adversarial_buff"
)

// When the wearer defeats an enemy, increases SPD by 10/12/14/16/18% for
// 2 turn(s)
func init() {
	lightcone.Register(key.Adversarial, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(AdversarialCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: onTriggerDeath,
		},
	})

	modifier.Register(AdversarialBuff, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_BUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   AdversarialCheck,
		Source: owner,
		State:  0.08 + 0.02*float64(lc.Imposition),
	})
}

func onTriggerDeath(mod *modifier.Instance, target key.TargetID) {
	amt := mod.State().(float64)

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:     AdversarialBuff,
		Source:   mod.Owner(),
		Duration: 2,
		Stats:    info.PropMap{prop.SPDPercent: amt},
	})
}
