package march7th

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	MarchE2 key.Reason = "march7th-e2"
)

func init() {

}

func (c *char) initEidolons() {

}

func (c *char) e1() {
	for _, target := range c.engine.Enemies() {
		if c.engine.HasModifier(target, MarchUltFreeze) {
			c.engine.ModifyEnergy(info.ModifyAttribute{
				Key:    MarchE2,
				Source: c.id,
				Target: target,
				Amount: 5,
			})
		}
	}
}

func (c *char) e2(target key.TargetID) {

}
