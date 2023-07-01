package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/key"
)

type HPChangeEventHandler = handler.EventHandler[HPChangeEvent]
type HPChangeEvent struct {
	Target             key.TargetID
	OldHPRatio         float64
	NewHPRatio         float64
	OldHP              float64
	NewHP              float64
	IsHPChangeByDamage bool
}

type LimboWaitHealEventHandler = handler.CancelableEventHandler[LimboWaitHealEvent]
type LimboWaitHealEvent struct {
	Target      key.TargetID
	IsCancelled bool `exhaustruct:"optional"`
}

func (e LimboWaitHealEvent) Cancelled() handler.CancellableEvent {
	e.IsCancelled = true
	return e
}

type TargetDeathEventHandler = handler.EventHandler[TargetDeathEvent]
type TargetDeathEvent struct {
	Target key.TargetID
	Killer key.TargetID
}

type EnergyChangeEventHandler = handler.EventHandler[EnergyChangeEvent]
type EnergyChangeEvent struct {
	Target    key.TargetID
	OldEnergy float64
	NewEnergy float64
}

type StanceChangeEventHandler = handler.EventHandler[StanceChangeEvent]
type StanceChangeEvent struct {
	Target    key.TargetID
	OldStance float64
	NewStance float64
}

type StanceBreakEventHandler = handler.EventHandler[StanceBreakEvent]
type StanceBreakEvent struct {
	Target key.TargetID
	Source key.TargetID
}

type StanceResetEventHandler = handler.EventHandler[StanceResetEvent]
type StanceResetEvent struct {
	Target key.TargetID
}

type SPChangeEventHandler = handler.EventHandler[SPChangeEvent]
type SPChangeEvent struct {
	OldSP int
	NewSP int
}
