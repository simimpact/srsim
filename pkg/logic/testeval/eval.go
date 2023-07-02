package testeval

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic"
)

var _ logic.Eval = (*Evaluator)(nil)

type TargetActionFunc func(engine engine.Engine, char info.CharInstance) logic.Action
type UltFunc func(engine engine.Engine) []logic.Action

type Evaluator struct {
	engine        engine.Engine
	targetActions map[key.TargetID]TargetActionFunc
	ult           UltFunc
}

func New() *Evaluator {
	return &Evaluator{
		engine:        nil,
		targetActions: make(map[key.TargetID]TargetActionFunc, 4),
		ult:           func(engine engine.Engine) []logic.Action { return nil },
	}
}

func (eval *Evaluator) UltLogic(f UltFunc) {
	eval.ult = f
}

func (eval *Evaluator) TargetAction(id key.TargetID, f TargetActionFunc) {
	eval.targetActions[id] = f
}

func (eval *Evaluator) Init(engine engine.Engine) error {
	eval.engine = engine
	return nil
}

func (eval *Evaluator) NextAction(id key.TargetID) (logic.Action, error) {
	if f, ok := eval.targetActions[id]; ok {
		char, _ := eval.engine.CharacterInstance(id)
		return f(eval.engine, char), nil
	}

	return logic.Action{Type: logic.InvalidAction}, fmt.Errorf("unknown target id: %v", id)
}

func (eval *Evaluator) UltCheck() ([]logic.Action, error) {
	return eval.ult(eval.engine), nil
}
