package key

type Modifier string
type Shield string
type TargetEvaluator string

type ActionType int

const (
	InvalidAction ActionType = iota
	ActionAttack
	ActionSkill
	ActionBurst
	EndActionType
)
