package qingque

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.Qingque, character.Config{
		Create:     NewInstance,
		Rarity:     4,
		Element:    model.DamageType_QUANTUM,
		Path:       model.Path_ERUDITION,
		MaxEnergy:  140,
		Promotions: promotions,
		Traces:     traces,
		SkillInfo: character.SkillInfo{
			Attack: character.Attack{
				SPAdd:      0,
				TargetType: model.TargetType_ENEMIES,
			},
			Skill: character.Skill{
				SPNeed:     1,
				TargetType: model.TargetType_SELF,
				CanUse: func(engine engine.Engine, instance info.CharInstance) bool {
					c := instance.(*char)
					return (c.tiles[0] != 4)
				},
			},
			Ult: character.Ult{
				TargetType: model.TargetType_ENEMIES,
			},
			Technique: character.Technique{
				TargetType: model.TargetType_SELF,
				IsAttack:   false,
			},
		},
	})
}

type char struct {
	engine      engine.Engine
	id          key.TargetID
	info        info.Character
	tiles       []int
	suits       []string
	unusedSuits []string
}

func NewInstance(engine engine.Engine, id key.TargetID, charInfo info.Character) info.CharInstance {
	c := &char{
		engine:      engine,
		id:          id,
		info:        charInfo,
		tiles:       []int{0, 0, 0},
		suits:       make([]string, 3),
		unusedSuits: []string{"Wan", "Tong", "Tiao"},
	}
	engine.Events().TurnStart.Subscribe(c.talentTurnStartListener)
	c.initTraces()
	c.initEidolons()
	return c
}
