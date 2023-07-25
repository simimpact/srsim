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

type skillState struct {
	damageBoost float64
}

func init() {
	modifier.Register(Skill, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		MaxCount:   4,
		Listeners: modifier.Listeners{
			OnAdd:    skillOnAdd,
			OnPhase2: skillOnPhase2,
		},
		CountAddWhenStack: 1,
	})
}
func (c *char) Skill(target key.TargetID, state info.ActionState) {
	extraDamage := 0.0
	if c.info.Traces["102"] {
		extraDamage = 0.1
	}
	c.engine.AddModifier(target, info.Modifier{
		Name:   Skill,
		Source: c.id,
		State: skillState{
			damageBoost: extraDamage + skill[c.info.SkillLevelIndex()],
		},
	})
	c.drawTile()
	c.drawTile()
	if c.info.Eidolon >= 4 && c.engine.Rand().Float64() < 0.24 {
		c.engine.AddModifier(target, info.Modifier{
			Name:   Autarky,
			Source: c.id,
		})
	}
	if c.tiles[0] == 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   Talent,
			Source: c.id,
			Stats:  info.PropMap{prop.ATKPercent: talent[c.info.TalentLevelIndex()]},
		})
	}
	c.engine.InsertAction(c.id)
}
func skillOnAdd(mod *modifier.Instance) {
	state := mod.State().(skillState)
	mod.AddProperty(prop.ATKPercent, mod.Count()*(state.damageBoost))
}
func skillOnPhase2(mod *modifier.Instance) {
	mod.RemoveSelf()
}
