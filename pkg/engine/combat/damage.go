package combat

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

func (mgr *Manager) res(h *info.Hit) float64 {
	res := h.Defender.GetProperty(prop.AllDamageRES)
	res += h.Defender.GetProperty(prop.DamageRES(h.DamageType)) - h.Attacker.GetProperty(prop.DamagePEN(h.DamageType))

	if res < -1 {
		res = -1
	} else if res > .9 {
		res = .9
	}
	res = 1 - res

	return res
}

func (mgr *Manager) vul(h *info.Hit) float64 {
	vul := 1.0 + h.Defender.GetProperty(prop.AllDamageTaken)
	vul += h.Attacker.GetProperty(prop.DamageTaken(h.DamageType))
	if vul > 3.5 {
		vul = 3.5
	}
	return vul
}
