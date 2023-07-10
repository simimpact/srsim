package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1                    key.Modifier = "natasha-e1-autoheal"
	E1PercentThreshold    float64      = 0.30
	E1HealScale           float64      = 0.15
	E1HealFlat            float64      = 400
	E2                    key.Modifier = "natasha-e2"
	E2ThresholdPercentage float64      = 0.30
	E2PercentageHOT       float64      = 0.06
	E2FlatHOT             float64      = 160
	E4                    key.Modifier = "natasha-e4"
)

type E1State struct {
	PercentageThreshold float64
	HealScale           float64
	HealFlat            float64
	alreadyUsed         bool
}

func init() {

	//Self heal if hp lower than 30% after gettting hit : Eidolon 1
	modifier.Register(E1, modifier.Config{
		//Refactor into own method maybe?
		Listeners: modifier.Listeners{
			OnAfterBeingAttacked: func(mod *modifier.Instance, e event.AttackEnd) {
				selfHealer := mod.Owner()
				selfHealState := mod.State().(E1State)
				if !selfHealState.alreadyUsed {
					lowEnough := mod.OwnerStats().CurrentHPRatio() <= selfHealState.PercentageThreshold
					if lowEnough {
						mod.Engine().InsertAbility(info.Insert{
							Execute: func() {
								mod.Engine().Heal(info.Heal{
									Targets: []key.TargetID{selfHealer},
									Source:  selfHealer,
									BaseHeal: info.HealMap{
										model.HealFormula_BY_HEALER_MAX_HP: selfHealState.HealScale,
									},
									HealValue: selfHealState.HealFlat,
								})
							},
							Source: selfHealer,
							AbortFlags: []model.BehaviorFlag{
								model.BehaviorFlag_STAT_CTRL,
								model.BehaviorFlag_DISABLE_ACTION},
							Priority: info.CharHealSelf,
						})
						selfHealState.alreadyUsed = true
					}
				}
			},
		},
	})

	//Register E2
	modifier.Register(E2, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: func(mod *modifier.Instance) {
				mod.Engine().Heal(info.Heal{
					Targets: []key.TargetID{mod.Owner()},
					Source:  mod.Source(),
					BaseHeal: info.HealMap{
						model.HealFormula_BY_HEALER_MAX_HP: E2PercentageHOT,
					},
					HealValue:   E2FlatHOT,
					UseSnapshot: true,
				})
			},
		},
	})

	//Register E4
	modifier.Register(E4, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterBeingAttacked: func(mod *modifier.Instance, e event.AttackEnd) {
				mod.Engine().ModifyEnergy(mod.Owner(), 5)
			},
		},
	})
}

// Called when a character calls NewInstance
func (c *char) initEidolons() {
	//Self heal if hp lower than 30% after gettting hit
	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(c.id,
			info.Modifier{
				Name:   E1,
				Source: c.id,
				State: E1State{
					PercentageThreshold: E1PercentThreshold,
					HealScale:           E1HealScale,
					HealFlat:            E1HealFlat,
					alreadyUsed:         false,
				},
			})
	}

	//Extra 5 energy on being hit
	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E4,
			Source: c.id,
		})
	}
}

// Add a HOT if heal target is 30% hp or lower when healed
// Should only be called by Nat's ult
func (c *char) e2(targets []key.TargetID) {
	if c.info.Eidolon >= 2 {
		for _, trg := range targets {
			targetQualifies := c.engine.Stats(trg).CurrentHPRatio() <= E2ThresholdPercentage
			if targetQualifies {
				c.engine.AddModifier(trg, info.Modifier{
					Name:     E2,
					Source:   c.id,
					Duration: 1,
				})
			}
		}
	}
}

// Create an extra attack instance of type pursued, only called by nat basic attack
func (c *char) e6(target key.TargetID) {
	if c.info.Eidolon >= 6 {
		c.engine.Attack(info.Attack{
			Targets:    []key.TargetID{target},
			Source:     c.id,
			DamageType: model.DamageType_PHYSICAL,
			AttackType: model.AttackType_PURSUED,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_MAX_HP: 0.4,
			},
			StanceDamage: 0.0,
			EnergyGain:   0.0,
			HitRatio:     1.0,
		})
	}

}
