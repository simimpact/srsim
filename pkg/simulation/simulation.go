package simulation

import (
	"encoding/binary"
	"math/rand"

	"github.com/simimpact/srsim/pkg/engine/logging"

	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/combat"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/queue"
	"github.com/simimpact/srsim/pkg/engine/shield"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/engine/target/enemy"
	"github.com/simimpact/srsim/pkg/engine/turn"
	"github.com/simimpact/srsim/pkg/gcs/eval"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Target interface {
	Exec(key.ActionType)
}

type Simulation struct {
	cfg  *model.SimConfig
	eval *eval.Eval
	seed int64

	// services
	IdGen    *key.TargetIDGenerator
	Random   *rand.Rand
	Event    *event.System
	Queue    *queue.Handler
	Modifier *modifier.Manager
	Attr     *attribute.Service
	Char     *character.Manager
	Enemy    *enemy.Manager
	Turn     *turn.Manager
	Combat   *combat.Manager
	Shield   *shield.Manager

	// state
	Sp            int
	Tp            int
	Targets       map[key.TargetID]info.TargetClass
	characters    []key.TargetID
	enemies       []key.TargetID
	neutrals      []key.TargetID
	TotalAV       float64
	Active        key.TargetID
	ActionTargets map[key.TargetID]bool
}

func RunWithLog(logger logging.Logger, cfg *model.SimConfig, eval *eval.Eval, seed int64) (*model.IterationResult, error) {
	logging.InitLogger(logger)
	return Run(cfg, eval, seed)
}

func Run(cfg *model.SimConfig, eval *eval.Eval, seed int64) (*model.IterationResult, error) {
	s := NewSimulation(cfg, eval, seed)
	return s.Run()
}

func NewSimulation(cfg *model.SimConfig, eval *eval.Eval, seed int64) *Simulation {
	s := &Simulation{
		cfg:  cfg,
		eval: eval,
		seed: seed,

		Event:  &event.System{},
		Queue:  queue.New(),
		Random: rand.New(rand.NewSource(seed)),
		IdGen:  key.NewTargetIDGenerator(),

		Sp:            3,
		Tp:            4, // TODO: define starting amount in config?
		Targets:       make(map[key.TargetID]info.TargetClass, 15),
		characters:    make([]key.TargetID, 0, 4),
		enemies:       make([]key.TargetID, 0, 5),
		neutrals:      make([]key.TargetID, 0, 5),
		ActionTargets: make(map[key.TargetID]bool, 10),
	}
	s.eval.Engine = s

	// init services

	// core stats
	s.Modifier = modifier.NewManager(s)
	s.Attr = attribute.New(s.Event, s.Modifier)

	// target management
	s.Char = character.New(s, s.Attr, s.eval)
	s.Enemy = enemy.New(s, s.Attr)

	// game logic
	s.Turn = turn.New(s.Event, s.Attr)
	s.Shield = shield.New(s.Event, s.Attr)
	s.Combat = combat.New(s.Event, s.Attr, s.Shield, s, s.Random)

	return s
}

// TODO: RunWithDebug

func RandSeed() (int64, error) {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(b[:])), nil
}

func (sim *Simulation) Events() *event.System {
	return sim.Event
}

func (sim *Simulation) Rand() *rand.Rand {
	return sim.Random
}
