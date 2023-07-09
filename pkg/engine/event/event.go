// package event provides an event system that all other functionalities can hook on to
// and either subscribe or emit events
package event

import "github.com/simimpact/srsim/pkg/engine/event/handler"

type System struct {
	Initialize  InitializeEventHandler
	BattleStart BattleStartEventHandler
	Termination TerminationEventHandler
	Phase1Start Phase1StartEventHandler
	Phase1End   Phase1EndEventHandler
	Phase2Start Phase2StartEventHandler
	Phase2End   Phase2EndEventHandler
	TurnStart   TurnStartEventHandler
	TurnEnd     TurnEndEventHandler
	ActionStart ActionStartEventHandler
	ActionEnd   ActionEndEventHandler
	InsertStart InsertStartEventHandler
	InsertEnd   InsertEndEventHandler
	TargetDeath TargetDeathEventHandler

	AttackStart AttackStartEventHandler
	AttackEnd   AttackEndEventHandler
	HitStart    HitStartEventHandler
	HitEnd      HitEndEventHandler
	HealStart   HealStartEventHandler
	HealEnd     HealEndEventHandler

	CharactersAdded CharactersAddedEventHandler
	EnemiesAdded    EnemiesAddedEventHandler

	ModifierAdded            ModifierAddedEventHandler
	ModifierResisted         ModifierResistedEventHandler
	ModifierRemoved          ModifierRemovedEventHandler
	ModifierExtendedDuration ModifierExtendedDurationEventHandler
	ModifierExtendedCount    ModifierExtendedCountEventHandler

	ShieldAdded   ShieldAddedEventHandler
	ShieldRemoved ShieldRemovedEventHandler
	ShieldChange  ShieldChangeEventHandler

	HPChange      HPChangeEventHandler
	LimboWaitHeal LimboWaitHealEventHandler
	EnergyChange  EnergyChangeEventHandler
	StanceChange  StanceChangeEventHandler
	StanceBreak   StanceBreakEventHandler
	StanceReset   StanceResetEventHandler
	SPChange      SPChangeEventHandler

	TurnTargetsAdded       TurnTargetsAddedEventHandler
	TurnReset              TurnResetEventHandler
	GaugeChange            GaugeChangeEventHandler
	CurrentGaugeCostChange CurrentGaugeCostChangeEventHandler

	// test placeholder until we get actual events defined
	Ping handler.EventHandler[int]
}
