package combat

import (
	"math/rand"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

func baseDamage(h *info.Hit) float64 {
	dmgMap := h.BaseDamage
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
			damage += v * BreakBaseDamage[h.Attacker.Level()]
		}
	}
	return damage
}

func bonusDamage(h *info.Hit) float64 {
	dmg := 1.0
	// If hit doesn't use break damage equation, adds dmg%
	// Otherwise, adds break effect%
	if !h.AsPureDamage {
		dmg += h.Attacker.DamagePercent(h.DamageType)
		if h.AttackType == model.AttackType_DOT {
			dmg += h.Attacker.GetProperty(prop.DOTDamagePercent)
		}
	}

	if h.BaseDamage[model.DamageFormula_BY_BREAK_DAMAGE] != 0 {
		dmg += h.Attacker.BreakEffect()
	}

	return dmg
}

func defMult(h *info.Hit) float64 {
	def := h.Defender.DEF()
	mult := 1 - (def / (def + 200 + 10*float64(h.Attacker.Level())))
	return mult
}

func res(h *info.Hit) float64 {
	res := h.Defender.DamageRES(h.DamageType) - (h.Attacker.GetProperty(prop.DamagePEN(h.DamageType)) + h.Attacker.GetProperty(prop.AllDamagePEN))

	if res < -1 {
		res = -1
	} else if res > .9 {
		res = .9
	}
	res = 1 - res

	return res
}

func vul(h *info.Hit) float64 {
	vul := 1.0 + h.Defender.GetProperty(prop.AllDamageTaken)
	vul += h.Attacker.GetProperty(prop.DamageTaken(h.DamageType))
	if vul > 3.5 {
		vul = 3.5
	}
	return vul
}

func toughness(h *info.Hit) float64 {
	if h.Defender.Stance() == 0 {
		return 1.0
	}
	return 0.9
}

func damageReduce(h *info.Hit) float64 {
	reduce := 1 - h.Defender.GetProperty(prop.AllDamageReduce)
	if reduce < 0.01 {
		reduce = 0.01
	}
	return reduce
}

func crit(h *info.Hit, rdm *rand.Rand) bool {
	if h.AttackType == model.AttackType_DOT || h.AttackType == model.AttackType_ELEMENT_DAMAGE || h.AsPureDamage {
		return false
	}
	return rdm.Float64() < h.Attacker.CritChance()
}

func critDmg(h *info.Hit, crit bool) float64 {
	if crit {
		return 1.0 + h.Attacker.CritDamage()
	}
	return 1.0
}
