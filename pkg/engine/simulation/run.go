package simulation

import (
	"errors"

	"github.com/simimpact/srsim/pkg/key"
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
	return beginTurn, nil
}

func beginTurn(s *Simulation) (stateFn, error) {
	//AVUpdate
	next := s.turnManager.AdvanceTurn()
	if next == key.TargetInvalid {
		return nil, errors.New("unexpected: turn manager returned an invalid target for next turn")
	}

	//DetermineTurn

	return action, nil
}

func action(s *Simulation) (stateFn, error) {
	return endTurn, nil
}

func endTurn(s *Simulation) (stateFn, error) {
	return exitCheck, nil
}

func exitCheck(s *Simulation) (stateFn, error) {
	//TODO: just call it quits for now
	if s.cfg.Settings.TtkMode {
		return nil, nil
	}
	if s.turnManager.CurrentCycle() >= int(s.cfg.Settings.CycleLimit) {
		return nil, nil
	}
	return beginTurn, nil
}
