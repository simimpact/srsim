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
	OnAdd func(mod *Instance)

	// Called when a modifier instance is removed, either forceably or due to the instance expiring.
	OnRemove func(mod *Instance)

	// Called when the duration for all modifiers instances of this shape are extended.
	OnExtendDuration func(mod *Instance)

	// Called when the count/stacks for all modifier instances of this shape are extended.
	// Will not be called if OnAdd is called (doesnt call on standard stacking behavior)
	OnExtendCount func(mod *Instance)

	// Called when any stat changes on the target this modifier is attached to. Will be called if
	// you modify properties within this call, so take care not to create a recursive loop.
	OnPropertyChange func(mod *Instance)

	// Called at the start of the turn, before the action takes place (used by DoTs).
	OnPhase1 func(mod *Instance)

	// Called at the end of the turn
	OnPhase2 func(mod *Instance)

	// ------------ attribute events

	// Called when the current HP of the attached target changes
	OnHPChange func(mod *Instance, e event.HPChange)

	// Called when attached target's current HP = 0. If returns true, will cancel the event and
	// prevent the TargetDeathEvent from occuring. Used by revives.
	OnLimboWaitHeal func(mod *Instance) bool

	// Called when the attached target has taken fatal damage and no revive will occur.
	OnBeforeDying func(mod *Instance)

	// Called when the attached target kills another target. The given target ID is the target that
	// has been killed.
	OnTriggerDeath func(mod *Instance, target key.TargetID)

	// Called whe nthe attached start
	OnEnergyChange func(mod *Instance, e event.EnergyChange)

	// Called when the attached target stance changes
	OnStanceChange func(mod *Instance, e event.StanceChange)

	// Called when the attached target causes another target to go into a break state (0 stance).
	OnTriggerBreak func(mod *Instance, target key.TargetID)

	// Called when the attached target goes into a break state (stance reached 0).
	OnBeingBreak func(mod *Instance)

	// Called when the attached target break status ends (stance resets to max).
	OnEndBreak func(mod *Instance)

	// ------------ shield events

	// Called when a shield has been added to the attached target.
	OnShieldAdded func(mod *Instance, e event.ShieldAdded)

	// Called when a shield has been removed from the attached target.
	OnShieldRemoved func(mod *Instance, e event.ShieldRemoved)

	// ------------ combat events

	// Called when an attack starts and the attached target is the attacker.
	OnBeforeAttack func(mod *Instance, e event.AttackStart)

	// Called when an attack starts and the attached target is one of the targets being attacked.
	OnBeforeBeingAttacked func(mod *Instance, e event.AttackStart)

	// Called after an attack finishes (after all hits) and the attached target is the attacker
	OnAfterAttack func(mod *Instance, e event.AttackEnd)

	// Called after an attack finishes (after all hits) and the attached target was hit by the attack.
	OnAfterBeingAttacked func(mod *Instance, e event.AttackEnd)

	// Called before any hit occurs and the attached target is the attacker. Hit data is mutable.
	OnBeforeHitAll func(mod *Instance, e event.HitStart)

	// Called before a qualified hit occurs and the attached target is the attacker. "Qualified" hit
	// means it is not of AttackType DOT, PURSUED, or ELEMENT_DAMAGE. Hit data is mutable.
	OnBeforeHit func(mod *Instance, e event.HitStart)

	// called before any hit occurs and the attached target is the defender. Hit data is mutable.
	OnBeforeBeingHitAll func(mod *Instance, e event.HitStart)

	// Called before a qualified hit occurs and the attached target is the defender. "Qualified" hit
	// means it is not of AttackType DOT, PURSUED, or ELEMENT_DAMAGE. Hit data is mutable.
	OnBeforeBeingHit func(mod *Instance, e event.HitStart)

	// Called after any hit occurs and the attached target is the attacker.
	OnAfterHitAll func(mod *Instance, e event.HitEnd)

	// Called after a qualified hit occurs and the attached target is the attacker. "Qualified" hit
	// means it is not of AttackType DOT, PURSUED, or ELEMENT_DAMAGE.
	OnAfterHit func(mod *Instance, e event.HitEnd)

	// Called after any hit occurs and the attached target is the defender.
	OnAfterBeingHitAll func(mod *Instance, e event.HitEnd)

	// Called after a qualified hit occurs and the attached target is the defender. "Qualified" hit
	// means it is not of AttackType DOT, PURSUED, or ELEMENT_DAMAGE.
	OnAfterBeingHit func(mod *Instance, e event.HitEnd)

	// Called before performing a heal and the attached target is the healer. Heal data is mutable.
	OnBeforeDealHeal func(mod *Instance, e *event.HealStart)

	// Called before performing a heal and the attached target is the receiver. Heal data is mutable.
	OnBeforeBeingHeal func(mod *Instance, e *event.HealStart)

	// Called after a heal is performed and the attached target is the healer.
	OnAfterDealHeal func(mod *Instance, e event.HealEnd)

	// Called after a heal is performed and the attached target is the receiver
	OnAfterBeingHeal func(mod *Instance, e event.HealEnd)

	// ------------ sim events

	// Called when an action starts being executed (attack, skill, ult)
	OnBeforeAction func(mod *Instance, e event.ActionStart)

	// Called when an action finishes being executed (attack, skill, ult)
	OnAfterAction func(mod *Instance, e event.ActionEnd)
}

