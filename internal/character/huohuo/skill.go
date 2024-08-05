package huohuo

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill       key.Heal   = "huohuo-skill"
	SkillReason key.Reason = "huohuo-skill"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.DispelStatus(target, info.Dispel{
		Status: model.StatusType_STATUS_DEBUFF,
		Order:  model.DispelOrder_LAST_ADDED,
		Count:  1,
	})
	c.engine.Heal(info.Heal{
		Key:     Skill,
		Targets: []key.TargetID{target},
		Source:  c.id,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: skillMainRate[c.info.SkillLevelIndex()],
		},
		HealValue: skillMainValue[c.info.SkillLevelIndex()],
	})
	c.engine.Heal(info.Heal{
		Key:     Skill,
		Targets: c.engine.AdjacentTo(target),
		Source:  c.id,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: skillAdjacentRate[c.info.SkillLevelIndex()],
		},
		HealValue: skillAdjacentValue[c.info.SkillLevelIndex()],
	})
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Target: c.id,
		Source: c.id,
		Amount: 30,
		Key:    SkillReason,
	})

	c.TalentRound = 2
	if c.info.Eidolon >= 1 {
		c.TalentRound = 3
	}
	c.DispelCount = 6
	for _, target := range c.engine.Characters() {
		if c.info.Eidolon >= 1 {
			c.engine.AddModifier(target, info.Modifier{
				Name:   TalentBuff,
				Source: c.id,
				Stats:  info.PropMap{prop.SPDPercent: 0.12},
			})
			continue
		}
		c.engine.AddModifier(target, info.Modifier{
			Name:   TalentBuff,
			Source: c.id,
		})
	}
}
