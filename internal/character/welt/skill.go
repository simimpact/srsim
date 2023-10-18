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
		c.engine.AddModifier(e.Defender, info.Modifier{
			Name:     spdDebuff,
			Source:   c.id,
			Chance:   skillChance[c.info.SkillLevelIndex()],
			Stats:    info.PropMap{prop.SPDPercent: 0.1},
			Duration: 2,
		})
	})
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

	c.applyE1Pursued(target, 0.8*skillAtk[c.info.SkillLevelIndex()])
	// extra random attacks
	// TODO : confirm : DM uses IncludeLimbo but exclude targets w/ HP <= 0
	for i := 0; i < 2; i++ {
		chosenTarget := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Enemies(),
			Max:     1,
		})
		c.engine.Attack(info.Attack{
			Key:        Skill,
			Source:     c.id,
			Targets:    chosenTarget,
			AttackType: model.AttackType_SKILL,
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
