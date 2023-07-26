package blade

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Technique key.Attack = "blade-technique"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.ModifyHPByRatio(info.ModifyHPByRatio{
		Key:       key.Reason(Technique),
		Target:    c.id,
		Source:    c.id,
		Ratio:     -0.2,
		RatioType: model.ModifyHPRatioType_MAX_HP,
		Floor:     1,
	})

	c.engine.Attack(info.Attack{
		Key:        Technique,
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_WIND,
		AttackType: model.AttackType_MAZE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_MAX_HP: 0.4,
		},
		StanceDamage: 60.0,
		EnergyGain:   0.0,
	})
}
