package serval

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	ServalTalent key.Attack = "serval-talent"
)

func onAfterHit(c *char, target key.TargetID, state info.ActionState) {
	if c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_ELECTRIC) {
		c.engine.Attack(info.Attack{
			Key:        ServalTalent,
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_THUNDER,
			AttackType: model.AttackType_PURSUED,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: talent[c.info.SkillLevelIndex()],
			},
			StanceDamage: 0.0,
			EnergyGain:   0.0,
		})
	}
}
