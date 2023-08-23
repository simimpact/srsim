package seele

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1Check  = "seele-e1-check"
	E4       = "seele-e4"
	E6       = "seele-e6"
	E6Debuff = "seele-e6-debuff"
)

// E1 : When dealing DMG to an enemy whose HP percentage is 80% or lower,
// CRIT Rate increases by 15%.
// E4 : Seele regenerates 15 Energy when she defeats an enemy.
// E6 : After Seele uses her Ultimate, inflict the target enemy with Butterfly Flurry
// for 1 turn(s). Enemies suffering from Butterfly Flurry will take Additional
// Quantum DMG equal to 15% of Seele's Ultimate DMG every time they are attacked.
// If the target enemy is defeated by the Butterfly Flurry DMG triggered by
// other allies' attacks, Seele's Talent will not be triggered.
// When Seele is knocked down, the Butterfly Flurry inflicted on the enemies will be removed.

func init() {
	modifier.Register(E1Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: boostCR,
		},
		CanModifySnapshot: true,
	})
	modifier.Register(E4, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: addFlatEnergy,
		},
	})
	modifier.Register(E6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeDying: removeDebuffOnEnemies,
		},
	})

	modifier.Register(E6Debuff, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterBeingAttacked: addPursuedDmg,
		},
		StatusType: model.StatusType_STATUS_DEBUFF,
		Stacking:   modifier.ReplaceBySource,
	})
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E1Check,
			Source: c.id,
		})
	}
	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E4,
			Source: c.id,
		})
	}
	if c.info.Eidolon >= 6 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6,
			Source: c.id,
		})
	}
}

func boostCR(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().HPRatio(e.Defender) <= 0.8 {
		e.Hit.Attacker.AddProperty(E6, prop.CritChance, 0.15)
	}
}

func addFlatEnergy(mod *modifier.Instance, target key.TargetID) {
	mod.Engine().ModifyEnergy(info.ModifyAttribute{
		Key:    E4,
		Target: mod.Owner(),
		Source: mod.Owner(),
		Amount: 15.0,
	})
}

// when seele is dying, remove all instances of e6 debuffs on all enemies.
func removeDebuffOnEnemies(mod *modifier.Instance) {
	for _, enemy := range mod.Engine().Enemies() {
		mod.Engine().RemoveModifier(enemy, E6Debuff)
	}
}

func addPursuedDmg(mod *modifier.Instance, e event.AttackEnd) {
	// fetch seele's atk value. confirm if this is snapshot or dynamic
	atkAmt := mod.Engine().Stats(mod.Source()).ATK()

	// apply dmg to mod owner from mod applier(source)
	mod.Engine().Attack(info.Attack{
		Key:        E6Debuff,
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		DamageType: model.DamageType_QUANTUM,
		AttackType: model.AttackType_PURSUED,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: atkAmt * 0.15,
		},
	})
}
