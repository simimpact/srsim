package danheng

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SkillCritCheck key.Modifier = "dan-heng-skill-crit-check"
	SkillSpeedDown key.Modifier = "dan-heng-skill-speed-down"
)

func init() {
	modifier.Register(SkillSpeedDown, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_DOWN},
	})

	modifier.Register(SkillCritCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterHit: func(mod *modifier.Instance, e event.HitEndEvent) {
				if e.IsCrit {
					slowAmt := mod.State().(float64)
					mod.Engine().AddModifier(e.Defender, info.Modifier{
						Name:     SkillSpeedDown,
						Source:   e.Attacker,
						Duration: 2,
						Chance:   1.0,
						Stats:    info.PropMap{prop.SPDPercent: -slowAmt},
					})
				}
			},
		},
	})
}

var skillHits = []float64{0.3, 0.15, 0.15, 0.4}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// add modifier which will check for crits and try to apply slow each hit
	slowAmt := 0.12
	if c.info.Eidolon >= 6 {
		slowAmt += 0.08
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   SkillCritCheck,
		Source: c.id,
		State:  slowAmt,
	})

	// 4 hits
	for _, hitRatio := range skillHits {
		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_WIND,
			AttackType: model.AttackType_SKILL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
			},
			StanceDamage: 60.0,
			EnergyGain:   30.0,
			HitRatio:     hitRatio,
		})
	}

	state.EndAttack()
	c.a4()

	// remove crit check mod
	c.engine.RemoveModifier(c.id, SkillCritCheck)
}
