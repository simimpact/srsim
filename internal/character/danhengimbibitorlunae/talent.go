package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent = "danhengimbibitorlunae-talent"
)

type talentState struct {
	damageBoost float64
}

func (c *char) init() {
	maxcount := 6.0
	countadd := 1.0
	if c.info.Eidolon >= 1 {
		maxcount = 10.0
		countadd = 2.0
	}
	modifier.Register(Talent, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		MaxCount:   maxcount,
		Listeners: modifier.Listeners{
			OnAdd:    talentOnAdd,
			OnPhase2: talentOnPhase2,
		},
		CountAddWhenStack: countadd,
	})
}

func (c *char) AddTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
		State: talentState{
			damageBoost: talent[c.info.TalentLevelIndex()],
		},
	})
}

func talentOnAdd(mod *modifier.Instance) {
	state := mod.State().(talentState)
	mod.SetProperty(prop.AllDamagePercent, mod.Count()*(state.damageBoost))
}
func talentOnPhase2(mod *modifier.Instance) {
	mod.RemoveSelf()
}
