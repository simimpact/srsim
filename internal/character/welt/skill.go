package welt

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill                  = "welt-skill"
	spdDebuff key.Modifier = "welt-spd-down"
)

// Deals Imaginary DMG equal to 72% of Welt's ATK to a single enemy
// and further deals DMG 2 extra times, with each time dealing Imaginary DMG equal to
// 72% of Welt's ATK to a random enemy. On hit, there is a 75% base chance
// to reduce the enemy's SPD by 10% for 2 turn(s).

func init() {
	modifier.Register(Skill, modifier.Config{})

	modifier.Register(spdDebuff, modifier.Config{})
}

func initSkill() {

}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// attack
	c.engine.Attack(info.Attack{
		Key:        Skill,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_IMAGINARY,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillAtk[c.info.SkillLevelIndex()],
		},
		StanceDamage: 30,
		EnergyGain:   10,
	})

	// extra random attacks
	// NOTE : confirm attackType etc.
	for i := 0; i < 2; i++ {
		chosenTarget := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Enemies(),
			Max:     1,
		})
		c.engine.Attack(info.Attack{
			Key:        Skill,
			Source:     c.id,
			Targets:    chosenTarget,
			AttackType: model.AttackType_PURSUED,
			DamageType: model.DamageType_IMAGINARY,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skillAtk[c.info.SkillLevelIndex()],
			},
			StanceDamage: 30,
			EnergyGain:   10,
		})
	}

	state.EndAttack()
}
