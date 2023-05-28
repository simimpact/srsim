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
		Path:       model.Path_ERUDITION,
		MaxEnergy:  120,
		Promotions: promotions,
		Traces:     traces,
	})
}

type char struct {
	engine engine.Engine
	id     key.TargetID
	info   info.Character
}

func NewInstance(engine engine.Engine, id key.TargetID, info info.Character) character.CharInstance {
	c := &char{
		engine: engine,
		id:     id,
		info:   info,
	}

	c.a2()
	c.a4()
	c.a6()

	return c
}

func (c *char) Attack(target key.TargetID) {

}

func (c *char) Skill(target key.TargetID) {

}

func (c *char) Burst(target key.TargetID) {

}
