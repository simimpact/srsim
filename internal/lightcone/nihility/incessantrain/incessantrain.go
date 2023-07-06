package incessantrain

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Desc : Increases the wearer's Effect Hit Rate by 24%.
// When the wearer deals DMG to an enemy that currently has 3 or more debuffs,
// increases the wearer's CRIT Rate by 12%. After the wearer uses their Basic ATK, Skill, or Ultimate,
// there is a 100% base chance to implant Aether Code on a random hit target that does not yet have it.
// Targets with Aether Code receive 12% increased DMG for 1 turn.

func init() {
	lightcone.Register(key.IncessantRain, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
