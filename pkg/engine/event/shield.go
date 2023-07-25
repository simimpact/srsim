package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type ShieldAddedEventHandler = handler.EventHandler[ShieldAdded]
type ShieldAdded struct {
	ID           key.Shield  `json:"id"`
	Info         info.Shield `json:"info"`
	ShieldHealth float64     `json:"shield_health"`
}

type ShieldRemovedEventHandler = handler.EventHandler[ShieldRemoved]
type ShieldRemoved struct {
	ID     key.Shield   `json:"id"`
	Target key.TargetID `json:"target"`
}

type ShieldChangeEventHandler = handler.EventHandler[ShieldChange]
type ShieldChange struct {
	ID     key.Shield   `json:"id"`
	Target key.TargetID `json:"target"`
	OldHP  float64      `json:"old_hp"`
	NewHP  float64      `json:"new_hp"`
}
