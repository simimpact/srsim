package event

import "sort"

type MutableListener[E event] func(event *E)

// Mutable EventHandler that allows listeners to mutate the emitted event.
// Like the PriorityEventHandler, will execute listeners in order of priority (ascending)
type MutableEventHandler[E event] struct {
	listeners mutableListeners[E]
}

// Emit a mutable event to all subscribed listeners, in order of priority (ascending)
func (handler *MutableEventHandler[E]) Emit(event *E) {
	for _, listener := range handler.listeners {
		listener.listener(event)
	}
}

// Subscribe a listener to this event handler with the given priority. Listeners are executed
// in ascending order.
func (handler *MutableEventHandler[E]) Subscribe(listener MutableListener[E], priority int) {
	ml := mutableListener[E]{ listener: listener, priority: priority }
	handler.listeners = append(handler.listeners, ml)
	sort.Sort(handler.listeners)
}

type mutableListener[E event] struct {
	listener MutableListener[E]
	priority int
}

type mutableListeners[E event] []mutableListener[E]

func (a mutableListeners[E]) Len() int { return len(a) }
func (a mutableListeners[E]) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a mutableListeners[E]) Less(i ,j int) bool { return a[i].priority < a[j].priority }