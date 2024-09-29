package luocha

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill         = "luocha-skill"
	SkillInsert   = "luocha-skill-insert"
	SkillInsertCD = "luocha-skill-insert-cooldown"
)

var IsInsert bool

func init() {
	modifier.Register(SkillInsert, modifier.Config{})

	modifier.Register(SkillInsertCD, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
	})
}

func (c *char) initSkillInsert() {
	c.engine.Events().HPChange.Subscribe(c.SkillInsertListener)
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// do A2
	if c.info.Traces["102"] {
		c.engine.DispelStatus(target, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
	}

	// do E2
	if c.info.Eidolon >= 2 {
		if c.engine.HPRatio(target) < 0.5 {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   E2HealBoost,
				Source: c.id,
			})
		} else {
			c.engine.AddModifier(target, info.Modifier{
				Name:   E2Shield,
				Source: c.id,
			})
		}
	}

	// heal target
	c.engine.Heal(info.Heal{
		Key:     Skill,
		Targets: []key.TargetID{target},
		Source:  c.id,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_ATK: skillPer[c.info.SkillLevelIndex()],
		},
		HealValue: skillFlat[c.info.SkillLevelIndex()],
	})

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Skill,
		Target: c.id,
		Source: c.id,
		Amount: 30,
	})

	// add 1 stack of Abyss Flower if no Field active
	if !c.engine.HasModifier(c.id, Field) {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   AbyssFlower,
			Source: c.id,
		})
	}

	// something with inserts
	if IsInsert {
		IsInsert = false
	}
}

func (c *char) SkillInsertListener(e event.HPChange) {
	// bypass if hp change is positive
	if e.NewHP > e.OldHP {
		return
	}

	// bypass if on cooldown
	if c.engine.HasModifier(c.id, SkillInsertCD) {
		return
	}

	// bypass if CC'd or unable to act
	cond1 := c.engine.HasBehaviorFlag(c.id, model.BehaviorFlag_STAT_CTRL)
	cond2 := c.engine.HasBehaviorFlag(c.id, model.BehaviorFlag_DISABLE_ACTION)
	if cond1 || cond2 {
		return
	}

	trg := c.engine.Retarget(info.Retarget{
		Targets: c.engine.Enemies(),
		Filter: func(target key.TargetID) bool {
			// missing filters
			return true
		},
		Max: 1,
	})

	if c.engine.HPRatio(trg[0]) <= 0.5 {
		// 	// apply another mod that applies another mod...
	}
}
