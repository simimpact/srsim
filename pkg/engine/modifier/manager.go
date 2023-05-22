package modifier

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type activeModifiers []*info.ModifierInstance

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

func (mgr *Manager) StartTurn() {
	mgr.turnCount += 1
}

// Decrease duration of modifiers and remove any modifiers that have expired (0 count or duration)
func (mgr *Manager) Tick(target key.TargetID, time BattlePhase) {
	i := 0
	for _, mod := range mgr.targets[target] {
		// only update modifier if its tick moment is for this given BattlePhase
		if modifierCatalog[mod.Name].TickMoment != time {
			mgr.targets[target][i] = mod
			i++
			continue
		}

		// if on application turn and TickImmediately is false, can skip this check
		if mgr.turnCount == mod.RenewTurn() && !mod.TickImmediately {
			mgr.targets[target][i] = mod
			i++
			continue
		}

		remove := false

		// only remove mods based on count if their count is 0
		if mod.Count == 0 {
			remove = true
		}

		// only decrease and remove duration of mods that have non-negative durations
		if mod.Duration >= 0 {
			mod.Duration -= 1
			if mod.Duration <= 0 {
				mod.Duration = 0
				remove = true
			}
		}

		if !remove {
			mgr.targets[target][i] = mod
			i++
			continue
		}

		mgr.remove(target, mod)
	}
	mgr.targets[target] = mgr.targets[target][:i]
}
