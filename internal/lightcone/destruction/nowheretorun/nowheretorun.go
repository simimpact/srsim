package nowheretorun

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Increases the wearer's ATK by 48%.
// Whenever the wearer defeats an enemy, they restore HP equal to 24% of their ATK.

func init() {
	lightcone.Register(key.NowheretoRun, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(key.NowheretoRun, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: onTriggerDeath,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}

func onTriggerDeath(mod *modifier.ModifierInstance, target key.TargetID) {

}
