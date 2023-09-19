package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.DanHengImbibitorLunae, character.Config{
		Create:     NewInstance,
		Rarity:     5,
		Element:    model.DamageType_IMAGINARY,
		Path:       model.Path_DESTRUCTION,
		MaxEnergy:  140,
		Promotions: promotions,
		Traces:     traces,
		SkillInfo: character.SkillInfo{
			Attack: character.Attack{
				SPAdd:      0,
				TargetType: model.TargetType_ENEMIES,
			},
			Skill: character.Skill{
				SPNeed:     0,
				TargetType: model.TargetType_SELF,
				CanUse: func(engine engine.Engine, instance info.CharInstance) bool {
					c := instance.(*char)
					return engine.HasModifier(c.id, Point)
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
	engine.Events().ActionEnd.Subscribe(c.E6ActionEndListener)
	c.initSkill()
	c.initTalent()
	c.initTraces()
	return c
}
