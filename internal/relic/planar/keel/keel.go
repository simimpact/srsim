package keel

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

// 2pc: Increases the wearer's Effect RES by 10%.
//      When the wearer's Effect RES is at 30% or higher, all allies' CRIT DMG increases by 10%.

// TO-DO: how to handle `applied` and pointer logic, where to put the stats

const (
	check    = "broken-keel"
	keelcdmg = "broken-keel-cdmg"
)

type state struct {
	applied *bool
}

func init() {
	relic.Register(key.BrokenKeel, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.EffectRES: 0.1},
			},
			{
				MinCount:     2,
				CreateEffect: Create,
			},
		},
	})
	modifier.Register(check, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd:            onCheck,
			OnPropertyChange: onCheck,
		},
	})
	modifier.Register(keelcdmg, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

func Create(engine engine.Engine, owner key.TargetID) {
	engine.AddModifier(owner, info.Modifier{
		Name:   check,
		Source: owner,
	})
	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		appliedInit := false
		for char := range event.CharInfo {
			engine.AddModifier(char, info.Modifier{
				Name:   keelcdmg,
				Source: owner,
				State: state{
					applied: &appliedInit,
				},
			})
		}
	})
}

func onCheck(mod *modifier.Instance) {
	stats := mod.OwnerStats()
	applied := mod.State().(*bool)

	if stats.EffectRES() >= 0.3 && !*applied {
		for _, c := range mod.Engine().Characters() {
			mod.Engine().AddModifier(c, info.Modifier{
				Name:   keelcdmg,
				Source: mod.Owner(),
				Stats:  info.PropMap{prop.CritDMG: 0.1},
			})
		}
		*mod.State().(*bool) = true
	}

	if stats.EffectRES() < 0.3 && *applied {
		for _, c := range mod.Engine().Characters() {
			mod.Engine().RemoveModifierFromSource(c, mod.Owner(), keelcdmg)
		}
		*mod.State().(*bool) = false
	}
}
