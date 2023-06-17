// package event provides an event system that all other functionalities can hook on to
// and either subscribe or emit events
package event

import "github.com/simimpact/srsim/pkg/engine/event/handler"

type System struct {
	Initialize  InitializeEventHandler
	BattleStart BattleStartEventHandler
	TurnEnd     TurnEndEventHandler
	Termination TerminationEventHandler
	ActionStart ActionStartEventHandler
	ActionEnd   ActionEndEventHandler
	UltStart    UltStartEventHandler
	UltEnd      UltEndEventHandler
	InsertStart InsertStartEventHandler
	InsertEnd   InsertEndEventHandler

	AttackStart AttackStartEventHandler
	AttackEnd   AttackEndEventHandler
	HitStart    HitStartEventHandler
	HitEnd      HitEndEventHandler
	HealStart   HealStartEventHandler
	HealEnd     HealEndEventHandler

	CharacterAdded CharacterAddedEventHandler
	EnemyAdded     EnemyAddedEventHandler

	ModifierAdded            ModifierAddedEventHandler
	ModifierResisted         ModifierResistedEventHandler
	ModifierRemoved          ModifierRemovedEventHandler
	ModifierExtendedDuration ModifierExtendedDurationEventHandler
	ModifierExtendedCount    ModifierExtendedCountEventHandler

	ShieldAdded   ShieldAddedEventHandler
	ShieldRemoved ShieldRemovedEventHandler
	ShieldChange  ShieldChangeEventHandler

	HPChange       HPChangeEventHandler
	LimboWaitHeal  LimboWaitHealEventHandler
	TargetDeath    TargetDeathEventHandler
	EnergyChange   EnergyChangeEventHandler
	StanceChange   StanceChangeEventHandler
	StanceBreak    StanceBreakEventHandler
	StanceBreakEnd StanceBreakEndEventHandler
	SPChange       SPChangeEventHandler

	TurnTargetsAdded       TurnTargetsAddedEventHandler
	TurnStart              TurnStartEventHandler
	TurnReset              TurnResetEventHandler
	GaugeChange            GaugeChangeEventHandler
	CurrentGaugeCostChange CurrentGaugeCostChangeEventHandler

	// test placeholder until we get actual events defined
	Ping handler.EventHandler[int]
}
