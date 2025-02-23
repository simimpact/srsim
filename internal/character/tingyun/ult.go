package tingyun

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult     = "tingyun-ult"
	UltBuff = "tingyun-ult-buff"
)

func init() {
	modifier.Register(UltBuff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		CanDispel:  true,
		Duration:   2,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	energyAmount := 50.0
	if c.info.Eidolon >= 6 {
		energyAmount += 10.0
	}
	c.engine.ModifyEnergyFixed(info.ModifyAttribute{
		Key:    Technique,
		Target: target,
		Source: c.id,
		Amount: energyAmount,
	})

	c.engine.AddModifier(target, info.Modifier{
		Name:            UltBuff,
		Source:          c.id,
		Stats:           info.PropMap{prop.AllDamagePercent: ult[c.info.UltLevelIndex()]},
		TickImmediately: true,
	})

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Ult,
		Target: c.id,
		Source: c.id,
		Amount: 5.0,
	})
}
