package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/key"
)

type TurnTargetsAddedEventHandler = handler.EventHandler[TurnTargetsAdded]
type TurnTargetsAdded struct {
	Targets   []key.TargetID
	TurnOrder []TurnStatus
}

type TurnResetEventHandler = handler.EventHandler[TurnReset]
type TurnReset struct {
	ResetTarget key.TargetID
	GaugeCost   float64
	TurnOrder   []TurnStatus
}

type GaugeChangeEventHandler = handler.EventHandler[GaugeChange]
type GaugeChange struct {
	Target    key.TargetID
	OldGauge  float64
	NewGauge  float64
	TurnOrder []TurnStatus
}

type CurrentGaugeCostChangeEventHandler = handler.EventHandler[CurrentGaugeCostChange]
type CurrentGaugeCostChange struct {
	OldCost float64
	NewCost float64
}

type TurnStatus struct {
	ID    key.TargetID
	Gauge float64
	AV    float64
}
