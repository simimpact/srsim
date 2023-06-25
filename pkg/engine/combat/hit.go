package combat

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
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

	// remainingDamage := mgr.shld.AbsorbDamage(hit.Defender.ID(), damage)

	// NOTE:
	// * BaseDamage multipliers, EnergyGain, and StanceDamage should be scaled by HitRatio
	// * dots & element damage do not crit (unknown if also ByPureDamage?)
	// * ByPureDamage = true means a "simplified" damage function (the break damage equation)

	mgr.event.HitEnd.Emit(event.HitEndEvent{
		Attacker:         hit.Attacker.ID(),
		Defender:         hit.Defender.ID(),
		AttackType:       hit.AttackType,
		DamageType:       hit.DamageType,
		HPDamage:         0, // TODO
		BaseDamage:       mgr.baseDamage(hit) * hit.HitRatio,
		BonusDamage:      0, // TODO
		TotalDamage:      0, // TODO
		ShieldDamage:     0, // TODO
		HPRatioRemaining: mgr.attr.HPRatio(hit.Defender.ID()),
		IsCrit:           false, // TODO
		UseSnapshot:      hit.UseSnapshot,
	})
}

// Base DMG = (MV + Extra MV) * Scaling Attribute + Extra DMG
// k = scaling attribute
// v = (MV + Extra MV)
// TODO: how to handle 'Extra DMG'?
func (mgr *Manager) baseDamage(h *info.Hit) float64 {
	var dmgMap info.DamageMap = h.BaseDamage
	var damage float64
	for k, v := range dmgMap {
		switch k {
		case model.DamageFormula_BY_ATK:
			damage = v * mgr.attr.Stats(h.Attacker.ID()).ATK()
		case model.DamageFormula_BY_DEF:
			damage = v * mgr.attr.Stats(h.Attacker.ID()).DEF()
		case model.DamageFormula_BY_MAX_HP:
			damage = v * mgr.attr.Stats(h.Attacker.ID()).MaxHP()
			// TODO: Figure out how to handle scale on break
			// case model.DamageFormula_BY_BREAK_DAMAGE:
		}
	}
	return damage

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
		BaseDamage:   baseDamage,
		EnergyGain:   atk.EnergyGain,
		StanceDamage: atk.StanceDamage,
		HitRatio:     ratio,
		AsPureDamage: atk.AsPureDamage,
		DamageValue:  atk.DamageValue,
		UseSnapshot:  atk.UseSnapshot,
	}
}
