package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2 key.Modifier = "silverwolf-e2"
	E6 key.Modifier = "silverwolf-e6"
)

func init() {
	// When an enemy enters battle, reduces their Effect RES by 20%.
	modifier.Register(E2, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
		TickMoment: modifier.ModifierPhase1End,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.SetProperty(prop.EffectRES, -0.2)
			},
		},
	})

	// For every debuff the target enemy has, the DMG dealt by Silver Wolf increases by 20%, up to a limit of 100%.
	modifier.Register(E6, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: func(mod *modifier.Instance, e event.HitStart) {
				debuffCount := mod.Engine().ModifierStatusCount(e.Defender, model.StatusType_STATUS_DEBUFF)
				if debuffCount > 5 {
					debuffCount = 5
				}
				e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.2*float64(debuffCount))
			},
		},
	})
}

// After using her Ultimate to attack enemies, Silver Wolf regenerates 7
// Energy for every debuff that the target enemy currently has. This effect
// can be triggered up to 5 time(s) in each use of her Ultimate.
func (c *char) e1(target key.TargetID) {
	if c.info.Eidolon >= 1 {
		debuffCount := c.engine.ModifierStatusCount(target, model.StatusType_STATUS_DEBUFF)
		if debuffCount > 5 {
			debuffCount = 5
		}
		c.engine.ModifyEnergy(c.id, float64(7*debuffCount))
	}
}

// After using her Ultimate to attack enemies, deals Additional Quantum DMG
// equal to 20% of Silver Wolf's ATK for every debuff currently on the enemy
// target. This effect can be triggered for a maximum of 5 time(s) during
// each use of her Ultimate.
func (c *char) e4(target key.TargetID) {
	if c.info.Eidolon >= 4 {
		debuffCount := c.engine.ModifierStatusCount(target, model.StatusType_STATUS_DEBUFF)
		if debuffCount > 5 {
			debuffCount = 5
		}
		for i := 0; i < debuffCount; i++ {
			c.engine.Attack(info.Attack{
				Source:     c.id,
				Targets:    []key.TargetID{target},
				DamageType: model.DamageType_QUANTUM,
				AttackType: model.AttackType_PURSUED,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: 0.2,
				},
			})
		}
	}
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 2 {
		c.engine.Events().EnemiesAdded.Subscribe(func(e event.EnemiesAdded) {
			for _, enemy := range e.Enemies {
				c.engine.AddModifier(enemy.ID, info.Modifier{
					Name:   E2,
					Source: c.id,
				})
			}
		})

		c.engine.Events().TargetDeath.Subscribe(func(event event.TargetDeath) {
			if event.Target == c.id {
				for _, trg := range c.engine.Enemies() {
					c.engine.RemoveModifier(trg, E2)
				}
			}
		})
	}

	if c.info.Eidolon >= 6 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6,
			Source: c.id,
		})
	}
}
