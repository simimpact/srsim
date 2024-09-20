package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	a2 key.Reason = "luka-a2"
)

func (c *char) initTalent() {
	// Luka should not stack e4 with this initial stack
	c.fightingSpirit = 1
}

func (c *char) incrementFightingSpritBy(amt int) {
	addValue := 0
	if c.fightingSpirit >= 3 {
		addValue = 4 - c.fightingSpirit
	} else {
		addValue = amt
	}
	c.fightingSpirit += amt
	if c.fightingSpirit > 4 {
		c.fightingSpirit = 4
	}

	if c.info.Eidolon >= 4 && addValue > 0 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   e4,
			Source: c.id,
			Count:  float64(amt),
		})
	}

	if c.info.Traces["102"] && addValue > 0 {
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    a2,
			Target: c.id,
			Source: c.id,
			Amount: float64(3 * addValue),
		})
	}
}
