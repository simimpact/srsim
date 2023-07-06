package landauschoice

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Desc : The wearer is more likely to be attacked, and DMG taken is reduced by 18%.
// Apparent modifiers : Aggro and DmgTakenReduce
// DM Listeners : OnStack = aggroAddedRatio + allDamageReduce, OnStart = addModifier
// Conclusion : on Create, add singular modifier w/ all calcs. register mod at init.

func init() {
	lightcone.Register(key.LandausChoice, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
