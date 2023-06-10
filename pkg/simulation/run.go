package simulation

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type stateFn func(*simulation) (stateFn, error)

func (s *simulation) run() (*model.IterationResult, error) {
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
func initialize(s *simulation) (stateFn, error) {
	// subscribe any internal hooks for core loop
	s.subscribe()

	// enable all registered startup hooks
	for k, hook := range hooks {
		if err := hook(s); err != nil {
			return nil, fmt.Errorf("error executing hook %v", k)
		}
	}

	// TODO: stats collectors should enable here?

	// want to emit after all hooks in the event that they subscribe to these
	s.event.Initialize.Emit(event.InitializeEvent{
		Config: s.cfg,
		Seed:   s.seed,
	})

	// initialize all characters. This is done as part of initialize rather than startBattle due to
	// us wanting to retain characters in the same state between multiple waves (if ever supported)
	for _, char := range s.cfg.Characters {
		id := s.idGen.New()
		if err := s.charManager.AddCharacter(id, char); err != nil {
			return nil, fmt.Errorf("error initializing character %w", err)
		}

		s.targets[id] = TargetCharacter
		s.characters = append(s.characters, id)
	}

	return startBattle, nil
}

// start the battle. this would be called at the beginning of every wave. This should:
//  1. create all enemies
//  2. add all active targets to the turn order (in order of characters > neutrals > enemies)
//  3. emit a BattleStart event containing information about the state of the wave
//  4. execute the engage logic (if this is the first wave of the fight)
//  5. check and run and burst executions (case of bursts occurring at 0 AV)
func startBattle(s *simulation) (stateFn, error) {
	for _, enemy := range s.cfg.Enemies {
		id := s.idGen.New()
		if err := s.enemyManager.AddEnemy(id, enemy); err != nil {
			return nil, fmt.Errorf("error initializing enemy %w", err)
		}

		s.targets[id] = TargetEnemy
		s.enemies = append(s.enemies, id)
	}

	// add all the targets to the turn order at once. Big a mega array to accomplish this
	// the order of the copies matter (want characters > neutrals > enemies for cases of ties)
	all := make([]key.TargetID, len(s.characters)+len(s.enemies)+len(s.neutrals))
	copy(all, s.characters)
	copy(all[len(s.characters):], s.neutrals)
	copy(all[len(s.characters)+len(s.neutrals):], s.enemies)
	s.turnManager.AddTargets(all...)

	// emit BattleStart event to log the "start state" of everything
	charStats := make([]*info.Stats, len(s.characters))
	for i, t := range s.characters {
		charStats[i] = s.attributeService.Stats(t)
	}
	enemyStats := make([]*info.Stats, len(s.enemies))
	for i, t := range s.enemies {
		enemyStats[i] = s.attributeService.Stats(t)
	}
	neutralStats := make([]*info.Stats, len(s.neutrals))
	for i, t := range s.neutrals {
		neutralStats[i] = s.attributeService.Stats(t)
	}
	s.event.BattleStart.Emit(event.BattleStartEvent{
		CharInfo:     s.charManager.Characters(),
		EnemyInfo:    s.enemyManager.Enemies(),
		CharStats:    charStats,
		EnemyStats:   enemyStats,
		NeutralStats: neutralStats,
	})

	return engage, nil
}

// start turn. This will determine which target is taking their turn and will progress time.
// This does not directly emit a TurnStartEvent since the underlying turn manager does that for us.
func beginTurn(s *simulation) (stateFn, error) {
	// determine who's turn it is and increase total AV
	next, av, err := s.turnManager.StartTurn()
	if !s.IsValid(next) || err != nil {
		return nil, fmt.Errorf(
			"unexpected: turn manager returned an invalid target for next turn %w", err)
	}
	s.active = next
	s.totalAV += av

	s.modManager.Tick(s.active, info.TurnStart)
	return phase1, nil
}

// phase1 is the time between the start of the turn and the action being performed. This is when
// stuff like dots deal damage, stance gets reset, and any bursts happen prior to the action.
func phase1(s *simulation) (stateFn, error) {
	// super special case where if they are frozen at the start of the turn, need to skip break
	// reset because that is just how frozen works. Need to make this check BEFORE the tick, since
	// once the tick happens frozen will be removed.
	isFrozen := s.HasBehaviorFlag(s.active, model.BehaviorFlag_STAT_CTRL_FROZEN)

	// tick any modifiers that listen for phase1 (primarily dots)
	// TODO: skillEffect is here invalid. Is there a skill effect for dots?
	s.modManager.Tick(s.active, info.ModifierPhase1)

	// skip all other phase1 logic when frozen and go straight to phase2
	if isFrozen {
		return phase2, nil
	}

	// reset the stance if this is start of enemy turn and their stance is 0
	if s.IsEnemy(s.active) && s.attributeService.Stance(s.active) <= 0 {
		info, err := s.EnemyInfo(s.active)
		if err != nil {
			return nil, fmt.Errorf("error when getting enemy info in phase1 %w", err)
		}
		if err := s.attributeService.SetStance(s.active, s.active, info.MaxStance); err != nil {
			return nil, fmt.Errorf("error when reseting target stance %w", err)
		}
	}

	// skip the action if this target has the DISABLE_ACTION flag
	if s.HasBehaviorFlag(s.active, model.BehaviorFlag_DISABLE_ACTION) {
		return phase2, nil
	}

	return s.executeQueue(info.InsertAbilityPhase1, action)
}

// phase2 is the time after action and before end of turn. This is where follow up attacks occur,
// bursts that occur after action end, and modifiers tick prior to the end of the turn
func phase2(s *simulation) (stateFn, error) {
	// start of phase2 is treated as an "ActionEnd" for any clean up. We have it here instead of
	// inside of action for the cases where the action was skipped.
	s.turnManager.ResetTurn()
	s.modManager.Tick(s.active, info.ActionEnd)

	// execute anything that is in the execution queue. any follow ups, bursts, etc.
	if next, err := s.executeQueue(info.InsertAbilityPhase2, endTurn); next == nil || err != nil {
		return nil, err
	}

	// tick all phase2 modifiers before finally ending the turn
	s.modManager.Tick(s.active, info.ModifierPhase2)
	return endTurn, nil
}

// finalize that this is the end of the turn. Mainly just emitting the turn end event
func endTurn(s *simulation) (stateFn, error) {
	// TODO: cleanup check (reuse code from death subscription)

	// emit TurnEnd event to log the current state of all remaining targets
	charStats := make([]*info.Stats, len(s.characters))
	for i, t := range s.characters {
		charStats[i] = s.attributeService.Stats(t)
	}
	enemyStats := make([]*info.Stats, len(s.enemies))
	for i, t := range s.enemies {
		enemyStats[i] = s.attributeService.Stats(t)
	}
	neutralStats := make([]*info.Stats, len(s.neutrals))
	for i, t := range s.neutrals {
		neutralStats[i] = s.attributeService.Stats(t)
	}
	s.event.TurnEnd.Emit(event.TurnEndEvent{
		Characters: charStats,
		Enemies:    enemyStats,
		Neutrals:   neutralStats,
	})

	return s.exitCheck(beginTurn)
}

// check if we want to exit the sim. If not, return the next state that was passed in
func (s *simulation) exitCheck(next stateFn) (stateFn, error) {
	var reason model.TerminationReson
	if len(s.characters) == 0 {
		reason = model.TerminationReson_BATTLE_LOSS
	} else if len(s.enemies) == 0 {
		reason = model.TerminationReson_BATTLE_WIN
	} else if int(s.totalAV/100) >= int(s.cfg.Settings.CycleLimit) {
		reason = model.TerminationReson_TIMEOUT
	}

	if reason != model.TerminationReson_INVALID_TERMINATION {
		s.event.Termination.Emit(event.TerminationEvent{
			TotalAV: s.totalAV,
			Reason:  reason,
		})
		return nil, nil
	}
	return next, nil
}
