package luka

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.e1Check(target)

	if c.info.Eidolon >= 2 && c.engine.Stats(target).IsWeakTo(model.DamageType_PHYSICAL) {
		c.incrementFightingSprit()
	}

	if c.info.Traces["101"] {
		c.engine.DispelStatus(target, info.Dispel{
			Status: model.StatusType_STATUS_BUFF,
			Order:  model.DispelOrder_LAST_ADDED,
		})
	}

	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_PHYSICAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		StanceDamage: 60,
		EnergyGain:   30,
	})

	c.applyBleed(target)
	c.incrementFightingSprit()

}

func (c *char) applyBleed(target key.TargetID) {
	c.engine.AddModifier(target, info.Modifier{
		Name:   common.Bleed,
		Source: c.id,
		State:  common.BleedState{},
	})
}
