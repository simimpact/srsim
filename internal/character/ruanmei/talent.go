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
	Talent              = "ruanmei-talent"
	TalentSpdBuff       = "ruanmei-talent-spd-buff"
	TalentBreakListener = "ruanmei-talent-break-listener"
	TalentBreak         = "ruanmei-talent-break-damage"
)

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd:         addSpdAndA2,
			OnRemove:      removeUltResPen,
			OnBeforeDying: removeTalentA2E2,
		},
	})
	modifier.Register(TalentSpdBuff, modifier.Config{
		Stacking:      modifier.Refresh,
		StatusType:    model.StatusType_STATUS_BUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
	})
	modifier.Register(TalentBreakListener, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeingBreak: doTalentBreakDamage,
		},
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
	})

	// Add mods to upcoming allies
	// Slightly inaccurate as this should:
	// - trigger after any entity gets created,
	// - check if the target is part of allied team,
	// - apply the relevant mods to the target.
	c.engine.Events().CharactersAdded.Subscribe(
		func(event event.CharactersAdded) {
			for _, trg := range c.engine.Characters() {
				c.engine.AddModifier(trg, info.Modifier{
					Name:   TalentSpdBuff,
					Source: c.id,
					Stats:  info.PropMap{prop.SPDPercent: 0.1},
				})
				if c.info.Traces["101"] {
					c.engine.AddModifier(trg, info.Modifier{
						Name:   A2,
						Source: c.id,
						Stats:  info.PropMap{prop.BreakEffect: 0.2},
					})
				}
				if c.info.Eidolon >= 2 {
					c.engine.AddModifier(trg, info.Modifier{
						Name:   E2,
						Source: c.id,
					})
				}
			}
		},
	)

	// Add mods to upcoming enemies (similar inaccuracy as above)
	c.engine.Events().EnemiesAdded.Subscribe(
		func(event event.EnemiesAdded) {
			for _, trg := range c.engine.Enemies() {
				if c.info.Eidolon >= 4 {
					c.engine.AddModifier(trg, info.Modifier{
						Name:   E4Listener,
						Source: c.id,
					})
				}
				c.engine.AddModifier(trg, info.Modifier{
					Name:   TalentBreakListener,
					Source: c.id,
				})
			}
		},
	)

	// Remove TalentBreakListener from all enemies when Ruan Mei dies
	c.engine.Events().TargetDeath.Subscribe(func(event event.TargetDeath) {
		if event.Target == c.id {
			for _, trg := range c.engine.Enemies() {
				c.engine.RemoveModifierFromSource(trg, c.id, TalentBreakListener)
			}
		}
	})

	// Add mods on BattleStart
	c.engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		if c.info.Eidolon >= 2 {
			for _, trg := range c.engine.Characters() {
				c.engine.AddModifier(trg, info.Modifier{
					Name:   E2,
					Source: c.id,
				})
			}
		}
		if c.info.Eidolon >= 4 {
			for _, trg := range c.engine.Enemies() {
				c.engine.AddModifier(trg, info.Modifier{
					Name:   E4Listener,
					Source: c.id,
				})
			}
		}
		if c.info.Traces["103"] {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   A6,
				Source: c.id,
			})
		}
	})
}

// Add mods to existing allies
func addSpdAndA2(mod *modifier.Instance) {
	for _, trg := range mod.Engine().Characters() {
		if trg != mod.Owner() {
			mod.Engine().AddModifier(trg, info.Modifier{
				Name:   TalentSpdBuff,
				Source: mod.Owner(),
				Stats:  info.PropMap{prop.SPDPercent: 0.1},
			})
		}
		mod.Engine().AddModifier(trg, info.Modifier{
			Name:   A2,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.BreakEffect: 0.2},
		})
	}
}

func removeUltResPen(mod *modifier.Instance) {
	for _, trg := range mod.Engine().Characters() {
		mod.Engine().RemoveModifier(trg, UltResPenAlly)
	}
}

func removeTalentA2E2(mod *modifier.Instance) {
	for _, trg := range mod.Engine().Characters() {
		mod.Engine().RemoveModifier(trg, TalentSpdBuff)
		mod.Engine().RemoveModifier(trg, A2)
		mod.Engine().RemoveModifier(trg, E2)
	}
}

// Applies Talent's Break Damage directly, without adding another mod
func doTalentBreakDamage(mod *modifier.Instance) {
	// Team member check skipped as it is deemed unnecessary

	rm, _ := mod.Engine().CharacterInfo(mod.Source())
	mult := talentBreakDamage[rm.TalentLevelIndex()]
	if rm.Eidolon >= 6 {
		mult += 2
	}
	maxStanceMult := ((mod.Engine().MaxStance(mod.Owner()) / 30) + 2) / 4
	mod.Engine().Attack(info.Attack{
		Key:          TalentBreak,
		Targets:      []key.TargetID{mod.Owner()},
		Source:       mod.Source(),
		DamageType:   model.DamageType_ICE,
		AttackType:   model.AttackType_ELEMENT_DAMAGE,
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_BREAK_DAMAGE: mult * maxStanceMult},
		AsPureDamage: true,
	})
}
