package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type InitializeEventHandler = handler.EventHandler[Initialize]
type Initialize struct {
	Config *model.SimConfig `json:"config"`
	Seed   int64            `json:"seed"`
	// TODO: sim metadata (build date, commit hash, etc)?
}

type CharactersAddedEventHandler = handler.EventHandler[CharactersAdded]
type CharactersAdded struct {
	Characters []CharInfo `json:"characters"`
}

type CharInfo struct {
	ID   key.TargetID    `json:"id"`
	Info *info.Character `json:"info"`
}

type EnemiesAddedEventHandler = handler.EventHandler[EnemiesAdded]
type EnemiesAdded struct {
	Enemies []EnemyInfo `json:"enemies"`
}

type EnemyInfo struct {
	ID   key.TargetID `json:"id"`
	Info *info.Enemy  `json:"info"`
}

type BattleStartEventHandler = handler.EventHandler[BattleStart]
type BattleStart struct {
	CharInfo     map[key.TargetID]info.Character `json:"char_info"`
	EnemyInfo    map[key.TargetID]info.Enemy     `json:"enemy_info"`
	CharStats    []*info.Stats                   `json:"char_stats"`
	EnemyStats   []*info.Stats                   `json:"enemy_stats"`
	NeutralStats []*info.Stats                   `json:"neutral_stats"`
}

type Phase1StartEventHandler = handler.EventHandler[Phase1Start]
type Phase1Start struct{}

type Phase1EndEventHandler = handler.EventHandler[Phase1End]
type Phase1End struct{}

type Phase2StartEventHandler = handler.EventHandler[Phase2Start]
type Phase2Start struct{}

type Phase2EndEventHandler = handler.EventHandler[Phase2End]
type Phase2End struct{}

type TurnStartEventHandler = handler.EventHandler[TurnStart]
type TurnStart struct {
	Active     key.TargetID     `json:"active"`
	TargetType info.TargetClass `json:"target_type"`
	DeltaAV    float64          `json:"delta_av"`
	TotalAV    float64          `json:"total_av"`
	TurnOrder  []TurnStatus     `json:"turn_order"`
}

type TurnEndEventHandler = handler.EventHandler[TurnEnd]
type TurnEnd struct {
	Characters []*info.Stats `json:"characters"`
	Enemies    []*info.Stats `json:"enemies"`
	Neutrals   []*info.Stats `json:"neutrals"`
}

type TerminationEventHandler = handler.EventHandler[Termination]
type Termination struct {
	TotalAV float64                 `json:"total_av"`
	Reason  model.TerminationReason `json:"reason"`
}

type ActionStartEventHandler = handler.EventHandler[ActionStart]
type ActionStart struct {
	Owner      key.TargetID     `json:"owner"`
	AttackType model.AttackType `json:"attack_type"`
	IsInsert   bool             `json:"is_insert"`
}

type ActionEndEventHandler = handler.EventHandler[ActionEnd]
type ActionEnd struct {
	Owner      key.TargetID          `json:"owner"`
	Targets    map[key.TargetID]bool `json:"targets"`
	AttackType model.AttackType      `json:"attack_type"`
	IsInsert   bool                  `json:"is_insert"`
}

type InsertStartEventHandler = handler.EventHandler[InsertStart]
type InsertStart struct {
	Owner      key.TargetID         `json:"owner"`
	AbortFlags []model.BehaviorFlag `json:"abort_flags"`
	Priority   info.InsertPriority  `json:"priority"`
}

type InsertEndEventHandler = handler.EventHandler[InsertEnd]
type InsertEnd struct {
	Owner      key.TargetID          `json:"owner"`
	Targets    map[key.TargetID]bool `json:"targets"`
	AbortFlags []model.BehaviorFlag  `json:"abort_flags"`
	Priority   info.InsertPriority   `json:"priority"`
}

type TargetDeathEventHandler = handler.EventHandler[TargetDeath]
type TargetDeath struct {
	Target key.TargetID `json:"target"`
	Killer key.TargetID `json:"killer"`
}
