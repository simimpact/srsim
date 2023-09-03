package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	invigoration = "invigoration"
	invigLocal   = "invigoration-local"
	revive       = "bailu-revive"
)

type invigStruct struct {
	healPercent, healFlat float64
	healsLeft             int
}

// After an ally with Invigoration is hit, restores the ally's HP
// for 5.4% of Bailu's Max HP plus 144. This effect can trigger 2 time(s).
// When an ally receives a killing blow, they will not be knocked down.
// Bailu immediately heals the ally for 18% of Bailu's Max HP plus 480 HP.
// This effect can be triggered 1 time per battle

func init() {
	modifier.Register(invigoration, modifier.Config{
		Stacking: modifier.Prolong,
		Listeners: modifier.Listeners{
			OnAdd: addInvigLocal,
		},
	})
	modifier.Register(invigLocal, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterBeingHitAll: healOnBeingHit,
		},
		Stacking: modifier.Replace,
	})
}

func (c *char) initTalent() {
	// team revive logic
	reviveCountLeft := 1
	revPercent := revivePercent[c.info.TalentLevelIndex()]
	revFlat := reviveFlat[c.info.TalentLevelIndex()]
	c.engine.Events().LimboWaitHeal.Subscribe(func(e event.LimboWaitHeal) bool {
		if e.IsCancelled || reviveCountLeft <= 0 {
			return false
		}
		c.addHeal(revive, revPercent, revFlat, []key.TargetID{e.Target})
		reviveCountLeft--
		return true
	}, 1)
}

// used to set dynamic local (per-character) value for invigoration heal trigger count.
func addInvigLocal(mod *modifier.Instance) {
	state := mod.State().(invigStruct)
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   invigLocal,
		Source: mod.Source(),
		State:  &state,
	})
}

func healOnBeingHit(mod *modifier.Instance, e event.HitEnd) {
	state := mod.State().(*invigStruct)
	if state.healsLeft < 0 {
		mod.RemoveSelf()
	}
	mod.Engine().Heal(info.Heal{
		Key:       invigoration,
		Source:    mod.Source(),
		Targets:   []key.TargetID{mod.Owner()},
		BaseHeal:  info.HealMap{model.HealFormula_BY_HEALER_MAX_HP: state.healPercent},
		HealValue: state.healFlat,
	})
	state.healsLeft--
}

func (c *char) addInvigoration(target key.TargetID, duration int) {
	c.engine.AddModifier(target, info.Modifier{
		Name:     invigoration,
		Source:   c.id,
		Duration: duration,
		State: invigStruct{
			healPercent: talentPercent[c.info.TalentLevelIndex()],
			healFlat:    talentFlat[c.info.TalentLevelIndex()],
			healsLeft:   2,
		},
	})
}
