package modifier

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type activeModifiers []*Instance

type Eval interface {
	EvalModifiers(target key.TargetID) info.ModifierState
}

type Manager struct {
	engine    engine.Engine
	targets   map[key.TargetID]activeModifiers
	turnCount int
}

func NewManager(engine engine.Engine) *Manager {
	mgr := &Manager{
		engine:    engine,
		targets:   make(map[key.TargetID]activeModifiers, 10),
		turnCount: 0,
	}
	mgr.subscribe()
	return mgr
}

func (mgr *Manager) HasFlag(target key.TargetID, flags ...model.BehaviorFlag) bool {
	flagSet := make(map[model.BehaviorFlag]struct{}, len(flags))
	for _, flag := range flags {
		flagSet[flag] = struct{}{}
	}

	for _, mod := range mgr.targets[target] {
		for _, flag := range mod.BehaviorFlags() {
			if _, ok := flagSet[flag]; ok {
				return true
			}
		}
	}
	return false
}

func (mgr *Manager) HasModifier(target key.TargetID, name key.Modifier) bool {
	for _, mod := range mgr.targets[target] {
		if mod.name == name {
			return true
		}
	}
	return false
}

func (mgr *Manager) HasModifierFromSource(target, source key.TargetID, name key.Modifier) bool {
	for _, mod := range mgr.targets[target] {
		if mod.name == name && mod.source == source {
			return true
		}
	}
	return false
}

func (mgr *Manager) GetModifiers(target key.TargetID, name key.Modifier) []info.Modifier {
	out := make([]info.Modifier, 0, 5)
	for _, mod := range mgr.targets[target] {
		if mod.name == name {
			out = append(out, mod.ToModel())
		}
	}
	return out
}

func (mgr *Manager) ModifierStackCount(target, source key.TargetID, modifier key.Modifier) float64 {
	count := 0.0
	for _, mod := range mgr.targets[target] {
		if mod.name == modifier && mod.source == source {
			count += mod.count
		}
	}
	return count
}

// makes a copy for safe iteration
func (mgr *Manager) itr(target key.TargetID) activeModifiers {
	out := make(activeModifiers, len(mgr.targets[target]))
	copy(out, mgr.targets[target])
	return out
}
