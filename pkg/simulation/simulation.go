package simulation

import (
	"context"
	"math/rand"

	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/engine/turn"
	"github.com/simimpact/srsim/pkg/gcs"
	"github.com/simimpact/srsim/pkg/gcs/eval"
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

	// services
	idGen            *key.TargetIDGenerator
	rand             *rand.Rand
	event            *event.System
	modManager       *modifier.Manager
	attributeService *attribute.Service
	charManager      *character.Manager
	turnManager      *turn.TurnCtrl
	eval             *eval.Eval

	//??
	res *model.IterationResult
}

func Run(ctx context.Context, list *gcs.ActionList, cfg *model.SimConfig) (*model.IterationResult, error) {
	s := &Simulation{
		cfg:       cfg,
		actioners: make(map[key.TargetID]Target),

		rand:  rand.New(rand.NewSource(1)), // TODO: seed
		event: &event.System{},
		idGen: key.NewTargetIDGenerator(),

		res: &model.IterationResult{
			TargetIdMapping: make(map[int32]string),
		},
	}

	// init services
	s.modManager = modifier.NewManager(s)
	s.attributeService = attribute.New(s, s.modManager)
	s.charManager = character.New(s, s.attributeService)
	// turnManager: turn.New(),
	s.eval = eval.New(list.Program, ctx)

	return s.run()
}
