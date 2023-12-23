package welt

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Deals Imaginary DMG equal to 150% of Welt's ATK to all enemies,
// with a 100% base chance for enemies hit by this ability to be Imprisoned for 1 turn.
// Imprisoned enemies have their actions delayed by 40% and SPD reduced by 10%.

const (
	Ult key.Attack   = "welt-ult"
	A2  key.Modifier = "welt-a2"
)

func init() {
	modifier.Register(A2, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
	})
}

var attackHits = []float64{0.1, 0.9}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	// A2 : When using Ultimate, there is a 100% base chance to
	// 		increase the DMG received by the targets by 12% for 2 turn(s).
	if c.info.Traces["101"] {
		for _, target := range c.engine.Enemies() {
			c.engine.AddModifier(target, info.Modifier{
				Name:     A2,
				Source:   c.id,
				Chance:   1,
				Duration: 2,
				Stats:    info.PropMap{prop.AllDamageTaken: 0.12},
			})
		}
	}

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
		state.EndAttack()

		// imprison logic : DM uses common.Confine
		c.engine.AddModifier(target, info.Modifier{
			Name:   common.Imprisonment,
			Source: c.id,
			State: common.ImprisonState{
				SpeedDownRatio: 0.1,
				DelayRatio:     ultDelay[c.info.UltLevelIndex()],
			},
			Chance:   1,
			Duration: 1,
		})
	}
}
