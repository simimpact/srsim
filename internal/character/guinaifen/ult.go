package guinaifen

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Ult = "guinaifen-ult"

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Ult,
		Targets:    c.engine.Enemies(),
		Source:     c.id,
		AttackType: model.AttackType_ULT,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		EnergyGain:   5,
		StanceDamage: 60,
	})

	// triggers all Burns on all enemies
	for _, trg := range c.engine.Enemies() {
		for _, dot := range c.engine.GetModifersByBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_BURN) {
			dot.State.(common.TriggerableDot).TriggerDot(dot, ultDetonateBurn[c.info.UltLevelIndex()], c.engine, trg)
		}
	}

	c.engine.EndAttack()
}
