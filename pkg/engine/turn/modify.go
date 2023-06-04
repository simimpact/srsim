package turn

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) SetGauge(target key.TargetID, amt float64) error {
	if _, ok := mgr.targetIndex[target]; !ok {
		return fmt.Errorf("unknown target: %v", target)
	}

	prev := mgr.target(target).gauge
	mgr.target(target).gauge = amt

	if mgr.target(target).gauge == prev {
		return nil
	}

	// create map of TargetID -> AV so you only need to calc once within this call
	targetAV := make(map[key.TargetID]float64, len(mgr.order))
	for k := range mgr.targetIndex {
		targetAV[k] = mgr.av(k)
	}

	// TODO:
	// 1. update gauge of target
	// 2. move target to top of order
	// 3. sort target down based on AV. In the event of tie, this target should be at top of order.
	//		If there is an active turn and it is not this target, this target should go below the active
	//		turn (so index 1 instead of 0 when 0 gauge/AV)
	// 4. emit GaugeChangeEvent

	// TODO: this is also needed for TurnStart emit & TurnReset emit, should be abstracted
	status := make([]event.TurnStatus, len(mgr.order))
	for i, t := range mgr.order {
		status[i] = event.TurnStatus{
			ID:    t.id,
			Gauge: t.gauge,
			AV:    targetAV[t.id],
			// TODO: should we also add speed to this?
		}
	}

	mgr.event.GaugeChange.Emit(event.GaugeChangeEvent{
		Target:    target,
		OldGauge:  prev,
		NewGauge:  mgr.target(target).gauge,
		TurnOrder: status,
	})
	return nil
}

func (mgr *Manager) ModifyGaugeNormalized(target key.TargetID, amt float64) error {
	if _, ok := mgr.targetIndex[target]; !ok {
		return fmt.Errorf("unknown target: %v", target)
	}

	return mgr.SetGauge(target, mgr.target(target).gauge+amt*BASE_GAUGE)
}

func (mgr *Manager) ModifyGaugeAV(target key.TargetID, amt float64) error {
	if _, ok := mgr.targetIndex[target]; !ok {
		return fmt.Errorf("unknown target: %v", target)
	}

	added := mgr.attr.Stats(target).SPD() * amt // SPD * AV = gauge
	return mgr.SetGauge(target, mgr.target(target).gauge+added)
}

func (mgr *Manager) ModifyCurrentGaugeCost(amt float64) {
	mgr.SetCurrentGaugeCost(mgr.gaugeCost + amt)
}

func (mgr *Manager) SetCurrentGaugeCost(amt float64) {
	prev := mgr.gaugeCost
	mgr.gaugeCost = amt

	if prev == mgr.gaugeCost {
		return
	}

	mgr.event.CurrentGaugeCostChange.Emit(event.CurrentGaugeCostChangeEvent{
		OldCost: prev,
		NewCost: mgr.gaugeCost,
	})
}
