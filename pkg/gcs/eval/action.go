package eval

import (
	"errors"
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/action"
	"github.com/simimpact/srsim/pkg/key"
)

func (e *Eval) NextAction(target key.TargetID) (action.Action, error) {
	t, ok := e.targetNode[target]
	if !ok {
		return action.Action{}, errors.New("not found action callback")
	}
	act, err := e.evalTargetNode(t, key.ActionAttack, key.ActionSkill)
	if err != nil {
		return action.Action{}, err
	}
	if act.Type == key.InvalidAction {
		act, ok = e.defaultActions[target]
		if !ok {
			return action.Action{}, errors.New("not found default action")
		}
	}

	act.Target = target
	return act, nil
}

func (e *Eval) UltCheck() ([]action.Action, error) {
	result := make([]action.Action, 0)
	for _, t := range e.ultNodes {
		act, err := e.evalTargetNode(t, key.ActionUlt)
		if err != nil {
			return nil, err
		}
		if act.Type != key.InvalidAction {
			act.Target = t.target
			result = append(result, act)
		}
	}
	return result, nil
}

func (e *Eval) evalTargetNode(t TargetNode, checkType ...key.ActionType) (action.Action, error) {
	obj, err := e.evalNode(t.node, t.env)
	if err != nil {
		return action.Action{}, err
	}
	if obj.Typ() != typRet {
		return action.Action{}, errors.New("the function must return the value")
	}
	res := obj.(*retval).res
	if res.Typ() != typAct && res.Typ() != typNull {
		return action.Action{}, fmt.Errorf("the return value must be action or null, got %v", obj.Typ())
	}

	var act action.Action
	if res.Typ() == typAct {
		act = (res).(*actionval).val
	}

	// check required types
	if act.Type != key.InvalidAction && len(checkType) > 0 {
		found := false
		for _, v := range checkType {
			if act.Type == v {
				found = true
				break
			}
		}

		if !found {
			return action.Action{}, fmt.Errorf("wrong action type, got %v", obj.Typ())
		}
	}

	return act, nil
}
