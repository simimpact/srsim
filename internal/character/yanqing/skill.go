package yanqing

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill key.Attack = "yanqing-skill"
)

var skillHits = []float64{0.25, 0.25, 0.25, 0.25}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range skillHits {
		c.engine.Attack(info.Attack{
			Key:          Skill,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_ICE,
			AttackType:   model.AttackType_SKILL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()]},
			StanceDamage: 60,
			EnergyGain:   30,
			HitRatio:     hitRatio,
		})
	}
	c.tryFollow(target)
	c.addTalent()
}
