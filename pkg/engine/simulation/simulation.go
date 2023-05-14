package simulation

import (
	"context"

	"github.com/simimpact/srsim/pkg/engine/system"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Target interface {
	Exec(key.ActionType)
}

type Simulation struct {
	cfg *model.SimConfig
	//some states and stuff
	actioners map[key.TargetID]Target
	//key services
	turnManager system.TurnManager


	//??
	res *model.IterationResult
}

func Run(ctx context.Context, cfg *model.SimConfig) (*model.IterationResult, error) {
	s := &Simulation{
		cfg: cfg,
		actioners: make(map[key.TargetID]Target),
		res: &model.IterationResult{
			TargetIdMapping: make(map[int32]string),
		},
	}


	return s.run()
}
