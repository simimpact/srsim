package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	vuln = "luka-ult-vuln"
)

func init() {
	modifier.Register(
		vuln, modifier.Config{
			Stacking:   modifier.ReplaceBySource,
			Duration:   3,
			StatusType: model.StatusType_STATUS_DEBUFF,
		},
	)
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.e1Check(target)

	c.incrementFightingSprit()
	c.engine.AddModifier(target, info.Modifier{
		Name:     vuln,
		Source:   c.id,
		Duration: 3,
		Stats: info.PropMap{
			prop.AllDamageTaken: ultVuln[c.info.UltLevelIndex()],
		},
	})
}
