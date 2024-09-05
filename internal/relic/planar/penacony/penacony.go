package penacony

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	mod     key.Modifier = "penacony-land-of-dreams"
	dmgBuff key.Modifier = "penacony-buff"
)

// Increases wearer's Energy Regeneration Rate by 5%.
// Increases DMG by 10% for all other allies that are of the same Type as the wearer.
func init() {
	relic.Register(key.PenaconyLandOfDreams, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.EnergyRegen: 0.05},
			},
			{
				MinCount: 2,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   mod,
						Source: owner,
					})
				},
			},
		},
	})

	modifier.Register(dmgBuff, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

func Create(engine engine.Engine, owner key.TargetID) {
	holderInfo, _ := engine.CharacterInfo(owner)

	engine.Events().BattleStart.Subscribe(func(e event.BattleStart) {
		for _, char := range engine.Characters() {
			charInfo, _ := engine.CharacterInfo(char)
			if charInfo.Element == holderInfo.Element {
				engine.AddModifier(char, info.Modifier{
					Name:   dmgBuff,
					Source: owner,
					Stats:  info.PropMap{prop.AllDamagePercent: 0.1},
				})
			}
		}
	})

	engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if e.Target == owner {
			for _, char := range engine.Characters() {
				engine.RemoveModifierFromSource(char, owner, dmgBuff)
			}
		}
	})
}
