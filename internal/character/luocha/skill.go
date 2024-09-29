package luocha

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill       = "luocha-skill"
	E2HealBoost = "luocha-e2-healboost"
	E2Shield    = "luocha-e2-shield"
)

func init() {
	modifier.Register(E2HealBoost, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeforeDealHeal: applyE2HealBoost,
		},
	})

	modifier.Register(E2Shield, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		// CanDispel: true,
		Duration: 2,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.Engine().AddShield(E2Shield, info.Shield{
					Source:      mod.Source(),
					Target:      mod.Owner(),
					BaseShield:  info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_ATK: 0.18},
					ShieldValue: 240,
				})
			},
			OnRemove: func(mod *modifier.Instance) {
				mod.Engine().RemoveShield(E2Shield, mod.Owner())
			},
		},
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// do A2
	if c.info.Traces["102"] {
		c.engine.DispelStatus(target, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
	}

	// do E2
	if c.info.Eidolon >= 2 {
		if c.engine.HPRatio(target) < 0.5 {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   E2HealBoost,
				Source: c.id,
			})
		} else {
			c.engine.AddModifier(target, info.Modifier{
				Name:   E2Shield,
				Source: c.id,
			})
		}
	}

	// heal target
	c.engine.Heal(info.Heal{
		Key:     Skill,
		Targets: []key.TargetID{target},
		Source:  c.id,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_ATK: skillPer[c.info.SkillLevelIndex()],
		},
		HealValue: skillFlat[c.info.SkillLevelIndex()],
	})

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Skill,
		Target: c.id,
		Source: c.id,
		Amount: 30,
	})

	// add 1 stack of Abyssal Flower
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   abyssFlower,
		Source: c.id,
	})

	// something with inserts
}

func applyE2HealBoost(mod *modifier.Instance, e *event.HealStart) {
	e.Healer.AddProperty(E2HealBoost, prop.HealBoost, 0.3)
	mod.RemoveSelf()
}
