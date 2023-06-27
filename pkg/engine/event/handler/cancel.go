package handler

import (
	"github.com/simimpact/srsim/pkg/engine/logging"
	"sort"
)

type CancelableListener[E Event] func(event E) bool

// Cancelable EventHandler that allows listeners to cancel the Event (preventing other listeners
// from being called). Like the PriorityEventHandler, will execute listeners in order of priority
type CancelableEventHandler[E cancellableEvent] struct {
	listeners cancelableListeners[E]
}

// Emit a cancelable Event to all subscribed listeners, in order of priority (ascending).
// The first listener to respond true will cancel the Event and prevent other listeners from being
// called.
func (handler *CancelableEventHandler[E]) Emit(event E) bool {
	for _, listener := range handler.listeners {
		if listener.listener(event) {
			event.Cancelled()
			logging.Log(event)
			return true
		}
	}
	logging.Log(event)
	return false
}

// Subscribe a listener to this Event handler with the given priority. Listeners are executed
// in ascending order.
func (handler *CancelableEventHandler[E]) Subscribe(listener CancelableListener[E], priority int) {
	ml := cancelableListener[E]{listener: listener, priority: priority}
	handler.listeners = append(handler.listeners, ml)
	sort.Sort(handler.listeners)
}

type cancelableListener[E Event] struct {
	listener CancelableListener[E]
	priority int
}

type cancelableListeners[E Event] []cancelableListener[E]

func (a cancelableListeners[E]) Len() int           { return len(a) }
func (a cancelableListeners[E]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a cancelableListeners[E]) Less(i, j int) bool { return a[i].priority < a[j].priority }
