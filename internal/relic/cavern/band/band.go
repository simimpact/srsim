package band

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
	check = "band-of-sizzling-thunder"
	buff  = "band-of-sizzling-thunder-buff"
)

// 2pc: Increases Lightning DMG by 10%.
// 4pc: When the wearer uses their Skill, increases the wearer's ATK by 20% for 1 turn(s).

func init() {
	relic.Register(key.BandOfSizzlingThunder, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.ThunderDamagePercent: 0.1},
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
			OnBeforeAction: onBeforeSkill,
		},
	})
	modifier.Register(buff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
		Duration:   1,
	})
}

func onBeforeSkill(mod *modifier.Instance, e event.ActionStart) {
	if e.AttackType == model.AttackType_SKILL {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   buff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.ATKPercent: 0.2},
		})
	}
}
