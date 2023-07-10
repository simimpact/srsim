package qingque

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 key.Modifier = "qingque-a2"
	A6 key.Modifier = "qingque-a6"
)

// A2:
// 	Restores 1 Skill Point when using the Skill. This effect can only be triggered 1 time per battle.
// A4:
//	Using the Skill increases DMG Boost effect of attacks by an extra 10%.
// A6:
//	Qingque's SPD increases by 10% for 1 turn after using the Enhanced Basic ATK.

func init() {
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAction: A2ActionEndListener,
		},
	})
	modifier.Register(A6, modifier.Config{
		StatusType:    model.StatusType_STATUS_BUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
	})
}

func (c *char) initTraces() {
	if c.info.Traces["101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
		})
	}
}

func A2ActionEndListener(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType == model.AttackType_SKILL {
		mod.Engine().ModifySP(1)
		mod.RemoveSelf()
	}
}

func (c *char) a6() {
	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:     A6,
			Source:   c.id,
			Stats:    info.PropMap{prop.SPDPercent: 0.1},
			Duration: 1,
		})
	}
}
