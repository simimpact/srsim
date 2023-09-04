package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// E6 : Bailu can heal allies who received a killing blow 1 more time(s) in a single battle.

const (
	E2 key.Modifier = "bailu-e2"
	E4 key.Modifier = "bailu-e4"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(E4, modifier.Config{
		Stacking:          modifier.Replace,
		MaxCount:          3,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_BUFF,
	})
}

func (c *char) initEidolons() {
	// E4 : Every healing provided by the Skill makes the recipient deal 10% more DMG for 2 turn(s).
	// This effect can stack up to 3 time(s).
	c.engine.Events().HealEnd.Subscribe(func(e event.HealEnd) {
		if c.info.Eidolon < 4 {
			return
		}
		if e.Healer == c.id && e.Key == Skill {
			c.engine.AddModifier(e.Target, info.Modifier{
				Name:     E4,
				Source:   c.id,
				Stats:    info.PropMap{prop.AllDamagePercent: 0.1},
				Duration: 2,
			})
		}
	})
}
