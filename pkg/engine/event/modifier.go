package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type ModifierAddedEventHandler = handler.EventHandler[ModifierAdded]
type ModifierAdded struct {
	Target   key.TargetID
	Modifier info.Modifier
	Chance   float64
}

type ModifierResistedEventHandler = handler.EventHandler[ModifierResisted]
type ModifierResisted struct {
	Target     key.TargetID
	Source     key.TargetID
	Modifier   key.Modifier
	Chance     float64
	BaseChance float64
	EHR        float64
	EffectRES  float64
	DebuffRES  float64
}

type ModifierRemovedEventHandler = handler.EventHandler[ModifierRemoved]
type ModifierRemoved struct {
	Target   key.TargetID
	Modifier info.Modifier
}

type ModifierExtendedDurationEventHandler = handler.EventHandler[ModifierExtendedDuration]
type ModifierExtendedDuration struct {
	Target   key.TargetID
	Modifier info.Modifier
	OldValue int
	NewValue int
}

type ModifierExtendedCountEventHandler = handler.EventHandler[ModifierExtendedCount]
type ModifierExtendedCount struct {
	Target   key.TargetID
	Modifier info.Modifier
	OldValue float64
	NewValue float64
}
