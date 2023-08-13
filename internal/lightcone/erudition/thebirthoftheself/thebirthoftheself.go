package thebirthoftheself

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Increases DMG dealt by the wearer's follow-up attacks by 24%.
// If the current HP of the target enemy is below or equal to 50%,
// increases DMG dealt by follow-up attacks by an extra 24%.

func init() {
	lightcone.Register(key.TheBirthoftheSelf, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
