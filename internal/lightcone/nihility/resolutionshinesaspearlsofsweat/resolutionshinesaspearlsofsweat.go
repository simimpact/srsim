package resolutionshinesaspearlsofsweat

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// When the wearer hits an enemy and if the hit enemy is not already Ensnared,
// then there is a 60% base chance to Ensnare the hit enemy.
// Ensnared enemies' DEF decreases by 12% for 1 turn(s).

func init() {
	lightcone.Register(key.ResolutionShinesAsPearlsofSweat, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
