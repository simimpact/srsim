package modifier

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

type Listeners struct {
	// ------------ listeners for modifier internal processes

	// Called when a new modifier instance is added. Note: if using Replace or ReplaceBySource,
	// this will always be a fresh instance when stacking. If using Merge, OnAdd will be called
	// on the old instance.
	OnAdd func(mod *ModifierInstance)

	// Called when a modifier instance is removed, either forceably or due to the instance expiring.
	OnRemove func(mod *ModifierInstance)

	// Called when the duration for all modifiers instances of this shape are extended.
	OnExtendDuration func(mod *ModifierInstance)

	// Called when the count/stacks for all modifier instances of this shape are extended.
	// Will not be called if OnAdd is called (doesnt call on standard stacking behavior)
	OnExtendCount func(mod *ModifierInstance)

	// Called when any stat changes on the target this modifier is attached to. Will be called if
	// you modify properties within this call, so take care not to create a recursive loop.
	OnPropertyChange func(mod *ModifierInstance)

	// Called at the start of the turn, before the action takes place (used by DoTs).
	OnPhase1 func(mod *ModifierInstance)

	// Called at the end of the turn
	OnPhase2 func(mod *ModifierInstance)

	// ------------ attribute events

	// Called when the current HP of the attached target changes
	OnHPChange func(mod *ModifierInstance, e event.HPChangeEvent)

	// Called when attached target's current HP = 0. If returns true, will cancel the event and
	// prevent the TargetDeathEvent from occuring. Used by revives.
	OnLimboWaitHeal func(mod *ModifierInstance) bool

	// Called when the attached target kills another target. The given target ID is the target that
	// has been killed.
	OnTriggerDeath func(mod *ModifierInstance, target key.TargetID)

	// Called whe nthe attached start
	OnEnergyChange func(mod *ModifierInstance, e event.EnergyChangeEvent)

	// Called when the attached target stance changes
	OnStanceChange func(mod *ModifierInstance, e event.StanceChangeEvent)

	// Called when the attached target causes another target to go into a break state (0 stance).
	OnTriggerBreak func(mod *ModifierInstance, target key.TargetID)

	// Called when the attached target goes into a break state (stance reached 0).
	OnBeingBreak func(mod *ModifierInstance)

	// Called when the attached target break status ends (stance resets to max).
	OnEndBreak func(mod *ModifierInstance)

	// ------------ combat events

	// Called when an attack starts and the attached target is the attacker.
	OnBeforeAttack func(mod *ModifierInstance, e event.AttackStartEvent)

	// Called when an attack starts and the attached target is one of the targets being attacked.
	OnBeforeBeingAttacked func(mod *ModifierInstance, e event.AttackStartEvent)

	// Called after an attack finishes (after all hits) and the attached target is the attacker
	OnAfterAttack func(mod *ModifierInstance, e event.AttackEndEvent)

	// Called after an attack finishes (after all hits) and the attached target was hit by the attack.
	OnAfterBeingAttacked func(mod *ModifierInstance, e event.AttackEndEvent)

	// Called before any hit occurs and the attached target is the attacker. Hit data is mutable.
	OnBeforeHitAll func(mod *ModifierInstance, e event.HitStartEvent)

	// Called before a qualified hit occurs and the attached target is the attacker. "Qualified" hit
	// means it is not of AttackType DOT, PURSUED, or ELEMENT_DAMAGE. Hit data is mutable.
	OnBeforeHit func(mod *ModifierInstance, e event.HitStartEvent)

	// called before any hit occurs and the attached target is the defender. Hit data is mutable.
	OnBeforeBeingHitAll func(mod *ModifierInstance, e event.HitStartEvent)

	// Called before a qualified hit occurs and the attached target is the defender. "Qualified" hit
	// means it is not of AttackType DOT, PURSUED, or ELEMENT_DAMAGE. Hit data is mutable.
	OnBeforeBeingHit func(mod *ModifierInstance, e event.HitStartEvent)

	// Called after any hit occurs and the attached target is the attacker.
	OnAfterHitAll func(mod *ModifierInstance, e event.HitEndEvent)

	// Called after a qualified hit occurs and the attached target is the attacker. "Qualified" hit
	// means it is not of AttackType DOT, PURSUED, or ELEMENT_DAMAGE.
	OnAfterHit func(mod *ModifierInstance, e event.HitEndEvent)

	// Called after any hit occurs and the attached target is the defender.
	OnAfterBeingHitAll func(mod *ModifierInstance, e event.HitEndEvent)

	// Called after a qualified hit occurs and the attached target is the defender. "Qualified" hit
	// means it is not of AttackType DOT, PURSUED, or ELEMENT_DAMAGE.
	OnAfterBeingHit func(mod *ModifierInstance, e event.HitEndEvent)

	// Called before performing a heal and the attached target is the healer. Heal data is mutable.
	OnBeforeDealHeal func(mod *ModifierInstance, e *event.HealStartEvent)

	// Called before performing a heal and the attached target is the receiver. Heal data is mutable.
	OnBeforeBeingHeal func(mod *ModifierInstance, e *event.HealStartEvent)

	// Called after a heal is performed and the attached target is the healer.
	OnAfterDealHeal func(mod *ModifierInstance, e event.HealEndEvent)

	// Called after a heal is performed and the attached target is the receiver
	OnAfterBeingHeal func(mod *ModifierInstance, e event.HealEndEvent)
}

