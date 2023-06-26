package handler

// Interface for all Event definitions to ensure common functionality across all events
// TODO: May want to get rid of/replace with requiring protos for all Event payloads
type Event interface {
	// TODO: helper functions that should exist for all events
	// Ex: ToLog()
	// Ex: String()
	// ToProto()
}

type cancellableEvent interface {
	Cancelled()
}
