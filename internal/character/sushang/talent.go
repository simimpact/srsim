package sushang

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TalentBuff key.Modifier = "sushang_talent_buff"
)

func init() {
	// Only for E0-E5
	modifier.Register(TalentBuff, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_BUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
	})
}

func (c *char) initTalent() {
	c.engine.Events().StanceBreak.Subscribe(func(e event.StanceBreakEvent) {
		c.addTalentBuff()
	})
}

func (c *char) addTalentBuff() {
	if c.info.Eidolon >= 6 {
		c.addE6Stack()
	} else {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   TalentBuff,
			Source: c.id,
			Stats: info.PropMap{
				prop.SPDPercent: talent[c.info.TalentLevelIndex()],
			},
			Duration: 2,
		})
	}
}
