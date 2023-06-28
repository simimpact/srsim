package gepard

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent key.Modifier = "gepard-talent"
)

type talentState struct {
	revivePerc float64
}

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnLimboWaitHeal: func(mod *modifier.ModifierInstance) bool {

				// Dispel all debuffs
				mod.Engine().DispelStatus(mod.Owner(), info.Dispel{
					Status: model.StatusType_STATUS_DEBUFF,
					Order:  model.DispelOrder_LAST_ADDED,
				})

				// Queue Heal
				mod.Engine().InsertAbility(info.Insert{
					Execute: func() {
						mod.Engine().SetHP(
							mod.Owner(), mod.Owner(), mod.OwnerStats().MaxHP()*mod.State().(talentState).revivePerc)
					},
					Source:   mod.Owner(),
					Priority: info.CharReviveSelf,
				})

				mod.RemoveSelf()
				return true
			},
		},
	})
}

func (c *char) talent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
		State: talentState{
			revivePerc: talent[c.info.TalentLevelIndex()],
		},
	})
}
