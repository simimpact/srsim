package blade

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent       key.Attack   = "blade-talent"
	IsAttack     key.Modifier = "blade_is_Attack"
	GainedCharge key.Modifier = "blade_gained_charge"
)

func (c *char) Talent() {
	// Follow-up Attack
	c.engine.InsertAbility(info.Insert{
		Execute: func() {
			c.engine.Attack(info.Attack{
				Key:        Talent,
				Source:     c.id,
				Targets:    c.engine.Enemies(),
				DamageType: model.DamageType_WIND,
				AttackType: model.AttackType_INSERT,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK:    talentAtk[c.info.TalentLevelIndex()],
					model.DamageFormula_BY_MAX_HP: talentHP[c.info.TalentLevelIndex()],
				},
				StanceDamage: 30.0,
				EnergyGain:   10.0,
			})
		},
		Key:        key.Insert(Talent),
		Source:     c.id,
		Priority:   info.CharInsertAttackSelf,
		AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
	})

	// Heal
	c.engine.Heal(info.Heal{
		Key:      key.Heal(Talent),
		Targets:  []key.TargetID{c.id},
		Source:   c.id,
		BaseHeal: info.HealMap{model.HealFormula_BY_TARGET_MAX_HP: 0.25},
	})
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
