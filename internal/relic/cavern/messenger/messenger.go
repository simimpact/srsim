package messenger

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	check = "messenger-traversing-hackerspace"
	buff  = "messenger-traversing-hackerspace-buff"
)

// 2pc: Increases SPD by 6%.
// 4pc: When the wearer uses their Ultimate on an ally,
//      SPD for all allies increases by 12% for 1 turn(s). This effect cannot be stacked.

func init() {
	relic.Register(key.MessengerTraversingHackerspace, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.SPDPercent: 0.06},
			},
			{
				MinCount: 4,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   check,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: onBeforeUlt,
		},
	})
	modifier.Register(buff, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd: onAdd,
		},
	})
}

func onBeforeUlt(mod *modifier.Instance, e event.ActionStart) {
	if e.AttackType == model.AttackType_ULT {
		for _, char := range mod.Engine().Characters() {
			mod.Engine().AddModifier(char, info.Modifier{
				Name:     buff,
				Source:   mod.Owner(),
				Duration: 1,
			})
		}
	}
}

func onAdd(mod *modifier.Instance) {
	mod.AddProperty(prop.SPDPercent, 0.12)
}
