package xueyi

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent    = "xueyi-talent"
	TalentFua = "xueyi-talent-fua"
)

func init() {

}

func (c *char) initTalent() {
	if c.info.Eidolon >= 6 {
		c.stackReq = 6
	}
	c.engine.Events().StanceChange.Subscribe(c.handleStanceChange)
}

func (c *char) handleStanceChange(e event.StanceChange) {
	if e.Key == TalentFua {
		return
	}
	if e.Key == TalentFua {
		return
	}
	if e.Source != c.id {
		if c.engine.IsCharacter(e.Source) {
			c.incrementTalentStacks(1)
		}
		return
	}
	diff := e.OldStance - e.NewStance
	increment := 0
	if diff > 0 && diff <= 30 {
		increment = 1
	} else {
		increment = int(diff / 30)
	}
	c.incrementTalentStacks(increment)
}

func (c *char) incrementTalentStacks(amt int) {
	if amt > 8 {
		amt = 8
	}
	c.curStacks += amt
	if c.curStacks > c.stackReq+6 {
		c.curStacks = c.stackReq + 6
	}
	if c.curStacks >= c.stackReq {
		c.curStacks -= c.stackReq
		// Activate talent fua
		c.engine.InsertAbility(info.Insert{
			Key:        TalentFua,
			Source:     c.id,
			AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
			Priority:   info.CharInsertAttackSelf,
			Execute:    c.talentFua,
		})
	}
}

func (c *char) talentFua() {

}
