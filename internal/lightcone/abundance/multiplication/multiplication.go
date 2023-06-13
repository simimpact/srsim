package multiplication

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Multiplication = "multiplication"

func init() {
	lightcone.Register(key.Multiplication, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})

	modifier.Register(Multiplication, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: func(mod *modifier.ModifierInstance, e event.ActionEvent) {
				rank := mod.State().(int)
				if e.AttackType == model.AttackType_NORMAL {
					mod.Engine().ModifyCurrentGaugeCost(-0.1 - float64(rank)*0.02)
				}
			},
		},
	})

}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   Multiplication,
		Source: owner,
		State:  lc.Rank,
	})
}
