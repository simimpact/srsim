package welt

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent key.Attack = "welt-talent"
	E2     key.Reason = "welt-e2"
)

// When hitting an enemy that is already Slowed,
// Welt deals Additional Imaginary DMG equal to 60% of his ATK to the enemy.

var triggerFlags = []model.BehaviorFlag{
	model.BehaviorFlag_STAT_SPEED_DOWN,
}

func (c *char) initTalent() {
	c.engine.Events().HitEnd.Subscribe(func(e event.HitEnd) {
		// negative check for early return
		if e.Attacker != c.id || !c.engine.HasBehaviorFlag(e.Defender, triggerFlags...) {
			return
		}

		// additional pursued dmg
		c.engine.Attack(info.Attack{
			Key:        Talent,
			Targets:    []key.TargetID{e.Defender},
			Source:     c.id,
			AttackType: model.AttackType_PURSUED,
			DamageType: model.DamageType_IMAGINARY,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: talentAtk[c.info.TalentLevelIndex()],
			},
		})

		// E2 : When his Talent is triggered, Welt regenerates 3 Energy.
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    E2,
			Target: c.id,
			Source: c.id,
			Amount: 3,
		})
	})
}
