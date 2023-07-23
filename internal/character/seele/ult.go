package seele

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult key.Attack = "seele-ult"
)

// Seele enters the buffed state and deals Quantum DMG equal to 425% of her ATK to a single enemy.

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	// add buffedState mod
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   BuffedState,
		Source: c.id,
		Stats:  info.PropMap{prop.AllDamagePercent: talent[c.info.TalentLevelIndex()]},
	})

	// deal ult dmg
	c.engine.Attack(info.Attack{
		Key:        Ult,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_QUANTUM,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		StanceDamage: 90.0,
		EnergyGain:   5.0,
	})
}
