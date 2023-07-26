package blade

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Hellscape key.Modifier = "blade-hellscape"
)

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
		statsPropMap.Modify(prop.CritChance, 0.15)
	}

	// TODO: Duration Behaviour?
	c.engine.AddModifier(c.id, info.Modifier{
		Name:     Hellscape,
		Source:   c.id,
		Stats:    statsPropMap,
		Duration: 3,
	})

	c.engine.InsertAction(c.id)
}
