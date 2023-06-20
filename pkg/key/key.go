package key

type Modifier string
type Shield string
type TargetEvaluator TargetID

type ActionType string

const (
	InvalidAction ActionType = "invalid"
	ActionAttack  ActionType = "attack"
	ActionSkill   ActionType = "skill"
	ActionBurst   ActionType = "burst"
	EndActionType ActionType = "end"
)
