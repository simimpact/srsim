package blade

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.Blade, character.Config{
		Create:     NewInstance,
		Rarity:     5,
		Element:    model.DamageType_WIND,
		Path:       model.Path_DESTRUCTION,
		MaxEnergy:  130,
		Promotions: promotions,
		Traces:     traces,
		SkillInfo: character.SkillInfo{
			Attack: character.Attack{
				SPAdd:      0,
				TargetType: model.TargetType_ENEMIES,
			},
			Skill: character.Skill{
				SPNeed:     1,
				TargetType: model.TargetType_SELF,
				CanUse: func(engine engine.Engine, instance info.CharInstance) bool {
					c := instance.(*char)
					return !c.engine.HasModifier(c.id, Hellscape)
				},
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
	engine    engine.Engine
	id        key.TargetID
	info      info.Character
	hpLoss    float64
	charge    int
	maxCharge int
}

func NewInstance(engine engine.Engine, id key.TargetID, charInfo info.Character) info.CharInstance {
	c := &char{
		engine:    engine,
		id:        id,
		info:      charInfo,
		hpLoss:    0.0,
		charge:    0,
		maxCharge: 5,
	}

	if c.info.Eidolon >= 6 {
		c.maxCharge = 4
	}

	c.initTraces()

	engine.Events().HPChange.Subscribe(c.hpLossListener)
	engine.Events().AttackStart.Subscribe(c.onBeforeBeingHitListener)
	engine.Events().AttackEnd.Subscribe(c.onListenAfterAttackListener)

	return c
}

func (c *char) hpLossListener(e event.HPChange) {
	if e.Target != c.id {
		return
	}

	if e.OldHP >= e.NewHP {
		return
	}

	if c.engine.HasModifier(c.id, IsAttack) {
		if !c.engine.HasModifier(c.id, GainedCharge) {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:     Talent,
				Source:   c.id,
				MaxCount: float64(c.maxCharge),
			})
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   GainedCharge,
				Source: c.id,
			})
		}
	} else {
		c.charge++
	}

	if c.charge >= c.maxCharge {
		c.Talent()
		c.charge = 0
	}

	hpChange := e.NewHP - e.OldHP
	c.hpLoss += hpChange

	// E4
	// TODO: Determine Timing?

	if c.info.Eidolon >= 4 {
		if e.OldHPRatio > 0.5 && e.NewHPRatio <= 0.5 && c.engine.ModifierStackCount(c.id, c.id, E4) < 2 {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   E4,
				Source: c.id,
			})
		}
	}
}
