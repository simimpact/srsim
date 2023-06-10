package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Character struct {
	Key          key.Character
	Level        int
	Ascension    int
	Eidolon      int
	Traces       map[string]bool
	AbilityLevel AbilityLevel
	Path         model.Path
	Element      model.DamageType
	LightCone    LightCone
	Relics       map[key.Relic]int
}

type AbilityLevel struct {
	Attack int
	Skill  int
	Ult    int
	Talent int
}

type LightCone struct {
	Key       key.LightCone
	Level     int
	Ascension int
	Rank      int
	Path      model.Path
}

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

	// Returns the SkillEffect of this action. This should match what was defined in the
	// character config.
	SkillEffect() model.SkillEffect

	// Returns info/metadata on the character, such as eidolon level, set of enabled traces, current
	// level, max level, level of each skill, etc
	CharacterInfo() Character

	// Will end the current active attack. When this happens is different for each skill implementation
	// so it is important that it is correctly called at the right time.
	EndAttack()
}
