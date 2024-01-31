package himeko

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	ultimate  = "himeko-ult"
	ultEnergy = "himeko-ult-energy"
)

func init() {
	modifier.Register(ultimate, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: ultDeathListener,
		},
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   ultEnergy,
		Source: c.id,
	})

	c.engine.Attack(info.Attack{
		Key:        ultimate,
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

	// E6
	if c.info.Eidolon >= 6 {
		for i := 0; i < 2; i++ {
			c.engine.Retarget(info.Retarget{
				Targets: c.engine.Enemies(),
				Filter: func(target key.TargetID) bool {
					return true
				},
				Max:          1,
				IncludeLimbo: false,
			})
		}
	}

	c.engine.EndAttack()
	c.engine.RemoveModifier(c.id, ultEnergy)
}

func ultDeathListener(mod *modifier.Instance, target key.TargetID) {
	mod.Engine().ModifyEnergy(
		info.ModifyAttribute{
			Key:    ultimate,
			Target: mod.Owner(),
			Source: mod.Source(),
			Amount: 5,
		},
	)
}
