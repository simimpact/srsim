package solitaryhealing

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Increases the wearer's Break Effect by 20%. When the wearer uses their Ultimate,
// increases DoT dealt by the wearer by 24%, lasting for 2 turn(s).
// When a target enemy suffering from DoT imposed by the wearer is defeated,
// regenerates 4 Energy for the wearer.

func init() {
	lightcone.Register(key.SolitaryHealing, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
