package seele

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Technique key.Attack = "seele-technique"
)

// After using her Technique, Seele gains Stealth for 20 second(s).
// While Stealth is active, Seele cannot be detected by enemies.
// And when entering battle by attacking enemies, Seele will immediately enter the buffed state.

// IMPL : ignore maze stealth. set to onBattleStart : add buffedState mod to seele.

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	// A4 : While Seele is in the buffed state, her Quantum RES PEN increases by 20%.
	resPenAmt := 0.0
	if c.info.Traces["102"] {
		resPenAmt = 0.2
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   BuffedState,
		Source: c.id,
		Stats: info.PropMap{
			prop.AllDamagePercent: talent[c.info.TalentLevelIndex()],
			prop.QuantumPEN:       resPenAmt,
		},
	})
	// add the minimal toughness damage atk? (observed effect, not in the DM?)
	c.engine.Attack(info.Attack{
		Key:        key.Attack(BuffedState),
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_QUANTUM,
		AttackType: model.AttackType_MAZE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.0,
		},
		StanceDamage: 60.0,
		EnergyGain:   0.0,
	})
}
