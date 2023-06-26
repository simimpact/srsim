package teststub

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/gcs/eval"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/simimpact/srsim/testframe/testcfg"
	"github.com/stretchr/testify/suite"
	"time"
)

type Stub struct {
	suite.Suite
	eventPipe chan handler.Event
	cfg       *model.SimConfig
	eval      *eval.Eval
	simulator *simulation.Simulation
}

func (s *Stub) SetupTest() {
	s.eventPipe = make(chan handler.Event)
	s.cfg = testcfg.TestConfigTwoElites()
	s.eval = testcfg.StandardTestEval()
}

func (s *Stub) TearDownTest() {

}

func (s *Stub) StartSimulation() {
	l := NewTestLogger(s.eventPipe)
	go func() {
		simulation.RunWithLog(l, s.cfg, s.eval, 0)
	}()
}

type EventChecker func(e handler.Event) (bool, error)

func (s *Stub) Expect(checkers ...EventChecker) {
	for {
		var e handler.Event
		select {
		case e = <-s.eventPipe:
		case <-time.After(1 * time.Second):
			s.FailNow("Event not intercepted")
		}
		for i := range checkers {
			toContinue, err := checkers[i](e)
			if toContinue {
				continue
			} else {
				if err != nil {
					s.FailNow("Event Checker err %v", err)
					return
				}
				break
			}
		}
	}
}
