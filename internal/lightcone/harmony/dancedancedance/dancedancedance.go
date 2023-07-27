package dancedancedance

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// When the wearer uses their Ultimate, all allies' actions are Advanced Forward by 16%

func init() {
	lightcone.Register(key.DanceDanceDance, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
