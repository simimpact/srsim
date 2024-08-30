package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	a2 key.Reason = "luka-a2"
)

func init() {

}

func (c *char) initTalent() {
	c.fightingSpirit = 1
}

func (c *char) incrementFightingSprit() {
	c.fightingSpirit += 1

	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   e4,
			Source: c.id,
		})
	}

	if c.info.Traces["102"] {
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    a2,
			Target: c.id,
			Source: c.id,
			Amount: 3,
		})
	}
}
