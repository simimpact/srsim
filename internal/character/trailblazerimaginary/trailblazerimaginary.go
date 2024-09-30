package trailblazerimaginary

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.TrailblazerImaginary, character.Config{
		Create:     NewInstance,
		Rarity:     5,
		Element:    model.DamageType_IMAGINARY,
		Path:       model.Path_HARMONY,
		MaxEnergy:  140,
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
				TargetType: model.TargetType_ALLIES,
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
	ultLifeTime int
	E1Used      bool
}

func NewInstance(engine engine.Engine, id key.TargetID, charInfo info.Character) info.CharInstance {
	c := &char{
		engine:      engine,
		id:          id,
		info:        charInfo,
		ultLifeTime: 0,
		E1Used:      false,
	}
	engine.Events().TurnStart.Subscribe(c.buffListener)
	engine.Events().StanceBreak.Subscribe(c.talentListener)
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(id, info.Modifier{
			Name:   E2Buff,
			Source: c.id,
			Stats: info.PropMap{
				prop.EnergyRegen: 0.25,
			},
		})
	}
	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(id, info.Modifier{
			Name:   E4ListenerBuff,
			Source: c.id,
		})
	}
	return c
}
