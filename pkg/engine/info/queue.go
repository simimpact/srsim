package info

//go:generate stringer -type=InsertPriority

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// metadata on what is being inserted into the turn queue
type Insert struct {
	// The logic to run when this insert is executed during the turn
	Execute func()

	// The target this insert is associated with (used to look up their current state)
	Source key.TargetID

	// Before this insert executes, will check the Source for if it contains any of the given
	// abort flags. If an abort flag matches, this insert will not be executed.
	AbortFlags []model.BehaviorFlag

	// The priority of this insert within the turn queue
	Priority InsertPriority
}

type InsertPriority int

// TODO: specific insert priorities for specific casses (IE: himiko vs herta, clara vs march)
const (
	CharReviveSelf          InsertPriority = 45
	CharHealSelf            InsertPriority = 48
	CharReviveOthers        InsertPriority = 55
	CharHealOthers          InsertPriority = 58
	CharBuffSelf            InsertPriority = 65
	CharInsertAttackSelf    InsertPriority = 75
	CharInsertAttackOthers  InsertPriority = 115
	EnemyReviveSelf         InsertPriority = 145
	EnemyHealSelf           InsertPriority = 148
	EnemyReviveOthers       InsertPriority = 155
	EnemyHealOthers         InsertPriority = 158
	EnemyBuffSelf           InsertPriority = 165
	EnemyInsertAttackSelf   InsertPriority = 175
	EnemyInsertAttackOthers InsertPriority = 215
	CharInsertAction        InsertPriority = 500
	EnemyInsertAction       InsertPriority = 1000
)

func (i InsertPriority) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}
