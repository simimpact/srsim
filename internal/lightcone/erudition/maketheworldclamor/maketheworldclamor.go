package maketheworldclamor

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
	clamor = "make-the-world-clamor"
)

// The wearer regenerates 20 Energy immediately upon entering battle,
// and increases Ultimate DMG by 32%.

func init() {
	lightcone.Register(key.MaketheWorldClamor, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
	modifier.Register(clamor, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: buffUltDamage,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.Events().BattleStart.Subscribe(func(e event.BattleStart) {
		engine.ModifyEnergy(info.ModifyAttribute{
			Key:    clamor,
			Target: owner,
			Source: owner,
			Amount: 20,
		})
	})
}

func buffUltDamage(mod *modifier.Instance, e event.HitStart) {

}
