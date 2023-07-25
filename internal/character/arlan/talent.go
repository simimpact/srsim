package arlan

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent     key.Modifier = "arlan-talent"
	TalentBuff key.Modifier = "arlan-talent-buff"
)

type talentState struct {
	maxBonusDamage float64
}

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: addRemoveCheck,
			OnHPChange: func(mod *modifier.Instance, e event.HPChange) {
				addRemoveCheck(mod)
			},
		},
	})

	modifier.Register(TalentBuff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd: buffUpdate,
			OnHPChange: func(mod *modifier.Instance, e event.HPChange) {
				buffUpdate(mod)
			},
		},
	})
}

func addRemoveCheck(mod *modifier.Instance) {
	hp := mod.Engine().HPRatio(mod.Owner())
	if hp >= 1.0 {
		mod.Engine().RemoveModifier(mod.Owner(), TalentBuff)
	} else {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   TalentBuff,
			Source: mod.Owner(),
			State:  mod.State(),
		})
	}
}

func buffUpdate(mod *modifier.Instance) {
	addedBonusDamage := (1 - mod.Engine().HPRatio(mod.Owner())) * mod.State().(talentState).maxBonusDamage
	mod.SetProperty(prop.AllDamagePercent, addedBonusDamage)
}

func (c *char) addTalentMod() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
		State: talentState{
			maxBonusDamage: talent[c.info.TalentLevelIndex()],
		},
	})
}
