package simulation

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/model"
)

type stateFn func(*Simulation) (stateFn, error)

func (s *Simulation) run() (*model.IterationResult, error) {
	var err error
	//TODO: per Kyle this is totally unnecessary; for that reason alone this will stay
	//because what's better than another future Kyle problem?
	for state := initState; state != nil; {
		state, err = state(s)
		if err != nil {
			//handle error here
			return nil, err
		}
	}
	return nil, nil
}

func initState(s *Simulation) (stateFn, error) {
	s.setupServices()
	return startBattle, nil
}

func startBattle(s *Simulation) (stateFn, error) {
	burstCheck(s)
	executeQueue(s)
	return beginTurn, nil
}

func beginTurn(s *Simulation) (stateFn, error) {
	//AVUpdate
	next, err := s.turnManager.StartTurn()
	if !s.IsValid(next) || err != nil {
		return nil, fmt.Errorf(
			"unexpected: turn manager returned an invalid target for next turn %w", err)
	}

	//DetermineTurn

	burstCheck(s)
	executeQueue(s)

	return action, nil
}

func action(s *Simulation) (stateFn, error) {
	// TODO: FindAction (NextAction?) / ExecuteAction
	burstCheck(s)
	return endTurn, nil
}

func endTurn(s *Simulation) (stateFn, error) {
	executeQueue(s)
	return exitCheck, nil
}

func exitCheck(s *Simulation) (stateFn, error) {
	//TODO: just call it quits for now
	if s.cfg.Settings.TtkMode {
		return nil, nil
	}
	if int(s.turnManager.TotalAV()/100) >= int(s.cfg.Settings.CycleLimit) {
		return nil, nil
	}
	return beginTurn, nil
}

func burstCheck(s *Simulation) bool {
	bursts := s.eval.BurstCheck()
	for _, value := range bursts {
		fmt.Printf("burst: target (%v) type (%v)\n", value.Target, value.Type)
		// TODO: execute the burst?
	}
	return len(bursts) > 0
}

func executeQueue(s *Simulation) {
	for {
		// TODO: ExecuteAction
		if !burstCheck(s) { // nothing to execute
			break
		}
	}
}
