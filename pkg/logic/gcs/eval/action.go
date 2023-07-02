package eval

import (
	"errors"
	"fmt"

	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic"
)

func (e *Eval) NextAction(target key.TargetID) (logic.Action, error) {
	t, ok := e.targetNode[target]
	if !ok {
		return logic.Action{}, errors.New("not found action callback")
	}
	act, err := e.evalTargetNode(t, logic.ActionAttack, logic.ActionSkill)
	if err != nil {
		return logic.Action{}, err
	}
	if act.Type == logic.InvalidAction {
		act, ok = e.defaultActions[target]
		if !ok {
			return logic.Action{}, errors.New("not found default action")
		}
	}

	act.Target = target
	return act, nil
}

func (e *Eval) UltCheck() ([]logic.Action, error) {
	result := make([]logic.Action, 0)
	for _, t := range e.ultNodes {
		act, err := e.evalTargetNode(t, logic.ActionUlt)
		if err != nil {
			return nil, err
		}
		if act.Type != logic.InvalidAction {
			act.Target = t.target
			result = append(result, act)
		}
	}
	return result, nil
}

func (e *Eval) evalTargetNode(t TargetNode, checkType ...logic.ActionType) (logic.Action, error) {
	obj, err := e.evalNode(t.node, t.env)
	if err != nil {
		return logic.Action{}, err
	}
	if obj.Typ() != typRet {
		return logic.Action{}, errors.New("the function must return the value")
	}
	res := obj.(*retval).res
	if res.Typ() != typAct && res.Typ() != typNull {
		return logic.Action{}, fmt.Errorf("the return value must be action or null, got %v", obj.Typ())
	}

	var act logic.Action
	if res.Typ() == typAct {
		act = (res).(*actionval).val
	}

	// check required types
	if act.Type != logic.InvalidAction && len(checkType) > 0 {
		found := false
		for _, v := range checkType {
			if act.Type == v {
				found = true
				break
			}
		}

		if !found {
			return logic.Action{}, fmt.Errorf("wrong action type, got %v", obj.Typ())
		}
	}

	return act, nil
}
