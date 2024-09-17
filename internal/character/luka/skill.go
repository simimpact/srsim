package luka

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	LukaBleed            = "luka-bleed"
	Skill     key.Attack = "luka-skill"
)

func init() {
	modifier.Register(
		LukaBleed,
		modifier.Config{
			Stacking:      modifier.ReplaceBySource,
			Duration:      3,
			BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_DOT, model.BehaviorFlag_STAT_DOT_BLEED},
		},
	)
}

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
		Key:        Skill,
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
	state.EndAttack()
}

// Luka's bleed has custom logic that differs slightly from common bleeds
type BleedState struct {
	DamagePercentage    float64
	EnemyHealthRatioCap float64
}

func (c *char) applyBleed(target key.TargetID) {
	c.engine.AddModifier(target, info.Modifier{
		Name:   LukaBleed,
		Source: c.id,
		State: BleedState{
			DamagePercentage:    skillDotCap[c.info.SkillLevelIndex()],
			EnemyHealthRatioCap: 0.24,
		},
		Chance:   1,
		Duration: 3,
	})
}

// Implementation of: github.com/srsim/internal/global/common/triggerable_dot.go
func (b *BleedState) TriggerDot(dot info.Modifier, ratio float64, engine engine.Engine, target key.TargetID) {
	owner := engine.Stats(dot.Source)
	targetStats := engine.Stats(target)
	bleedDamage := b.EnemyHealthRatioCap * targetStats.MaxHP()
	skillCap := b.DamagePercentage * owner.ATK()
	if bleedDamage > (skillCap) {
		bleedDamage = skillCap
	}
	engine.Attack(
		info.Attack{
			Key:        LukaBleed,
			Source:     dot.Source,
			Targets:    []key.TargetID{target},
			AttackType: model.AttackType_DOT,
			DamageType: model.DamageType_PHYSICAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: 0,
			},
			DamageValue:  bleedDamage,
			AsPureDamage: true,
		},
	)
}
