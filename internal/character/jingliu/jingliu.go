package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.Jingliu, character.Config{
		Create:     NewInstance,
		Rarity:     5,
		Element:    model.DamageType_ICE,
		Path:       model.Path_DESTRUCTION,
		MaxEnergy:  140,
		Promotions: promotions,
		Traces:     traces,
		SkillInfo: character.SkillInfo{
			Attack: character.Attack{
				SPAdd:      1,
				TargetType: model.TargetType_ENEMIES,
			},
			Skill: character.Skill{
				SPNeed:     0,
				TargetType: model.TargetType_ENEMIES,
				CanUse: func(engine engine.Engine, instance info.CharInstance) bool {
					c := instance.(*char)
					if c.isEnhanced {
						return true
					}
					return engine.SP() >= 1
				},
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
	engine     engine.Engine
	id         key.TargetID
	info       info.Character
	Syzygy     int
	isEnhanced bool
	afterUlt   bool
}

func NewInstance(engine engine.Engine, id key.TargetID, charInfo info.Character) info.CharInstance {
	c := &char{
		engine:     engine,
		id:         id,
		info:       charInfo,
		Syzygy:     0,
		isEnhanced: false,
		afterUlt:   false,
	}
	engine.Events().TurnEnd.Subscribe(c.checkSyzygy)
	engine.Events().ActionStart.Subscribe(c.E1Listener)
	engine.Events().HitStart.Subscribe(c.E2Listener)
	engine.Events().HitStart.Subscribe(c.A6Listener)
	return c
}
