package lushaka

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
	name = "lushaka-the-sunken-seas"
)

// Increases the wearer's Energy Regeneration Rate by 5%.
// If the wearer is not the first character in the team lineup,
// then increases the ATK of the first character in the team lineup by 12%.
func init() {
	relic.Register(key.Lushaka, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.EnergyRegen: 0.05},
				CreateEffect: Create,
			},
		},
	})

	modifier.Register(name, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		CanDispel:  true,
	})
}

func Create(engine engine.Engine, owner key.TargetID) {
	firstCharacter := engine.Characters()[0]

	engine.Events().BattleStart.Subscribe(func(e event.BattleStart) {
		if firstCharacter != owner {
			engine.AddModifier(firstCharacter, info.Modifier{
				Name:   name,
				Source: owner,
				Stats:  info.PropMap{prop.ATKPercent: 0.12},
			})
		}
	})

	engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if e.Target == owner {
			engine.RemoveModifierFromSource(firstCharacter, owner, name)
		}
	})
}
