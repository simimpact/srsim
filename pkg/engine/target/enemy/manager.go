package enemy

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type Manager struct {
	engine engine.Engine
	attr   attribute.Manager
	info   map[key.TargetID]info.Enemy
}

func New(engine engine.Engine, attr attribute.Manager) *Manager {
	return &Manager{
		engine: engine,
		attr:   attr,
		info:   make(map[key.TargetID]info.Enemy, 5),
	}
}

func (mgr *Manager) Get(id key.TargetID) (Instance, error) {
	panic("not implemented")
}

func (mgr *Manager) Info(id key.TargetID) (info.Enemy, error) {
	if char, ok := mgr.info[id]; ok {
		return char, nil
	}
	return info.Enemy{}, fmt.Errorf("target is not an enemy: %v", id)
}

func (mgr *Manager) Enemies() map[key.TargetID]info.Enemy {
	return mgr.info
}
