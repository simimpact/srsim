package gallagher

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var (
	basicHitSplit = []float64{0.5, 0.5}
	ebaHitSplit   = []float64{0.25, 0.15, 0.6}
)

const (
	Normal       key.Attack   = "gallgher-normal"
	NectarBlitz  key.Attack   = "gallagher-nectar-blitz"
	AtkReduction key.Modifier = "gallagher-atk-reduction"
)

func init() {
	modifier.Register(AtkReduction, modifier.Config{
		StatusType: model.StatusType_STATUS_DEBUFF,
		Stacking:   modifier.ReplaceBySource,
		Duration:   2,
	})
}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	if c.isEnhanced {
		c.EnhancedBasic(target, state)
	} else {
		c.Basic(target, state)
	}
}

func (c *char) Basic(target key.TargetID, state info.ActionState) {
	for index, ratio := range basicHitSplit {
		c.engine.Attack(
			info.Attack{
				Source:   c.id,
				Targets:  []key.TargetID{target},
				Key:      Normal,
				HitIndex: index,
				HitRatio: ratio,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
				},
				AttackType:   model.AttackType_NORMAL,
				DamageType:   model.DamageType_FIRE,
				StanceDamage: 30,
				EnergyGain:   20,
			},
		)
	}

	state.EndAttack()
}

func (c *char) EnhancedBasic(target key.TargetID, state info.ActionState) {
	for index, ratio := range ebaHitSplit {
		c.engine.Attack(
			info.Attack{
				Source:   c.id,
				Targets:  []key.TargetID{target},
				Key:      NectarBlitz,
				HitIndex: index,
				HitRatio: ratio,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: enhancedBasic[c.info.AttackLevelIndex()],
				},
				AttackType:   model.AttackType_NORMAL,
				DamageType:   model.DamageType_FIRE,
				StanceDamage: 90,
				EnergyGain:   20,
			},
		)
	}
	state.EndAttack()

	c.engine.AddModifier(target, info.Modifier{
		Name:     AtkReduction,
		Source:   c.id,
		Duration: 2,
		Stats: info.PropMap{
			prop.ATKPercent: -ebaAtkReduction[c.info.AttackLevelIndex()],
		},
	})

	c.isEnhanced = false
}
