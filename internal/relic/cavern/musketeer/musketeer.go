package musketeer

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	mod = key.Modifier("musketeer-of-wild-wheat")
)

func init() {
	relic.Register(key.MusketeerOfWildWheat, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{model.Property_ATK_PERCENT: 0.12},
			},
			{
				MinCount: 4,
				Stats:    info.PropMap{model.Property_SPD_PERCENT: 0.06},
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
			// TODO: implement OnBeforeHit hook
		},
	})
}
