package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type ModifierAddedEventHandler = handler.EventHandler[ModifierAddedEvent]
type ModifierAddedEvent struct {
	Target   key.TargetID
	Modifier info.Modifier
	Chance   float64
}

type ModifierResistedEventHandler = handler.EventHandler[ModifierResistedEvent]
type ModifierResistedEvent struct {
	Target     key.TargetID
	Source     key.TargetID
	Modifier   key.Modifier
	Chance     float64
	BaseChance float64
	EHR        float64
	EffectRES  float64
	DebuffRES  float64
}

type ModifierRemovedEventHandler = handler.EventHandler[ModifierRemovedEvent]
type ModifierRemovedEvent struct {
	Target   key.TargetID
	Modifier info.Modifier
}

type ModifierExtendedDurationEventHandler = handler.EventHandler[ModifierExtendedDurationEvent]
type ModifierExtendedDurationEvent struct {
	Target   key.TargetID
	Modifier info.Modifier
	OldValue int
	NewValue int
}

type ModifierExtendedCountEventHandler = handler.EventHandler[ModifierExtendedCountEvent]
type ModifierExtendedCountEvent struct {
	Target   key.TargetID
	Modifier info.Modifier
	OldValue float64
	NewValue float64
}
