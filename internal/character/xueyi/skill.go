package xueyi

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill key.Attack = "xueyi-skill"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// Main target
	c.engine.Attack(info.Attack{
		Targets:    []key.TargetID{target},
		Key:        Skill,
		Source:     c.id,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_QUANTUM,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillmain[c.info.AttackLevelIndex()],
		},
		EnergyGain:   30,
		StanceDamage: 60,
	})

	// Adjacents
	c.engine.Attack(info.Attack{
		Targets:    c.engine.AdjacentTo(target),
		Key:        Skill,
		Source:     c.id,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_QUANTUM,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skilladj[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30,
	})

	state.EndAttack()
}
