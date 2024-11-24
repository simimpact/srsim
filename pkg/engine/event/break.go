package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/key"
)

type BreakExtendEventHandler = handler.EventHandler[BreakExtend]
type BreakExtend struct {
	Key    key.Reason   `json:"key"`
	Target key.TargetID `json:"target"`
}
