package danheng

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// After Dan Heng uses his Technique, his ATK increases by 40% at the start
// of the next battle for 3 turn(s).

const (
	Technique key.Modifier = "dan-heng-technique"
)

func init() {
	modifier.Register(Technique, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:     Technique,
		Source:   c.id,
		Duration: 3,
		Stats:    info.PropMap{prop.ATKPercent: 0.4},
	})
}
