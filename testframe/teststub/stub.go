package teststub

import (
	"fmt"
	"time"

	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/logging"
	"github.com/simimpact/srsim/pkg/logic/gcs/eval"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/simimpact/srsim/testframe/testcfg"
	"github.com/stretchr/testify/suite"
)

type Stub struct {
	suite.Suite
	autoContinue  bool
	eventPipe     chan handler.Event
	haltSignaller chan struct{}
	cfg           *model.SimConfig
	eval          *eval.Eval
	simulator     *simulation.Simulation
}

func (s *Stub) SetupTest() {
	s.eventPipe = make(chan handler.Event)
	s.haltSignaller = make(chan struct{})
	s.cfg = testcfg.TestConfigTwoElites()
	s.eval = testcfg.StandardTestEval()
	s.autoContinue = true
}

func (s *Stub) TearDownTest() {
	fmt.Println("Test Finished")
	logging.Singleton = logging.NewNilLogger()
	select {
	case <-s.eventPipe:
		s.haltSignaller <- struct{}{}
	default:
	}
	close(s.eventPipe)
	close(s.haltSignaller)
}

func (s *Stub) StartSimulation() {
	l := NewTestLogger(s.eventPipe, s.haltSignaller)
	logging.InitLogger(l)
	s.simulator = simulation.NewSimulation(s.cfg, s.eval, 0)
	go func() {
		itres, err := s.simulator.Run()
		if err != nil {
			s.FailNow("Simulation run error %v", err)
		}
		fmt.Println(itres)
	}()
}

// EventChecker returns True if the checker passes, and False if the checker fails.
// All checkers must pass for an Expect to conclude.
// If any checker fails, subsequent checkers will not run. If it also returns an error, the Expect fails.
// If it fails but does not return an error, Expect will continue to run.
type EventChecker func(e handler.Event) (bool, error)

func (s *Stub) Expect(checkers ...EventChecker) {
	for {
		var e handler.Event
		select {
		case e = <-s.eventPipe:
		case <-time.After(1 * time.Second):
			s.FailNow("Event not intercepted")
		}
		var toContinue bool
		var err error
		for i := range checkers {
			toContinue, err = checkers[i](e)
			if toContinue {
				continue
			}

			if err != nil {
				s.FailNow("Event Checker err", err)
				return
			}
			break
		}
		if s.autoContinue || !toContinue {
			s.haltSignaller <- struct{}{}
		}
		if toContinue {
			return
		}
	}
}

// Continue resumes the simulation. This must be called after each Expect if AutoContinue is disabled.
func (s *Stub) Continue() {
	s.haltSignaller <- struct{}{}
}

func (s *Stub) SetAutoContinue(cont bool) {
	s.autoContinue = cont
}
