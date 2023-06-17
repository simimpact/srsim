package handler

import "sort"

type CancelableListener[E event] func(event E) bool

// Cancelable EventHandler that allows listeners to cancel the event (preventing other listeners
// from being called). Like the PriorityEventHandler, will execute listeners in order of priority
type CancelableEventHandler[E event] struct {
	listeners cancelableListeners[E]
}

// Emit a cancelable event to all subscribed listeners, in order of priority (ascending).
// The first listener to respond true will cancel the event and prevent other listeners from being
// called.
func (handler *CancelableEventHandler[E]) Emit(event E) bool {
	for _, listener := range handler.listeners {
		if listener.listener(event) {
			Singleton.Log(event)
			return true
		}
	}
	Singleton.Log(event)
	return false
}

// Subscribe a listener to this event handler with the given priority. Listeners are executed
// in ascending order.
func (handler *CancelableEventHandler[E]) Subscribe(listener CancelableListener[E], priority int) {
	ml := cancelableListener[E]{listener: listener, priority: priority}
	handler.listeners = append(handler.listeners, ml)
	sort.Sort(handler.listeners)
}

type cancelableListener[E event] struct {
	listener CancelableListener[E]
	priority int
}

type cancelableListeners[E event] []cancelableListener[E]

func (a cancelableListeners[E]) Len() int           { return len(a) }
func (a cancelableListeners[E]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a cancelableListeners[E]) Less(i, j int) bool { return a[i].priority < a[j].priority }
