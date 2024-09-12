package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	ultimate = "luka-ult"
	vuln     = "luka-ult-vuln"
)

func init() {
	modifier.Register(
		vuln, modifier.Config{
			Stacking:   modifier.ReplaceBySource,
			Duration:   3,
			StatusType: model.StatusType_STATUS_DEBUFF,
		},
	)
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.e1Check(target)

	c.incrementFightingSprit()
	c.engine.AddModifier(target, info.Modifier{
		Name:     vuln,
		Source:   c.id,
		Duration: 3,
		Stats: info.PropMap{
			prop.AllDamageTaken: ultVuln[c.info.UltLevelIndex()],
		},
	})

	c.engine.Attack(info.Attack{
		Key:          ultimate,
		Targets:      []key.TargetID{target},
		Source:       c.id,
		AttackType:   model.AttackType_ULT,
		DamageType:   model.DamageType_PHYSICAL,
		StanceDamage: 90,
		EnergyGain:   5,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
	})

	state.EndAttack()
}
