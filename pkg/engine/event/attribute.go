package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/key"
)

type HPChangeEventHandler = handler.EventHandler[HPChange]
type HPChange struct {
	Target             key.TargetID `json:"target"`
	OldHPRatio         float64      `json:"old_hp_ratio"`
	NewHPRatio         float64      `json:"new_hp_ratio"`
	OldHP              float64      `json:"old_hp"`
	NewHP              float64      `json:"new_hp"`
	IsHPChangeByDamage bool         `json:"is_hp_change_by_damage"`
}

type LimboWaitHealEventHandler = handler.CancelableEventHandler[LimboWaitHeal]
type LimboWaitHeal struct {
	Target      key.TargetID `json:"target"`
	IsCancelled bool         `exhaustruct:"optional" json:"is_cancelled"`
}

func (e LimboWaitHeal) Cancelled() handler.CancellableEvent {
	e.IsCancelled = true
	return e
}

type EnergyChangeEventHandler = handler.EventHandler[EnergyChange]
type EnergyChange struct {
	Target    key.TargetID `json:"target"`
	OldEnergy float64      `json:"old_energy"`
	NewEnergy float64      `json:"new_energy"`
}

type StanceChangeEventHandler = handler.EventHandler[StanceChange]
type StanceChange struct {
	Target    key.TargetID `json:"target"`
	OldStance float64      `json:"old_stance"`
	NewStance float64      `json:"new_stance"`
}

type StanceBreakEventHandler = handler.EventHandler[StanceBreak]
type StanceBreak struct {
	Target key.TargetID `json:"target"`
	Source key.TargetID `json:"source"`
}

type StanceResetEventHandler = handler.EventHandler[StanceReset]
type StanceReset struct {
	Target key.TargetID `json:"target"`
}

type SPChangeEventHandler = handler.EventHandler[SPChange]
type SPChange struct {
	OldSP int `json:"old_sp"`
	NewSP int `json:"new_sp"`
}
