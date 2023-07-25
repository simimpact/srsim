package modifier

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

// Decrease duration of modifiers and remove any modifiers that have expired (0 count or duration)
func (mgr *Manager) Tick(target key.TargetID, phase info.BattlePhase) {
	switch phase {
	case info.TurnStart:
		mgr.turnCount += 1
	case info.ModifierPhase1:
		for _, mod := range mgr.itr(target) {
			f := mod.listeners.OnPhase1
			if f != nil {
				f(mod)
			}
		}
		mgr.modifierPhaseEnd(target, ModifierPhase1End)
	case info.ActionEnd:
		// all modifiers added before action end should be flagged that they can tick on phase2
		for _, mod := range mgr.targets[target] {
			mod.canTickImmediatelyPhase2 = true
		}
	case info.ModifierPhase2:
		for _, mod := range mgr.itr(target) {
			f := mod.listeners.OnPhase2
			if f != nil {
				f(mod)
			}
		}
		mgr.modifierPhaseEnd(target, ModifierPhase2End)
	}
}

func (mgr *Manager) modifierPhaseEnd(target key.TargetID, time TickMoment) {
	i := 0
	var removedMods []*Instance
	for _, mod := range mgr.targets[target] {
		// only update modifier if its tick moment is for this given BattlePhase
		if modifierCatalog[mod.name].TickMoment != time {
			mgr.targets[target][i] = mod
			i++
			continue
		}

		// on phase 2, if tickImmediately is true canTickImmediately must also be true
		tickImmediately := mod.tickImmediately
		if time == ModifierPhase2End {
			tickImmediately = tickImmediately && mod.canTickImmediatelyPhase2
		}

		// if on application turn and TickImmediately is false, can skip this check
		if mgr.turnCount == mod.renewTurn && !tickImmediately {
			mgr.targets[target][i] = mod
			i++
			continue
		}

		remove := false

		// only remove mods based on count if their count is 0
		if mod.count == 0 {
			remove = true
		}

		// only decrease and remove duration of mods that have non-negative durations
		if mod.duration >= 0 {
			mod.duration -= 1
			if mod.duration <= 0 {
				mod.duration = 0
				remove = true
			}
		}

		if !remove {
			mgr.targets[target][i] = mod
			i++
			continue
		}

		removedMods = append(removedMods, mod)
	}
	mgr.targets[target] = mgr.targets[target][:i]
	mgr.emitRemove(target, removedMods)
}
