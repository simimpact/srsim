package simulation

import (
	"encoding/binary"
	"math/rand"

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

type simulation struct {
	cfg  *model.SimConfig
	eval *eval.Eval
	seed int64

	// services
	idGen    *key.TargetIDGenerator
	rand     *rand.Rand
	event    *event.System
	queue    *queue.Handler
	modifier *modifier.Manager
	attr     *attribute.Service
	char     *character.Manager
	enemy    *enemy.Manager
	turn     *turn.Manager
	combat   *combat.Manager
	shield   *shield.Manager

	// state
	sp            int
	tp            int
	targets       map[key.TargetID]info.TargetClass
	characters    []key.TargetID
	enemies       []key.TargetID
	neutrals      []key.TargetID
	totalAV       float64
	active        key.TargetID
	actionTargets map[key.TargetID]bool
}

func Run(cfg *model.SimConfig, eval *eval.Eval, seed int64) (*model.IterationResult, error) {
	s := &simulation{
		cfg:  cfg,
		eval: eval,
		seed: seed,

		event: &event.System{},
		queue: queue.New(),
		rand:  rand.New(rand.NewSource(seed)),
		idGen: key.NewTargetIDGenerator(),

		sp:            3,
		tp:            4, // TODO: define starting amount in config?
		targets:       make(map[key.TargetID]info.TargetClass, 15),
		characters:    make([]key.TargetID, 0, 4),
		enemies:       make([]key.TargetID, 0, 5),
		neutrals:      make([]key.TargetID, 0, 5),
		actionTargets: make(map[key.TargetID]bool, 10),
	}
	s.eval.Engine = s

	// init services

	// core stats
	s.modifier = modifier.NewManager(s)
	s.attr = attribute.New(s.event, s.modifier)

	// target management
	s.char = character.New(s, s.attr, s.eval)
	s.enemy = enemy.New(s, s.attr)

	// game logic
	s.turn = turn.New(s.event, s.attr)
	s.shield = shield.New(s.event, s.attr)
	s.combat = combat.New(s.event, s.attr, s.shield)

	return s.run()
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

func (sim *simulation) Events() *event.System {
	return sim.event
}

func (sim *simulation) Rand() *rand.Rand {
	return sim.rand
}
