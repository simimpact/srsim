package modifier

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
)

type Listeners struct {
	OnAdd            func(engine engine.Engine, modifier *info.ModifierInstance)
	OnRemove         func(engine engine.Engine, modifier *info.ModifierInstance)
	OnExtendDuration func(engine engine.Engine, modifier *info.ModifierInstance)
	OnExtendCount    func(engine engine.Engine, modifier *info.ModifierInstance)
}

func (mgr *Manager) subscribe() {

}
