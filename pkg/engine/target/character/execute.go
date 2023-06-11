package character

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target"
	"github.com/simimpact/srsim/pkg/engine/target/evaltarget"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// TODO: This function needs to execute the AST evaluator to determine what action it should perform.
// Unlike ExecuteUlt, we do not know what the action will be until we evaluate it right before
// its execution. Current logic in this function is a major placeholder instead of calling AST.
//
// At a high level the flow of this should be:
// 1. get the skill info data for this target
// 2. evaluate the ast to determine the action to execute (giving the skill info + current engine state)
// 3. AST returns what action/AttackType + a TargetEvaluator key
// 4. execute TargetEvaluator to determine primary target
// 5. build and return ExecutableAction
func (mgr *Manager) ExecuteAction(id key.TargetID, isInsert bool) (target.ExecutableAction, error) {
	skillInfo, err := mgr.SkillInfo(id)
	if err != nil {
		return target.ExecutableAction{}, err
	}
	char := mgr.instances[id]

	// TODO: this is hardcoded action behavior logic. This should be doing logic eval instead
	// of something hardcoded
	// TODO: eval.NextAction?
	// TODO: determine skillEffect & attackType from eval
	//
	// current hardcoded logic: use skill if possible, otherwise attack
	check := skillInfo.Skill.CanUse
	if mgr.engine.SP() > skillInfo.Skill.SPCost && (check == nil || check(mgr.engine, char)) {

		// TODO: this is placeholder evaltarget. Need to get what evaluator to use from AST
		primaryTarget, err := evaltarget.Evaluate(mgr.engine, evaltarget.Info{
			Source:      id,
			Evaluator:   evaltarget.LowestHP,
			TargetType:  skillInfo.Skill.ValidTargets,
			SourceClass: info.ClassCharacter,
		})
		if err != nil {
			return target.ExecutableAction{}, err
		}

		return target.ExecutableAction{
			Execute: func() {
				char.Skill(primaryTarget, actionState{
					mgr:         mgr,
					source:      id,
					isInsert:    isInsert,
					skillEffect: skillInfo.Skill.SkillEffect,
				})
			},
			SPChange:    -skillInfo.Skill.SPCost,
			SkillEffect: skillInfo.Skill.SkillEffect,
			AttackType:  model.AttackType_SKILL,
			IsInsert:    isInsert,
		}, nil
	}

	// TODO: this is placeholder evaltarget. Need to get what evaluator to use from AST
	primaryTarget, err := evaltarget.Evaluate(mgr.engine, evaltarget.Info{
		Source:      id,
		Evaluator:   evaltarget.LowestHP,
		TargetType:  skillInfo.Attack.ValidTargets,
		SourceClass: info.ClassCharacter,
	})
	if err != nil {
		return target.ExecutableAction{}, err
	}

	return target.ExecutableAction{
		Execute: func() {
			char.Attack(primaryTarget, actionState{
				mgr:         mgr,
				source:      id,
				isInsert:    isInsert,
				skillEffect: skillInfo.Attack.SkillEffect,
			})
		},
		SPChange:    +1,
		SkillEffect: skillInfo.Attack.SkillEffect,
		AttackType:  model.AttackType_NORMAL,
		IsInsert:    isInsert,
	}, nil
}

// TODO: This should take in the following
//   - id of the target
//   - TargetEvaluator key for what evaluation logic to use
//   - What UltType to use (to support case of MC)
//
// This should be a simple function that just:
//  1. gets the SkillEffect of the relevant UltType
//  2. find the method to execute in the character instance based on UltType
//  3. call TargetEvaluator to determine the primary target
//  4. return ExecutableUlt w/ this information bundled
func (mgr *Manager) ExecuteUlt(id key.TargetID) (target.ExecutableUlt, error) {
	skillInfo, err := mgr.SkillInfo(id)
	if err != nil {
		return target.ExecutableUlt{}, err
	}
	char := mgr.instances[id]

	// TODO: This is hardcoded ult behavior.
	if singleUlt, ok := char.(info.SingleUlt); ok {
		primaryTarget, err := evaltarget.Evaluate(mgr.engine, evaltarget.Info{
			Source:      id,
			Evaluator:   evaltarget.LowestHP,
			TargetType:  skillInfo.Attack.ValidTargets,
			SourceClass: info.ClassCharacter,
		})
		if err != nil {
			return target.ExecutableUlt{}, err
		}

		return target.ExecutableUlt{
			Execute: func() {
				singleUlt.Ult(primaryTarget, actionState{
					mgr:         mgr,
					source:      id,
					isInsert:    true,
					skillEffect: skillInfo.Ult.SkillEffect,
				})
			},
			SkillEffect: skillInfo.Ult.SkillEffect,
		}, nil
	} else if multiUlt, ok := char.(info.MultiUlt); ok {
		primaryTarget, err := evaltarget.Evaluate(mgr.engine, evaltarget.Info{
			Source:      id,
			Evaluator:   evaltarget.LowestHP,
			TargetType:  skillInfo.Attack.ValidTargets,
			SourceClass: info.ClassCharacter,
		})
		if err != nil {
			return target.ExecutableUlt{}, err
		}

		return target.ExecutableUlt{
			Execute: func() {
				multiUlt.UltAttack(primaryTarget, actionState{
					mgr:         mgr,
					source:      id,
					isInsert:    true,
					skillEffect: skillInfo.Ult.SkillEffect,
				})
			},
			SkillEffect: skillInfo.Ult.SkillEffect,
		}, nil
	}
	return target.ExecutableUlt{}, fmt.Errorf("unknown ult signature for char %v", id)
}

type actionState struct {
	mgr         *Manager
	source      key.TargetID
	isInsert    bool
	skillEffect model.SkillEffect
}

func (a actionState) IsInsert() bool {
	return a.isInsert
}

func (a actionState) SkillEffect() model.SkillEffect {
	return a.skillEffect
}

func (a actionState) EndAttack() {
	a.mgr.engine.EndAttack()
}

func (a actionState) CharacterInfo() info.Character {
	return a.mgr.info[a.source]
}
