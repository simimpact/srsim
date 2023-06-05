package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/model"
)

type InitializeEventHandler = handler.EventHandler[InitializeEvent]
type InitializeEvent struct {
	Config *model.SimConfig
	Seed   int64
	// TODO: sim metadata (build date, commit hash, etc)?
}

// TODO: event data (current state of enemies and characters?)
// can we/should we reuse this event to occur at the start of each wave?
type BattleStartEventHandler = handler.EventHandler[struct{}]

type ActionStartEvent struct {
}

type ActionEndEvent struct {
}
