package shield

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) AbsorbDamage(target key.TargetID, damage float64) float64 {
	// TODO: ACTUAL SHIELD LOGIC GOES HERE
	// 1. for every shield on this target, attempt to deal "damage" amount of damage to the shield
	// 2. for each shield, a ShieldChangeEvent should be emitted (how much damage was done to the shield)
	// 3. if shield HP <= damage, remove that shield (mgr.RemoveShield to generate event)
	// 4. if all shields HP <= damage, return will be damage - max(ShieldHP) [the damage not blocked by shield]
	// 5. if there are shields which HP > damage, return should be 0
	// 6. remaining state should be only shields that starting HP > damage

	// placeholder just for some basic event emission
	for _, shield := range mgr.targets[target] {
		prevHP := shield.HP
		shield.HP -= damage
		if shield.HP < 0 {
			shield.HP = 0
		}
		mgr.event.ShieldChange.Emit(event.ShieldChange{
			ID:     shield.name,
			Target: target,
			OldHP:  prevHP,
			NewHP:  shield.HP,
		})

		if shield.HP == 0 {
			mgr.RemoveShield(shield.name, target)
			damage -= prevHP
		}

		if shield.HP != 0 {
			damage = 0
		}
	}

	return damage
}
