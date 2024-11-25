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
	modifier.Register(TalentBuff, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_BUFF,
		CanDispel:     true,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
		Listeners: modifier.Listeners{
			OnAdd: talentOnAdd,
		},
		CountAddWhenStack: 1,
	})
}

func (c *char) initTalent() {
	c.engine.Events().StanceBreak.Subscribe(func(e event.StanceBreak) {
		c.addTalentBuff()
	})
}

func (c *char) addTalentBuff() {
	maxCount := 1.0
	if c.info.Eidolon >= 6 {
		maxCount = 2.0
	}

	c.engine.AddModifier(c.id, info.Modifier{
		Name:     TalentBuff,
		Source:   c.id,
		Duration: 2,
		State:    talent[c.info.TalentLevelIndex()],
		MaxCount: maxCount,
	})
}

func talentOnAdd(mod *modifier.Instance) {
	mod.AddProperty(prop.SPDPercent, mod.Count()*mod.State().(float64))
}
