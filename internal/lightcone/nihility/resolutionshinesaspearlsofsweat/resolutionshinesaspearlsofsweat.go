package resolutionshinesaspearlsofsweat

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	sweat    key.Modifier = "resolution-shines-as-pearls-of-sweat"
	Ensnared key.Modifier = "resolution-shines-as-pearls-of-sweat-ensnared"
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
	modifier.Register(sweat, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: applyEnsnared,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}

func applyEnsnared(mod *modifier.Instance, e event.HitStart) {

}
