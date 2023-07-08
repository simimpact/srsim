package clara

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult        key.Modifier = "clara-ult" // MAvatar_Klara_00_Ultra_WarriorMode
	UltCounter key.Modifier = "clara-ult-enhanced-counter"
)

type ultCounterNum (int)

func init() {
	modifier.Register(Ult, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_BURST},
		Stacking:      modifier.Refresh,
		Duration:      2,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				state := mod.State().(State)

				// extra res
				mod.AddProperty(prop.AllDamageReduce, ult_cut[state.ultLevelIndex])
				// extra aggro
				mod.AddProperty(prop.AggroPercent, 5)

				amt := 2
				if mod.Engine().HasModifier(mod.Owner(), E6) {
					amt = 3
				}

				mod.Engine().AddModifier(mod.Owner(), info.Modifier{
					Name:  UltCounter,
					State: ultCounterNum(amt),
				})
			},
		},
	})

	// mod check in talent.go
	modifier.Register(UltCounter, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterHit: func(mod *modifier.Instance, e event.HitEnd) {
				// after counter
				if e.AttackType == model.AttackType_INSERT {
					remaining := mod.State().(ultCounterNum)
					remaining -= 1
					if remaining <= 0 { // executes 2 counters
						mod.RemoveSelf()
					}
				}
			},
		},
	})
}

// After Clara uses Ultimate, DMG dealt to her is reduced by an extra *%, and
// she has greatly increased chances of being attacked by enemies for 2 turn(s).
// In addition, Svarog's Counter is enhanced. When an ally is attacked, Svarog
// immediately launches a Counter, and its DMG multiplier against the enemy
// increases by *%. Enemies adjacent to it take 50% of the DMG dealt to the
// target enemy. Enhanced Counter(s) can take effect 2 time(s).
func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Ult,
		Source: c.id,
	})

	c.e2()

	c.engine.ModifyEnergy(c.id, 5.0)
}
