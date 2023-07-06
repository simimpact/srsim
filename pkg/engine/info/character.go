package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Character struct {
	Key          key.Character     `json:"key"`
	Level        int               `json:"level"`
	Ascension    int               `json:"ascension"`
	Eidolon      int               `json:"eidolon"`
	Traces       map[string]bool   `json:"traces"`
	AbilityLevel AbilityLevel      `json:"ability_level"`
	Path         model.Path        `json:"path"`
	Element      model.DamageType  `json:"element"`
	LightCone    LightCone         `json:"light_cone"`
	Relics       map[key.Relic]int `json:"relics"`
}

type AbilityLevel struct {
	Attack int `json:"attack"`
	Skill  int `json:"skill"`
	Ult    int `json:"ult"`
	Talent int `json:"talent"`
}

type LightCone struct {
	Key        key.LightCone `json:"key"`
	Level      int           `json:"level"`
	Ascension  int           `json:"ascension"`
	Imposition int           `json:"imposition"`
	Path       model.Path    `json:"path"`
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

	// Returns info/metadata on the character, such as eidolon level, set of enabled traces, current
	// level, max level, level of each skill, etc
	CharacterInfo() Character

	// Will end the current active attack. When this happens is different for each skill implementation
	// so it is important that it is correctly called at the right time.
	EndAttack()
}

// gets the current attack level in base-0. Useful for indexing by level in implementation
func (i Character) AttackLevelIndex() int {
	return i.AbilityLevel.Attack - 1
}

// gets the current skill level in base-0. Useful for indexing by level in implementation
func (i Character) SkillLevelIndex() int {
	return i.AbilityLevel.Skill - 1
}

// gets the current ult level in base-0. Useful for indexing by level in implementation
func (i Character) UltLevelIndex() int {
	return i.AbilityLevel.Ult - 1
}

// gets the current talent level in base-0. Useful for indexing by level in implementation
func (i Character) TalentLevelIndex() int {
	return i.AbilityLevel.Talent - 1
}
