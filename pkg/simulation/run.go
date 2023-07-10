package simulation

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/hook"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type stateFn func(*Simulation) (stateFn, error)

func (sim *Simulation) Run() (*model.IterationResult, error) {
	var err error
	// TODO: per Kyle this is totally unnecessary; for that reason alone this will stay
	// because what's better than another future Kyle problem?
	for state := initialize; state != nil; {
		state, err = state(sim)
		if err != nil {
			// handle error here
			return nil, err
		}
	}
	// TODO: create IterationResult
	return nil, nil
}

// initialize the sim and create characters from the config to prep for execution
func initialize(sim *Simulation) (stateFn, error) {
	// subscribe any internal hooks for core loop
	sim.subscribe()

	// enable all registered startup hooks
	for k, hook := range hook.StartupHooks() {
		if err := hook(sim); err != nil {
			return nil, fmt.Errorf("error executing hook %v", k)
		}
	}

	// TODO: stats collectors should enable here?

	// want to emit after all hooks in the event that they subscribe to these
	sim.Event.Initialize.Emit(event.Initialize{
		Config: sim.cfg,
		Seed:   sim.seed,
	})

	// initialize all characters. This is done as part of initialize rather than startBattle due to
	// us wanting to retain characters in the same state between multiple waves (if ever supported)
	for _, char := range sim.cfg.Characters {
		id := sim.IDGen.New()
		sim.Targets[id] = info.ClassCharacter
		sim.characters = append(sim.characters, id)

		if err := sim.Char.AddCharacter(id, char); err != nil {
			return nil, fmt.Errorf("error initializing character %w", err)
		}
	}

	// emit event saying that all characters are initialized
	chars := make([]event.CharInfo, 0, len(sim.characters))
	for _, id := range sim.characters {
		info, _ := sim.Char.Info(id)
		chars = append(chars, event.CharInfo{
			ID:   id,
			Info: &info,
		})
	}
	sim.Event.CharactersAdded.Emit(event.CharactersAdded{
		Characters: chars,
	})

	// run the script to register callbacks
	if err := sim.eval.Init(sim); err != nil {
		return nil, err
	}

	return startBattle, nil
}

// start the battle. this would be called at the beginning of every wave. This should:
//  1. create all enemies
//  2. add all active targets to the turn order (in order of characters > neutrals > enemies)
//  3. emit a BattleStart event containing information about the state of the wave
//  4. execute the engage logic (if this is the first wave of the fight)
//  5. check and run and burst executions (case of bursts occurring at 0 AV)
func startBattle(sim *Simulation) (stateFn, error) {
	for _, enemy := range sim.cfg.Enemies {
		id := sim.IDGen.New()
		sim.Targets[id] = info.ClassEnemy
		sim.enemies = append(sim.enemies, id)

		if err := sim.Enemy.AddEnemy(id, enemy); err != nil {
			return nil, fmt.Errorf("error initializing enemy %w", err)
		}
	}

	// emit event saying that all enemies are initialized
	enemies := make([]event.EnemyInfo, 0, len(sim.enemies))
	for _, id := range sim.enemies {
		info, _ := sim.Enemy.Info(id)
		enemies = append(enemies, event.EnemyInfo{
			ID:   id,
			Info: &info,
		})
	}
	sim.Event.EnemiesAdded.Emit(event.EnemiesAdded{
		Enemies: enemies,
	})

	// add all the targets to the turn order at once. Big a mega array to accomplish this
	// the order of the copies matter (want characters > neutrals > enemies for cases of ties)
	all := make([]key.TargetID, len(sim.characters)+len(sim.enemies)+len(sim.neutrals))
	copy(all, sim.characters)
	copy(all[len(sim.characters):], sim.neutrals)
	copy(all[len(sim.characters)+len(sim.neutrals):], sim.enemies)
	sim.Turn.AddTargets(all...)

	// emit BattleStart event to log the "start state" of everything
	snap := sim.createSnapshot()
	sim.Event.BattleStart.Emit(event.BattleStart{
		CharInfo:     sim.Char.Characters(),
		EnemyInfo:    sim.Enemy.Enemies(),
		CharStats:    snap.characters,
		EnemyStats:   snap.enemies,
		NeutralStats: snap.neutrals,
	})

	return engage, nil
}

// run engagement logic for the first wave of a battle. This is any techniques + if the engagement
// of the battle should be a "weakness" engage (player's favor) or "ambush" engage (enemy's favor)
func engage(sim *Simulation) (stateFn, error) {
	// TODO: waveCount & only do this call if is first wave
	// TODO: execute any techniques + engagement logic
	// TODO: weakness engage vs ambush
	// TODO: emit EngageEvent
	return sim.executeQueue(info.BattleStart, beginTurn)
}

