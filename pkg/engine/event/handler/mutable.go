package handler

import (
	"github.com/simimpact/srsim/pkg/engine/logging"
	"sort"
)

type MutableListener[E Event] func(event *E)

// Mutable EventHandler that allows listeners to mutate the emitted Event.
// Like the PriorityEventHandler, will execute listeners in order of priority (ascending)
type MutableEventHandler[E Event] struct {
	listeners mutableListeners[E]
}

// Emit a mutable Event to all subscribed listeners, in order of priority (ascending)
func (handler *MutableEventHandler[E]) Emit(event *E) {
	for _, listener := range handler.listeners {
		listener.listener(event)
	}
	logging.Log(event)
}

// Subscribe a listener to this Event handler with the given priority. Listeners are executed
// in ascending order.
func (handler *MutableEventHandler[E]) Subscribe(listener MutableListener[E], priority int) {
	ml := mutableListener[E]{listener: listener, priority: priority}
	handler.listeners = append(handler.listeners, ml)
	sort.Sort(handler.listeners)
}

type mutableListener[E Event] struct {
	listener MutableListener[E]
	priority int
}

type mutableListeners[E Event] []mutableListener[E]

func (a mutableListeners[E]) Len() int           { return len(a) }
func (a mutableListeners[E]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a mutableListeners[E]) Less(i, j int) bool { return a[i].priority < a[j].priority }
