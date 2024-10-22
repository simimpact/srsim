package guinaifen

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill = "guinaifen-skill"
	E1    = "guinaifen-e1"
)

func init() {
	modifier.Register(E1, modifier.Config{
		StatusType: model.StatusType_STATUS_DEBUFF,
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// E1
	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(target, info.Modifier{
			Name:     E1,
			Source:   c.id,
			Duration: 2,
			Chance:   1,
			Stats:    info.PropMap{prop.EffectRES: -0.1},
		})
	}

	// get target list for Burn
	targetList := c.engine.Retarget(info.Retarget{
		Targets: append([]key.TargetID{target}, c.engine.AdjacentTo(target)...),
		Max:     3,
	})

	// apply Burn with Skill's chance
	for _, trg := range targetList {
		multiplier := skillBurn[c.info.SkillLevelIndex()]
		chance := 1.0
		// E2
		if c.info.Eidolon >= 2 && c.engine.HasBehaviorFlag(trg, model.BehaviorFlag_STAT_DOT_BURN) {
			multiplier += 0.4
		}
		c.applyBurn(target, multiplier, chance)
	}

	// Main target
	c.engine.Attack(info.Attack{
		Key:        Skill,
		Targets:    []key.TargetID{target},
		Source:     c.id,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillMain[c.info.SkillLevelIndex()],
		},
		EnergyGain:   30,
		StanceDamage: 60,
	})

	// Adjacent targets
	c.engine.Attack(info.Attack{
		Key:        Skill,
		Targets:    c.engine.AdjacentTo(target),
		Source:     c.id,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillAdj[c.info.SkillLevelIndex()],
		},
		EnergyGain:   0,
		StanceDamage: 30,
	})

	c.engine.EndAttack()
}

// Skill's burn application method, uses Talent's burn application function
func (c *char) applyBurn(target key.TargetID, multiplier, chance float64) {
	applyBurn(c.engine, c.id, target, multiplier, chance)
}
