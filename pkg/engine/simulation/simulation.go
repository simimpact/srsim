package simulation

import (
	"context"

	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Simulation struct {
	cfg *model.SimConfig
	//some states and stuff
	//key services
	turnManager TurnManager
}

type TurnManager interface {
	AdvanceTurn() key.TargetID //advance to next turn and update AV accordingly
	CurrentCycle() int         //current cycle count
}

func Run(ctx context.Context, cfg *model.SimConfig) (*model.IterationResult, error) {
	s := &Simulation{
		cfg: cfg,
	}


	return s.run()
}
