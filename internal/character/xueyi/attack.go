package xueyi

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Normal key.Attack = "xueyi-normal"
)

var hitsplit = []float64{0.4, 0.6}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	for i, split := range hitsplit {
		c.engine.Attack(info.Attack{
			Targets:    []key.TargetID{target},
			Key:        Normal,
			Source:     c.id,
			AttackType: model.AttackType_NORMAL,
			DamageType: model.DamageType_QUANTUM,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
			},
			HitIndex:     i,
			HitRatio:     split,
			EnergyGain:   20,
			StanceDamage: 30,
		})
	}

	state.EndAttack()
}
