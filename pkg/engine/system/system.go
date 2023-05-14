package system

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type TurnManager interface {
	AdvanceTurn() key.TargetID //advance to next turn and update AV accordingly
	CurrentCycle() int         //current cycle count
}

type Attributes interface {
	Stat(key.TargetID, model.AttributeType) float64
}

type CharacterServices struct {
	TurnService TurnManager
	AttributeService Attributes
}

type EnemyServices struct {

}

type RelicServices struct {

}

type LightconeServices struct {

}