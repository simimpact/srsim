package welt

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill     key.Attack   = "welt-skill"
	spdDebuff key.Modifier = "welt-spd-down"
	E6        key.Attack   = "welt-e6"
)

// Deals Imaginary DMG equal to 72% of Welt's ATK to a single enemy
// and further deals DMG 2 extra times, with each time dealing Imaginary DMG equal to
// 72% of Welt's ATK to a random enemy. On hit, there is a 75% base chance
// to reduce the enemy's SPD by 10% for 2 turn(s).

func init() {
	modifier.Register(spdDebuff, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_SPEED_DOWN,
		},
	})
}

func (c *char) initSkill() {
	// onAfterHit event listener. add spdDown with chance.
	c.engine.Events().HitEnd.Subscribe(func(e event.HitEnd) {
		if e.Attacker != c.id || e.Key != Skill {
			return
		}

		// E4 : Base chance for Skill to inflict SPD Reduction increases by 35%.
		debuffChance := skillChance[c.info.SkillLevelIndex()]
		if c.info.Eidolon >= 4 {
			debuffChance += 0.35
		}

		c.engine.AddModifier(e.Defender, info.Modifier{
			Name:     spdDebuff,
			Source:   c.id,
			Chance:   debuffChance,
			Stats:    info.PropMap{prop.SPDPercent: 0.1},
			Duration: 2,
		})
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// targeted attack
	c.applySkillAtk(Skill, []key.TargetID{target})
	// extra attack from E1 ult activation
	c.applyE1Pursued(target, 0.8*skillAtk[c.info.SkillLevelIndex()])

	// extra random attacks
	for i := 0; i < 2; i++ {
		chosenTarget := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Enemies(),
			Max:     1,
		})
		c.applySkillAtk(Skill, chosenTarget)
	}
	// E6 : When using Skill, deals DMG for 1 extra time to a random enemy.
	if c.info.Eidolon >= 6 {
		c.applySkillAtk(E6, []key.TargetID{target})
	}

	state.EndAttack()
}

func (c *char) applySkillAtk(atkKey key.Attack, targets []key.TargetID) {
	// E6 check for energy distribution.
	energyAmt := 10.0
	if c.info.Eidolon >= 6 {
		energyAmt = 7.5
	}
	c.engine.Attack(info.Attack{
		Key:        atkKey,
		Source:     c.id,
		Targets:    targets,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_IMAGINARY,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillAtk[c.info.SkillLevelIndex()],
		},
		StanceDamage: 30,
		EnergyGain:   energyAmt,
	})
}
