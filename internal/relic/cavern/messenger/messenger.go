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
				MinCount:     2,
				Stats:        info.PropMap{prop.SPDPercent: 0.06},
				CreateEffect: nil,
			},
			{
				MinCount: 4,
				Stats:    nil,
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
			OnBeforeAction: BuffSelf,
			OnAfterAction:  BuffAllies,
		},
	})
	modifier.Register(buff, modifier.Config{
		Stacking:      modifier.Replace,
		StatusType:    model.StatusType_STATUS_BUFF,
		CanDispel:     true,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
		Duration:      1,
	})
}

// Workaround for missing "eligible" target list in ActionStart:
// Apply buff to owner when OnBeforeAction is triggered
// KNOWN BUG: this will also buff owners equipping this 4p set even if they do not target allies with an Ult
func BuffSelf(mod *modifier.Instance, e event.ActionStart) {
	if e.AttackType == model.AttackType_ULT {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   buff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.SPDPercent: 0.12},
		})
	}
}

// Apply buff to other allies when OnAfterAction is triggered
// INACCURACY: this should apply instead when OnBeforeAction is triggered
func BuffAllies(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType == model.AttackType_ULT {
		for _, char := range mod.Engine().Characters() {
			if char == mod.Owner() {
				continue
			}
			mod.Engine().AddModifier(char, info.Modifier{
				Name:   buff,
				Source: mod.Owner(),
				Stats:  info.PropMap{prop.SPDPercent: 0.12},
			})
		}
	}
}
