package eagle

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

const eagle = "eagle-of-twilight-thunder"

// 2pc: Increases Wind DMG by 10%.
// 4pc: After the wearer uses their Ultimate, their action is Advanced Forward by 25%.

func init() {
	relic.Register(key.EagleOfTwilightLine, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.WindDamagePercent: 0.1},
			},
			{
				MinCount: 4,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   eagle,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(eagle, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAction: onAfterUltimate,
		},
	})
}

func onAfterUltimate(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType == model.AttackType_ULT {
		mod.Engine().ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    eagle,
			Target: mod.Owner(),
			Source: mod.Owner(),
			Amount: -0.25,
		})
	}
}
