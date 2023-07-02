package logic

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
)

type Eval interface {
	Init(engine.Engine) error
	NextAction(key.TargetID) (Action, error)
	UltCheck() ([]Action, error)
}

type Action struct {
	Type            key.ActionType
	Target          key.TargetID        `exhaustruct:"optional"`
	TargetEvaluator key.TargetEvaluator `exhaustruct:"optional"`
}
