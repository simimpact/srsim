package arrows

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Arrows key.Modifier = "arrows"
)

// At the beginning of the battle, increases the wearer's CRIT Rate
// by 12/15/18/21/24% for 3 turn(s).
func init() {
	lightcone.Register(key.Arrows, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(Arrows, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.09 + 0.03*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:     Arrows,
		Source:   owner,
		Duration: 3,
		Stats:    info.PropMap{prop.CritChance: amt},
	})
}
