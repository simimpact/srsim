package bronya

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill key.Modifier = "bronya-skill"
)

type skillState struct {
	bonusDamage float64
}

func init() {
	modifier.Register(Skill, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				mod.SetProperty(prop.AllDamagePercent, mod.State().(skillState).bonusDamage)
			},
		},
		Duration: 1,
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
		c.engine.SetGauge(target, 0)
	}

	buffDuration := 1

	// E6 Duration increase
	if c.info.Eidolon >= 6 {
		buffDuration = 2
	}

	// Damage increase
	c.engine.AddModifier(target, info.Modifier{
		Name:   Skill,
		Source: c.id,
		State: skillState{
			bonusDamage: skill[c.info.AbilityLevel.Skill-1],
		},
		Duration: buffDuration,
	})

	// Try E2
	c.e2(target)

	// Add energy
	c.engine.ModifyEnergy(c.id, 30.0)
}
