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
	SkillEffect = "danhengimbibitorlunae-skill-effect" // MAvatar_DanHengIL_00_Skill02_CriticalDamage
)

type skillState struct {
	CritDMGBoost float64
}

// enhance attack,has 3 type,use 1/2/3 skill point

func (c *char) initSkill() {
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
	if c.attackLevel == 3 {
		return false
	}
	total := c.engine.SP()
	total += c.point
	total -= c.attackLevel
	return total > 0
}

// FIXME: enhance and attack is in 1 action, and won't trigger skill listener

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.attackLevel++
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
