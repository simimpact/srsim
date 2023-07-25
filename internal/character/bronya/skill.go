package bronya

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill = "bronya-skill"
)

func init() {
	modifier.Register(Skill, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// Try E1
	c.e1()

	// Dispel
	c.engine.DispelStatus(target, info.Dispel{
		Status: model.StatusType_STATUS_DEBUFF,
		Order:  model.DispelOrder_LAST_ADDED,
		Count:  1,
	})

	// Action forward
	if target != c.id {
		c.engine.SetGauge(info.ModifyAttribute{
			Key:    Skill,
			Target: target,
			Source: c.id,
			Amount: 0,
		})
	}

	buffDuration := 1

	// E6 Duration increase
	if c.info.Eidolon >= 6 {
		buffDuration = 2
	}

	// Damage increase
	c.engine.AddModifier(target, info.Modifier{
		Name:     Skill,
		Source:   c.id,
		Stats:    info.PropMap{prop.AllDamagePercent: skill[c.info.SkillLevelIndex()]},
		Duration: buffDuration,
	})

	// Try E2
	c.e2(target)

	// Add energy
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Skill,
		Target: c.id,
		Source: c.id,
		Amount: 30.0,
	})
}
