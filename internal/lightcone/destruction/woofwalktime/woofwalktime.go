package woofwalktime

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	mod key.Modifier = "woof-walk-time"
)

// Increases the wearer's ATK by 10%, and increases their DMG to enemies
// afflicted with Burn or Bleed by 16%. This also applies to DoT.

// impl note : literally same as fermata. diff DoT check and initial buff

func init() {
	lightcone.Register(key.WoofWalkTime, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
