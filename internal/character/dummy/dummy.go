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
	})
}

type char struct {
}

func NewInstance(engine engine.Engine, id key.TargetID, info info.Character) character.CharInstance {
	c := &char{}

	for _, id := range info.Traces {
		switch {
		case id == "1201201" && info.Level >= 1:
			info.BaseStats.Modify(model.Property_ATK_PERCENT, 0.04)
		case id == "1201202" && info.Ascension >= 2:
			info.BaseStats.Modify(model.Property_QUANTUM_DMG_PERCENT, 0.032)
		case id == "1201203" && info.Ascension >= 3:
			info.BaseStats.Modify(model.Property_ATK_PERCENT, 0.0400)
		case id == "1201204" && info.Ascension >= 3:
			info.BaseStats.Modify(model.Property_DEF_PERCENT, 0.0500)
		case id == "1201205" && info.Ascension >= 4:
			info.BaseStats.Modify(model.Property_ATK_PERCENT, 0.0600)
		case id == "1201206" && info.Ascension >= 5:
			info.BaseStats.Modify(model.Property_QUANTUM_DMG_PERCENT, 0.0480)
		case id == "1201207" && info.Ascension >= 5:
			info.BaseStats.Modify(model.Property_ATK_PERCENT, 0.0600)
		case id == "1201208" && info.Ascension >= 6:
			info.BaseStats.Modify(model.Property_DEF_PERCENT, 0.0750)
		case id == "1201209" && info.Level >= 75:
			info.BaseStats.Modify(model.Property_QUANTUM_DMG_PERCENT, 0.0640)
		case id == "1201210" && info.Level >= 80:
			info.BaseStats.Modify(model.Property_ATK_PERCENT, 0.0800)
		}
	}

	return c
}

func (c *char) Attack(engine engine.Engine, target key.TargetID) {

}

func (c *char) Skill(engine engine.Engine, target key.TargetID) {

}

func (c *char) Burst(engine engine.Engine, target key.TargetID) {

}
