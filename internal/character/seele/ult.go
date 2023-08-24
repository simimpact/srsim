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

// Enters the buffed state and deals Quantum DMG equal to 425% of her ATK to a single enemy.
func (c *char) Ult(target key.TargetID, state info.ActionState) {
	// A4 : While Seele is in the buffed state, her Quantum RES PEN increases by 20%.
	penAmt := 0.0
	if c.info.Traces["102"] {
		penAmt = 0.2
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Resurgence,
		Source: c.id,
		Stats: info.PropMap{
			prop.AllDamagePercent: talent[c.info.TalentLevelIndex()],
			prop.QuantumPEN:       penAmt,
		},
	})

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