func (mgr *Manager) subscribe() {
	events := mgr.engine.Events()

	// sim events
	events.ActionStart.Subscribe(mgr.actionStart)
	events.ActionEnd.Subscribe(mgr.actionEnd)

	// attribute events
	events.HPChange.Subscribe(mgr.hpChange)
	events.LimboWaitHeal.Subscribe(mgr.limboWaitHeal, 100)
	events.TargetDeath.Subscribe(mgr.targetDeath)
	events.EnergyChange.Subscribe(mgr.energyChange)
	events.StanceChange.Subscribe(mgr.stanceChange)
	events.StanceBreak.Subscribe(mgr.stanceBreak)
	events.StanceReset.Subscribe(mgr.stanceBreakEnd)

	// shield events
	events.ShieldAdded.Subscribe(mgr.shieldAdded)
	events.ShieldRemoved.Subscribe(mgr.shieldRemoved)

	// combat events
	events.AttackStart.Subscribe(mgr.attackStart)
	events.AttackEnd.Subscribe(mgr.attackEnd)
	events.HitStart.Subscribe(mgr.hitStart)
	events.HitEnd.Subscribe(mgr.hitEnd)
	events.HealStart.Subscribe(mgr.healStart, 100)
	events.HealEnd.Subscribe(mgr.healEnd)
}

func (mgr *Manager) emitPropertyChange(target key.TargetID) {
	for _, mod := range mgr.itr(target) {
		f := mod.listeners.OnPropertyChange
		if f != nil {
			f(mod)
		}
	}
}

func (mgr *Manager) emitAdd(target key.TargetID, mod *Instance, chance float64) {
	f := mod.listeners.OnAdd
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierAdded.Emit(event.ModifierAdded{
		Target:   target,
		Modifier: mod.ToModel(),
		Chance:   chance,
	})
}

func (mgr *Manager) emitRemove(target key.TargetID, mods []*Instance) {
	for _, mod := range mods {
		if len(mod.stats) > 0 {
			mgr.emitPropertyChange(target)
		}

		f := mod.listeners.OnRemove
		if f != nil {
			f(mod)
		}
		mgr.engine.Events().ModifierRemoved.Emit(event.ModifierRemoved{
			Target:   target,
			Modifier: mod.ToModel(),
		})
	}
}

func (mgr *Manager) emitExtendDuration(target key.TargetID, mod *Instance, old int) {
	f := mod.listeners.OnExtendDuration
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierExtendedDuration.Emit(event.ModifierExtendedDuration{
		Target:   target,
		Modifier: mod.ToModel(),
		OldValue: old,
		NewValue: mod.duration,
	})
}

func (mgr *Manager) emitExtendCount(target key.TargetID, mod *Instance, old float64) {
	f := mod.listeners.OnExtendCount
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierExtendedCount.Emit(event.ModifierExtendedCount{
		Target:   target,
		Modifier: mod.ToModel(),
		OldValue: old,
		NewValue: mod.count,
	})
}

func (mgr *Manager) attackStart(e event.AttackStart) {
	for _, mod := range mgr.itr(e.Attacker) {
		f := mod.listeners.OnBeforeAttack
		if f != nil {
			f(mod, e)
		}
	}
	for _, target := range e.Targets {
		for _, mod := range mgr.itr(target) {
			f := mod.listeners.OnBeforeBeingAttacked
			if f != nil {
				f(mod, e)
			}
		}
	}
}

