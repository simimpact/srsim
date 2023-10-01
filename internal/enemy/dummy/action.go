package dummy

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Single = "dummy-single"
	Blast  = "dummy-blast"
	Bounce = "dummy-bounce"
	AOE    = "dummy-aoe"
)

// Action implements info.EnemyInstance.
func (e *impl) Action(target key.TargetID, state info.ActionState) {
	switch e.attack {
	case None:
		// no attack, enemy does nothing
	case SingleAttack:
		e.Single(target)
	case BlastAttack:
		e.Blast(target)
	case BounceAttack:
		e.Bounce(target)
	case AoeAttack:
		e.Aoe()
	}
}

func (e *impl) Single(target key.TargetID) {
	ratio := 1 / float64(e.hitCount)
	for i := 0; i < e.hitCount; i++ {
		e.engine.Attack(info.Attack{
			Key:        Single,
			HitIndex:   i,
			Source:     e.id,
			Targets:    []key.TargetID{target},
			DamageType: e.damageType,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: e.damagePercent,
			},
			HitRatio:   ratio,
			EnergyGain: e.energyGain,
		})
	}
}

func (e *impl) Blast(target key.TargetID) {
	ratio := 1 / float64(e.hitCount)
	for i := 0; i < e.hitCount; i++ {
		atk := info.Attack{
			Key:        Blast,
			HitIndex:   i,
			Source:     e.id,
			Targets:    []key.TargetID{target},
			DamageType: e.damageType,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: e.damagePercent,
			},
			HitRatio:   ratio,
			EnergyGain: e.energyGain,
		}

		// primary target hit
		e.engine.Attack(atk)

		// hit adjacent targets to primary
		atk.Targets = e.engine.AdjacentTo(target)
		e.engine.Attack(atk)
	}
}

func (e *impl) Aoe() {
	ratio := 1 / float64(e.hitCount)
	for i := 0; i < e.hitCount; i++ {
		e.engine.Attack(info.Attack{
			Key:        AOE,
			HitIndex:   i,
			Source:     e.id,
			Targets:    e.engine.Characters(),
			DamageType: e.damageType,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: e.damagePercent,
			},
			HitRatio:   ratio,
			EnergyGain: e.energyGain,
		})
	}
}

func (e *impl) Bounce(target key.TargetID) {
	for i := 0; i < e.hitCount; i++ {
		target := []key.TargetID{target}

		if i > 0 {
			targets := e.engine.Characters()
			allDead := allTargetsDead(e.engine, targets)

			target = e.engine.Retarget(info.Retarget{
				Targets:      targets,
				Max:          1,
				IncludeLimbo: true,
				Filter: func(t key.TargetID) bool {
					return allDead || e.engine.HPRatio(t) > 0
				},
			})
		}

		e.engine.Attack(info.Attack{
			Key:        Bounce,
			HitIndex:   i,
			Source:     e.id,
			Targets:    target,
			DamageType: e.damageType,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: e.damagePercent,
			},
			EnergyGain: e.energyGain,
		})
	}
}

func allTargetsDead(engine engine.Engine, targets []key.TargetID) bool {
	for _, t := range targets {
		if engine.HPRatio(t) > 0 {
			return false
		}
	}
	return true
}
