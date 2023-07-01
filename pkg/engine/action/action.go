package action

import "github.com/simimpact/srsim/pkg/key"

type Action struct {
	Type            key.ActionType
	Target          key.TargetID        `exhaustruct:"optional"`
	TargetEvaluator key.TargetEvaluator `exhaustruct:"optional"`
}
