package info

import "github.com/simimpact/srsim/pkg/key"

// these need to be in a separate file for typescript schema generations

type CharInstance interface {
	Attack(target key.TargetID, state ActionState)
	Skill(target key.TargetID, state ActionState)
	Technique(target key.TargetID, state ActionState)
}

type SingleUlt interface {
	Ult(target key.TargetID, state ActionState)
}

type MultiUlt interface {
	UltAttack(target key.TargetID, state ActionState)
	UltSkill(target key.TargetID, state ActionState)
}

// This represents the current state of the action that the implementation is executing. It provides
// some data helpers + important interactions for the action that cannot be found in Engine
type ActionState interface {
	// Returns true if this action being executed is an insert. This will be true for cases such as:
	// 	- Seele extra turn
	// 	- QQ skill reuse
	IsInsert() bool

	// Returns info/metadata on the character, such as eidolon level, set of enabled traces, current
	// level, max level, level of each skill, etc
	CharacterInfo() Character

	// Will end the current active attack. When this happens is different for each skill implementation
	// so it is important that it is correctly called at the right time.
	EndAttack()
}
