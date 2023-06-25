package combat

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) totalDamage(h *info.Hit, base float64, dmg float64) float64 {
	def := h.Defender.DEF()
	def_mult := 1 - (def / (def + 200 + 10*float64(h.Attacker.Level())))

	res := mgr.res(h)
	vul := mgr.vul(h)

	// ByPureDamage equation does not scale on DMG%, and break effect applies.
	breakDmg := 1.0
	if h.AttackType == model.AttackType_ELEMENT_DAMAGE {
		breakDmg += float64(h.Attacker.BreakEffect())
		dmg = 1.0
	}

	// toughness multiplier
	toughness_multiplier := 0.9
	if h.Defender.Stance() == 0 {
		toughness_multiplier = 1
	}

	// TODO: weaken

	total := base * dmg * def_mult * res * vul * breakDmg * toughness_multiplier
	return total
}

func (mgr *Manager) res(h *info.Hit) float64 {
	// We don't currently have normal dmg pen/res, dot pen/res, etc. If we do, we need to add it in here.
	res := float64(model.Property_ALL_DMG_RES)
	switch h.DamageType {
	case model.DamageType_PHYSICAL:
		res -= float64(h.Defender.GetProperty(prop.PhysicalDamageRES) - h.Attacker.GetProperty(prop.PhysicalPEN))
	case model.DamageType_FIRE:
		res -= float64(h.Defender.GetProperty(prop.FireDamageRES) - h.Attacker.GetProperty(prop.FirePEN))
	case model.DamageType_ICE:
		res -= float64(h.Defender.GetProperty(prop.IceDamageRES) - h.Attacker.GetProperty(prop.IcePEN))
	case model.DamageType_WIND:
		res -= float64(h.Defender.GetProperty(prop.WindDamageRES) - h.Attacker.GetProperty(prop.WindPEN))
	case model.DamageType_THUNDER:
		res -= float64(h.Defender.GetProperty(prop.ThunderDamageRES) - h.Attacker.GetProperty(prop.ThunderPEN))
	case model.DamageType_QUANTUM:
		res -= float64(h.Defender.GetProperty(prop.QuantumDamageRES) - h.Attacker.GetProperty(prop.QuantumPEN))
	case model.DamageType_IMAGINARY:
		res -= float64(h.Defender.GetProperty(prop.ImaginaryDamageRES) - h.Attacker.GetProperty(prop.ImaginaryPEN))
	}
	if res < -1 {
		res = -1
	} else if res > .9 {
		res = .9
	}
	return res
}

func (mgr *Manager) vul(h *info.Hit) float64 {
	vul := 1.0 + float64(h.Defender.GetProperty(prop.AllDamageTaken))
	switch h.DamageType {
	case model.DamageType_PHYSICAL:
		vul += float64(h.Defender.GetProperty(prop.PhysicalDamageTaken))
	case model.DamageType_FIRE:
		vul += float64(h.Defender.GetProperty(prop.FireDamageTaken))
	case model.DamageType_ICE:
		vul += float64(h.Defender.GetProperty(prop.IceDamageTaken))
	case model.DamageType_WIND:
		vul += float64(h.Defender.GetProperty(prop.WindDamageTaken))
	case model.DamageType_THUNDER:
		vul += float64(h.Defender.GetProperty(prop.ThunderDamageTaken))
	case model.DamageType_QUANTUM:
		vul += float64(h.Defender.GetProperty(prop.QuantumDamageTaken))
	case model.DamageType_IMAGINARY:
		vul += float64(h.Defender.GetProperty(prop.ImaginaryDamageTaken))
	}
	if vul > 1.35 {
		vul = 1.35
	}
	return vul
}
