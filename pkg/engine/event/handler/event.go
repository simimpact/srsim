package handler

// Interface for all event definitions to ensure common functionality across all events
// TODO: May want to get rid of/replace with requiring protos for all event payloads
type event interface {
	// TODO: helper functions that should exist for all events
	// Ex: ToLog()
	// Ex: String()
	// ToProto()
}
