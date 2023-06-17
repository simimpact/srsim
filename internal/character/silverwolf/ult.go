package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	UltDefDown key.Modifier = "silverwolf-ult-def-down"
)

func init() {
	modifier.Register(UltDefDown, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_DEF_DOWN},
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		TickMoment:    modifier.ModifierPhase1End,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(target, info.Modifier{
		Name:            UltDefDown,
		Source:          c.id,
		Duration:        3,
		Chance:          ultChance[c.info.UltLevelIndex()],
		Stats:           info.PropMap{prop.DEFPercent: -ultDefDown[c.info.UltLevelIndex()]},
		TickImmediately: true,
	})

	c.engine.Attack(info.Attack{
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

	c.e4(target)
	c.e1(target)
}