func (mgr *Manager) attackEnd(e event.AttackEnd) {
	for _, mod := range mgr.itr(e.Attacker) {
		f := mod.listeners.OnAfterAttack
		if f != nil {
			f(mod, e)
		}
	}
	for _, target := range e.Targets {
		for _, mod := range mgr.itr(target) {
			f := mod.listeners.OnAfterBeingAttacked
			if f != nil {
				f(mod, e)
			}
		}
	}
}

func (mgr *Manager) hitStart(e event.HitStart) {
	qualified := e.Hit.AttackType.IsQualified()
	snapshot := e.Hit.UseSnapshot

	for _, mod := range mgr.itr(e.Attacker) {
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

	for _, mod := range mgr.itr(e.Defender) {
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

func (mgr *Manager) hitEnd(e event.HitEnd) {
	qualified := e.AttackType.IsQualified()
	snapshot := e.UseSnapshot

	for _, mod := range mgr.itr(e.Attacker) {
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

	for _, mod := range mgr.itr(e.Defender) {
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

func (mgr *Manager) healStart(e *event.HealStart) {
	snapshot := e.UseSnapshot

	for _, mod := range mgr.itr(e.Healer.ID()) {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnBeforeDealHeal
		if f != nil {
			f(mod, e)
		}
	}

	for _, mod := range mgr.itr(e.Target.ID()) {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnBeforeBeingHeal
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) healEnd(e event.HealEnd) {
	snapshot := e.UseSnapshot

	for _, mod := range mgr.itr(e.Healer) {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnAfterDealHeal
		if f != nil {
			f(mod, e)
		}
	}

	for _, mod := range mgr.itr(e.Target) {
		if snapshot && !mod.modifySnapshot {
			continue
		}

		f := mod.listeners.OnAfterBeingHeal
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) hpChange(e event.HPChange) {
	for _, mod := range mgr.itr(e.Target) {
		f := mod.listeners.OnHPChange
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) limboWaitHeal(e event.LimboWaitHeal) bool {
	for _, mod := range mgr.itr(e.Target) {
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

func (mgr *Manager) targetDeath(e event.TargetDeath) {
	for _, mod := range mgr.itr(e.Target) {
		f := mod.listeners.OnBeforeDying
		if f != nil {
			f(mod)
		}
	}

	for _, mod := range mgr.itr(e.Killer) {
		f := mod.listeners.OnTriggerDeath
		if f != nil {
			f(mod, e.Target)
		}
	}
}

func (mgr *Manager) energyChange(e event.EnergyChange) {
	for _, mod := range mgr.itr(e.Target) {
		f := mod.listeners.OnEnergyChange
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) stanceChange(e event.StanceChange) {
	for _, mod := range mgr.itr(e.Target) {
		f := mod.listeners.OnStanceChange
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) stanceBreak(e event.StanceBreak) {
	for _, mod := range mgr.itr(e.Source) {
		f := mod.listeners.OnTriggerBreak
		if f != nil {
			f(mod, e.Target)
		}
	}
	for _, mod := range mgr.itr(e.Target) {
		f := mod.listeners.OnBeingBreak
		if f != nil {
			f(mod)
		}
	}
}

func (mgr *Manager) stanceBreakEnd(e event.StanceReset) {
	for _, mod := range mgr.itr(e.Target) {
		f := mod.listeners.OnEndBreak
		if f != nil {
			f(mod)
		}
	}
}

func (mgr *Manager) actionStart(e event.ActionStart) {
	for _, mod := range mgr.itr(e.Owner) {
		f := mod.listeners.OnBeforeAction
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) actionEnd(e event.ActionEnd) {
	for _, mod := range mgr.itr(e.Owner) {
		f := mod.listeners.OnAfterAction
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) shieldAdded(e event.ShieldAdded) {
	for _, mod := range mgr.itr(e.Info.Target) {
		f := mod.listeners.OnShieldAdded
		if f != nil {
			f(mod, e)
		}
	}
}

func (mgr *Manager) shieldRemoved(e event.ShieldRemoved) {
	for _, mod := range mgr.itr(e.Target) {
		f := mod.listeners.OnShieldRemoved
		if f != nil {
			f(mod, e)
		}
	}
}
