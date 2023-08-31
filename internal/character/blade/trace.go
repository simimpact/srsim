package blade

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 = "blade-A2"
	A4 = "blade-A4"
	A6 = "blade-A6"
)

func init() {
	// A2
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeDealHeal: func(mod *modifier.Instance, e *event.HealStart) {
				if e.Target.CurrentHPRatio() <= 0.5 {
					e.Target.AddProperty(key.Reason(A2), prop.HealTaken, 0.2)
				}
			},
		},
	})

	// A4
	modifier.Register(A4, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: func(mod *modifier.Instance, e event.AttackEnd) {
				if e.Key != EnhancedNormal {
					return
				}

				for _, target := range e.Targets {
					if mod.Engine().Stance(target) <= 0 {
						// Heal
						mod.Engine().Heal(info.Heal{
							Key:       A4,
							Targets:   []key.TargetID{mod.Owner()},
							Source:    mod.Owner(),
							BaseHeal:  info.HealMap{model.HealFormula_BY_TARGET_MAX_HP: 0.05},
							HealValue: 100,
						})

						return
					}
				}
			},
		},
	})

	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: func(mod *modifier.Instance, e event.HitStart) {
				if e.Hit.Key != Talent {
					return
				}

				e.Hit.Attacker.AddProperty(key.Reason(A6), prop.AllDamagePercent, 0.2)
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
			Name:   A4,
			Source: c.id,
		})
	}

	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}