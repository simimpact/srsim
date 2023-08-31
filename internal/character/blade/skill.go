package blade

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Hellscape key.Modifier = "blade-hellscape"
	E2        key.Modifier = "blade-e2"
)

func init() {
	modifier.Register(Hellscape, modifier.Config{
		Stacking:   modifier.Replace,
		Duration:   3,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E2, modifier.Config{
		Stacking:   modifier.Replace,
		Duration:   3,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.ModifyHPByRatio(info.ModifyHPByRatio{
		Key:       key.Reason(Hellscape),
		Target:    c.id,
		Source:    c.id,
		Ratio:     -0.3,
		RatioType: model.ModifyHPRatioType_MAX_HP,
		Floor:     1,
	})

	statsPropMap := info.PropMap{prop.AllDamagePercent: skill[c.info.SkillLevelIndex()]}

	// E2
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E2,
			Source: c.id,
			Stats:  info.PropMap{prop.CritChance: 0.15},
		})
	}

	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Hellscape,
		Source: c.id,
		Stats:  statsPropMap,
	})

	c.engine.InsertAction(c.id)
}
