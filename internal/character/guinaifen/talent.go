package guinaifen

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent           = "guinaifen-talent"
	FirekissListener = "guinaifen-firekiss-listener"
	Firekiss         = "guinaifen-firekiss"
	E4Listener       = "guinaifen-e4-listener"
)

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAttack: A2,
		},
	})

	modifier.Register(FirekissListener, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeforeBeingHitAll: checkFirekiss,
		},
	})

	modifier.Register(Firekiss, modifier.Config{
		StatusType: model.StatusType_STATUS_DEBUFF,
		Stacking:   modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnAdd: FirekissOnStack,
		},
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
	})

	// apply Firekiss listener to all enemies at start
	c.engine.Events().EnemiesAdded.Subscribe(func(e event.EnemiesAdded) {
		for _, trg := range c.engine.Enemies() {
			c.engine.AddModifier(trg, info.Modifier{
				Name:   FirekissListener,
				Source: c.id,
			})
		}
	})

	// apply E4 listener to all enemies at start
	if c.info.Eidolon >= 4 {
		c.engine.Events().EnemiesAdded.Subscribe(func(e event.EnemiesAdded) {
			for _, trg := range c.engine.Enemies() {
				c.engine.AddModifier(trg, info.Modifier{
					Name:   E4Listener,
					Source: c.id,
				})
			}
		})
	}

	// remove listeners when source dies
	c.engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if e.Target == c.id {
			for _, trg := range c.engine.Enemies() {
				c.engine.RemoveModifierFromSource(trg, c.id, FirekissListener)
				if c.info.Eidolon >= 4 {
					c.engine.RemoveModifierFromSource(trg, c.id, E4Listener)
				}
			}
		}
	})

}

func checkFirekiss(mod *modifier.Instance, e event.HitStart) {
	gui, _ := mod.Engine().CharacterInfo(mod.Source())
	maxStacks := 3
	if gui.Eidolon >= 6 {
		maxStacks = 4
	}

	// check if damage was from Burn using workaround
	if e.Hit.AttackType == model.AttackType_DOT && e.Hit.DamageType == model.DamageType_FIRE {
		mod.Engine().AddModifier(e.Defender, info.Modifier{
			Name:              Firekiss,
			Source:            mod.Source(),
			Chance:            1,
			Duration:          3,
			MaxCount:          float64(maxStacks),
			CountAddWhenStack: 1,
		})
	}
}

// calculate the received damage increase
func FirekissOnStack(mod *modifier.Instance) {
	gui, _ := mod.Engine().CharacterInfo(mod.Source())
	mod.SetProperty(prop.AllDamageTaken, mod.Count()*talent[gui.TalentLevelIndex()])
}

// apply DoT on Normal
func A2(mod *modifier.Instance, e event.AttackStart) {
	gui, _ := mod.Engine().CharacterInfo(mod.Owner())
	if gui.Traces["101"] && e.AttackType == model.AttackType_NORMAL {
		target := e.Targets[0]
		multiplier := skillBurn[gui.SkillLevelIndex()]
		chance := 0.8
		if gui.Eidolon >= 2 && mod.Engine().HasBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_BURN) {
			multiplier += 0.4
		}
		applyBurn(mod.Engine(), mod.Owner(), target, multiplier, chance)
	}
}

// Talent's burn application function, with the Burn multiplier and its base chance as inputs
func applyBurn(engine engine.Engine, source, target key.TargetID, multiplier, chance float64) {
	engine.AddModifier(target, info.Modifier{
		Name:   common.Burn,
		Source: source,
		State: &common.BurnState{
			DamagePercentage:    multiplier,
			DamageValue:         0,
			DEFDamagePercentage: 0,
		},
		Chance:   chance,
		Duration: 2,
	})
}
