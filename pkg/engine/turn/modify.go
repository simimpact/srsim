package turn

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *manager) SetGauge(data info.ModifyAttribute) error {
	if _, ok := mgr.targetIndex[data.Target]; !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}

	prev := mgr.target(data.Target).gauge
	mgr.target(data.Target).gauge = data.Amount

	if mgr.target(data.Target).gauge == prev {
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

	mgr.event.GaugeChange.Emit(event.GaugeChange{
		Key:       data.Key,
		Target:    data.Target,
		Source:    data.Source,
		OldGauge:  prev,
		NewGauge:  mgr.target(data.Target).gauge,
		TurnOrder: status,
	})
	return nil
}

func (mgr *manager) ModifyGaugeNormalized(data info.ModifyAttribute) error {
	if _, ok := mgr.targetIndex[data.Target]; !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}

	data.Amount = mgr.target(data.Target).gauge + data.Amount*BaseGauge
	return mgr.SetGauge(data)
}

func (mgr *manager) ModifyGaugeAV(data info.ModifyAttribute) error {
	if _, ok := mgr.targetIndex[data.Target]; !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}

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
