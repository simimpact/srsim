package character

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target"
	"github.com/simimpact/srsim/pkg/engine/target/evaltarget"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic"
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

	act, err := mgr.eval.NextAction(id)
	if err != nil {
		return target.ExecutableAction{}, err
	}

	check := skillInfo.Skill.CanUse
	useSkill := act.Type == logic.ActionSkill
	if useSkill && mgr.engine.SP() >= skillInfo.Skill.SPNeed && (check == nil || check(mgr.engine, char)) {
		primaryTarget, err := evaltarget.Evaluate(mgr.engine, evaltarget.Info{
			Source:      id,
			Evaluator:   act.TargetEvaluator,
			TargetType:  skillInfo.Skill.TargetType,
			SourceClass: info.ClassCharacter,
		})
		if err != nil {
			return target.ExecutableAction{}, err
		}

		return target.ExecutableAction{
			Execute: func() {
				char.Skill(primaryTarget, actionState{
					engine:   mgr.engine,
					isInsert: isInsert,
				})
			},
			SPDelta:    -skillInfo.Skill.SPNeed,
			AttackType: model.AttackType_SKILL,
			IsInsert:   isInsert,
			Key:        mgr.info[id].Key.String(),
		}, nil
	}

	primaryTarget, err := evaltarget.Evaluate(mgr.engine, evaltarget.Info{
		Source:      id,
		Evaluator:   act.TargetEvaluator,
		TargetType:  skillInfo.Attack.TargetType,
		SourceClass: info.ClassCharacter,
	})
	if err != nil {
		return target.ExecutableAction{}, err
	}

	return target.ExecutableAction{
		Execute: func() {
			char.Attack(primaryTarget, actionState{
				engine:   mgr.engine,
				isInsert: isInsert,
			})
		},
		SPDelta:    skillInfo.Attack.SPAdd,
		AttackType: model.AttackType_NORMAL,
		IsInsert:   isInsert,
		Key:        mgr.info[id].Key.String(),
	}, nil
}

// TODO: This should take in the following
//   - id of the target
//   - TargetEvaluator key for what evaluation logic to use
//   - What UltType to use (to support case of MC)
//
// This should be a simple function that just:
//  1. find the method to execute in the character instance based on UltType
//  2. call TargetEvaluator to determine the primary target
//  3. return ExecutableUlt w/ this information bundled
func (mgr *Manager) ExecuteUlt(act logic.Action) (target.ExecutableUlt, error) {
	id := act.Target
	skillInfo, err := mgr.SkillInfo(id)
	if err != nil {
		return target.ExecutableUlt{}, err
	}
	char := mgr.instances[id]

	if singleUlt, ok := char.(info.SingleUlt); ok {
		if act.Type != logic.ActionUlt { // if key.ActionUltAttack or key.ActionUltSkill is used
			return target.ExecutableUlt{}, fmt.Errorf("wrong action key; expected ult, got %s", string(act.Type))
		}

		primaryTarget, err := evaltarget.Evaluate(mgr.engine, evaltarget.Info{
			Source:      id,
			Evaluator:   act.TargetEvaluator,
			TargetType:  skillInfo.Ult.TargetType,
			SourceClass: info.ClassCharacter,
		})
		if err != nil {
			return target.ExecutableUlt{}, err
		}

		return target.ExecutableUlt{
			Execute: func() {
				singleUlt.Ult(primaryTarget, actionState{
					engine:   mgr.engine,
					isInsert: false,
				})
			},
		}, nil
	}

	if multiUlt, ok := char.(info.MultiUlt); ok {
		if act.Type != logic.ActionUltAttack && act.Type != logic.ActionUltSkill { // if key.ActionUlt is used
			return target.ExecutableUlt{}, fmt.Errorf(
				"wrong action key; expected ult_attack or ult_skill, got %s", string(act.Type))
		}

		primaryTarget, err := evaltarget.Evaluate(mgr.engine, evaltarget.Info{
			Source:      id,
			Evaluator:   evaltarget.LowestHP,
			TargetType:  skillInfo.Ult.TargetType,
			SourceClass: info.ClassCharacter,
		})
		if err != nil {
			return target.ExecutableUlt{}, err
		}

		return target.ExecutableUlt{
			Execute: func() {
				multiUlt.UltAttack(primaryTarget, actionState{
					engine:   mgr.engine,
					isInsert: false,
				})
			},
		}, nil
	}
	return target.ExecutableUlt{}, fmt.Errorf("unknown ult signature for char %v", id)
}

type actionState struct {
	engine   engine.Engine
	isInsert bool
}

func (a actionState) IsInsert() bool {
	return a.isInsert
}
func (a actionState) EndAttack() {
	a.engine.EndAttack()
}
