package danheng

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent   key.Modifier = "dan-heng-talent"
	TalentCD key.Modifier = "dan-heng-talent-cd"
)

type talentState struct {
	penAmt float64
	cd     int
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

// subscribe to all action ends to see if dan heng was ever the target of an ally action
func (c *char) talentActionEndListener(e event.ActionEvent) {
	// must be an ally (neutral targets dont count)
	if !c.engine.IsCharacter(e.Owner) {
		return
	}

	// cannot self trigger
	if c.id == e.Owner {
		return
	}

	// ignore if dan heng was not the target
	if !e.Targets[c.id] {
		return
	}

	// still on CD
	if c.engine.HasModifier(c.id, TalentCD) {
		return
	}

	// add talent modifier
	cd := 2
	if c.info.Eidolon >= 2 {
		cd -= 1
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: e.Owner,
		State: talentState{
			penAmt: talent[c.info.TalentLevelIndex()],
			cd:     cd,
		},
	})
}

func talentBeforeHitAll(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	state := mod.State().(talentState)

	/// only give pen to normal, skill, and ult hits. pursued will not be buffed
	if e.Hit.AttackType == model.AttackType_NORMAL ||
		e.Hit.AttackType == model.AttackType_SKILL ||
		e.Hit.AttackType == model.AttackType_ULT {
		e.Hit.Attacker.AddProperty(prop.WindPEN, state.penAmt)
	}
}

// after buffed action completes, add CD and remove talent buff
func talentAfterAction(mod *modifier.ModifierInstance, e event.ActionEvent) {
	state := mod.State().(talentState)

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:            TalentCD,
		Source:          mod.Owner(),
		Duration:        state.cd,
		TickImmediately: e.AttackType == model.AttackType_ULT,
	})

	mod.RemoveSelf()
}
