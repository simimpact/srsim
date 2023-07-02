package teststub

import "github.com/simimpact/srsim/pkg/key"

// Continue resumes the simulation. This must be called after each Expect if AutoContinue is disabled.
func (s *Stub) Continue() {
	s.haltSignaller <- struct{}{}
}

func (s *Stub) SetAutoContinue(cont bool) {
	s.autoContinue = cont
}

func (s *Stub) SetAutoRun(cont bool) {
	s.autoRun = cont
}

// TerminateRun pipes a command with an astronomical AV to immediately exceed the cycle limit, ending the run
func (s *Stub) TerminateRun() {
	go func() {
		s.turnPipe <- TurnCommand{Next: s.Characters.GetCharacterID(0), Av: 100000}
	}()
}

// NextTurn queues the next turn without using up any AV cost
func (s *Stub) NextTurn(id key.TargetID) {
	go func() {
		s.turnPipe <- TurnCommand{Next: id, Av: 0}
	}()
}
