package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SkillEffect  = "danhengimbibitorlunae-skill-effect"
	EnhanceLevel = "danhengimbibitorlunae-enhancelevel"
)

type skillState struct {
	CritDMGBoost float64
}

// enhance attack,has 3 type,use 1/2/3 skill point

func (c *char) initSkill() {
	modifier.Register(EnhanceLevel, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          3,
		StatusType:        model.StatusType_UNKNOWN_STATUS,
		CountAddWhenStack: 1,
	})
	modifier.Register(SkillEffect, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		MaxCount:   4,
		Listeners: modifier.Listeners{
			OnAdd: SkillOnAdd,
		},
		CountAddWhenStack: 1,
	})
}

func canUseSkill(engine engine.Engine, instance info.CharInstance) bool {
	c := instance.(*char)
	total := c.engine.SP()
	if c.engine.HasModifier(c.id, Point) {
		total += int(c.engine.ModifierStackCount(c.id, c.id, Point))
	}
	if c.engine.HasModifier(c.id, EnhanceLevel) {
		level := int(c.engine.ModifierStackCount(c.id, c.id, EnhanceLevel))
		if level == 3 {
			return false
		}
		total -= level
	}
	return total > 0
}

// FIXME: enhance and attack is in 1 action

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   EnhanceLevel,
		Source: c.id,
	})
	c.engine.InsertAction(c.id)
}

// add skill effect stack when attack
func (c *char) AddSkill() {
	outroarDuration := 1
	// if E4, the effect has 1 more turn and can be refreshed
	if c.info.Eidolon >= 4 {
		outroarDuration = 2
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:            SkillEffect,
		Source:          c.id,
		TickImmediately: true,
		Duration:        outroarDuration,
		State: skillState{
			CritDMGBoost: skill[c.info.SkillLevelIndex()],
		},
	})
}

func SkillOnAdd(mod *modifier.Instance) {
	state := mod.State().(skillState)
	mod.SetProperty(prop.CritDMG, mod.Count()*(state.CritDMGBoost))
}
