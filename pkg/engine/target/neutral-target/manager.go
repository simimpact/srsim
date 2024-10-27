package neutraltarget

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
	instances map[key.TargetID]info.NeutralTargetInstance
	info      map[key.TargetID]info.NeutralTarget
}

func New(engine engine.Engine, attr attribute.Manager, eval logic.Eval) *Manager {
	return &Manager{
		engine:    engine,
		attr:      attr,
		eval:      eval,
		instances: make(map[key.TargetID]info.NeutralTargetInstance, 4),
		info:      make(map[key.TargetID]info.NeutralTarget, 4),
	}
}

func (mgr *Manager) Get(id key.TargetID) (info.NeutralTargetInstance, error) {
	if instance, ok := mgr.instances[id]; ok {
		return instance, nil
	}

	return nil, fmt.Errorf("Target id was not a neutral target: %v", id)
}

func (mgr *Manager) Info(id key.TargetID) (info.NeutralTarget, error) {
	if info, ok := mgr.info[id]; ok {
		return info, nil
	}

	return info.NeutralTarget{}, fmt.Errorf("Target id was not a neutral target: %v", id)
}

func (mgr *Manager) Neutrals() map[key.TargetID]info.NeutralTarget {
	return mgr.info
}

func (mgr *Manager) AttackInfo(id key.TargetID) (Attack, error) {
	if target, ok := mgr.info[id]; ok {
		return neutralCatalog[target.Key].Attack, nil
	}

	return Attack{}, fmt.Errorf("Id was not a neutral %v", id)
}
