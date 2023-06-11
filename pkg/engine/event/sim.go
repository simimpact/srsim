package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type InitializeEventHandler = handler.EventHandler[InitializeEvent]
type InitializeEvent struct {
	Config *model.SimConfig
	Seed   int64
	// TODO: sim metadata (build date, commit hash, etc)?
}

type BattleStartEventHandler = handler.EventHandler[BattleStartEvent]
type BattleStartEvent struct {
	CharInfo     map[key.TargetID]info.Character
	EnemyInfo    map[key.TargetID]info.Enemy
	CharStats    []*info.Stats
	EnemyStats   []*info.Stats
	NeutralStats []*info.Stats
}

type TurnEndEventHandler = handler.EventHandler[TurnEndEvent]
type TurnEndEvent struct {
	Characters []*info.Stats
	Enemies    []*info.Stats
	Neutrals   []*info.Stats
}

type TerminationEventHandler = handler.EventHandler[TerminationEvent]
type TerminationEvent struct {
	TotalAV float64
	Reason  model.TerminationReson
}

type ActionStartEventHandler = handler.EventHandler[ActionEvent]
type ActionEndEventHandler = handler.EventHandler[ActionEvent]
type ActionEvent struct {
	Target     key.TargetID
	AttackType model.AttackType
	IsInsert   bool
}

type UltStartEventHandler = handler.EventHandler[ActionEvent]
type UltEndEventHandler = handler.EventHandler[ActionEvent]

type InsertStartEventHandler = handler.EventHandler[InsertEvent]
type InsertEndEventHandler = handler.EventHandler[InsertEvent]
type InsertEvent struct {
	Target     key.TargetID
	AbortFlags []model.BehaviorFlag
	Priority   info.InsertPriority
}
