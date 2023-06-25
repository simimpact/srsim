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
		BonusDamage:      mgr.bonusDamage(hit),
		TotalDamage:      0, // TODO
		ShieldDamage:     0, // TODO
		HPRatioRemaining: mgr.attr.HPRatio(hit.Defender.ID()),
		IsCrit:           mgr.crit(hit),
		UseSnapshot:      hit.UseSnapshot,
	})
}

// BASE DAMAGE:
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
		case model.DamageFormula_BY_BREAK_DAMAGE:
			damage = v // TODO: Fact check this
		}
	}
	return damage
}

func (mgr *Manager) crit(h *info.Hit) bool {
	if h.AttackType == model.AttackType_DOT || h.AttackType == model.AttackType_ELEMENT_DAMAGE {
		return false
	}
	return mgr.rdm.Float64() > mgr.attr.Stats(h.Attacker.ID()).CritChance()
}

func (mgr *Manager) bonusDamage(h *info.Hit) float64 {
	dmg := 1 + float64(model.Property_ALL_DMG_PERCENT)
	switch h.DamageType {
	case model.DamageType_PHYSICAL:
		dmg += float64(model.Property_PHYSICAL_DMG_PERCENT)
	case model.DamageType_FIRE:
		dmg += float64(model.Property_FIRE_DMG_PERCENT)
	case model.DamageType_ICE:
		dmg += float64(model.Property_ICE_DMG_PERCENT)
	case model.DamageType_WIND:
		dmg += float64(model.Property_WIND_DMG_PERCENT)
	case model.DamageType_THUNDER:
		dmg += float64(model.Property_THUNDER_DMG_PERCENT)
	case model.DamageType_QUANTUM:
		dmg += float64(model.Property_QUANTUM_DMG_PERCENT)
	case model.DamageType_IMAGINARY:
		dmg += float64(model.Property_IMAGINARY_DMG_PERCENT)
	}

	// By my understanding, all other dmg% should be handled in AllDMGPercent
	if h.AttackType == model.AttackType_DOT {
		dmg += float64(model.Property_DOT_DMG_PERCENT)
	}

	return dmg
}

// RES

// Vul

// STANCE/TOUGHNESS

// Total Damage

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
