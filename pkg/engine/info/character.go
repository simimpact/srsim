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
