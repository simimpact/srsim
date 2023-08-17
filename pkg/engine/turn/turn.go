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
	TurnOrder() []key.TargetID
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

// TotalAV returns the total AV tracked by the Manager.
func (mgr *manager) TotalAV() float64 {
	return mgr.totalAV
}

// TurnOrder returns an ordered array of TargetIDs from the Manager's turnOrder array.
func (mgr *manager) TurnOrder() []key.TargetID {
	targetOrder := make([]key.TargetID, mgr.order.Len())
	for i, v := range mgr.order.order {
		targetOrder[i] = v.id
	}
	return targetOrder
}

// EventTurnStatus returns an array of event.TurnStatus structs populated with the current ID, Gauge, and AV of each target in the Manager's turnOrder.
func (mgr *manager) EventTurnStatus() []event.TurnStatus {
	turnStatus := make([]event.TurnStatus, mgr.order.Len())
	for i, v := range mgr.order.order {
		turnStatus[i] = event.TurnStatus {
			ID : v.id,
			Gauge : v.gauge,
			AV : mgr.order.av(i),
		}
	}
	return turnStatus
}

// target checks whether a key.TargetID exists in the Manager's turnOrder array.
// If it exists, it returns a pointer to the target. Otherwise, it returns nil.
func (mgr *manager) target(id key.TargetID) *target {
	for _, t := range mgr.order.order {
		if t.id == id {
			return t
		}
	}
	return nil
}

// av returns the current AV of the given target based on their current gauge and speed.
// This call is "expensive", so avoid calling it multiple times in the same logic.
// TODO: might want to change this into a util function that also takes in the speed?
func (mgr *manager) av(id key.TargetID) float64 {
	return mgr.target(id).gauge / mgr.attr.Stats(id).SPD()
}

// AddTargets adds a target to the Manager's turnOrder.
// This is a variadic for performance (cheaper to add multiple at once rather than one at a time)
func (mgr *manager) AddTargets(ids ...key.TargetID) {
	for _, id := range ids {
		t := &target{
			id:    id,
			gauge: BaseGauge,
		}
		mgr.order.order = append(mgr.order.order, t)
	}

	sort.Stable(mgr.order)

	mgr.event.TurnTargetsAdded.Emit(event.TurnTargetsAdded{
		Targets:   ids,
		TurnOrder: mgr.EventTurnStatus(),
	})
}

// RemoveTarget removes a target from the Manager's turnOrder.
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

// StartTurn processes changes to target's gauges on the start of a new turn.
//	1. Mark target at top of order as "active"
//	2. get that target's current AV
//	3. add AV to "TotalAV" (keeps track of how much AV has progressed over entire sim)
//	4. loop through all targets in the order, subtracing this AV from them (active target gauge = 0)
//			new_gauge = current_gauge - av * speed
func (mgr *manager) StartTurn() (key.TargetID, float64, []event.TurnStatus, error) {
	if mgr.activeTurn {
		return -1, 0, nil, fmt.Errorf("cannot start turn when already in an active turn: %+v", mgr)
	}

	mgr.gaugeCost = 1.0
	mgr.activeTurn = true
	mgr.activeTarget = mgr.order.order[0].id
	av := mgr.av(mgr.activeTarget)

	for _, t := range mgr.order.order {
		t.gauge -= av * mgr.attr.Stats(t.id).SPD()
	}

	mgr.totalAV += av
	return mgr.activeTarget, av, mgr.EventTurnStatus(), nil
}

// ResetTurn resets the gauge of the target taking their turn (target at top of stack) and updates the Manager's turnOrder.
// New gauge is set at BaseGauge * gaugeCost
// Target should then be moved to the bottom of the turn order and then sorted up to the correct
// position based on their AV. In the event of ties, this target should be last in the order.
// Note that this is the same behavior as AddTarget (but different from SetGauge which goes top-down).
// 1. update gauge of active target
// 2. move target to bottom of order
// 3. sort target up based on AV. In the event of tie, this target should be at bottom of order.
// 4. update targetIndexes for all targets that moved in the order (or just repopulate all)
// 5. emit TurnResetEvent
func (mgr *manager) ResetTurn() error {
	if !mgr.activeTurn {
		return fmt.Errorf(
			"target at top of order must have 0 gauge to call reset (their turn is active) %+v", mgr.order.order[0])
	}

	mgr.activeTurn = false
	mgr.order.order[0].gauge = BaseGauge * mgr.gaugeCost
	
	// It would be more efficient to loop through mgr.order ourselves to determine this single target's placement instead of resorting the whole array when no other elements are changing.
	// Unless we are also checking for other SPD changes that happened during the turn, in which case sort.Stable() is better to use, but only after we move the element to the end
	// so as to ensure that, in the case of a tie, it is properly at the tail end of the tied elements.
	mgr.order.order = append(mgr.order.order[1:], mgr.order.order[0])
	sort.Stable(mgr.order)

	// resetTargetAV := BaseGauge / mgr.attr.Stats(mgr.order.order[0].id).SPD()
	// for i, _ := range mgr.order.order[1:] {
	// 	if mgr.order.av(i) > resetTargetAV {
	// 		formerActiveTurn := mgr.order.order[0]
	// 		mgr.order.order = append(mgr.order.order[1:i+1], mgr.order.order[i:]...)
	// 		mgr.order.order[i] = formerActiveTurn
	// 		break
	// 	}
	// 	if i == mgr.order.Len() {
	// 		mgr.order.order = append(mgr.order.order[1:], mgr.order.order[0])
	// 	}
	// }

	mgr.event.TurnReset.Emit(event.TurnReset{
		ResetTarget: mgr.activeTarget,
		GaugeCost:   mgr.gaugeCost,
		TurnOrder:   mgr.EventTurnStatus(),
	})
	return nil
}
