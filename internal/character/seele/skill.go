package seele

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SkillAttack    key.Attack   = "seele-skill"
	SkillSpeedBuff key.Modifier = "seele-skill-speed-up"
)

// Increases Seele's SPD by 25% for 2 turn(s)
// and deals Quantum DMG equal to 220% of Seele's ATK to a single enemy.

func init() {
	modifier.Register(SkillSpeedBuff, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_BUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
	})
}

var skillHits = []float64{0.2, 0.1, 0.1, 0.6}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// E2 : The SPD Boost effect of Seele's Skill can stack up to 2 time(s).
	maxCountAmt := 1.0
	countAddAmt := 0.0
	if c.info.Eidolon >= 2 {
		maxCountAmt = 2.0
		countAddAmt = 1.0
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:              SkillSpeedBuff,
		Source:            c.id,
		Duration:          2,
		Stats:             info.PropMap{prop.SPDPercent: 0.25},
		MaxCount:          maxCountAmt,
		CountAddWhenStack: countAddAmt,
	})

	// attack
	for i, hitRatio := range skillHits {
		c.engine.Attack(info.Attack{
			Key:        SkillAttack,
			HitIndex:   i,
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_QUANTUM,
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
}
