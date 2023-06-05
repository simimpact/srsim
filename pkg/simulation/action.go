package simulation

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
)

// run engagement logic for the first wave of a battle. This is any techniques + if the engagement
// of the battle should be a "weakness" engage (player's favor) or "ambush" engage (enemy's favor)
func engage(s *simulation) (stateFn, error) {
	// TODO: waveCount & only do this call if is first wave
	// TODO: execute any techniques + engagement logic
	// TODO: weakness engage vs ambush
	// TODO: emit EngageEvent
	return nil, nil
}

func action(s *simulation) (stateFn, error) {
	// TODO: ActionStartEvent
	// TODO: attempt to execute an action
	// TODO: loop over action execution attempts until terminating action
	// TODO: SP increase depending on action type?
	// TODO: ActionEndEvent
	// TODO: FindAction (NextAction?) / ExecuteAction
	burstCheck(s)
	return phase2, nil
}

// execute everything on the queue. After queue execution is complete, return the next stateFn
// as the next state to run. This logic will run the exitCheck on each execution. If an exit
// condition is met, will return that state instead
func (s *simulation) executeQueue(phase info.BattlePhase, next stateFn) (stateFn, error) {
	// if active is not a character, cannot prform any queue execution until after ActionEnd
	if phase < info.ActionEnd && !s.IsCharacter(s.active) {
		burstCheck(s)
		return s.exitCheck(next)
	}

	// loop over each entry of the queue, executing and then performing the necessary cleanup
	// After cleanup, run burstCheck which may add more to the queue

	// burstCheck(s)
	// while !s.queue.IsEmpty() {
	// 	 e := s.queue.Pop()
	//	 check abort flags and skip if a flag matches
	// 	 emit start event depending on type
	//   e.execute()
	//	 emit AttackEnd event if it has not ended yet
	//	 emit end event depending on type
	//	 burstCheck(s)
	//	 exitCheck(s)
	// }
	return next, nil
}

func burstCheck(s *simulation) bool {
	bursts := s.eval.BurstCheck()
	for _, value := range bursts {
		fmt.Printf("burst: target (%v) type (%v)\n", value.Target, value.Type)
		// TODO: execute the burst?
	}
	return len(bursts) > 0
}
