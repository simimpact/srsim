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
	E2     key.Modifier = "clara-e2"
	E4     key.Modifier = "clara-e4"
	E4Buff key.Modifier = "clara-e4-buff"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Stacking:   modifier.Refresh,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E4Buff, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E4, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterBeingAttacked: func(mod *modifier.Instance, e event.AttackEnd) {
				mod.Engine().AddModifier(mod.Owner(), info.Modifier{
					Name:            E4Buff,
					Source:          mod.Owner(),
					Duration:        1,
					Stats:           info.PropMap{prop.AllDamageReduce: 0.3},
					TickImmediately: true,
				})
			},
		},
	})
}

// After using the Ultimate, ATK increases by 30% for 2 turn(s).
func (c *char) e2() {
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:            E2,
			Source:          c.id,
			Stats:           info.PropMap{prop.ATKPercent: 0.3},
			Duration:        2,
			TickImmediately: true,
		})
	}
}

func (c *char) initEidolons() {
	// After Clara is hit, the DMG taken by Clara is reduced by 30%. This effect
	// lasts until the start of her next turn.
	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E4,
			Source: c.id,
		})
	}
}
