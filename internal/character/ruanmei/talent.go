package ruanmei

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent              = "ruanmei-talent"
	TalentSpdBuff       = "ruanmei-talent-spd-buff"
	TalentBreakListener = "ruanmei-talent-break-listener"
	UltResPen           = "ruanmei-ult-res-pen"
	A2                  = "ruanmei-a2"
	E2                  = "ruanmei-e2"
	E4                  = "ruanmei-e4"
	E4Listener          = "ruanmei-e4-listener"
)

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd:    addSpdAndA2,
			OnRemove: removeUltResPen,
		},
	})
	modifier.Register(TalentBreakListener, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeingBreak: doTalentBreakDamage,
		},
	})
	modifier.Register(UltResPen, modifier.Config{
		Stacking: modifier.Refresh,
	})
	modifier.Register(A2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(E2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: applyE2,
		},
	})
	// E4 is summarized to 1 buff mod
	modifier.Register(E4, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		// CanDispel: true,
	})
	// Causes known bug of the breaking hit not benefitting from E4 as this should apply with OnBeforeBeingBreak
	modifier.Register(E4Listener, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeingBreak: func(mod *modifier.Instance) {
				mod.Engine().AddModifier(mod.Source(), info.Modifier{
					Name:     E4,
					Source:   mod.Source(),
					Duration: 3,
				})
			},
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
	// - check if it is part of allied team,
	// - apply the mods to the target.
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

	// Add mods to upcoming enemies
	c.engine.Events().EnemiesAdded.Subscribe(
		func(event event.EnemiesAdded) {
			for _, trg := range c.engine.Enemies() {
				if c.info.Eidolon >= 4 {
					c.engine.AddModifier(trg, info.Modifier{
						Name:   E4,
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

	// Remove TalentBreakListener when Ruan Mei dies
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
		mod.Engine().RemoveModifier(trg, UltResPen)
	}
}

func doTalentBreakDamage(mod *modifier.Instance) {

}

func applyE2(mod *modifier.Instance, e event.HitStart) {

}
