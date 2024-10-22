package hunterofglacialforest

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
	mod  = key.Modifier("hunter-of-glacial-forest")
	buff = key.Modifier("hunter-of-glacial-forest-buff")
)

// 2pc: Increase Ice DMG by 10%
// 4pc: After the wearer uses their Ultimate, their CRIT DMG increases by 25% for 2 turn(s)

func init() {
	relic.Register(key.HunterOfGlacialForest, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.IceDamagePercent: 0.10},
				CreateEffect: nil,
			},
			{
				MinCount: 4,
				Stats:    nil,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   mod,
						Source: owner,
					})
				},
			},
		},
	})

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: onBeforeAction,
		},
	})

	modifier.Register(buff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func onBeforeAction(mod *modifier.Instance, e event.ActionStart) {
	if e.AttackType == model.AttackType_ULT {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     buff,
			Source:   mod.Owner(),
			Duration: 2,
			Stats:    info.PropMap{prop.CritDMG: 0.25},
		})
	}
}
