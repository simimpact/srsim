package key

type TargetID int

const (
	TargetInvalid TargetID = 0
)

type CharacterKey string

type LightConeKey string

type RelicKey string

type EnemyKey string

type ActionType int

const (
	InvalidAction ActionType = iota
	ActionAttack
	ActionSkill
	ActionBurst
	EndActionType
)