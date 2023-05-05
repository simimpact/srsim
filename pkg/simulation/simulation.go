package simulation

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Simulation struct {
	SimulationConfig
	//some states and stuff
	//key services
	turnManager TurnManager
}

type SimulationConfig struct {
	cfg *model.SimConfig
}

type TurnManager interface {
	AdvanceTurn() key.TargetID //advance to next turn and update AV accordingly
	CurrentCycle() int         //current cycle count
}

func New(cfg SimulationConfig) (*Simulation, error) {
	s := &Simulation{
		SimulationConfig: cfg,
	}

	return s, nil
}
