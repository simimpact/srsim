package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SparkleTalent = "red-herring"
)

func init() {
	modifier.Register(SparkleTalent, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		MaxCount: 3,
		Listeners: modifier.Listeners{
			OnAdd:    adjustBuff,
			OnRemove: removeE2,
		},
		CanDispel:  true,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

type talentState struct {
	DmgPercentPerStack float64
	isE2               bool
}

func (c *char) initTalent() {
	c.engine.Events().SPChange.Subscribe(c.talent)
	c.engine.Events().ModifierAdded.Subscribe(c.adjustTalentBuff)
	c.engine.Events().ModifierRemoved.Subscribe(c.revertTalentBuff)
	//TODO: Modify maximum sp count once the api for doing so is defined.
}

func (c *char) talent(e event.SPChange) {
	spChange := e.OldSP - e.NewSP
	if spChange > 0 && c.engine.IsCharacter(e.Source) {
		for _, teammate := range c.engine.Characters() {
			c.engine.AddModifier(teammate, info.Modifier{
				Name:     SparkleTalent,
				Source:   c.id,
				Duration: 2,
				Count:    float64(spChange),
				State: talentState{
					DmgPercentPerStack: talent[c.info.TalentLevelIndex()],
					isE2:               c.info.Eidolon >= 2,
				},
			})
		}
	}
}

func adjustBuff(mod *modifier.Instance) {
	curstate := mod.State().(talentState)
	// TODO: Maybe refactor the two buffs to be tied to the same modifier like in dm? end res is same but might be more readable
	mod.SetProperty(prop.AllDamagePercent, curstate.DmgPercentPerStack*mod.Count())
	if curstate.isE2 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:      E2,
			Source:    mod.Source(),
			CanDispel: false,
			Count:     mod.Count(),
		})
	}
}

func (c *char) adjustTalentBuff(e event.ModifierAdded) {
	if e.Modifier.Name == Cipher && e.Modifier.Source == c.id {
		for _, char := range c.engine.Characters() {
			curstacks := c.engine.ModifierStackCount(char, c.id, SparkleTalent)
			c.engine.AddModifier(char, info.Modifier{
				Name:     SparkleTalent,
				Source:   c.id,
				Duration: 2,
				Count:    curstacks,
				State: talentState{
					DmgPercentPerStack: talent[c.info.TalentLevelIndex()] + ultimate[c.info.UltLevelIndex()],
					isE2:               c.info.Eidolon >= 2,
				},
			})
		}
	}
}

func (c *char) revertTalentBuff(e event.ModifierRemoved) {
	if e.Modifier.Name == Cipher && e.Modifier.Source == c.id {
		for _, char := range c.engine.Characters() {
			curstacks := c.engine.ModifierStackCount(char, c.id, SparkleTalent)
			c.engine.AddModifier(char, info.Modifier{
				Name:     SparkleTalent,
				Source:   c.id,
				Duration: 2,
				Count:    curstacks,
				State: talentState{
					DmgPercentPerStack: talent[c.info.TalentLevelIndex()],
					isE2:               c.info.Eidolon >= 2,
				},
			})
		}
	}
}

func removeE2(mod *modifier.Instance) {
	for _, char := range mod.Engine().Characters() {
		mod.Engine().RemoveModifier(char, E2)
	}
}
