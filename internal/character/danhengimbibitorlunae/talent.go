package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent = "danhengimbibitorlunae-talent" // MAvatar_DanHengIL_00_AllDamageTypeAddedRatio
)

type talentState struct {
	damageBoost float64
}

func (c *char) initTalent() {
	modifier.Register(Talent, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		MaxCount:   6,
		Listeners: modifier.Listeners{
			OnAdd: talentOnAdd,
		},
		CountAddWhenStack: 1,
	})
}

func (c *char) AddTalent() {
	maxcount := 6.0
	countadd := 1.0
	// if E1,maxcount+4 and 1 more stack for each hit
	if c.info.Eidolon >= 1 {
		maxcount = 10.0
		countadd = 2.0
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:              Talent,
		Source:            c.id,
		TickImmediately:   true,
		Duration:          1,
		MaxCount:          maxcount,
		CountAddWhenStack: countadd,
		State: talentState{
			damageBoost: talent[c.info.TalentLevelIndex()],
		},
	})
}

func talentOnAdd(mod *modifier.Instance) {
	state := mod.State().(talentState)
	mod.SetProperty(prop.AllDamagePercent, mod.Count()*(state.damageBoost))
}
