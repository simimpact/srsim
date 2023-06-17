package qingque

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	if c.tiles[0] == 0 {
		c.drawTile()
		c.drawTile()
	}
}
