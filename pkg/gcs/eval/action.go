package eval

import (
	"errors"
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/action"
	"github.com/simimpact/srsim/pkg/key"
)

func (e *Eval) NextAction(target key.TargetID) action.Action {
	t, ok := e.targetNode[target]
	if !ok {
		return action.Action{Target: target, Type: key.ActionAttack}
	}
	useSkill, err := e.evalTargetNode(t)
	if err != nil {
		e.Err <- err
		return action.Action{Target: target, Type: key.ActionAttack}
	}

	actionType := key.ActionAttack
	if useSkill {
		actionType = key.ActionSkill
	}
	return action.Action{Target: target, Type: actionType}
}

func (e *Eval) BurstCheck() []action.Action {
	result := make([]action.Action, 0)
	for _, t := range e.burstNodes {
		useBurst, err := e.evalTargetNode(t)
		if err != nil {
			e.Err <- err
			break
		}
		if useBurst {
			result = append(result, action.Action{
				Target: t.target,
				Type:   key.ActionBurst,
			})
		}
	}
	return result
}

func (e *Eval) evalTargetNode(t TargetNode) (bool, error) {
	obj, err := e.evalNode(t.node, t.env)
	if err != nil {
		return false, err
	}
	if obj.Typ() != typRet {
		return false, errors.New("the function must return the value")
	}
	res := obj.(*retval).res
	if res.Typ() != typNum {
		return false, fmt.Errorf("the return value must be number, got %v", obj.Typ())
	}
	return ntob(res.(*number)), nil
}
