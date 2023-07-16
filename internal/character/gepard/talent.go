package gepard

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent key.Modifier = "gepard-talent"
	Revive              = "gepard-revive"
	A4     key.Reason   = "gepard-a4"
	E6     key.Reason   = "gepard-e6"
)

type talentState struct {
	revivePerc float64
	a4Active   bool
	e6Active   bool
}

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnLimboWaitHeal: talentRevive,
		},
	})
}

func talentRevive(mod *modifier.Instance) bool {
	// Dispel all debuffs
	mod.Engine().DispelStatus(mod.Owner(), info.Dispel{
		Status: model.StatusType_STATUS_DEBUFF,
		Order:  model.DispelOrder_LAST_ADDED,
	})

	// Queue Heal
	mod.Engine().InsertAbility(info.Insert{
		Execute: func() {
			// Set HP to specified Percentage
			mod.Engine().SetHP(info.ModifyAttribute{
				Key:    Revive,
				Target: mod.Owner(),
				Source: mod.Owner(),
				Amount: mod.OwnerStats().MaxHP() * mod.State().(talentState).revivePerc,
			})

			// If A4, restore Energy to 100% (Energy Cost is 100)
			if mod.State().(talentState).a4Active {
				mod.Engine().ModifyEnergyFixed(info.ModifyAttribute{
					Key:    A4,
					Target: mod.Owner(),
					Source: mod.Owner(),
					Amount: 100,
				})
			}

			// If E6, action forward
			if mod.State().(talentState).e6Active {
				mod.Engine().SetGauge(info.ModifyAttribute{
					Key:    E6,
					Target: mod.Owner(),
					Source: mod.Owner(),
					Amount: 0,
				})
			}

			mod.RemoveSelf()
		},
		Key:        Revive,
		Source:     mod.Owner(),
		Priority:   info.CharReviveSelf,
		AbortFlags: nil,
	})

	return true
}

func (c *char) talent() {
	revivePerc := talent[c.info.TalentLevelIndex()]

	if c.info.Eidolon >= 6 {
		revivePerc += 0.5
	}

	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
		State: talentState{
			revivePerc: revivePerc,
			a4Active:   c.info.Traces["102"],
			e6Active:   c.info.Eidolon >= 6,
		},
	})
}
