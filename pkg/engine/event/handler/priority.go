package handler

import (
	"github.com/simimpact/srsim/pkg/engine/logging"
	"sort"
)

// Priority EventHandler that will execute listeners in order of priority (ascending)
type PriorityEventHandler[E Event] struct {
	listeners priorityListeners[E]
}

// Emit an Event to all subscribed listeners, in order of priority (ascending)
func (handler *PriorityEventHandler[E]) Emit(event E) {
	for _, listener := range handler.listeners {
		listener.listener(event)
	}
	logging.Log(event)
}

// Subscribe a listener to this Event handler with the given priority. Listeners are executed
// in ascending order.
func (handler *PriorityEventHandler[E]) Subscribe(listener Listener[E], priority int) {
	pl := priorityListener[E]{listener: listener, priority: priority}
	handler.listeners = append(handler.listeners, pl)
	sort.Sort(handler.listeners)
}

type priorityListener[E Event] struct {
	listener Listener[E]
	priority int
}

type priorityListeners[E Event] []priorityListener[E]

func (a priorityListeners[E]) Len() int           { return len(a) }
func (a priorityListeners[E]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a priorityListeners[E]) Less(i, j int) bool { return a[i].priority < a[j].priority }
