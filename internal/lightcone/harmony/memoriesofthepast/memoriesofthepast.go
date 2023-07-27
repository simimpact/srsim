package memoriesofthepast

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
	memories key.Modifier = "memories-of-the-past"
)

// Increases the wearer's Break Effect by 28%. When the wearer attacks,
// additionally regenerates 4 Energy. This effect can only be triggered 1 time per turn.

func init() {
	lightcone.Register(key.MemoriesofthePast, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})
	modifier.Register(memories, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: addEnergy,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}

// add energy on attack IF not on cd.
func addEnergy(mod *modifier.Instance, e event.AttackEnd) {

}
