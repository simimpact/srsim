package testeval

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic"
)

type ActionEval func(engine engine.Engine, id key.TargetID) logic.Action
type UltFunc func(engine engine.Engine) []logic.Action

func Default() ActionEval {
	return func(e engine.Engine, id key.TargetID) logic.Action {
		return logic.Action{
			Type:            logic.ActionAttack,
			Target:          id,
			TargetEvaluator: key.TargetEvaluator(e.Enemies()[0]),
		}
	}
}
