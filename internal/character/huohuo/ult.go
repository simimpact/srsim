package huohuo

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	UltBuff key.Modifier = "huohuo-ult"
	Ult     key.Reason   = "huohuo-ult"
)

func init() {
	modifier.Register(UltBuff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		Duration:   2,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	for _, target := range c.engine.Characters() {
		if c.id == target {
			continue
		}
		c.engine.ModifyEnergyFixed(info.ModifyAttribute{
			Key:    Ult,
			Target: target,
			Source: c.id,
			Amount: c.engine.MaxEnergy(target) * ultEnergy[c.info.UltLevelIndex()],
		})
		c.engine.AddModifier(target, info.Modifier{
			Name:            UltBuff,
			Source:          c.id,
			TickImmediately: false,
			Stats:           info.PropMap{prop.ATKPercent: ultAttack[c.info.UltLevelIndex()]},
		})
	}
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Target: c.id,
		Source: c.id,
		Amount: 5,
		Key:    Ult,
	})
}
