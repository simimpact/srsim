package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SkillEffect             = "danhengimbibitorlunae-skill-effect"
	EnhanceLevel            = "danhengimbibitorlunae-enhancelevel"
	Skill        key.Reason = "danhengimbibitorlunae-skill"
)

type skillState struct {
	CritDMGBoost float64
}

// enhance attack,has 3 type,use 1/2/3 skill point
// it shouldn't trigger the action listener(like YuKong's trace), but I don't know how to do this

func (c *char) initSkill() {
	modifier.Register(EnhanceLevel, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          3,
		StatusType:        model.StatusType_UNKNOWN_STATUS,
		CountAddWhenStack: 1,
	})
	// if E4, the effect has 1 more turn and can be refreshed
	if c.info.Eidolon < 4 {
		modifier.Register(SkillEffect, modifier.Config{
			StatusType: model.StatusType_STATUS_BUFF,
			Stacking:   modifier.ReplaceBySource,
			MaxCount:   4,
			Listeners: modifier.Listeners{
				OnAdd:    SkillOnAdd,
				OnPhase2: SkillOnPhase2,
			},
			CountAddWhenStack: 1,
		})
	} else {
		modifier.Register(SkillEffect, modifier.Config{
			StatusType: model.StatusType_STATUS_BUFF,
			Stacking:   modifier.ReplaceBySource,
			MaxCount:   4,
			Duration:   2,
			Listeners: modifier.Listeners{
				OnAdd: SkillOnAdd,
			},
			CountAddWhenStack: 1,
		})
	}
}

// TODO this skill can be used when it has an effect and without skillpoint

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	if c.engine.HasModifier(c.id, Point) {
		c.engine.ExtendModifierCount(c.id, Point, -1)
		c.engine.ModifySP(info.ModifySP{
			Key:    Skill,
			Source: c.id,
			Amount: 1,
		})
	}
	c.engine.InsertAction(c.id)
}

// add skill effect stack when attack
func (c *char) AddSkill() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   SkillEffect,
		Source: c.id,
		State: skillState{
			CritDMGBoost: skill[c.info.SkillLevelIndex()],
		},
	})
}

func SkillOnAdd(mod *modifier.Instance) {
	state := mod.State().(skillState)
	mod.SetProperty(prop.CritDMG, mod.Count()*(state.CritDMGBoost))
}
func SkillOnPhase2(mod *modifier.Instance) {
	mod.RemoveSelf()
}
