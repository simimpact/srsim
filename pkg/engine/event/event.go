// package event provides an event system that all other functionalities can hook on to
// and either subscribe or emit events
package event

type System struct {
	AttackEvent AttackEventHandler
	DamageEvent DamageEventHandler

	// test placeholder until we get actual events defined
	Ping EventHandler[int]
}

// Interface for all event definitions to ensure common functionality across all events
// TODO: May want to get rid of/replace with requiring protos for all event payloads
type event interface {
	// TODO: helper functions that should exist for all events
	// Ex: ToLog()
	// Ex: String()
	// ToProto()
}