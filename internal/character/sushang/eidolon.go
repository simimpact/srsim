package sushang

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2 key.Modifier = "sushang-e2"
	E4 key.Modifier = "sushang-e4"
	E6 key.Modifier = "sushang-e6"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E4, modifier.Config{})

	modifier.Register(E6, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		StatusType:        model.StatusType_STATUS_BUFF,
		BehaviorFlags:     []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
		MaxCount:          2,
		CountAddWhenStack: 1,
		Listeners: modifier.Listeners{
			OnAdd: e6OnAdd,
		},
	})
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E4,
			Source: c.id,
			Stats: info.PropMap{
				prop.BreakEffect: 0.4,
			},
		})
	}

	if c.info.Eidolon >= 6 {
		c.addE6Stack()
	}
}

func (c *char) e2() {
	if c.info.Eidolon >= 2 {
		if !c.engine.HasModifier(c.id, E2) {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   E2,
				Source: c.id,
				Stats: info.PropMap{
					prop.AllDamageReduce: 0.2,
				},
				Duration: 1,
			})
		}
	}
}

func (c *char) addE6Stack() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:     E6,
		Source:   c.id,
		Duration: 2,
		State:    talent[c.info.TalentLevelIndex()],
	})
}

func e6OnAdd(mod *modifier.ModifierInstance) {
	mod.AddProperty(prop.SPDPercent, mod.Count()*mod.State().(float64))
}
