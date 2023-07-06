package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/key"
)

type TurnTargetsAddedEventHandler = handler.EventHandler[TurnTargetsAdded]
type TurnTargetsAdded struct {
	Targets   []key.TargetID `json:"targets"`
	TurnOrder []TurnStatus   `json:"turn_order"`
}

type TurnResetEventHandler = handler.EventHandler[TurnReset]
type TurnReset struct {
	ResetTarget key.TargetID `json:"reset_target"`
	GaugeCost   float64      `json:"gauge_cost"`
	TurnOrder   []TurnStatus `json:"turn_order"`
}

type GaugeChangeEventHandler = handler.EventHandler[GaugeChange]
type GaugeChange struct {
	Target    key.TargetID `json:"target"`
	OldGauge  float64      `json:"old_gauge"`
	NewGauge  float64      `json:"new_gauge"`
	TurnOrder []TurnStatus `json:"turn_order"`
}

type CurrentGaugeCostChangeEventHandler = handler.EventHandler[CurrentGaugeCostChange]
type CurrentGaugeCostChange struct {
	OldCost float64 `json:"old_cost"`
	NewCost float64 `json:"new_cost"`
}

type TurnStatus struct {
	ID    key.TargetID `json:"id"`
	Gauge float64      `json:"gauge"`
	AV    float64      `json:"av"`
}
