package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	e1 = "luka-e1"
	e2 = "luka-e2"
	e4 = "luka-e4"
)

func init() {
	modifier.Register(e1, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		Count:      1,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})

	modifier.Register(e4, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		CountAddWhenStack: 1,
		MaxCount:          4,
		Listeners: modifier.Listeners{
			OnAdd: stackAttack,
		},
	})
}

func (c *char) e1Check(target key.TargetID) {
	if c.info.Eidolon >= 1 && c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_BLEED) {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   e1,
			Source: c.id,
			Stats: info.PropMap{
				prop.AllDamagePercent: 0.15,
			},
			Duration: 2,
		})
	}
}

func stackAttack(mod *modifier.Instance) {
	mod.SetProperty(prop.ATKConvert, mod.Count()*0.05)
}
