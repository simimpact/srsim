package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 = "sparkle-a2"
	A6 = "sparkle-a6"
)

var (
	buffMap = map[int]float64{1: 0.05, 2: 0.15, 3: 0.3, 0: 0.0}
)

func init() {
	modifier.Register(A6, modifier.Config{
		CanDispel:  true,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) initTraces() {
	c.engine.Events().BattleStart.Subscribe(c.initA6)
}

func (c *char) initA6(e event.BattleStart) {
	if c.info.Traces["103"] {
		qua := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Characters(),
			Filter: func(target key.TargetID) bool {
				charinfo, _ := c.engine.CharacterInfo(target)
				return charinfo.Element == model.DamageType_QUANTUM
			},
			IncludeLimbo: false,
		})

		quacount := len(qua)
		if quacount < 0 {
			quacount = 0
		} else if quacount > 3 {
			quacount = 3
		}

		for _, char := range c.engine.Characters() {
			charinfo, _ := c.engine.CharacterInfo(char)
			atkBuff := 0.15
			if charinfo.Element == model.DamageType_QUANTUM {
				atkBuff += buffMap[quacount]
			}
			c.engine.AddModifier(char, info.Modifier{
				Name:   A6,
				Source: c.id,
				Stats: info.PropMap{
					prop.ATKPercent: atkBuff,
				},
			})
		}
	}
}
