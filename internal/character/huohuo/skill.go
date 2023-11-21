package huohuo

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill key.Heal = "huohuo-skill"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
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
	c.engine.DispelStatus(target, info.Dispel{
		Status: model.StatusType_STATUS_DEBUFF,
		Order:  model.DispelOrder_LAST_ADDED,
		Count:  1,
	})

	c.TalentRound = 2
	c.DispelCount = 6
	targets := c.engine.Characters()
	for _, target := range targets {
		c.engine.AddModifier(target, info.Modifier{
			Name:   TalentBuff,
			Source: c.id,
		})
	}
}
