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
	advForwardAmt := 0.14 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   dance,
		Source: owner,
		State:  advForwardAmt,
	})
}

func advForwardOnUlt(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType != model.AttackType_ULT {
		return
	}
	// apply advance forward to all characters.
	for _, char := range mod.Engine().Characters() {
		mod.Engine().ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    dance,
			Target: char,
			Source: mod.Owner(),
			Amount: mod.State().(float64),
		})
	}
}
