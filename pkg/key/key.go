package key

type Modifier string
type Shield string
type TargetEvaluator TargetID

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
