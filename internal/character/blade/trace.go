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
	A2Check = "blade-A2-check"
	A2Buff  = "blade-A2-buff"
	A4      = "blade-A4"
	A6      = "blade-A6"
)

func init() {
	// A2
	modifier.Register(A2Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: a2HPCheck,
			OnHPChange: func(mod *modifier.Instance, e event.HPChange) {
				a2HPCheck(mod)
			},
		},
	})

	modifier.Register(A2Buff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})

	// A4
	modifier.Register(A4, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: func(mod *modifier.Instance, e event.AttackEnd) {
				if e.Key != EnhancedNormalPrimary && e.Key != EnhancedNormalAdjacent {
					return
				}

				if mod.Engine().Stance(e.Targets[0]) <= 0 {
					// Heal
					mod.Engine().Heal(info.Heal{
						Key:       A4,
						Targets:   []key.TargetID{mod.Owner()},
						Source:    mod.Owner(),
						BaseHeal:  info.HealMap{model.HealFormula_BY_TARGET_MAX_HP: 0.05},
						HealValue: 100,
					})
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
			Name:   A2Check,
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

func a2HPCheck(mod *modifier.Instance) {
	if mod.Engine().HPRatio(mod.Owner()) <= 0.5 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   A2Buff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.HealTaken: 0.2},
		})
	} else {
		mod.Engine().RemoveModifier(mod.Owner(), A2Buff)
	}
}
