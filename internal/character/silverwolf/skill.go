package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine"
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

var damageTypeToResProperty = map[model.DamageType]prop.Property{
	model.DamageType_PHYSICAL:  prop.PhysicalDamageRES,
	model.DamageType_FIRE:      prop.FireDamageRES,
	model.DamageType_ICE:       prop.IceDamageRES,
	model.DamageType_THUNDER:   prop.ThunderDamageRES,
	model.DamageType_WIND:      prop.WindDamageRES,
	model.DamageType_QUANTUM:   prop.QuantumDamageRES,
	model.DamageType_IMAGINARY: prop.ImaginaryDamageRES,
}

func init() {
	modifier.Register(SkillResDown, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
	})

	modifier.Register(SkillWeakType, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_ATTACH_WEAKNESS},
		Stacking:      modifier.Replace,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				dmgType, ok := chooseWeaknessType(mod.Engine(), mod.Owner())
				if !ok {
					mod.RemoveSelf()
					return
				}
				mod.AddWeakness(dmgType)
				mod.SetProperty(damageTypeToResProperty[dmgType], -0.2)
			},
		},
	})
}

func chooseWeaknessType(engine engine.Engine, target key.TargetID) (model.DamageType, bool) {
	types := info.WeaknessMap{}
	// add team damage types
	for _, char := range engine.Characters() {
		info, _ := engine.CharacterInfo(char)
		types[info.Element] = true
	}
	// remove enemy damage types
	for t := model.DamageType_PHYSICAL; t <= model.DamageType_IMAGINARY; t++ {
		if engine.Stats(target).IsWeakTo(t) {
			types.Remove(t)
		}
	}
	keys := []model.DamageType{}
	for t := range types {
		keys = append(keys, t)
	}
	if len(keys) == 0 {
		return model.DamageType_INVALID_DAMAGE_TYPE, false
	}
	return keys[engine.Rand().Intn(len(keys))], true
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// A6:
	//	If there are 3 or more debuff(s) affecting the enemy when the Skill is
	//	used, then the Skill decreases the enemy's All-Type RES by an additional
	//	3%.
	allDamageDown := -skillResDown[c.info.SkillLevelIndex()]
	if c.info.Traces["103"] && c.engine.ModifierStatusCount(target, model.StatusType_STATUS_DEBUFF) >= 3 {
		allDamageDown -= 0.03
	}
	c.engine.AddModifier(target, info.Modifier{
		Name:     SkillResDown,
		Source:   c.id,
		Duration: 2,
		Chance:   1,
		Stats:    info.PropMap{prop.AllDamageRES: allDamageDown},
	})

	// A4:
	//	The duration of the Weakness implanted by Silver Wolf's Skill increases
	//	by 1 turn(s).
	duration := 2
	if c.info.Traces["102"] {
		duration += 1
	}
	c.engine.AddModifier(target, info.Modifier{
		Name:     SkillWeakType,
		Source:   c.id,
		Duration: duration,
		Chance:   skillChance[c.info.SkillLevelIndex()],
	})

	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_QUANTUM,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		StanceDamage: 60.0,
		EnergyGain:   30.0,
	})

	state.EndAttack()
}
