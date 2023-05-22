package key


type EnemyKey string

const (
	DummyEnemy EnemyKey = "dummy_enemy"
)

type ActionType int

const (
	InvalidAction ActionType = iota
	ActionAttack
	ActionSkill
	ActionBurst
	EndActionType
)