package march7th

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill  = "march7th-skill"
	E6heal = "march7th-e6hot"
)

type shieldState struct {
	shieldPercentage float64
	shieldFlat       float64
}

func init() {
	modifier.Register(Skill, modifier.Config{
		Duration: 3,
		Stacking: modifier.Replace,
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
				if mod.Engine().Stats(mod.Owner()).CurrentHPRatio() >= 30 {
					mod.AddProperty(prop.AggroPercent, 5)
				}
			},
			OnRemove: func(mod *modifier.Instance) {
				mod.Engine().RemoveShield(Skill, mod.Owner())
			},
			OnPhase1: func(mod *modifier.Instance) {
				march7th, _ := mod.Engine().CharacterInfo(mod.Source())
				if march7th.Eidolon >= 6 {
					mod.Engine().Heal(info.Heal{
						Key:     E6heal,
						Targets: []key.TargetID{mod.Owner()},
						Source:  mod.Source(),
						BaseHeal: info.HealMap{
							model.HealFormula_BY_TARGET_MAX_HP: 0.4,
						},
						HealValue:   106,
						UseSnapshot: true,
					})
				}
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

	c.engine.AddModifier(target, info.Modifier{
		Name:   Skill,
		Source: c.id,
		State: shieldState{
			shieldPercentage: skill[c.info.SkillLevelIndex()],
			shieldFlat:       skillflat[c.info.SkillLevelIndex()],
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
