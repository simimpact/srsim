package logic

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
)

type Eval interface {
	Init(engine.Engine) error
	NextAction(key.TargetID) (Action, error)
	DefaultAction(key.TargetID) (Action, error)
	UltCheck() ([]Action, error)
}

type Action struct {
	Type            ActionType
	Target          key.TargetID        `exhaustruct:"optional"`
	TargetEvaluator key.TargetEvaluator `exhaustruct:"optional"`
}

type ActionType string

const (
	InvalidAction ActionType = ""
	ActionAttack  ActionType = "attack"
	ActionSkill   ActionType = "skill"
	ActionUlt     ActionType = "ult"
	// ActionUltAttack / ActionUltSkill is used to support case of MC
	ActionUltAttack ActionType = "ult_attack"
	ActionUltSkill  ActionType = "ult_skill"
	EndActionType   ActionType = "end"
)
