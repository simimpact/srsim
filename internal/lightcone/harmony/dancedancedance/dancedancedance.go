package dancedancedance

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
	dance = "dance-dance-dance"
)

// When the wearer uses their Ultimate, all allies' actions are Advanced Forward by 16%

func init() {
	lightcone.Register(key.DanceDanceDance, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})
	modifier.Register(dance, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAction: advForwardOnUlt,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}

func advForwardOnUlt(mod *modifier.Instance, e event.ActionEnd) {

}
