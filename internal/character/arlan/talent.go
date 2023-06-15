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
	Talent key.Modifier = "arlan-talent"
)

type talentState struct {
	maxBonusDamage float64
}

func init() {
	modifier.Register(Talent, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			// set bonus damage on combat start
			OnAdd: func(mod *modifier.ModifierInstance) {
				addedBonusDamage := (1 - mod.Engine().HPRatio(mod.Owner())) * mod.State().(talentState).maxBonusDamage
				mod.SetProperty(prop.AllDamagePercent, addedBonusDamage)
			},

			// update bonus damage based on new HP
			OnHPChange: func(mod *modifier.ModifierInstance, e event.HPChangeEvent) {
				addedBonusDamage := (1 - mod.Engine().HPRatio(mod.Owner())) * mod.State().(talentState).maxBonusDamage
				mod.SetProperty(prop.AllDamagePercent, addedBonusDamage)
			},
		},
	})

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
