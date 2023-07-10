package clara

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// A2:
// When attacked, this character has a 35.00 % fixed chance to remove a debuff placed on them.
// A4:
// The chance to resist Crowd Control Debuffs increases by 35.00 %.
// A6:
// Increases Svarog's Counter DMG by 30.00 %.

const (
	A2 key.Modifier = "clara-a2"
	A4 key.Modifier = "clara-a4"
	A6 key.Modifier = "clara-a6"
)

func init() {
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeBeingAttacked: func(mod *modifier.Instance, e event.AttackStart) {
				if mod.Engine().Rand().Float32() < 0.35 {
					mod.Engine().DispelStatus(mod.Owner(), info.Dispel{
						Status: model.StatusType_STATUS_DEBUFF,
						Order:  model.DispelOrder_LAST_ADDED,
						Count:  1,
					})
				}
			},
		},
	})

	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: func(mod *modifier.Instance, e event.HitStart) {
				if e.Hit.AttackType == model.AttackType_INSERT {
					e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.3)
				}
			},
		},
	})
}

func (c *char) initTraces() {
	if c.info.Traces["101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
		})
	}

	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:      A4,
			Source:    c.id,
			DebuffRES: info.DebuffRESMap{model.BehaviorFlag_STAT_CTRL: 0.35},
		})
	}

	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}
