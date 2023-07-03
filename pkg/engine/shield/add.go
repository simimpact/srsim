package shield

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) AddShield(id key.Shield, shield info.Shield) {
	// TODO: LOGIC FOR ADDING A SHIELD
	// 1. Check if the target already has this shield, if so remove old (mgr.RemoveShield)
	// 2. Get the stats for the source and target
	// 3. Compute shield HP/create ShieldInstance given the add paramsm
	// 4. add shield to mgr.targets[shield.target]
	// 5. emit ShieldAdded event

	// emit to signify shield added
	mgr.event.ShieldAdded.Emit(event.ShieldAdded{
		ID:           id,
		Info:         shield,
		ShieldHealth: 0, // TODO: populate
	})
}
