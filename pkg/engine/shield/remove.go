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

	delete(mgr.targets[target], id)

	mgr.event.ShieldRemoved.Emit(event.ShieldRemoved{
		ID:     id,
		Target: target,
	})
}
