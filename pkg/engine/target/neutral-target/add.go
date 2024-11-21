package neutraltarget

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

// TODO: Add some way for neutrals to refer back to their owners/summoners to use their stats?
func (mgr *Manager) AddNeutral(id key.TargetID, key key.NeutralTarget) error {
	config, ok := neutralCatalog[key]
	if !ok {
		return fmt.Errorf("Invalid neutral target %v", key)
	}

	info := info.NeutralTarget{
		Key:     key,
		Element: config.Element,
	}

	mgr.info[id] = info
	mgr.instances[id] = config.Create(mgr.engine, id, info)
	return nil
}
