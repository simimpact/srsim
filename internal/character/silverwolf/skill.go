package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SkillResDown  key.Modifier = "silverwolf-skill-res-down"
	SkillWeakType key.Modifier = "silverwolf-skill-weak-type"
)

func init() {
	modifier.Register(SkillResDown, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
	})

	modifier.Register(SkillWeakType, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				types := info.WeaknessMap{}
				for _, char := range mod.Engine().Characters() {
					info, _ := mod.Engine().CharacterInfo(char)
					types[info.Element] = true
				}
				for t := model.DamageType_PHYSICAL; t <= model.DamageType_IMAGINARY; t++ {
					if mod.OwnerStats().IsWeakTo(t) {
						delete(types, t)
					}
				}
				keys := []model.DamageType{}
				for t := range types {
					keys = append(keys, t)
				}
				mod.AddWeakness(keys[mod.Engine().Rand().Intn(len(keys))])
			},
		},
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(target, info.Modifier{
		Name:     SkillResDown,
		Source:   c.id,
		Duration: 2,
		Chance:   1,
		Stats:    info.PropMap{prop.AllDamageRES: -skillResDown[c.info.AbilityLevel.Skill-1]},
	})

	c.engine.AddModifier(target, info.Modifier{
		Name:     SkillWeakType,
		Source:   c.id,
		Duration: 2,
		Chance:   skillChance[c.info.AbilityLevel.Skill-1],
	})

	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_QUANTUM,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.AbilityLevel.Skill-1],
		},
		StanceDamage: 60.0,
		EnergyGain:   30.0,
	})

	state.EndAttack()
}
