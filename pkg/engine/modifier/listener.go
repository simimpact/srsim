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

	// ------------ combat events

	// Called when an attack starts and the attached target is the attacker.
	OnBeforeAttack func(mod *ModifierInstance, e event.AttackStartEvent)
	// Called when an attack starts and the attached target is one of the targets being attacked.
	OnBeforeBeingAttacked func(mod *ModifierInstance, e event.AttackStartEvent)
	// Called after an attack finishes (after all hits) and the attached target is the attacker
	OnAfterAttack func(mod *ModifierInstance, e event.AttackEndEvent)
	// Called after an attack finishes (after all hits) and the attached target was hit by the attack.
	OnAfterBeingAttacked func(mod *ModifierInstance, e event.AttackEndEvent)
	// Called before a hit occurs and the attached target is the attacker. Hit data is mutable
	// to allow modifiers to modify any stats prior to the damage calculation.
	OnBeforeHit func(mod *ModifierInstance, e event.BeforeHitEvent)
	// Called before a hit occurs and the attached target is the defender. Hit data is mutable
	// to allow modifiers to modify any stats prior to the damage calculation.
	OnBeforeBeingHit func(mod *ModifierInstance, e event.BeforeHitEvent)
	// Called after a hit occurs and the attached target is the attacker.
	OnAfterHit func(mod *ModifierInstance, e event.AfterHitEvent)
	// Called after a hit occurs and the attached target is the defender.
	OnAfterBeingHit func(mod *ModifierInstance, e event.AfterHitEvent)
	// Called before performing a heal and the attached target is the healer. Heal data is mutable.
	OnBeforeDealHeal func(mod *ModifierInstance, e *event.BeforeHealEvent)
	// Called before performing a heal and the attached target is the receiver. Heal data is mutable.
	OnBeforeBeingHeal func(mod *ModifierInstance, e *event.BeforeHealEvent)
	// Called after a heal is performed and the attached target is the healer.
	OnAfterDealHeal func(mod *ModifierInstance, e event.AfterHealEvent)
	// Called after a heal is performed and the attached target is the receiver
	OnAfterBeingHeal func(mod *ModifierInstance, e event.AfterHealEvent)
}

func (mgr *Manager) subscribe() {
	events := mgr.engine.Events()

	// attribute events
	events.HPChange.Subscribe(mgr.hpChange)
	events.LimboWaitHeal.Subscribe(mgr.limboWaitHeal, 100)
	events.TargetDeath.Subscribe(mgr.targetDeath)

	// combat events
	events.AttackStart.Subscribe(mgr.attackStart)
	events.AttackEnd.Subscribe(mgr.attackEnd)
	events.BeforeHit.Subscribe(mgr.beforeHit)
	events.AfterHit.Subscribe(mgr.afterHit)
	events.BeforeHeal.Subscribe(mgr.beforeHeal, 100)
	events.AfterHeal.Subscribe(mgr.afterHeal)
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

func (mgr *Manager) beforeHit(e event.BeforeHitEvent) {
	for _, mod := range mgr.targets[e.Attacker] {
		f := mod.listeners.OnBeforeHit
		if f != nil {
			f(mod, e)
		}
	}
	for _, mod := range mgr.targets[e.Defender] {
		f := mod.listeners.OnBeforeBeingHit
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) afterHit(e event.AfterHitEvent) {
	for _, mod := range mgr.targets[e.Attacker] {
		f := mod.listeners.OnAfterHit
		if f != nil {
			f(mod, e)
		}
	}
	for _, mod := range mgr.targets[e.Defender] {
		f := mod.listeners.OnAfterBeingHit
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) beforeHeal(e *event.BeforeHealEvent) {
	for _, mod := range mgr.targets[e.Healer.ID()] {
		f := mod.listeners.OnBeforeDealHeal
		if f != nil {
			f(mod, e)
		}
	}
	for _, mod := range mgr.targets[e.Target.ID()] {
		f := mod.listeners.OnBeforeBeingHeal
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) afterHeal(e event.AfterHealEvent) {
	for _, mod := range mgr.targets[e.Healer] {
		f := mod.listeners.OnAfterDealHeal
		if f != nil {
			f(mod, e)
		}
	}
	for _, mod := range mgr.targets[e.Target] {
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
