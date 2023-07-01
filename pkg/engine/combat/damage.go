package combat

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) totalDamage(h *info.Hit, base float64, dmg float64, crit bool) float64 {
	def := h.Defender.DEF()
	def_mult := 1 - (def / (def + 200 + 10*float64(h.Attacker.Level())))

	res := mgr.res(h)
	vul := mgr.vul(h)

	// ByPureDamage equation does not scale on DMG%, and break effect applies.
	breakDmg := 1.0
	if h.AttackType == model.AttackType_ELEMENT_DAMAGE {
		breakDmg += h.Attacker.BreakEffect()
		dmg = 1.0
	}

	// toughness multiplier
	toughness_multiplier := 0.9
	if h.Defender.Stance() == 0 {
		toughness_multiplier = 1
	}

	fatigue := 1 - h.Attacker.GetProperty(prop.Fatigue)
	AllDamageReduce := 1 - h.Defender.GetProperty(prop.AllDamageReduce)
	if AllDamageReduce < 0.01 {
		AllDamageReduce = 0.01
	}

	crit_dmg := 1.0
	if crit {
		crit_dmg += h.Attacker.CritDamage()
	}

	total := base * dmg * def_mult * res * vul * breakDmg * toughness_multiplier * fatigue * AllDamageReduce * crit_dmg
	return total
}

func (mgr *Manager) res(h *info.Hit) float64 {
	// We don't currently have basic/ult dmg pen/res, dot pen/res, etc. If we do, we need to add it in here.
	res := h.Defender.GetProperty(prop.AllDamageRES)
	res += h.Defender.GetProperty(prop.DamageRES(h.DamageType)) - h.Attacker.GetProperty(prop.DamagePEN(h.DamageType))
	res = 1 - res

	if res < -1 {
		res = -1
	} else if res > .9 {
		res = .9
	}
	return res
}

func (mgr *Manager) vul(h *info.Hit) float64 {
	vul := 1.0 + h.Defender.GetProperty(prop.AllDamageTaken)
	vul += h.Attacker.GetProperty(prop.DamageTaken(h.DamageType))
	if vul > 1.35 {
		vul = 1.35
	}
	return vul
}
