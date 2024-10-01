package luocha

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult = "luocha-ult"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	// Do E6
	if c.info.Eidolon >= 6 {
		for _, trg := range c.engine.Enemies() {
			c.engine.AddModifier(trg, info.Modifier{
				Name:   E6,
				Source: c.id,
				Stats:  info.PropMap{prop.AllDamageRES: 0.2},
			})
		}
	}

	// Dispel 1 buff on all enemies
	for _, trg := range c.engine.Enemies() {
		c.engine.DispelStatus(trg, info.Dispel{
			Status: model.StatusType_STATUS_BUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
	}

	// Add 1 stack of Abyss Flower if no Field active
	if !c.engine.HasModifier(c.id, Field) {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   AbyssFlower,
			Source: c.id,
		})
	}

	// Do damage
	c.engine.Attack(info.Attack{
		Key:        Ult,
		AttackType: model.AttackType_ULT,
		DamageType: model.DamageType_IMAGINARY,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		Targets:      c.engine.Enemies(),
		Source:       c.id,
		EnergyGain:   5,
		StanceDamage: 60,
	})
}
