package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type InitializeEventHandler = handler.EventHandler[Initialize]
type Initialize struct {
	Config *model.SimConfig
	Seed   int64
	// TODO: sim metadata (build date, commit hash, etc)?
}

type BattleStartEventHandler = handler.EventHandler[BattleStart]
type BattleStart struct {
	CharInfo     map[key.TargetID]info.Character
	EnemyInfo    map[key.TargetID]info.Enemy
	CharStats    []*info.Stats
	EnemyStats   []*info.Stats
	NeutralStats []*info.Stats
}

type TurnStartEventHandler = handler.EventHandler[TurnStart]
type TurnStart struct {
	Active    key.TargetID
	DeltaAV   float64
	TotalAV   float64
	TurnOrder []TurnStatus
}

type TurnEndEventHandler = handler.EventHandler[TurnEnd]
type TurnEnd struct {
	Characters []*info.Stats
	Enemies    []*info.Stats
	Neutrals   []*info.Stats
}

type TerminationEventHandler = handler.EventHandler[Termination]
type Termination struct {
	TotalAV float64
	Reason  model.TerminationReason
}

type ActionStartEventHandler = handler.EventHandler[ActionStart]
type ActionStart struct {
	Owner      key.TargetID
	AttackType model.AttackType
	IsInsert   bool
}

type ActionEndEventHandler = handler.EventHandler[ActionEnd]
type ActionEnd struct {
	Owner      key.TargetID
	Targets    map[key.TargetID]bool
	AttackType model.AttackType
	IsInsert   bool
}

type InsertStartEventHandler = handler.EventHandler[InsertStart]
type InsertStart struct {
	Owner      key.TargetID
	AbortFlags []model.BehaviorFlag
	Priority   info.InsertPriority
}

type InsertEndEventHandler = handler.EventHandler[InsertEnd]
type InsertEnd struct {
	Owner      key.TargetID
	Targets    map[key.TargetID]bool
	AbortFlags []model.BehaviorFlag
	Priority   info.InsertPriority
}
