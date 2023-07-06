package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type ModifierAddedEventHandler = handler.EventHandler[ModifierAdded]
type ModifierAdded struct {
	Target   key.TargetID  `json:"target"`
	Modifier info.Modifier `json:"modifier"`
	Chance   float64       `json:"chance"`
}

type ModifierResistedEventHandler = handler.EventHandler[ModifierResisted]
type ModifierResisted struct {
	Target        key.TargetID `json:"target"`
	Source        key.TargetID `json:"source"`
	Modifier      key.Modifier `json:"modifier"`
	Chance        float64      `json:"chance"`
	BaseChance    float64      `json:"base_chance"`
	EffectHitRate float64      `json:"effect_hit_rate"`
	EffectRES     float64      `json:"effect_res"`
	DebuffRES     float64      `json:"debuff_res"`
}

type ModifierRemovedEventHandler = handler.EventHandler[ModifierRemoved]
type ModifierRemoved struct {
	Target   key.TargetID  `json:"target"`
	Modifier info.Modifier `json:"modifier"`
}

type ModifierExtendedDurationEventHandler = handler.EventHandler[ModifierExtendedDuration]
type ModifierExtendedDuration struct {
	Target   key.TargetID  `json:"target"`
	Modifier info.Modifier `json:"modifier"`
	OldValue int           `json:"old_value"`
	NewValue int           `json:"new_value"`
}

type ModifierExtendedCountEventHandler = handler.EventHandler[ModifierExtendedCount]
type ModifierExtendedCount struct {
	Target   key.TargetID  `json:"target"`
	Modifier info.Modifier `json:"modifier"`
	OldValue float64       `json:"old_value"`
	NewValue float64       `json:"new_value"`
}
