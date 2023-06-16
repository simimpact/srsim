package qingque

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill key.Modifier = "qingque-skill"
)

func init() {
	modifier.Register(Skill, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		MaxCount:   4,
		Listeners: modifier.Listeners{
			OnAdd: skillOnAdd,
		},
	})
}
func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(target, info.Modifier{
		Name:   Skill,
		Source: c.id,
	})
	c.engine.InsertAction(c.id)
}
func skillOnAdd(mod *modifier.ModifierInstance) {
	extraDamage := 0.0
	instance, err := mod.Engine().CharacterInstance(mod.Owner())
	if err != nil {
		// bad stuff idk how to deal with this
	}
	c := instance.(*char)
	mod.Owner()
	if c.info.Traces["1201102"] {
		extraDamage = 0.1
	}
	mod.AddProperty(prop.ATKPercent, mod.Count()*(extraDamage+skill[c.info.SkillLevelIndex()]))
}
