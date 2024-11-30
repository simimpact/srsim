package gallagher

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1 = "gallagher-e1"
	E2 = "gallagher-e2"
	E6 = "gallagher-e6"
)

func init() {
	modifier.Register(E1, modifier.Config{})

	modifier.Register(E2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		CanDispel:  true,
		Duration:   2,
	})

	modifier.Register(E6, modifier.Config{})
}

func (c *char) initEidolons() {
	c.engine.Events().BattleStart.Subscribe(c.E1Buffs)

	if c.info.Eidolon >= 6 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6,
			Source: c.id,
			Stats: info.PropMap{
				prop.BreakEffect:         0.2,
				prop.AllStanceDMGPercent: 0.2,
			},
		})
	}
}

func (c *char) E1Buffs(e event.BattleStart) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   E1,
		Source: c.id,
		Stats: info.PropMap{
			prop.EffectRES: 0.5,
		},
	})

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Target: c.id,
		Source: c.id,
		Key:    E1,
		Amount: 20,
	})
}
