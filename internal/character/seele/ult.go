package seele

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult key.Attack = "seele-ult"
)

// Enters the buffed state and deals Quantum DMG equal to 425% of her ATK to a single enemy.
func (c *char) Ult(target key.TargetID, state info.ActionState) {
	// enter buffed state w/o extra turns.
	c.enterResurgence(false)
	// ult attack.
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
	state.EndAttack()
	// E6 : After Ultimate, inflict the target enemy with Butterfly Flurry for 1 turn(s).
	if c.info.Eidolon >= 6 {
		c.engine.AddModifier(target, info.Modifier{
			Name:   E6Debuff,
			Source: c.id,
			State:  ult[c.info.UltLevelIndex()],
		})
	}
}
