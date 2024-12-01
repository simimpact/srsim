package bronya

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1         key.Reason   = "bronya-e1"
	E1Cooldown key.Modifier = "bronya-e1-cooldown"
	E2Hover    key.Modifier = "bronya-e2-hover"
	E2Buff     key.Modifier = "bronya-e2-buff"
	E4Cooldown key.Modifier = "bronya-e4-buff"
	Insert                  = "bronya-follow-up"
)

func init() {
	modifier.Register(E2Hover, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAction: func(mod *modifier.Instance, e event.ActionEnd) {
				mod.Engine().AddModifier(mod.Owner(), info.Modifier{
					Name:   E2Buff,
					Source: mod.Source(),
				})
				mod.RemoveSelf()
			},
		},
	})

	modifier.Register(E2Buff, modifier.Config{
		StatusType:    model.StatusType_STATUS_BUFF,
		Stacking:      modifier.ReplaceBySource,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
		CanDispel:     true,
		Duration:      1,
	})

	modifier.Register(E4Cooldown, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
		Duration:   1,
	})
}

func (c *char) e1() {
	// If E1+
	if c.info.Eidolon >= 1 {
		// If not on CD
		if !c.engine.HasModifier(c.id, E1Cooldown) {
			// 50% Chance
			if c.engine.Rand().Float32() < 0.5 {
				// Add SP
				c.engine.ModifySP(info.ModifySP{
					Key:    E1,
					Source: c.id,
					Amount: 1,
				})
				// Set on CD
				c.engine.AddModifier(c.id, info.Modifier{
					Name:     E1Cooldown,
					Source:   c.id,
					Duration: 1,
				})
			}
		}
	}
}

func (c *char) e2(target key.TargetID) {
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(target, info.Modifier{
			Name:   E2Hover,
			Source: c.id,
			Stats:  info.PropMap{prop.SPDPercent: 0.3},
		})
	}
}

func (c *char) e4Listener(e event.AttackEnd) {
	// Assumed to be E4+ from subscription

	// has to be Ally Attacker
	if !c.engine.IsCharacter(e.Attacker) {
		return
	}

	// must not be Bronya
	if c.id == e.Attacker {
		return
	}

	// Off CD
	if c.engine.HasModifier(c.id, E4Cooldown) {
		return
	}

	toPickFrom := make([]key.TargetID, 0, len(e.Targets))

	for _, trg := range e.Targets {
		if c.engine.Stats(trg).IsWeakTo(model.DamageType_WIND) {
			toPickFrom = append(toPickFrom, trg)
		}
	}

	target := toPickFrom[c.engine.Rand().Intn(len(toPickFrom))]

	// Follow-up Attack
	c.engine.InsertAbility(info.Insert{
		Execute: func() {
			c.engine.Attack(info.Attack{
				Key:        Insert,
				Source:     c.id,
				Targets:    []key.TargetID{target},
				DamageType: model.DamageType_WIND,
				AttackType: model.AttackType_INSERT,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()] * 0.8,
				},
				StanceDamage: 30.0,
				EnergyGain:   5.0,
			})
		},
		Key:        Insert,
		Source:     c.id,
		Priority:   info.CharInsertAttackOthers,
		AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
	})

	// Set on CD
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   E4Cooldown,
		Source: c.id,
	})
}
