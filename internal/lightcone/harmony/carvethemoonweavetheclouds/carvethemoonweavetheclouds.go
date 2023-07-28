package carvethemoonweavetheclouds

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// At the start of the battle and whenever the wearer's turn begins,
// one of the following effects is applied randomly: All allies' ATK increases by 10%,
// all allies' CRIT DMG increases by 12%, or all allies' Energy Regeneration Rate
// increases by 6%. The applied effect cannot be identical to the last effect applied,
// and will replace the previous effect. The applied effect will be removed when the wearer
// has been knocked down. Effects of the similar type cannot be stacked.

func init() {
	lightcone.Register(key.CarvetheMoonWeavetheClouds, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
