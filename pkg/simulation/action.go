package simulation

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/queue"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// TODO: Unknown TC
//	- Does ActionStart & ActionEnd happen for inserted actions? If  so, this means
//		LC such as "In the Name of the World" buff these insert actions
//	- Do Insert abilities (follow up attacks, counters, etc) count as an Action (similar to above)?

// TODO: support more than just characters
func (sim *simulation) InsertAction(target key.TargetID) error {
	if !sim.IsCharacter(target) {
		return fmt.Errorf("target must be a character to insert an action %v", target)
	}

	sim.queue.Insert(queue.Task{
		Source:   target,
		Priority: info.CharInsertAction,
		AbortFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_CTRL,
			model.BehaviorFlag_DISABLE_ACTION,
		},
		Execute: func() { sim.executeCharAction(target, true) },
	})
	return nil
}

func (sim *simulation) InsertAbility(i info.Insert) {
	sim.queue.Insert(queue.Task{
		Source:     i.Source,
		Priority:   i.Priority,
		AbortFlags: i.AbortFlags,
		Execute:    func() { sim.executeInsert(i) },
	})
}

func (sim *simulation) ultCheck() {
	bursts := sim.eval.BurstCheck()
	for _, value := range bursts {
		sim.queue.Insert(queue.Task{
			Source:   value.Target,
			Priority: info.CharInsertUlt,
			AbortFlags: []model.BehaviorFlag{
				model.BehaviorFlag_STAT_CTRL,
				model.BehaviorFlag_DISABLE_ACTION,
			},
			Execute: func() { sim.executeCharUlt(value.Target) },
		})
	}
}

// run engagement logic for the first wave of a battle. This is any techniques + if the engagement
// of the battle should be a "weakness" engage (player's favor) or "ambush" engage (enemy's favor)
func engage(s *simulation) (stateFn, error) {
	// TODO: waveCount & only do this call if is first wave
	// TODO: execute any techniques + engagement logic
	// TODO: weakness engage vs ambush
	// TODO: emit EngageEvent
	return s.executeQueue(info.BattleStart, beginTurn)
}

func action(s *simulation) (stateFn, error) {
	switch s.targets[s.active] {
	case TargetCharacter:
		s.executeCharAction(s.active, false)
	case TargetEnemy:
		// TODO:
	case TargetNeutral:
		// TODO:
	}

	s.skillEffect = model.SkillEffect_INVALID_SKILL_EFFECT
	s.ultCheck()
	return phase2, nil
}

// execute everything on the queue. After queue execution is complete, return the next stateFn
// as the next state to run. This logic will run the exitCheck on each execution. If an exit
// condition is met, will return that state instead
func (s *simulation) executeQueue(phase info.BattlePhase, next stateFn) (stateFn, error) {
	// always ult check when calling executeQueue
	s.ultCheck()

	// if active is not a character, cannot prform any queue execution until after ActionEnd
	if phase < info.ActionEnd && !s.IsCharacter(s.active) {
		return s.exitCheck(next)
	}

	for !s.queue.IsEmpty() {
		insert := s.queue.Pop()

		// if source has no HP, skip this insert
		if s.attributeService.HPRatio(insert.Source) <= 0 {
			continue
		}

		// if the source has an abort flag, skip this insert
		if s.HasBehaviorFlag(insert.Source, insert.AbortFlags...) {
			continue
		}

		insert.Execute()

		// attempt to exit. If can exit, stop sim now
		if next, err := s.exitCheck(next); next == nil || err != nil {
			return next, err
		}
		s.ultCheck()
	}

	s.skillEffect = model.SkillEffect_INVALID_SKILL_EFFECT
	return next, nil
}

func (sim *simulation) executeCharAction(target key.TargetID, isInsert bool) {
	// TODO: sim.eval.NextAction?
	// TODO: determine skillEffect & attackType from eval
	// TODO: SP increase/decrease depending on action chosen
	// TODO: Target Evaluator to determine target to use
	var skillEffect model.SkillEffect
	var attackType model.AttackType

	sim.skillEffect = skillEffect
	sim.event.ActionStart.Emit(event.ActionEvent{
		Target:      target,
		SkillEffect: skillEffect,
		AttackType:  attackType,
		IsInsert:    isInsert,
	})

	// TODO: execute action here

	// end attack if in one. no-op if not in an attack
	sim.combatManager.EndAttack()

	sim.event.ActionEnd.Emit(event.ActionEvent{
		Target:      target,
		SkillEffect: skillEffect,
		AttackType:  attackType,
		IsInsert:    isInsert,
	})
}

// TODO: may need to take in a burst type for MC case of having dual bursts?
func (sim *simulation) executeCharUlt(target key.TargetID) {
	// TODO: get ult execution function + this ult's skill effect
	var skillEffect model.SkillEffect

	sim.skillEffect = skillEffect
	sim.event.UltStart.Emit(event.ActionEvent{
		Target:      target,
		SkillEffect: skillEffect,
		AttackType:  model.AttackType_ULT,
		IsInsert:    true,
	})

	// TODO: execute ult here

	// end attack if in one. no-op if not in an attack
	sim.combatManager.EndAttack()

	sim.event.UltEnd.Emit(event.ActionEvent{
		Target:      target,
		SkillEffect: skillEffect,
		AttackType:  model.AttackType_ULT,
		IsInsert:    true,
	})
}

func (sim *simulation) executeInsert(i info.Insert) {
	sim.skillEffect = i.SkillEffect
	sim.event.InsertStart.Emit(event.InsertEvent{
		Target:      i.Source,
		SkillEffect: i.SkillEffect,
		AbortFlags:  i.AbortFlags,
		Priority:    i.Priority,
	})

	// execute insert
	i.Execute()

	// end attack if in one. no-op if not in an attack
	sim.combatManager.EndAttack()

	sim.event.InsertEnd.Emit(event.InsertEvent{
		Target:      i.Source,
		SkillEffect: i.SkillEffect,
		AbortFlags:  i.AbortFlags,
		Priority:    i.Priority,
	})
}
