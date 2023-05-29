package modifier

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
)

type activeModifiers []*ModifierInstance

type Manager struct {
	engine    engine.Engine
	targets   map[key.TargetID]activeModifiers
	turnCount int
}

func NewManager(engine engine.Engine) *Manager {
	mgr := &Manager{
		engine:  engine,
		targets: make(map[key.TargetID]activeModifiers),
	}
	mgr.subscribe()
	return mgr
}
