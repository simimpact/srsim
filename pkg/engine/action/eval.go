package action

import (
	"github.com/simimpact/srsim/pkg/gcs/ast"
	"github.com/simimpact/srsim/pkg/key"
)

type BurstNode struct {
	target key.TargetID
	node ast.Node
}

type ActionEvaluater interface {
	NextAction(target key.TargetID) Action
	BurstCheck() (Action, bool)
}

type EvalAction struct {
	targetNode map[key.TargetID]ast.Node
	burstNodes []BurstNode
}

func (e *EvalAction) NextAction(target key.TargetID) Action {
	t, ok := e.targetNode[target]
	if !ok {
		//most likely if nothing is registered we'll want to just return
		//attack here
		panic("error handle here")
	}
	return e.eval(t)
}

func (e *EvalAction) BurstCheck() (Action, bool) {
	for _, v := range e.burstNodes {
		res := e.eval(v.node)
		if res.Type != key.InvalidAction {
			return res, true
		}
	}
	return Action{Type: key.InvalidAction}, false
}

func (e *EvalAction) eval(n ast.Node) Action {
	//your node eval code goes here
	//note that this node should be an expr that evaluates to an action
	//so we should have some sort of special type ActionStmt defined
	//in the ast
}