// start turn. This will determine which target is taking their turn and will progress time.
// This does not directly emit a TurnStartEvent since the underlying turn manager does that for us.
func beginTurn(sim *Simulation) (stateFn, error) {
	// determine who's turn it is and increase total AV
	next, av, turnOrder, err := sim.Turn.StartTurn()
	if !sim.IsValid(next) || err != nil {
		return nil, fmt.Errorf(
			"unexpected: turn manager returned an invalid target for next turn %w", err)
	}
	sim.Active = next
	sim.TotalAV += av

	sim.Event.TurnStart.Emit(event.TurnStart{
		Active:     next,
		TargetType: sim.Targets[next],
		DeltaAV:    av,
		TotalAV:    sim.TotalAV,
		TurnOrder:  turnOrder,
	})
	sim.Modifier.Tick(sim.Active, info.TurnStart)
	return phase1, nil
}

// phase1 is the time between the start of the turn and the action being performed. This is when
// stuff like dots deal damage, stance gets reset, and any bursts happen prior to the action.
func phase1(sim *Simulation) (stateFn, error) {
	sim.Event.Phase1Start.Emit(event.Phase1Start{})

	// tick any modifiers that listen for phase1 (primarily dots)
	sim.Modifier.Tick(sim.Active, info.ModifierPhase1)
	sim.deathCheck(false)

	// skip the action if this target has the DISABLE_ACTION flag
	if sim.HasBehaviorFlag(sim.Active, model.BehaviorFlag_DISABLE_ACTION) {
		return phase2, nil
	}

	// reset the stance if this is start of enemy turn and their stance is 0
	if sim.IsEnemy(sim.Active) && sim.Attr.Stance(sim.Active) <= 0 {
		info, err := sim.Enemy.Info(sim.Active)
		if err != nil {
			return nil, fmt.Errorf("error when getting enemy info in phase1 %w", err)
		}
		if err := sim.Attr.SetStance(sim.Active, sim.Active, info.MaxStance); err != nil {
			return nil, fmt.Errorf("error when reseting target stance %w", err)
		}
	}

	next, err := sim.executeQueue(info.InsertAbilityPhase1, action)
	if err == nil {
		sim.Event.Phase1End.Emit(event.Phase1End{})
	}
	return next, err
}

// actually execute the action for this turn and then move on to phase2 once done
func action(sim *Simulation) (stateFn, error) {
	if err := sim.executeAction(sim.Active, false); err != nil {
		return nil, fmt.Errorf("unknown error executing action %w", err)
	}
	sim.deathCheck(false)
	return phase2, nil
}

// phase2 is the time after action and before end of turn. This is where follow up attacks occur,
// bursts that occur after action end, and modifiers tick prior to the end of the turn
func phase2(sim *Simulation) (stateFn, error) {
	// start of phase2 is treated as an "ActionEnd" for any clean up. We have it here instead of
	// inside of action for the cases where the action was skipped.
	sim.Turn.ResetTurn()
	sim.Modifier.Tick(sim.Active, info.ActionEnd)
	sim.Event.Phase2Start.Emit(event.Phase2Start{})

	// execute anything that is in the execution queue. any follow ups, bursts, etc.
	if next, err := sim.executeQueue(info.InsertAbilityPhase2, endTurn); next == nil || err != nil {
		return nil, err
	}

	// tick all phase2 modifiers before finally ending the turn
	sim.Modifier.Tick(sim.Active, info.ModifierPhase2)
	sim.Event.Phase2End.Emit(event.Phase2End{})
	return endTurn, nil
}

// finalize that this is the end of the turn. Mainly just emitting the turn end event
func endTurn(sim *Simulation) (stateFn, error) {
	sim.deathCheck(true) // treat limbo'd targets as dead for edge case

	// emit TurnEnd event to log the current state of all remaining targets
	snap := sim.createSnapshot()
	sim.Event.TurnEnd.Emit(event.TurnEnd{
		Characters: snap.characters,
		Enemies:    snap.enemies,
		Neutrals:   snap.neutrals,
	})

	return sim.exitCheck(beginTurn)
}

// check if we want to exit the sim. If not, return the next state that was passed in
func (sim *Simulation) exitCheck(next stateFn) (stateFn, error) {
	var reason model.TerminationReason
	switch {
	case len(sim.characters) == 0:
		reason = model.TerminationReason_BATTLE_LOSS
	case len(sim.enemies) == 0:
		reason = model.TerminationReason_BATTLE_WIN
	case int(sim.TotalAV/100) >= int(sim.cfg.Settings.CycleLimit):
		reason = model.TerminationReason_TIMEOUT
	}

	if reason != model.TerminationReason_INVALID_TERMINATION {
		sim.Event.Termination.Emit(event.Termination{
			TotalAV: sim.TotalAV,
			Reason:  reason,
		})
		return nil, nil
	}
	return next, nil
}
