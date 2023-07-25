package character

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic"
)

type Manager struct {
	engine    engine.Engine
	attr      attribute.Manager
	eval      logic.Eval
	instances map[key.TargetID]info.CharInstance
	info      map[key.TargetID]info.Character
}

func New(engine engine.Engine, attr attribute.Manager, eval logic.Eval) *Manager {
	return &Manager{
		engine:    engine,
		attr:      attr,
		eval:      eval,
		instances: make(map[key.TargetID]info.CharInstance, 4),
		info:      make(map[key.TargetID]info.Character, 4),
	}
}

func (mgr *Manager) Get(id key.TargetID) (info.CharInstance, error) {
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

func (mgr *Manager) Characters() map[key.TargetID]info.Character {
	return mgr.info
}

func (mgr *Manager) SkillInfo(id key.TargetID) (SkillInfo, error) {
	if i, ok := mgr.info[id]; ok {
		return characterCatalog[i.Key].SkillInfo, nil
	}
	return SkillInfo{}, fmt.Errorf("target is not a character: %v", id)
}
