package sagacity

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
	checker  key.Modifier = "sagacity-check"
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
	modifier.Register(checker, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: buffAtkOnUlt,
		},
	})
	modifier.Register(sagacity, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	atkAmt := 0.18 + 0.06*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   checker,
		Source: owner,
		State:  atkAmt,
	})
}

func buffAtkOnUlt(mod *modifier.Instance, e event.ActionStart) {
	if e.AttackType != model.AttackType_ULT {
		return
	}
	atkAmt := mod.State().(float64)
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:     sagacity,
		Source:   mod.Owner(),
		Stats:    info.PropMap{prop.ATKPercent: atkAmt},
		Duration: 2,
	})
}
