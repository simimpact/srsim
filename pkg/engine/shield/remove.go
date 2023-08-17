package shield

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) RemoveShield(id key.Shield, target key.TargetID) {
	// if target does not have the shield, then nothing to do
	if !mgr.HasShield(target, id) {
		return
	}

	i := 0
	for _, shield := range mgr.targets[target] {
		if shield.name != id {
			mgr.targets[target][i] = shield
			i++
		}
	}
	mgr.targets[target] = mgr.targets[target][:i]

	mgr.event.ShieldRemoved.Emit(event.ShieldRemoved{
		ID:     id,
		Target: target,
	})
}
