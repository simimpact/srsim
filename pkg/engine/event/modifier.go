package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type ModifierAddedEventHandler = handler.EventHandler[ModifierAddedEvent]
type ModifierAddedEvent struct {
	Target   key.TargetID
	Modifier *info.ModifierInstance
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
	Modifier *info.ModifierInstance
}

type ModifierExtendedEventHandler = handler.EventHandler[ModifierExtendedEvent]
type ModifierExtendedEvent struct {
	Target    key.TargetID
	Modifier  *info.ModifierInstance
	Operation string
	OldValue  int
	NewValue  int
}
