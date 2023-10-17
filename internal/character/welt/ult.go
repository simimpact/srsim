package welt

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Deals Imaginary DMG equal to 150% of Welt's ATK to all enemies,
// with a 100% base chance for enemies hit by this ability to be Imprisoned for 1 turn.
// Imprisoned enemies have their actions delayed by 40% and SPD reduced by 10%.

const (
	Ult key.Attack = "welt-ult"
)

var attackHits = []float64{0.1, 0.9}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	// global AOE attack
	for _, target := range c.engine.Enemies() {
		// multiple hits attack.
		for i, hitRatio := range attackHits {
			c.engine.Attack(info.Attack{
				Key:        Ult,
				HitIndex:   i,
				HitRatio:   hitRatio,
				Targets:    []key.TargetID{target},
				Source:     c.id,
				AttackType: model.AttackType_ULT,
				DamageType: model.DamageType_IMAGINARY,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: ultAtk[c.info.UltLevelIndex()],
				},
				StanceDamage: 60,
				EnergyGain:   5,
			})
		}

		// imprison logic : DM uses common.Confine
		c.engine.AddModifier(target, info.Modifier{
			Name:   common.Imprisonment,
			Source: c.id,
			State: common.ImprisonState{
				SpeedDownRatio: 0.1,
				DelayRatio:     0.4,
			},
			Chance:   1,
			Duration: 1,
		})
	}
}
