// package turn provides a naiva implementation of the TurnManager
package turn

import (
	"fmt"
	"sort"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

const BaseGauge float64 = 10000.0

type Manager interface {
	TotalAV() float64
	AddTargets(ids ...key.TargetID)
	RemoveTarget(id key.TargetID)
	StartTurn() (key.TargetID, float64, []event.TurnStatus, error)
	ResetTurn() error
	engine.Turn
}

type target struct {
	id    key.TargetID
	gauge float64
}

type turnOrder struct {
	attr  attribute.Getter
	order []*target
}

func (a *turnOrder) av(i int) float64 {
	return a.order[i].gauge / a.attr.Stats(a.order[i].id).SPD()
}

func (a *turnOrder) Len() int           { return len(a.order) }
func (a *turnOrder) Swap(i, j int)      { a.order[i], a.order[j] = a.order[j], a.order[i] }
func (a *turnOrder) Less(i, j int) bool { return a.av(i) < a.av(j) }

type manager struct {
	event *event.System
	attr  attribute.Getter

	// TODO: I'd create a custom data type/struct that contains both order & targetIndex. It then
	// manages all operations on the order array and keeps the index map up to date for easy access
	order        *turnOrder
	gaugeCost    float64
	activeTurn   bool
	activeTarget key.TargetID
	totalAV      float64
}

func New(e *event.System, attr attribute.Getter) Manager {
	mgr := &manager{
		event: e,
		attr:  attr,
		order: &turnOrder{
			attr:  attr,
			order: make([]*target, 0, 10),
		},
		gaugeCost:    1.0,
		activeTurn:   false,
		activeTarget: 0,
		totalAV:      0,
	}
	return mgr
}

func (mgr *manager) TotalAV() float64 {
	return mgr.totalAV
}

func (mgr *manager) target(id key.TargetID) *target {
	for _, t := range mgr.order.order {
		if t.id == id {
			return t
		}
	}
	return nil
}

// returns the current AV of the given target based on their current gauge and speed.
// This call is "expensive", so avoid calling it multiple times in the same logic.
// TODO: might want to change this into a util function that also takes in the speed?
func (mgr *manager) av(id key.TargetID) float64 {
	return mgr.target(id).gauge / mgr.attr.Stats(id).SPD()
}

// This is a variadic for performance (cheaper to add multiple at once rather than one at a time)
func (mgr *manager) AddTargets(ids ...key.TargetID) {
	for _, id := range ids {
		t := &target{
			id:    id,
			gauge: BaseGauge,
		}
		mgr.order.order = append(mgr.order.order, t)
	}

	// TODO: sort the order array based on each target's AV. This sort algorithm must be stable.
	//		update targetIndexes based off the new positions post sort.
	sort.Stable(mgr.order)

	mgr.event.TurnTargetsAdded.Emit(event.TurnTargetsAdded{
		Targets:   ids,
		TurnOrder: []event.TurnStatus{}, // TODO: populate
	})
}

func (mgr *manager) RemoveTarget(id key.TargetID) {
	idx := 0
	for i, t := range mgr.order.order {
		if t.id == id {
			idx = i
			break
		}
	}
	mgr.order.order = append(mgr.order.order[:idx], mgr.order.order[idx+1:]...)
}

func (mgr *manager) StartTurn() (key.TargetID, float64, []event.TurnStatus, error) {
	if mgr.activeTurn {
		return -1, 0, nil, fmt.Errorf("cannot start turn when already in an active turn: %+v", mgr)
	}

	// reset gauge cost for this new turn
	mgr.gaugeCost = 1.0
	mgr.activeTurn = true
	mgr.activeTarget = mgr.order.order[0].id
	av := mgr.av(mgr.activeTarget)

	// mgr.order.order[0].gauge = 0 // set gauge/av of active to 0

	for _, t := range mgr.order.order {
		t.gauge -= av * mgr.attr.Stats(t.id).SPD()
	}

	// TODO: Set the active turn to the target at the top of the order (target with lowest AV).
	//		Get the current AV of this target and subtract it from all targets in the order to simulate
	//		"progressing time forward" by the given AV.
	//
	//	1. Mark target at top of order as "active"
	//	2. get that target's current AV
	//	3. add AV to "TotalAV" (keeps track of how much AV has progressed over entire sim)
	//	4. loop through all targets in the order, subtracing this AV from them (active target gauge = 0)
	//			new_gauge = current_gauge - av * speed

	mgr.totalAV += av
	return mgr.activeTarget, av, nil, nil
}

func (mgr *manager) ResetTurn() error {
	if !mgr.activeTurn {
		return fmt.Errorf(
			"target at top of order must have 0 gauge to call reset (their turn is active) %+v", mgr.order.order[0])
	}
	mgr.activeTurn = false
	mgr.order.order[0].gauge = BaseGauge * mgr.gaugeCost
	sort.Stable(mgr.order)

	// Resets the gauge of the target taking their turn (target at top of stack). New gauge is set at
	// BaseGauge * gaugeCost
	//
	// target should then be moved to the bottom of the turn order and then sorted up to the correct
	// position based on their AV. In the event of ties, this target should be last in the order
	// Note that this is the same behavior as AddTarget (but different from SetGauge which goes top-down)

	// TODO:
	// 1. update gauge of active target
	// 2. move target to bottom of order
	// 3. sort target up based on AV. In the event of tie, this target should be at bottom of order.
	// 4. update targetIndexes for all targets that moved in the order (or just repopulate all)
	// 5. emit TurnResetEvent

	mgr.event.TurnReset.Emit(event.TurnReset{
		ResetTarget: mgr.activeTarget,
		GaugeCost:   mgr.gaugeCost,
		TurnOrder:   []event.TurnStatus{}, // TODO: need to populate based on new order
	})
	return nil
}
