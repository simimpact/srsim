package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	invigoration = "invigoration"
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
	// TODO :
	// replace old invigoration that's attached to bailu with event subscribers.
	// 1. make sure all c.addInvigoration applies invig to all team members.
	// 2. make sure to add onDeath subs to bailu to remove all active invigorations
	//    on team members.
	modifier.Register(invigoration, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterBeingHitAll: healOnBeingHit,
		},
		Stacking:          modifier.Prolong,
		CanModifySnapshot: true,
	})
}

func (c *char) initTalent() {
	// talent revive logics.
	reviveCountLeft := 1
	// E6 : Bailu can heal allies who received a killing blow 1 more time(s)
	//      in a single battle.
	if c.info.Eidolon >= 6 {
		reviveCountLeft = 2
	}
	revPercent := revivePercent[c.info.TalentLevelIndex()]
	revFlat := reviveFlat[c.info.TalentLevelIndex()]

	c.engine.Events().LimboWaitHeal.Subscribe(func(e event.LimboWaitHeal) bool {
		if e.IsCancelled || reviveCountLeft <= 0 || c.engine.IsEnemy(e.Target) {
			return false
		}
		// dispel and heal logics :
		reviveCountLeft--
		// all debuffs dispel.
		c.engine.DispelStatus(e.Target, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_FIRST_ADDED,
		})
		// "revive" heal
		c.addHeal(revive, revPercent, revFlat, []key.TargetID{e.Target})
		return true
	}, 1)

	// bailu death : remove all active invigoration modifiers
	c.engine.Events().TargetDeath.Subscribe((func(e event.TargetDeath) {
		if e.Target == c.id {
			for _, char := range c.engine.Characters() {
				c.engine.RemoveModifier(char, invigoration)
			}
		}
	}))
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

// TODO : change this logic to an event subscriber.
func energyOnRemove(mod *modifier.Instance) {
	// E1 : If the target ally's current HP is equal to their Max HP when
	// Invigoration ends, regenerates 8 extra Energy for this target.
	charInfo, _ := mod.Engine().CharacterInfo(mod.Owner())
	if charInfo.Eidolon < 1 {
		return
	}

	if mod.Engine().HPRatio(mod.Owner()) == 1.0 {
		mod.Engine().ModifyEnergy(info.ModifyAttribute{
			Key:    invigoration,
			Target: mod.Owner(),
			Source: mod.Source(),
			Amount: 8.0,
		})
	}
}

// adds invigoration re-heal with independent heal counters to chars.
func (c *char) addInvigoration(target key.TargetID, duration int) {
	// A4 : Invigoration can trigger 1 more time(s).
	healsLeft := 2
	if c.info.Traces["102"] {
		healsLeft = 3
	}
	// A6 : Characters with Invigoration receive 10% less DMG.
	dmgTakenAmt := 0.0
	if c.info.Traces["103"] {
		dmgTakenAmt = -0.1
	}

	c.engine.AddModifier(target, info.Modifier{
		// added mod is the target version(not one that's attached to bailu)
		Name:     invigoration,
		Source:   c.id,
		Duration: duration,
		State: invigStruct{
			healPercent: talentPercent[c.info.TalentLevelIndex()],
			healFlat:    talentFlat[c.info.TalentLevelIndex()],
			healsLeft:   healsLeft,
		},
		Stats: info.PropMap{prop.AllDamageTaken: dmgTakenAmt},
	})
}
