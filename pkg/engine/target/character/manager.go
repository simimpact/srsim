package character

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type register interface {
	AddTarget(key key.TargetID, base attribute.BaseStats) error
}

type Manager struct {
	engine    engine.Engine
	attr      register
	instances map[key.TargetID]CharInstance
	info      map[key.TargetID]info.Character
}

func New(engine engine.Engine, attr register) *Manager {
	return &Manager{
		engine:    engine,
		attr:      attr,
		instances: make(map[key.TargetID]CharInstance),
		info:      make(map[key.TargetID]info.Character),
	}
}

func (mgr *Manager) Get(id key.TargetID) (CharInstance, error) {
	if instance, ok := mgr.instances[id]; ok {
		return instance, nil
	}
	return nil, fmt.Errorf("target is not a character: %v", id)
}

func (mgr *Manager) Info(id key.TargetID) (info.Character, error) {
	if char, ok := mgr.info[id]; ok {
		return char, nil
	}
	return info.Character{}, fmt.Errorf("target is not a character: %v", id)
}
