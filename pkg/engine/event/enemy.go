package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type EnemyAddedEventHandler = handler.EventHandler[EnemyAdded]
type EnemyAdded struct {
	ID   key.TargetID
	Info info.Enemy
}
