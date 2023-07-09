package clara

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Technique key.Modifier = "clara-technique-aggro"
)

func init() {
	modifier.Register(Technique, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   2,
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Technique,
		Source: c.id,
		Stats:  info.PropMap{prop.AggroPercent: 5},
	})
}
