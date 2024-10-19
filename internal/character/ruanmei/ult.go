package ruanmei

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult           = "ruanmei-ult"
	UltResPenAlly = "ruanmei-ult-res-pen-ally"
	UltBuffAlly   = "ruanmei-ult-buff-ally"
	UltDebuff     = "ruanmei-thanatopium-rebloom"
	UltDebuffCD   = "ruanmei-thanatopium-rebloom-cooldown"
)

func init() {
	// 2 mods summarized to 1 mod that represents both logic and visual/Buff status type
	modifier.Register(Ult, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		TickMoment: modifier.ModifierPhase1End,
		Listeners: modifier.Listeners{
			OnAdd:    addResPen,
			OnRemove: removeUltMods,
		},
	})
	// 2 mods summarized to 1 mod that represents both logic and visual/Buff status type
	modifier.Register(UltResPenAlly, modifier.Config{
		Stacking: modifier.Refresh,
	})
	// Purely for visual/Buff status type for other allies
	modifier.Register(UltBuffAlly, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(UltDebuff, modifier.Config{
		Listeners: modifier.Listeners{
			// OnAllowAction: removeReset, (unknown mechanic, will be ignored)
			OnHPChange: removeCDAndSelf,
			OnEndBreak: doUltImprintWithRemove,
			// OnDispel: removeCD, (missing listener)
			// OnBreakExtendAnim: doUltImprint, (missing listener)
		},
	})
	modifier.Register(UltDebuffCD, modifier.Config{
		Listeners: modifier.Listeners{
			OnLimboWaitHeal: func(mod *modifier.Instance) bool {
				mod.RemoveSelf()
				return false
			},
			OnEndBreak: func(mod *modifier.Instance) {
				mod.RemoveSelf()
			},
		},
	})
}

func (c *char) initUlt() {
	if c.info.Eidolon >= 1 {
		// Apply E1 to allies created while Ult is active
		c.engine.Events().CharactersAdded.Subscribe(func(event event.CharactersAdded) {
			for _, char := range event.Characters {
				trg := char.ID
				c.engine.AddModifier(trg, info.Modifier{
					Name:   E1,
					Source: c.id,
				})
			}
		})
	}
	// Apply TR Debuff with AttackEnd while RM has Ult
	c.engine.Events().AttackEnd.Subscribe(func(event event.AttackEnd) {
		if !c.engine.IsCharacter(event.Attacker) {
			return
		}
		if c.engine.HasModifier(c.id, Ult) {
			for _, trg := range event.Targets {
				c.engine.AddModifier(trg, info.Modifier{
					Name:   UltDebuff,
					Source: c.id,
				})
				c.engine.AddModifier(trg, info.Modifier{
					Name:   UltDebuffCD,
					Source: c.id,
				})
			}
		}
	})
	// Remove UltDebuff after it procs its damage
	c.engine.Events().TurnEnd.Subscribe(func(event event.TurnEnd) {
		for _, trg := range c.engine.Enemies() {
			if c.engine.HasModifierFromSource(trg, c.id, UltDebuff) {
				// Get UltDebuff's dynamic value
			}
		}
	})
	// Remove UltDebuff from all enemies if RM dies
	c.engine.Events().TargetDeath.Subscribe(func(event event.TargetDeath) {
		if event.Target == c.id {
			for _, trg := range c.engine.Enemies() {
				c.engine.RemoveModifierFromSource(trg, c.id, UltDebuff)
			}
		}
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	durationUlt := 2
	if c.info.Eidolon >= 6 {
		durationUlt += 1
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:     Ult,
		Source:   c.id,
		Duration: durationUlt,
		State:    ultResPen[c.info.UltLevelIndex()],
	})
	if c.info.Eidolon >= 1 {
		for _, trg := range c.engine.Characters() {
			c.engine.AddModifier(trg, info.Modifier{
				Name:   E1,
				Source: c.id,
			})
		}
	}
	for _, trg := range c.engine.Characters() {
		c.engine.AddModifier(trg, info.Modifier{
			Name:   UltBuffAlly,
			Source: c.id,
		})
	}
}

// Missing AllDamagePEN
func addResPen(mod *modifier.Instance) {
	for _, trg := range mod.Engine().Characters() {
		if trg == mod.Owner() {
			mod.AddProperty(prop.IcePEN, mod.State().(float64))
		} else {
			mod.Engine().AddModifier(trg, info.Modifier{
				Name:   UltResPenAlly,
				Source: mod.Owner(),
				Stats:  info.PropMap{prop.IcePEN: mod.State().(float64)},
			})
		}
	}
}

func removeUltMods(mod *modifier.Instance) {
	for _, trg := range mod.Engine().Characters() {
		mod.Engine().RemoveModifier(trg, UltResPenAlly)
		mod.Engine().RemoveModifier(trg, UltBuffAlly)
		mod.Engine().RemoveModifier(trg, E1)
	}
}

func removeCDAndSelf(mod *modifier.Instance, e event.HPChange) {
	if mod.Engine().HPRatio(mod.Owner()) == 0 {
		mod.Engine().RemoveModifier(mod.Owner(), UltDebuffCD)
		mod.RemoveSelf()
	}
}

func doUltImprintWithRemove(mod *modifier.Instance) {
	doUltImprint(mod)
	mod.Engine().RemoveModifier(mod.Owner(), UltDebuffCD)
	mod.RemoveSelf()
}

func doUltImprint(mod *modifier.Instance) {
	mod.Engine().InsertAbility(info.Insert{
		Key: UltDebuff,
		Execute: func() {

		},
		Source:   mod.Source(),
		Priority: info.CharInsertAttackSelf,
	})
}

// Need to do logic for "..._Count" dynamic value (declare inside UltDebuff)
