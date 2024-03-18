package reforgedremembrance

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	rememberance = "reforged-rememberance"
)

// Increases the wearer's Effect Hit Rate by 40%. When the wearer deals DMG to an enemy
// inflicted with Wind Shear, Burn, Shock, or Bleed, each respectively grants 1 stack of Prophet,
// stacking up to 4 time(s). In a single battle, only 1 stack of Prophet can be granted for each
// type of DoT. Every stack of Prophet increases wearer's ATK by 5% and enables the DoT dealt
// to ignore 7.2% of the target's DEF.
func init() {
	lightcone.Register(key.ReforgedRemembrance, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	ehrAmt := 0.4 + 0.05*float64(lc.Imposition)
	modState := state{
		atkBuff: 0.05 + 0.01*float64(lc.Imposition),
		defShred: 0.72 + 0.07*float64(lc.Imposition),
	}
}
