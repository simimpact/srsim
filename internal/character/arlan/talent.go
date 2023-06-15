package arlan

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
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
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: talentBeforeHitAll,
			OnAfterAction:  talentAfterAction,
		},
	})
}

func (c *char) addTalentMod() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
		State: talentState{
			maxBonusDamage: talent[c.info.AbilityLevel.Talent],
		},
	})
}
