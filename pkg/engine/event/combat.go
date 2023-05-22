package event

import "github.com/simimpact/srsim/pkg/engine/event/handler"

type AttackEventHandler = handler.EventHandler[AttackEvent]
type AttackEvent struct {
	// TODO
}

type DamageEventHandler = handler.EventHandler[DamageEvent]
type DamageEvent struct {
	// TODO
}
