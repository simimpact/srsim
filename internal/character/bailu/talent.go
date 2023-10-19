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

// mostly used to track amount of heals left for each chars with invigoration
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
		Listeners: modifier.Listeners{
			OnAfterBeingHitAll: healOnBeingHit,
		},
		Stacking:          modifier.ReplaceBySource,
		CanModifySnapshot: true,
	})
}

func (c *char) initTalent() {
	// talent revive logics.
	// E6 : Bailu can heal allies who received a killing blow 1 more time(s)
	//      in a single battle.
	if c.info.Eidolon >= 6 {
		c.reviveLeft = 2
	}
	revPercent := revivePercent[c.info.TalentLevelIndex()]
	revFlat := reviveFlat[c.info.TalentLevelIndex()]

	cancelFlags := []model.BehaviorFlag{
		model.BehaviorFlag_STAT_CTRL,
		model.BehaviorFlag_DISABLE_ACTION,
	}

	// revive only if available and dying target is a character.
	c.engine.Events().LimboWaitHeal.Subscribe(func(e event.LimboWaitHeal) bool {
		if e.IsCancelled ||
			c.reviveLeft <= 0 ||
			!c.engine.IsCharacter(e.Target) ||
			// cannot revive if under CC
			c.engine.HasBehaviorFlag(c.id, cancelFlags...) {
			// TODO : confirm what return should be when conditions aren't met
			return false
		}
		c.reviveLeft--

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

// heals on being hit w/ heal limit. attached to invigoration modifier
func healOnBeingHit(mod *modifier.Instance, e event.HitEnd) {
	state := mod.State().(*invigStruct)
	// do nothing if there are no more invigoration heals
	if state.healsLeft <= 0 {
		return
	}
	mod.Engine().Heal(info.Heal{
		Key:         invigoration,
		Source:      mod.Source(),
		Targets:     []key.TargetID{mod.Owner()},
		BaseHeal:    info.HealMap{model.HealFormula_BY_HEALER_MAX_HP: state.healPercent},
		HealValue:   state.healFlat,
		UseSnapshot: true,
	})
	state.healsLeft--
}

// adds invigoration re-heal with independent heal counters to chars.
// attached to char struct for easier reuse in ult and technique.
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