func (mgr *Manager) subscribe() {
	events := mgr.engine.Events()

	// attribute events
	events.HPChange.Subscribe(mgr.hpChange)
	events.LimboWaitHeal.Subscribe(mgr.limboWaitHeal, 100)
	events.TargetDeath.Subscribe(mgr.targetDeath)
	events.EnergyChange.Subscribe(mgr.energyChange)
	events.StanceChange.Subscribe(mgr.stanceChange)
	events.StanceBreak.Subscribe(mgr.stanceBreak)
	events.StanceBreakEnd.Subscribe(mgr.stanceBreakEnd)

	// combat events
	events.AttackStart.Subscribe(mgr.attackStart)
	events.AttackEnd.Subscribe(mgr.attackEnd)
	events.HitStart.Subscribe(mgr.hitStart)
	events.HitEnd.Subscribe(mgr.hitEnd)
	events.HealStart.Subscribe(mgr.healStart, 100)
	events.HealEnd.Subscribe(mgr.healEnd)
}

func (mgr *Manager) emitPropertyChange(target key.TargetID) {
	for _, mod := range mgr.targets[target] {
		f := mod.listeners.OnPropertyChange
		if f != nil {
			f(mod)
		}
	}
}

func (mgr *Manager) emitAdd(target key.TargetID, mod *ModifierInstance, chance float64) {
	f := mod.listeners.OnAdd
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierAdded.Emit(event.ModifierAddedEvent{
		Target:   target,
		Modifier: mod.ToModel(),
		Chance:   chance,
	})
}

func (mgr *Manager) emitRemove(target key.TargetID, mods []*ModifierInstance) {
	for _, mod := range mods {
		if len(mod.stats) > 0 {
			mgr.emitPropertyChange(target)
		}

		f := mod.listeners.OnRemove
		if f != nil {
			f(mod)
		}
		mgr.engine.Events().ModifierRemoved.Emit(event.ModifierRemovedEvent{
			Target:   target,
			Modifier: mod.ToModel(),
		})
	}
}

func (mgr *Manager) emitExtendDuration(target key.TargetID, mod *ModifierInstance, old int) {
	f := mod.listeners.OnExtendDuration
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierExtendedDuration.Emit(event.ModifierExtendedDurationEvent{
		Target:   target,
		Modifier: mod.ToModel(),
		OldValue: old,
		NewValue: mod.duration,
	})
}

func (mgr *Manager) emitExtendCount(target key.TargetID, mod *ModifierInstance, old float64) {
	f := mod.listeners.OnExtendCount
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierExtendedCount.Emit(event.ModifierExtendedCountEvent{
		Target:   target,
		Modifier: mod.ToModel(),
		OldValue: old,
		NewValue: mod.count,
	})
}

func (mgr *Manager) attackStart(e event.AttackStartEvent) {
	for _, mod := range mgr.targets[e.Attacker] {
		f := mod.listeners.OnBeforeAttack
		if f != nil {
			f(mod, e)
		}
	}
	for _, target := range e.Targets {
		for _, mod := range mgr.targets[target] {
			f := mod.listeners.OnBeforeBeingAttacked
			if f != nil {
				f(mod, e)
			}
		}
	}
}

