package combat

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) performHit(hit *info.Hit) {
	mgr.event.HitStart.Emit(event.HitStartEvent{
		Attacker: hit.Attacker.ID(),
		Defender: hit.Defender.ID(),
		Hit:      hit,
	})

	// ACTUAL DAMAGE STUFF GOES HERE:
	// 1. Calculate Damage given the hit info (using stats contained within the hit)
	// 2. Given calc'd damage, call shield.AbsorbDamage(hit.defender, amt) float64
	// 3. AbsorbDamage returns the remaining damage
	// 4. ModifyHP of the remaining damage
	// 5. Emit DamageResultEvent

	// NOTE:
	// * BaseDamage multipliers, EnergyGain, and StanceDamage should be scaled by HitRatio
	// * dots & element damage do not crit (unknown if also ByPureDamage?)
	// * ByPureDamage = true means a "simplified" damage function (the break damage equation)

	mgr.event.HitEnd.Emit(event.HitEndEvent{
		Attacker:         hit.Attacker.ID(),
		Defender:         hit.Defender.ID(),
		AttackType:       hit.AttackType,
		DamageType:       hit.DamageType,
		AttackEffect:     hit.AttackEffect,
		BaseDamage:       0, // TODO
		BonusDamage:      0, // TODO
		TotalDamage:      0, // TODO
		ShieldDamage:     0, // TODO
		HPDamage:         0, // TODO
		HPRatioRemaining: mgr.attr.HPRatio(hit.Defender.ID()),
		IsCrit:           false, // TODO
	})
}

func (mgr *Manager) newHit(target key.TargetID, atk info.Attack) *info.Hit {
	// set HitRatio to 1 if unspecified
	ratio := atk.HitRatio
	if ratio <= 0 {
		ratio = 1
	}

	// make a copy of the base damage info
	baseDamage := make(info.DamageMap, len(atk.BaseDamage))
	for k, v := range atk.BaseDamage {
		baseDamage[k] = v
	}

	return &info.Hit{
		Attacker:     mgr.attr.Stats(atk.Source),
		Defender:     mgr.attr.Stats(target),
		AttackType:   atk.AttackType,
		DamageType:   atk.DamageType,
		AttackEffect: mgr.attackInfo.attackEffect,
		BaseDamage:   baseDamage,
		EnergyGain:   atk.EnergyGain,
		StanceDamage: atk.StanceDamage,
		HitRatio:     ratio,
		AsPureDamage: atk.AsPureDamage,
		DamageValue:  atk.DamageValue,
	}
}
