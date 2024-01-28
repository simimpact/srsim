package himeko

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	himeko_talent = "himeko-talent"
)

func init() {
	modifier.Register(himeko_talent, modifier.Config{
		Listeners: modifier.Listeners{},
	})
}

func (c *char) initTalent() {
	c.engine.Events().StanceBreak.Subscribe(c.breakListener)
}

func (c *char) breakListener(e event.StanceBreak) {
	targ, _ := c.engine.EnemyInfo(e.Target)

	// Himeko talent immediately maxes out if an elite/boss is broken
	if targ.Rank == model.EnemyRank_BIG_BOSS || targ.Rank == model.EnemyRank_ELITE || targ.Rank == model.EnemyRank_LITTLE_BOSS {
		c.talentStacks += 3
	} else {
		c.talentStacks++
	}

	if c.talentStacks >= 3 {

	}
}
