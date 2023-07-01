package modifier

//go:generate mockgen -destination=../../mock/mock_modifier.go -package=mock -mock_names Eval=MockModifier github.com/simimpact/srsim/pkg/engine/modifier Eval

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type activeModifiers []*ModifierInstance

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
		engine:  engine,
		targets: make(map[key.TargetID]activeModifiers, 10),
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

func (mgr *Manager) GetModifiers(target key.TargetID, name key.Modifier) []info.Modifier {
	out := make([]info.Modifier, 0, 5)
	for _, mod := range mgr.targets[target] {
		if mod.name == name {
			out = append(out, mod.ToModel())
		}
	}
	return out
}
