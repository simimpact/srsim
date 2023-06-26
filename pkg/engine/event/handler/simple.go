package handler

import "github.com/simimpact/srsim/pkg/engine/logging"

type Listener[E Event] func(event E)

// Simple EventHandler that on Emit will run all listeners in the order of their subscription
type EventHandler[E Event] struct {
	listeners []Listener[E]
}

// Emit an Event to all subscribed listeners, in the order they subscribed (non-deterministic order)
func (handler *EventHandler[E]) Emit(event E) {
	for _, listener := range handler.listeners {
		listener(event)
	}
	logging.Log(event)
}

// Subscribe a listener to this Event handler to be executed when Emit is called
func (handler *EventHandler[E]) Subscribe(listener Listener[E]) {
	handler.listeners = append(handler.listeners, listener)
}
