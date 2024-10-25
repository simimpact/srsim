package turn

import (
	"sort"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
)

// SetGauge sets the gauge of the target as detailed in the input ModifyAttribute struct.
//  1. Update gauge of target
//  2. Move target to top of order
//  3. Sort target down based on AV. In the event of tie, this target should be at top of order.
//     If there is an active turn and it is not this target, this target should go below the active
//     turn (so index 1 instead of 0 when 0 gauge/AV)
//  4. Emit GaugeChangeEvent
func (mgr *manager) SetGauge(data info.ModifyAttribute) error {
	previousGauge := mgr.target(data.Target).gauge

	// if there's no change to Gauge, exit early
	if previousGauge == int64(data.Amount) {
		return nil
	}

	mgr.target(data.Target).gauge = int64(data.Amount)

	// find target index in mgr.orderHandler.turnOrder
	targetIndex, err := mgr.orderHandler.FindTargetIndex(data.Target)
	if err != nil {
		return err
	}

	// set start index to 1 only if there is an active turn and it is not this target. Do not want to
	// make this target the active target if not their turn.

	startIndex := 0
	if mgr.activeTurn && targetIndex != 0 {
		startIndex = 1
	}

	prev := mgr.orderHandler.turnOrder[targetIndex]
	for i := startIndex; i <= targetIndex; i++ {
		mgr.orderHandler.turnOrder[i], prev = prev, mgr.orderHandler.turnOrder[i]
	}

	sort.Stable(mgr.orderHandler)

	mgr.event.GaugeChange.Emit(event.GaugeChange{
		Key:       data.Key,
		Target:    data.Target,
		Source:    data.Source,
		OldGauge:  previousGauge,
		NewGauge:  mgr.target(data.Target).gauge,
		TurnOrder: mgr.EventTurnStatus(),
	})
	return nil
}

func (mgr *manager) ModifyGaugeNormalized(data info.ModifyAttribute) error {
	data.Amount = float64(mgr.target(data.Target).gauge) + data.Amount*float64(BaseGauge)
	return mgr.SetGauge(data)
}

func (mgr *manager) ModifyGaugeAV(data info.ModifyAttribute) error {
	added := mgr.attr.Stats(data.Target).SPD() * data.Amount // SPD * AV = gauge
	data.Amount = float64(mgr.target(data.Target).gauge) + added

	return mgr.SetGauge(data)
}

func (mgr *manager) ModifyCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
	data.Amount = float64(mgr.gaugeCost) + data.Amount
	mgr.SetCurrentGaugeCost(data)
}

func (mgr *manager) SetCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
	prev := mgr.gaugeCost
	mgr.gaugeCost = int64(data.Amount)

	if prev == mgr.gaugeCost {
		return
	}

	mgr.event.CurrentGaugeCostChange.Emit(event.CurrentGaugeCostChange{
		Key:     data.Key,
		Source:  data.Source,
		OldCost: prev,
		NewCost: mgr.gaugeCost,
	})
}
