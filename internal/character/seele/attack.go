package seele

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Normal key.Attack = "seele-normal"
	A6     key.Reason = "seele-a6"
)

var attackHits = []float64{0.3, 0.7}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range attackHits {
		c.engine.Attack(info.Attack{
			Key:        Normal,
			HitIndex:   i,
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_QUANTUM,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
			},
			StanceDamage: 30.0,
			EnergyGain:   20.0,
			HitRatio:     hitRatio,
		})
	}
	state.EndAttack()
	// A6 : After using a Basic ATK, Seele's next action will be Advanced Forward by 20%.
	if c.info.Traces["103"] {
		c.advanceForward(A6, state.IsInsert())
	}
}
