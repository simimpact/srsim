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

	mgr.event.HitEnd.Emit(event.HitEndEvent{
		Attacker:         hit.Attacker.ID(),
		Defender:         hit.Defender.ID(),
		AttackType:       hit.AttackType,
		DamageType:       hit.DamageType,
		HPDamage:         0, // TODO
		BaseDamage:       mgr.baseDamage(hit) * hit.HitRatio,
		BonusDamage:      mgr.bonusDamage(hit),
		TotalDamage:      mgr.totalDamage(hit, mgr.baseDamage(hit)*hit.HitRatio, mgr.bonusDamage(hit)),
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
	attacker := mgr.attr.Stats(h.Attacker.ID())
	dmg := 1 + float64(attacker.GetProperty(prop.AllDamagePercent))
	switch h.DamageType {
	case model.DamageType_PHYSICAL:
		dmg += float64(attacker.GetProperty(prop.PhysicalDamagePercent))
	case model.DamageType_FIRE:
		dmg += float64(attacker.GetProperty(prop.FireDamagePercent))
	case model.DamageType_ICE:
		dmg += float64(attacker.GetProperty(prop.IceDamagePercent))
	case model.DamageType_WIND:
		dmg += float64(attacker.GetProperty(prop.WindDamagePercent))
	case model.DamageType_THUNDER:
		dmg += float64(attacker.GetProperty(prop.ThunderDamagePercent))
	case model.DamageType_QUANTUM:
		dmg += float64(attacker.GetProperty(prop.QuantumDamagePercent))
	case model.DamageType_IMAGINARY:
		dmg += float64(attacker.GetProperty(prop.ImaginaryDamagePercent))
	}

	// By my understanding, all other dmg% should be handled in AllDMGPercent
	if h.AttackType == model.AttackType_DOT {
		dmg += float64(attacker.GetProperty(prop.DOTDamagePercent))
	}

	return dmg
}

// TODO: STANCE/TOUGHNESS

// TOTAL DAMAGE
// TODO: It appears that there is only one RES type for the entire sim. Change this when we get enemies.
func (mgr *Manager) totalDamage(h *info.Hit, base float64, dmg float64) float64 {
	// TODO: Check if DEF shred is applied already
	attacker := mgr.attr.Stats(h.Attacker.ID())
	defender := mgr.attr.Stats(h.Defender.ID())

	def := defender.DEF()
	def_mult := 1 - (def / (def + 200 + 10*float64(defender.Level())))

	// We don't currently have normal dmg pen/res, dot pen/res, etc. If we do, we need to add it in here.
	res := float64(model.Property_ALL_DMG_RES)
	switch h.DamageType {
	case model.DamageType_PHYSICAL:
		res -= float64(defender.GetProperty(prop.PhysicalDamageRES) - attacker.GetProperty(prop.PhysicalPEN))
	case model.DamageType_FIRE:
		res -= float64(defender.GetProperty(prop.FireDamageRES) - attacker.GetProperty(prop.FirePEN))
	case model.DamageType_ICE:
		res -= float64(defender.GetProperty(prop.IceDamageRES) - attacker.GetProperty(prop.IcePEN))
	case model.DamageType_WIND:
		res -= float64(defender.GetProperty(prop.WindDamageRES) - attacker.GetProperty(prop.WindPEN))
	case model.DamageType_THUNDER:
		res -= float64(defender.GetProperty(prop.ThunderDamageRES) - attacker.GetProperty(prop.ThunderPEN))
	case model.DamageType_QUANTUM:
		res -= float64(defender.GetProperty(prop.QuantumDamageRES) - attacker.GetProperty(prop.QuantumPEN))
	case model.DamageType_IMAGINARY:
		res -= float64(defender.GetProperty(prop.ImaginaryDamageRES) - attacker.GetProperty(prop.ImaginaryPEN))
	}
	if res < -1 {
		res = -1
	} else if res > .9 {
		res = .9
	}

	vul := 1.0 + float64(defender.GetProperty(prop.AllDamageTaken))
	switch h.DamageType {
	case model.DamageType_PHYSICAL:
		vul += float64(defender.GetProperty(prop.PhysicalDamageTaken))
	case model.DamageType_FIRE:
		vul += float64(defender.GetProperty(prop.FireDamageTaken))
	case model.DamageType_ICE:
		vul += float64(defender.GetProperty(prop.IceDamageTaken))
	case model.DamageType_WIND:
		vul += float64(defender.GetProperty(prop.WindDamageTaken))
	case model.DamageType_THUNDER:
		vul += float64(defender.GetProperty(prop.ThunderDamageTaken))
	case model.DamageType_QUANTUM:
		vul += float64(defender.GetProperty(prop.QuantumDamageTaken))
	case model.DamageType_IMAGINARY:
		vul += float64(defender.GetProperty(prop.ImaginaryDamageTaken))
	}
	if vul > 1.35 {
		vul = 1.35
	}

	return base * dmg * def_mult * res * vul
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
