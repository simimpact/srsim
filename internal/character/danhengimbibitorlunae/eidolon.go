package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E6Count  = "danhengimbibitorlunae-e6count"
	E6Effect = "danhengimbibitorlunae-e6effect"
)

// E1 increase talent max stack by 4 and gain 1 extra stack when hit
// E2 100% forward and 1 more ult stack
// E4 skill buff 1 more turn
// E6 20% imaginary pen for attack3 when ally use ult, max 3 stack

func init() {
	modifier.Register(E6Count, modifier.Config{
		StatusType:        model.StatusType_STATUS_BUFF,
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          3,
		CountAddWhenStack: 1,
	})
	modifier.Register(E6Effect, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) E6ActionEndListener(e event.ActionEnd) {
	if e.AttackType == model.AttackType_ULT && e.Owner != c.id {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6Count,
			Source: c.id,
		})
	}
}
