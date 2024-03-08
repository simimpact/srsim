// package dummy implements a dummy character for testing purposes
package dummy

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.DummyCharacter, character.Config{
		Create:     NewInstance,
		Rarity:     4,
		Element:    model.DamageType_QUANTUM,
		Path:       model.Path_ABUNDANCE,
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
				TargetType: model.TargetType_ALLIES,
			},
			Ult: character.Ult{
				TargetType: model.TargetType_SELF,
			},
			Technique: character.Technique{
				TargetType: model.TargetType_ALLIES,
				IsAttack:   true,
			},
		},
	})
}

type char struct {
	engine engine.Engine
	id     key.TargetID
	info   info.Character
}

func NewInstance(engine engine.Engine, id key.TargetID, charInfo info.Character) info.CharInstance {
	c := &char{
		engine: engine,
		id:     id,
		info:   charInfo,
	}

	c.a2()
	c.a4()
	c.a6()

	return c
}

func (c *char) Attack(target key.TargetID, state info.ActionState) {

}

func (c *char) Skill(target key.TargetID, state info.ActionState) {

}

func (c *char) Ult(target key.TargetID, state info.ActionState) {

}

func (c *char) Technique(target key.TargetID, state info.ActionState) {

}
