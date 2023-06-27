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

func (s *Simulation) Run() (*model.IterationResult, error) {
	var err error
	//TODO: per Kyle this is totally unnecessary; for that reason alone this will stay
	//because what's better than another future Kyle problem?
	for state := initialize; state != nil; {
		state, err = state(s)
		if err != nil {
			//handle error here
			return nil, err
		}
	}
	// TODO: create IterationResult
	return nil, nil
}

// initialize the sim and create characters from the config to prep for execution
func initialize(s *Simulation) (stateFn, error) {
	// subscribe any internal hooks for core loop
	s.subscribe()

	// enable all registered startup hooks
	for k, hook := range hook.StartupHooks() {
		if err := hook(s); err != nil {
			return nil, fmt.Errorf("error executing hook %v", k)
		}
	}

	// TODO: stats collectors should enable here?

	// want to emit after all hooks in the event that they subscribe to these
	s.Event.Initialize.Emit(event.InitializeEvent{
		Config: s.cfg,
		Seed:   s.seed,
	})

	// initialize all characters. This is done as part of initialize rather than startBattle due to
	// us wanting to retain characters in the same state between multiple waves (if ever supported)
	for _, char := range s.cfg.Characters {
		id := s.IdGen.New()
		if err := s.Char.AddCharacter(id, char); err != nil {
			return nil, fmt.Errorf("error initializing character %w", err)
		}

		s.Targets[id] = info.ClassCharacter
		s.characters = append(s.characters, id)
	}

	// run the script to register callbacks
	err := s.eval.Init()
	if err != nil {
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
func startBattle(s *Simulation) (stateFn, error) {
	for _, enemy := range s.cfg.Enemies {
		id := s.IdGen.New()
		if err := s.Enemy.AddEnemy(id, enemy); err != nil {
			return nil, fmt.Errorf("error initializing enemy %w", err)
		}

		s.Targets[id] = info.ClassEnemy
		s.enemies = append(s.enemies, id)
	}

	// add all the targets to the turn order at once. Big a mega array to accomplish this
	// the order of the copies matter (want characters > neutrals > enemies for cases of ties)
	all := make([]key.TargetID, len(s.characters)+len(s.enemies)+len(s.neutrals))
	copy(all, s.characters)
	copy(all[len(s.characters):], s.neutrals)
	copy(all[len(s.characters)+len(s.neutrals):], s.enemies)
	s.Turn.AddTargets(all...)

	// emit BattleStart event to log the "start state" of everything
	snap := s.createSnapshot()
	s.Event.BattleStart.Emit(event.BattleStartEvent{
		CharInfo:     s.Char.Characters(),
		EnemyInfo:    s.Enemy.Enemies(),
		CharStats:    snap.characters,
		EnemyStats:   snap.enemies,
		NeutralStats: snap.neutrals,
	})

	return engage, nil
}

// run engagement logic for the first wave of a battle. This is any techniques + if the engagement
// of the battle should be a "weakness" engage (player's favor) or "ambush" engage (enemy's favor)
func engage(s *Simulation) (stateFn, error) {
	// TODO: waveCount & only do this call if is first wave
	// TODO: execute any techniques + engagement logic
	// TODO: weakness engage vs ambush
	// TODO: emit EngageEvent
	return s.executeQueue(info.BattleStart, beginTurn)
}

// start turn. This will determine which target is taking their turn and will progress time.
// This does not directly emit a TurnStartEvent since the underlying turn manager does that for us.
func beginTurn(s *Simulation) (stateFn, error) {
	// determine who's turn it is and increase total AV
	next, av, err := s.Turn.StartTurn()
	if !s.IsValid(next) || err != nil {
		return nil, fmt.Errorf(
			"unexpected: turn manager returned an invalid target for next turn %w", err)
	}
	s.Active = next
	s.TotalAV += av

	s.Modifier.Tick(s.Active, info.TurnStart)
	return phase1, nil
}

// phase1 is the time between the start of the turn and the action being performed. This is when
// stuff like dots deal damage, stance gets reset, and any bursts happen prior to the action.
func phase1(s *Simulation) (stateFn, error) {
	// tick any modifiers that listen for phase1 (primarily dots)
	s.Modifier.Tick(s.Active, info.ModifierPhase1)

	// skip the action if this target has the DISABLE_ACTION flag
	if s.HasBehaviorFlag(s.Active, model.BehaviorFlag_DISABLE_ACTION) {
		return phase2, nil
	}

	// reset the stance if this is start of enemy turn and their stance is 0
	if s.IsEnemy(s.Active) && s.Attr.Stance(s.Active) <= 0 {
		info, err := s.Enemy.Info(s.Active)
		if err != nil {
			return nil, fmt.Errorf("error when getting enemy info in phase1 %w", err)
		}
		if err := s.Attr.SetStance(s.Active, s.Active, info.MaxStance); err != nil {
			return nil, fmt.Errorf("error when reseting target stance %w", err)
		}
	}

	return s.executeQueue(info.InsertAbilityPhase1, action)
}

// actually execute the action for this turn and then move on to phase2 once done
func action(s *Simulation) (stateFn, error) {
	if err := s.executeAction(s.Active, false); err != nil {
		return nil, fmt.Errorf("unknown error executing action %w", err)
	}
	return phase2, nil
}

// phase2 is the time after action and before end of turn. This is where follow up attacks occur,
// bursts that occur after action end, and modifiers tick prior to the end of the turn
func phase2(s *Simulation) (stateFn, error) {
	// start of phase2 is treated as an "ActionEnd" for any clean up. We have it here instead of
	// inside of action for the cases where the action was skipped.
	s.Turn.ResetTurn()
	s.Modifier.Tick(s.Active, info.ActionEnd)

	// execute anything that is in the execution queue. any follow ups, bursts, etc.
	if next, err := s.executeQueue(info.InsertAbilityPhase2, endTurn); next == nil || err != nil {
		return nil, err
	}

	// tick all phase2 modifiers before finally ending the turn
	s.Modifier.Tick(s.Active, info.ModifierPhase2)
	return endTurn, nil
}

// finalize that this is the end of the turn. Mainly just emitting the turn end event
func endTurn(s *Simulation) (stateFn, error) {
	// check for special case where a target was supposed to be revived but never did (reviver died?)
	s.deathCheck(s.characters)
	s.deathCheck(s.enemies)

	// emit TurnEnd event to log the current state of all remaining targets
	snap := s.createSnapshot()
	s.Event.TurnEnd.Emit(event.TurnEndEvent{
		Characters: snap.characters,
		Enemies:    snap.enemies,
		Neutrals:   snap.neutrals,
	})

	return s.exitCheck(beginTurn)
}

// check if we want to exit the sim. If not, return the next state that was passed in
func (s *Simulation) exitCheck(next stateFn) (stateFn, error) {
	var reason model.TerminationReason
	if len(s.characters) == 0 {
		reason = model.TerminationReason_BATTLE_LOSS
	} else if len(s.enemies) == 0 {
		reason = model.TerminationReason_BATTLE_WIN
	} else if int(s.TotalAV/100) >= int(s.cfg.Settings.CycleLimit) {
		reason = model.TerminationReason_TIMEOUT
	}

	if reason != model.TerminationReason_INVALID_TERMINATION {
		s.Event.Termination.Emit(event.TerminationEvent{
			TotalAV: s.TotalAV,
			Reason:  reason,
		})
		return nil, nil
	}
	return next, nil
}
