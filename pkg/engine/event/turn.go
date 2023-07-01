package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/key"
)

type TurnTargetsAddedEventHandler = handler.EventHandler[TurnTargetsAddedEvent]
type TurnTargetsAddedEvent struct {
	Targets   []key.TargetID
	TurnOrder []TurnStatus
}

type TurnResetEventHandler = handler.EventHandler[TurnResetEvent]
type TurnResetEvent struct {
	ResetTarget key.TargetID
	GaugeCost   float64
	TurnOrder   []TurnStatus
}

type GaugeChangeEventHandler = handler.EventHandler[GaugeChangeEvent]
type GaugeChangeEvent struct {
	Target    key.TargetID
	OldGauge  float64
	NewGauge  float64
	TurnOrder []TurnStatus
}

type CurrentGaugeCostChangeEventHandler = handler.EventHandler[CurrentGaugeCostChangeEvent]
type CurrentGaugeCostChangeEvent struct {
	OldCost float64
	NewCost float64
}

type TurnStatus struct {
	ID    key.TargetID
	Gauge float64
	AV    float64
}
