package info

type BattlePhase int

const (
	UnknownPhase BattlePhase = iota
	BattleStart
	TurnStart
	ModifierPhase1
	InsertAbilityPhase1
	ActionStart
	ActionEnd
	InsertAbilityPhase2
	ModifierPhase2
	TurnEnd
)
