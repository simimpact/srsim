package fleet

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	mod key.Modifier = "fleet-of-the-ageless"
	buff key.Modifier ="fleet-of-the-ageless-atk-buff"
)

// Increases the wearer's Max HP by 12%. When the wearer's SPD reaches 120 or higher,
// all allies' ATK increases by 8%.
func init() {
	relic.Register(key.FleetOfTheAgeless, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.HPPercent: 0.12},
			},
			{
				MinCount: 2,
				CreateEffect: create(engine.Engine, key.TargetID),
			},
		},
	})

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd:            check,
			OnPropertyChange: check,
		},
	})

	modifier.Register(buff, modifier.Config{})
}


func create (engine engine.Engine, owner  key.TargetID){
	engine.Events().BattleStart.Subscribe(func(event event.BattleStart){
		for char := range event.CharInfo {
			engine.AddModifier(char, mod)
		}
	})
}

func check(mod *modifier.Instance) {
	stats := mod.OwnerStats()
	applied := mod.State().(*bool)

	if stats.SPD() >= 120 && !applied {
		for _, c := range engine.Characters() {
			engine.AddModifier(i, info.Modifier{
				Name:   buff,
				Source: owner,
				Stats: info.PropMap{prop.ATKPercent, 0.08}
				
			})
		}
		mod.State().(*bool) = true
	}

	if stats.SPD() < 120 && applied {
		for _, c := range engine.Characters() {
			mod.Engine().RemoveModifierFromSource(c, mod.Owner(), buff)
		}
		mod.State().(*bool) = false
	}
}


