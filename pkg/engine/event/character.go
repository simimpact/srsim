package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type CharacterAddedEventHandler = handler.EventHandler[CharacterAdded]
type CharacterAdded struct {
	ID   key.TargetID
	Info info.Character
}
