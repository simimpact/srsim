package blade

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent       = "blade-talent"
	IsAttack     = "blade-is-Attack"
	GainedCharge = "blade-gained-charge"
)

var talentHits = []float64{1.0 / 3, 1.0 / 3, 1.0 / 3}

func (c *char) Talent() {
	e6AddMod := 0.0

	// E6
	if c.info.Eidolon >= 6 {
		e6AddMod = 0.5
	}

	// Follow-up Attack
	for i, hitRatio := range talentHits {
		c.engine.InsertAbility(info.Insert{
			Execute: func() {
				c.engine.Attack(info.Attack{
					Key:        Talent,
					HitIndex:   i,
					Source:     c.id,
					Targets:    c.engine.Enemies(),
					DamageType: model.DamageType_WIND,
					AttackType: model.AttackType_INSERT,
					BaseDamage: info.DamageMap{
						model.DamageFormula_BY_ATK:    talentAtk[c.info.TalentLevelIndex()],
						model.DamageFormula_BY_MAX_HP: (talentHP[c.info.TalentLevelIndex()] + e6AddMod),
					},
					StanceDamage: 30.0,
					EnergyGain:   10.0,
					HitRatio:     hitRatio,
				})

				c.engine.EndAttack()

				// Heal
				c.engine.Heal(info.Heal{
					Key:      Talent,
					Targets:  []key.TargetID{c.id},
					Source:   c.id,
					BaseHeal: info.HealMap{model.HealFormula_BY_TARGET_MAX_HP: 0.25},
				})

				c.charge = 0
			},
			Key:        Talent,
			Source:     c.id,
			Priority:   info.CharInsertAttackSelf,
			AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
		})
	}
}

func (c *char) onBeforeBeingHitListener(e event.AttackStart) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   IsAttack,
		Source: c.id,
	})
}

func (c *char) onListenAfterAttackListener(e event.AttackEnd) {
	c.engine.RemoveModifier(c.id, IsAttack)
	c.engine.RemoveModifier(c.id, GainedCharge)
}
