package hanya

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1         = "hanya-e1"
	E1Cooldown = "hanya-e1-cooldown"
	E2         = "hanya-e2"
)

func init() {
	modifier.Register(E1, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnTriggerDeath: E1HanyaAdv,
		},
	})

	modifier.Register(E1Cooldown, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase2: E1HanyaCD,
		},
	})

	modifier.Register(E2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		Duration:   1,
	})
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E2,
			Source: c.id,
			Stats: info.PropMap{
				prop.SPDConvert: 0.2,
			},
		})
	}
}

func E1HanyaAdv(mod *modifier.Instance, target key.TargetID) {
	if !mod.Engine().HasModifier(mod.Source(), E1Cooldown) {
		mod.Engine().ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    E1,
			Source: mod.Owner(),
			Target: mod.Source(),
			Amount: -1500,
		})
	}
}

func E1HanyaCD(mod *modifier.Instance) {
	mod.RemoveSelf()
}
