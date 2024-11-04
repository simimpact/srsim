package gallagher

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Besotted key.Modifier = "gallagher-besotted"
)

func init() {
	modifier.Register(Besotted, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Duration: 2,
		Listeners: modifier.Listeners{
			OnBeforeBeingHitAll:  besottedBreakVuln,
			OnAfterBeingAttacked: besottedHeal,
		},
	})
}

type BesottedState struct {
	a6Active  bool
	breakVuln float64
	healAmt   float64
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	besottedDur := 2
	if c.info.Eidolon >= 4 {
		besottedDur += 1
	}

	for _, enemy := range c.engine.Enemies() {
		c.engine.AddModifier(enemy, info.Modifier{
			Name: Besotted,
			State: &BesottedState{
				a6Active:  c.info.Traces["103"],
				breakVuln: talent[c.info.TalentLevelIndex()],
				healAmt:   talent_heal[c.info.TalentLevelIndex()],
			},
		})
	}
}

func besottedBreakVuln(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_ELEMENT_DAMAGE {
		e.Hit.Defender.AddProperty(
			key.Reason(Besotted),
			prop.AllDamageTaken,
			mod.State().(BesottedState).breakVuln,
		)
	}
}

func besottedHeal(mod *modifier.Instance, e event.AttackEnd) {
	if !mod.Engine().IsCharacter(e.Attacker) {
		return
	}
	state := mod.State().(BesottedState)
	healtargets := []key.TargetID{e.Attacker}
	// Heal everyone if nectar blitz and gallagher a6 is activated, otherwise just heal attacker
	if e.Key == NectarBlitz && state.a6Active {
		healtargets = mod.Engine().Characters()
	}
	mod.Engine().Heal(info.Heal{
		Key:       key.Heal(Besotted),
		Source:    mod.Source(),
		Targets:   healtargets,
		HealValue: state.healAmt,
	})
}
