// root package for all core logic that powers srsim
package engine

//go:generate mockgen -destination=../mock/mock_engine.go -package=mock github.com/simimpact/srsim/pkg/engine Engine

// only event & info are allowed to be imported from engine here
import (
	"math/rand"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Interface that defines how implementations can interact with all engine components. This should
// be thought of as the API definition for all things within the engine. Any implementations will
// be given access to the engine implementation (simulation, outside of the engine package) and call
// these methods to interact.
type Engine interface {
	// Accessor to all event handlers for event subscription and emission
	Events() *event.System

	// Random number generator
	Rand() *rand.Rand

	Modifier
	Attribute
	Combat
	Shield
	Turn
	Validator
	Info
	Target

	// TODO: Execute Queue
	//	For callback + skill methods, need an "AttackState" passed in which allows you to do operations
	//	such as decide when `AttackEnd` happens (if left uncalled, will happen after all logic executes)
	// TODO: Skill Point (Boost Point). Other sim metadata calls?
}

type Modifier interface {
	// Adds a new modifier to the given target. At minimum, instance must specify the name of the
	// modifier and the source.
	//
	// Returns true if the modifier was successfully added, otherwise false. Will always return
	// false if error is non-nil
	AddModifier(target key.TargetID, instance info.Modifier) (bool, error)

	// Removes all instances of a modifier from the target
	RemoveModifier(target key.TargetID, modifier key.Modifier)

	// Removes all instances of a modifier from the target, only where source matches the given source
	RemoveModifierFromSource(target, source key.TargetID, modifier key.Modifier)

	// Extends the duration of all instances of the given modifier by the amount
	ExtendModifierDuration(target key.TargetID, modifier key.Modifier, amt int)

	// Extends the count of all instances of the given modifier by the amount. Will not extend past
	// the modifiers MaxCount
	ExtendModifierCount(target key.TargetID, modifier key.Modifier, amt float64)

	// Returns true if the target has at least one instance of the modifier
	HasModifier(target key.TargetID, modifier key.Modifier) bool

	// Returns the total count of modifiers that are of the given StatusType (Buff or Debuff)
	ModifierCount(target key.TargetID, statusType model.StatusType) int

	// Returns true if the target has the given behavior flag from an attached modifier. If multiple
	// flags are passed, will return true if at least one is attached
	HasBehaviorFlag(target key.TargetID, flags ...model.BehaviorFlag) bool
}

type Attribute interface {
	// Gets a snapshot of the current target's stats. Any modifications to these stats will
	// only be applied to the snapshot.
	Stats(target key.TargetID) *info.Stats

	// Sets the target HP to the given amount. Source target is used for tracking who owns this HP
	// modification in the event that the modification kills the target.
	SetHP(target, source key.TargetID, amt float64) error

	// Modifies the target HP by the given ratio (% of health). The ratio data can include a floor value to
	// ensure that the target HP does not go below the given threshold. Source target is used for
	// tracking who owns this HP modification in the event that the modification kills the target.
	ModifyHPByRatio(target, source key.TargetID, data info.ModifyHPByRatio) error

	// Modifies the target HP by the given flat amount. Source target is used for tracking who owns
	// this HP modification in the event that the modification kills the target.
	ModifyHPByAmount(target, source key.TargetID, amt float64) error

	// Modifies the target stance by the given flat amount. Source target is used for tracking who
	// owns this stance modification in the event that the stance reaches 0 and a break is triggered.
	// This stance modification will also scale with the source's ALL_STANCE_DMG_PERCENT.
	ModifyStance(target, source key.TargetID, amt float64) error

	// Modifies the target energy by the given flat amount. Energy amount added will be multiplied
	// by the target's current Energy Regeneration amount.
	ModifyEnergy(target key.TargetID, amt float64) error

	// Modifies the target energy by the given flat amount. This amount is fixed and will not be
	// increased by the target's Energy Regeneration.
	ModifyEnergyFixed(target key.TargetID, amt float64) error
}

type Combat interface {
	// Performs the given attack where Source is the attacker and Targets are all targets that
	// are being hit
	Attack(atk info.Attack)

	// Performs the given heal where Source is the healer and Targets are all targets that are
	// being healed
	Heal(heal info.Heal)
}

type Shield interface {
	// TODO:
	AddShield()
	RemoveShield()
}

type Turn interface {
	ModifyGauge(target key.TargetID, modifyType model.ModifyGauge, amt float64)
	SetGauge(target key.TargetID, amt float64)
	// TODO: need ModifyCurrentSkillDelayCost? (in dm used to modify gauge for next turn, during current turn)
}

type Validator interface {
	// Check if the given TargetID is valid
	IsValid(target key.TargetID) bool
}

type Info interface {
	// Metadata for the given character, such as their current level, ascension, traces, etc.
	CharacterInfo(target key.TargetID) (info.Character, error)

	// Metadata for the given enemy, such as their current level and weaknesses.
	EnemyInfo(target key.TargetID) (model.Enemy, error)
}

type Target interface {
	// returns true if the given TargetID is for a character
	IsCharacter(target key.TargetID) bool

	// returns true if the given TargetID is for an enemy
	IsEnemy(target key.TargetID) bool

	// returns the ids of targets that are adjacent to the given targent (empty if there are none)
	AdjacentTo(target key.TargetID) []key.TargetID

	// returns a list of all character target ids
	Characters() []key.TargetID

	// returns a list of all enemy target ids
	Enemies() []key.TargetID

	// returns a list of all neutral target ids (these are special cases, such as the Lightning-Lord)
	Neutrals() []key.TargetID

	// TODO: target type, (Light, Dark, Neutral)
	AddTarget() key.TargetID
}
