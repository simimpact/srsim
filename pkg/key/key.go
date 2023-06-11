package key

type Enemy string

const (
	DummyEnemy Enemy = "dummy_enemy"
)

type ActionType int

const (
	InvalidAction ActionType = iota
	ActionAttack
	ActionSkill
	ActionBurst
	EndActionType
)

type TargetEvaluator string
