package turn

import (
	"sort"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
)

// SetGauge sets the gauge of the target as detailed in the input ModifyAttribute struct.
// 1. Update gauge of target
// 2. Move target to top of order
// 3. Sort target down based on AV. In the event of tie, this target should be at top of order.
//		If there is an active turn and it is not this target, this target should go below the active
//		turn (so index 1 instead of 0 when 0 gauge/AV)
// 4. Emit GaugeChangeEvent
func (mgr *manager) SetGauge(data info.ModifyAttribute) error {

	previousGauge := mgr.target(data.Target).gauge

	// if there's no change to Gauge, exit early
	if previousGauge == data.Amount {
		return nil
	}

	mgr.target(data.Target).gauge = data.Amount

	// find target index in mgr.orderHandler.turnOrder
	targetIndex, err := mgr.orderHandler.FindTargetIndex(data.Target)
	if err != nil {
		return err
	}

	// targetIndex == 0 indicates its already at the start of turnOrder, so no change needs to be made.
	// if there is an activeTurn, set our target to index 1; otherwise set to index 0.

	if targetIndex == 0 {
	} else {
		mgr.orderHandler.turnOrder = append([]*target{mgr.target(data.Target)}, append(mgr.orderHandler.turnOrder[:targetIndex], mgr.orderHandler.turnOrder[targetIndex+1:]...)...)
		if mgr.activeTarget != data.Target {
			switchValue := mgr.orderHandler.turnOrder[0]
			mgr.orderHandler.turnOrder[0] = mgr.orderHandler.turnOrder[1]
			mgr.orderHandler.turnOrder[1] = switchValue
		}
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
	data.Amount = mgr.target(data.Target).gauge + data.Amount*BaseGauge
	return mgr.SetGauge(data)
}

func (mgr *manager) ModifyGaugeAV(data info.ModifyAttribute) error {
	added := mgr.attr.Stats(data.Target).SPD() * data.Amount // SPD * AV = gauge
	data.Amount = mgr.target(data.Target).gauge + added

	return mgr.SetGauge(data)
}

func (mgr *manager) ModifyCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
	data.Amount = mgr.gaugeCost + data.Amount
	mgr.SetCurrentGaugeCost(data)
}

func (mgr *manager) SetCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
	prev := mgr.gaugeCost
	mgr.gaugeCost = data.Amount

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
