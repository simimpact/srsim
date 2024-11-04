package gallagher

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.Gallagher, character.Config{
		Create:     NewInstance,
		Rarity:     4,
		Element:    model.DamageType_FIRE,
		Path:       model.Path_ABUNDANCE,
		MaxEnergy:  110,
		Promotions: promotions,
		Traces:     traces,
		SkillInfo: character.SkillInfo{
			Attack: character.Attack{
				SPAdd:      1,
				TargetType: model.TargetType_ENEMIES,
			},
			Skill: character.Skill{
				SPNeed:     1,
				TargetType: model.TargetType_ALLIES,
			},
			Ult: character.Ult{
				TargetType: model.TargetType_ENEMIES,
			},
			Technique: character.Technique{
				TargetType: model.TargetType_SELF,
				IsAttack:   true,
			},
		},
	})
}

type char struct {
	engine     engine.Engine
	id         key.TargetID
	info       info.Character
	isEnhanced bool
}

func NewInstance(engine engine.Engine, id key.TargetID, charInfo info.Character) info.CharInstance {
	c := &char{
		engine:     engine,
		id:         id,
		info:       charInfo,
		isEnhanced: false,
	}

	return c
}
