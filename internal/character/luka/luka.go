package luka

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.Luka, character.Config{
		Create:     NewInstance,
		Rarity:     4,
		Element:    model.DamageType_PHYSICAL,
		Path:       model.Path_NIHILITY,
		MaxEnergy:  130,
		Promotions: promotions,
		Traces:     traces,
		SkillInfo: character.SkillInfo{
			Attack: character.Attack{
				SPAdd:      1,
				TargetType: model.TargetType_ENEMIES,
			},
			Skill: character.Skill{
				SPNeed:     1,
				TargetType: model.TargetType_ENEMIES,
			},
			Ult: character.Ult{
				TargetType: model.TargetType_ENEMIES,
			},
			Technique: character.Technique{
				TargetType: model.TargetType_ENEMIES,
				IsAttack:   true,
			},
		},
	})
}

type char struct {
	engine         engine.Engine
	id             key.TargetID
	info           info.Character
	fightingSpirit int
}

func NewInstance(engine engine.Engine, id key.TargetID, charInfo info.Character) info.CharInstance {
	c := &char{
		engine:         engine,
		id:             id,
		info:           charInfo,
		fightingSpirit: 0,
	}

	c.initTalent()
	return c
}
