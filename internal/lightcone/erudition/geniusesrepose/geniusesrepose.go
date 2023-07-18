package geniusesrepose

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	mod key.Modifier = "geniuses-repose"
)

// DESC : Increases the wearer's ATK by 16%.
// When the wearer defeats an enemy, the wearer's CRIT DMG increases by 24% for 3 turn(s).

func init() {
	lightcone.Register(key.GeniusesRepose, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
