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
	E2Buff key.Modifier = "clara-e2-buff"
	E4     key.Modifier = "clara-e4"
	E4Buff key.Modifier = "clara-e4-buff"
	E6     key.Modifier = "clara-e6"
)

func init() {
	modifier.Register(E2Buff, modifier.Config{
		Stacking:   modifier.Refresh,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E4Buff, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
	})

	modifier.Register(E4, modifier.Config{
		Stacking:   modifier.Refresh,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAfterBeingAttacked: func(mod *modifier.Instance, e event.AttackEnd) {
				mod.Engine().AddModifier(mod.Owner(), info.Modifier{
					Name:     E4Buff,
					Source:   mod.Owner(),
					Duration: 1,
					Stats:    info.PropMap{prop.AllDamageReduce: 0.3},
				})
			},
		},
	})

	// tag modifier, don't need listeners for E6 cause talent handles all
	// enemies attacks
	modifier.Register(E6, modifier.Config{})
}

// Using Skill will not remove Marks of Counter on the enemy.
func (c *char) e1() {
	if c.info.Eidolon == 0 {
		for _, enemy := range c.engine.Enemies() {
			c.engine.RemoveModifier(enemy, TalentMark)
		}
	}
}

// After using the Ultimate, ATK increases by 30% for 2 turn(s).
func (c *char) e2() {
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:     E2Buff,
			Source:   c.id,
			Stats:    info.PropMap{prop.ATKPercent: 0.3},
			Duration: 2,
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

	// After other allies are hit, Svarog also has a 50% fixed chance to trigger a
	// Counter on the attacker and mark them with a Mark of Counter. When using
	// Ultimate, the number of Enhanced Counters increases by 1.
	if c.info.Eidolon >= 6 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6,
			Source: c.id,
		})
	}
}
