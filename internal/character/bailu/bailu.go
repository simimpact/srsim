package bailu

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.Bailu, character.Config{
		Create:     NewInstance,
		Rarity:     5,
		Element:    model.DamageType_THUNDER,
		Path:       model.Path_ABUNDANCE,
		MaxEnergy:  100,
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
				TargetType: model.TargetType_ALLIES,
			},
			Technique: character.Technique{
				TargetType: model.TargetType_ALLIES,
				IsAttack:   false,
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

	c.initTalent()
	c.initTraces()
	c.initEidolons()

	return c
}

// Add heal action with BaseHeal based on Bailu's HP.
func (c *char) addHeal(key key.Heal, healPercent, healFlat float64, target []key.TargetID) {
	c.engine.Heal(info.Heal{
		Key:     key,
		Source:  c.id,
		Targets: target,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: healPercent,
		},
		HealValue: healFlat,
	})
}
