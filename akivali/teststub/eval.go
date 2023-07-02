package teststub

import (
	"fmt"

	"github.com/simimpact/srsim/akivali/testcfg/testeval"
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic"
)

var _ logic.Eval = (*evaluator)(nil)

type evaluator struct {
	engine     engine.Engine
	charAction map[key.TargetID]testeval.ActionEval
	ult        testeval.UltFunc
}

func newEvaluator() *evaluator {
	return &evaluator{
		engine:     nil,
		charAction: make(map[key.TargetID]testeval.ActionEval, 4),
		ult:        func(engine engine.Engine) []logic.Action { return nil },
	}
}

func (e *evaluator) registerAction(id key.TargetID, f testeval.ActionEval) {
	e.charAction[id] = f
}

func (e *evaluator) Init(engine engine.Engine) error {
	e.engine = engine
	return nil
}

func (e *evaluator) NextAction(id key.TargetID) (logic.Action, error) {
	if f, ok := e.charAction[id]; ok {
		return f(e.engine, id), nil
	}

	return logic.Action{Type: logic.InvalidAction}, fmt.Errorf("unknown target id: %v", id)
}

func (e *evaluator) UltCheck() ([]logic.Action, error) {
	return e.ult(e.engine), nil
}
