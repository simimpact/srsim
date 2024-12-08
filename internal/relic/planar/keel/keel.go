package keel

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

// 2pc: Increases the wearer's Effect RES by 10%.
//      When the wearer's Effect RES is at 30% or higher, all allies' CRIT DMG increases by 10%.

const (
	check    = "broken-keel"
	keelcdmg = "broken-keel-cdmg"
)

type state struct {
	// "flag" to reduce redundancy when adding and removing buff
	applied bool
}

func init() {
	relic.Register(key.BrokenKeel, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.EffectRES: 0.1},
				CreateEffect: nil,
			},
			{
				MinCount:     2,
				Stats:        nil,
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
		Stacking:      modifier.ReplaceBySource,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_REMOVE_WHEN_SOURCE_DEAD},
	})
}

func Create(engine engine.Engine, owner key.TargetID) {
	engine.AddModifier(owner, info.Modifier{
		Name:   check,
		Source: owner,
		State:  &state{applied: false},
	})
	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		for char := range event.CharInfo {
			engine.AddModifier(char, info.Modifier{
				Name:   keelcdmg,
				Source: owner,
			})
		}
	})
	// remove buff from all chars on planar holder's death
	engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if e.Target == owner {
			for _, char := range engine.Characters() {
				engine.RemoveModifierFromSource(char, owner, keelcdmg)
			}
		}
	})
}

func onCheck(mod *modifier.Instance) {
	stats := mod.OwnerStats()
	st := mod.State().(*state)

	if stats.EffectRES() >= 0.3 && !st.applied {
		for _, c := range mod.Engine().Characters() {
			mod.Engine().AddModifier(c, info.Modifier{
				Name:   keelcdmg,
				Source: mod.Owner(),
				Stats:  info.PropMap{prop.CritDMG: 0.1},
			})
		}
		st.applied = true
	}

	if stats.EffectRES() < 0.3 && st.applied {
		for _, c := range mod.Engine().Characters() {
			mod.Engine().RemoveModifierFromSource(c, mod.Owner(), keelcdmg)
		}
		st.applied = false
	}
}
