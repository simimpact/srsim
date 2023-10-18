package sagacity

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	sagacity key.Modifier = "sagacity"
)

// When the wearer uses their Ultimate, increases ATK by 24% for 2 turn(s).

func init() {
	lightcone.Register(key.Sagacity, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
	modifier.Register(sagacity, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: buffAtkOnUlt,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	atkAmt := 0.18 + 0.06*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   sagacity,
		Source: owner,
		State:  atkAmt,
	})
}

func buffAtkOnUlt(mod *modifier.Instance, e event.ActionStart) {

}
