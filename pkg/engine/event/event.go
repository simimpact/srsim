// package event provides an event system that all other functionalities can hook on to
// and either subscribe or emit events
package event

import "github.com/simimpact/srsim/pkg/engine/event/handler"

type System struct {
	AttackStart  AttackStartEventHandler
	AttackEnd    AttackEndEventHandler
	BeforeHit    BeforeHitEventHandler
	DamageResult DamageResultEventHandler
	AfterHit     AfterHitEventHandler
	BeforeHeal   BeforeHealEventHandler
	AfterHeal    AfterHealEventHandler

	CharacterAdded CharacterAddedEventHandler

	ModifierAdded            ModifierAddedEventHandler
	ModifierResisted         ModifierResistedEventHandler
	ModifierRemoved          ModifierRemovedEventHandler
	ModifierExtendedDuration ModifierExtendedDurationEventHandler
	ModifierExtendedCount    ModifierExtendedCountEventHandler

	HPChange      HPChangeEventHandler
	LimboWaitHeal LimboWaitHealEventHandler
	TargetDeath   TargetDeathEventHandler

	// test placeholder until we get actual events defined
	Ping handler.EventHandler[int]
}
