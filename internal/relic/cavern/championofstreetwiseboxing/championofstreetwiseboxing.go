package championofstreetwiseboxing

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	boxing  key.Modifier = "champion-of-streetwise-boxing"
	atkBuff key.Modifier = "champion-of-streetwise-boxing-atk-buff"
)

func init() {
	// 2pc : Increases Physical DMG by 10%.
	// 4pc : After the wearer attacks or is hit, their ATK increases by 5%
	//       for the rest of the battle. This effect can stack up to 5 time(s).

	// register
	relic.Register(key.ChampionOfStreetwiseBoxing, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.PhysicalDamagePercent: 0.1},
			},
			{
				MinCount: 4,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   boxing,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(boxing, modifier.Config{})
	modifier.Register(atkBuff, modifier.Config{})
}
