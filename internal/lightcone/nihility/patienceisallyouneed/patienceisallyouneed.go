package patienceisallyouneed

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Increases DMG dealt by the wearer by 24%. After every attack launched by wearer,
// their SPD increases by 4.8%, stacking up to 3 times. If the wearer hits an enemy target
// that is not afflicted by Erode, there is a 100% base chance to inflict Erode to the target.
// Enemies afflicted with Erode are also considered to be Shocked and will receive
// Lightning DoT at the start of each turn equal to 60% of the wearer's ATK, lasting for 1 turn(s).

func init() {
	lightcone.Register(key.PatienceIsAllYouNeed, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
