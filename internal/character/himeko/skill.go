package himeko

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	a4    = "himeko-a4"
	skill = "himeko-skill"
)

func init() {
	modifier.Register(a4, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: a4listener,
		},
	})
}

var skillSplitMain = []float64{0.2, 0.2, 0.05, 0.05, 0.05, 0.05, 0.05, 0.4}
var skillSplitAdj = []float64{0.2, 0.2, 0.6}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   a4,
			Source: c.id,
		})
	}

	// Main target
	for index, ratio := range skillSplitMain {
		c.engine.Attack(info.Attack{
			Key:          skill,
			HitIndex:     index,
			HitRatio:     ratio,
			Targets:      []key.TargetID{target},
			Source:       c.id,
			AttackType:   model.AttackType_SKILL,
			DamageType:   model.DamageType_FIRE,
			EnergyGain:   20,
			StanceDamage: 30,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skillMain[c.info.SkillLevelIndex()],
			},
		})
	}

	// Adjacent targets
	for index, ratio := range skillSplitAdj {
		c.engine.Attack(info.Attack{
			Key:          skill,
			HitIndex:     index,
			HitRatio:     ratio,
			Targets:      c.engine.AdjacentTo(target),
			Source:       c.id,
			AttackType:   model.AttackType_SKILL,
			DamageType:   model.DamageType_FIRE,
			EnergyGain:   20,
			StanceDamage: 30,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skillAdj[c.info.SkillLevelIndex()],
			},
		})
	}

	c.engine.RemoveModifier(c.id, a4)

	c.engine.EndAttack()
}

func a4listener(mod *modifier.Instance, e event.HitStart) {
	isSkill := e.Hit.AttackType == model.AttackType_SKILL
	targetHasBurn := mod.Engine().HasBehaviorFlag(e.Defender, model.BehaviorFlag_STAT_DOT_BURN)
	if isSkill && targetHasBurn {
		e.Hit.Attacker.AddProperty(a4, prop.AllDamagePercent, 0.2)
	}
}