func (mgr *Manager) attackEnd(e event.AttackEndEvent) {
	for _, mod := range mgr.targets[e.Attacker] {
		f := mod.listeners.OnAfterAttack
		if f != nil {
			f(mod, e)
		}
	}
	for _, target := range e.Targets {
		for _, mod := range mgr.targets[target] {
			f := mod.listeners.OnAfterBeingAttacked
			if f != nil {
				f(mod, e)
			}
		}
	}
}

func (mgr *Manager) hitStart(e event.HitStartEvent) {
	qualified := e.Hit.AttackType.IsQualified()
	snapshot := e.Hit.UseSnapshot

	for _, mod := range mgr.targets[e.Attacker] {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnBeforeHitAll
		if f != nil {
			f(mod, e)
		}

		f = mod.listeners.OnBeforeHit
		if f != nil && qualified {
			f(mod, e)
		}
	}

	for _, mod := range mgr.targets[e.Defender] {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnBeforeBeingHitAll
		if f != nil {
			f(mod, e)
		}

		f = mod.listeners.OnBeforeBeingHit
		if f != nil && qualified {
			f(mod, e)
		}
	}
}

func (mgr *Manager) hitEnd(e event.HitEndEvent) {
	qualified := e.AttackType.IsQualified()
	snapshot := e.UseSnapshot

	for _, mod := range mgr.targets[e.Attacker] {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnAfterHitAll
		if f != nil {
			f(mod, e)
		}

		f = mod.listeners.OnAfterHit
		if f != nil && qualified {
			f(mod, e)
		}
	}

	for _, mod := range mgr.targets[e.Defender] {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnAfterBeingHitAll
		if f != nil {
			f(mod, e)
		}

		f = mod.listeners.OnAfterBeingHit
		if f != nil && qualified {
			f(mod, e)
		}
	}
}

func (mgr *Manager) healStart(e *event.HealStartEvent) {
	snapshot := e.UseSnapshot

	for _, mod := range mgr.targets[e.Healer.ID()] {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnBeforeDealHeal
		if f != nil {
			f(mod, e)
		}
	}

	for _, mod := range mgr.targets[e.Target.ID()] {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnBeforeBeingHeal
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) healEnd(e event.HealEndEvent) {
	snapshot := e.UseSnapshot

	for _, mod := range mgr.targets[e.Healer] {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnAfterDealHeal
		if f != nil {
			f(mod, e)
		}
	}

	for _, mod := range mgr.targets[e.Target] {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnAfterBeingHeal
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) hpChange(e event.HPChangeEvent) {
	for _, mod := range mgr.targets[e.Target] {
		f := mod.listeners.OnHPChange
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) limboWaitHeal(e event.LimboWaitHealEvent) bool {
	for _, mod := range mgr.targets[e.Target] {
		f := mod.listeners.OnLimboWaitHeal
		if f != nil {
			result := f(mod)
			if result {
				return true
			}
		}
	}
	return false
}

func (mgr *Manager) targetDeath(e event.TargetDeathEvent) {
	for _, mod := range mgr.targets[e.Killer] {
		f := mod.listeners.OnTriggerDeath
		if f != nil {
			f(mod, e.Target)
		}
	}
}

func (mgr *Manager) energyChange(e event.EnergyChangeEvent) {
	for _, mod := range mgr.targets[e.Target] {
		f := mod.listeners.OnEnergyChange
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) stanceChange(e event.StanceChangeEvent) {
	for _, mod := range mgr.targets[e.Target] {
		f := mod.listeners.OnStanceChange
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) stanceBreak(e event.StanceBreakEvent) {
	for _, mod := range mgr.targets[e.Source] {
		f := mod.listeners.OnTriggerBreak
		if f != nil {
			f(mod, e.Target)
		}
	}
	for _, mod := range mgr.targets[e.Target] {
		f := mod.listeners.OnBeingBreak
		if f != nil {
			f(mod)
		}
	}
}

func (mgr *Manager) stanceBreakEnd(e event.StanceBreakEndEvent) {
	for _, mod := range mgr.targets[e.Target] {
		f := mod.listeners.OnEndBreak
		if f != nil {
			f(mod)
		}
	}
}
