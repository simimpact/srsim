package herta

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill        = "herta-skill"
	skillHPCheck = "herta-skill-check"
)

func init() {
	modifier.Register(skillHPCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: skillBeforeHitListener,
		},
	},
	)
}

var hitsplit = []float64{0.3, 0.7}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   skillHPCheck,
		Source: c.id,
	})
	for index, ratio := range hitsplit {
		c.engine.Attack(info.Attack{
			Key:     Skill,
			Source:  c.id,
			Targets: c.engine.Enemies(),
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
			},
			DamageType:   model.DamageType_ICE,
			AttackType:   model.AttackType_SKILL,
			HitIndex:     index,
			HitRatio:     ratio,
			StanceDamage: 30,
			EnergyGain:   30,
		})
	}

	c.engine.RemoveModifier(c.id, skillHPCheck)

	c.engine.EndAttack()
}

func skillBeforeHitListener(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().HPRatio(e.Defender) >= 0.5 {
		a2increase := 0.0
		herta, _ := mod.Engine().CharacterInfo(e.Attacker)
		// A2
		if herta.Traces["101"] {
			a2increase = 0.25
		}

		e.Hit.Attacker.AddProperty(skillHPCheck, prop.AllDamagePercent, 0.2+a2increase)
	}
}
