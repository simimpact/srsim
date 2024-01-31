package himeko

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.Himeko, character.Config{
		Create:     NewInstance,
		Rarity:     5,
		Element:    model.DamageType_FIRE,
		Path:       model.Path_ERUDITION,
		MaxEnergy:  120,
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
				IsAttack:   false,
			},
		},
	})
}

type char struct {
	engine       engine.Engine
	id           key.TargetID
	info         info.Character
	talentStacks int
	canAttack    bool
}

func NewInstance(engine engine.Engine, id key.TargetID, charInfo info.Character) info.CharInstance {
	c := &char{
		engine:       engine,
		id:           id,
		info:         charInfo,
		talentStacks: 1,
		canAttack:    true,
	}

	c.initTalent()
	c.initTraces()
	c.initEidolons()

	return c
}
