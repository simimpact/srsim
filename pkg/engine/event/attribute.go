package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/key"
)

type HPChangeEventHandler = handler.EventHandler[HPChange]
type HPChange struct {
	Target             key.TargetID
	OldHPRatio         float64
	NewHPRatio         float64
	OldHP              float64
	NewHP              float64
	IsHPChangeByDamage bool
}

type LimboWaitHealEventHandler = handler.CancelableEventHandler[LimboWaitHeal]
type LimboWaitHeal struct {
	Target      key.TargetID
	IsCancelled bool `exhaustruct:"optional"`
}

func (e LimboWaitHeal) Cancelled() handler.CancellableEvent {
	e.IsCancelled = true
	return e
}

type TargetDeathEventHandler = handler.EventHandler[TargetDeath]
type TargetDeath struct {
	Target key.TargetID
	Killer key.TargetID
}

type EnergyChangeEventHandler = handler.EventHandler[EnergyChange]
type EnergyChange struct {
	Target    key.TargetID
	OldEnergy float64
	NewEnergy float64
}

type StanceChangeEventHandler = handler.EventHandler[StanceChange]
type StanceChange struct {
	Target    key.TargetID
	OldStance float64
	NewStance float64
}

type StanceBreakEventHandler = handler.EventHandler[StanceBreak]
type StanceBreak struct {
	Target key.TargetID
	Source key.TargetID
}

type StanceResetEventHandler = handler.EventHandler[StanceReset]
type StanceReset struct {
	Target key.TargetID
}

type SPChangeEventHandler = handler.EventHandler[SPChange]
type SPChange struct {
	OldSP int
	NewSP int
}
