package clara

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// marked enemy attack instance
	markedEnemies := make([]key.TargetID, 0, 5)
	for _, enemy := range c.engine.Enemies() {
		if c.engine.HasModifier(enemy, TalentMark) {
			markedEnemies = append(markedEnemies, enemy)
		}
	}

	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    markedEnemies,
		DamageType: model.DamageType_PHYSICAL,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
	})

	// usual skill attack
	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_PHYSICAL,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		StanceDamage: 30.0,
		EnergyGain:   30,
	})

	state.EndAttack()

	// E1: Using Skill will not remove Marks of Counter on the enemy.
	if c.info.Eidolon == 0 {
		for _, enemy := range c.engine.Enemies() {
			c.engine.RemoveModifier(enemy, TalentMark)
		}
	}
}
