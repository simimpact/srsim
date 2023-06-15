package bronya

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Technique key.Modifier = "bronya-technique"
)

func init() {
	modifier.Register(Technique, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				mod.SetProperty(prop.ATKPercent, 0.15)
			},
		},
		Duration: 2,
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	targets := c.engine.Characters()

	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name:   Technique,
			Source: c.id,
		})
	}

	c.engine.Events().CharacterAdded.Subscribe(func(e event.CharacterAddedEvent) {
		c.engine.AddModifier(e.Id, info.Modifier{
			Name:   Technique,
			Source: c.id,
		})
	})
}
