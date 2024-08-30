package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Normal         key.Attack = "luka-normal"
	EnhancedNormal key.Attack = "luka-enhanced-normal"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	c.e1Check(target)
	if c.fightingSpirit < 2 {
		c.basicAttack(target, state)
	} else {
		c.enhancedBasic(target, state)
	}
}

func (c *char) enhancedBasic(target key.TargetID, state info.ActionState) {

}

func (c *char) basicAttack(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Normal,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_PHYSICAL,
		Source:     c.id,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30,
		EnergyGain:   20,
	})

	state.EndAttack()
}
