package arlan

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// A2:
// 	If the current HP percentage is 30% or lower when defeating an enemy, immediately restores HP equal to 20% of Max HP.
// A4:
//	The chance to resist DoT Debuffs increases by 50%.
// A6:
//	Upon entering battle, if Arlan's HP is less than or equal to 50%, he can nullify all DMG received except for DoTs until he is attacked.

const (
	A2 key.Modifier = "arlan-a2"
	A4 key.Modifier = "arlan-a4"
	A6 key.Modifier = "arlan-a6"
)

func init() {
	// A2 Register
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: func(mod *modifier.Instance, target key.TargetID) {
				if mod.Engine().HPRatio(mod.Owner()) <= 0.3 {
					mod.Engine().Heal(info.Heal{
						Targets:  []key.TargetID{mod.Owner()},
						Source:   mod.Owner(),
						BaseHeal: info.HealMap{model.HealFormula_BY_TARGET_MAX_HP: 0.2},
					})
				}
			},
		},
	})

	// A6 Register
	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				// TODO: https://github.com/simimpact/srsim/issues/13
			},
			OnAfterBeingAttacked: func(mod *modifier.Instance, e event.AttackEnd) {
				mod.RemoveSelf()
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
			DebuffRES: info.DebuffRESMap{model.BehaviorFlag_STAT_DOT: 0.5},
		})
	}

	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}
