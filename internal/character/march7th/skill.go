package march7th

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill = "march7th-shield"
	E6    = "march7th-e6"
)

type shieldState struct {
	shieldPercentage float64
	shieldFlat       float64
	healPercentage   float64
	healFlat         float64
}

func init() {
	modifier.Register(Skill, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   3,
		Stacking:   modifier.Replace,
		CanDispel:  true,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.Engine().AddShield(Skill, info.Shield{
					Source: mod.Source(),
					Target: mod.Owner(),
					BaseShield: info.ShieldMap{
						model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: mod.State().(shieldState).shieldPercentage,
					},
					ShieldValue: mod.State().(shieldState).shieldFlat,
				})
				if mod.Engine().Stats(mod.Owner()).CurrentHPRatio() >= 0.30 {
					mod.AddProperty(prop.AggroPercent, 5)
				}
			},
			OnRemove: func(mod *modifier.Instance) {
				mod.Engine().RemoveShield(Skill, mod.Owner())
			},
			OnPhase1: func(mod *modifier.Instance) {
				mod.Engine().Heal(info.Heal{
					Key:     E6,
					Targets: []key.TargetID{mod.Owner()},
					Source:  mod.Source(),
					BaseHeal: info.HealMap{
						model.HealFormula_BY_TARGET_MAX_HP: mod.State().(shieldState).healPercentage,
					},
					HealValue:   mod.State().(shieldState).healFlat,
					UseSnapshot: true,
				})
			},
		},
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// A2 check
	if c.info.Traces["101"] {
		c.engine.DispelStatus(target, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
	}

	// A4 check
	shieldDur := 3
	if c.info.Traces["102"] {
		shieldDur += 1
	}

	// E6 Check
	e6HealPercentage := 0.0
	e6HealFlat := 0.0
	if c.info.Eidolon >= 6 {
		e6HealPercentage = 0.04
		e6HealFlat = 106
	}

	c.engine.AddModifier(target, info.Modifier{
		Name:   Skill,
		Source: c.id,
		State: shieldState{
			shieldPercentage: skill[c.info.SkillLevelIndex()],
			shieldFlat:       skillflat[c.info.SkillLevelIndex()],
			healPercentage:   e6HealPercentage,
			healFlat:         e6HealFlat,
		},
		Duration: shieldDur,
	})

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Source: c.id,
		Target: c.id,
		Amount: 30,
		Key:    Skill,
	})
}
