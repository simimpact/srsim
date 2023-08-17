package shield

import (
	"math"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) AbsorbDamage(target key.TargetID, damage float64) float64 {
	// Check if shielded or damage is non-positive
	if !mgr.IsShielded(target) || damage <= 0 {
		return damage
	}

	var removedShields []*Instance
	i := 0
	damageOut := damage
	newMaxShieldHP := 0.0
	oldMaxShieldHP := mgr.MaxShield(target)
	var maxShieldID key.Shield
	// Apply damage to all shields attached to entity
	for _, shield := range mgr.targets[target] {
		remaining := math.Dim(damage, shield.hp)
		shield.hp = math.Dim(shield.hp, damage)
		// Outgoing damage is always the lowest damage remaining after initial damage is subtracted from shields
		if remaining < damageOut {
			damageOut = remaining
		}
		// Max shield is always the highest shield hp remaining after damage is applied
		if shield.hp > newMaxShieldHP {
			newMaxShieldHP = shield.hp
			maxShieldID = shield.name
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
		ID:        maxShieldID,
		OldHP:     oldMaxShieldHP,
		NewHP:     newMaxShieldHP,
		DamageIn:  damage,
		DamageOut: damageOut,
	})
	return damageOut
}
