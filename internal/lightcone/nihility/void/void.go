package void

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	ehrBuff key.Modifier = "void"
)

// At the start of the battle, the wearer's Effect Hit Rate increases by 20% for 3 turn(s).

func init() {
	lightcone.Register(key.Void, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(ehrBuff, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
