package blade

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Hellscape        = "blade-hellscape"
	HellscapeDmgBuff = "blade-hellscape-dmg-buff"
)

func init() {
	modifier.Register(Hellscape, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Duration: 3,
		Listeners: modifier.Listeners{
			OnRemove: func(mod *modifier.Instance) {
				mod.Engine().RemoveModifier(mod.Owner(), HellscapeDmgBuff)
				mod.Engine().RemoveModifier(mod.Owner(), E2)
			},
		},
	})
	modifier.Register(HellscapeDmgBuff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.ModifyHPByRatio(info.ModifyHPByRatio{
		Key:       Hellscape,
		Target:    c.id,
		Source:    c.id,
		Ratio:     -0.3,
		RatioType: model.ModifyHPRatioType_MAX_HP,
		Floor:     1,
	})

	statsPropMap := info.PropMap{prop.AllDamagePercent: skill[c.info.SkillLevelIndex()]}

	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Hellscape,
		Source: c.id,
	})

	// E2
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E2,
			Source: c.id,
			Stats:  info.PropMap{prop.CritChance: 0.15},
		})
	}

	c.engine.AddModifier(c.id, info.Modifier{
		Name:   HellscapeDmgBuff,
		Source: c.id,
		Stats:  statsPropMap,
	})

	c.engine.InsertAction(c.id)
}
