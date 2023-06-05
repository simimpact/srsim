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

// executes the initial function + any chain it returns
func (s *simulation) execute(init stateFn) error {
	var err error
	for state := init; state != nil; {
		state, err = state(s)
		if err != nil {
			//handle error here
			return err
		}
	}
	return nil
}

func initialize(s *simulation) (stateFn, error) {
	// TODO: startup hooks
	// TODO: death subscription

	s.event.Initialize.Emit(event.InitializeEvent{
		Config: s.cfg,
		Seed:   s.seed,
	})

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

	// TODO: figure out data to add to battle start
	s.event.BattleStart.Emit(struct{}{})

	if err := s.execute(engage); err != nil {
		return nil, fmt.Errorf("error attempting to perform engage %w", err)
	}

	if err := executeQueue(s, info.BattleStart); err != nil {
		return nil, fmt.Errorf("error attempting to execute queue in battle start %w", err)
	}
	return beginTurn, nil
}

func engage(s *simulation) (stateFn, error) {
	// TODO: waveCount & only do this call if is first wave
	// TODO: execute any techniques + engagement logic
	// TODO: weakness engage vs ambush
	// TODO: emit EngageEvent
	return nil, nil
}

func beginTurn(s *simulation) (stateFn, error) {
	// determine who's turn it is and increase AV
	next, av, err := s.turnManager.StartTurn()
	if !s.IsValid(next) || err != nil {
		return nil, fmt.Errorf(
			"unexpected: turn manager returned an invalid target for next turn %w", err)
	}
	s.active = next
	s.totalAV += av

	s.modManager.Tick(s.active, info.TurnStart)

	// TODO: unsure if we want a burstCheck on TurnStart?
	return modifierPhase1, nil
}

func modifierPhase1(s *simulation) (stateFn, error) {
	s.modManager.Tick(s.active, info.ModifierPhase1)

	// special case of frozen where stance reset never happens
	// TODO: is this a ctrl thing, does stun do this too?
	if s.HasBehaviorFlag(s.active, model.BehaviorFlag_STAT_CTRL_FROZEN) {
		return actionEnd, nil
	}

	// reset the stance if start of enemy turn and stance is 0
	if s.IsEnemy(s.active) && s.attributeService.Stance(s.active) <= 0 {
		// TODO: get enemy info to know what their max stance is
		if err := s.attributeService.SetStance(s.active, s.active, 5000); err != nil {
			return nil, fmt.Errorf("error to reset target stance %w", err)
		}
	}

	// skip the action if this target has the DISABLE_ACTION flag
	if s.HasBehaviorFlag(s.active, model.BehaviorFlag_DISABLE_ACTION) {
		return actionEnd, nil
	}

	return insertPhase1, nil
}

func insertPhase1(s *simulation) (stateFn, error) {
	if err := executeQueue(s, info.InsertAbilityPhase1); err != nil {
		return nil, fmt.Errorf("error attempting to execute queue in insert phase 1 %w", err)
	}
	return action, nil
}

func action(s *simulation) (stateFn, error) {
	// TODO: ActionStartEvent
	// TODO: attempt to execute an action
	// TODO: loop over action execution attempts until terminating action
	// TODO: SP increase depending on action type?
	// TODO: ActionEndEvent
	// TODO: FindAction (NextAction?) / ExecuteAction
	burstCheck(s)
	return actionEnd, nil
}

func actionEnd(s *simulation) (stateFn, error) {
	s.turnManager.ResetTurn()
	s.modManager.Tick(s.active, info.ActionEnd)
	return nil, nil
}

func insertPhase2(s *simulation) (stateFn, error) {
	if err := executeQueue(s, info.InsertAbilityPhase2); err != nil {
		return nil, fmt.Errorf("error attempting to execute queue in insert phase 1 %w", err)
	}
	return modifierPhase2, nil
}

func modifierPhase2(s *simulation) (stateFn, error) {
	s.modManager.Tick(s.active, info.ModifierPhase2)
	return endTurn, nil
}

func endTurn(s *simulation) (stateFn, error) {
	return exitCheck, nil
}

func exitCheck(s *simulation) (stateFn, error) {
	//TODO: just call it quits for now
	if s.cfg.Settings.TtkMode {
		return nil, nil
	}
	if int(s.totalAV/100) >= int(s.cfg.Settings.CycleLimit) {
		return nil, nil
	}
	return beginTurn, nil
}

func burstCheck(s *simulation) bool {
	bursts := s.eval.BurstCheck()
	for _, value := range bursts {
		fmt.Printf("burst: target (%v) type (%v)\n", value.Target, value.Type)
		// TODO: execute the burst?
	}
	return len(bursts) > 0
}

func executeQueue(s *simulation, phase info.BattlePhase) error {
	// if active is not a character, cannot prform any queue execution until after ActionEnd
	if phase < info.ActionEnd && !s.IsCharacter(s.active) {
		burstCheck(s)
		return nil
	}

	// loop over each entry of the queue, executing and then performing the necessary cleanup
	// After cleanup, run burstCheck which may add more to the queue
	// stop looping when queue is empty after a burst check

	// burstCheck(s)
	// while !s.queue.IsEmpty() {
	// 	 e := s.queue.Pop()
	// 	 emit start event depending on type
	//   e.execute()
	//	 emit AttackEnd event if it has not ended yet
	//	 emit end event depending on type
	//	 burstCheck(s)
	// }
	return nil
}
