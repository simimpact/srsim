package gallagher

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill key.Heal = "gallagher-skill"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	if c.info.Eidolon >= 2 {
		c.engine.DispelStatus(target, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
		c.engine.AddModifier(target, info.Modifier{
			Name:   E2,
			Source: c.id,
			Stats: info.PropMap{
				prop.EffectRES: 0.3,
			},
			Duration: 2,
		})
	}

	c.engine.Heal(info.Heal{
		Key:     Skill,
		Source:  c.id,
		Targets: []key.TargetID{target},
		// Gallagher is always a flat heal
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_ATK: 0,
		},
		HealValue: skill[c.info.SkillLevelIndex()],
	})
}
