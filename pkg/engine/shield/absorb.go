package shield

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) AbsorbDamage(target key.TargetID, damage float64) float64 {
	// Check if incoming damage = 0
	if damage != 0 {
		var removedShields []*Instance
		i := 0
		// If not 0, then run loop to apply damage to all shields
		for _, shield := range mgr.targets[target] {
			shield.hp -= damage
			if shield.hp < 0 {
				shield.hp = 0
			}
			// Generate list of shields to remove
			if shield.hp == 0 {
				removedShields = append(removedShields, shield)
			} else {
				mgr.targets[target][i] = shield
				i++
			}
		}
		mgr.targets[target] = mgr.targets[target][:i]

		// Getting outgoing damage
		maxShield := mgr.MaxShield(target)
		maxShieldHP := mgr.targets[target][maxShield]
		newMaxShieldHP := maxShieldHP.hp - damage
		damageOut := -newMaxShieldHP
		if damageOut < 0 {
			damageOut = 0
		}

		// Event emission to remove shields that have 0 hp
		for _, shield := range removedShields {
			mgr.event.ShieldRemoved.Emit(event.ShieldRemoved{
				ID:     shield.name,
				Target: target,
			})
		}

		// Event emission for shield hp change
		mgr.event.ShieldChange.Emit(event.ShieldChange{
			Target:    target,
			NewHP:     newMaxShieldHP,
			DamageIn:  damage,
			DamageOut: damageOut,
		})
		return damageOut
	}
	return damage
}
