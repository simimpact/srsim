package carvethemoonweavetheclouds

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	carveCheck key.Modifier = "carve-the-moon-weave-the-cloud"
	carveAtk   key.Modifier = "carve-the-moon-weave-the-cloud-atk"
	carveCDmg  key.Modifier = "carve-the-moon-weave-the-cloud-cdmg"
	carveERR   key.Modifier = "carve-the-moon-weave-the-cloud-err"
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
	modifier.Register(carveCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeDying: removeBuffs,
			OnPhase1:      applyTeamBuffRandomly,
		},
	})
	modifier.Register(carveAtk, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(carveCDmg, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(carveERR, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}

func removeBuffs(mod *modifier.Instance) {

}

func applyTeamBuffRandomly(mod *modifier.Instance) {

}
