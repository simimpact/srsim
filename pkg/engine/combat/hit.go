package combat

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
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

	base := mgr.baseDamage(hit)*hit.HitRatio + hit.DamageValue
	bonus := mgr.bonusDamage(hit)
	total := mgr.totalDamage(hit, base, bonus)
	hpUpdate := [2]float64{mgr.shld.AbsorbDamage(hit.Defender.ID(), total), total}

	mgr.attr.ModifyHPByAmount(hit.Defender.ID(), hit.Attacker.ID(), total, true)
	mgr.attr.ModifyStance(hit.Defender.ID(), hit.Attacker.ID(), hit.StanceDamage)
	if mgr.target.IsCharacter(hit.Attacker.ID()) {
		mgr.attr.ModifyEnergy(hit.Attacker.ID(), hit.EnergyGain)
	} else {
		mgr.attr.ModifyEnergy(hit.Defender.ID(), hit.EnergyGain)
	}

	mgr.event.HitEnd.Emit(event.HitEndEvent{
		Attacker:         hit.Attacker.ID(),
		Defender:         hit.Defender.ID(),
		AttackType:       hit.AttackType,
		DamageType:       hit.DamageType,
		HPDamage:         hpUpdate[0],
		BaseDamage:       base,
		BonusDamage:      bonus,
		TotalDamage:      total,
		ShieldDamage:     hpUpdate[1], // TODO: Do we need this? AbsorbDamage already emits a shield modifying event
		HPRatioRemaining: mgr.attr.HPRatio(hit.Defender.ID()),
		IsCrit:           mgr.crit(hit),
		UseSnapshot:      hit.UseSnapshot,
	})
}

// BASE DAMAGE:
// Base DMG = (MV + Extra MV) * Scaling Attribute + Extra DMG
// k = scaling attribute
// v = (MV + Extra MV)
func (mgr *Manager) baseDamage(h *info.Hit) float64 {
	var dmgMap info.DamageMap = h.BaseDamage
	damage := 0.0
	for k, v := range dmgMap {
		switch k {
		case model.DamageFormula_BY_ATK:
			damage += v * h.Attacker.ATK()
		case model.DamageFormula_BY_DEF:
			damage += v * h.Attacker.DEF()
		case model.DamageFormula_BY_MAX_HP:
			damage += v * h.Attacker.MaxHP()
		case model.DamageFormula_BY_BREAK_DAMAGE:
			damage += v // TODO: Fact check this
		}
	}
	return damage
}

func (mgr *Manager) crit(h *info.Hit) bool {
	if h.AttackType == model.AttackType_DOT || h.AttackType == model.AttackType_ELEMENT_DAMAGE || h.AsPureDamage {
		return false
	}
	return mgr.rdm.Float64() < h.Attacker.CritChance()
}

func (mgr *Manager) bonusDamage(h *info.Hit) float64 {
	dmg := 1 + float64(h.Attacker.GetProperty(prop.AllDamagePercent))
	switch h.DamageType {
	case model.DamageType_PHYSICAL:
		dmg += float64(h.Attacker.GetProperty(prop.PhysicalDamagePercent))
	case model.DamageType_FIRE:
		dmg += float64(h.Attacker.GetProperty(prop.FireDamagePercent))
	case model.DamageType_ICE:
		dmg += float64(h.Attacker.GetProperty(prop.IceDamagePercent))
	case model.DamageType_WIND:
		dmg += float64(h.Attacker.GetProperty(prop.WindDamagePercent))
	case model.DamageType_THUNDER:
		dmg += float64(h.Attacker.GetProperty(prop.ThunderDamagePercent))
	case model.DamageType_QUANTUM:
		dmg += float64(h.Attacker.GetProperty(prop.QuantumDamagePercent))
	case model.DamageType_IMAGINARY:
		dmg += float64(h.Attacker.GetProperty(prop.ImaginaryDamagePercent))
	}

	// By my understanding, all other dmg% should be handled in AllDMGPercent
	if h.AttackType == model.AttackType_DOT {
		dmg += float64(h.Attacker.GetProperty(prop.DOTDamagePercent))
	}

	return dmg
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
